package maps

import (
	"testing"
)

func TestGet(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is test",
	}

	t.Run("test get key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is test"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("test get non existent key", func(t *testing.T) {
		_, err := dictionary.Search("not exist")
		if err == nil {
			t.Fatal("expected to get an error.")
		}
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is test",
	}

	t.Run("test add new key", func(t *testing.T) {
		dictionary.Add("test", "this is test")
		got, _ := dictionary.Search("test")
		want := "this is test"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("test add existing key", func(t *testing.T) {
		err := dictionary.Add("test", "this is new test")
		if err == nil {
			t.Fatal("is has to be an error while adding an existing key")
		}
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is test",
	}

	t.Run("test update existing key", func(t *testing.T) {
		dictionary.Update("test", "new value")
		got, _ := dictionary.Search("test")
		want := "new value"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("test update new key", func(t *testing.T) {
		err := dictionary.Update("new key", "new value")
		if err == nil {
			t.Fatal("is has to be an error while updating new key")
		}
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is test",
	}

	t.Run("test delete existing key", func(t *testing.T) {
		dictionary.Delete("test")
		_, err := dictionary.Search("test")
		if err == nil {
			t.Fatal("should have been deleted")
		}
	})

	t.Run("test delete existing key", func(t *testing.T) {
		err := dictionary.Delete("someKey")
		if err == nil {
			t.Fatal("key does not exist")
		}
	})
}
