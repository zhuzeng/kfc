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

// BigqueryDatasetsGetter has a method to return a BigqueryDatasetInterface.
// A group's client should implement this interface.
type BigqueryDatasetsGetter interface {
	BigqueryDatasets(namespace string) BigqueryDatasetInterface
}

// BigqueryDatasetInterface has methods to work with BigqueryDataset resources.
type BigqueryDatasetInterface interface {
	Create(*v1alpha1.BigqueryDataset) (*v1alpha1.BigqueryDataset, error)
	Update(*v1alpha1.BigqueryDataset) (*v1alpha1.BigqueryDataset, error)
	UpdateStatus(*v1alpha1.BigqueryDataset) (*v1alpha1.BigqueryDataset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BigqueryDataset, error)
	List(opts v1.ListOptions) (*v1alpha1.BigqueryDatasetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BigqueryDataset, err error)
	BigqueryDatasetExpansion
}

// bigqueryDatasets implements BigqueryDatasetInterface
type bigqueryDatasets struct {
	client rest.Interface
	ns     string
}

// newBigqueryDatasets returns a BigqueryDatasets
func newBigqueryDatasets(c *GoogleV1alpha1Client, namespace string) *bigqueryDatasets {
	return &bigqueryDatasets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bigqueryDataset, and returns the corresponding bigqueryDataset object, and an error if there is any.
func (c *bigqueryDatasets) Get(name string, options v1.GetOptions) (result *v1alpha1.BigqueryDataset, err error) {
	result = &v1alpha1.BigqueryDataset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bigquerydatasets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BigqueryDatasets that match those selectors.
func (c *bigqueryDatasets) List(opts v1.ListOptions) (result *v1alpha1.BigqueryDatasetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BigqueryDatasetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bigquerydatasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bigqueryDatasets.
func (c *bigqueryDatasets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bigquerydatasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a bigqueryDataset and creates it.  Returns the server's representation of the bigqueryDataset, and an error, if there is any.
func (c *bigqueryDatasets) Create(bigqueryDataset *v1alpha1.BigqueryDataset) (result *v1alpha1.BigqueryDataset, err error) {
	result = &v1alpha1.BigqueryDataset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bigquerydatasets").
		Body(bigqueryDataset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a bigqueryDataset and updates it. Returns the server's representation of the bigqueryDataset, and an error, if there is any.
func (c *bigqueryDatasets) Update(bigqueryDataset *v1alpha1.BigqueryDataset) (result *v1alpha1.BigqueryDataset, err error) {
	result = &v1alpha1.BigqueryDataset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bigquerydatasets").
		Name(bigqueryDataset.Name).
		Body(bigqueryDataset).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *bigqueryDatasets) UpdateStatus(bigqueryDataset *v1alpha1.BigqueryDataset) (result *v1alpha1.BigqueryDataset, err error) {
	result = &v1alpha1.BigqueryDataset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bigquerydatasets").
		Name(bigqueryDataset.Name).
		SubResource("status").
		Body(bigqueryDataset).
		Do().
		Into(result)
	return
}

// Delete takes name of the bigqueryDataset and deletes it. Returns an error if one occurs.
func (c *bigqueryDatasets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bigquerydatasets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bigqueryDatasets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bigquerydatasets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched bigqueryDataset.
func (c *bigqueryDatasets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BigqueryDataset, err error) {
	result = &v1alpha1.BigqueryDataset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bigquerydatasets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
