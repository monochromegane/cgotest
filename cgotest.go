package cgotest

import (
	"fmt"
	"os/user"
)

func GetHome() string {
	var homeDir string
	usr, err := user.Current()
	if err == nil {
		fmt.Println("user.Current() is success.")
		homeDir = usr.HomeDir
	} else {
		fmt.Printf("user.Current() is failure. %v\n", err)
		homeDir = "nothing"
	}
	return homeDir
}
