package application

import (
	"context"
	"fmt"

	"lolquizz/internal/domain/event"
	"lolquizz/internal/domain/game"
	"lolquizz/internal/domain/room"
)

type QuestionProvider interface {
	GetQuestions(ctx context.Context, count int) ([]*game.Question, error)
}

type GameService struct {
	rooms     room.Repository
	games     map[game.GameId]*game.Game
	roomGames map[game.RoomId]game.GameId
	eventBus  event.Publisher
	questions QuestionProvider
	idGen     func() string
}

func NewGameService(rooms room.Repository, events event.Publisher, questions QuestionProvider, idGen func() string) *GameService {
	return &GameService{
		rooms:     rooms,
		games:     make(map[game.GameId]*game.Game),
		roomGames: make(map[game.RoomId]game.GameId),
		eventBus:  events,
		questions: questions,
		idGen:     idGen,
	}
}

func (s *GameService) StartGame(ctx context.Context, roomId game.RoomId, hostId game.PlayerId) error {
	r, err := s.rooms.FindById(ctx, roomId)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if err := r.StartGame(hostId); err != nil {
		return err
	}

	settings := r.Settings
	questions, err := s.questions.GetQuestions(ctx, settings.QuestionCount)
	if err != nil {
		return fmt.Errorf("get questions: %w", err)
	}

	g, err := game.NewGame(game.GameId(s.idGen()), roomId, questions, settings)
	if err != nil {
		return fmt.Errorf("create game: %w", err)
	}

	s.games[g.Id] = g
	s.roomGames[roomId] = g.Id

	if err := s.rooms.Save(ctx, r); err != nil {
		return fmt.Errorf("save room: %w", err)
	}

	//TODO rework
	// question := g.CurrentQuestion()
	// s.eventBus.PublishToRoom(roomId, &game.QuestionStartedEvent{
	// 	RoomId:       roomId,
	// 	QuestionText: question.Text,
	// 	Duration:     question.Duration,
	// })

	return nil
}
