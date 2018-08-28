/*
Copyright 2018 the Heptio Ark contributors.

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

package fake

import (
	ark_v1 "github.com/heptio/ark/pkg/apis/ark/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackupStorageLocations implements BackupStorageLocationInterface
type FakeBackupStorageLocations struct {
	Fake *FakeArkV1
	ns   string
}

var backupstoragelocationsResource = schema.GroupVersionResource{Group: "ark.heptio.com", Version: "v1", Resource: "backupstoragelocations"}

var backupstoragelocationsKind = schema.GroupVersionKind{Group: "ark.heptio.com", Version: "v1", Kind: "BackupStorageLocation"}

// Get takes name of the backupStorageLocation, and returns the corresponding backupStorageLocation object, and an error if there is any.
func (c *FakeBackupStorageLocations) Get(name string, options v1.GetOptions) (result *ark_v1.BackupStorageLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupstoragelocationsResource, c.ns, name), &ark_v1.BackupStorageLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.BackupStorageLocation), err
}

// List takes label and field selectors, and returns the list of BackupStorageLocations that match those selectors.
func (c *FakeBackupStorageLocations) List(opts v1.ListOptions) (result *ark_v1.BackupStorageLocationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupstoragelocationsResource, backupstoragelocationsKind, c.ns, opts), &ark_v1.BackupStorageLocationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &ark_v1.BackupStorageLocationList{ListMeta: obj.(*ark_v1.BackupStorageLocationList).ListMeta}
	for _, item := range obj.(*ark_v1.BackupStorageLocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupStorageLocations.
func (c *FakeBackupStorageLocations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupstoragelocationsResource, c.ns, opts))

}

// Create takes the representation of a backupStorageLocation and creates it.  Returns the server's representation of the backupStorageLocation, and an error, if there is any.
func (c *FakeBackupStorageLocations) Create(backupStorageLocation *ark_v1.BackupStorageLocation) (result *ark_v1.BackupStorageLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupstoragelocationsResource, c.ns, backupStorageLocation), &ark_v1.BackupStorageLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.BackupStorageLocation), err
}

// Update takes the representation of a backupStorageLocation and updates it. Returns the server's representation of the backupStorageLocation, and an error, if there is any.
func (c *FakeBackupStorageLocations) Update(backupStorageLocation *ark_v1.BackupStorageLocation) (result *ark_v1.BackupStorageLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupstoragelocationsResource, c.ns, backupStorageLocation), &ark_v1.BackupStorageLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.BackupStorageLocation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupStorageLocations) UpdateStatus(backupStorageLocation *ark_v1.BackupStorageLocation) (*ark_v1.BackupStorageLocation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupstoragelocationsResource, "status", c.ns, backupStorageLocation), &ark_v1.BackupStorageLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.BackupStorageLocation), err
}

// Delete takes name of the backupStorageLocation and deletes it. Returns an error if one occurs.
func (c *FakeBackupStorageLocations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupstoragelocationsResource, c.ns, name), &ark_v1.BackupStorageLocation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupStorageLocations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backupstoragelocationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &ark_v1.BackupStorageLocationList{})
	return err
}

// Patch applies the patch and returns the patched backupStorageLocation.
func (c *FakeBackupStorageLocations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *ark_v1.BackupStorageLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupstoragelocationsResource, c.ns, name, data, subresources...), &ark_v1.BackupStorageLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.BackupStorageLocation), err
}
