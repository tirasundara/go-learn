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
		{
			"Array",
			[2]Profile{
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

	t.Run("with Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		got := []string{}

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with Channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
