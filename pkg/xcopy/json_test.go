package xcopy_test

import (
	"reflect"
	"testing"

	"github.com/agussyahrilmubarok/gox/pkg/xcopy"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJSONCopy(t *testing.T) {
	original := Person{Name: "Alice", Age: 30}
	var copied Person

	xcopy.ByJSON(&original, &copied)

	if !reflect.DeepEqual(original, copied) {
		t.Errorf("Expected %+v, got %+v", original, copied)
	}

	// Modify original and check copied remains unchanged
	original.Name = "Bob"
	original.Age = 40

	if copied.Name != "Alice" || copied.Age != 30 {
		t.Errorf("Copied struct changed unexpectedly: %+v", copied)
	}
}
