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

// EventProcessingFactory service to create event processing iterators
type EventProcessingFactory service

// EventPFP struct
type EventPFP struct {
	PFilter           PFilter       `json:"pFilter,omitempty"`
	VecFieldsToOrder  FieldsToOrder `json:"vecFieldsToOrder,omitempty"`
	VecFieldsToReturn []string      `json:"vecFieldsToReturn"`
	LifetimeSEC       int64         `json:"lifetimeSec"`
}

// CreateEventProcessing Create event processing iterator.
func (epf *EventProcessingFactory) CreateEventProcessing(ctx context.Context, params EventPFP) (*StrIteratorId,
	[]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", epf.client.Server+"/api/v1.0/EventProcessingFactory.CreateEventProcessing",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	strIteratorId := new(StrIteratorId)
	raw, err := epf.client.Do(ctx, request, &strIteratorId)
	return strIteratorId, raw, err
}

// CreateEventProcessing2 Create event processing iterator with filter.
func (epf *EventProcessingFactory) CreateEventProcessing2(ctx context.Context, params EventPFP) (*StrIteratorId,
	[]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", epf.client.Server+"/api/v1.0/EventProcessingFactory.CreateEventProcessing2",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	strIteratorId := new(StrIteratorId)
	raw, err := epf.client.Do(ctx, request, &strIteratorId)
	return strIteratorId, raw, err
}

// EventPFH struct
type EventPFH struct {
	StrDomainName     string          `json:"strDomainName"`
	StrHostName       string          `json:"strHostName"`
	StrProduct        string          `json:"strProduct"`
	StrVersion        string          `json:"strVersion"`
	PFilter           *PFilter        `json:"pFilter"`
	VecFieldsToReturn []string        `json:"vecFieldsToReturn"`
	VecFieldsToOrder  []FieldsToOrder `json:"vecFieldsToOrder"`
	LifetimeSEC       int64           `json:"lifetimeSec"`
}

// CreateEventProcessingForHost Create event processing iterator for host.
func (epf *EventProcessingFactory) CreateEventProcessingForHost(ctx context.Context, params EventPFH) (*StrIteratorId,
	[]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", epf.client.Server+"/api/v1.0/EventProcessingFactory.CreateEventProcessingForHost",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	strIteratorId := new(StrIteratorId)
	raw, err := epf.client.Do(ctx, request, &strIteratorId)
	return strIteratorId, raw, err
}

// CreateEventProcessingForHost2 Create event processing iterator with filter for host.
func (epf *EventProcessingFactory) CreateEventProcessingForHost2(ctx context.Context, params EventPFH) (*StrIteratorId,
	[]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", epf.client.Server+"/api/v1.0/EventProcessingFactory.CreateEventProcessingForHost2",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	strIteratorId := new(StrIteratorId)
	raw, err := epf.client.Do(ctx, request, &strIteratorId)
	return strIteratorId, raw, err
}
