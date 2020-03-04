// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 12:25

package kruntime

import "testing"

func TestNumCPU(t *testing.T) {
	t.Log(NumCPU())
}

func TestPathRunDir(t *testing.T) {
	t.Log(PathRunDir())
}

func TestPathRunExe(t *testing.T) {
	t.Log(PathRunExe())
}

func TestPathWork(t *testing.T) {
	t.Log(PathWork())
}