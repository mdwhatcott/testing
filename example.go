package testing

type bowling struct {
	score  int
	throw  int
	throws [21]int
}

func (this *bowling) recordRoll(pins int) {
	this.throws[this.throw] = pins
	this.throw++
}

func (this *bowling) calculateScore() int {
	this.throw = 0
	this.score = 0
	for frame := 0; frame < 10; frame++ {
		this.score += this.scoreCurrentFrame()
		this.throw += this.advanceFrame()
	}
	return this.score
}
func (this *bowling) scoreCurrentFrame() int {
	if this.currentFrameIsStrike() {
		return this.scoreStrikeFrame()
	} else if this.currentFrameIsSpare() {
		return this.scoreSpareFrame()
	} else {
		return this.scoreRegularFrame()
	}
}
func (this *bowling) currentFrameIsStrike() bool {
	return this.pins(0) == 10
}
func (this *bowling) currentFrameIsSpare() bool {
	return this.frameScore() == 10
}
func (this *bowling) scoreStrikeFrame() int {
	return 10 + this.pins(1) + this.pins(2)
}
func (this *bowling) scoreSpareFrame() int {
	return 10 + this.pins(2)
}
func (this *bowling) scoreRegularFrame() int {
	return this.frameScore()
}
func (this *bowling) frameScore() int {
	return this.pins(0) + this.pins(1)
}
func (this *bowling) pins(offset int) int {
	return this.throws[this.throw+offset]
}
func (this *bowling) advanceFrame() int {
	if this.currentFrameIsStrike() {
		return 1
	} else {
		return 2
	}
}
