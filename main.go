package main

import (
	"fmt"

	"github.com/wflysnow/my-sample-operator/pkg/apis/samplecontroller/v1alpha1"
)

func main() {
	temVar := &v1alpha1.Foo{}
	fmt.Println("this is sample-operator project #v", temVar)
}
