package maze

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Printf("The current time is: %v\n", time.Now())

	LoadMaze("maze.txt")
}
