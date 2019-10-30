package main

import (
	"fmt"
	"github.com/samtcifihi/TMG/src/goban"
	_ "local/userio"
)

func main() {
	fmt.Println("Welcome to TMG")

	// testPyrPile := new(pyramid.PyrPile)

	// testPyrPile.Add("red", 1)
	// testPyrPile.Add("green", 2)
	// testPyrPile.Add("blue", 3)

	// fmt.Println("Here are the test pyramids:")

	// testPyr1 := testPyrPile.GetPyr(0)
	// testPyr2 := testPyrPile.GetPyr(1)
	// testPyr3 := testPyrPile.GetPyr(2)

	// fmt.Println("1: Color = ", testPyr1.Color())
	// fmt.Println("1: Size = ", testPyr1.Size())
	// fmt.Println("2: Color = ", testPyr2.Color())
	// fmt.Println("2: Size = ", testPyr2.Size())
	// fmt.Println("3: Color = ", testPyr3.Color())
	// fmt.Println("3: Size = ", testPyr3.Size())

	testGoban := goban.NewSq(9, 9)

	fmt.Println("Size: ", testGoban.Height(), "x", testGoban.Width())
	// fmt.Println("2-2 point: ", testGoban.GetPoint(1, 1))

	testGoban.Print()

	fmt.Println("Thank you for playing TMG.")
}
