package game

import (
	"time"

	"lolquizz/internal/domain/room"
)

type GamePhase int

const (
	PhaseQuestions GamePhase = iota
	PhaseAnswers
	PhaseResults
	PhaseFinished
)

type Game struct {
	Id              GameId
	RoomId          RoomId
	Phase           GamePhase
	Questions       []*Question
	Scores          map[PlayerId]int
	currentQuestion int
}

func NewGame(id GameId, roomId RoomId, questions []*Question, settings *room.Settings) (*Game, error) {
	if len(questions) == 0 {
		return nil, ErrNoQuestions
	}

	return &Game{
		Id:              id,
		RoomId:          roomId,
		Phase:           PhaseQuestions,
		Questions:       questions,
		Scores:          make(map[PlayerId]int),
		currentQuestion: 0,
	}, nil
}

func (g *Game) CurrentQuestion() *Question {
	if g.currentQuestion >= len(g.Questions) {
		return nil
	}
	return g.Questions[g.currentQuestion]
}

func (g *Game) NextRound() {
	g.currentQuestion++
	if g.currentQuestion >= len(g.Questions) {
		g.Phase = PhaseAnswers
		return
	}
}

func (g *Game) SubmitAnswer(playerId PlayerId, answer string) error {
	if g.Phase != PhaseQuestions {
		return ErrNotInQuestionsPhase
	}

	question := g.CurrentQuestion()
	if question == nil {
		return ErrNoQuestions
	}

	question.Answer(playerId, answer, time.Now())
	return nil
}

func (g *Game) JudgeAnswer(playerId PlayerId, correct bool) error {
	if g.Phase != PhaseAnswers {
		return ErrNotInAnswersPhase
	}

	question := g.CurrentQuestion()
	if question == nil {
		return ErrNoQuestions
	}
	answer := question.GetAnswer(playerId)
	if answer == nil {
		return ErrNoAnswer
	}

	if correct {
		g.Scores[playerId] += answer.Points()
	}
	return nil
}
