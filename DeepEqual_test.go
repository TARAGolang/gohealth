package gohealth

import (
	"fmt"
	"reflect"
	"testing"
)

func DeepEqual(expected, obtained interface{}, t *testing.T) {
	if !reflect.DeepEqual(expected, obtained) {

		s := fmt.Sprintf(`Not deep equal:
Expected: %#v
Obtained: %#v
`, expected, obtained)

		t.Error(s)
	}
}
