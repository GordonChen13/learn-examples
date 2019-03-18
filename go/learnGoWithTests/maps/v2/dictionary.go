package v2

import "errors"

type Dictionary map[string]string
var NotFoundError = errors.New("word not found")
var DuplicateWordError = errors.New("duplicate word error")
var UpdateNotExitsError = errors.New("update not exists word error")
var DeleteNotExistsError = errors.New("delete not exists word error")

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", NotFoundError
	}

	return value, nil
}

func (d Dictionary) Add(word, value string) (error)  {
	if _, ok := d[word]; ok {
		return  DuplicateWordError
	}
	d[word] = value
	return nil
}

func (d Dictionary) Update(word, value string) (error)  {
	if _, ok := d[word]; !ok {
		return  UpdateNotExitsError
	}
	d[word] = value
	return nil
}

func (d Dictionary) Delete(word string) (error)  {
	if _, ok := d[word]; !ok {
		return  DeleteNotExistsError
	}
    delete(d, word)
	return nil
}
