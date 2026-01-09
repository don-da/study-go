package maps

type Dictionary map[string]string

const (
	ErrUnknownWord      = DictionaryErr("could not find the word you are looking for")
	ErrWordExists       = DictionaryErr("word already exists in dictionary!")
	ErrWordDoesNotExist = DictionaryErr("word doesn't exist!")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]
	if !ok {
		return "", ErrUnknownWord
	}

	return word, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrUnknownWord:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrUnknownWord:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrUnknownWord:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
