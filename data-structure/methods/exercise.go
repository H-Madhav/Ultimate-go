package main

import "fmt"

type player struct {
	name   string
	atBats int
	hits   int
}

func (p player) average() float64 {
	if p.atBats == 0 {
		return 0.0
	}
	return float64(p.hits) / float64(p.atBats)
}

func main() {
	players := []player{{"Madhav", 1, 100}, {"Keshav", 2, 100}}

	for _, player := range players {
		fmt.Println("Name:", player.name, "Avg:", player.average())
	}
}
