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

// AlbListenerRulesGetter has a method to return a AlbListenerRuleInterface.
// A group's client should implement this interface.
type AlbListenerRulesGetter interface {
	AlbListenerRules(namespace string) AlbListenerRuleInterface
}

// AlbListenerRuleInterface has methods to work with AlbListenerRule resources.
type AlbListenerRuleInterface interface {
	Create(*v1alpha1.AlbListenerRule) (*v1alpha1.AlbListenerRule, error)
	Update(*v1alpha1.AlbListenerRule) (*v1alpha1.AlbListenerRule, error)
	UpdateStatus(*v1alpha1.AlbListenerRule) (*v1alpha1.AlbListenerRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AlbListenerRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AlbListenerRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AlbListenerRule, err error)
	AlbListenerRuleExpansion
}

// albListenerRules implements AlbListenerRuleInterface
type albListenerRules struct {
	client rest.Interface
	ns     string
}

// newAlbListenerRules returns a AlbListenerRules
func newAlbListenerRules(c *AwsV1alpha1Client, namespace string) *albListenerRules {
	return &albListenerRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the albListenerRule, and returns the corresponding albListenerRule object, and an error if there is any.
func (c *albListenerRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AlbListenerRule, err error) {
	result = &v1alpha1.AlbListenerRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alblistenerrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AlbListenerRules that match those selectors.
func (c *albListenerRules) List(opts v1.ListOptions) (result *v1alpha1.AlbListenerRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AlbListenerRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alblistenerrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested albListenerRules.
func (c *albListenerRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("alblistenerrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a albListenerRule and creates it.  Returns the server's representation of the albListenerRule, and an error, if there is any.
func (c *albListenerRules) Create(albListenerRule *v1alpha1.AlbListenerRule) (result *v1alpha1.AlbListenerRule, err error) {
	result = &v1alpha1.AlbListenerRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("alblistenerrules").
		Body(albListenerRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a albListenerRule and updates it. Returns the server's representation of the albListenerRule, and an error, if there is any.
func (c *albListenerRules) Update(albListenerRule *v1alpha1.AlbListenerRule) (result *v1alpha1.AlbListenerRule, err error) {
	result = &v1alpha1.AlbListenerRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alblistenerrules").
		Name(albListenerRule.Name).
		Body(albListenerRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *albListenerRules) UpdateStatus(albListenerRule *v1alpha1.AlbListenerRule) (result *v1alpha1.AlbListenerRule, err error) {
	result = &v1alpha1.AlbListenerRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alblistenerrules").
		Name(albListenerRule.Name).
		SubResource("status").
		Body(albListenerRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the albListenerRule and deletes it. Returns an error if one occurs.
func (c *albListenerRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alblistenerrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *albListenerRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alblistenerrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched albListenerRule.
func (c *albListenerRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AlbListenerRule, err error) {
	result = &v1alpha1.AlbListenerRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("alblistenerrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
