// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 12:26

package kstr

import "testing"

func TestUpperFirst(t *testing.T) {
	t.Log(UpperFirst("abcd"))
}

func TestLowerFirst(t *testing.T) {
	t.Log(LowerFirst("ABCD"))
}

func TestThousandsSeparator(t *testing.T) {
	t.Log(ThousandsSeparator("TestString"))
	t.Log(ThousandsSeparator("1002942821.42"))
}

func TestIsUpperPrefix(t *testing.T) {
	t.Log(IsUpperPrefix("Kercylan"))
	t.Log(IsUpperPrefix("kercylan"))
}

func TestIsLowerPrefix(t *testing.T) {
	t.Log(IsLowerPrefix("Kercylan"))
	t.Log(IsLowerPrefix("kercylan"))
}

func TestRemoveLast(t *testing.T) {
	t.Log(RemoveLast("ABCD"))
}

func TestRemoveFirst(t *testing.T) {
	t.Log(RemoveFirst("ABCD"))
}

func TestDistinct(t *testing.T) {
	t.Log(Distinct("ABCCCDDEFG", "C"))
}

func TestFormatSpeedyFloat32(t *testing.T) {
	t.Log(FormatSpeedyFloat32("123.45"))
}

func TestFormatSpeedyFloat64(t *testing.T) {
	t.Log(FormatSpeedyFloat64("123.45"))
}

func TestFormatSpeedyInt(t *testing.T) {
	t.Log(FormatSpeedyInt("12345"))
}

func TestFormatSpeedyInt64(t *testing.T) {
	t.Log(FormatSpeedyInt64("12345"))
}

func TestKV(t *testing.T) {
	t.Log(KV("a=b"))
	t.Log(KV("a:b", ":"))
}

func TestHideSensitivity(t *testing.T) {
	t.Log(HideSensitivity("17717777777"))
	t.Log(HideSensitivity("xusixxxx@gmail.com"))
}

