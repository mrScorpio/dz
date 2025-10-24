package main

import (
	concurrency "github.com/mrScorpio/dz/1-concurrency"
	randomApi "github.com/mrScorpio/dz/2-random-api"
)

func main() {
	concurrency.CreateRoutines()
	randomApi.CrWebServ()
}
