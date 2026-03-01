package game

import "errors"

var (
	ErrNoQuestions         = errors.New("No questions")
	ErrNoAnswer            = errors.New("No answer")
	ErrNotInQuestionsPhase = errors.New("Not in questions phase")
	ErrNotInAnswersPhase   = errors.New("Not in answers phase")
)
