package solver

import (
	"log"

	"github.com/Alveel/almaze/pkg/maze"
	"github.com/Alveel/almaze/pkg/models"
)

func ULDR(m *models.Maze, p *models.Player) {
	solved := false
	directions := []int{maze.UP, maze.LEFT, maze.DOWN, maze.RIGHT}
	p.WalkedRoute = append(p.WalkedRoute, m.Entrance)

MazeLoop:
	for !solved {
		//log.Printf("Current location: X%d/Y%d\n", p.CurrentField.X, p.CurrentField.Y)

		if p.CurrentField == m.Exit {
			solved = true
			log.Println("Exit found!")
			break MazeLoop
		}

		// First try to move in the same direction as the player is already facing.
		nf, err := maze.TryToMove(m, p)
		if err != nil {
			//log.Println(err.Error())
		} else {
			p.WalkedRoute = append(p.WalkedRoute, nf)
			p.CurrentField = nf
			continue MazeLoop
		}

		// Otherwise, try to move in other directions
	DirectionLoop:
		for _, direction := range directions {
			// Skip if it's the direction Player is facing, because we already tried that.
			if direction == p.FacingDirection {
				continue DirectionLoop
			}
			p.FacingDirection = direction
			nf, err := maze.TryToMove(m, p)
			if err != nil {
				//log.Println(err.Error())
				continue DirectionLoop
			}
			if maze.AlreadyVisited(nf, p.WalkedRoute) {
				//log.Printf("Current location: X%d/Y%d\n", p.CurrentField.X, p.CurrentField.Y)
				//log.Printf("Already visited X%d/Y%d", nf.X, nf.Y)
				continue DirectionLoop
			}

			p.WalkedRoute = append(p.WalkedRoute, nf)
			p.CurrentField = nf
			continue MazeLoop
		}
	}
}
