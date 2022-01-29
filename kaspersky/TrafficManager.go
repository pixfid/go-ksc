/*
 * MIT License
 *
 * Copyright (c) [2020] [Semchenko Aleksandr]
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package kaspersky

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"net/http"
)

// TrafficManager service allows to limit network traffic speed between KSC server and Network agents or between servers within server hierarchy.
type TrafficManager service

//TrafficRestrictions struct
type TrafficRestrictions struct {
	TrafficPRestrictions TrafficPRestrictions `json:"pRestriction"`
}

// TrafficPRestrictions struct
type TrafficPRestrictions struct {
	// TrfmRestrFromHour time period start hour, 0-23
	TrfmRestrFromHour int64 `json:"TRFM_RESTR_FROM_HOUR"`

	//TrfmRestrFromMin time period start minute, 0-59
	TrfmRestrFromMin int64 `json:"TRFM_RESTR_FROM_MIN"`

	// TrfmRestrIp4High High border of IP addresses diapason
	TrfmRestrIp4High string `json:"TRFM_RESTR_IP4_HIGH"`

	// TrfmRestrIp4Low Low border of IP addresses diapason
	TrfmRestrIp4Low string `json:"TRFM_RESTR_IP4_LOW"`

	// TrfmRestrIp4Mask IP subnet mask
	TrfmRestrIp4Mask string `json:"TRFM_RESTR_IP4_MASK,omitempty"`

	// TrfmRestrIp4Subnet IP subnet
	TrfmRestrIp4Subnet string `json:"TRFM_RESTR_IP4_SUBNET,omitempty"`

	// TrfmRestrLimit limit for all other time, kilobytes per second
	TrfmRestrLimit int64 `json:"TRFM_RESTR_LIMIT"`

	// TrfmRestrTimeLimit limit for specified time, kilobytes per second
	TrfmRestrTimeLimit int64 `json:"TRFM_RESTR_TIME_LIMIT"`

	// TrfmRestrToHour time period start hour, 0-23
	TrfmRestrToHour int64 `json:"TRFM_RESTR_TO_HOUR"`

	// TrfmRestrToMin time period end minute, 0-59
	TrfmRestrToMin int64 `json:"TRFM_RESTR_TO_MIN"`
}

// AddRestriction Add traffic restriction.
func (tm *TrafficManager) AddRestriction(ctx context.Context, params TrafficRestrictions) (*PxgValInt, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", tm.client.Server+"/api/v1.0/TrafficManager.AddRestriction", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := tm.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// DeleteRestriction Remove traffic restriction.
func (tm *TrafficManager) DeleteRestriction(ctx context.Context, nRestrictionId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nRestrictionId": %d}`, nRestrictionId))
	request, err := http.NewRequest("POST", tm.client.Server+"/api/v1.0/TrafficManager.DeleteRestriction", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := tm.client.Request(ctx, request, nil)
	return raw, err
}

// GetRestrictions Returns all currently active restrictions list.
func (tm *TrafficManager) GetRestrictions(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", tm.client.Server+"/api/v1.0/TrafficManager.GetRestrictions", nil)
	if err != nil {
		return nil, err
	}

	raw, err := tm.client.Request(ctx, request, nil)
	return raw, err
}

// UpdateRestriction Modify existing traffic restriction settings.
func (tm *TrafficManager) UpdateRestriction(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", tm.client.Server+"/api/v1.0/TrafficManager.UpdateRestriction", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := tm.client.Request(ctx, request, nil)
	return raw, err
}
