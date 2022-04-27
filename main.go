package main

import (
	"os";
	"discord-cli/args"
)

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
		args.MemberStats()

	case "list":
		args.ServerListStats()
	}
}