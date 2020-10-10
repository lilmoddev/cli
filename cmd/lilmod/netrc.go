package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bgentry/go-netrc/netrc"
)

// writeNETRC writes or updates a netrc file
func writeNETRC(host, user, password, path string) error {
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("could not get homedir: %w", err)
		}
		path = filepath.Join(home, ".netrc")
	}
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("could not open netrc file: %w", err)
	}
	defer file.Close()
	creds, err := netrc.Parse(file)
	if err != nil {
		return fmt.Errorf("could not parse netrc file %q: %w", path, err)
	}
	m := creds.FindMachine(host)
	if m == nil {
		m = creds.NewMachine(host, user, password, "")
	} else {
		m.UpdateLogin(user)
		m.UpdatePassword(password)
	}
	bts, err := creds.MarshalText()
	if err != nil {
		return fmt.Errorf("could not marshal netrc: %w", err)
	}
	err = ioutil.WriteFile(path, bts, 0660)
	if err != nil {
		return fmt.Errorf("could not write netrc to file: %w", err)
	}
	return nil
}
