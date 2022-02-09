package maze

import (
	"almaze/pkg/models"
	"bufio"
	"log"
	"os"
)

func isWalkable(field string) bool {
	if field == " " {
		return true
	}
	return false
}

func LoadMaze(mazeFile string) models.Maze {
	file, err := os.Open(mazeFile)
	if err != nil {
		log.Fatalf("Failed loading maze file %s", mazeFile)
	}

	// Create empty Maze and []MazeField objects
	var maze = new(models.Maze)
	var mazeFields []models.MazeField

	s := bufio.NewScanner(file)
	// Parse maze line by line
	s.Split(bufio.ScanLines)

	curLine, maxWidth := 0, 0

	// Scan each line
	for s.Scan() {
		// Find each rune/character in a line
		data := []rune(s.Text())
		for i := 0; i < len(data); i++ {
			// Find the maximum width of the maze
			if maxWidth < i {
				maxWidth = i
			}

			// Create a new MazeField
			mf := models.NewMazeField(curLine, i, isWalkable(string(data[i])))
			//log.Printf("MazeField X: %d, Y: %d, Walkable: %t", mf.X, mf.Y, mf.Walkable)
			mazeFields = append(mazeFields, *mf)
		}
		curLine++
	}

	// Configure the maze
	maze.Width = maxWidth
	maze.Height = curLine
	maze.Fields = mazeFields
	maze.Entrance = FindEntrance(*maze)
	maze.Exit = FindExit(*maze)
	log.Printf("Maze width: %d, height: %d", maze.Width, maze.Height)

	return *maze
}
