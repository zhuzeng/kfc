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

// MssqlElasticpoolsGetter has a method to return a MssqlElasticpoolInterface.
// A group's client should implement this interface.
type MssqlElasticpoolsGetter interface {
	MssqlElasticpools(namespace string) MssqlElasticpoolInterface
}

// MssqlElasticpoolInterface has methods to work with MssqlElasticpool resources.
type MssqlElasticpoolInterface interface {
	Create(*v1alpha1.MssqlElasticpool) (*v1alpha1.MssqlElasticpool, error)
	Update(*v1alpha1.MssqlElasticpool) (*v1alpha1.MssqlElasticpool, error)
	UpdateStatus(*v1alpha1.MssqlElasticpool) (*v1alpha1.MssqlElasticpool, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MssqlElasticpool, error)
	List(opts v1.ListOptions) (*v1alpha1.MssqlElasticpoolList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MssqlElasticpool, err error)
	MssqlElasticpoolExpansion
}

// mssqlElasticpools implements MssqlElasticpoolInterface
type mssqlElasticpools struct {
	client rest.Interface
	ns     string
}

// newMssqlElasticpools returns a MssqlElasticpools
func newMssqlElasticpools(c *AzurermV1alpha1Client, namespace string) *mssqlElasticpools {
	return &mssqlElasticpools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mssqlElasticpool, and returns the corresponding mssqlElasticpool object, and an error if there is any.
func (c *mssqlElasticpools) Get(name string, options v1.GetOptions) (result *v1alpha1.MssqlElasticpool, err error) {
	result = &v1alpha1.MssqlElasticpool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mssqlelasticpools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MssqlElasticpools that match those selectors.
func (c *mssqlElasticpools) List(opts v1.ListOptions) (result *v1alpha1.MssqlElasticpoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MssqlElasticpoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mssqlelasticpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mssqlElasticpools.
func (c *mssqlElasticpools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mssqlelasticpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mssqlElasticpool and creates it.  Returns the server's representation of the mssqlElasticpool, and an error, if there is any.
func (c *mssqlElasticpools) Create(mssqlElasticpool *v1alpha1.MssqlElasticpool) (result *v1alpha1.MssqlElasticpool, err error) {
	result = &v1alpha1.MssqlElasticpool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mssqlelasticpools").
		Body(mssqlElasticpool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mssqlElasticpool and updates it. Returns the server's representation of the mssqlElasticpool, and an error, if there is any.
func (c *mssqlElasticpools) Update(mssqlElasticpool *v1alpha1.MssqlElasticpool) (result *v1alpha1.MssqlElasticpool, err error) {
	result = &v1alpha1.MssqlElasticpool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mssqlelasticpools").
		Name(mssqlElasticpool.Name).
		Body(mssqlElasticpool).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mssqlElasticpools) UpdateStatus(mssqlElasticpool *v1alpha1.MssqlElasticpool) (result *v1alpha1.MssqlElasticpool, err error) {
	result = &v1alpha1.MssqlElasticpool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mssqlelasticpools").
		Name(mssqlElasticpool.Name).
		SubResource("status").
		Body(mssqlElasticpool).
		Do().
		Into(result)
	return
}

// Delete takes name of the mssqlElasticpool and deletes it. Returns an error if one occurs.
func (c *mssqlElasticpools) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mssqlelasticpools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mssqlElasticpools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mssqlelasticpools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mssqlElasticpool.
func (c *mssqlElasticpools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MssqlElasticpool, err error) {
	result = &v1alpha1.MssqlElasticpool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mssqlelasticpools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
