package member

import (
	"discord-cli/api";
	"net/http";
	"encoding/json"
)

func GetMemberInfo() map[string]string {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", api.GetEndpoint() + "/users/@me", nil)
	req.Header.Set("Authorization", api.GetBearer())

	res, _ := client.Do(req)

	var info map[string]string

	json.NewDecoder(res.Body).Decode(&info)

	return info
}

func GetAllServers() []map[string]interface{} {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", api.GetEndpoint() + "/users/@me/guilds", nil)
	req.Header.Set("Authorization", api.GetBearer())

	res, _ := client.Do(req)

	var servers []map[string]interface{}

	json.NewDecoder(res.Body).Decode(&servers)

	return servers
}