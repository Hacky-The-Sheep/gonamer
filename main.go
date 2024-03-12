package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Compile the regex for later use
	r, _ := regexp.Compile(`^.*\d\d\d\d\.`)

	for i := 1; i < len(os.Args); i++ {
		mName := os.Args[i]

		fmt.Println("")
		fmt.Println("Movie Name:", mName)
		idx := r.FindStringIndex(mName)
		ogMatch := strings.Split(mName[idx[0]:idx[1]-1], ".")
		match := strings.Join(ogMatch, " ")

		newTitle := fmt.Sprintf("%s(%s).%s", match[0:len(match)-4], match[len(match)-4:], mName[len(mName)-3:])
		var answer string

		fmt.Println("The movie will be renamed to:", newTitle)
		fmt.Print("y to accept, n to skip: ")
		fmt.Scan(&answer)
		if answer == "y" {
			os.Rename(mName, newTitle)
			fmt.Println("The file has been renamed!")
		} else {
			continue
		}

		if len(os.Args) > 2 {
			continue
		} else {
			break
		}
	}

}
