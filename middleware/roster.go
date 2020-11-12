package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sabrestats/models"
)

const SabresID = "4416d559-0f24-11e2-8525-18a905767e44"
const SportsRadarUrl = "http://api.sportradar.us/nhl/trial/v7/en" // hierarchy.json?api_key="

//func CreateLines() []models.Player {
//    team := getTeam()
//    players := team.Players
//    return nil
//}

func assembleGET(feed, target string) string {
	url := fmt.Sprintf("%s/%s/%s/%s.json?api_key=%s", SportsRadarUrl, feed, SabresID, target, os.Getenv("SPORTRADAR_NHL_EICHELS"))
	return url
}

/**********************************************************************************************************************/
// Below functions are calls to the SportsRadar API and are limited per month to 1,000
func UpdateRoster() *models.Team {
	resp, err := http.Get(assembleGET("teams", "profile"))
	if err != nil {
		log.Println(fmt.Sprintf("error getting roster: %+v", err))
		panic(err)
	}
	defer resp.Body.Close()

	log.Printf("Roster response status: %v", resp.Status)

	team := new(models.Team)
	err = json.NewDecoder(resp.Body).Decode(team)
	if err != nil {
		log.Printf("Unable to retrieve roster: %+v\n", err)
	}

	log.Println(fmt.Sprintf("%+v", team))
	log.Println(team.Players)
	return team
}

func GetSeasonalStats() *models.Team {
	resp, err := http.Get(assembleGET("seasons/2019/REG/teams", "statistics"))
	if err != nil {
		log.Println(fmt.Sprintf("unable to update seasonal statistics: %+v", err))
	}
	defer resp.Body.Close()

	log.Println(fmt.Sprintf("Updating stats"))

	teamStats := new(models.Team)
	err = json.NewDecoder(resp.Body).Decode(teamStats)
	if err != nil {
		log.Println(fmt.Sprintf("unable to decode team stats: %+v", err))
	}

	return teamStats
}

/******************** No Depth charts until the start of the season ***************************************************/

func getDepthChart() map[string]interface{} {
	resp, err := http.Get(assembleGET("teams", "depth_chart"))
	if err != nil {
		log.Println(fmt.Sprintf("error getting roster: %+v", err))
		panic(err)
	}

	result := make(map[string]interface{})

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Could not unmarshal depth chart body: %+v", err)
	}

	depth := result["positions"].(map[string]interface{})
	log.Println(fmt.Sprintf("depth: %+v", depth))

	// returns a map of [positions]position
	return depth
}

func ForwardsDepth() []models.Player {
	// Get depth chart
	// Assemble each line according to depth chart
	// Remove extra players
	depth := getDepthChart()
	//var lineup []models.Player

	LWGroup := depth["LW"].(map[string]interface{})
	LW := LWGroup["players"].([]interface{})
	log.Println(fmt.Sprintf("LW: %+v", LW))

	//for i := 0; i < 4; i++ {
	//    lineup = append(lineup, depth["LW"].(models.Player))
	//    lineup = append(lineup, depth["C"].(models.Player))
	//    lineup = append(lineup, depth["RW"].(models.Player))
	//}
	//
	//return lineup
	return nil
}
