package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type yopur number in, just press ENTER when ready."

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)
}
func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	// display a welcome/instructions

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Guess the number game!")
	fmt.Println("----------------------")
	fmt.Println("")

	fmt.Println("Think of an number between 1-10", prompt)
	reader.ReadString('\n')
	// take them through the game
	fmt.Println("Multiply line by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("divide the result by the number you originally though", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer

	fmt.Println("The answer is", answer)

}
