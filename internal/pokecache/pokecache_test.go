package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct{
		key string
		val []byte
	} {
		{
			key: "example.com",
			val: []byte("test"),	
		},
		{
			key: "anothertest.com",
			val: []byte("this better be there!"),	
		},
		{
			key: "",
			val: []byte(""),	
		},
		{
			key: " tes  ttt ",
			val: []byte("   te  st   "),	
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case: %d", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key: %s", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Values not matching: %v != %v", val, c.val)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	cases := []struct {
		intervalTime time.Duration
		waitTime time.Duration
		key string
		found bool
	} {
		{
			intervalTime: 5 * time.Millisecond,
			waitTime: 10 * time.Millisecond,
			key: "KEY",
			found: false,
		}, 
		{
			intervalTime: 5 * time.Millisecond,
			waitTime: 9 * time.Millisecond,
			key: "KEY2",
			found: true,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case: %d", i), func(t *testing.T) {
			cache := NewCache(c.intervalTime)
			time.Sleep(1 * time.Millisecond)
			cache.Add("KEY", []byte("VALUE"))

			time.Sleep(c.waitTime)

			_, ok := cache.Get("KEY")
			if ok != c.found {
				if c.found {
					t.Errorf("entry was NOT found when it was meant to")
				} else {
					t.Errorf("entry was found when it was meant to be deleted")
				}
				return
			}
		})
	}
}