package nhlapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Roster struct to store rosters
type Roster struct {
	Person struct {
		ID       int    `json:"id"`
		FullName string `json:"fullName"`
		Link     string `json:"link"`
	} `json:"person"`
	JerseyNumber string `json:"jerseyNumber"`
	Position     struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"position"`
	Link string `json:"link"`
}

type nhlRosterResponse struct {
	Rosters []Roster `json:"roster"`
}

// GetRosters will get the roster given a team id
func GetRosters(teamID int) ([]Roster, error) {

	res, err := http.Get(fmt.Sprintf("%s/teams/%d/roster", baseURL, teamID))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response nhlRosterResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Rosters, err
}
