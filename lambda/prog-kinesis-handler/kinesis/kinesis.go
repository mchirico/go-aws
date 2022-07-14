package kinesis

import (
	"fmt"
	"github.com/mchirico/go-aws/kinesis"
	"log"
	"time"
)

func Put(data []byte) {
	p := kinesis.NewP("mmc")
	now := time.Now().Format(time.RFC822Z)
	_, err := p.Put(fmt.Sprint("%s: lambda",now), data)
	if err != nil {
		log.Println(err)
	}

}
