package main

import (
	"log"
	"os"
)

func main() {
	//Clear the terminal
	Clear()

	//Main logic
	if len(os.Args) > 1 {
		//Check arg
		run, err := CheckArgs(os.Args)
		if err != nil {
			log.Fatalf("We countered the following error. %s", err)
			run()
		}
	} else {
		Help()
	}
}
