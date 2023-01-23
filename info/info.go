package info

import "fmt"

const mainText = "BMI Calculator"
const seperator = "-----------------"
const WeightPrompt = "Please enter the weight in (kg): "
const HeightPrompt = "Please enter the height in (m): "

func WelcomeText() {
	fmt.Println(mainText)
	fmt.Println(seperator)
}
