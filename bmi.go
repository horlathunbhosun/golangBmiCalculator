package main

import (
	"fmt"
	"github.com/horlathunbhosun/bmicalculator/info"
	"strconv"
	"strings"
)

func main() {
	//Output information
	fmt.Println(info.MainText)
	fmt.Println(info.Seperator)
	//Prompt for user input(weight + height)

	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')
	//Save that user input in variables
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	//calculate the BMI (weight/ (height * height))
	bmi := weight / (height * height)

	fmt.Printf("The BMI is %.2f", bmi)
	//Output the calculated BMI

}
