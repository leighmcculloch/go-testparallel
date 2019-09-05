package testdisabled

import (
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	valueM := reflect.ValueOf(t).Elem()
	isParallelField := valueM.FieldByName("isParallel")
	isParallel := isParallelField.Bool()

	if !isParallel {
		t.Logf("isParallel is %v", isParallel)
	} else {
		t.Errorf("isParallel is %v, want true", isParallel)
	}
}
