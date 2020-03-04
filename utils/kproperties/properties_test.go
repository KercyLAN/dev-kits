// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 11:03

package kproperties

import (
	"fmt"
	"github.com/KercyLAN/dev-kits/utils/kruntime"
	"testing"
)

var properties *Properties
func init() {
	var err error
	properties, err = New(kruntime.PathWork() + "\\test.properties")
	if err != nil {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	properties, err := New(kruntime.PathWork() + "\\test.properties")
	if err != nil {
		panic(err)
	}else {
		t.Log(properties)
	}
}

func TestProperties_Each(t *testing.T) {
	properties.Each(func(key string, value string) {
		t.Log(fmt.Sprintf("%v = %v", key, value))
	})
}

func TestProperties_GetBool(t *testing.T) {
	t.Log(properties.GetBool("kproperties.bool.true"))
	t.Log(properties.GetBool("kproperties.bool.false"))
}

func TestProperties_GetFloat32(t *testing.T) {
	t.Log(properties.GetFloat32("kproperties.float"))
	t.Log(properties.GetFloat32("kproperties.float.space"))
}

func TestProperties_GetFloat64(t *testing.T) {
	t.Log(properties.GetFloat64("kproperties.float"))
	t.Log(properties.GetFloat64("kproperties.float.space"))
}

func TestProperties_GetInt(t *testing.T) {
	t.Log(properties.GetInt("kproperties.int"))
	t.Log(properties.GetInt("kproperties.int.space"))
}

func TestProperties_GetInt8(t *testing.T) {
	t.Log(properties.GetInt8("kproperties.int"))
	t.Log(properties.GetInt8("kproperties.int.space"))
}

func TestProperties_GetInt16(t *testing.T) {
	t.Log(properties.GetInt16("kproperties.int"))
	t.Log(properties.GetInt16("kproperties.int.space"))
}

func TestProperties_GetInt32(t *testing.T) {
	t.Log(properties.GetInt32("kproperties.int"))
	t.Log(properties.GetInt32("kproperties.int.space"))
}

func TestProperties_GetInt64(t *testing.T) {
	t.Log(properties.GetInt64("kproperties.int"))
	t.Log(properties.GetInt64("kproperties.int.space"))
}

func TestProperties_GetUint(t *testing.T) {
	t.Log(properties.GetUint("kproperties.int"))
	t.Log(properties.GetUint("kproperties.int.space"))
}

func TestProperties_GetUint8(t *testing.T) {
	t.Log(properties.GetUint8("kproperties.int"))
	t.Log(properties.GetUint8("kproperties.int.space"))
}

func TestProperties_GetUint16(t *testing.T) {
	t.Log(properties.GetUint16("kproperties.int"))
	t.Log(properties.GetUint16("kproperties.int.space"))
}

func TestProperties_GetUint32(t *testing.T) {
	t.Log(properties.GetUint32("kproperties.int"))
	t.Log(properties.GetUint32("kproperties.int.space"))
}

func TestProperties_GetUint64(t *testing.T) {
	t.Log(properties.GetUint64("kproperties.int"))
	t.Log(properties.GetUint64("kproperties.int.space"))
}

func TestProperties_GetInterface(t *testing.T) {
	t.Log(properties.GetInterface("kproperties.string"))
}

func TestProperties_GetString(t *testing.T) {
	t.Log(properties.GetString("kproperties.string"))
}

func TestProperties_HasKey(t *testing.T) {
	t.Log(properties.HasKey("kproperties.string"))
	t.Log(properties.HasKey("kproperties.strings"))
}