/*


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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TheiaSpec defines the desired state of Theia
type TheiaSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Template TheiaTemplateSpec `json:"template,omitempty"`
}

// TheiaTemplateSpec defines the pod spec for the Theia
type TheiaTemplateSpec struct {
	Spec corev1.PodSpec `json:"spec,omitempty"`
}

// TheiaStatus defines the observed state of Theia
type TheiaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Conditions is an array of current conditions
	Conditions []TheiaCondition `json:"conditions"`
	// ReadyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
	ReadyReplicas int32 `json:"readyReplicas"`
	// ContainerState is the state of underlying container.
	ContainerState corev1.ContainerState `json:"containerState"`
}

// TheiaCondition defines the conditions of Theia status
type TheiaCondition struct {
	// Type is the type of the condition. Possible values are Running|Waiting|Terminated
	Type string `json:"type"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// (brief) reason the container is in the current state
	// +optional
	Reason string `json:"reason,omitempty"`
	// Message regarding why the container is in the current state.
	// +optional
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true

// Theia is the Schema for the theia API
type Theia struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TheiaSpec   `json:"spec,omitempty"`
	Status TheiaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TheiaList contains a list of Theia
type TheiaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Theia `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Theia{}, &TheiaList{})
}
