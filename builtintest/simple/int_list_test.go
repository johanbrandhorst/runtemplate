package simple

import (
	"testing"

	"sort"

)

func intRangeOf(from, to int) []int {
	n := 1 + to - from
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = i + from
	}
	return a
}

func TestNewIntList(t *testing.T) {
	a := NewXIntList(1, 2, 3)

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}

	if a.Get(1) != 2 {
		t.Errorf("Expected 2 but got %d", a.Get(1))
	}

	if a.IsSet() {
		t.Error("Expected not a set")
	}

	if !a.IsSequence() {
		t.Error("Expected a sequence")
	}
}


func TestIntListClone(t *testing.T) {
	a := NewXIntList(1, 2)

	b := a.Clone()

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}

	c := a.Clone().Tail()

	if a.Equals(c) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a, c)
	}
}

func TestIntListSend(t *testing.T) {
	a := NewXIntList(1, 2, 3, 4)

	b := BuildXIntListFromChan(a.Send())

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}
}

func TestIntListHeadTailLastInit(t *testing.T) {
	a := NewXIntList(1, 2, 3, 4)

	if a.Head() != 1 {
		t.Errorf("Expected 1 but got %d", a.Head())
	}

	if a.Last() != 4 {
		t.Errorf("Expected 4 but got %d", a.Last())
	}

	tail := a.Tail()
	if !tail.Equals(NewXIntList(2, 3, 4)) {
		t.Errorf("Expected '2, 3, 4' but got '%+v'", tail)
	}

	init := a.Init()
	if !init.Equals(NewXIntList(1, 2, 3)) {
		t.Errorf("Expected '1, 2, 3' but got '%+v'", init)
	}
}

func TestIntListFilter(t *testing.T) {
	a := NewXIntList(1, 2, 3, 4)

	b := a.Filter(func(v int) bool {
		return v > 2
	})

	if !b.Equals(NewXIntList(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b)
	}
}

func TestIntListPartition(t *testing.T) {
	a := NewXIntList(1, 2, 3, 4)

	b, c := a.Partition(func(v int) bool {
		return v > 2
	})

	if !b.Equals(NewXIntList(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b)
	}

	if !c.Equals(NewXIntList(1, 2)) {
		t.Errorf("Expected '1, 2' but got '%+v'", c)
	}
}


func TestIntListSort(t *testing.T) {
	a := NewXIntList(13, 4, 7, -2, 9)

	sort.Sort(a)

	if !a.Equals(NewXIntList(-2, 4, 7, 9, 13)) {
		t.Errorf("Expected '3, 4' but got '%+v'", a)
	}
}

func TestIntListReverse(t *testing.T) {
	a := NewXIntList(13, 4, 7, -2, 9)

	b := a.Reverse()

	if b.Equals(a) {
		t.Errorf("Expected reverse of '%+v' but got '%+v'", a, b)
	}

	c := b.Reverse()

	if !c.Equals(a) {
		t.Errorf("Expected '%+v' but got '%+v'", a, c)
	}
}

func TestIntListShuffle(t *testing.T) {
	a := NewXIntList(intRangeOf(1, 1000)...)

	b := a.Shuffle()

	if b.Equals(a) {
		t.Errorf("Expected shuffle of '%+v' but got '%+v'", a, b)
	}

	// prove that the same set of numbers is present
	sort.Sort(b)

	if !b.Equals(a) {
		t.Errorf("Expected '%+v' but got '%+v'", a, b)
	}
}


func TestIntListTake(t *testing.T) {
	a := NewXIntList(intRangeOf(1, 100)...)

	b := a.Take(30)

	if b.Len() != 30 || b.Head() != 1 || b.Last() != 30 {
		t.Errorf("Expected list from 1 to 30 but got '%+v'", b)
	}

	c := a.TakeLast(30)

	if c.Len() != 30 || c.Head() != 71 || c.Last() != 100 {
		t.Errorf("Expected list from 71 to 100 but got '%+v'", c)
	}

	d := a.Take(101)

	if d.Len() != 100 || d.Head() != 1 || d.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", d)
	}

	e := a.TakeLast(101)

	if e.Len() != 100 || e.Head() != 1 || e.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", e)
	}
}

func TestIntListDrop(t *testing.T) {
	a := NewXIntList(intRangeOf(1, 100)...)

	b := a.Drop(70)

	if b.Len() != 30 || b.Head() != 71 || b.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", b)
	}

	c := a.DropLast(75)

	if c.Len() != 25 || c.Head() != 1 || c.Last() != 25 {
		t.Errorf("Expected list from 1 to 25 but got '%+v'", c)
	}

	d := a.Drop(101)

	if d.Len() != 0 {
		t.Errorf("Expected empty list but got '%+v'", d)
	}

	e := a.DropLast(101)

	if e.Len() != 0 {
		t.Errorf("Expected empty list but got '%+v'", e)
	}
}

func TestIntListTakeWhile(t *testing.T) {
	a := NewXIntList(intRangeOf(1, 100)...)

	b := a.TakeWhile(func(v int) bool {
		return v <= 20
	})

	if b.Len() != 20 || b.Head() != 1 || b.Last() != 20 {
		t.Errorf("Expected list from 1 to 20 but got '%+v'", b)
	}

	c := a.TakeWhile(func(v int) bool {
		return true
	})

	if c.Len() != 100 || c.Head() != 1 || c.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", b)
	}
}

func TestIntListDropWhile(t *testing.T) {
	a := NewXIntList(intRangeOf(1, 100)...)

	b := a.DropWhile(func(v int) bool {
		return v <= 80
	})

	if b.Len() != 20 || b.Head() != 81 || b.Last() != 100 {
		t.Errorf("Expected list from 81 to 100 but got '%+v'", b)
	}

	c := a.DropWhile(func(v int) bool {
		return true
	})

	if c.Len() != 0 {
		t.Errorf("Expected empty list but got '%+v'", b)
	}
}

func TestIntListDistinctBy(t *testing.T) {
	a := NewXIntList(1, 1, 1, 2, 1, 2, 3, 4, 5, 3, 3, 5)

	c := a.DistinctBy(func(v1, v2 int) bool {
		return v1 == v2
	})

	if !c.Equals(NewXIntList(1, 2, 3, 4, 5)) {
		t.Errorf("Expected 1 to 5 but got '%+v'", c)
	}
}

func TestIntListIndexWhere(t *testing.T) {
	a := NewXIntList(intRangeOf(1, 100)...)

	b := a.IndexWhere(func(v int) bool {
		return v >= 47
	})

	if b != 46 {
		t.Errorf("Expected 46 but got %d", b)
	}

	c := a.IndexWhere(func(v int) bool {
		return false
	})

	if c != -1 {
		t.Errorf("Expected -1 but got %d", c)
	}

	d := a.IndexWhere2(func(v int) bool {
		return v % 3 == 0
	}, 10)

	if d != 11 {
		t.Errorf("Expected 11 but got %d", d)
	}
}

func TestIntListLastIndexWhere(t *testing.T) {
	a := NewXIntList(intRangeOf(1, 100)...)

	b := a.LastIndexWhere(func(v int) bool {
		return v <= 47
	})

	if b != 46 {
		t.Errorf("Expected 46 but got %d", b)
	}

	c := a.LastIndexWhere(func(v int) bool {
		return false
	})

	if c != -1 {
		t.Errorf("Expected -1 but got %d", c)
	}

	d := a.LastIndexWhere2(func(v int) bool {
		return v % 3 == 0
	}, 61)

	if d != 59 {
		t.Errorf("Expected 59 but got %d", d)
	}
}
