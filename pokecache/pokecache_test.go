package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCaching(t *testing.T) {
	const interval = time.Millisecond * 100

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
		{
			key: "https://gogiel.com",
			val: []byte("evenmoretestdata"),
		},
		{
			key: "https://twitor.com/something",
			val: []byte("among us <- funny btw if it wasnt clear"),
		},
		{
			key: "https://www.youtube.com/watch?v=H9aC5AGY9YU",
			val: []byte("juan."),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}

			time.Sleep(interval)
			_, ok = cache.Get(c.key)
			if ok {
				t.Errorf("expected to NOT find key")
				return
			}
		})
	}
}
