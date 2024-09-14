package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Code Controllers ")
	fmt.Println("If Else")

	var result string

	loginCount := 18

	if loginCount < 10 {
		result = "Less than 10"
	} else if loginCount < 20 {
		result = "Less than 20 bro..."

	} else {
		result = "You just leave it Da!"

	}
	fmt.Println(result)

	//Defin on the go
	if num := 4; num < 10 {
		result = "Less bro"

	} else {

		result = "You just leave it Da!"
	}

	fmt.Println("Switch Case")

	// rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	// Switch case breaks by default

	switch diceNumber {
	case 1:
		fmt.Println(diceNumber)
		fallthrough
	case 2:
		fmt.Println(diceNumber)
	case 3:
		fmt.Println(diceNumber)
		fallthrough
	case 4:
		fmt.Println(diceNumber)
	case 5:
		fmt.Println(diceNumber)
	case 6:
		fmt.Println(diceNumber)
	}

	fmt.Println("Switch Case")

	days := []string{"Sunday", "MoonDay", "TuesDays", "WednesDay", "ReinDay"}

	fmt.Println(days)

	for day := 0; day < len(days); day++ {
		fmt.Println("Day ", days[day])
	}

	for idx := range days {
		fmt.Println("Index ", idx)
	}

	for idx, v := range days {
		fmt.Println("Day ", v, "with index", idx)
	}

	rougue := 5

	for rougue < 15 {
		if rougue == 7 {
			rougue++
			continue
		}
		if rougue == 13 {
			goto moon
		}

		if rougue == 14 {
			break
		}
		fmt.Println(" value ", rougue)
		rougue++ // there is no ++rougue
	}

	//Go to this section directly.
moon:
	fmt.Println("Moon Happened..")

}
