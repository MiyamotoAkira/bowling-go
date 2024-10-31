package bowling_structs_v3

const MAX_FRAME_SCORE = 10
const FIRST_ROLL = 0
const SECOND_ROLL = 1
const THIRD_ROLL = 2
const MAX_FRAMES = 10

type standardRoll struct {
	score  int
	scored bool
}

func (roll *standardRoll) score_pins(score int) {
	roll.score = score
	roll.scored = true
}

type standardFrame struct {
	rolls        [3]standardRoll
	isFinalFrame bool
	bonusScore   [2]standardRoll
}

var EMPTY_FRAME standardFrame

func (frame *standardFrame) isFullFrame() bool {
	if frame.isFinalFrame {
		return frame.rolls[THIRD_ROLL].scored
	} else {
		return frame.isStrike() || frame.rolls[SECOND_ROLL].scored
	}
}

func (frame *standardFrame) isStrike() bool {
	return frame.rolls[FIRST_ROLL].score == MAX_FRAME_SCORE
}

func (frame *standardFrame) isSpare() bool {
	return !frame.isStrike() && (frame.rolls[FIRST_ROLL].score+frame.rolls[SECOND_ROLL].score == MAX_FRAME_SCORE)
}

func (frame *standardFrame) rollPins(score int) {
	var roll_number = FIRST_ROLL
	if frame.isFinalFrame {
		if frame.rolls[SECOND_ROLL].scored {
			roll_number = THIRD_ROLL
		} else if frame.rolls[FIRST_ROLL].scored {
			roll_number = SECOND_ROLL
		}

	} else {
		if frame.rolls[FIRST_ROLL].scored {
			roll_number = SECOND_ROLL
		}
	}

	frame.rolls[roll_number].score_pins(score)
}

func (frame *standardFrame) score() int {
	total := 0
	total += frame.rolls[FIRST_ROLL].score
	total += frame.rolls[SECOND_ROLL].score
	if frame.isFinalFrame {
		total += frame.rolls[THIRD_ROLL].score
	}
	total += frame.bonusScore[FIRST_ROLL].score
	total += frame.bonusScore[SECOND_ROLL].score

	return total
}

func (frame *standardFrame) addBonus(score int) {
	if frame.isStrike() {
		if !frame.bonusScore[FIRST_ROLL].scored {
			frame.bonusScore[FIRST_ROLL].score_pins(score)
		} else if !frame.bonusScore[SECOND_ROLL].scored {
			frame.bonusScore[SECOND_ROLL].score_pins(score)
		}
	}

	if frame.isSpare() {
		if !frame.bonusScore[FIRST_ROLL].scored {
			frame.bonusScore[FIRST_ROLL].score_pins(score)
		}
	}
}

type game struct {
	frames              [MAX_FRAMES]standardFrame
	currentFrameIndex   int
	currentFrame        *standardFrame
	previousFrame       *standardFrame
	previousToLastFrame *standardFrame
}

func (bowlingGame *game) startGame() {
	bowlingGame.frames = [MAX_FRAMES]standardFrame{}
	bowlingGame.frames[MAX_FRAMES-1].isFinalFrame = true
	bowlingGame.currentFrameIndex = 0
	bowlingGame.currentFrame = &bowlingGame.frames[bowlingGame.currentFrameIndex]
	bowlingGame.previousFrame = &EMPTY_FRAME
	bowlingGame.previousToLastFrame = &EMPTY_FRAME
}

func (bowlingGame *game) rollPins(score int) {
	bowlingGame.currentFrame.rollPins(score)
	bowlingGame.previousFrame.addBonus(score)
	bowlingGame.previousToLastFrame.addBonus(score)

	if bowlingGame.currentFrame.isFullFrame() {
		bowlingGame.selectFrames()
	}
}

func (bowlingGame *game) selectFrames() {
	if !bowlingGame.currentFrame.isFinalFrame {
		bowlingGame.currentFrameIndex++
		bowlingGame.currentFrame = &bowlingGame.frames[bowlingGame.currentFrameIndex]
		if bowlingGame.currentFrameIndex-1 >= 0 {
			bowlingGame.previousFrame = &bowlingGame.frames[bowlingGame.currentFrameIndex-1]
		}
		if bowlingGame.currentFrameIndex-2 >= 0 {
			bowlingGame.previousToLastFrame = &bowlingGame.frames[bowlingGame.currentFrameIndex-2]
		}
	}
}

func (bowlingGame *game) scoreGame() int {
	total := 0
	for _, v := range bowlingGame.frames {
		total += v.score()
	}

	return total
}

func Roll(score int) {
	bowlingGame.rollPins(score)
}

func Score() int {
	return bowlingGame.scoreGame()
}

var bowlingGame game

func StartNewGame() {
	bowlingGame.startGame()
}
