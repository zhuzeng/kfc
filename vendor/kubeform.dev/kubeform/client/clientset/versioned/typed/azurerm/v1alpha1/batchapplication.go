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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BatchApplicationsGetter has a method to return a BatchApplicationInterface.
// A group's client should implement this interface.
type BatchApplicationsGetter interface {
	BatchApplications(namespace string) BatchApplicationInterface
}

// BatchApplicationInterface has methods to work with BatchApplication resources.
type BatchApplicationInterface interface {
	Create(*v1alpha1.BatchApplication) (*v1alpha1.BatchApplication, error)
	Update(*v1alpha1.BatchApplication) (*v1alpha1.BatchApplication, error)
	UpdateStatus(*v1alpha1.BatchApplication) (*v1alpha1.BatchApplication, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BatchApplication, error)
	List(opts v1.ListOptions) (*v1alpha1.BatchApplicationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BatchApplication, err error)
	BatchApplicationExpansion
}

// batchApplications implements BatchApplicationInterface
type batchApplications struct {
	client rest.Interface
	ns     string
}

// newBatchApplications returns a BatchApplications
func newBatchApplications(c *AzurermV1alpha1Client, namespace string) *batchApplications {
	return &batchApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the batchApplication, and returns the corresponding batchApplication object, and an error if there is any.
func (c *batchApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.BatchApplication, err error) {
	result = &v1alpha1.BatchApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("batchapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BatchApplications that match those selectors.
func (c *batchApplications) List(opts v1.ListOptions) (result *v1alpha1.BatchApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BatchApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("batchapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested batchApplications.
func (c *batchApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("batchapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a batchApplication and creates it.  Returns the server's representation of the batchApplication, and an error, if there is any.
func (c *batchApplications) Create(batchApplication *v1alpha1.BatchApplication) (result *v1alpha1.BatchApplication, err error) {
	result = &v1alpha1.BatchApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("batchapplications").
		Body(batchApplication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a batchApplication and updates it. Returns the server's representation of the batchApplication, and an error, if there is any.
func (c *batchApplications) Update(batchApplication *v1alpha1.BatchApplication) (result *v1alpha1.BatchApplication, err error) {
	result = &v1alpha1.BatchApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("batchapplications").
		Name(batchApplication.Name).
		Body(batchApplication).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *batchApplications) UpdateStatus(batchApplication *v1alpha1.BatchApplication) (result *v1alpha1.BatchApplication, err error) {
	result = &v1alpha1.BatchApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("batchapplications").
		Name(batchApplication.Name).
		SubResource("status").
		Body(batchApplication).
		Do().
		Into(result)
	return
}

// Delete takes name of the batchApplication and deletes it. Returns an error if one occurs.
func (c *batchApplications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("batchapplications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *batchApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("batchapplications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched batchApplication.
func (c *batchApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BatchApplication, err error) {
	result = &v1alpha1.BatchApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("batchapplications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}