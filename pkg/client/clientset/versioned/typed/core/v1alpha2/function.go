/*
Copyright 2022 The OpenFunction Authors.

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

package v1alpha2

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha2 "github.com/openfunction/apis/core/v1alpha2"
	scheme "github.com/openfunction/pkg/client/clientset/versioned/scheme"
)

// FunctionsGetter has a method to return a FunctionInterface.
// A group's client should implement this interface.
type FunctionsGetter interface {
	Functions(namespace string) FunctionInterface
}

// FunctionInterface has methods to work with Function resources.
type FunctionInterface interface {
	Create(ctx context.Context, function *v1alpha2.Function, opts v1.CreateOptions) (*v1alpha2.Function, error)
	Update(ctx context.Context, function *v1alpha2.Function, opts v1.UpdateOptions) (*v1alpha2.Function, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.Function, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.FunctionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Function, err error)
	FunctionExpansion
}

// functions implements FunctionInterface
type functions struct {
	client rest.Interface
	ns     string
}

// newFunctions returns a Functions
func newFunctions(c *CoreV1alpha2Client, namespace string) *functions {
	return &functions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the function, and returns the corresponding function object, and an error if there is any.
func (c *functions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Function, err error) {
	result = &v1alpha2.Function{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("functions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Functions that match those selectors.
func (c *functions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.FunctionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.FunctionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("functions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested functions.
func (c *functions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("functions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a function and creates it.  Returns the server's representation of the function, and an error, if there is any.
func (c *functions) Create(ctx context.Context, function *v1alpha2.Function, opts v1.CreateOptions) (result *v1alpha2.Function, err error) {
	result = &v1alpha2.Function{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("functions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(function).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a function and updates it. Returns the server's representation of the function, and an error, if there is any.
func (c *functions) Update(ctx context.Context, function *v1alpha2.Function, opts v1.UpdateOptions) (result *v1alpha2.Function, err error) {
	result = &v1alpha2.Function{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("functions").
		Name(function.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(function).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the function and deletes it. Returns an error if one occurs.
func (c *functions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("functions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *functions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("functions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched function.
func (c *functions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Function, err error) {
	result = &v1alpha2.Function{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("functions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
