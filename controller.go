package main

import (
	samplescheme "my-sample-operator/pkg/apis/samplecontroller/v1alpha1"
	clientset "my-sample-operator/pkg/generated/clientset/versioned"
	informers "my-sample-operator/pkg/generated/informers/externalversions/samplecontroller/v1alpha1"
	listers "my-sample-operator/pkg/generated/listers/samplecontroller/v1alpha1"

	appsinformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	appListers "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

const controllerAgentName = "sample-controller"

const (
	SuccessSynced         = "Synced"
	ErrResourceExists     = "ErrResourceExists"
	MessageResourceExists = "Resource %q already exists and is not managed by Foo"
	MessageResourceSynced = "Foo synced successful"
)

// controller is the controller implementation for Foo resources
type Controller struct {
	kubeClientset    kubernetes.Interface
	sampleClientset  clientset.Interface
	deploymentLister appListers.DeploymentLister
	deploymentSynced cache.InformerSynced
	fooLister        listers.FooLister
	fooSynced        cache.InformerSynced
	workqueue        workqueue.RateLimitingInterface
	recorder         record.EventRecorder
}

func NewController(
	kubeClientset kubernetes.Interface,
	exampleClientset clientset.Interface,
	kubeInformerFactory appsinformers.DeploymentInformer,
	exampleInformerFactory informers.FooInformer) *Controller {
	// Creater event broadcaster
	// Add sample-controller types to the default kubernetes Scheme so Events can be
	// logged for sample-controller types
	utilruntime.Must(samplescheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	// eventBroadcaster.StartStructuredLogging(0)
	return &Controller{}
}

func (c *Controller) Run(threadiness int, stop <-chan struct{}) error {
	return nil
}
