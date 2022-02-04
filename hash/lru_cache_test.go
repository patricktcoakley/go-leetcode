package hash

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	got := cache.Get(1)
	want := 1

	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
	}

	cache.Put(3, 3)

	got = cache.Get(2)
	want = -1

	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
	}

	cache.Put(4, 4)

	got = cache.Get(1)
	want = -1

	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
	}

	got = cache.Get(3)
	want = 3

	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
	}

	got = cache.Get(4)
	want = 4

	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
	}

}
