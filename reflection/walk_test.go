package walk

import (
	"reflect"
	"testing"
)

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
