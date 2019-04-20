package handling

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/maxp36/rembook/handling/generated/prisma"
)

type addBookRequest struct {
	Name        string
	Description string
}

type addBookResponse struct {
	Book prisma.Book `json:"book,omitempty"`
	Err  error       `json:"err,omitempty"`
}

func (r addBookResponse) error() error { return r.Err }

func makeAddBookEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addBookRequest)
		book, err := s.AddBook(ctx, req.Name, req.Description)
		return addBookResponse{Book: book, Err: err}, nil
	}
}

type getBookRequest struct {
	ID string `json:"id"`
}

type getBookResponse struct {
	Book prisma.Book `json:"book,omitempty"`
	Err  error       `json:"err,omitempty"`
}

func (r getBookResponse) error() error { return r.Err }

func makeGetBookEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getBookRequest)
		book, err := s.GetBook(ctx, req.ID)
		return getBookResponse{Book: book, Err: err}, nil
	}
}

type deleteBookRequest struct {
	ID string `json:"id"`
}

type deleteBookResponse struct {
	Book prisma.Book `json:"book,omitempty"`
	Err  error       `json:"err,omitempty"`
}

func (r deleteBookResponse) error() error { return r.Err }

func makeDeleteBookEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteBookRequest)
		book, err := s.DeleteBook(ctx, req.ID)
		return deleteBookResponse{Book: book, Err: err}, nil
	}
}

type listBooksRequest struct{}

type listBooksResponse struct {
	Books []prisma.Book `json:"books,omitempty"`
	Err   error         `json:"err,omitempty"`
}

func (r listBooksResponse) error() error { return r.Err }

func makeListBooksEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listBooksRequest)
		books, err := s.Books(ctx)
		return listBooksResponse{Books: books, Err: err}, nil
	}
}

type addChapterRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	BookID      string `json:"book_id"`
}

type addChapterResponse struct {
	Chapter prisma.Chapter `json:"chapter,omitempty"`
	Err     error          `json:"err,omitempty"`
}

func (r addChapterResponse) error() error { return r.Err }

func makeAddChapterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addChapterRequest)
		chapter, err := s.AddChapter(ctx, req.Name, req.Description, req.BookID)
		return addChapterResponse{Chapter: chapter, Err: err}, nil
	}
}

type getChapterRequest struct {
	ID string `json:"id"`
}

type getChapterResponse struct {
	Chapter prisma.Chapter `json:"chapter,omitempty"`
	Err     error          `json:"err,omitempty"`
}

func (r getChapterResponse) error() error { return r.Err }

func makeGetChapterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getChapterRequest)
		chapter, err := s.GetChapter(ctx, req.ID)
		return getChapterResponse{Chapter: chapter, Err: err}, nil
	}
}

type deleteChapterRequest struct {
	ID string `json:"id"`
}

type deleteChapterResponse struct {
	Chapter prisma.Chapter `json:"chapter,omitempty"`
	Err     error          `json:"err,omitempty"`
}

func (r deleteChapterResponse) error() error { return r.Err }

func makeDeleteChapterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteChapterRequest)
		chapter, err := s.DeleteChapter(ctx, req.ID)
		return deleteChapterResponse{Chapter: chapter, Err: err}, nil
	}
}

type listChaptersRequest struct {
	BookID string `json:"book_id"`
}

type listChaptersResponse struct {
	Chapters []prisma.Chapter `json:"chapters,omitempty"`
	Err      error            `json:"err,omitempty"`
}

func (r listChaptersResponse) error() error { return r.Err }

func makeListChaptersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listChaptersRequest)
		chapters, err := s.Chapters(ctx, req.BookID)
		return listChaptersResponse{Chapters: chapters, Err: err}, nil
	}
}
