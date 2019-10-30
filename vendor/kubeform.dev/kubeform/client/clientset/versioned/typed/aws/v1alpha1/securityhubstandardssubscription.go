/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SecurityhubStandardsSubscriptionsGetter has a method to return a SecurityhubStandardsSubscriptionInterface.
// A group's client should implement this interface.
type SecurityhubStandardsSubscriptionsGetter interface {
	SecurityhubStandardsSubscriptions(namespace string) SecurityhubStandardsSubscriptionInterface
}

// SecurityhubStandardsSubscriptionInterface has methods to work with SecurityhubStandardsSubscription resources.
type SecurityhubStandardsSubscriptionInterface interface {
	Create(*v1alpha1.SecurityhubStandardsSubscription) (*v1alpha1.SecurityhubStandardsSubscription, error)
	Update(*v1alpha1.SecurityhubStandardsSubscription) (*v1alpha1.SecurityhubStandardsSubscription, error)
	UpdateStatus(*v1alpha1.SecurityhubStandardsSubscription) (*v1alpha1.SecurityhubStandardsSubscription, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SecurityhubStandardsSubscription, error)
	List(opts v1.ListOptions) (*v1alpha1.SecurityhubStandardsSubscriptionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecurityhubStandardsSubscription, err error)
	SecurityhubStandardsSubscriptionExpansion
}

// securityhubStandardsSubscriptions implements SecurityhubStandardsSubscriptionInterface
type securityhubStandardsSubscriptions struct {
	client rest.Interface
	ns     string
}

// newSecurityhubStandardsSubscriptions returns a SecurityhubStandardsSubscriptions
func newSecurityhubStandardsSubscriptions(c *AwsV1alpha1Client, namespace string) *securityhubStandardsSubscriptions {
	return &securityhubStandardsSubscriptions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the securityhubStandardsSubscription, and returns the corresponding securityhubStandardsSubscription object, and an error if there is any.
func (c *securityhubStandardsSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.SecurityhubStandardsSubscription, err error) {
	result = &v1alpha1.SecurityhubStandardsSubscription{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securityhubstandardssubscriptions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecurityhubStandardsSubscriptions that match those selectors.
func (c *securityhubStandardsSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.SecurityhubStandardsSubscriptionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SecurityhubStandardsSubscriptionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securityhubstandardssubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested securityhubStandardsSubscriptions.
func (c *securityhubStandardsSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("securityhubstandardssubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a securityhubStandardsSubscription and creates it.  Returns the server's representation of the securityhubStandardsSubscription, and an error, if there is any.
func (c *securityhubStandardsSubscriptions) Create(securityhubStandardsSubscription *v1alpha1.SecurityhubStandardsSubscription) (result *v1alpha1.SecurityhubStandardsSubscription, err error) {
	result = &v1alpha1.SecurityhubStandardsSubscription{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("securityhubstandardssubscriptions").
		Body(securityhubStandardsSubscription).
		Do().
		Into(result)
	return
}

// Update takes the representation of a securityhubStandardsSubscription and updates it. Returns the server's representation of the securityhubStandardsSubscription, and an error, if there is any.
func (c *securityhubStandardsSubscriptions) Update(securityhubStandardsSubscription *v1alpha1.SecurityhubStandardsSubscription) (result *v1alpha1.SecurityhubStandardsSubscription, err error) {
	result = &v1alpha1.SecurityhubStandardsSubscription{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("securityhubstandardssubscriptions").
		Name(securityhubStandardsSubscription.Name).
		Body(securityhubStandardsSubscription).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *securityhubStandardsSubscriptions) UpdateStatus(securityhubStandardsSubscription *v1alpha1.SecurityhubStandardsSubscription) (result *v1alpha1.SecurityhubStandardsSubscription, err error) {
	result = &v1alpha1.SecurityhubStandardsSubscription{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("securityhubstandardssubscriptions").
		Name(securityhubStandardsSubscription.Name).
		SubResource("status").
		Body(securityhubStandardsSubscription).
		Do().
		Into(result)
	return
}

// Delete takes name of the securityhubStandardsSubscription and deletes it. Returns an error if one occurs.
func (c *securityhubStandardsSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securityhubstandardssubscriptions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *securityhubStandardsSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securityhubstandardssubscriptions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched securityhubStandardsSubscription.
func (c *securityhubStandardsSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecurityhubStandardsSubscription, err error) {
	result = &v1alpha1.SecurityhubStandardsSubscription{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("securityhubstandardssubscriptions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}