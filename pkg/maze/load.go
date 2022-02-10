package maze

import (
	"almaze/pkg/models"
	"bufio"
	"log"
	"os"
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

	// Create empty []MazeField object
	var mazeFields []models.MazeField

	s := bufio.NewScanner(file)
	// Parse maze line by line
	s.Split(bufio.ScanLines)

	curLine, maxWidth := 0, 0

	// Scan each line
	for s.Scan() {
		curLine++
		// Find each rune/character in a line
		data := []rune(s.Text())
		curWidth := len(data)
		for i := 0; i < curWidth; i++ {
			// Create a new MazeField
			mf := models.NewMazeField(i+1, curLine, isWall(data[i]))
			mazeFields = append(mazeFields, *mf)
		}
		// Find the maximum width of the maze
		if maxWidth < curWidth {
			maxWidth = curWidth
		}
	}

	// Instantiate the maze
	maze := models.NewMaze(maxWidth, curLine, mazeFields)
	log.Printf("Maze width: %d, height: %d", maze.Width, maze.Height)
	maze.Entrance = FindEntrance(maze)
	maze.Exit = FindExit(maze)
	log.Printf("Entrance: %dX/%dY", maze.Entrance.X, maze.Entrance.Y)
	log.Printf("Exit: %dX/%dY", maze.Exit.X, maze.Exit.Y)

	return maze
}
