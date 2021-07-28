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
	"log"
	"net/http"
)

// SrvRi "Remote installation" task results
type SrvRi service

// ShouldForceReboot Check if the administrator has started a forced reboot of the host.
// This method is only for Network Agents.
func (sr *SrvRi) ShouldForceReboot(ctx context.Context, wstrHostID, wstrTaskID string) (*PxgValBool, error) {
	postData := []byte(fmt.Sprintf(`{"wstrHostID" : "%s", "wstrTaskID" : "%s"}`, wstrHostID, wstrTaskID))

	request, err := http.NewRequest("POST", sr.client.Server+"/api/v1.0/SrvRi.ShouldForceReboot",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	result := new(PxgValBool)
	raw, err := sr.client.Do(ctx, request, &result)

	if sr.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

type RiTask struct {
	WstrHostID string `json:"wstrHostID"`
	PData      struct {
		Type  string `json:"type"`
		Value struct {
			Type  string `json:"type"`
			Value struct {
				KLRISRVRITASKID   string `json:"KLRI_SRVRI_TASK_ID"`
				KLRISRVRIRESDESCR string `json:"KLRI_SRVRI_RES_DESCR"`
				KLRISRVRIRESCODE  int    `json:"KLRI_SRVRI_RES_CODE"`
			} `json:"value"`
		} `json:"value"`
	} `json:"pData"`
}

// SetRiTaskResults Send to server "Remote installation" task results. This method is only for Network Agents.
func (sr *SrvRi) SetRiTaskResults(ctx context.Context, params RiTask) (*PropagationState, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sr.client.Server+"/api/v1.0/SrvRi.SetRiTaskResults",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sr.client.Do(ctx, request, nil)

	if sr.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return nil, err
}

func (sr *SrvRi) SetRebootConfirmedHosts(ctx context.Context, wstrHostID, wstrTaskID string) (*PxgValBool, error) {
	postData := []byte(fmt.Sprintf(`{"strDomain": "%s", "nType" : "%s" }`, wstrHostID, wstrTaskID))

	request, err := http.NewRequest("POST", sr.client.Server+"/api/v1.0/SrvRi.SetRebootConfirmedHosts",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	result := new(PxgValBool)
	raw, err := sr.client.Do(ctx, request, &result)

	if sr.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}
