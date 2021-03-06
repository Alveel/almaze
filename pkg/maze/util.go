package maze

import (
	"fmt"
	"log"

	"github.com/Alveel/almaze/pkg/models"
)

//FindExits find exits in the maze by looping over the outer rows, first horizontally, then vertically
func FindExits(maze *models.Maze) (*models.MazeField, *models.MazeField) {
	var exits []*models.MazeField

	// Find exits on Y-axis
	horizontalWalls := [2]*models.MazeLine{
		maze.Lines[0],
		maze.Lines[maze.Height-1],
	}
	// TODO: find way to prevent for-loop nesting
	for _, ml := range horizontalWalls {
		for _, mf := range ml.Fields {
			if !mf.Wall {
				log.Printf("Exit found on horizontal wall at X%d/Y%d", mf.X, mf.Y)
				exits = append(exits, mf)
			}
		}
	}

	// Find exits on Y-axis
	// TODO: find way to prevent for-loop nesting
	// First loop through all lines
	for _, ml := range maze.Lines {
		// Then loop through all the fields
		for _, mf := range ml.Fields {
			// If the field is at the left-most or right-most of the maze, and it's NOT a wall, it's an exit.
			if (mf.X == 0 || mf.X == maze.Width-1) && !mf.Wall {
				log.Printf("Exit found on vertical wall at X%d/Y%d", mf.X, mf.Y)
				exits = append(exits, mf)
			}
		}
	}

	exitCount := len(exits)
	if exitCount != 2 {
		panic(fmt.Sprintf("Expected 2 entrances/exits, found %d!", exitCount))
	}

	return exits[0], exits[1]
}

//AlreadyVisited check if the tf (target field) has already been visited.
func AlreadyVisited(tf *models.MazeField, vf []*models.MazeField) bool {
	for _, f := range vf {
		if f == tf {
			return true
		}
	}
	return false
}
