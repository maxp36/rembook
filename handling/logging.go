package handling

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/maxp36/rembook/handling/generated/prisma"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) AddBook(ctx context.Context, name string, description string) (book prisma.Book, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "add_book",
			"name", name,
			"description", description,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.AddBook(ctx, name, description)
}

func (s *loggingService) GetBook(ctx context.Context, id string) (book prisma.Book, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "get_book",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetBook(ctx, id)
}

func (s *loggingService) DeleteBook(ctx context.Context, id string) (book prisma.Book, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "delete_book",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteBook(ctx, id)
}

func (s *loggingService) Books(ctx context.Context) (books []prisma.Book, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "list_books",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.Books(ctx)
}

func (s *loggingService) AddChapter(ctx context.Context, name string, description string, bookID string) (chapter prisma.Chapter, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "add_chapter",
			"name", name,
			"description", description,
			"bookID", bookID,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.AddChapter(ctx, name, description, bookID)
}

func (s *loggingService) GetChapter(ctx context.Context, id string) (chapter prisma.Chapter, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "get_chapter",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetChapter(ctx, id)
}

func (s *loggingService) DeleteChapter(ctx context.Context, id string) (chapter prisma.Chapter, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "delete_chapter",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteChapter(ctx, id)
}

func (s *loggingService) Chapters(ctx context.Context, bookID string) (chapters []prisma.Chapter, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "list_chapters",
			"bookID", bookID,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.Chapters(ctx, bookID)
}
