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

type PasswordPolicyPasswordSettingsLockout struct {
	AutoUnlockMinutes               int64
	AutoUnlockMinutesPtr            *int64 `json:"autoUnlockMinutes,omitempty"`
	MaxAttempts                     int64
	MaxAttemptsPtr                  *int64   `json:"maxAttempts,omitempty"`
	ShowLockoutFailures             *bool    `json:"showLockoutFailures,omitempty"`
	UserLockoutNotificationChannels []string `json:"userLockoutNotificationChannels,omitempty"`
}

func NewPasswordPolicyPasswordSettingsLockout() *PasswordPolicyPasswordSettingsLockout {
	return &PasswordPolicyPasswordSettingsLockout{}
}

func (a *PasswordPolicyPasswordSettingsLockout) IsPolicyInstance() bool {
	return true
}

func (a *PasswordPolicyPasswordSettingsLockout) MarshalJSON() ([]byte, error) {
	type Alias PasswordPolicyPasswordSettingsLockout
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.AutoUnlockMinutes != 0 {
		result.AutoUnlockMinutesPtr = Int64Ptr(a.AutoUnlockMinutes)
	}
	if a.MaxAttempts != 0 {
		result.MaxAttemptsPtr = Int64Ptr(a.MaxAttempts)
	}
	return json.Marshal(&result)
}

func (a *PasswordPolicyPasswordSettingsLockout) UnmarshalJSON(data []byte) error {
	type Alias PasswordPolicyPasswordSettingsLockout

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.AutoUnlockMinutesPtr != nil {
		a.AutoUnlockMinutes = *result.AutoUnlockMinutesPtr
		a.AutoUnlockMinutesPtr = result.AutoUnlockMinutesPtr
	}
	if result.MaxAttemptsPtr != nil {
		a.MaxAttempts = *result.MaxAttemptsPtr
		a.MaxAttemptsPtr = result.MaxAttemptsPtr
	}
	return nil
}
