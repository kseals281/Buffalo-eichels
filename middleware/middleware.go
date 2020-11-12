package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sabrestats/models"
)

// initSocketConnectionPool initializes a Unix socket connection pool for
// a Cloud SQL instance of SQL Server.
func initSocketConnectionPool() (*sql.DB, error) {
	var (
		dbUser                 = os.Getenv("DB_USER")                  // e.g. 'my-db-user'
		dbPwd                  = os.Getenv("DB_PASS")                  // e.g. 'my-db-password'
		instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
		dbName                 = os.Getenv("DB_NAME")                  // e.g. 'my-database'
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return dbPool, nil
}

func Forwards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	eichel := []models.Player{
		{
			FullName:        "Taylor Hall",
			Height:          73,
			Weight:          204,
			PrimaryPosition: "LW",
			JerseyNumber:    4,
			Statistics: &models.Statistics{
				Total: &models.TotalStatistics{
					GamesPlayed: 82,
					Goals:       33,
					Assists:     50,
					Points:      83,
				},
			},
		}, {
			FullName:        "Jack Eichel",
			Height:          74,
			Weight:          203,
			PrimaryPosition: "C",
			JerseyNumber:    9,
			Statistics: &models.Statistics{
				Total: &models.TotalStatistics{
					GamesPlayed: 82,
					Goals:       35,
					Assists:     50,
					Points:      85,
				},
			},
		}, {
			FullName:        "Sam Reinhart",
			Height:          71,
			Weight:          196,
			PrimaryPosition: "RW",
			JerseyNumber:    23,
			Statistics: &models.Statistics{
				Total: &models.TotalStatistics{
					GamesPlayed: 82,
					Goals:       24,
					Assists:     29,
					Points:      53,
				},
			},
		}, {
			FullName:        "Taylor Hall",
			Height:          73,
			Weight:          204,
			PrimaryPosition: "LW",
			JerseyNumber:    4,
			Statistics: &models.Statistics{
				Total: &models.TotalStatistics{
					GamesPlayed: 82,
					Goals:       33,
					Assists:     50,
					Points:      83,
				},
			},
		}, {
			FullName:        "Jack Eichel",
			Height:          74,
			Weight:          203,
			PrimaryPosition: "C",
			JerseyNumber:    9,
			Statistics: &models.Statistics{
				Total: &models.TotalStatistics{
					GamesPlayed: 82,
					Goals:       35,
					Assists:     50,
					Points:      85,
				},
			},
		}, {
			FullName:        "Sam Reinhart",
			Height:          71,
			Weight:          196,
			PrimaryPosition: "RW",
			JerseyNumber:    23,
			Statistics: &models.Statistics{
				Total: &models.TotalStatistics{
					GamesPlayed: 82,
					Goals:       24,
					Assists:     29,
					Points:      53,
				},
			},
		},
	}
	json.NewEncoder(w).Encode(eichel)
}

func Defencemen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	dahlin := []models.Player{
		{
			FullName:        "Rasmus Dahlin",
			Height:          73,
			Weight:          176,
			PrimaryPosition: "D",
			JerseyNumber:    26,
			Statistics: &models.Statistics{
				Total: &models.TotalStatistics{
					GamesPlayed: 82,
					Goals:       13,
					Assists:     42,
					Points:      55,
				},
			},
		}, {
			FullName:        "Henri Jokiharju",
			Height:          70,
			Weight:          180,
			PrimaryPosition: "D",
			JerseyNumber:    10,
			Statistics: &models.Statistics{
				Total: &models.TotalStatistics{
					GamesPlayed: 82,
					Goals:       9,
					Assists:     12,
					Points:      21,
				},
			},
		},
	}
	json.NewEncoder(w).Encode(dahlin)
}

//func Goalies(w http.ResponseWriter, r *http.Request) {
//    w.Header().Set("Access-Control-Allow-Origin", "*")
//    miller := []models.Goalie{
//        {
//            FullName:        "Linus Ullmark",
//            Height:          "6'6\"",
//            Weight:          210,
//            PrimaryPosition: "G",
//            JerseyNumber:    35,
//            Games:           34,
//            Started:         34,
//            Wins:            17,
//            Losses:          14,
//            Overtime:        3,
//            SavePct:         .915,
//            GoalsAgainst:    2.69,
//        },
//    }
//    json.NewEncoder(w).Encode(miller)
//}
