package main

import (
	"discord-cli/args"
	"fmt"
	"os"
)

func doesTokenExist() bool {
	user := os.Getenv("USER")

	fi, _ := os.Open("/home/" + user + "/.discord-cli/TOKENSECRET.txt")
	si, _ := fi.Stat()

	if si.Size() == 0 {
		fmt.Println("You have to run \"discord-cli auth\" first")
		return false
	}

	return true
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) <= 0 {
		os.Exit(func() int {
			return 0
		}())
	}

	switch arguments[0] {
	case "auth":
		args.Auth()

	case "stats":
		if !doesTokenExist() {
			return
		}
		args.MemberStats()

	case "list":
		if !doesTokenExist() {
			return
		}
		args.ServerListStats()
	}
}
