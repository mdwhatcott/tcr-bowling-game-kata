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
	return this.AssertEqual(expected, this.game.Score())
}
func (this *GameFixture) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		this.game.Roll(pins)
	}
}
func (this *GameFixture) TestGutterGame() {
	this.rollMany(20, 0)
	this.assertScore(0)
}
