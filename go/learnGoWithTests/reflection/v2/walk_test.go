package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}
type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string field",
			struct {
				Name string
				City string
			}{"Chris", "ShenZhen"},
			[]string{"Chris", "ShenZhen"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{"Chris", Profile{33, "ShenZhen"}},
			[]string{"Chris", "ShenZhen"},
		},
		{
			"Arrays",
			[2]Profile {
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Pointer to things",
			&Person{"Chris", Profile{33, "ShenZhen"}},
			[]string{"Chris", "ShenZhen"},
		},
		{
			"Slices",
			&Person{"Chris", Profile{33, "ShenZhen"}},
			[]string{"Chris", "ShenZhen"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
