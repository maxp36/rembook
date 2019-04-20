package handling

import (
	"context"
	"errors"

	"github.com/maxp36/rembook/handling/generated/prisma"
)

// ErrInvalidArgument is returned if method's argument are invalid
var ErrInvalidArgument = errors.New("invalid argument")

// Service is th instance of the Handling service
type Service interface {
	AddBook(ctx context.Context, name string, description string) (prisma.Book, error)
	GetBook(ctx context.Context, id string) (prisma.Book, error)
	DeleteBook(ctx context.Context, id string) (prisma.Book, error)
	Books(ctx context.Context) ([]prisma.Book, error)

	AddChapter(ctx context.Context, name string, description string, bookID string) (prisma.Chapter, error)
	GetChapter(ctx context.Context, id string) (prisma.Chapter, error)
	DeleteChapter(ctx context.Context, id string) (prisma.Chapter, error)
	Chapters(ctx context.Context, bookID string) ([]prisma.Chapter, error)
}

type service struct{}

var client = prisma.New(nil)

func (s *service) AddBook(ctx context.Context, name, description string) (prisma.Book, error) {
	if name == "" || description == "" {
		return prisma.Book{}, ErrInvalidArgument
	}

	book, err := client.CreateBook(prisma.BookCreateInput{
		Name:        name,
		Description: description,
	}).Exec(ctx)

	if err != nil {
		return prisma.Book{}, err
	}

	return *book, nil
}

func (s *service) GetBook(ctx context.Context, id string) (prisma.Book, error) {
	if id == "" {
		return prisma.Book{}, ErrInvalidArgument
	}

	book, err := client.Book(prisma.BookWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		return prisma.Book{}, err
	}

	return *book, nil
}

func (s *service) DeleteBook(ctx context.Context, id string) (prisma.Book, error) {
	if id == "" {
		return prisma.Book{}, ErrInvalidArgument
	}

	book, err := client.DeleteBook(prisma.BookWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		return prisma.Book{}, err
	}

	return *book, nil
}

func (s *service) Books(ctx context.Context) ([]prisma.Book, error) {
	books, err := client.Books(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *service) AddChapter(ctx context.Context, name string, description string, bookID string) (prisma.Chapter, error) {
	if name == "" || description == "" || bookID == "" {
		return prisma.Chapter{}, ErrInvalidArgument
	}

	chapter, err := client.CreateChapter(prisma.ChapterCreateInput{
		Name:        name,
		Description: description,
		Book: prisma.BookCreateOneWithoutChaptersInput{
			Connect: &prisma.BookWhereUniqueInput{
				ID: &bookID,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return prisma.Chapter{}, err
	}

	return *chapter, nil
}

func (s *service) GetChapter(ctx context.Context, id string) (prisma.Chapter, error) {
	if id == "" {
		return prisma.Chapter{}, ErrInvalidArgument
	}

	chapter, err := client.Chapter(prisma.ChapterWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		return prisma.Chapter{}, err
	}

	return *chapter, nil
}

func (s *service) DeleteChapter(ctx context.Context, id string) (prisma.Chapter, error) {
	if id == "" {
		return prisma.Chapter{}, ErrInvalidArgument
	}

	chapter, err := client.DeleteChapter(prisma.ChapterWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		return prisma.Chapter{}, err
	}

	return *chapter, nil
}

func (s *service) Chapters(ctx context.Context, bookID string) ([]prisma.Chapter, error) {
	if bookID == "" {
		return nil, ErrInvalidArgument
	}

	chapters, err := client.Book(prisma.BookWhereUniqueInput{
		ID: &bookID,
	}).Chapters(nil).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return chapters, nil
}

// NewService returns the new instance of Handling service
func NewService() Service {
	return &service{}
}
