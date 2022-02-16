package maze

import (
	"fmt"
	"strings"

	"github.com/Alveel/almaze/pkg/models"
)

type MoveError struct {
	CurrentField models.MazeField
	TargetField  models.MazeField
	msg          string
}

func (e *MoveError) Error() string {
	return fmt.Sprintf("Error moving: %s", e.msg)
}

type TurnDirection interface {
	TurnLeft() int
	TurnRight() int
	TurnAround() int
}

type Turner struct {
	Maze   *models.Maze
	Player *models.Player
}

//MoveStraight can just use the current FacingDirection of Player.
func MoveStraight(m *models.Maze, p *models.Player) (*models.MazeField, error) {
	return TryToMove(m, p)
}

// TurnLeft
// Directions are clockwise as defined in constants.go, so moving left means going counterclockwise.
// If we loop over (under?) the directions, wrap around. (Am I even using terminology that makes sense?)
func (mv Turner) TurnLeft() int {
	if mv.Player.FacingDirection-1 < UP {
		return LEFT
	}
	return mv.Player.FacingDirection - 1
}

// TurnRight
// Directions are clockwise as defined in constants.go, so moving right means going clockwise.
// If we loop over (under?) the directions, wrap around. (Am I even using terminology that makes sense?)
func (mv Turner) TurnRight() int {
	if mv.Player.FacingDirection+1 > LEFT {
		return UP
	}
	return mv.Player.FacingDirection + 1
}

// TurnAround to turn around; backtrack
func (mv Turner) TurnAround() int {
	if mv.Player.FacingDirection+2 > LEFT {
		return mv.Player.FacingDirection - 4
	}
	return mv.Player.FacingDirection + 2
}

//TryToMove in Maze from Player.CurrentField MazeField in direction Player.FacingDirection
func TryToMove(m *models.Maze, p *models.Player) (*models.MazeField, error) {
	// Create instance to set empty fields. Will be overwritten if a move is valid.
	var tf = new(models.MazeField)
	var reasons []string

	// TryToMove in given direction
	if p.FacingDirection == UP {
		if p.CurrentField.Y <= 0 {
			reasons = append(reasons, "out of bounds")
		} else {
			tf = m.Lines[p.CurrentField.Y-1].Fields[p.CurrentField.X]
			//log.Println("Moving up")
		}
	} else if p.FacingDirection == DOWN {
		if p.CurrentField.Y+1 >= m.Height {
			reasons = append(reasons, "out of bounds")
		} else {
			tf = m.Lines[p.CurrentField.Y+1].Fields[p.CurrentField.X]
			//log.Println("Moving down")
		}
	} else if p.FacingDirection == LEFT {
		if p.CurrentField.X <= 0 {
			reasons = append(reasons, "out of bounds")
		} else {
			tf = m.Lines[p.CurrentField.Y].Fields[p.CurrentField.X-1]
			//log.Println("Moving left")
		}
	} else if p.FacingDirection == RIGHT {
		if p.CurrentField.X+1 >= m.Width {
			reasons = append(reasons, "out of bounds")
		} else {
			tf = m.Lines[p.CurrentField.Y].Fields[p.CurrentField.X+1]
			//log.Println("Moving right")
		}
	}

	// Check for move validity
	if tf.Wall {
		reasons = append(reasons, "hit wall")
	}

	if len(reasons) > 0 {
		// Create MoveError
		moveError := &MoveError{
			CurrentField: *p.CurrentField,
			TargetField:  *tf,
			msg: fmt.Sprintf("Failed moving from field X%d/Y%d to field X%d/Y%d (reason: %s)",
				p.CurrentField.X, p.CurrentField.Y,
				tf.X, tf.Y,
				strings.Join(reasons, ", "),
			),
		}
		return p.CurrentField, moveError
	}

	return tf, nil
}
