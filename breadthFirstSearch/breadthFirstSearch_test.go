package breadthFirstSearch

import "testing"

var people = map[string][]string{
	"you": {"alice", "bob", "claire", },
	"bob": {"abuja", "peggy"},
	"alice": {"peggy"},
	"claire": {"thom", "johnny"},
	"abuja": {"fed", "def"},
	"peggy": {"red", "der"},
	"thom": {"mike"},
	"johny": {"putin", "muller"},
	"fed": {"muller", "ivan"},
	"def": {},
	"red": {"putin", "ivan", "max"},
	"der": {},
	"mike": {},
	"putin": {},
	"muller": {},
}

func TestBreadthFirstSearch(t *testing.T) {
	item := BreadthFirstSearch(people)
	if !item {
		t.Errorf("TestBreadthFirstSearch for OK Failed - error")
	}
}