package maze

import (
	"bufio"
	"log"
	"os"

	"github.com/Alveel/almaze/pkg/models"
)

func isWall(field rune) bool {
	if field == ' ' {
		return false
	}
	return true
}

//LoadMaze attempts to load a maze from a text file. Everything that is not a whitespace is seen as a wall.
func LoadMaze(mazeFile string) models.Maze {
	// Load the text file
	file, err := os.Open(mazeFile)
	if err != nil {
		log.Fatalf("Failed loading maze file %s", mazeFile)
	}

	// Create empty []MazeLine object, which will be filled with MazeLine[MazeField]
	var mazeLines []*models.MazeLine

	s := bufio.NewScanner(file)
	// Parse maze line by line
	s.Split(bufio.ScanLines)

	currentLineNumber, maxWidth := 0, 0

	// Scan each line
	for s.Scan() {
		var ml models.MazeLine
		// Find each rune/character in a line
		data := []rune(s.Text())
		currentLineWidth := len(data)
		for i := 0; i < currentLineWidth; i++ {
			// Create a new MazeField
			mf := models.NewMazeField(i, currentLineNumber, isWall(data[i]))
			ml.Fields = append(ml.Fields, mf)
		}
		// Find the maximum width of the maze
		if maxWidth < currentLineWidth {
			maxWidth = currentLineWidth
		}

		mazeLines = append(mazeLines, &ml)
		currentLineNumber++
	}

	// Instantiate the maze
	maze := models.NewMaze(maxWidth, currentLineNumber, mazeLines)
	log.Printf("Maze width: %d, height: %d, %s", maze.Width, maze.Height, LineBreak)
	maze.Entrance, maze.Exit = FindExits(&maze)
	log.Printf("Entrance: %dX/%dY%s", maze.Entrance.X, maze.Entrance.Y, LineBreak)
	log.Printf("Entrance pointer: %p", maze.Entrance)
	log.Printf("Exit: %dX/%dY%s", maze.Exit.X, maze.Exit.Y, LineBreak)
	log.Printf("Exit pointer: %p", maze.Exit)

	return maze
}
