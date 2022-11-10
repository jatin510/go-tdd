package maps

type Dictionary map[string]string

type dictionaryErr string

func (e dictionaryErr) Error() string {
	return string(e)
}

var (
	ErrNotFound   = dictionaryErr("could not find the word you were looking for")
	ErrWordExists = dictionaryErr("cannot add word because it already exists")
)

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
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
