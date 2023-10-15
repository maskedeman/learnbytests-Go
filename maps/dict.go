package maps

const (
	ErrNotFound          = DictionaryErr("cant find the word you are looking for")
	ErrWordExists        = DictionaryErr("can't add words, it already exists")
	ErrWordDoesNotExists = DictionaryErr("can't update word cause it doesn't exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
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
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = definition
	default:
		return err

	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
