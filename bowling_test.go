package bowling_test

import (
	"bowling"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SingleRoll(t *testing.T) {
	bowling.Clear()
	bowling.Roll(5)
	assert.Equal(t, 5, bowling.Score())
}

func Test_TwoRolls(t *testing.T) {
	bowling.Clear()
	bowling.Roll(5)
	bowling.Roll(5)
	assert.Equal(t, 10, bowling.Score())
}

func Test_WithStrikeAndFollowUpRolls(t *testing.T) {
	bowling.Clear()
	bowling.Roll(10)
	bowling.Roll(4)
	bowling.Roll(3)
	assert.Equal(t, 24, bowling.Score())
}

func Test_WithStrikeAndWithSingleFollowUpRoll(t *testing.T) {
	bowling.Clear()
	bowling.Roll(10)
	bowling.Roll(4)
	assert.Equal(t, 18, bowling.Score())
}

func Test_WithStrikeAndWithNoFollowUPRoll(t *testing.T) {
	bowling.Clear()
	bowling.Roll(10)
	assert.Equal(t, 10, bowling.Score())
}
