package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Rein is going to Review the Time pkg")

	presentTime := time.Now()
	// time format is very specific and fixed
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	fmt.Println(time.Date(2022, time.April, 0, 0, 0, 0, 0, time.UTC).Format("01-02-2006 Monday"))
}
