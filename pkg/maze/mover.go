package maze

import (
	"almaze/pkg/models"
	"fmt"
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
	return fmt.Sprintf("Error moving %s", e.msg)
}

func Move(maze models.Maze, cf models.MazeField, direction int) (models.MazeField, error) {
	var tf models.MazeField

	if direction == UP {
		tf = maze.Lines[cf.Y-1].Fields[cf.X]
	} else if direction == DOWN {
		tf = maze.Lines[cf.Y+1].Fields[cf.X]
	} else if direction == LEFT {
		tf = maze.Lines[cf.Y].Fields[cf.X-1]
	} else if direction == RIGHT {
		tf = maze.Lines[cf.Y].Fields[cf.X+1]
	}

	if !tf.Wall {
		return tf, nil
	}

	moveError := &MoveError{
		CurrentField: cf,
		TargetField:  tf,
		msg:          fmt.Sprintf("Failed moving from field X%d/Y%d to field X%d/Y%d (%s)", cf.X, cf.Y, tf.X, tf.Y, "hit wall"),
	}
	return cf, moveError
}
