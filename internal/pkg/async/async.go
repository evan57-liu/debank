package async

import (
	"log"
	"runtime/debug"
)

// Go creates a goroutine.
// It is very dangerous to use the go keyword to directly create a goroutine.
// Once a panic occurs inside the goroutine without calling the recover() function,
// the entire application will exit.
// We need to pay attention to the following two points:
// 1. Panic will only trigger to defer of the current Goroutine.
// 2. Recover will only take effect when called in defer.
func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[async] panic recovered: %v\n%s", err, string(debug.Stack()))
			}
		}()

		f()
	}()
}
