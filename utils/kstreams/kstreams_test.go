// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 12:32

package kstreams

import (
	"errors"
	"testing"
)

const template = `line 1
line 2
line 3`

func TestEachLine(t *testing.T) {
	EachLine(template, func(index int, line string) {
		t.Log(index, line)
	})
}

func TestEachLineOff(t *testing.T) {
	err := EachLineOff(template, func(index int, line string) error {
		if index == 1 {
			return errors.New("stop!")
		}
		t.Log(index, line)
		return nil
	})
	if err != nil {
		t.Log(err)
	}
}

func TestEachMapSort(t *testing.T) {
	mInt := map[int]string {
		3: "line 3",
		1: "line 1",
		2: "line 2",
	}
	EachMapSort(mInt, func(key int, value string) {
		t.Log(key, value)
	})
	t.Log("------------------------------------")
	mString := map[string]string {
		"3": "line 3",
		"1": "line 1",
		"b": "line b",
		"2": "line 2",
		"a": "line a",
	}
	EachMapSort(mString, func(key string, value string) {
		t.Log(key, value)
	})
}
