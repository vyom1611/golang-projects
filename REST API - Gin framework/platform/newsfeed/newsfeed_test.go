package newsfeed

import (
	"testing"
)

//Testing all handler functions for the API

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{"an item", "with body"})

	if len(feed.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{"an item", "with body"})
	results := feed.GetAll()

	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
