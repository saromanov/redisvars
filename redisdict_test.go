package redisvars

import (
	"testing"
)

func TestBasicGetting(t *testing.T) {
	rv := New("localhost:6379")
	dict := rv.NewDict()
	key := "A"
	value := "B"
	dict.Delete(key)
	dict.Set(key, value)
	dict.Commit()
	result := dict.Get(key)
	if result != value {
		t.Errorf("%s not match %s", result, value)
	}
}

func TestGetAfterOverwrite(t *testing.T) {
	rv := New("localhost:6379")
	dict := rv.NewDict()
	key := "A"
	value := "B"
	dict.Delete(key)
	dict.Set(key, value)
	dict.Commit()
	dict.Set(key, "C")
	result := dict.Get(key)
	if result != value {
		t.Errorf("%s not match %s", result, value)
	}
}

func TestGetFromEmpty(t *testing.T) {
	rv := New("localhost:6379")
	dict := rv.NewDict()
	key := "A"
	value := "B"
	dict.Delete(key)
	dict.Delete("C")
	dict.Set(key, value)
	dict.Commit()
	result := dict.Get("C")
	if result != "<nil>" {
		t.Errorf("%s not match %s", result, value)
	}
}

func TestSetAndGetManyElements(t *testing.T) {
	rv := New("localhost:6379")
	dict := rv.NewDict()
	keys := map[string]string{"A": "B", "C": "D", "E": "F"}
	for key, value := range keys {
		dict.Delete(key)
		dict.Set(key, value)
	}
	dict.Commit()
	for key, value := range keys {
		result := dict.Get(key)
		if result != value {
			t.Errorf("%s not match %s", result, value)
		}
	}
}
