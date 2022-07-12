/*
iam_cmb ... iam combination. FIXME (mmc), come up with
a better name.

*/

package iam

import (
	"fmt"
	"testing"
)

func TestWriteFile(t *testing.T) {
	data := []byte("junk\n1\n1")
	file := "/Users/mchirico/junk2.txt"
	err := WriteFile(file, data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestI_ListAccessKeys(t *testing.T) {
	i := NewI("k8s")
	result, err := i.ListAccessKeys()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestI_ListAttachedGroupPolicies(t *testing.T) {
	i := NewI("k8s")
	result, err := i.ListAttachedGroupPolicies("k8s")
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result.AttachedPolicies {
		fmt.Println(*v.PolicyName, *v.PolicyArn)
	}

}

func TestI_AccessReport(t *testing.T) {
	i := NewI("k8s")
	result, err := i.AccessReport()
	if err != nil {
		t.Fatal(err)
	}
	WriteFile("/tmp/report.csv", result.Content)
	fmt.Println(string(result.Content))

}
