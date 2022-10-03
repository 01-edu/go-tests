package solutions

import "fmt"

func DealAPackOfCards(deck []int) {
	for _, v := range deck {
		if v == 1 {
			fmt.Printf("Player 1: %v, %v, %v", deck[0], deck[1], deck[2])
			fmt.Println()
		} else if v == 3 {
			fmt.Printf("Player 2: %v, %v, %v", deck[3], deck[4], deck[5])
			fmt.Println()
		} else if v == 6 {
			fmt.Printf("Player 3: %v, %v, %v", deck[6], deck[7], deck[8])
			fmt.Println()
		} else if v == 9 {
			fmt.Printf("Player 4: %v, %v, %v", deck[9], deck[10], deck[11])
			fmt.Println()
		}
	}
}
