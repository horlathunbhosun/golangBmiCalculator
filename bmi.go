package main

import (
	"fmt"
	"github.com/horlathunbhosun/bmicalculator/info"
	"math/rand"
)

func main() {
	//Output information
	info.WelcomeText()

	weight, height := getUserMetrics()
	bmi := calculateBMI(weight, height)

	printBMI(bmi)
	//a, b := generateRandomNumber()
	//
	//sum := add(a, b)
	//
	//printNumber(sum)

}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}

//func returnReadline() (dat,string) {
//	dat, _ = reader.ReadString('\n')
//	 return
//}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func printNumber(num int) {
	fmt.Printf("The number is %v", num)
}

//func genarateRandomNumber() (int, int) {
//	randomNumber1 := rand.Intn(10)
//	randomNumber2 := rand.Intn(9)
//	return randomNumber1, randomNumber2
//}

func generateRandomNumber() (r1 int, r2 int) {
	r1 = rand.Intn(10)
	r2 = rand.Intn(9)
	return
}
