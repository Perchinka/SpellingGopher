package zenquotes

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"perchinka.github.io/spelling-gopher/internal/domain/quote"
)

func TestRepository_Random(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		body    string
		want    quote.Quote
		wantErr bool
	}{
		{
			name:   "valid single quote",
			status: 200,
			body:   `[{"q":"be water","a":"Bruce Lee","c":"8"}]`,
			want:   quote.Quote{Quote: "be water", Author: "Bruce Lee", CharacterCount: 8},
		},
		{name: "empty array", status: 200, body: `[]`, wantErr: true},
		{name: "rate limited", status: 429, body: `Too Many Requests`, wantErr: true},
		{name: "malformed json", status: 200, body: `{not json`, wantErr: true},
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
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}

}
