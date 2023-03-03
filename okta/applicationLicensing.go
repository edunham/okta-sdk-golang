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

type ApplicationLicensing struct {
	SeatCount    int64
	SeatCountPtr *int64 `json:"seatCount,omitempty"`
}

func (a *ApplicationLicensing) MarshalJSON() ([]byte, error) {
	type Alias ApplicationLicensing
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.SeatCount != 0 {
		result.SeatCountPtr = Int64Ptr(a.SeatCount)
	}
	return json.Marshal(&result)
}

func (a *ApplicationLicensing) UnmarshalJSON(data []byte) error {
	type Alias ApplicationLicensing

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.SeatCountPtr != nil {
		a.SeatCount = *result.SeatCountPtr
		a.SeatCountPtr = result.SeatCountPtr
	}
	return nil
}
