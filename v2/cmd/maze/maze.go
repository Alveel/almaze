package main

import (
	"github.com/Alveel/almaze/v2/pkg/models"
	"log"
	"time"

	"github.com/Alveel/almaze/v2/pkg/maze"
	"github.com/Alveel/almaze/v2/pkg/maze/solver"
)

func main() {
	log.Printf("The current time is: %s%s", time.Now().Format("2006-01-02 15:04:05"), maze.LineBreak)
	start := time.Now()

	//LoadMaze("maze.txt")
	myMaze := maze.LoadMaze("maze.txt")
	maze.DrawMaze(myMaze)
	player := models.NewPlayer(myMaze.Entrance, maze.DOWN)
	solver.ULDR(myMaze, player)
	//maze.DrawMaze(myMaze)
	log.Printf("The current time is: %s%s", time.Now().Format("2006-01-02 15:04:05"), maze.LineBreak)
	duration := time.Since(start)
	log.Printf("Execution took %dms%s", duration.Milliseconds(), maze.LineBreak)
}
