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
