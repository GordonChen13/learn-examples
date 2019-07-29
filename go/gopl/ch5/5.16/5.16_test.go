package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrJoin(t *testing.T) {
	cases := []struct{
		name string
		sep string
		input []string
		want string
	} {
		{"empty", ",", []string{}, ""},
		{"empty sep", "", []string{"a", "b"}, "ab"},
		{"multi", ",", []string{"a", "b", "c"}, "a,b,c"},
		{"multi with sep", ",", []string{"a,b", ",b", "c,"}, "a,b,,b,c,"},
		{"multi with chinese", "，", []string{"你好", "世界", "啊啊"}, "你好，世界，啊啊"},
	}

	for _,c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := StrJoin(c.sep, c.input...)
			assert.Equal(t, c.want, got)
		})
	}
}
