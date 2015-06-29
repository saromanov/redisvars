package redisvars

import
(
	"testing"
	"encoding/json"
)

type Foobar struct {
	Item string
	Data int
}

func TestRedisStruct(t *testing.T) {
	elem := &Foobar{"first",1}
	rv := New("localhost:6379")
	rstr := rv.NewStruct()
	rstr.SetStruct("foobar1",elem)
	rstr.CommitStruct()
	res := rstr.GetStruct("foobar1")
	var foobar Foobar
	err := json.Unmarshal([]byte(res), &foobar)
	if err != nil {
		t.Errorf("Found error")
	}

	if foobar.Item != "first" {
		t.Errorf("%s not match %s", "first", foobar.Item)
	}
	if foobar.Data != 1 {
		t.Errorf("%d not match %d", 1, foobar.Data)
	}
}