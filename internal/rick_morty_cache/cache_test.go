package rickmortycache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Errorf("Error creating the cache")
	}
}

func TestAddGetFromCache(t *testing.T) {
	cache := NewCache(time.Minute)

	cases := []struct {
		key string
		expected []byte 
	}{
		{
			key: "key1",
			expected: []byte("value1"),
		},
		{
			key: "key2",
			expected: []byte("value2"),
		},
	}

	for _ , cs := range cases {
		cache.Add(cs.key, cs.expected)
		actual, ok := cache.Get(cs.key)
		if !ok {
			t.Errorf("Key not found: %s", cs.key)
		}

		if string(actual) != string(cs.expected) {
			t.Errorf("Value returned is incorrect: %s vs %s", string(actual), string(cs.expected))
		}
	}
}

func TestReapLoop(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	key := "key1"
	value := "value1"
	cache.Add(key, []byte(value))

	time.Sleep(interval + time.Millisecond)

	_ , ok := cache.Get(key)
	if ok {
		t.Errorf("The key should not be present: %s", key)
	}

	key = "key2"
	value = "value2"

	cache.Add(key, []byte(value))

	time.Sleep(interval / 2)

	_ , ok = cache.Get(key)
	if !ok {
		t.Errorf("The key should be present: %s", key)
	}
}