package main

import (
	"fmt"
)

const (
	IMCThinnessLimit   = 18.5
	IMCNormalLimit     = 25.0
	IMCOverweightLimit = 30.0
)

func main() {

	// weight, height := 70.5, 1.75

	var (
		weight float64
		height float64
		name   string
	)

	fmt.Print("Enter your name : ")
	fmt.Scan(&name)

	fmt.Print("Enter your weight (kg) : ")
	fmt.Scan(&weight)

	fmt.Print("Enter your height (m) : ")
	fmt.Scan(&height)

	fmt.Println("Name :", name)

	var IMC float64 = weight / (height * height)
	fmt.Printf("IMC : %.2f\n", IMC)
	var category string = findIMCCategory(IMC)
	fmt.Println("Category :", category)
}

func findIMCCategory(IMC float64) string {
	if IMC < IMCThinnessLimit {
		return "Thin"
	} else if IMC < IMCNormalLimit {
		return "Normal"
	} else if IMC < IMCOverweightLimit {
		return "Overweight"
	} else {
		return "Obese"
	}
}
