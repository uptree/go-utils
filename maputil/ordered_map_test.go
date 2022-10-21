package maputil

import (
	"encoding/json"
	"sort"
	"strings"
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestOrderedMap(t *testing.T) {
	o := NewOrderedMap()
	// number
	o.Set("number", 3)
	v, _ := o.Get("number")
	assert.Equalf(t, 3, v.(int), "FAIL")
	// string
	o.Set("string", "x")
	v, _ = o.Get("string")
	assert.Equalf(t, "x", v.(string), "FAIL")
	// string slice
	o.Set("strings", []string{
		"t",
		"u",
	})
	v, _ = o.Get("strings")
	assert.Equalf(t, "t", v.([]string)[0], "FAIL")
	assert.Equalf(t, "u", v.([]string)[1], "FAIL")
	// mixed slice
	o.Set("mixed", []interface{}{
		1,
		"1",
	})
	v, _ = o.Get("mixed")
	assert.Equalf(t, 1, v.([]interface{})[0].(int), "FAIL")
	assert.Equalf(t, "1", v.([]interface{})[1].(string), "FAIL")

	// overriding existing key
	o.Set("number", 4)
	v, _ = o.Get("number")
	assert.Equalf(t, 4, v.(int), "FAIL")

	// Keys method
	keys := o.Keys()
	expectedKeys := []string{
		"number",
		"string",
		"strings",
		"mixed",
	}
	for i, key := range keys {
		assert.Equalf(t, expectedKeys[i], key, "FAIL")
	}
	for i, key := range expectedKeys {
		assert.Equalf(t, expectedKeys[i], key, "FAIL")
	}
	// delete
	o.Delete("strings")
	o.Delete("not a key being used")
	assert.Equalf(t, 3, len(o.Keys()), "FAIL")
	_, ok := o.Get("strings")
	assert.Equalf(t, false, ok, "FAIL")
}

func TestBlankMarshalJSON(t *testing.T) {
	o := NewOrderedMap()
	// blank map
	b, err := json.Marshal(o)
	assert.Nil(t, err)
	assert.Equalf(t, `{}`, string(b), "FAIL")

	// convert to indented json
	bi, err := json.MarshalIndent(o, "", "  ")
	assert.Nil(t, err)
	assert.Equalf(t, `{}`, string(bi), "FAIL")
}

func TestMarshalJSON(t *testing.T) {
	o := NewOrderedMap()
	// number
	o.Set("number", 3)
	// string
	o.Set("string", "x")
	// string
	o.Set("specialstring", "\\.<>[]{}_-")
	// new value keeps key in old position
	o.Set("number", 4)
	// keys not sorted alphabetically
	o.Set("z", 1)
	o.Set("a", 2)
	o.Set("b", 3)
	// slice
	o.Set("slice", []interface{}{
		"1",
		1,
	})
	// orderedmap
	v := NewOrderedMap()
	v.Set("e", 1)
	v.Set("a", 2)
	o.Set("orderedmap", v)
	// escape key
	o.Set("test\n\r\t\\\"ing", 9)
	// convert to json
	b, err := json.Marshal(o)
	assert.Nil(t, err)
	// check json is correctly ordered
	expected := `{"number":4,"string":"x","specialstring":"\\.\u003c\u003e[]{}_-","z":1,"a":2,"b":3,"slice":["1",1],"orderedmap":{"e":1,"a":2},"test\n\r\t\\\"ing":9}`
	assert.Equalf(t, expected, string(b), "FAIL")

	// convert to indented json
	bi, err := json.MarshalIndent(o, "", "  ")
	assert.Nil(t, err)
	ei := `{
  "number": 4,
  "string": "x",
  "specialstring": "\\.\u003c\u003e[]{}_-",
  "z": 1,
  "a": 2,
  "b": 3,
  "slice": [
    "1",
    1
  ],
  "orderedmap": {
    "e": 1,
    "a": 2
  },
  "test\n\r\t\\\"ing": 9
}`
	assert.Equalf(t, ei, string(bi), "FAIL")
}

func TestMarshalJSONNoEscapeHTML(t *testing.T) {
	o := NewOrderedMap()
	o.SetEscapeHTML(false)
	// string special characters
	o.Set("specialstring", "\\.<>[]{}_-")
	// convert to json
	b, err := o.MarshalJSON()
	assert.Nil(t, err)

	s := strings.Replace(string(b), "\n", "", -1)
	// check json is correctly ordered
	assert.Equalf(t, `{"specialstring":"\\.<>[]{}_-"}`, s, "FAIL")
}

func TestMarshalJSONNoEscapeHTMLRecursive(t *testing.T) {
	src := `{"x":"<>","y":[{"z":["<>"]}]}`
	o := NewOrderedMap()
	o.SetEscapeHTML(false)
	err := json.Unmarshal([]byte(src), &o)
	assert.Nil(t, err)

	b, err := o.MarshalJSON()
	assert.Nil(t, err)

	s := strings.Replace(string(b), "\n", "", -1)
	assert.Equalf(t, src, s, "FAIL")
}

func TestUnmarshalJSONStruct(t *testing.T) {
	var v struct {
		Data *OrderedMap `json:"data"`
	}

	err := json.Unmarshal([]byte(`{ "data": { "x": 1 } }`), &v)
	assert.Nil(t, err)

	x, ok := v.Data.Get("x")
	assert.Equalf(t, true, ok, "FAIL")
	assert.Equalf(t, float64(1), x, "FAIL")
}

func TestOrderedMap_SortKeys(t *testing.T) {
	s := `
{
  "b": 2,
  "a": 1,
  "c": 3
}
`
	o := NewOrderedMap()
	json.Unmarshal([]byte(s), &o)

	o.SortKeys(sort.Strings)

	// Check the root keys
	expectedKeys := []string{
		"a",
		"b",
		"c",
	}
	k := o.Keys()
	for i := range k {
		if k[i] != expectedKeys[i] {
			t.Error("SortKeys root key order", i, k[i], "!=", expectedKeys[i])
		}
	}
}

func TestOrderedMap_Sort(t *testing.T) {
	s := `
{
  "b": 2,
  "a": 1,
  "c": 3
}
`
	o := NewOrderedMap()
	json.Unmarshal([]byte(s), &o)
	o.Sort(func(a *Pair, b *Pair) bool {
		return a.value.(float64) > b.value.(float64)
	})

	// Check the root keys
	expectedKeys := []string{
		"c",
		"b",
		"a",
	}
	k := o.Keys()
	for i := range k {
		assert.Equalf(t, expectedKeys[i], k[i], "FAIL")
	}
}
