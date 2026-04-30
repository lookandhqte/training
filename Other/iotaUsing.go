package other

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return []string{"North", "East", "South", "West"}[d]
}

// func main() {
// 	north := North
// 	fmt.Println(north)
// }
