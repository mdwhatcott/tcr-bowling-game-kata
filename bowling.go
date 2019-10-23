package bowling

type Game struct {
	score int
}

func (this *Game) Roll(pins int) {
	this.score += pins
}

func (this *Game) Score() int {
	return this.score
}
