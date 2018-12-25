package main

import "testing"

func TestMakeStatusDict(t *testing.T) {
	statusDict := MakeStatusDict()

	if l := len(statusDict.S); l != 0 {
		t.Error("empty statusDict")
	}
}

func TestAddToStatusDict(t *testing.T) {
	var id int

	statusDict := MakeStatusDict()
	id = statusDict.Add("one")
	if id != 0 {
		t.Error("insert")
	}
	id = statusDict.Add("two")
	if id != 1 {
		t.Error("insert")
	}
	if l := len(statusDict.S); l != 2 {
		t.Error("empty statusDict")
	}
}

func TestAddSameStatusDict(t *testing.T) {
	var id int

	statusDict := MakeStatusDict()
	id = statusDict.Add("one")
	if id != 0 {
		t.Error("insert")
	}
	id = statusDict.Add("one")
	if id != 0 {
		t.Error("update")
	}

}
