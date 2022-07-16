package dynamoDB

import (
	"fmt"
	"testing"
)

func TestD_List(t *testing.T) {
	d := NewD()
	result, err := d.List()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
