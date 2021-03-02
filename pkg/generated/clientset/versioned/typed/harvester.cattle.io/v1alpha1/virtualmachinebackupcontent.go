/*
Copyright 2021 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/rancher/harvester/pkg/apis/harvester.cattle.io/v1alpha1"
	scheme "github.com/rancher/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualMachineBackupContentsGetter has a method to return a VirtualMachineBackupContentInterface.
// A group's client should implement this interface.
type VirtualMachineBackupContentsGetter interface {
	VirtualMachineBackupContents(namespace string) VirtualMachineBackupContentInterface
}

// VirtualMachineBackupContentInterface has methods to work with VirtualMachineBackupContent resources.
type VirtualMachineBackupContentInterface interface {
	Create(ctx context.Context, virtualMachineBackupContent *v1alpha1.VirtualMachineBackupContent, opts v1.CreateOptions) (*v1alpha1.VirtualMachineBackupContent, error)
	Update(ctx context.Context, virtualMachineBackupContent *v1alpha1.VirtualMachineBackupContent, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineBackupContent, error)
	UpdateStatus(ctx context.Context, virtualMachineBackupContent *v1alpha1.VirtualMachineBackupContent, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineBackupContent, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VirtualMachineBackupContent, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VirtualMachineBackupContentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineBackupContent, err error)
	VirtualMachineBackupContentExpansion
}

// virtualMachineBackupContents implements VirtualMachineBackupContentInterface
type virtualMachineBackupContents struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineBackupContents returns a VirtualMachineBackupContents
func newVirtualMachineBackupContents(c *HarvesterV1alpha1Client, namespace string) *virtualMachineBackupContents {
	return &virtualMachineBackupContents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineBackupContent, and returns the corresponding virtualMachineBackupContent object, and an error if there is any.
func (c *virtualMachineBackupContents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineBackupContent, err error) {
	result = &v1alpha1.VirtualMachineBackupContent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinebackupcontents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineBackupContents that match those selectors.
func (c *virtualMachineBackupContents) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualMachineBackupContentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineBackupContentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinebackupcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineBackupContents.
func (c *virtualMachineBackupContents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinebackupcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualMachineBackupContent and creates it.  Returns the server's representation of the virtualMachineBackupContent, and an error, if there is any.
func (c *virtualMachineBackupContents) Create(ctx context.Context, virtualMachineBackupContent *v1alpha1.VirtualMachineBackupContent, opts v1.CreateOptions) (result *v1alpha1.VirtualMachineBackupContent, err error) {
	result = &v1alpha1.VirtualMachineBackupContent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinebackupcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineBackupContent).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualMachineBackupContent and updates it. Returns the server's representation of the virtualMachineBackupContent, and an error, if there is any.
func (c *virtualMachineBackupContents) Update(ctx context.Context, virtualMachineBackupContent *v1alpha1.VirtualMachineBackupContent, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineBackupContent, err error) {
	result = &v1alpha1.VirtualMachineBackupContent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinebackupcontents").
		Name(virtualMachineBackupContent.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineBackupContent).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *virtualMachineBackupContents) UpdateStatus(ctx context.Context, virtualMachineBackupContent *v1alpha1.VirtualMachineBackupContent, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineBackupContent, err error) {
	result = &v1alpha1.VirtualMachineBackupContent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinebackupcontents").
		Name(virtualMachineBackupContent.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineBackupContent).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualMachineBackupContent and deletes it. Returns an error if one occurs.
func (c *virtualMachineBackupContents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinebackupcontents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineBackupContents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinebackupcontents").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualMachineBackupContent.
func (c *virtualMachineBackupContents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineBackupContent, err error) {
	result = &v1alpha1.VirtualMachineBackupContent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinebackupcontents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
