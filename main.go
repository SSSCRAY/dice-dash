package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// 🧱
	// 🎲
	// ❌

	score := 0

	fmt.Println("Get ready!")
	fmt.Println("Счет:", score)
	fmt.Println("")

	for {
		fmt.Println("------------------------")
		fmt.Println("Приближаюсь к стене")
		fmt.Println("🎲      🧱")
		fmt.Println("")

		fmt.Println("Прыгаю!")
		fmt.Println("      🎲")
		fmt.Println("      🧱")
		fmt.Println("")

		if rand.Intn(4) == 1 {
			fmt.Println("Я врезался в стену")
			fmt.Println("     💥🧱")
			fmt.Println("")
			break
		}

		fmt.Println("Я успешно перепрыгнул через стену")
		fmt.Println("      🧱  🎲")
		fmt.Println("")

		score += 5
		fmt.Println("Счет:", score)
		fmt.Println("------------------------")
		fmt.Println("")
	}

	fmt.Println("Game over!")
	fmt.Println("Ваш итог:", score)

}
