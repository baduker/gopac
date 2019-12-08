package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var maze []string

func loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	return nil
}

func printScreen() {
	for _, line := range maze {
		fmt.Print(line)
	}
}

func readInput() (string, error) {
	buffer := make([]byte, 100)
	cnt, err := os.Stdin.Read(buffer)

	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	}
	return "", nil
}

func initialize() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatal("Unable to activate cbreak mode: ", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatal("Unable to activate cooked mode: ", err)
	}
}

func main() {
	//	initialize the game
	err := loadMaze("maze.txt")
	if err != nil {
		log.Print("Unable to load the maze. :(", err)
		return
	}

	// game loop
	for {
		//	update screen
		printScreen()
		// process input
		input, err := readInput()
		if err != nil {
			log.Println("Error reading input!", err)
		}
		// process collisions

		// check if game is over
		if input == "ESC" {
			break
		}
		// Temp: break infinite loop
		break

		// repeat
	}
}
