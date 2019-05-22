package handling

import (
	"context"

	"github.com/maxp36/rembook/handling/generated/prisma"
	"github.com/opentracing/opentracing-go"
)

type tracingService struct {
	tracer opentracing.Tracer
	Service
}

func NewTracingService(tracer opentracing.Tracer, s Service) Service {
	return &tracingService{
		tracer:  tracer,
		Service: s,
	}
}

func (s *tracingService) AddBook(ctx context.Context, name, description string) (prisma.Book, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "AddBook")
	defer span.Finish()
	return s.Service.AddBook(ctx, name, description)
}

func (s *tracingService) GetBook(ctx context.Context, id string) (prisma.Book, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "GetBook")
	defer span.Finish()
	return s.Service.GetBook(ctx, id)
}

func (s *tracingService) DeleteBook(ctx context.Context, id string) (prisma.Book, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "DeleteBook")
	defer span.Finish()
	return s.Service.DeleteBook(ctx, id)
}

func (s *tracingService) Books(ctx context.Context) ([]prisma.Book, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "Books")
	defer span.Finish()
	return s.Service.Books(ctx)
}

func (s *tracingService) AddChapter(ctx context.Context, name string, description string, bookID string) (prisma.Chapter, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "AddChapter")
	defer span.Finish()
	return s.Service.AddChapter(ctx, name, description, bookID)
}

func (s *tracingService) GetChapter(ctx context.Context, id string) (prisma.Chapter, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "GetChapter")
	defer span.Finish()
	return s.Service.GetChapter(ctx, id)
}

func (s *tracingService) DeleteChapter(ctx context.Context, id string) (prisma.Chapter, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "DeleteChapter")
	defer span.Finish()
	return s.Service.DeleteChapter(ctx, id)
}

func (s *tracingService) Chapters(ctx context.Context, bookID string) ([]prisma.Chapter, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "Chapters")
	defer span.Finish()
	return s.Service.Chapters(ctx, bookID)
}
