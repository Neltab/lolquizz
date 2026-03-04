package game

import "errors"

var (
	ErrNoQuestions         = errors.New("no questions")
	ErrNoAnswer            = errors.New("no answer")
	ErrNotInQuestionsPhase = errors.New("not in questions phase")
	ErrNotInAnswersPhase   = errors.New("not in answers phase")
)
