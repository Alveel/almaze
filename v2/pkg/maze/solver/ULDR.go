package solver

import (
	"github.com/Alveel/almaze/v2/pkg/maze"
	"github.com/Alveel/almaze/v2/pkg/models"
	"log"
)

func ULDR(m *models.Maze, p *models.Player) {
	solved := false
	directions := []int{maze.UP, maze.LEFT, maze.DOWN, maze.RIGHT}
	p.WalkedRoute = append(p.WalkedRoute, m.Entrance)

	for !solved {
		//log.Printf("Current location: X%d/Y%d\n", p.CurrentField.X, p.CurrentField.Y)

		if p.CurrentField == m.Exit {
			solved = true
			log.Println("Exit found!")
			break
		}

		// First try to move in the same direction as the player is already facing.
		nf, err := maze.Move(m, p, p.FacingDirection)
		if err != nil {
			//log.Println(err.Error())
		} else {
			p.WalkedRoute = append(p.WalkedRoute, nf)
			p.CurrentField = nf
			continue
		}

		// Otherwise, try to move in other directions
		for _, direction := range directions {
			// Skip if it's the direction Player is facing, because we already tried that.
			if direction == p.FacingDirection {
				continue
			}
			nf, err := maze.Move(m, p, direction)
			if err != nil {
				//log.Println(err.Error())
				continue
			}
			if maze.AlreadyVisited(nf, p.WalkedRoute) {
				//log.Printf("Current location: X%d/Y%d\n", p.CurrentField.X, p.CurrentField.Y)
				//log.Printf("Already visited X%d/Y%d", nf.X, nf.Y)
				continue
			}

			p.WalkedRoute = append(p.WalkedRoute, nf)
			p.CurrentField = nf
			p.FacingDirection = direction
			break // break out of direction loop
		}
	}
}
