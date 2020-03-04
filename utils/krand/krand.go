// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-28 19:37

package krand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(999)))
}

