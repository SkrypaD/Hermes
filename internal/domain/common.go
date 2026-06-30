package domain

import "strconv"

type ID int64

func (i *ID) Convert(str_val string) (ID, error) {
	val, err := strconv.ParseInt(str_val, 10, 64)
	if err != nil {
		return ID(-1), err
	}
	*i = ID(val)
	return *i, nil
}
