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
	"fmt"
	"net/http"
)

//	HostTagsApi Class Reference
//	Interface allows to acquire and manage tags for hosts. It is additional interface for common ListTags.
//
//	List of all members.
type HostTagsApi service

//HostTags struct using in HostTagsApi.GetHostTags
type HostTags struct {
	//ppHostTags	(array) collection of (params) objects where each of them has the following structure:
	PPHostTags []PPHostTags `json:"PxgRetVal"`
}

type PPHostTags struct {
	Type      string    `json:"type"`
	PPHostTag PPHostTag `json:"value"`
}

type PPHostTag struct {
	//Value of the tag
	KLHSTTagValue string `json:"KLHST_TagValue"`

	//true if tag has been set by product
	KlhstIsTagSetByProduct bool `json:"KLHST_IS_TAG_SET_BY_PRODUCT"`

	//true if tag has been set by host tag rule
	KlhstIsTagSetByHosttagrule bool `json:"KLHST_IS_TAG_SET_BY_HOSTTAGRULE"`
}

//	Get tags for the host.
//
//	Parameters:
//	- szwHostId	[in] (string) - host identifier ( guid )
//	- pParams	(params) reserved.
//
//	Return:
//	- ppHostTags	(array) collection of (params) objects where each of them has the following structure:
//	|- "KLHST_TagValue" (string). Value of the tag
//	|- "KLHST_IS_TAG_SET_BY_PRODUCT" (bool). true if tag has been set by product
//	|- "KLHST_IS_TAG_SET_BY_HOSTTAGRULE" (bool). true if tag has been set by host tag rule
func (kc *HostTagsApi) GetHostTags(ctx context.Context, szwHostId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwHostId": "%s"}`, szwHostId))
	request, err := http.NewRequest("POST", kc.client.Server+"/api/v1.0/HostTagsApi.GetHostTags", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := kc.client.Do(ctx, request, nil)
	return raw, err
}
