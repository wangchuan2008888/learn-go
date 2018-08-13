package my_font

import (
	"sort"
	"errors"
)

type Font struct {
	family fontFamily
	size   int
}

type fontFamily struct {
	name string
}

var fonts = [] string{"aaa", "bbb", "ccc"}

func new_font(name string) (*fontFamily, error) {
	index := sort.SearchStrings(fonts, name)
	if index == len(fonts) {
		return nil, errors.New("invalide font name")
	} else {
		return &fontFamily{name: name}, nil
	}
}

func New(name string, size int) (*Font, error) {
	font, err := new_font(name)
	if err != nil || size<5 || size>255 {
		return nil, err
	} else {
		return &Font{*font, size}, nil

	}
}
