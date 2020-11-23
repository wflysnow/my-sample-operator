package main

import (
	"flag"
	clientset "my-sample-operator/pkg/generated/clientset/versioned"
	"my-sample-operator/pkg/generated/informers/externalversions"
	"my-sample-operator/pkg/signals"
	"time"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

var (
	masterURL  string
	kubeconfig string
)

func init() {
	flag.StringVar(&masterURL, "master", "", "The address of the kubernetes API server. Overrides any value in kubeconfig.Only required if out-of-cluster.")
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig.Only requied if out-of-cluter.")
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	stopCh := signals.SetupSignalHandler()
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		// panic(err.Error())
		klog.Fatalf("Error building kubeconfig :%s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}
	exampleClient, err := clientset.NewForConfig(cfg)
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, 30*time.Second)
	exampleInformerFactory := externalversions.NewSharedInformerFactory(exampleClient, 39*time.Second)
	controller := NewController(kubeClient,
		exampleClient,
		kubeInformerFactory.Apps().V1().Deployments(),
		exampleInformerFactory.Samplecontroller().V1alpha1().Foos())
	kubeInformerFactory.Start(stopCh)
	exampleInformerFactory.Start(stopCh)
	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error runing controller: %s", err.Error())
	}
}
