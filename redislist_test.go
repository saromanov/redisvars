package redisvars

import (
	"testing"
)

func TestSetInList(t *testing.T) {
	rv := New("localhost:6379")
	lst := rv.NewList()
	lst.SetList("A", []string{"A", "B", "C"})
	lst.CommitList()
	result := lst.GetList("A")
	if result[0] != "C" || result[1] != "B" || result[2] != "A" {
		t.Errorf("Not match")
	}
}
