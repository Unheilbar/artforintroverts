package scheduler

import (
	"context"
	"fmt"
	"time"
)

func RunRefreshJob(ctx context.Context, refresh func() error, interval uint) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(interval))
	for {
		err := refresh()
		if err != nil {
			fmt.Printf("err during refresh state %s \n", err)
		}
		select {
		case <-ticker.C:
			continue
		case <-ctx.Done():
			return
		}
	}
}
