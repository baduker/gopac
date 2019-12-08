package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

		// process collisions

		// check if game is over

		// Temp: break infinite loop
		break

		// repeat
	}
}
