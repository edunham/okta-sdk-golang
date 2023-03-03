/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// Code generated by okta openapi generator. DO NOT EDIT.

package okta

import "encoding/json"

type AuthenticatorProviderConfiguration struct {
	AuthPort         int64
	AuthPortPtr      *int64                                           `json:"authPort,omitempty"`
	Host             string                                           `json:"host,omitempty"`
	HostName         string                                           `json:"hostName,omitempty"`
	InstanceId       string                                           `json:"instanceId,omitempty"`
	IntegrationKey   string                                           `json:"integrationKey,omitempty"`
	SecretKey        string                                           `json:"secretKey,omitempty"`
	SharedSecret     string                                           `json:"sharedSecret,omitempty"`
	UserNameTemplate *AuthenticatorProviderConfigurationUserNamePlate `json:"userNameTemplate,omitempty"`
}

func (a *AuthenticatorProviderConfiguration) MarshalJSON() ([]byte, error) {
	type Alias AuthenticatorProviderConfiguration
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.AuthPort != 0 {
		result.AuthPortPtr = Int64Ptr(a.AuthPort)
	}
	return json.Marshal(&result)
}

func (a *AuthenticatorProviderConfiguration) UnmarshalJSON(data []byte) error {
	type Alias AuthenticatorProviderConfiguration

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.AuthPortPtr != nil {
		a.AuthPort = *result.AuthPortPtr
		a.AuthPortPtr = result.AuthPortPtr
	}
	return nil
}
