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

type MoveDirection interface {
	MoveStraight() (*models.MazeField, error)
	MoveLeft() (*models.MazeField, error)
	MoveRight() (*models.MazeField, error)
}

type Mover struct {
	Maze   *models.Maze
	Player *models.Player
}

func (mv Mover) MoveStraight() (*models.MazeField, error) {
	return Move(mv.Maze, mv.Player, mv.Player.FacingDirection)
}

// MoveLeft
// Directions are clockwise as defined in constants.go, so moving left means going counterclockwise.
// If we loop over (under?) the directions, wrap around. (Am I even using terminology that makes sense?)
func (mv Mover) MoveLeft() (*models.MazeField, error) {
	if mv.Player.FacingDirection-1 <= UP {
		return Move(mv.Maze, mv.Player, LEFT)
	}
	return Move(mv.Maze, mv.Player, mv.Player.FacingDirection-1)
}

// MoveRight
// Directions are clockwise as defined in constants.go, so moving right means going clockwise.
// If we loop over (under?) the directions, wrap around. (Am I even using terminology that makes sense?)
func (mv Mover) MoveRight() (*models.MazeField, error) {
	if mv.Player.FacingDirection+1 >= LEFT {
		return Move(mv.Maze, mv.Player, UP)
	}
	return Move(mv.Maze, mv.Player, mv.Player.FacingDirection+1)
}

//Move from given MazeField in Maze
// TODO: maybe it's better to have a direction in which the "cursor" is looking, and move straight/left/right/backwards?
func Move(m *models.Maze, p *models.Player, direction int) (*models.MazeField, error) {
	// Create instance to set empty fields. Will be overwritten if a move is valid.
	var tf = new(models.MazeField)
	var reasons []string

	// Move in given direction
	if direction == UP {
		if p.CurrentField.Y <= 0 {
			reasons = append(reasons, "out of bounds")
		} else {
			tf = m.Lines[p.CurrentField.Y-1].Fields[p.CurrentField.X]
			//log.Println("Moving up")
		}
	} else if direction == DOWN {
		if p.CurrentField.Y+1 >= m.Height {
			reasons = append(reasons, "out of bounds")
		} else {
			tf = m.Lines[p.CurrentField.Y+1].Fields[p.CurrentField.X]
			//log.Println("Moving down")
		}
	} else if direction == LEFT {
		if p.CurrentField.X <= 0 {
			reasons = append(reasons, "out of bounds")
		} else {
			tf = m.Lines[p.CurrentField.Y].Fields[p.CurrentField.X-1]
			//log.Println("Moving left")
		}
	} else if direction == RIGHT {
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
