package main

import (
	"io"
	"log"
	"os"
	"sync"
	"time"

	"api-eater/nhlapi"
)

func main() {

	// start time for benchmark
	now := time.Now()

	// Open rosters.txt file to write to
	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Error opening 'rosters.txt': %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	// Fetch all team names from API
	teams, err := nhlapi.GetAllTeams()
	if err != nil {
		log.Fatalf("Error getting all teams: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(len(teams))

	// unbuffered channel
	results := make(chan []nhlapi.Roster)

	// iterate every team to get roster
	for _, team := range teams {
		go func(team nhlapi.Team) {
			roster, err := nhlapi.GetRosters(team.ID)
			if err != nil {
				log.Fatalf("error getting roster: %v", err)
			}

			// send roster to results channel
			results <- roster

			wg.Done()
		}(team)

	}

	go func() {
		wg.Wait()
		close(results)
	}()

	display(results)

	// stop time for benchmark
	log.Printf("took %v", time.Now().Sub(now).String())
}

func display(results chan []nhlapi.Roster) {
	for r := range results {
		for _, ros := range r {
			log.Println("-------------------------------")
			log.Printf("ID: %v\n", ros.Person.ID)
			log.Printf("Name: %s\n", ros.Person.FullName)
			log.Printf("Position: %s\n", ros.Position.Code)
			log.Printf("Number: %s\n", ros.JerseyNumber)
		}
	}
}
