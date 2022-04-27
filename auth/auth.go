package auth

import (
	"net/http";
	"github.com/joho/godotenv";
	"os";
	"net/url";
	"log";
	"encoding/json";
	"fmt";
	"io/ioutil"
)

func Host(ch *chan bool) {
	godotenv.Load()
	clientId := os.Getenv("CLIENTID")

	authUrl := "https://discord.com/api/oauth2/authorize?client_id=" + clientId + "&redirect_uri=http%3A%2F%2F127.0.0.1%3A3369&response_type=code&scope=identify%20guilds%20guilds.members.read"

	fmt.Println("Authorize this app here: " + authUrl)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		GetToken(w, r, ch)
 	})
	http.ListenAndServe("127.0.0.1:3369", nil)
}

func GetToken(w http.ResponseWriter, request *http.Request, ch *chan bool) {
	tokenURL := "https://discord.com/api/v8/oauth2/token"

	godotenv.Load()
	clientId := os.Getenv("CLIENTID")
	clientSecret := os.Getenv("CLIENTSECRET")

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://127.0.0.1:3369")
	data.Set("code", request.URL.Query().Get("code"))

	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
        log.Fatalln(err)
    }

	decoder := json.NewDecoder(resp.Body)
	var r map[string]string
	err = decoder.Decode(&r)

	user := os.Getenv("USER")

	err = ioutil.WriteFile("/home/" + user + "/.discord-cli/TOKENSECRET.txt", []byte(r["access_token"]), 0777)
	if err != nil {
        log.Fatalln(err)
    }

	*ch <- true
}