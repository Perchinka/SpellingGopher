package quote

import (
	"context"
	"errors"
	"testing"
)

type fakeRepo struct {
	quote Quote
	err   error
	calls int
}

func (f *fakeRepo) Random(ctx context.Context) (Quote, error) {
	f.calls++
	return f.quote, f.err
}

func TestService_Random(t *testing.T) {
	tests := []struct {
		name    string
		repo    *fakeRepo
		want    Quote
		wantErr bool
	}{
		{
			name: "returns quote from repo",
			repo: &fakeRepo{quote: Quote{Quote: "be water", Author: "Lee"}},
			want: Quote{Quote: "be water", Author: "Lee"},
		},
		{
			name:    "propagates repo error",
			repo:    &fakeRepo{err: errors.New("boom")},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewService(tt.repo)
			got, err := svc.Random(context.Background())
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}
