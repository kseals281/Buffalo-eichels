package models

type Team struct {
	ID      string   `json:"id"`
	Name    string   `json:"name,omitempty"`
	Market  string   `json:"market,omitempty"`
	Alias   string   `json:"alias,omitempty"`
	Venue   Arena    `json:"venue,omitempty"`
	Players []Player `json:"players,omitempty"`
}

type Arena struct {
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	Capacity int    `json:"capacity,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	Country  string `json:"country,omitempty"`
	Timezone string `json:"time_zone,omitempty"`
}

type SeasonToDate struct {
	OwnRecord Record   `json:"own_record"`
	Opponents Record   `json:"opponents"`
	Players   []Player `json:"players"`
}

type Player struct {
	ID              string      `json:"id"`
	FullName        string      `json:"full_name,omitempty"`
	First           string      `json:"first_name,omitempty"`
	Last            string      `json:"last_name,omitempty"`
	Abbr            string      `json:"abbr_name,omitempty"`
	Height          int         `json:"height,omitempty"`
	Weight          int         `json:"weight,omitempty"`
	Position        string      `json:"position,omitempty"`
	PrimaryPosition string      `json:"primary_position,omitempty"`
	JerseyNumber    int         `json:"jersey_number,omitempty"`
	Exp             int         `json:"experience,omitempty"`
	Statistics      *Statistics `json:"statistics,omitempty"`
	//Goaltending  GoalieStats `json:"goaltending"`
}

type Record struct {
	Statistics Statistics `json:"statistics"`
}

type Statistics struct {
	Total *TotalStatistics `json:"total"`
}

type TotalStatistics struct {
	GamesPlayed int `json:"games_played"`
	Goals       int `json:"goals"`
	Assists     int `json:"assists"`
	Points      int `json:"points"`
}

//type GoalieStats struct {
//    Games           int     `json:"games_played"`
//    Wins            int     `json:"wins"`
//    Losses          int     `json:"losses"`
//    OTL             int     `json:"overtime_losses"`
//    SavesPct        float64 `json:"saves_pct"`
//    GoalsAgainstAvg float64 `json:"avg_goals_against"`
//}
