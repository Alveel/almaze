package solver

import (
	"log"

	"github.com/Alveel/almaze/v2/pkg/maze"
	"github.com/Alveel/almaze/v2/pkg/models"
)

func WallFollower(m models.Maze) {
	solved := false
	directions := []int{maze.UP, maze.LEFT, maze.DOWN, maze.RIGHT}
	lastDirection := maze.UP

	for !solved {
		log.Printf("Current location: Y%d/Y%d\n", m.CurrentField.X, m.CurrentField.Y)
		if m.CurrentField == m.Exit {
			solved = true
			log.Println("Exit found!")
			break
		}

		// First try to move in the same direction as last move.
		nf, err := maze.Move(m, lastDirection)
		if err != nil {
			log.Println(err.Error())
		} else {
			nf.Visited = true
			m.CurrentField = &nf
			continue
		}

		// Otherwise, try to move in other directions
		for _, direction := range directions {
			// Skip this direction if it's the same as last direction
			if direction == lastDirection {
				continue
			}
			nf, err = maze.Move(m, direction)
			if err != nil {
				log.Println(err.Error())
			} else {
				nf.Visited = true
				m.CurrentField = &nf
				continue
			}
		}
	}
}
