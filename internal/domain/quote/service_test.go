package quote

import (
	"context"
	"errors"
	"testing"
)

type fakeRepo struct {
	quote  Quote
	err    error
	calls  int
	gotCtx context.Context
}

func (f *fakeRepo) Random(ctx context.Context) (Quote, error) {
	f.calls++
	f.gotCtx = ctx
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
			repo: &fakeRepo{quote: Quote{Text: "be water", Author: "Lee"}},
			want: Quote{Text: "be water", Author: "Lee"},
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

func TestService_Random_CallsRepoOnce(t *testing.T) {
	repo := &fakeRepo{quote: Quote{Text: "be water"}}
	svc := NewService(repo)

	if _, err := svc.Random(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if repo.calls != 1 {
		t.Errorf("repo called %d times, want 1", repo.calls)
	}
}

func TestService_Random_ForwardsContext(t *testing.T) {
	type ctxKey struct{}
	repo := &fakeRepo{}
	svc := NewService(repo)

	ctx := context.WithValue(context.Background(), ctxKey{}, "v")
	if _, err := svc.Random(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if repo.gotCtx != ctx {
		t.Errorf("repo received ctx %v, want %v", repo.gotCtx, ctx)
	}
	if got := repo.gotCtx.Value(ctxKey{}); got != "v" {
		t.Errorf("ctx value = %v, want %q", got, "v")
	}
}
