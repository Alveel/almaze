package maze

import (
	"almaze/pkg/models"
	"log"
)

// TODO: clean this up, more code-reuse!
// Interfaces how?

//FindEntrance tries to find the entrance on either the top-most row or left-most column.
//If it cannot find the entrance, it panics.
func FindEntrance(maze models.Maze) *models.MazeField {
	for _, mf := range maze.Fields {
		// Find entrance on X-axis
		if mf.Y == 1 {
			if !mf.Wall {
				log.Printf("Entrance found on top border at X%d/Y%d", mf.X, mf.Y)
				return &mf
			}
		}

		// Find entrance on Y-axis
		if mf.X == 1 {
			if !mf.Wall {
				log.Printf("Entrance found on left border at X%d/Y%d", mf.X, mf.Y)
				return &mf
			}
		}
	}

	panic("Entrance not found!")
}

//FindExit tries to find the exit on the bottom-most row or right-most column.
//If it cannot find the exit, it panics.
func FindExit(maze models.Maze) *models.MazeField {
	for _, mf := range maze.Fields {
		// Find exit on X-axis
		if mf.Y == maze.Height {
			if !mf.Wall {
				log.Printf("Exit found on bottom border at X%d/Y%d", mf.X, mf.Y)
				return &mf
			}
		}

		// Find exit on Y-axis
		if mf.X == maze.Width {
			if !mf.Wall {
				log.Printf("Exit found on right border at X%d/Y%d", mf.X, mf.Y)
				return &mf
			}
		}
	}

	panic("Exit not found!")
}
