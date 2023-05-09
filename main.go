package main

import (
	"fmt"

	"training.go/helloworld/tool"
)

// func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {

// }

func main() {
	nbLines, nbOcc, lines, finalText, err := tool.FindReplaceFile("go-wiki.txt", "Go", "Perl")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("==Summary==")
	fmt.Printf("Number of occurrences of Go:%d\n", nbOcc)
	fmt.Printf("Number of lines:%d\n", nbLines)
	fmt.Printf("Lines:%s\n", lines)
	fmt.Println("==End of Summary==")
	fmt.Println("==Final Text==")
	fmt.Println(finalText)

	errw := tool.WriteNewFile(finalText)
	if errw != nil {
		fmt.Println(err)
	}
}
