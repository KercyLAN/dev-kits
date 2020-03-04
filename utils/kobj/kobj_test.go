// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 12:21

package kobj

import (
	"testing"
)

type testStruct struct {
	i int
}

func TestIsEmpty(t *testing.T) {
	var s *testStruct
	var m map[string]int
	var slice []string
	t.Log(IsEmpty(1))
	t.Log(IsEmpty(1.1))
	t.Log(IsEmpty(0.0))
	t.Log(IsEmpty("sd"))
	t.Log(IsEmpty(""))
	t.Log(IsEmpty(make([]string, 0)))
	t.Log(IsEmpty(nil))
	t.Log(IsEmpty(s))
	t.Log(IsEmpty(&testStruct{}))
	t.Log(IsEmpty(m))
	t.Log(IsEmpty(map[string]string{}))
	t.Log(IsEmpty(slice))
}

func TestIsAllEmpty(t *testing.T) {
	var s *testStruct
	var m map[string]int
	var slice []string
	t.Log(IsAllEmpty(0,0.0,"","",nil,s,m,slice))
	t.Log(IsAllEmpty(1,1.1,0.0,"sd","",make([]string, 0),nil,s,&testStruct{},m,map[string]string{},slice))
}