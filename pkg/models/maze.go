package models

type MazeField struct {
	X       int
	Y       int
	Wall    bool
	Visited bool
}

func NewMazeField(x, y int, wall bool) *MazeField {
	mf := new(MazeField)
	mf.X = x
	mf.Y = y
	mf.Wall = wall
	mf.Visited = false
	return mf
}

type MazeLine struct {
	X      int
	Fields []MazeField
}

type Maze struct {
	Width    int
	Height   int
	Lines    []MazeLine
	Entrance *MazeField
	Exit     *MazeField
}

func NewMaze(w, h int, f []MazeLine) Maze {
	maze := new(Maze)
	maze.Width = w
	maze.Height = h
	maze.Lines = f
	return *maze
}
