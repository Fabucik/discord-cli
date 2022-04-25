package main

import (
	"fmt";
	"discord-cli/auth";
	"discord-cli/member";
	"os";
	"log";
	"encoding/json";
	"io/ioutil"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		os.Exit(func() int {
			return 0
		}())
	}

	switch args[0] {
	case "auth":
		ch := make(chan bool, 1)
		go auth.Host(&ch)

		<-ch

		user := os.Getenv("USER")

		err := os.Mkdir("/home/" + user + "/.discord-cli", 0777)
		if err != nil {
			
		}

		_, err = os.Create("/home/" + user + "/.discord-cli/memberinfo.txt")
		if err != nil {
			
		}

		fmt.Println("Succefully authenticated!")
	
	case "stats":
		r := member.GetMemberInfo()

		user := os.Getenv("USER")

		mar, _ := json.Marshal(r)

		err := ioutil.WriteFile("/home/" + user + "/.discord-cli/memberinfo.txt", mar , 0666)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Username: " + r["username"] + r["discriminator"] + "\n" + "ID: " + r["id"])

		fmt.Println("Wrote output to " + "/home/" + user + "/.discord-cli/memberinfo.txt")
		
	}
}