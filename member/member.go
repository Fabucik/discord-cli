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

	var r map[string]string

	json.NewDecoder(res.Body).Decode(&r)

	return r
}