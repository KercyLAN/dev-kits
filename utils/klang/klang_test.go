// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 11:27

package klang

import (
	"fmt"
	"github.com/KercyLAN/dev-kits/utils/kruntime"
	"testing"
)

var lang = New("test", kruntime.PathWork(), LOCAL_NONE)

func TestNew(t *testing.T) {
	t.Log(New("test", kruntime.PathWork(), LOCAL_NONE))
}

func TestLang_FastSwitchSize(t *testing.T) {
	t.Log(fmt.Sprintf("before: %v", lang.FastSwitchSize()))
	change(LOCAL_zh)
	t.Log(fmt.Sprintf("later: %v", lang.FastSwitchSize()))
}

func TestLang_FastSwitchSizeLocals(t *testing.T) {
	t.Log(fmt.Sprintf("before: %v", lang.FastSwitchSizeLocals()))
	change(LOCAL_zh)
	t.Log(fmt.Sprintf("later: %v", lang.FastSwitchSizeLocals()))
}

func TestLang_Get(t *testing.T) {
	t.Log(fmt.Sprintf("before: %v", lang.Get("klang.hello")))
	t.Log(fmt.Sprintf("before: %v", lang.Get("klang.test")))
	change(LOCAL_zh)
	t.Log(fmt.Sprintf("later: %v", lang.Get("klang.hello")))
	t.Log(fmt.Sprintf("later: %v", lang.Get("klang.test")))
}

func TestLang_GetLocal(t *testing.T) {
	t.Log(lang.GetLocal())
}

func TestLang_Reset(t *testing.T) {
	t.Log(fmt.Sprintf("before: %v", lang))
	change(LOCAL_zh)
	t.Log(fmt.Sprintf("later: %v", lang))
	lang.Reset()
	t.Log(fmt.Sprintf("reset: %v", lang))
}

func TestLang_SetEncoder(t *testing.T) {
	lang.SetEncoder(EncoderGbkUtf8)
}

func TestLang_SetLocal(t *testing.T) {
	err := lang.SetLocal(LOCAL_zh)
	if err != nil {
		panic(err)
	}
}

func TestLang_SetSweepers(t *testing.T) {
	lang.SetSweepers(func(lang *Lang) {
		t.Log("Reach the threshold!")
	})
	change(LOCAL_en)
	change(LOCAL_th)
	change(LOCAL_tr)
	change(LOCAL_uk)
	change(LOCAL_ur)
	change(LOCAL_zh)
}

func TestLang_SetSweepersThresholdValue(t *testing.T) {
	lang.SetSweepersThresholdValue(2)
	lang.SetSweepers(func(lang *Lang) {
		t.Log("Reach the threshold!")
	})
	change(LOCAL_en)
	change(LOCAL_th)
	change(LOCAL_tr)
	change(LOCAL_uk)
	change(LOCAL_ur)
	change(LOCAL_zh)
}

func change(local Local) {
	err := lang.SetLocal(local)
	if err != nil {
		panic(err)
	}
}