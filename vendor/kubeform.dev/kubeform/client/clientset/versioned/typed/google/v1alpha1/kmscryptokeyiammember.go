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

// KmsCryptoKeyIamMembersGetter has a method to return a KmsCryptoKeyIamMemberInterface.
// A group's client should implement this interface.
type KmsCryptoKeyIamMembersGetter interface {
	KmsCryptoKeyIamMembers(namespace string) KmsCryptoKeyIamMemberInterface
}

// KmsCryptoKeyIamMemberInterface has methods to work with KmsCryptoKeyIamMember resources.
type KmsCryptoKeyIamMemberInterface interface {
	Create(*v1alpha1.KmsCryptoKeyIamMember) (*v1alpha1.KmsCryptoKeyIamMember, error)
	Update(*v1alpha1.KmsCryptoKeyIamMember) (*v1alpha1.KmsCryptoKeyIamMember, error)
	UpdateStatus(*v1alpha1.KmsCryptoKeyIamMember) (*v1alpha1.KmsCryptoKeyIamMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KmsCryptoKeyIamMember, error)
	List(opts v1.ListOptions) (*v1alpha1.KmsCryptoKeyIamMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsCryptoKeyIamMember, err error)
	KmsCryptoKeyIamMemberExpansion
}

// kmsCryptoKeyIamMembers implements KmsCryptoKeyIamMemberInterface
type kmsCryptoKeyIamMembers struct {
	client rest.Interface
	ns     string
}

// newKmsCryptoKeyIamMembers returns a KmsCryptoKeyIamMembers
func newKmsCryptoKeyIamMembers(c *GoogleV1alpha1Client, namespace string) *kmsCryptoKeyIamMembers {
	return &kmsCryptoKeyIamMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kmsCryptoKeyIamMember, and returns the corresponding kmsCryptoKeyIamMember object, and an error if there is any.
func (c *kmsCryptoKeyIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsCryptoKeyIamMember, err error) {
	result = &v1alpha1.KmsCryptoKeyIamMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmscryptokeyiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KmsCryptoKeyIamMembers that match those selectors.
func (c *kmsCryptoKeyIamMembers) List(opts v1.ListOptions) (result *v1alpha1.KmsCryptoKeyIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KmsCryptoKeyIamMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmscryptokeyiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kmsCryptoKeyIamMembers.
func (c *kmsCryptoKeyIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kmscryptokeyiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kmsCryptoKeyIamMember and creates it.  Returns the server's representation of the kmsCryptoKeyIamMember, and an error, if there is any.
func (c *kmsCryptoKeyIamMembers) Create(kmsCryptoKeyIamMember *v1alpha1.KmsCryptoKeyIamMember) (result *v1alpha1.KmsCryptoKeyIamMember, err error) {
	result = &v1alpha1.KmsCryptoKeyIamMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kmscryptokeyiammembers").
		Body(kmsCryptoKeyIamMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kmsCryptoKeyIamMember and updates it. Returns the server's representation of the kmsCryptoKeyIamMember, and an error, if there is any.
func (c *kmsCryptoKeyIamMembers) Update(kmsCryptoKeyIamMember *v1alpha1.KmsCryptoKeyIamMember) (result *v1alpha1.KmsCryptoKeyIamMember, err error) {
	result = &v1alpha1.KmsCryptoKeyIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmscryptokeyiammembers").
		Name(kmsCryptoKeyIamMember.Name).
		Body(kmsCryptoKeyIamMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kmsCryptoKeyIamMembers) UpdateStatus(kmsCryptoKeyIamMember *v1alpha1.KmsCryptoKeyIamMember) (result *v1alpha1.KmsCryptoKeyIamMember, err error) {
	result = &v1alpha1.KmsCryptoKeyIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmscryptokeyiammembers").
		Name(kmsCryptoKeyIamMember.Name).
		SubResource("status").
		Body(kmsCryptoKeyIamMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the kmsCryptoKeyIamMember and deletes it. Returns an error if one occurs.
func (c *kmsCryptoKeyIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmscryptokeyiammembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kmsCryptoKeyIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmscryptokeyiammembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kmsCryptoKeyIamMember.
func (c *kmsCryptoKeyIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsCryptoKeyIamMember, err error) {
	result = &v1alpha1.KmsCryptoKeyIamMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kmscryptokeyiammembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}