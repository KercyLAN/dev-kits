// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 12:22

package krand

import "testing"

func TestInt64(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Int64(0, 10))
	}
}

func TestName(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Name())
	}
}

func TestEnglishName(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(EnglishName())
	}
}

func TestHostName(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(HostName())
	}
}

func TestChineseName(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(ChineseName())
	}
}

func TestInt(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Int(0, 10))
	}
}

func TestNumberString(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(NumberString(0, 99999))
	}
}

func TestNumberStringRepair(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(NumberStringRepair(0, 99999))
	}
}