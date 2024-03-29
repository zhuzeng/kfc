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

// CloudwatchLogGroupsGetter has a method to return a CloudwatchLogGroupInterface.
// A group's client should implement this interface.
type CloudwatchLogGroupsGetter interface {
	CloudwatchLogGroups(namespace string) CloudwatchLogGroupInterface
}

// CloudwatchLogGroupInterface has methods to work with CloudwatchLogGroup resources.
type CloudwatchLogGroupInterface interface {
	Create(*v1alpha1.CloudwatchLogGroup) (*v1alpha1.CloudwatchLogGroup, error)
	Update(*v1alpha1.CloudwatchLogGroup) (*v1alpha1.CloudwatchLogGroup, error)
	UpdateStatus(*v1alpha1.CloudwatchLogGroup) (*v1alpha1.CloudwatchLogGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudwatchLogGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudwatchLogGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchLogGroup, err error)
	CloudwatchLogGroupExpansion
}

// cloudwatchLogGroups implements CloudwatchLogGroupInterface
type cloudwatchLogGroups struct {
	client rest.Interface
	ns     string
}

// newCloudwatchLogGroups returns a CloudwatchLogGroups
func newCloudwatchLogGroups(c *AwsV1alpha1Client, namespace string) *cloudwatchLogGroups {
	return &cloudwatchLogGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudwatchLogGroup, and returns the corresponding cloudwatchLogGroup object, and an error if there is any.
func (c *cloudwatchLogGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchLogGroup, err error) {
	result = &v1alpha1.CloudwatchLogGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatchloggroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudwatchLogGroups that match those selectors.
func (c *cloudwatchLogGroups) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchLogGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudwatchLogGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatchloggroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudwatchLogGroups.
func (c *cloudwatchLogGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatchloggroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudwatchLogGroup and creates it.  Returns the server's representation of the cloudwatchLogGroup, and an error, if there is any.
func (c *cloudwatchLogGroups) Create(cloudwatchLogGroup *v1alpha1.CloudwatchLogGroup) (result *v1alpha1.CloudwatchLogGroup, err error) {
	result = &v1alpha1.CloudwatchLogGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudwatchloggroups").
		Body(cloudwatchLogGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudwatchLogGroup and updates it. Returns the server's representation of the cloudwatchLogGroup, and an error, if there is any.
func (c *cloudwatchLogGroups) Update(cloudwatchLogGroup *v1alpha1.CloudwatchLogGroup) (result *v1alpha1.CloudwatchLogGroup, err error) {
	result = &v1alpha1.CloudwatchLogGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudwatchloggroups").
		Name(cloudwatchLogGroup.Name).
		Body(cloudwatchLogGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudwatchLogGroups) UpdateStatus(cloudwatchLogGroup *v1alpha1.CloudwatchLogGroup) (result *v1alpha1.CloudwatchLogGroup, err error) {
	result = &v1alpha1.CloudwatchLogGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudwatchloggroups").
		Name(cloudwatchLogGroup.Name).
		SubResource("status").
		Body(cloudwatchLogGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudwatchLogGroup and deletes it. Returns an error if one occurs.
func (c *cloudwatchLogGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudwatchloggroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudwatchLogGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudwatchloggroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudwatchLogGroup.
func (c *cloudwatchLogGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchLogGroup, err error) {
	result = &v1alpha1.CloudwatchLogGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudwatchloggroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
