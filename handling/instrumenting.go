package handling

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/maxp36/rembook/handling/generated/prisma"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) AddBook(ctx context.Context, name string, description string) (prisma.Book, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "add_book").Add(1)
		s.requestLatency.With("method", "add_book").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.AddBook(ctx, name, description)
}

func (s *instrumentingService) GetBook(ctx context.Context, id string) (prisma.Book, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "get_book").Add(1)
		s.requestLatency.With("method", "get_book").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetBook(ctx, id)
}

func (s *instrumentingService) DeleteBook(ctx context.Context, id string) (prisma.Book, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "delete_book").Add(1)
		s.requestLatency.With("method", "delete_book").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteBook(ctx, id)
}

func (s *instrumentingService) Books(ctx context.Context) ([]prisma.Book, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "list_books").Add(1)
		s.requestLatency.With("method", "list_books").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Books(ctx)
}

func (s *instrumentingService) AddChapter(ctx context.Context, name string, description string, bookID string) (prisma.Chapter, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "add_chapter").Add(1)
		s.requestLatency.With("method", "add_chapter").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.AddChapter(ctx, name, description, bookID)
}

func (s *instrumentingService) GetChapter(ctx context.Context, id string) (prisma.Chapter, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "get_chapter").Add(1)
		s.requestLatency.With("method", "get_chapter").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetChapter(ctx, id)
}

func (s *instrumentingService) DeleteChapter(ctx context.Context, id string) (prisma.Chapter, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "delete_chapter").Add(1)
		s.requestLatency.With("method", "delete_chapter").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteChapter(ctx, id)
}

func (s *instrumentingService) Chapters(ctx context.Context, bookID string) ([]prisma.Chapter, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "list_chapters").Add(1)
		s.requestLatency.With("method", "list_chapters").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Chapters(ctx, bookID)
}
