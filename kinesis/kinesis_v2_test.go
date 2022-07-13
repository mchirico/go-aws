package kinesis

import (
	"fmt"
	"testing"
	"time"
)

func TestNewP(t *testing.T) {
	p := NewP("mmc")
	p.Create()

	p.Put("key", []byte("1. Data 1 2 3..."))
	p.Put("key", []byte("2. Data 1 2 3..."))
	p.Put("key", []byte("3. Data 1 2 3..."))

}

func TestNewP2(t *testing.T) {
	p := NewP("mmc2")
	_, err := p.Create()
	if err != nil {
		t.FailNow()
	}

	p.Put("key", []byte("1. Data 1 2 3..."))
	p.Put("key", []byte("2. Data 1 2 3..."))
	p.Put("key3", []byte("3. Data 1 2 3..."))

}



func Test_put(t *testing.T) {
	p := NewP("mmc")
	p.Put("key", []byte("1. Data 1 2 3..."))
	p.Put("key", []byte("2. Data 1 2 3..."))
	p.Put("key", []byte("3. Data 1 2 3..."))
	for i := 0; i <= 12; i++ {
		p.Put("keyX", []byte(fmt.Sprintf("3. Data 1 2 3...%s", time.Now().Format(time.Kitchen))))
	}

}

func Test_put2(t *testing.T) {
	p := NewP("mmc2")
	p.Put("key", []byte("1. Data 1 2 3..."))
	p.Put("key", []byte("2. Data 1 2 3..."))
	p.Put("key", []byte("3. Data 1 2 3..."))
	for i := 0; i <= 12; i++ {
		p.Put("keyX", []byte(fmt.Sprintf("3. Data 1 2 3...%s", time.Now().Format(time.Kitchen))))
	}

}

func Test_Get(t *testing.T) {
	p := NewP("mmc")
	result, err := p.Get()
	if err != nil {
		t.FailNow()
	}
	for _, v := range result.Records {

		fmt.Println(string(v.Data))
		fmt.Println(*v.PartitionKey)
		fmt.Println(*v.SequenceNumber)
	}

}

func Test_Delete(t *testing.T) {
	p := NewP("mmc")
	p.Delete()

	p2 := NewP("mmc2")
	p2.Delete()
}
