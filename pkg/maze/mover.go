package maze

import (
	"almaze/pkg/models"
	"fmt"
	"log"
	"strings"
)

const (
	UP    int = 1
	DOWN  int = 2
	LEFT  int = 3
	RIGHT int = 4
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
func Move(m models.Maze, direction int) (models.MazeField, error) {
	cf := m.CurrentField
	var tf models.MazeField

	// Move in given direction
	if direction == UP {
		tf = m.Lines[cf.Y-1].Fields[cf.X]
		log.Println("Moving up")
	} else if direction == DOWN {
		tf = m.Lines[cf.Y+1].Fields[cf.X]
		log.Println("Moving down")
	} else if direction == LEFT {
		tf = m.Lines[cf.Y].Fields[cf.X-1]
		log.Println("Moving left")
	} else if direction == RIGHT {
		tf = m.Lines[cf.Y].Fields[cf.X+1]
		log.Println("Moving right")
	}

	// Check for move validity
	var reason strings.Builder
	if tf.Wall {
		reason.WriteString("hit wall")
	}
	if tf.Visited {
		log.Println("Already visited")
	}

	if reason.Len() > 0 {
		// Create MoveError
		moveError := &MoveError{
			CurrentField: *cf,
			TargetField:  tf,
			msg:          fmt.Sprintf("Failed moving from field X%d/Y%d to field X%d/Y%d (reason: %s)", cf.X, cf.Y, tf.X, tf.Y, reason),
		}
		return *cf, moveError
	}

	return tf, nil
}
