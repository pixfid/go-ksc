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

// ConEvents service to server events. This interface allow user to subscribe on server events and retrieve them.
type ConEvents service

type EventRetrieve struct {
	PEvents   []interface{} `json:"pEvents"`
	NPeriod   int64         `json:"nPeriod"`
	PxgRetVal bool          `json:"PxgRetVal"`
}

// Retrieve Use this method to retrieve events.
func (ce *ConEvents) Retrieve(ctx context.Context) (*EventRetrieve, error) {
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.Retrieve", nil)

	if err != nil {
		return nil, err
	}

	eventRetrieve := new(EventRetrieve)
	raw, err := ce.client.Do(ctx, request, &eventRetrieve)

	if ce.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return eventRetrieve, err
}

// EventSubscribeParams struct
type EventSubscribeParams struct {
	WstrEvent string           `json:"wstrEvent"`
	PFilter   ESubscribeFilter `json:"pFilter,omitempty"`
}

type ESubscribeFilter struct {
	Type  string     `json:"type"`
	Value ESubscribe `json:"value"`
}

type ESubscribe struct {
	ProductName string `json:"product_name,omitempty"`
}

// SubscribeEventResponse struct
type SubscribeEventResponse struct {
	NPeriod   int64 `json:"nPeriod"`
	PxgRetVal int64 `json:"PxgRetVal"`
}

// Subscribe on event. Use this method to subscribe on events. Method returns period of polling.
// You should use it between retrieve calls. Also attribute pFilter allow you to cut off unnecessary events.
func (ce *ConEvents) Subscribe(ctx context.Context, params EventSubscribeParams) (*SubscribeEventResponse, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.Subscribe", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	subscribeEventResponse := new(SubscribeEventResponse)
	raw, err := ce.client.Do(ctx, request, &subscribeEventResponse)

	if ce.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return subscribeEventResponse, err
}

// UnSubscribe from event. Use this method to unsubscribe from an event.
func (ce *ConEvents) UnSubscribe(ctx context.Context, nSubsId int64) error {
	postData := []byte(fmt.Sprintf(`{"nSubsId": %d}`, nSubsId))
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.UnSubscribe", bytes.NewBuffer(postData))

	if err != nil {
		return err
	}

	raw, err := ce.client.Do(ctx, request, nil)

	if ce.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

// IsAnyServiceConsoleAvailable Check any service console availability.
func (ce *ConEvents) IsAnyServiceConsoleAvailable(ctx context.Context) (*PxgValBool, error) {
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.IsAnyServiceConsoleAvailable", nil)

	if err != nil {
		return nil, err
	}

	result := new(PxgValBool)
	raw, err := ce.client.Do(ctx, request, &result)

	if ce.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

// IsServiceConsoleAvailable
// Checks whether service console for specified product is connected
// to KSC server and is able to execute Product backend commands
func (ce *ConEvents) IsServiceConsoleAvailable(ctx context.Context, wstrProdName, wstrProdVersion string) error {
	postData := []byte(fmt.Sprintf(`{"wstrProdName" : "%s", "wstrProdVersion" : "%s"}`, wstrProdName, wstrProdVersion))

	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.IsServiceConsoleAvailable",
		bytes.NewBuffer(postData))

	if err != nil {
		return err
	}

	raw, err := ce.client.Do(ctx, request, nil)

	if ce.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}
