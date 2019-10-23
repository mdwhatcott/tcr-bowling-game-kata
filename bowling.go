package bowling

type Game struct {
	score  int
	throw  int
	throws [21]int
}

func (this *Game) RecordRoll(pins int) {
	this.throws[this.throw] = pins
	this.throw++
}

func (this *Game) CalculateScore() int {
	this.throw = 0
	for frame := 0; frame < 10; frame++ {
		if this.currentFrameIsSpare() {
			this.scoreSpareFrame()
		} else {
			this.scoreRegularFrame()
		}
	}
	return this.score
}

func (this *Game) currentFrameIsSpare() bool {
	return this.frameScore() == 10
}

func (this *Game) scoreSpareFrame() {
	this.score += 10 + this.pins(2)
	this.throw += 2
}

func (this *Game) scoreRegularFrame() {
	this.score += this.frameScore()
	this.throw += 2
}

func (this *Game) frameScore() int {
	return this.pins(0) + this.pins(1)
}

func (this *Game) pins(offset int) int {
	return this.throws[this.throw+offset]
}
