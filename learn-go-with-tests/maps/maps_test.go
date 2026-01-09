package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown word")
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrUnknownWord)
	})
}

func TestAdd(t *testing.T) {
	word, definition := "test", "this is a test"

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("word already exists", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word, definition := "test", "this is a test"
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "delete test"}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		assertError(t, err, ErrUnknownWord)
	})

	t.Run("non-existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}
