package solver

import (
	"github.com/Alveel/almaze/pkg/maze"
	"github.com/Alveel/almaze/pkg/models"
	"log"
)

//WallFollower keeps the Player glued to the wall, following any turns until the exit is reached, if possible.
// This algorithm doesn't work if Player ends up on an "island", they will just keep walking in circles.
func WallFollower(m *models.Maze, p *models.Player) {
	solved := false

MazeLoop:
	for !solved {
		log.Printf("Current location: X%d/Y%d\n", p.CurrentField.X, p.CurrentField.Y)
		if p.CurrentField == m.Exit {
			solved = true
			log.Println("Exit found!")
			break
		}

		turner := maze.Turner{Maze: m, Player: p}

		// The idea here is to always try to turn right and move forward.
		// If that move is illegal, we turn back and move forwards.
		// If that move is also illegal we turn left and move forwards.
		turner.Player.FacingDirection = turner.TurnRight()

		for i := 0; i < 4; i++ {
			nf, err := maze.TryToMove(m, p)
			if err == nil {
				p.WalkedRoute = append(p.WalkedRoute, nf)
				p.CurrentField = nf
				// We have found a valid path! Continue with the MazeLoop or we will walk in the wrong direction.
				continue MazeLoop
			}
			turner.Player.FacingDirection = turner.TurnLeft()
		}

		log.Fatalf("Unable to move in any direction. Stuck.")
	}
}
