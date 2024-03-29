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

// KmsGrantsGetter has a method to return a KmsGrantInterface.
// A group's client should implement this interface.
type KmsGrantsGetter interface {
	KmsGrants(namespace string) KmsGrantInterface
}

// KmsGrantInterface has methods to work with KmsGrant resources.
type KmsGrantInterface interface {
	Create(*v1alpha1.KmsGrant) (*v1alpha1.KmsGrant, error)
	Update(*v1alpha1.KmsGrant) (*v1alpha1.KmsGrant, error)
	UpdateStatus(*v1alpha1.KmsGrant) (*v1alpha1.KmsGrant, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KmsGrant, error)
	List(opts v1.ListOptions) (*v1alpha1.KmsGrantList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsGrant, err error)
	KmsGrantExpansion
}

// kmsGrants implements KmsGrantInterface
type kmsGrants struct {
	client rest.Interface
	ns     string
}

// newKmsGrants returns a KmsGrants
func newKmsGrants(c *AwsV1alpha1Client, namespace string) *kmsGrants {
	return &kmsGrants{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kmsGrant, and returns the corresponding kmsGrant object, and an error if there is any.
func (c *kmsGrants) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsGrant, err error) {
	result = &v1alpha1.KmsGrant{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmsgrants").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KmsGrants that match those selectors.
func (c *kmsGrants) List(opts v1.ListOptions) (result *v1alpha1.KmsGrantList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KmsGrantList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmsgrants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kmsGrants.
func (c *kmsGrants) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kmsgrants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kmsGrant and creates it.  Returns the server's representation of the kmsGrant, and an error, if there is any.
func (c *kmsGrants) Create(kmsGrant *v1alpha1.KmsGrant) (result *v1alpha1.KmsGrant, err error) {
	result = &v1alpha1.KmsGrant{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kmsgrants").
		Body(kmsGrant).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kmsGrant and updates it. Returns the server's representation of the kmsGrant, and an error, if there is any.
func (c *kmsGrants) Update(kmsGrant *v1alpha1.KmsGrant) (result *v1alpha1.KmsGrant, err error) {
	result = &v1alpha1.KmsGrant{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmsgrants").
		Name(kmsGrant.Name).
		Body(kmsGrant).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kmsGrants) UpdateStatus(kmsGrant *v1alpha1.KmsGrant) (result *v1alpha1.KmsGrant, err error) {
	result = &v1alpha1.KmsGrant{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmsgrants").
		Name(kmsGrant.Name).
		SubResource("status").
		Body(kmsGrant).
		Do().
		Into(result)
	return
}

// Delete takes name of the kmsGrant and deletes it. Returns an error if one occurs.
func (c *kmsGrants) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmsgrants").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kmsGrants) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmsgrants").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kmsGrant.
func (c *kmsGrants) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsGrant, err error) {
	result = &v1alpha1.KmsGrant{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kmsgrants").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
