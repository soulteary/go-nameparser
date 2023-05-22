package bridge

import (
	"encoding/json"
	"fmt"
)

type HumanName struct {
	Text   string `json:"text"`
	Detail struct {
		Title    string `json:"title"`
		First    string `json:"first"`
		Middle   string `json:"middle"`
		Last     string `json:"last"`
		Suffix   string `json:"suffix"`
		Nickname string `json:"nickname"`
	} `json:"detail"`
}

func Parse(input string) (ret HumanName, err error) {
	var name HumanName
	err = json.Unmarshal([]byte(Convert(input)), &name)
	if err != nil {
		return ret, fmt.Errorf("Parsing JSON failed: %v", err)
	}
	return name, nil
}
