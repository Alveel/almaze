package models

type MazeField struct {
	X    int
	Y    int
	Wall bool
}

func NewMazeField(x, y int, wall bool) *MazeField {
	mf := new(MazeField)
	mf.X = x
	mf.Y = y
	mf.Wall = wall
	return mf
}

type MazeLine struct {
	Fields []*MazeField
}

type Maze struct {
	Width    int
	Height   int
	Lines    []*MazeLine
	Entrance *MazeField
	Exit     *MazeField
}

func NewMaze(w, h int, f []*MazeLine) Maze {
	maze := new(Maze)
	maze.Width = w
	maze.Height = h
	maze.Lines = f
	return *maze
}

type Player struct {
	CurrentField    *MazeField
	FacingDirection int
	WalkedRoute     []*MazeField
}

func NewPlayer(mf *MazeField, fd int) *Player {
	player := new(Player)
	player.CurrentField = mf
	player.FacingDirection = fd
	return player
}
