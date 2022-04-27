package api

import (
	"fmt";
	"io/ioutil";
	"os"
)

func GetBearer() string {
	user := os.Getenv("USER")

	f, err := ioutil.ReadFile("/home/" + user + "/.discord-cli/TOKENSECRET.txt")

	if err != nil {
		fmt.Println("err")
	}

	return "Bearer " + string(f)
}

func GetEndpoint() string {
	return "https://discord.com/api/v8"
}