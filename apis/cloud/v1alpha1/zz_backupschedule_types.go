/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BackupScheduleObservation struct {
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IDPolicy *string `json:"idPolicy,omitempty" tf:"id_policy,omitempty"`

	NextSnapshot *string `json:"nextSnapshot,omitempty" tf:"next_snapshot,omitempty"`
}

type BackupScheduleParameters struct {

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyItemDaily []PolicyItemDailyParameters `json:"policyItemDaily,omitempty" tf:"policy_item_daily,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyItemHourly []PolicyItemHourlyParameters `json:"policyItemHourly,omitempty" tf:"policy_item_hourly,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyItemMonthly []PolicyItemMonthlyParameters `json:"policyItemMonthly,omitempty" tf:"policy_item_monthly,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyItemWeekly []PolicyItemWeeklyParameters `json:"policyItemWeekly,omitempty" tf:"policy_item_weekly,omitempty"`

	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ReferenceHourOfDay *int64 `json:"referenceHourOfDay,omitempty" tf:"reference_hour_of_day,omitempty"`

	// +kubebuilder:validation:Optional
	ReferenceMinuteOfHour *int64 `json:"referenceMinuteOfHour,omitempty" tf:"reference_minute_of_hour,omitempty"`

	// +kubebuilder:validation:Optional
	RestoreWindowDays *int64 `json:"restoreWindowDays,omitempty" tf:"restore_window_days,omitempty"`

	// +kubebuilder:validation:Optional
	UpdateSnapshots *bool `json:"updateSnapshots,omitempty" tf:"update_snapshots,omitempty"`
}

type PolicyItemDailyObservation struct {
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyItemDailyParameters struct {

	// +kubebuilder:validation:Required
	FrequencyInterval *int64 `json:"frequencyInterval" tf:"frequency_interval,omitempty"`

	// +kubebuilder:validation:Required
	RetentionUnit *string `json:"retentionUnit" tf:"retention_unit,omitempty"`

	// +kubebuilder:validation:Required
	RetentionValue *int64 `json:"retentionValue" tf:"retention_value,omitempty"`
}

type PolicyItemHourlyObservation struct {
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyItemHourlyParameters struct {

	// +kubebuilder:validation:Required
	FrequencyInterval *int64 `json:"frequencyInterval" tf:"frequency_interval,omitempty"`

	// +kubebuilder:validation:Required
	RetentionUnit *string `json:"retentionUnit" tf:"retention_unit,omitempty"`

	// +kubebuilder:validation:Required
	RetentionValue *int64 `json:"retentionValue" tf:"retention_value,omitempty"`
}

type PolicyItemMonthlyObservation struct {
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyItemMonthlyParameters struct {

	// +kubebuilder:validation:Required
	FrequencyInterval *int64 `json:"frequencyInterval" tf:"frequency_interval,omitempty"`

	// +kubebuilder:validation:Required
	RetentionUnit *string `json:"retentionUnit" tf:"retention_unit,omitempty"`

	// +kubebuilder:validation:Required
	RetentionValue *int64 `json:"retentionValue" tf:"retention_value,omitempty"`
}

type PolicyItemWeeklyObservation struct {
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyItemWeeklyParameters struct {

	// +kubebuilder:validation:Required
	FrequencyInterval *int64 `json:"frequencyInterval" tf:"frequency_interval,omitempty"`

	// +kubebuilder:validation:Required
	RetentionUnit *string `json:"retentionUnit" tf:"retention_unit,omitempty"`

	// +kubebuilder:validation:Required
	RetentionValue *int64 `json:"retentionValue" tf:"retention_value,omitempty"`
}

// BackupScheduleSpec defines the desired state of BackupSchedule
type BackupScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupScheduleParameters `json:"forProvider"`
}

// BackupScheduleStatus defines the observed state of BackupSchedule.
type BackupScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupSchedule is the Schema for the BackupSchedules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlasjet}
type BackupSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupScheduleSpec   `json:"spec"`
	Status            BackupScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupScheduleList contains a list of BackupSchedules
type BackupScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupSchedule `json:"items"`
}

// Repository type metadata.
var (
	BackupSchedule_Kind             = "BackupSchedule"
	BackupSchedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupSchedule_Kind}.String()
	BackupSchedule_KindAPIVersion   = BackupSchedule_Kind + "." + CRDGroupVersion.String()
	BackupSchedule_GroupVersionKind = CRDGroupVersion.WithKind(BackupSchedule_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupSchedule{}, &BackupScheduleList{})
}