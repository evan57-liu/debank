package async

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// Group 封装 errgroup.Group 并增加 recover 机制
type Group struct {
	*errgroup.Group
}

// WithContext 创建带有 context 的 Group
func WithContext(ctx context.Context) (*Group, context.Context) {
	g, ctx := errgroup.WithContext(ctx)

	return &Group{g}, ctx
}

// Go 封装 errgroup.Group.Go 方法，增加 recover 处理，防止 panic 导致程序崩溃
func (g *Group) Go(f func() error) {
	g.Group.Go(func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				// 捕获 panic，并转换为 error 返回，防止 goroutine 崩溃
				err = fmt.Errorf("panic recovered: %v", r)
			}
		}()

		return f()
	})
}
