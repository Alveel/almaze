package maze

import (
	"fmt"
	"strings"

	"github.com/Alveel/almaze/pkg/models"
)

func DrawMaze(m models.Maze) {
	var sb strings.Builder

	for _, ml := range m.Lines {
		for _, mf := range ml.Fields {
			if mf.Wall {
				sb.WriteString(WallSymbol)
			} else {
				sb.WriteString(" ")
			}

			// End of line reached
			if mf.X == m.Width-1 {
				sb.WriteString(LineBreak)
			}
		}
	}

	fmt.Print(sb.String())
}

//func DrawRoute(m models.Maze, p models.Player) {
//	var sb strings.Builder
//
//}
