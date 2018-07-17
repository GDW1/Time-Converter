package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

var goLoc string /*location recognized by go*/

//function that takes the inputed string and cchecks if any time zones are assosiated with it
func stringToLoc(locSt string) *time.Location {
	strings.ToLower(locSt) //makes the string lowercase
	//this block corrects and spaces by replacing them with a underscore
	if strings.Contains(locSt, " ") {
		locSt = strings.Replace(locSt, " ", "_", -1)
	}
	//This itterates hrough the entire array of possible locations that are assosiated with the time library
	for _, v := range totalZones {
		//checks if the string the range is on contains the inputed string
		if strings.Contains(strings.ToLower(v), locSt) {
			//if so, load the strang as the loation and return that as the value
			loc, _ := time.LoadLocation(v)
			goLoc = v
			return loc
		} else if locSt == "san fransisco" || locSt == "san jose" {
			loc, _ := time.LoadLocation("America/Los_Angeles")
			goLoc = "America/Los_Angeles"
			return loc
			//runs all time zones
		} else if locSt == "developer_mode" || locSt == "developer" {
			allTimeZones(locSt)
			return nil
		}
	}
	//otherwise, return AS NIL to be proccessed by the main function
	fmt.Println("That city is not in the index. Please select a major city near you")
	return nil

}

//the function that is called on first
func main() {
	fmt.Println("Welcome to the timezone converter:")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your nearest major city")
	var text string
	fmt.Print("-> ")
	//reads for inputed values and waits for newline(\n) or the enter key to be inputed
	text, _ = reader.ReadString('\n')
	//takes out the newline as a value to make sure that it can be processed by the stringToLoc function
	text = strings.Replace(text, "\n", "", -1)
	//makes sure that the returned location has a value
	if stringToLoc(text) != nil {
		//sets the variable now to the time in the current location
		now := time.Now().In(stringToLoc(text))
		fmt.Printf("\n")
		//prints is using the fmt.Printf command with standard Format hh:mm:ss
		fmt.Printf("The current time for %v (%v) is %v\n", text, goLoc, now.Format("15:04:05"))

	} else {
		if text == "developer mode" || text == "developer" {
		} else {
			reader := bufio.NewReader(os.Stdin)
			//f the location is returned null the program will ask if the user would like to try again
			fmt.Print("To try again type y\n")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			//forces recursion if the user wants to try again
			if text == "y" {
				main()
			}
		}
	}
}

//developer function
func allTimeZones(text string) {
	var byTimeWString = make(map[int]string)
	//var byTime []int
	//again itterates through the entire array
	for i, v := range totalZones {
		//time.Sleep(100 * time.Millisecond)
		//loads the value of each location
		loc, _ := time.LoadLocation(v)
		now := time.Now().In(loc)
		fmt.Println(i, v, now.Format("Jan,3 15:06:03")) //prints the number in the array that the value is assosiated to then the value and then the time
		//fmt.Println()                                   //spacing
		hour := time.Now().In(loc).Hour()
		//byTime[i] = hour
		byTimeWString[hour] = v
	}
	var keys []int
	for k := range byTimeWString {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", byTimeWString[k])
	}
	//fmt.Println(byTimeWString)
}
