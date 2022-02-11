package maze

import (
	"fmt"
	"log"
	"strings"
	"time"

	"almaze/pkg/models"
)

var wallSymbol = "@"
var visitedSymbol = "."

func DrawMaze(maze models.Maze) string {
	var sb strings.Builder

	for _, ml := range maze.Lines {
		for _, mf := range ml.Fields {
			if mf.Wall {
				sb.WriteString(wallSymbol)
			} else if mf.Visited {
				sb.WriteString(visitedSymbol)
			} else {
				sb.WriteString(" ")
			}

			// End of line reached
			if mf.X == maze.Width {
				sb.WriteString("\n")
			}
		}
	}

	return sb.String()
}

func Run() {
	now := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("The current time is: %s\n", now)

	//LoadMaze("maze.txt")
	myMaze := LoadMaze("maze.txt")
	mazeString := DrawMaze(myMaze)
	fmt.Print(mazeString)
}
