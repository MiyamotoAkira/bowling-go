package bowling_structs

const MAX_FRAME_SCORE = 10
const FIRST_ROLL = 0
const SECOND_ROLL = 1
const THIRD_ROLL = 2

type standardRoll struct {
	score  int
	scored bool
}

func (roll *standardRoll) isScored() bool {
	return roll.scored
}

func (roll *standardRoll) scorePins(score int) {
	roll.score = score
	roll.scored = true
}

type frame interface {
	rollPins(score int)
	isFullFrame() bool
	isStrike() bool
	isSpare() bool
	frameScore() int
	frameScoreV2(nextFrame frame, nextNextFrame frame) int
	firstRollScore() int
}

type standardFrame struct {
	rolls [2]standardRoll
}

type finalFrame struct {
	rolls [3]standardRoll
}

func (frame *standardFrame) isFullFrame() bool {
	return frame.isStrike() || (frame.rolls[FIRST_ROLL].isScored() && frame.rolls[SECOND_ROLL].isScored())
}

func (frame *finalFrame) isFullFrame() bool {
	return frame.isStrike() || (frame.rolls[FIRST_ROLL].isScored() && frame.rolls[SECOND_ROLL].isScored())
}

func (frame *standardFrame) isStrike() bool {
	return frame.rolls[FIRST_ROLL].score == MAX_FRAME_SCORE
}

func (frame *finalFrame) isStrike() bool {
	return frame.rolls[FIRST_ROLL].score == MAX_FRAME_SCORE
}

func (frame *finalFrame) firstRollScore() int {
	return frame.rolls[FIRST_ROLL].score
}

func (frame *standardFrame) firstRollScore() int {
	return frame.rolls[FIRST_ROLL].score
}

func (frame *standardFrame) isSpare() bool {
	return !frame.isStrike() && (frame.rolls[FIRST_ROLL].score+frame.rolls[SECOND_ROLL].score == MAX_FRAME_SCORE)
}

func (frame *finalFrame) isSpare() bool {
	return !frame.isStrike() && (frame.rolls[FIRST_ROLL].score+frame.rolls[SECOND_ROLL].score == MAX_FRAME_SCORE)
}

func (frame *standardFrame) rollPins(score int) {
	var rollNumber = FIRST_ROLL
	if frame.rolls[FIRST_ROLL].isScored() {
		rollNumber = SECOND_ROLL
	}

	frame.rolls[rollNumber].scorePins(score)
}

func (frame *finalFrame) rollPins(score int) {
	var rollNumber = FIRST_ROLL
	if frame.rolls[SECOND_ROLL].isScored() {
		rollNumber = THIRD_ROLL
	} else if frame.rolls[FIRST_ROLL].isScored() {
		rollNumber = SECOND_ROLL
	}

	frame.rolls[rollNumber].scorePins(score)
}

func (frame *standardFrame) frameScore() int {
	return frame.rolls[FIRST_ROLL].score + frame.rolls[SECOND_ROLL].score
}

func (frame *finalFrame) frameScore() int {
	return frame.rolls[FIRST_ROLL].score + frame.rolls[SECOND_ROLL].score
}

func (frame *finalFrame) fullFrameScore() int {
	return frame.rolls[FIRST_ROLL].score + frame.rolls[SECOND_ROLL].score + frame.rolls[THIRD_ROLL].score
}

func (frame *standardFrame) frameScoreV2(nextFrame frame, nextNextFrame frame) int {
	total := frame.rolls[FIRST_ROLL].score + frame.rolls[SECOND_ROLL].score
	if frame.isStrike() {
		if nextFrame != nil {
			total += nextFrame.frameScore()

			if nextFrame.isStrike() && nextNextFrame != nil {
				total += nextNextFrame.firstRollScore()
			}
		}
	}

	if frame.isSpare() && nextFrame != nil {
		total += nextFrame.firstRollScore()
	}

	return total
}

func (frame *finalFrame) frameScoreV2(nextFrame frame, nextNextFrame frame) int {
	return frame.rolls[FIRST_ROLL].score + frame.rolls[SECOND_ROLL].score + frame.rolls[THIRD_ROLL].score
}

type game struct {
	frames       [9]standardFrame
	finalFrame   finalFrame
	currentFrame int
}

func (bowlingGame *game) startGame() {
	bowlingGame.frames = [9]standardFrame{}
	bowlingGame.finalFrame = finalFrame{}
	bowlingGame.currentFrame = 0
}

func (bowlingGame *game) rollPins(score int) {
	if bowlingGame.currentFrame > 8 {
		bowlingGame.finalFrame.rollPins(score)
	} else {
		bowlingGame.frames[bowlingGame.currentFrame].rollPins(score)

		if bowlingGame.frames[bowlingGame.currentFrame].isFullFrame() {
			bowlingGame.currentFrame++
		}
	}
}

func (bowlingGame *game) scoreGame() int {
	total := 0
	for i, v := range bowlingGame.frames {

		var nextFrame frame
		var nextNextFrame frame
		if i == 8 {
			nextFrame = &bowlingGame.finalFrame
		} else if i == 7 {
			nextFrame = &bowlingGame.frames[i+1]
			nextNextFrame = &bowlingGame.finalFrame

		} else {
			nextFrame = &bowlingGame.frames[i+1]
			nextNextFrame = &bowlingGame.frames[i+2]
		}
		total += v.frameScoreV2(nextFrame, nextNextFrame)
	}

	total += bowlingGame.finalFrame.fullFrameScore()

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
