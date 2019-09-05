package testenabled

import (
	"os"
	"reflect"
	"testing"

	"4d63.com/testparallel"
)

func TestMain(m *testing.M) {
	testparallel.All(m)
	os.Exit(m.Run())
}

func TestAll(t *testing.T) {
	valueM := reflect.ValueOf(t).Elem()
	isParallelField := valueM.FieldByName("isParallel")
	isParallel := isParallelField.Bool()

	if isParallel {
		t.Logf("isParallel is %v", isParallel)
	} else {
		t.Errorf("isParallel is %v, want true", isParallel)
	}
}
