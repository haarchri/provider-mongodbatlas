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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this PartyIntegration
func (mg *PartyIntegration) GetTerraformResourceType() string {
	return "mongodbatlas_third_party_integration"
}

// GetConnectionDetailsMapping for this PartyIntegration
func (tr *PartyIntegration) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"api_key": "spec.forProvider.apiKeySecretRef", "api_token": "spec.forProvider.apiTokenSecretRef", "license_key": "spec.forProvider.licenseKeySecretRef", "microsoft_teams_webhook_url": "spec.forProvider.microsoftTeamsWebhookUrlSecretRef", "password": "spec.forProvider.passwordSecretRef", "read_token": "spec.forProvider.readTokenSecretRef", "routing_key": "spec.forProvider.routingKeySecretRef", "secret": "spec.forProvider.secretSecretRef", "service_discovery": "spec.forProvider.serviceDiscoverySecretRef", "service_key": "spec.forProvider.serviceKeySecretRef", "user_name": "spec.forProvider.userNameSecretRef", "write_token": "spec.forProvider.writeTokenSecretRef"}
}

// GetObservation of this PartyIntegration
func (tr *PartyIntegration) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this PartyIntegration
func (tr *PartyIntegration) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this PartyIntegration
func (tr *PartyIntegration) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this PartyIntegration
func (tr *PartyIntegration) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this PartyIntegration
func (tr *PartyIntegration) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this PartyIntegration using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *PartyIntegration) LateInitialize(attrs []byte) (bool, error) {
	params := &PartyIntegrationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *PartyIntegration) GetTerraformSchemaVersion() int {
	return 0
}