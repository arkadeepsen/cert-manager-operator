// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/openshift/cert-manager-operator/pkg/operator/clientset/versioned/typed/operator/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeOperatorV1alpha1 struct {
	*testing.Fake
}

func (c *FakeOperatorV1alpha1) CertManagers() v1alpha1.CertManagerInterface {
	return &FakeCertManagers{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeOperatorV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
