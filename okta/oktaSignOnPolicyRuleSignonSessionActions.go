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

type OktaSignOnPolicyRuleSignonSessionActions struct {
	MaxSessionIdleMinutes        int64
	MaxSessionIdleMinutesPtr     *int64 `json:"maxSessionIdleMinutes"`
	MaxSessionLifetimeMinutes    int64
	MaxSessionLifetimeMinutesPtr *int64 `json:"maxSessionLifetimeMinutes"`
	UsePersistentCookie          *bool  `json:"usePersistentCookie,omitempty"`
}

func NewOktaSignOnPolicyRuleSignonSessionActions() *OktaSignOnPolicyRuleSignonSessionActions {
	return &OktaSignOnPolicyRuleSignonSessionActions{
		MaxSessionIdleMinutes:        120,
		MaxSessionIdleMinutesPtr:     Int64Ptr(120),
		MaxSessionLifetimeMinutes:    0,
		MaxSessionLifetimeMinutesPtr: Int64Ptr(0),
		UsePersistentCookie:          boolPtr(false),
	}
}

func (a *OktaSignOnPolicyRuleSignonSessionActions) IsPolicyInstance() bool {
	return true
}

func (a *OktaSignOnPolicyRuleSignonSessionActions) MarshalJSON() ([]byte, error) {
	type Alias OktaSignOnPolicyRuleSignonSessionActions
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.MaxSessionIdleMinutes != 0 {
		result.MaxSessionIdleMinutesPtr = Int64Ptr(a.MaxSessionIdleMinutes)
	}
	if a.MaxSessionLifetimeMinutes != 0 {
		result.MaxSessionLifetimeMinutesPtr = Int64Ptr(a.MaxSessionLifetimeMinutes)
	}
	return json.Marshal(&result)
}

func (a *OktaSignOnPolicyRuleSignonSessionActions) UnmarshalJSON(data []byte) error {
	type Alias OktaSignOnPolicyRuleSignonSessionActions

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.MaxSessionIdleMinutesPtr != nil {
		a.MaxSessionIdleMinutes = *result.MaxSessionIdleMinutesPtr
		a.MaxSessionIdleMinutesPtr = result.MaxSessionIdleMinutesPtr
	}
	if result.MaxSessionLifetimeMinutesPtr != nil {
		a.MaxSessionLifetimeMinutes = *result.MaxSessionLifetimeMinutesPtr
		a.MaxSessionLifetimeMinutesPtr = result.MaxSessionLifetimeMinutesPtr
	}
	return nil
}
