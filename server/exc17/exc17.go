package exc17

import (
	"log"
	"time"
	"unicode/utf8"
)

//Find the Characters Counterpart Char Code
//Create a function that takes a single character as an argument and
//returns the char code of its lowercased / uppercased counterpart.

func character_Counterpart(char string) (rune, int) {
	result, size := utf8.DecodeRuneInString(char)
	return result, size
}

//Time Around the World
//In this challenge, the goal is to calculate what time it is in two different cities.
//You're given a string cityA and the related string timestamp (time in cityA)
//with the date formatted in full U.S. notation, as in this example

func calculateTimeDifference(cityA, cityB, timestampA string) (time.Time, time.Time) {
	layout := "2006-01-02 15:04:05"
	timeinCityA, err := time.Parse(layout, timestampA)
	if err != nil {
		log.Println("Error in timestamp")
	}

	locationCityA, err := time.LoadLocation(cityA)
	if err != nil {
		log.Println("Error in timestamp")
	}
	timeinCityB := timeinCityA.In(locationCityA)
	return timeinCityA, timeinCityB
}
