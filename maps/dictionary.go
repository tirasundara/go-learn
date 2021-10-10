package maps

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("Could not find you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}
