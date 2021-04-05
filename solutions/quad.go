package solutions

import "fmt"

func printQuadLines(x, y int, start, middle, end string) {
	if x < 1 || y < 1 {
		return
	}

	drawLine := func(x int, s string) {
		start, middle, end := s[0], s[1], s[2]
		if x >= 1 {
			fmt.Printf("%c", start)
		}
		if x > 2 {
			for i := 0; i < (x - 2); i++ {
				fmt.Printf("%c", middle)
			}
		}
		if x > 1 {
			fmt.Printf("%c", end)
		}
		fmt.Println()
	}

	if y >= 1 {
		drawLine(x, start)
	}
	if y > 2 {
		for i := 0; i < y-2; i++ {
			drawLine(x, middle)
		}
	}
	if y > 1 {
		drawLine(x, end)
	}
}
