package handling

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHandler makes handlers for all service methods and sets routes.
func MakeHandler(s Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	addBookHandler := kithttp.NewServer(
		makeAddBookEndpoint(s),
		decodeAddBookRequest,
		encodeResponse,
		opts...,
	)
	getBookHandler := kithttp.NewServer(
		makeGetBookEndpoint(s),
		decodeGetBookRequest,
		encodeResponse,
		opts...,
	)
	deleteBookHandler := kithttp.NewServer(
		makeDeleteBookEndpoint(s),
		decodeDeleteBookRequest,
		encodeResponse,
		opts...,
	)
	listBooksHandler := kithttp.NewServer(
		makeListBooksEndpoint(s),
		decodeListBooksRequest,
		encodeResponse,
		opts...,
	)

	addChapterHandler := kithttp.NewServer(
		makeAddChapterEndpoint(s),
		decodeAddChapterRequest,
		encodeResponse,
		opts...,
	)
	getChapterHandler := kithttp.NewServer(
		makeGetChapterEndpoint(s),
		decodeGetChapterRequest,
		encodeResponse,
		opts...,
	)
	deleteChapterHandler := kithttp.NewServer(
		makeDeleteChapterEndpoint(s),
		decodeDeleteChapterRequest,
		encodeResponse,
		opts...,
	)
	listChaptersHandler := kithttp.NewServer(
		makeListChaptersEndpoint(s),
		decodeListChaptersRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/handling/v1/books", addBookHandler).Methods("POST")
	r.Handle("/handling/v1/books", listBooksHandler).Methods("GET")
	r.Handle("/handling/v1/books/{id}", getBookHandler).Methods("GET")
	r.Handle("/handling/v1/books/{id}", deleteBookHandler).Methods("DELETE")

	r.Handle("/handling/v1/books/{book_id}/chapters", addChapterHandler).Methods("POST")
	r.Handle("/handling/v1/books/{book_id}/chapters", listChaptersHandler).Methods("GET")
	r.Handle("/handling/v1/books/{book_id}/chapters/{id}", getChapterHandler).Methods("GET")
	r.Handle("/handling/v1/books/{book_id}/chapters/{id}", deleteChapterHandler).Methods("DELETE")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeAddBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return addBookRequest{
		Name:        body.Name,
		Description: body.Description,
	}, nil
}

func decodeGetBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return getBookRequest{ID: id}, nil
}

func decodeDeleteBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return deleteBookRequest{ID: id}, nil
}

func decodeListBooksRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listBooksRequest{}, nil
}

func decodeAddChapterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	bookID, ok := vars["book_id"]
	if !ok {
		return nil, errBadRoute
	}

	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return addChapterRequest{
		Name:        body.Name,
		Description: body.Description,
		BookID:      bookID,
	}, nil
}

func decodeGetChapterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}

	return getChapterRequest{ID: id}, nil
}

func decodeDeleteChapterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}

	return deleteChapterRequest{ID: id}, nil
}

func decodeListChaptersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	bookID, ok := vars["book_id"]
	if !ok {
		return nil, errBadRoute
	}

	return listChaptersRequest{BookID: bookID}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// EncodeError encodes errors from business-logic.
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
