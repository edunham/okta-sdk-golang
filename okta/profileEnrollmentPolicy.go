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

import (
	"encoding/json"
	"time"
)

type ProfileEnrollmentPolicy struct {
	Embedded    interface{}           `json:"_embedded,omitempty"`
	Links       interface{}           `json:"_links,omitempty"`
	Conditions  *PolicyRuleConditions `json:"conditions,omitempty"`
	Created     *time.Time            `json:"created,omitempty"`
	Description string                `json:"description,omitempty"`
	Id          string                `json:"id,omitempty"`
	LastUpdated *time.Time            `json:"lastUpdated,omitempty"`
	Name        string                `json:"name,omitempty"`
	Priority    int64
	PriorityPtr *int64 `json:"priority,omitempty"`
	Status      string `json:"status,omitempty"`
	System      *bool  `json:"system,omitempty"`
	Type        string `json:"type,omitempty"`
}

func NewProfileEnrollmentPolicy() *ProfileEnrollmentPolicy {
	return &ProfileEnrollmentPolicy{
		Type: "PROFILE_ENROLLMENT",
	}
}

func (a *ProfileEnrollmentPolicy) IsPolicyInstance() bool {
	return true
}

func (a *ProfileEnrollmentPolicy) MarshalJSON() ([]byte, error) {
	type Alias ProfileEnrollmentPolicy
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.Priority != 0 {
		result.PriorityPtr = Int64Ptr(a.Priority)
	}
	return json.Marshal(&result)
}

func (a *ProfileEnrollmentPolicy) UnmarshalJSON(data []byte) error {
	type Alias ProfileEnrollmentPolicy

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.PriorityPtr != nil {
		a.Priority = *result.PriorityPtr
		a.PriorityPtr = result.PriorityPtr
	}
	return nil
}
