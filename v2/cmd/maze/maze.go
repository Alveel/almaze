package main

import (
	"almaze/pkg/maze"
	"almaze/pkg/maze/solver"
	"log"
	"time"
)

func main() {
	log.Printf("The current time is: %s%s", time.Now().Format("2006-01-02 15:04:05"), maze.LineBreak)
	start := time.Now()

	//LoadMaze("maze.txt")
	myMaze := maze.LoadMaze("maze.txt")
	maze.DrawMaze(myMaze)
	time.Sleep(100 * time.Millisecond)
	solver.ULDR(myMaze)
	//maze.DrawMaze(myMaze)
	log.Printf("The current time is: %s%s", time.Now().Format("2006-01-02 15:04:05"), maze.LineBreak)
	duration := time.Since(start)
	log.Printf("Execution took %dms%s", duration.Milliseconds(), maze.LineBreak)
}
