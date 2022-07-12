package kinesis

import (
	"github.com/mchirico/go-aws/kinesis"
)

func Put() {
	p := kinesis.NewP("mmc")
	p.Put("lambda", []byte("1. Data 1 2 3..."))

}
