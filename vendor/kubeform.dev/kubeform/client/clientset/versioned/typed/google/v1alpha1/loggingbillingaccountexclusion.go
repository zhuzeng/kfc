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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LoggingBillingAccountExclusionsGetter has a method to return a LoggingBillingAccountExclusionInterface.
// A group's client should implement this interface.
type LoggingBillingAccountExclusionsGetter interface {
	LoggingBillingAccountExclusions(namespace string) LoggingBillingAccountExclusionInterface
}

// LoggingBillingAccountExclusionInterface has methods to work with LoggingBillingAccountExclusion resources.
type LoggingBillingAccountExclusionInterface interface {
	Create(*v1alpha1.LoggingBillingAccountExclusion) (*v1alpha1.LoggingBillingAccountExclusion, error)
	Update(*v1alpha1.LoggingBillingAccountExclusion) (*v1alpha1.LoggingBillingAccountExclusion, error)
	UpdateStatus(*v1alpha1.LoggingBillingAccountExclusion) (*v1alpha1.LoggingBillingAccountExclusion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LoggingBillingAccountExclusion, error)
	List(opts v1.ListOptions) (*v1alpha1.LoggingBillingAccountExclusionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoggingBillingAccountExclusion, err error)
	LoggingBillingAccountExclusionExpansion
}

// loggingBillingAccountExclusions implements LoggingBillingAccountExclusionInterface
type loggingBillingAccountExclusions struct {
	client rest.Interface
	ns     string
}

// newLoggingBillingAccountExclusions returns a LoggingBillingAccountExclusions
func newLoggingBillingAccountExclusions(c *GoogleV1alpha1Client, namespace string) *loggingBillingAccountExclusions {
	return &loggingBillingAccountExclusions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the loggingBillingAccountExclusion, and returns the corresponding loggingBillingAccountExclusion object, and an error if there is any.
func (c *loggingBillingAccountExclusions) Get(name string, options v1.GetOptions) (result *v1alpha1.LoggingBillingAccountExclusion, err error) {
	result = &v1alpha1.LoggingBillingAccountExclusion{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loggingbillingaccountexclusions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LoggingBillingAccountExclusions that match those selectors.
func (c *loggingBillingAccountExclusions) List(opts v1.ListOptions) (result *v1alpha1.LoggingBillingAccountExclusionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LoggingBillingAccountExclusionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loggingbillingaccountexclusions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loggingBillingAccountExclusions.
func (c *loggingBillingAccountExclusions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("loggingbillingaccountexclusions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a loggingBillingAccountExclusion and creates it.  Returns the server's representation of the loggingBillingAccountExclusion, and an error, if there is any.
func (c *loggingBillingAccountExclusions) Create(loggingBillingAccountExclusion *v1alpha1.LoggingBillingAccountExclusion) (result *v1alpha1.LoggingBillingAccountExclusion, err error) {
	result = &v1alpha1.LoggingBillingAccountExclusion{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("loggingbillingaccountexclusions").
		Body(loggingBillingAccountExclusion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a loggingBillingAccountExclusion and updates it. Returns the server's representation of the loggingBillingAccountExclusion, and an error, if there is any.
func (c *loggingBillingAccountExclusions) Update(loggingBillingAccountExclusion *v1alpha1.LoggingBillingAccountExclusion) (result *v1alpha1.LoggingBillingAccountExclusion, err error) {
	result = &v1alpha1.LoggingBillingAccountExclusion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loggingbillingaccountexclusions").
		Name(loggingBillingAccountExclusion.Name).
		Body(loggingBillingAccountExclusion).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *loggingBillingAccountExclusions) UpdateStatus(loggingBillingAccountExclusion *v1alpha1.LoggingBillingAccountExclusion) (result *v1alpha1.LoggingBillingAccountExclusion, err error) {
	result = &v1alpha1.LoggingBillingAccountExclusion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loggingbillingaccountexclusions").
		Name(loggingBillingAccountExclusion.Name).
		SubResource("status").
		Body(loggingBillingAccountExclusion).
		Do().
		Into(result)
	return
}

// Delete takes name of the loggingBillingAccountExclusion and deletes it. Returns an error if one occurs.
func (c *loggingBillingAccountExclusions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loggingbillingaccountexclusions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loggingBillingAccountExclusions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loggingbillingaccountexclusions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched loggingBillingAccountExclusion.
func (c *loggingBillingAccountExclusions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoggingBillingAccountExclusion, err error) {
	result = &v1alpha1.LoggingBillingAccountExclusion{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("loggingbillingaccountexclusions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
