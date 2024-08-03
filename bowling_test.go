package bowling_test

import (
	"bowling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BowlingTestSuite struct {
	suite.Suite
}

func (suite *BowlingTestSuite) SetupTest() {
	bowling.StartNewGame()
}

func TestBowling(t *testing.T) {
	suite.Run(t, new(BowlingTestSuite))
}

func (suite *BowlingTestSuite) Test_SingleRoll() {
	bowling.Roll(5)
	assert.Equal(suite.T(), 5, bowling.Score())
}

func (suite *BowlingTestSuite) Test_TwoRolls() {
	bowling.Roll(5)
	bowling.Roll(5)
	assert.Equal(suite.T(), 10, bowling.Score())
}

func (suite *BowlingTestSuite) Test_WithStrikeAndFollowUpRolls() {
	bowling.Roll(10)
	bowling.Roll(4)
	bowling.Roll(3)
	assert.Equal(suite.T(), 24, bowling.Score())
}

func (suite *BowlingTestSuite) Test_WithStrikeAndWithSingleFollowUpRoll() {
	bowling.Roll(10)
	bowling.Roll(4)
	assert.Equal(suite.T(), 18, bowling.Score())
}

func (suite *BowlingTestSuite) Test_WithStrikeAndWithNoFollowUPRoll() {
	bowling.Roll(10)
	assert.Equal(suite.T(), 10, bowling.Score())
}
