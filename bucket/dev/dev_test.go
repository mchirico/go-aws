package dev

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mchirico/go-aws/client"
	"io"
	"testing"
)

var count int

func More() (line []byte, err error) {
	count += 1
	var buf bytes.Buffer
	buf.Write([]byte("hello\n"))
	if count > 10 {
		return nil, io.EOF
	}
	return buf.Bytes(), nil
}

func Test_Upload(t *testing.T) {

	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	var buf bytes.Buffer
	buf.Write([]byte(in))
	s := &Stream{
		Buffer: buf,
		F:      More,
		Reader: bytes.NewReader([]byte(in)),
	}

	_ = s
	r, err := Upload(context.TODO(),
		client.Config(),
		"sharepoint-poc", "key", s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
