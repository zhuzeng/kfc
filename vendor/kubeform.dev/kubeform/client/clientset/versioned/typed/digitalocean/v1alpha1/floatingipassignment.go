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

	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FloatingIPAssignmentsGetter has a method to return a FloatingIPAssignmentInterface.
// A group's client should implement this interface.
type FloatingIPAssignmentsGetter interface {
	FloatingIPAssignments(namespace string) FloatingIPAssignmentInterface
}

// FloatingIPAssignmentInterface has methods to work with FloatingIPAssignment resources.
type FloatingIPAssignmentInterface interface {
	Create(*v1alpha1.FloatingIPAssignment) (*v1alpha1.FloatingIPAssignment, error)
	Update(*v1alpha1.FloatingIPAssignment) (*v1alpha1.FloatingIPAssignment, error)
	UpdateStatus(*v1alpha1.FloatingIPAssignment) (*v1alpha1.FloatingIPAssignment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FloatingIPAssignment, error)
	List(opts v1.ListOptions) (*v1alpha1.FloatingIPAssignmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FloatingIPAssignment, err error)
	FloatingIPAssignmentExpansion
}

// floatingIPAssignments implements FloatingIPAssignmentInterface
type floatingIPAssignments struct {
	client rest.Interface
	ns     string
}

// newFloatingIPAssignments returns a FloatingIPAssignments
func newFloatingIPAssignments(c *DigitaloceanV1alpha1Client, namespace string) *floatingIPAssignments {
	return &floatingIPAssignments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the floatingIPAssignment, and returns the corresponding floatingIPAssignment object, and an error if there is any.
func (c *floatingIPAssignments) Get(name string, options v1.GetOptions) (result *v1alpha1.FloatingIPAssignment, err error) {
	result = &v1alpha1.FloatingIPAssignment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("floatingipassignments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FloatingIPAssignments that match those selectors.
func (c *floatingIPAssignments) List(opts v1.ListOptions) (result *v1alpha1.FloatingIPAssignmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FloatingIPAssignmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("floatingipassignments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested floatingIPAssignments.
func (c *floatingIPAssignments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("floatingipassignments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a floatingIPAssignment and creates it.  Returns the server's representation of the floatingIPAssignment, and an error, if there is any.
func (c *floatingIPAssignments) Create(floatingIPAssignment *v1alpha1.FloatingIPAssignment) (result *v1alpha1.FloatingIPAssignment, err error) {
	result = &v1alpha1.FloatingIPAssignment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("floatingipassignments").
		Body(floatingIPAssignment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a floatingIPAssignment and updates it. Returns the server's representation of the floatingIPAssignment, and an error, if there is any.
func (c *floatingIPAssignments) Update(floatingIPAssignment *v1alpha1.FloatingIPAssignment) (result *v1alpha1.FloatingIPAssignment, err error) {
	result = &v1alpha1.FloatingIPAssignment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("floatingipassignments").
		Name(floatingIPAssignment.Name).
		Body(floatingIPAssignment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *floatingIPAssignments) UpdateStatus(floatingIPAssignment *v1alpha1.FloatingIPAssignment) (result *v1alpha1.FloatingIPAssignment, err error) {
	result = &v1alpha1.FloatingIPAssignment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("floatingipassignments").
		Name(floatingIPAssignment.Name).
		SubResource("status").
		Body(floatingIPAssignment).
		Do().
		Into(result)
	return
}

// Delete takes name of the floatingIPAssignment and deletes it. Returns an error if one occurs.
func (c *floatingIPAssignments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("floatingipassignments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *floatingIPAssignments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("floatingipassignments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched floatingIPAssignment.
func (c *floatingIPAssignments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FloatingIPAssignment, err error) {
	result = &v1alpha1.FloatingIPAssignment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("floatingipassignments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
