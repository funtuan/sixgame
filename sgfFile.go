package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func openSgf(path string) [][]int8 {

	sgffile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	sgfString := string(sgffile)

	checkerboard := make([][]int8, size, size)
	for i := 0; i < size; i++ {
		checkerboard[i] = make([]int8, size)
	}

	// fmt.Println(sgfString)
	commands, vals := resolveText(sgfString)

	for index, command := range commands {
		// fmt.Println(command, index)
		switch command {
		case "SZ":
			if val, err := strconv.Atoi(vals[index]); err == nil {
				size = val
				checkerboard = make([][]int8, val, val)
				for i := 0; i < val; i++ {
					checkerboard[i] = make([]int8, val)
				}
			}
		case "AB":
			valRunes := []rune(vals[index])
			checkerboard[valRunes[0]-97][valRunes[1]-97] = 1
		case "AW":
			valRunes := []rune(vals[index])
			checkerboard[valRunes[0]-97][valRunes[1]-97] = -1
		}
	}
	return checkerboard
}

func resolveText(text string) ([]string, []string) {
	commands := []string{}
	vals := []string{}

	text = strings.Replace(text, "(", "", -1)
	text = strings.Replace(text, ")", "", -1)
	text = strings.Replace(text, ";", "", -1)
	// text = strings.Replace(text, "\n", "", -1)
	for len(text) > 2 {
		startIndex := strings.Index(text, "[")
		endIndex := strings.Index(text, "]")
		// fmt.Println(text)
		// fmt.Println(len(text))
		// fmt.Println(startIndex)
		command := text[0:startIndex]
		val := text[startIndex+1 : endIndex]

		if len(command) > 0 {
			commands = append(commands, command)
			vals = append(vals, val)
		} else {
			commands = append(commands, commands[len(commands)-1])
			vals = append(vals, val)
		}

		text = text[endIndex+1 : len(text)]
	}
	// fmt.Println(commands)
	// fmt.Println(vals)
	return commands, vals
}

func printCheckerboard(checkerboard [][]int8) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			switch checkerboard[x][y] {
			case 0:
				fmt.Print("．")
			case 1:
				fmt.Print("Ｘ")
			case -1:
				fmt.Print("Ｏ")
			}
		}
		fmt.Print("\n")
	}
}
