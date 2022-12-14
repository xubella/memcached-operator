/*
Copyright 2022.

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
	tektonv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MemcachedSpec defines the desired state of Memcached
type MemcachedSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Memcached. Edit memcached_types.go to remove/update
	Size int32 `json:"size"`

	URL string `json:"url"`
	// +optional
	Repository string `json:"repository"`
	// +optional
	Branch string `json:"branch"`
	// +optional
	Reversion string `json:"reversion"`
	// +optional
	BranchTrim string `json:"branch-trim"`

	// +listType=atomic
	// +optional
	Workspaces []tektonv1beta1.WorkspaceBinding `json:"workspaces,omitempty"`
}

// MemcachedStatus defines the observed state of Memcached
type MemcachedStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Nodes []string `json:"nodes"`

	duckv1beta1.Status `json:",inline"`

	// RunStatusFields inlines the status fields.
	MemcachedStatusFields `json:",inline"`
}

type MemcachedStatusFields struct {
	// StartTime is the time the build is actually started.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// CompletionTime is the time the build completed.
	// +optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`

	// Results reports any output result values to be consumed by later
	// tasks in a pipeline.
	// +optional
	//Results []RunResult `json:"results,omitempty"`

	// RetriesStatus contains the history of RunStatus, in case of a retry.
	// +optional
	//RetriesStatus []RunStatus `json:"retriesStatus,omitempty"`

	// ExtraFields holds arbitrary fields provided by the custom task
	// controller.
	//ExtraFields runtime.RawExtension `json:"extraFields,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Memcached is the Schema for the memcacheds API
type Memcached struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcachedSpec   `json:"spec,omitempty"`
	Status MemcachedStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MemcachedList contains a list of Memcached
type MemcachedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memcached `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Memcached{}, &MemcachedList{})
}

func (m *Memcached) GetStatusCondition() apis.ConditionAccessor {
	return &m.Status
}

var memCondSet = apis.NewBatchConditionSet()

// MarkRunSucceeded changes the Succeeded condition to True with the provided reason and message.
func (m *MemcachedStatus) MarkRunSucceeded(reason, messageFormat string, messageA ...interface{}) {
	memCondSet.Manage(m).MarkTrueWithReason(apis.ConditionSucceeded, reason, messageFormat, messageA...)
	succeeded := m.GetCondition(apis.ConditionSucceeded)
	m.CompletionTime = &succeeded.LastTransitionTime.Inner
}

// MarkRunFailed changes the Succeeded condition to False with the provided reason and message.
func (m *MemcachedStatus) MarkRunFailed(reason, messageFormat string, messageA ...interface{}) {
	memCondSet.Manage(m).MarkFalse(apis.ConditionSucceeded, reason, messageFormat, messageA...)
	succeeded := m.GetCondition(apis.ConditionSucceeded)
	m.CompletionTime = &succeeded.LastTransitionTime.Inner
}
