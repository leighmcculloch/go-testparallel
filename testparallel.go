// Package testparallel provides a function for making all tests in a package
// run in Parallel without needing to add the line `t.Parallel()` to the
// beginning of every test in a package.
//
// Call All inside TestMain to enable parallel for all top-level tests:
//
//		func TestMain(m *testing.M) {
//			testparallel.All(m)
//			os.Exit(m.Run())
//		}
//
// Warning: Uses unsafe and is dependent on unexported fields, and may become
// unpredictable or not work with changes to the Go standard library. Use at
// your own risk.
package testparallel

import (
	"reflect"
	"testing"
	"unsafe"
)

// All wraps all the top-level tests in the testing.M with a function that
// first calls t.Parallel().
//
// Sub-tests are unaffected and still run in sequence.
//
// Call All inside TestMain:
//
//		func TestMain(m *testing.M) {
//			testparallel.All(m)
//			os.Exit(m.Run())
//		}
//
// Warning: Uses unsafe and is dependent on unexported fields, and may become
// unpredictable or not work with changes to the Go standard library. Use at
// your own risk.
func All(m *testing.M) {
	valueM := reflect.ValueOf(m).Elem()
	testsField := valueM.FieldByName("tests")
	tests := reflect.NewAt(testsField.Type(), unsafe.Pointer(testsField.UnsafeAddr())).Elem().Interface().([]testing.InternalTest)
	for i := range tests {
		f := tests[i].F
		tests[i].F = func(t *testing.T) {
			t.Parallel()
			f(t)
		}
	}
}
