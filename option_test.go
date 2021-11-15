package option

import (
	"strconv"
	"testing"
)

func TestSome(t *testing.T) {
	s := Some(1)
	if s.IsSome() != true {
		t.Fatal("Some(1).IsSome() should be true")
	}
}

func TestNone(t *testing.T) {
	s := None[int]()
	if s.IsSome() != false {
		t.Fatal("None().IsSome() should be false")
	}
}

func TestGet(t *testing.T) {
	s := Some(1)
	if Get(s) != 1 {
		t.Fatal("Some(1).Get() should be 1")
	}
}

func TestNoneGet(t *testing.T) {
	defer func() { recover() }()
	func() {
		None[int]().Get()
	}()
	t.Errorf("did not panic")
}

func TestOrElse(t *testing.T) {
	or := func() int {
		return 2
	}
	if OrElse(or, Some(1)) != 1 {
		t.Fatal("Should get 1 on Some")
	}
	if OrElse(or, None[int]()) != 2 {
		t.Fatal("Should get 2 on None")
	}
}

func TestOr(t *testing.T) {
	if Or(2, Some(1)) != 1 {
		t.Fatal("Should get 1 on Some")
	}
	if Or(2, None[int]()) != 2 {
		t.Fatal("Should get 2 on None")
	}
}

func TestMap(t *testing.T) {
	s := Some(1)
	if Map(func(v int) int { return v + 1 }, s).Get() != 2 {
		t.Fatal("Should map + 1 eq 2")
	}
	if Map(func(v int) string { return strconv.Itoa(v) }, s).Get() != "1" {
		t.Fatal("Should map to string \"1\"")
	}
}

func TestMapNone(t *testing.T) {
	n := None[int]()
	n = Map(func(v int) int { return v + 1 }, n)
	m := Map(func(v int) string { return strconv.Itoa(v) }, n)
	if n.IsSome() || m.IsSome() {
		t.Fatal("Should be None")
	}
}
