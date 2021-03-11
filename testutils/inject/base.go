package inject

import (
	"context"

	"go.viam.com/robotcore/api"
)

type Base struct {
	api.Base
	MoveStraightFunc func(ctx context.Context, distanceMillis int, millisPerSec float64, block bool) error
	SpinFunc         func(ctx context.Context, angleDeg float64, speed int, block bool) error
	WidthMillisFunc  func(ctx context.Context) (int, error)
	StopFunc         func(ctx context.Context) error
	CloseFunc        func(ctx context.Context) error
}

func (b *Base) MoveStraight(ctx context.Context, distanceMillis int, millisPerSec float64, block bool) error {
	if b.MoveStraightFunc == nil {
		return b.Base.MoveStraight(ctx, distanceMillis, millisPerSec, block)
	}
	return b.MoveStraightFunc(ctx, distanceMillis, millisPerSec, block)
}

func (b *Base) Spin(ctx context.Context, angleDeg float64, speed int, block bool) error {
	if b.SpinFunc == nil {
		return b.Base.Spin(ctx, angleDeg, speed, block)
	}
	return b.SpinFunc(ctx, angleDeg, speed, block)
}

func (b *Base) WidthMillis(ctx context.Context) (int, error) {
	if b.WidthMillisFunc == nil {
		return b.Base.WidthMillis(ctx)
	}
	return b.WidthMillisFunc(ctx)
}

func (b *Base) Stop(ctx context.Context) error {
	if b.StopFunc == nil {
		return b.Base.Stop(ctx)
	}
	return b.StopFunc(ctx)
}

func (b *Base) Close(ctx context.Context) error {
	if b.CloseFunc == nil {
		return b.Base.Close(ctx)
	}
	return b.CloseFunc(ctx)
}
