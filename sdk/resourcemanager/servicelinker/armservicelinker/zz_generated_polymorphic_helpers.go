//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicelinker

import "encoding/json"

func unmarshalAuthInfoBaseClassification(rawMsg json.RawMessage) (AuthInfoBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AuthInfoBaseClassification
	switch m["authType"] {
	case string(AuthTypeSecret):
		b = &SecretAuthInfo{}
	case string(AuthTypeServicePrincipalCertificate):
		b = &ServicePrincipalCertificateAuthInfo{}
	case string(AuthTypeServicePrincipalSecret):
		b = &ServicePrincipalSecretAuthInfo{}
	case string(AuthTypeSystemAssignedIdentity):
		b = &SystemAssignedIdentityAuthInfo{}
	case string(AuthTypeUserAssignedIdentity):
		b = &UserAssignedIdentityAuthInfo{}
	default:
		b = &AuthInfoBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAzureResourcePropertiesBaseClassification(rawMsg json.RawMessage) (AzureResourcePropertiesBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AzureResourcePropertiesBaseClassification
	switch m["type"] {
	case string(AzureResourceTypeKeyVault):
		b = &AzureKeyVaultProperties{}
	default:
		b = &AzureResourcePropertiesBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSecretInfoBaseClassification(rawMsg json.RawMessage) (SecretInfoBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SecretInfoBaseClassification
	switch m["secretType"] {
	case string(SecretTypeKeyVaultSecretReference):
		b = &KeyVaultSecretReferenceSecretInfo{}
	case string(SecretTypeKeyVaultSecretURI):
		b = &KeyVaultSecretURISecretInfo{}
	case string(SecretTypeRawValue):
		b = &ValueSecretInfo{}
	default:
		b = &SecretInfoBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTargetServiceBaseClassification(rawMsg json.RawMessage) (TargetServiceBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TargetServiceBaseClassification
	switch m["type"] {
	case string(TargetServiceTypeAzureResource):
		b = &AzureResource{}
	case string(TargetServiceTypeConfluentBootstrapServer):
		b = &ConfluentBootstrapServer{}
	case string(TargetServiceTypeConfluentSchemaRegistry):
		b = &ConfluentSchemaRegistry{}
	default:
		b = &TargetServiceBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
