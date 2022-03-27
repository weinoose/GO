package src

import "fmt"

type Driver struct {
	Name   string
	Number int
	Wins   int
	WDC    int
}

func ArgumentStruct(driver Driver) {
	fmt.Println(driver.Name, driver.Number, driver.Wins, driver.WDC)
}

func Struct() {
	var SV5 Driver
	SV5.Name = "Sebastian Vettel"
	SV5.Number = 5
	SV5.Wins = 53
	SV5.WDC = 4

	fmt.Println(SV5.Name, SV5.Number, SV5.Wins, SV5.WDC)
	ArgumentStruct(SV5)
}
