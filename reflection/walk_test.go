package walk

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
	t.Run("Test walk function", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"chris"},
				[]string{"chris"},
			},
			{
				"struct with two string field",
				struct {
					Name string
					City string
				}{"chris", "london"},
				[]string{"chris", "london"},
			},
			{
				"struct with non string field",
				struct {
					Name string
					age  int
				}{"chris", 24},
				[]string{"chris"},
			},
			{
				"nested fields",
				Person{
					"chris",
					struct {
						Age  int
						City string
					}{23, "london"},
				},
				[]string{"chris", "london"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				Walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}
			})
		}
	})
}
