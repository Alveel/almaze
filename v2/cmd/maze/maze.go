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

	//LoadMaze("maze.txt")
	myMaze := maze.LoadMaze("maze.txt")
	maze.DrawMaze(myMaze)
	player := models.NewPlayer(myMaze.Entrance, maze.DOWN)
	start := time.Now()
	solver.ULDR(&myMaze, player)
	//maze.DrawMaze(myMaze)
	duration := time.Since(start)
	log.Printf("The current time is: %s%s", time.Now().Format("2006-01-02 15:04:05"), maze.LineBreak)
	log.Printf("Execution took %dns%s", duration.Microseconds(), maze.LineBreak)
}
