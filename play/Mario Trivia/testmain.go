package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type NinStats interface {
	parse()
}

type ConsoleStats struct {
	name         string
	yearAppeared uint
}

func (cs ConsoleStats) parse() {
	return
}

// Takes an appropriate slice of strings and uses it to create a new NinStats
func makeNinStats(statsStringSlice []string) NinStats {

	// Convert strings to other types as necessary
	name := statsStringSlice[0]
	yearAppeared, _ := strconv.ParseUint(statsStringSlice[1], 10, 16)

	newNinStats := ConsoleStats{
		name,
		uint(yearAppeared),
	}
	return newNinStats
}

func main() {

	gameInfo := loadData()

	for again := "yes"; again != "no"; {
		askTrivia(gameInfo)
		fmt.Println("Would you like another question?")
		fmt.Scanf("%s\n", &again)
		again = strings.ToLower(again)
		for again != "yes" && again != "no" && again != "y" && again != "n" {
			fmt.Println("Would you like another question? Please enter \"yes\" or \"no\".")
			fmt.Scanln(&again)
		}
		if again == "n" {
			again = "no"
		}
	}
}

func loadData() []NinStats {
	// Load the array of data
	// Get data from the file
	fileStream, err := ioutil.ReadFile("MarioData.MDF")
	if err != nil {
		return nil
	}
	// Convert to a string
	fileString := string(fileStream)
	// Break it into slices by line
	lineSlice := strings.Split(fileString, "\n")

	var gameInfo []NinStats

	for _, i := range lineSlice {
		// Convert each slice of strings into a NinStats type and append to gameInfo slice
		gameInfo = append(gameInfo, makeNinStats(strings.Split(i, ",")))
	}

	return gameInfo
}

func askTrivia(infoArray []NinStats) {
	var userAnswer uint
	// var discardThis string
	// Pose a Nintendo trivia question based on the array of data presented
	// Seed random number gen
	seedTime := time.Now()
	rand.Seed(int64(seedTime.Nanosecond()))
	// Choose an element to base the question on
	elementNumber := rand.Intn(len(infoArray))
	// Choose which part of structure to ask a question about
	// ^^ For now just ask for the year
	// Pose the question that has been chosen
	fmt.Println("In what year was", infoArray[elementNumber].name, "released?")
	fmt.Scanln(&userAnswer)
	for userAnswer < 1980 || userAnswer > 2015 {
		fmt.Println("Please enter a year between 1980 and 2015.")
		fmt.Scanln(&userAnswer)
	}
	if userAnswer == infoArray[elementNumber].yearAppeared {
		fmt.Println("Correct! The", infoArray[elementNumber].name, "was released in", infoArray[elementNumber].yearAppeared)
	} else {
		fmt.Println("Wrong! The", infoArray[elementNumber].name, "was released in", infoArray[elementNumber].yearAppeared)
	}
}
