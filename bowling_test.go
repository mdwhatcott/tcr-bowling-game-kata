package bowling

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestGameFixture(t *testing.T) {
	gunit.Run(new(GameFixture), t)
}

type GameFixture struct {
	*gunit.Fixture
	game *Game
}

func (this *GameFixture) Setup() {
	this.game = new(Game)
}
func (this *GameFixture) assertScore(expected int) bool {
	return this.AssertEqual(expected, this.game.CalculateScore())
}
func (this *GameFixture) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		this.game.RecordRoll(pins)
	}
}
func (this *GameFixture) rollSeveral(throws ...int) {
	for _, throw := range throws {
		this.game.RecordRoll(throw)
	}
}
func (this *GameFixture) TestGutterGame() {
	this.rollMany(20, 0)
	this.assertScore(0)
}
func (this *GameFixture) TestAllOnes() {
	this.rollMany(20, 1)
	this.assertScore(20)
}
func (this *GameFixture) TestSpare() {
	this.rollSeveral(5, 5, 3)
	this.assertScore(16)
}
func (this *GameFixture) TestStrike() {
	this.rollSeveral(10, 3, 4)
	this.assertScore(24)
}
