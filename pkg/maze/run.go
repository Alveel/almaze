package maze

import (
	"fmt"
	"log"
	"strings"
	"time"

	"almaze/pkg/models"
)

var wallSymbol = "@"

func BuildMaze(maze models.Maze) string {
	var sb strings.Builder

	for _, mf := range maze.Fields {
		if mf.Wall {
			sb.WriteString(wallSymbol)
		} else {
			sb.WriteString(" ")
		}

		// End of line reached
		if mf.X == maze.Width {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func Run() {
	log.Printf("The current time is: %v\n", time.Now())

	//LoadMaze("maze.txt")
	myMaze := LoadMaze("maze.txt")
	mazeString := BuildMaze(myMaze)
	fmt.Print(mazeString)
}
