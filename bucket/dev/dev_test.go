package dev

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/mchirico/go-aws/client"
	"strings"
	"testing"
)

func Test_Upload(t *testing.T) {

	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`

	s := &Stream{
		Reader: csv.NewReader(strings.NewReader(in)),
	}

	r, err := Upload(context.TODO(),
		client.Config(),
		"bucket", "key", s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
