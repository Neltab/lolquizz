package application

import (
	"context"
	"fmt"
	"time"

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

func (s *GameService) StartGame(ctx context.Context, roomCode string, hostId game.PlayerId) error {
	r, err := s.rooms.FindByCode(ctx, roomCode)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if err := r.StartGame(hostId); err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	settings := r.Settings
	questions, err := s.questions.GetQuestions(ctx, settings.QuestionCount)
	if err != nil {
		return fmt.Errorf("get questions: %w", err)
	}

	g, err := game.NewGame(game.GameId(s.idGen()), r.Id, questions, settings)
	if err != nil {
		return fmt.Errorf("create game: %w", err)
	}

	s.games[g.Id] = g
	s.roomGames[r.Id] = g.Id

	if err := s.rooms.Save(ctx, r); err != nil {
		return fmt.Errorf("save room: %w", err)
	}

	go s.handleQuestions(g, r)

	return nil
}

func (s *GameService) handleQuestions(g *game.Game, r *room.Room) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for _, q := range g.Questions {
		s.eventBus.Publish(&game.QuestionStartedEvent{
			RoomId:   r.Id,
			Question: q,
			Game:     g,
		})
		<-ticker.C
	}

}

func (s *GameService) SubmitAnswer(ctx context.Context, roomCode string, playerId game.PlayerId, answer string) error {
	r, err := s.rooms.FindByCode(ctx, roomCode)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	gameId, ok := s.roomGames[r.Id]
	if !ok {
		return fmt.Errorf("no game in this room")
	}

	g, ok := s.games[gameId]
	if !ok {
		return fmt.Errorf("no game in this room")
	}

	if err := g.SubmitAnswer(playerId, answer); err != nil {
		return fmt.Errorf("submit answer: %w", err)
	}

	return nil
}

func (s *GameService) ReviewQuestion(ctx context.Context, roomCode string, playerId game.PlayerId) error {
	return nil
}

func (s *GameService) SubmitReview(ctx context.Context, roomCode string, playerId game.PlayerId, correct bool) error {
	return nil
}
