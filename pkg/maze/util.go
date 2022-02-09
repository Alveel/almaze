package maze

import (
	"almaze/pkg/models"
	"log"
)

// TODO: clean this up, more code-reuse!
// Interfaces how?

func FindEntrance(maze models.Maze) models.MazeField {
	for _, mf := range maze.Fields {
		// Find entrance on X-axis
		if mf.Y == 0 {
			if mf.Walkable {
				log.Printf("Top-most entrance found at X%d/Y%d", mf.X, mf.Y)
				return mf
			}
		}

		// Find entrance on Y-axis
		if mf.X == 0 {
			if mf.Walkable {
				log.Printf("Most-left entrance found at X%d/Y%d", mf.X, mf.Y)
				return mf
			}
		}
	}

	panic("Entrance not found!")
}

func FindExit(maze models.Maze) models.MazeField {
	for _, mf := range maze.Fields {
		// Find exit on X-axis
		if mf.Y == maze.Height {
			if mf.Walkable {
				log.Printf("Bottom-most exit found at X%d/Y%d", mf.X, mf.Y)
				return mf
			}
		}

		// Find exit on Y-axis
		if mf.X == maze.Width {
			if mf.Walkable {
				log.Printf("Rightmost exit found at X%d/Y%d", mf.X, mf.Y)
				return mf
			}
		}
	}

	panic("Exit! not found!")
}
