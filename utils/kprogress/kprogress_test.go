// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 10:47

package kprogress

import (
	"fmt"
	"testing"
	"time"
)

func TestProgress(t *testing.T) {
	for i := 1; i <= 100; i++ {
		time.Sleep(10)
		Progress(int64(i), 100, func(now int64, gross int64, progress float64, finish bool) {
			fmt.Println(fmt.Sprintf("progress: %v/%v %v％ %v", now, gross, progress, finish))
		})
	}
}
