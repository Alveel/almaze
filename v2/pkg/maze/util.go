package maze

import (
	"fmt"
	"log"

	"github.com/Alveel/almaze/v2/pkg/models"
)

func FindExits(maze models.Maze) (*models.MazeField, *models.MazeField) {
	var exits []models.MazeField

	// Find exits on Y-axis
	// We can
	horizontalWalls := [2]models.MazeLine{
		maze.Lines[0],
		maze.Lines[maze.Height-1],
	}
	for _, ml := range horizontalWalls {
		for _, mf := range ml.Fields {
			if !mf.Wall {
				log.Printf("Exit found on horizontal wall at Y%d/Y%d", mf.X, mf.Y)
				exits = append(exits, mf)
			}
		}
	}

	// Find exits on Y-axis
	// First loop through all lines
	for _, ml := range maze.Lines {
		// Then loop through all the fields
		for _, mf := range ml.Fields {
			// If the field is at the left-most or right-most of the maze, and it's a wall, it's an exit.
			if (mf.X == 1 || mf.X == maze.Width) && !mf.Wall {
				log.Printf("2Exit found on vertical wall at Y%d/Y%d", mf.X, mf.Y)
				exits = append(exits, mf)
			}
		}
	}

	exitCount := len(exits)
	if exitCount != 2 {
		panic(fmt.Sprintf("Expected 2 entrances/exits, found %d!", exitCount))
	}

	return &exits[0], &exits[1]
}
