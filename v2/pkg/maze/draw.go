package maze

import (
	"fmt"
	"strings"

	"github.com/Alveel/almaze/v2/pkg/models"
)

func DrawMaze(maze models.Maze) {
	var sb strings.Builder

	for _, ml := range maze.Lines {
		for _, mf := range ml.Fields {
			if mf.Wall {
				sb.WriteString(WallSymbol)
			} else if mf.Visited {
				sb.WriteString(VisitedSymbol)
			} else {
				sb.WriteString(" ")
			}

			// End of line reached
			if mf.X == maze.Width-1 {
				sb.WriteString(LineBreak)
			}
		}
	}

	fmt.Print(sb.String())
}
