package main

import (
	"io"
	"log"
	"os"
	"time"

	"api-eater/nhlapi"
)

func main() {

	// start time for benchmark
	now := time.Now()

	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening 'rosters.txt': %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlapi.GetAllTeams()
	if err != nil {
		log.Fatalf("Error getting all teams: %v", err)
	}

	for _, team := range teams {
		log.Println("-------------------------------")
		log.Printf("Name: %s", team.Name)

	}

	// stop time for benchmark
	log.Printf("took %v", time.Now().Sub(now).String())
}
