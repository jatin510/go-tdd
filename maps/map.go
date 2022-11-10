package maps

import "errors"

type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil

}
