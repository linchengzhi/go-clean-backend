package config

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	for {
		select {
		case <-ticker.C:
			num := rand.Int31n(5)
			fmt.Println(num)
			time.Sleep(time.Duration(num) * time.Second)
			break
		case <-time.After(100 * time.Second):
			t.Error("timeout")
		}
	}
}
