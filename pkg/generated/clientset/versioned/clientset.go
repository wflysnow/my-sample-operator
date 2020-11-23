/*
	may the force be with you
*/

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	samplecontrollerv1alpha1 "my-sample-operator/pkg/generated/clientset/versioned/typed/samplecontroller/v1alpha1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	SamplecontrollerV1alpha1() samplecontrollerv1alpha1.SamplecontrollerV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	samplecontrollerV1alpha1 *samplecontrollerv1alpha1.SamplecontrollerV1alpha1Client
}

// SamplecontrollerV1alpha1 retrieves the SamplecontrollerV1alpha1Client
func (c *Clientset) SamplecontrollerV1alpha1() samplecontrollerv1alpha1.SamplecontrollerV1alpha1Interface {
	return c.samplecontrollerV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.samplecontrollerV1alpha1, err = samplecontrollerv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.samplecontrollerV1alpha1 = samplecontrollerv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.samplecontrollerV1alpha1 = samplecontrollerv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
