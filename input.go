package main

import (
	"bufio"
	"fmt"
	"github.com/horlathunbhosun/bmicalculator/info"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {

	weight = getUserInput(info.WeightPrompt)
	height = getUserInput(info.HeightPrompt)

	//fmt.Print(info.WeightPrompt)
	//weightInput, _ := reader.ReadString('\n')
	//weightInput = strings.Replace(weightInput, "\n", "", -1)
	//weight, _ = strconv.ParseFloat(weightInput, 64)
	//
	//fmt.Print(info.HeightPrompt)
	//heightInput, _ := reader.ReadString('\n')
	//heightInput = strings.Replace(heightInput, "\n", "", -1)
	//height, _ = strconv.ParseFloat(heightInput, 64)

	return
}

func getUserInput(promptText string) (value float64) {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	value, _ = strconv.ParseFloat(userInput, 64)
	return
}
