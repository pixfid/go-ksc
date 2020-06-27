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

// ConEvents service to server events. This interface allow user to subscribe on server events and retrieve them.
type ConEvents service

// Retrieve Use this method to retrieve events.
func (ce *ConEvents) Retrieve(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.Retrieve",
		nil)

	if err != nil {
		return nil, err
	}

	raw, err := ce.client.Do(ctx, request, nil)
	return raw, err
}

// SubscribeEventParam struct
type SubscribeEventParam struct {
	WstrEvent string `json:"wstrEvent"`
}

// SubscribeEventResponse struct
type SubscribeEventResponse struct {
	NPeriod   int64 `json:"nPeriod"`
	PxgRetVal int64 `json:"PxgRetVal"`
}

// Subscribe on event. Use this method to subscribe on events. Method returns period of polling.
// You should use it between retrieve calls. Also attribute pFilter allow you to cut off unnecessary events.
func (ce *ConEvents) Subscribe(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.Subscribe", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := ce.client.Do(ctx, request, nil)
	return raw, err
}

// UnSubscribe from event. Use this method to unsubscribe from an event.
func (ce *ConEvents) UnSubscribe(ctx context.Context, nSubsId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nSubsId": %d}`, nSubsId))
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.UnSubscribe", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := ce.client.Do(ctx, request, nil)
	return raw, err
}
