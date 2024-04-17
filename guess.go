// guess - игра, в которой игрок должен угадать случайное число.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println("Я выбрал случайное число между 1 и 100")
	fmt.Println("Сможешь его угадать?")

	reader := bufio.NewReader((os.Stdin))

	succes := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("У тебя осталось", 10-guesses, "попыток.")

		fmt.Print("Сделай предположение: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Сорян, Твое предположение МЕНЬШЕ.")
		} else if guess > target {
			fmt.Println("Сорян, Твое предположение БОЛЬШЕ.")
		} else {
			succes = true
			fmt.Println("Вау ты крут! Угадал!")
			break
		}
	}

	if !succes {
		fmt.Println("Извини, ты не угадал мое число. Число было -", target)
	}

}
