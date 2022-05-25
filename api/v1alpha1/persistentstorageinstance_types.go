/*
 * Copyright 2021, 2022 Hewlett Packard Enterprise Development LP
 * Other additional copyright holders may be indicated within.
 *
 * The entirety of this work is licensed under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 *
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// PersistentStorageNameLabel is defined for resources that relate to the name of a DWS PersistentStorageInstance
	PersistentStorageNameLabel = "dws.cray.hpe.com/persistentstorage.name"

	// PersistentStorageNamespaceLabel is defined for resources that relate to the namespace of a DWS PersistentStorageInstance
	PersistentStorageNamespaceLabel = "dws.cray.hpe.com/persistentstorage.namespace"
)

// PersistentStorageInstanceSpec defines the desired state of PersistentStorageInstance
type PersistentStorageInstanceSpec struct {
	// Name is the name given to this persistent storage instance.
	Name string `json:"name"`

	// FsType describes the File System Type for this storage instance.
	// +kubebuilder:validation:Enum=raw;xfs;gfs2;lustre
	FsType string `json:"fsType"`

	// DWDirective is a copy of the #DW for this instance
	DWDirective string `json:"dwDirective"`
}

// PersistentStorageInstanceStatus defines the observed state of PersistentStorageInstance
type PersistentStorageInstanceStatus struct {
	// Servers refers to the Servers resource that provides the backing storage for this storage instance
	Servers corev1.ObjectReference `json:"servers,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PersistentStorageInstance is the Schema for the Persistentstorageinstances API
type PersistentStorageInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PersistentStorageInstanceSpec   `json:"spec,omitempty"`
	Status PersistentStorageInstanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PersistentStorageInstanceList contains a list of PersistentStorageInstance
type PersistentStorageInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PersistentStorageInstance `json:"items"`
}

func (p *PersistentStorageInstanceList) GetObjectList() []client.Object {
	objectList := []client.Object{}

	for i := range p.Items {
		objectList = append(objectList, &p.Items[i])
	}

	return objectList
}

func init() {
	SchemeBuilder.Register(&PersistentStorageInstance{}, &PersistentStorageInstanceList{})
}
