package kinesis

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestNewP(t *testing.T) {
	p := NewP("stream0")
	p.Create()

	p.Put("key", []byte("1. Data 1 2 3..."))
	p.Put("key", []byte("2. Data 1 2 3..."))
	p.Put("key", []byte("3. Data 1 2 3..."))

}

func TestNewPCreateonly(t *testing.T) {
	p := NewP("mmc")
	p.Create()

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
	p := NewP("stream0")

	for i := 0; i <= 12; i++ {
		p.Put(fmt.Sprintf("keyX:%d", i), []byte(fmt.Sprintf("3. Data 1 2 3...%s", time.Now().Format(time.Kitchen))))
	}

}

func Test_put2(t *testing.T) {
	p := NewP("mmc2")
	p.Put("key-lambda", []byte("Lambda *** TEST ***"))
	p.Put("key", []byte("2. Data 1 2 3..."))
	p.Put("key", []byte("3. Data 1 2 3..."))
	for i := 0; i <= 12; i++ {
		p.Put("keyX", []byte(fmt.Sprintf("3. Data 1 2 3...%s", time.Now().Format(time.Kitchen))))
	}

}

func Test_Get(t *testing.T) {
	p := NewP("stream0")
	result, err := p.Get()
	if err != nil {
		t.FailNow()
	}
	pkey := []string{}
	for _, v := range result.Records {

		//fmt.Println(string(v.Data))
		//fmt.Println(*v.PartitionKey)
		pkey = append(pkey, *v.PartitionKey)
		for i, v := range pkey {
			fmt.Println(i, v)
		}
		//fmt.Println(*v.SequenceNumber)
	}

	sort.Slice(pkey, func(i, j int) bool {
		return pkey[j] < pkey[i]
	})
	fmt.Println(pkey)
}

func Test_Register(t *testing.T) {
	p := NewP("mmc2")
	streamARN, err := p.StreamARN()
	if err != nil {
		t.FailNow()
	}
	fmt.Println(*streamARN)
	result, err := p.Register("prog2", *streamARN)
	if err != nil {
		t.FailNow()
	}
	p.SubscribeToShard()
	fmt.Println(*result.Consumer)
}

func Test_Subscribe(t *testing.T) {
	p := NewP("mmc2")
	result, err := p.SubscribeToShard()
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)

}

func Test_Delete(t *testing.T) {
	p := NewP("mmc")
	p.Delete()

	p2 := NewP("mmc2")
	p2.Delete()
}

func Test_Delete1(t *testing.T) {
	p := NewP("mmc")
	p.Delete()

}
