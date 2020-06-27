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

// HostTagsApi service allows to acquire and manage tags for hosts. It is additional service for common ListTags.
type HostTagsApi service

//HostTagsParams using in
type HostTagsParams struct {
	//SzwHostID host identifier ( guid )
	SzwHostID string `json:"szwHostId"`
	//PParams reserved.
	PParams Null `json:"pParams"`
}

//HostTags struct using in HostTagsApi.GetHostTags
type HostTags struct {
	HTags []HTags `json:"PxgRetVal"`
}

type HTags struct {
	Type      string    `json:"type"`
	HTagValue HTagValue `json:"value,omitempty"`
}

type HTagValue struct {
	// KLHSTTagValue Value of the tag
	KLHSTTagValue string `json:"KLHST_TagValue"`

	// KlhstIsTagSetByProduct true if tag has been set by product
	KlhstIsTagSetByProduct bool `json:"KLHST_IS_TAG_SET_BY_PRODUCT,omitempty"`

	// KlhstIsTagSetByHosttagrule true if tag has been set by host tag rule
	KlhstIsTagSetByHosttagrule bool `json:"KLHST_IS_TAG_SET_BY_HOSTTAGRULE,omitempty"`
}

// GetHostTags Get tags for the host.
func (kc *HostTagsApi) GetHostTags(ctx context.Context, params HostTagsParams) (*HostTags, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", kc.client.Server+"/api/v1.0/HostTagsApi.GetHostTags", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	hostTags := new(HostTags)
	raw, err := kc.client.Do(ctx, request, &hostTags)
	return hostTags, raw, err
}
