package main

import (
	"errors"
	"fmt"
)

/*
#################
###GLOBAL VARS###
#################
*/

//Help text
var helpText string = `############
###LiveDB###
############

[Database Traffic Simulator]

USAGE:

	./livedb <command>

COMMANDS:
	status		Provide status of 'customer' database
`

//modes defines what operation livedb will perform

/*
####################
###UTIL FUNCTIONS###
####################
*/

//Help Text
func Help() {
	fmt.Println(helpText)
}

//Clear the terminal
func Clear() {
	fmt.Print("\033[H\033[2J")
}

//CheckArgs validates arguments.
func CheckArgs(a []string) (func(), error) {
	if a[1] == "status" {
		return func() { Status() }, nil
	}
	return func() { Help() }, errors.New("We did not received a valid argument.")
}
