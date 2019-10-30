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

// ElasticsearchDomainsGetter has a method to return a ElasticsearchDomainInterface.
// A group's client should implement this interface.
type ElasticsearchDomainsGetter interface {
	ElasticsearchDomains(namespace string) ElasticsearchDomainInterface
}

// ElasticsearchDomainInterface has methods to work with ElasticsearchDomain resources.
type ElasticsearchDomainInterface interface {
	Create(*v1alpha1.ElasticsearchDomain) (*v1alpha1.ElasticsearchDomain, error)
	Update(*v1alpha1.ElasticsearchDomain) (*v1alpha1.ElasticsearchDomain, error)
	UpdateStatus(*v1alpha1.ElasticsearchDomain) (*v1alpha1.ElasticsearchDomain, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ElasticsearchDomain, error)
	List(opts v1.ListOptions) (*v1alpha1.ElasticsearchDomainList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElasticsearchDomain, err error)
	ElasticsearchDomainExpansion
}

// elasticsearchDomains implements ElasticsearchDomainInterface
type elasticsearchDomains struct {
	client rest.Interface
	ns     string
}

// newElasticsearchDomains returns a ElasticsearchDomains
func newElasticsearchDomains(c *AwsV1alpha1Client, namespace string) *elasticsearchDomains {
	return &elasticsearchDomains{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the elasticsearchDomain, and returns the corresponding elasticsearchDomain object, and an error if there is any.
func (c *elasticsearchDomains) Get(name string, options v1.GetOptions) (result *v1alpha1.ElasticsearchDomain, err error) {
	result = &v1alpha1.ElasticsearchDomain{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchdomains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ElasticsearchDomains that match those selectors.
func (c *elasticsearchDomains) List(opts v1.ListOptions) (result *v1alpha1.ElasticsearchDomainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ElasticsearchDomainList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested elasticsearchDomains.
func (c *elasticsearchDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a elasticsearchDomain and creates it.  Returns the server's representation of the elasticsearchDomain, and an error, if there is any.
func (c *elasticsearchDomains) Create(elasticsearchDomain *v1alpha1.ElasticsearchDomain) (result *v1alpha1.ElasticsearchDomain, err error) {
	result = &v1alpha1.ElasticsearchDomain{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("elasticsearchdomains").
		Body(elasticsearchDomain).
		Do().
		Into(result)
	return
}

// Update takes the representation of a elasticsearchDomain and updates it. Returns the server's representation of the elasticsearchDomain, and an error, if there is any.
func (c *elasticsearchDomains) Update(elasticsearchDomain *v1alpha1.ElasticsearchDomain) (result *v1alpha1.ElasticsearchDomain, err error) {
	result = &v1alpha1.ElasticsearchDomain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticsearchdomains").
		Name(elasticsearchDomain.Name).
		Body(elasticsearchDomain).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *elasticsearchDomains) UpdateStatus(elasticsearchDomain *v1alpha1.ElasticsearchDomain) (result *v1alpha1.ElasticsearchDomain, err error) {
	result = &v1alpha1.ElasticsearchDomain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticsearchdomains").
		Name(elasticsearchDomain.Name).
		SubResource("status").
		Body(elasticsearchDomain).
		Do().
		Into(result)
	return
}

// Delete takes name of the elasticsearchDomain and deletes it. Returns an error if one occurs.
func (c *elasticsearchDomains) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticsearchdomains").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *elasticsearchDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticsearchdomains").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched elasticsearchDomain.
func (c *elasticsearchDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElasticsearchDomain, err error) {
	result = &v1alpha1.ElasticsearchDomain{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("elasticsearchdomains").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}