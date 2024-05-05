package reactive

import (
	"context"
	"time"

	"github.com/base-go/baseql/baseqlpb"
)

func InvalidateAfter(ctx context.Context, d time.Duration) {
	r := NewResource()
	timer := time.AfterFunc(d, r.Invalidate)
	r.Cleanup(func() { timer.Stop() })
	AddDependency(ctx, r, &baseqlpb.ExpirationTime{Time: time.Now().Add(d)})
}

func InvalidateAt(ctx context.Context, t time.Time) {
	InvalidateAfter(ctx, t.Sub(time.Now()))
}
