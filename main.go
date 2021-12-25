package main

import "fmt"

func main() {

	SikuBawah(0, 10)
	Chess(8)
	// SikuBawahKi(0, 5)
	segitiga(6)
	Prima(50)

}

func segitiga(a int) {
	for x := 1; x <= a; x++ {
		for y := a; y >= x; y-- {
			fmt.Print(" ")
		}
		for z := 1; z <= x; z++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func Prima(a int) {
	for bil := 1; bil < a; bil++ {
		i := 0
		for bag := 1; bag < a; bag++ {
			if bil%bag == 0 {
				i++
			}
		}
		if (i == 2) && (bil != 1) {
			fmt.Println(bil)
		}
	}
}

func SikuBawah(a, b int) {
	for x := a; x <= b; x++ {
		for y := 0; y < x; y++ {
			fmt.Printf("#")
		}
		fmt.Println()
	}

}

/*func SikuBawahKi(a, b int) {
	for x := a; x <= b; x++ {
		for y := 0; y < x; y++ {
			fmt.Printf("1")
		}
		fmt.Println("0")
	}
}
*/
func Chess(lenght int) {
	var k int = 0
	for i := 1; i <= lenght; i++ {
		k = 0

		for space := 1; space <= lenght; space++ {
			if k%2 == 1 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
			k++
		}
		fmt.Println("")
	}
}

/*
func Ches(lenght int) {
	var k int = 0
	for i := 1; i <= lenght; i++ {
		k = 0

		for space := 1; space <= lenght; space++ {
			if k%2 == 1 {
				fmt.Print(" ")
			} else {
				fmt.Println("#")
			}
			k++
		}
		fmt.Println("")
	}
}
*/
