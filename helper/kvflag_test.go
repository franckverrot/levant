package helper

import (
	"reflect"
	"testing"
)

func TestHelper_Set(t *testing.T) {
	cases := []struct {
		Input  string
		Output map[string]string
		Error  bool
	}{
		{
			"key=value",
			map[string]string{"key": "value"},
			false,
		},

		{
			"key=",
			map[string]string{"key": ""},
			false,
		},

		{
			"key=foo=bar",
			map[string]string{"key": "foo=bar"},
			false,
		},

		{
			"key",
			nil,
			true,
		},
	}

	for _, tc := range cases {
		f := new(Flag)
		err := f.Set(tc.Input)
		if (err != nil) != tc.Error {
			t.Fatalf("bad error. Input: %#v", tc.Input)
		}

		actual := map[string]string(*f)
		if !reflect.DeepEqual(actual, tc.Output) {
			t.Fatalf("bad: %#v", actual)
		}
	}
}
