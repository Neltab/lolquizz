package memory

import (
	"context"
	"lolquizz/internal/domain/game"
	"sync"
)

type QuestionRepository struct {
	questions map[string]*game.Question
	mu        sync.RWMutex
}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{
		questions: make(map[string]*game.Question),
	}
}

func (r *QuestionRepository) GetQuestions(ctx context.Context, count int) ([]*game.Question, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	questions := make([]*game.Question, 0, count)
	for _, q := range r.questions {
		questions = append(questions, q)
	}
	return questions, nil
}
