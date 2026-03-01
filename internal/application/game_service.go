package application

import (
	"context"
	"fmt"

	"lolquizz/internal/domain/event"
	"lolquizz/internal/domain/game"
	"lolquizz/internal/domain/room"
)

type EventPublisher interface {
	PublishToRoom(roomId room.RoomId, event event.Event)
	PublishToPlayer(playerId room.PlayerId, event event.Event)
}

type QuestionProvider interface {
	GetQuestions(ctx context.Context, count int, difficulty string) ([]*game.Question, error)
}

type GameService struct {
	rooms     room.Repository
	games     map[game.GameId]*game.Game
	roomGames map[room.RoomId]game.GameId
	events    EventPublisher
	questions QuestionProvider
	idGen     func() string
}

func NewGameService(rooms room.Repository, events EventPublisher, questions QuestionProvider, idGen func() string) *GameService {
	return &GameService{
		rooms:     rooms,
		games:     make(map[game.GameId]*game.Game),
		roomGames: make(map[room.RoomId]game.GameId),
		events:    events,
		questions: questions,
		idGen:     idGen,
	}
}

func (s *GameService) StartGame(ctx context.Context, roomId room.RoomId, hostId room.PlayerId) error {
	r, err := s.rooms.FindById(ctx, roomId)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if err := r.StartGame(hostId); err != nil {
		return err
	}

	settings := r.Settings
	questions, err := s.questions.GetQuestions(ctx, settings.QuestionCount, settings.Difficulty)
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

	question := g.CurrentQuestion()
	s.events.PublishToRoom(roomId, &game.QuestionStartedEvent{
		RoomId:       roomId,
		QuestionText: question.Text,
		Duration:     question.Duration,
	})

	return nil
}
