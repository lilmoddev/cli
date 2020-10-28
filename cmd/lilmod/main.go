package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"cli.lilmod.dev/clipb"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "lilmod",
		Commands: []*cli.Command{
			{
				Name: "login",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "registry",
						Aliases: []string{"r"},
						Usage:   "The registry URL to verify your login credentials and retrieve your GOPROXY URL",
					},
					&cli.BoolFlag{
						Name:  "token-stdin",
						Usage: "Retrieve the token credential from stdin",
					},
				},
				Action: func(c *cli.Context) error {
					if err := login(c); err != nil {
						return cli.Exit(err.Error(), 1)
					}
					return nil
				},
			},
			{
				Name: "logout",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "registry", // to remove from .netrc
						Value: "https://dev.lilmod.dev",
					},
				},
				Action: func(c *cli.Context) error {
					cmd := exec.CommandContext(c.Context, "go", "env", "-w", "GOPROXY=proxy.golang.org,direct")
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					err := cmd.Run()
					if err != nil {
						return fmt.Errorf("could not rewrite GOPROXY: %w", err)
					}
					// TODO: remove from .netrc
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}

func login(c *cli.Context) error {
	registry := c.String("registry")
	var token string
	var err error
	if c.Bool("token-stdin") {
		_, err = fmt.Scanln(&token)
		if err != nil {
			return fmt.Errorf("could not scan token from stdin: %w", err)
		}
	} else {
		prompt := promptui.Prompt{
			Label: "Token",
			Validate: func(s string) error {
				if strings.Contains(s, " ") {
					return fmt.Errorf("must not contain spaces")
				}
				return nil
			},
		}
		token, err = prompt.Run()
		if err != nil {
			return fmt.Errorf("could not get token from prompt: %w", err)
		}
	}
	lilmodClient := clipb.NewAPIProtobufClient(registry, http.DefaultClient)
	resp, err := lilmodClient.Proxy(c.Context, &clipb.ProxyRequest{ProxyToken: token})
	if err != nil {
		// TODO: nicer message when invalid auth
		return fmt.Errorf("error authenticating: %w", err)
	}
	proxyURL := resp.GetURL()
	url, err := url.Parse(proxyURL)
	if err != nil {
		return fmt.Errorf("error parsing %q: %w", proxyURL, err)
	}
	writeNETRC(url.Hostname(), "lilmod", token, "")
	// TODO: integrate with existing GOPROXY/GONOSUMDB instead of overriding everything
	args := []string{"env", "-w", fmt.Sprintf("GOPROXY=%s", proxyURL)}
	if privatePaths := resp.GetPrivatePaths(); len(privatePaths) > 0 {
		args = append(args, fmt.Sprintf("GONOSUMDB=%s", strings.Join(privatePaths, ",")))
	}
	cmd := exec.CommandContext(c.Context, "go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("could not persist GOPROXY: %w", err)
	}
	return nil
}
