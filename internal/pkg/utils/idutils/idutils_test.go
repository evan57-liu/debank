package idutils

import (
	"sync"
	"sync/atomic"
	"testing"
)

//	go test -bench=. -benchtime=3s -run=^$
//
// 基准测试：单线程调用 GenID
func BenchmarkGenID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenID()
	}
}

// 基准测试：多线程并发调用 GenID
func BenchmarkGenIDParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = GenID()
		}
	})
}

// 压力测试：高并发生成 ID
func TestGenIDConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	numGoroutines := 10000 // 并发 Goroutine 数量
	idSet := sync.Map{}    // 使用 sync.Map 存储生成的 ID，避免数据竞争
	var duplicateCount int32

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				id := GenID()

				// 检查 ID 是否重复
				if _, loaded := idSet.LoadOrStore(id, struct{}{}); loaded {
					atomic.AddInt32(&duplicateCount, 1) // 记录重复的 ID 数量
				}
			}
		}()
	}
	wg.Wait()

	if duplicateCount > 0 {
		t.Errorf("生成的 ID 存在重复，重复次数: %d", duplicateCount)
	} else {
		t.Logf("所有 %d 个 ID 都是唯一的!", numGoroutines*100)
	}
}
