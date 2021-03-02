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

package fake

import (
	"context"

	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReplicas implements ReplicaInterface
type FakeReplicas struct {
	Fake *FakeLonghornV1beta1
	ns   string
}

var replicasResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "replicas"}

var replicasKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Replica"}

// Get takes name of the replica, and returns the corresponding replica object, and an error if there is any.
func (c *FakeReplicas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Replica, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(replicasResource, c.ns, name), &v1beta1.Replica{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Replica), err
}

// List takes label and field selectors, and returns the list of Replicas that match those selectors.
func (c *FakeReplicas) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ReplicaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(replicasResource, replicasKind, c.ns, opts), &v1beta1.ReplicaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ReplicaList{ListMeta: obj.(*v1beta1.ReplicaList).ListMeta}
	for _, item := range obj.(*v1beta1.ReplicaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicas.
func (c *FakeReplicas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(replicasResource, c.ns, opts))

}

// Create takes the representation of a replica and creates it.  Returns the server's representation of the replica, and an error, if there is any.
func (c *FakeReplicas) Create(ctx context.Context, replica *v1beta1.Replica, opts v1.CreateOptions) (result *v1beta1.Replica, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(replicasResource, c.ns, replica), &v1beta1.Replica{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Replica), err
}

// Update takes the representation of a replica and updates it. Returns the server's representation of the replica, and an error, if there is any.
func (c *FakeReplicas) Update(ctx context.Context, replica *v1beta1.Replica, opts v1.UpdateOptions) (result *v1beta1.Replica, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(replicasResource, c.ns, replica), &v1beta1.Replica{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Replica), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicas) UpdateStatus(ctx context.Context, replica *v1beta1.Replica, opts v1.UpdateOptions) (*v1beta1.Replica, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(replicasResource, "status", c.ns, replica), &v1beta1.Replica{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Replica), err
}

// Delete takes name of the replica and deletes it. Returns an error if one occurs.
func (c *FakeReplicas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(replicasResource, c.ns, name), &v1beta1.Replica{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(replicasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ReplicaList{})
	return err
}

// Patch applies the patch and returns the patched replica.
func (c *FakeReplicas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Replica, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(replicasResource, c.ns, name, pt, data, subresources...), &v1beta1.Replica{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Replica), err
}
