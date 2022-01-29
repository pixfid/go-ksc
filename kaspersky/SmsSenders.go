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

	"net/http"
)

// SmsSenders service provide to configure mobile devices as SMS senders.
type SmsSenders service

// HasAllowedSenders checks if there is a device allowed to send SMS
func (ss *SmsSenders) HasAllowedSenders(ctx context.Context) (*PxgValBool, []byte, error) {
	request, err := http.NewRequest("POST", ss.client.Server+"/api/v1.0/SmsSenders.HasAllowedSenders", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := ss.client.Request(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// PNewStatuses struct
type PNewStatuses struct {
	PNewStatus PNewStatus `json:"pNewStatuses,omitempty"`
}

// PNewStatus struct
type PNewStatus struct {
	Type            string          `json:"type,omitempty"`
	PNewStatusValue PNewStatusValue `json:"value,omitempty"`
}

// PNewStatusValue struct
type PNewStatusValue struct {
	Name              string `json:"name,omitempty"`
	BMayUseSMSSending bool   `json:"bMayUseSmsSending,omitempty"`
}

// AllowSenders change bMayUseSmsSending parameter for mobile devices
func (ss *SmsSenders) AllowSenders(ctx context.Context, params PNewStatuses) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ss.client.Server+"/api/v1.0/SmsSenders.AllowSenders",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ss.client.Request(ctx, request, nil)
	return raw, err
}
