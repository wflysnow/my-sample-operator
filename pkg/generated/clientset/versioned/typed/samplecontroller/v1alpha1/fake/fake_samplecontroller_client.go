/*
	may the force be with you
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "my-sample-operator/pkg/generated/clientset/versioned/typed/samplecontroller/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSamplecontrollerV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSamplecontrollerV1alpha1) Foos(namespace string) v1alpha1.FooInterface {
	return &FakeFoos{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSamplecontrollerV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}