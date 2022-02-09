package models

type MazeField struct {
	Visited  bool
	X        int
	Y        int
	Walkable bool
}

type Maze struct {
	Width    int
	Height   int
	Fields   []MazeField
	Entrance MazeField
	Exit     MazeField
}

func NewMazeField(x, y int, walk bool) *MazeField {
	mf := new(MazeField)
	mf.Visited = false
	mf.X = x
	mf.Y = y
	mf.Walkable = walk
	return mf
}
