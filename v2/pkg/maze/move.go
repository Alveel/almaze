package maze

import (
	"fmt"
	"strings"

	"github.com/Alveel/almaze/v2/pkg/models"
)

type MoveError struct {
	CurrentField models.MazeField
	TargetField  models.MazeField
	msg          string
}

func (e *MoveError) Error() string {
	return fmt.Sprintf("Error moving: %s", e.msg)
}

//Move from given MazeField in Maze
// TODO: maybe it's better to have a direction in which the "cursor" is looking, and move straight/left/right/backwards?
func Move(m models.Maze, p models.Player, direction int) (models.MazeField, error) {
	var tf models.MazeField
	var reasons []string

	// Move in given direction
	if direction == UP {
		//log.Println("Moving up")
		if p.CurrentField.Y <= 0 {
			reasons = append(reasons, "out of bounds")
		} else {
			tf = m.Lines[p.CurrentField.Y-1].Fields[p.CurrentField.X]
		}
	} else if direction == DOWN {
		//log.Println("Moving down")
		if p.CurrentField.Y+1 >= m.Height {
			reasons = append(reasons, "out of bounds")
		} else {
			tf = m.Lines[p.CurrentField.Y+1].Fields[p.CurrentField.X]
		}
	} else if direction == LEFT {
		//log.Println("Moving left")
		if p.CurrentField.X <= 0 {
			reasons = append(reasons, "out of bounds")
		} else {
			tf = m.Lines[p.CurrentField.Y].Fields[p.CurrentField.X-1]
		}
	} else if direction == RIGHT {
		//log.Println("Moving right")
		if p.CurrentField.X+1 >= m.Width {
			reasons = append(reasons, "out of bounds")
		} else {
			tf = m.Lines[p.CurrentField.Y].Fields[p.CurrentField.X+1]
		}
	}

	// Check for move validity
	if tf.Wall {
		reasons = append(reasons, "hit wall")
	}

	if len(reasons) > 0 {
		// Create MoveError
		moveError := &MoveError{
			CurrentField: p.CurrentField,
			TargetField:  tf,
			msg: fmt.Sprintf("Failed moving from field Y%d/Y%d to field Y%d/Y%d (reason: %s)",
				p.CurrentField.X, p.CurrentField.Y,
				tf.X, tf.Y,
				strings.Join(reasons, ", "),
			),
		}
		return p.CurrentField, moveError
	}

	return tf, nil
}
