package reflections

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
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
			}{"Tira"},
			[]string{"Tira"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{
				"John",
				"London",
			},
			[]string{"John", "London"},
		},
		{
			"Struct with non-string field",
			struct {
				Name string
				Age  int
			}{
				"Bob",
				33,
			},
			[]string{"Bob"},
		},
		{
			"Nested fields",
			Person{
				"Bob",
				Profile{
					33,
					"London",
				},
			},
			[]string{"Bob", "London"},
		},
		{
			"Pointers to thing",
			&Person{
				"Bob",
				Profile{
					33,
					"London",
				},
			},
			[]string{"Bob", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{27, "NY"},
			},
			[]string{"London", "NY"},
		},
	}

	for _, test := range cases {
		var got []string

		t.Run(test.Name, func(t *testing.T) {
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
