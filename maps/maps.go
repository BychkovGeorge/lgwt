package maps

import (
	"errors"
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if ok {
		return val, nil
	}
	return "", errors.New("value not exist")
}

func (d Dictionary) Add(word, value string) error {
	_, err := d.Search(word)
	if err == nil {
		return errors.New("value already exists")
	}
	d[word] = value
	return nil
}

func (d Dictionary) Update(word, value string) error {
	_, err := d.Search(word)
	if err != nil {
		return errors.New("value not exist")
	}
	d[word] = value
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return errors.New("value not exist")
	}
	delete(d, word)
	return nil
}
