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
	"log"
	"net/http"
)

// SrvIpmNewsAndStatistics In-product marketing support interface
type SrvIpmNewsAndStatistics service

type Parameters struct {
	PwchContentID string `json:"pwchContentId"`
}

type TrackingData struct {
	PxgRetVal []struct {
		Type  string `json:"type"`
		Value struct {
			ApplicationID      int    `json:"ApplicationId"`
			ApplicationVersion string `json:"ApplicationVersion"`
			HardwareID         string `json:"HardwareId"`
			LicenseID          string `json:"LicenseId"`
			Localization       string `json:"Localization"`
			LtsID              string `json:"LtsId"`
		} `json:"value"`
	} `json:"PxgRetVal"`
}

// GetTrackingData Gets TrackingData for specified content.
func (sins *SrvIpmNewsAndStatistics) GetTrackingData(ctx context.Context, params Parameters) (*TrackingData, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", sins.client.Server+"/api/v1.0/SrvIpmNewsAndStatistics.GetTrackingData",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	trackingData := new(TrackingData)
	raw, err := sins.client.Do(ctx, request, &trackingData)

	if sins.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return trackingData, err
}

// SendStatistics Send IPM statistics.
func (sins *SrvIpmNewsAndStatistics) SendStatistics(ctx context.Context, params Parameters) error {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", sins.client.Server+"/api/v1.0/SrvIpmNewsAndStatistics.SendStatistics",
		bytes.NewBuffer(postData))

	if err != nil {
		return err
	}

	raw, err := sins.client.Do(ctx, request, nil)

	if sins.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}
