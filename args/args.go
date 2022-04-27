package args

import (
	"fmt";
	"discord-cli/auth";
	"discord-cli/member";
	"os";
	"log";
	"encoding/json";
	"io/ioutil"
)

func Auth() {
	user := os.Getenv("USER")

	os.Mkdir("/home/" + user + "/.discord-cli", 0777)

	os.Create("/home/" + user + "/.discord-cli/TOKENSECRET.txt")

	ch := make(chan bool, 1)
	go auth.Host(&ch)

	<-ch

	os.Create("/home/" + user + "/.discord-cli/memberinfo.txt")

	fmt.Println("Succefully authenticated!")
}

func MemberStats() {
	info := member.GetMemberInfo()

	user := os.Getenv("USER")

	mar, _ := json.Marshal(info)

	err := ioutil.WriteFile("/home/" + user + "/.discord-cli/memberinfo.json", mar , 0666)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Username: " + info["username"] + "#" + info["discriminator"] + "\n" + "ID: " + info["id"])

	fmt.Println("Wrote output to " + "/home/" + user + "/.discord-cli/memberinfo.json")
}

func ServerListStats() {
	servers := member.GetAllServers()

	user := os.Getenv("USER")

	os.Create("/home/" + user + "/.discord-cli/serverlist.json")

	var serverString string

	for key, _ := range servers {
		if len(servers[key]) == 0 {
			break
		}
		serverString = serverString + servers[key]["name"].(string) + "\n"
	}

	marServers, _ := json.MarshalIndent(servers, "", "    ")

	err := ioutil.WriteFile("/home/" + user + "/.discord-cli/serverlist.json", marServers, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(serverString)
}