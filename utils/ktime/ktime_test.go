// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 12:38

package ktime

import (
	"testing"
	"time"
)

func TestInterval(t *testing.T) {
	d := time.Now()
	t.Log(Interval(d.Unix()))
	t.Log(Interval(d.Unix() - 10))
	t.Log(Interval(d.Unix() - 100))
	t.Log(Interval(d.Unix() - 10000))
	t.Log(Interval(d.Unix() - 100000))
	t.Log(Interval(d.Unix() - 100000000))
}
