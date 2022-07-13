package kinesis

import (
	"fmt"
	"github.com/mchirico/go-aws/kinesis"
	"log"
	"time"
)

func Put() {
	p := kinesis.NewP("mmc")
	now := time.Now().Format(time.RFC822Z)
	_, err := p.Put("lambda", []byte(fmt.Sprintf("%s: 1. Data 1 2 3...", now)))
	if err != nil {
		log.Println(err)
	}

}
