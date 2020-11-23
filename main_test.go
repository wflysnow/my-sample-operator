package main

import (
	"fmt"
	"my-sample-operator/pkg/generated/clientset/versioned"
	"my-sample-operator/pkg/generated/clientset/versioned/typed/samplecontroller/v1alpha1"
	"my-sample-operator/pkg/generated/informers/externalversions"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func TestClient(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("192.168.15.228:6443", "E:/config/config")
	if err != nil {
		panic(err.Error())
	}
	client, err := v1alpha1.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	fooList, err := client.Foos("test").List(metav1.ListOptions{})
	fmt.Println(fooList, err)

	clientset, err := versioned.NewForConfig(config)
	factory := externalversions.NewSharedInformerFactory(clientset, 30*time.Second)
	foo, err := factory.Samplecontroller().V1alpha1().Foos().Lister().Foos("test").Get("test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(foo, err)

}

func TestWaitUtil(t *testing.T) {
	stopCh := make(chan struct{})
	i := 0
	go wait.Util(func() {
		fmt.Printf("----%d----", i)
		i++
	}, time.Second, stopCh)
}
