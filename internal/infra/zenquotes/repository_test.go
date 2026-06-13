package zenquotes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"perchinka.github.io/spelling-gopher/internal/quote"
)

func TestRepository_Random(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		body    string
		want    quote.Quote
		wantErr error
	}{
		{
			name:   "valid single quote",
			status: 200,
			body:   `[{"q":"be water","a":"Bruce Lee","c":"8"}]`,
			want:   quote.Quote{Text: "be water", Author: "Bruce Lee"},
		},
		{name: "empty array", status: 200, body: `[]`, wantErr: ErrEmptyResponse},
		{name: "empty body", status: 200, body: ``, wantErr: ErrInvalidResponse},
		{name: "rate limited", status: 429, body: `Too Many Requests`, wantErr: ErrUnexpectedStatus},
		{name: "malformed json", status: 200, body: `{not json`, wantErr: ErrInvalidResponse},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
				fmt.Fprint(w, tt.body)
			}))
			defer srv.Close()

			repo := New(srv.Client())
			repo.baseURL = srv.URL

			got, err := repo.Random(context.Background())
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if tt.wantErr == nil && got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}

}

// TestRepository_RequestShape verifies the adapter issues a GET to the
// "/random" path (guards the baseURL + "/random" construction).
func TestRepository_RequestShape(t *testing.T) {
	var gotMethod, gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		fmt.Fprint(w, `[{"q":"be water","a":"Bruce Lee"}]`)
	}))
	defer srv.Close()

	repo := New(srv.Client())
	repo.baseURL = srv.URL

	if _, err := repo.Random(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/random" {
		t.Errorf("path = %q, want %q", gotPath, "/random")
	}
}

func TestRepository_TransportError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	client := srv.Client()
	url := srv.URL
	srv.Close() // server is unreachable

	repo := New(client)
	repo.baseURL = url

	_, err := repo.Random(context.Background())
	if err == nil {
		t.Fatal("expected a transport error, got nil")
	}
	for _, sentinel := range []error{ErrUnexpectedStatus, ErrEmptyResponse, ErrInvalidResponse} {
		if errors.Is(err, sentinel) {
			t.Errorf("transport error should not match sentinel %v, got %v", sentinel, err)
		}
	}
}

func TestRepository_ContextCanceled(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"q":"be water","a":"Bruce Lee"}]`)
	}))
	defer srv.Close()

	repo := New(srv.Client())
	repo.baseURL = srv.URL

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := repo.Random(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
}
