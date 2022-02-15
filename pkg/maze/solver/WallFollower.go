package solver

import (
	"log"

	"github.com/Alveel/almaze/pkg/maze"
	"github.com/Alveel/almaze/pkg/models"
)

func WallFollower(m *models.Maze, p *models.Player) {
	solved := false

	for !solved {
		log.Printf("Current location: Y%d/Y%d\n", p.CurrentField.X, p.CurrentField.Y)
		if &p.CurrentField == &m.Exit {
			solved = true
			log.Println("Exit found!")
			break
		}

		mover := maze.Mover{Maze: m, Player: p}
		// First try to move in the same direction as last move.
		//nf, err := maze.Move(m, p, lastDirection)
		nf, err := mover.MoveStraight()
		if err != nil {
			// Otherwise, we try to move right!
			nf, err := mover.MoveRight()
			if err != nil {
				log.Printf("Error moving right: %v", err.Error())
			} else {
				p.WalkedRoute = append(p.WalkedRoute, nf)
				p.CurrentField = nf
				// How to properly handle this? I already have some of this logic in the MoveDirection functions in move.go
				//p.FacingDirection = p.FacingDirection
			}
		} else {
			p.WalkedRoute = append(p.WalkedRoute, nf)
			p.CurrentField = nf
			continue
		}
	}
}
