package src

import (
	"fmt"
	"math/rand"
)

func ConditionalsPractice() {
	fmt.Print("Sayı tahmin yarışmasına hoş geldin. Kaç tane hakkının olmasını istersin: ")
	var owe int
	fmt.Scanln(&owe)
	number := rand.Intn(100)

	for i := 0; i <= owe; i++ {
		fmt.Print("Tahminin nedir?: ")
		var guess int
		fmt.Scanln(&guess)

		if owe > 0 {
			if guess < number {
				fmt.Print("Daha büyük bir tahmin yap.")
				owe = owe - 1
			} else if guess > number {
				fmt.Print("Daha küçük bir tahmin yap.")
				owe = owe - 1
			} else {
				fmt.Println("! BİLDİN !")
				i = owe
			}
		} else if owe == 0 {
			fmt.Println("! BİLEMEDİN !")
		}
	}
}
