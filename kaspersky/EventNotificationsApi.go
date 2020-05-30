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

//	EventNotificationsApi Class Reference
//
//	Allows to publish event with Administration Server as publisher.
//
//	List of all members.
type EventNotificationsApi service

//EventNotificationParams struct
type EventNotificationParams struct {
	WstrEventType string     `json:"wstrEventType"`
	PEventBody    PEventBody `json:"pEventBody"`
	TmBirthTime   string     `json:"tmBirthTime"`
}

type PEventBody struct {
	GnrlEaParam2 string `json:"GNRL_EA_PARAM_2"`
	GnrlEaParam5 string `json:"GNRL_EA_PARAM_5"`
	GnrlEaParam7 string `json:"GNRL_EA_PARAM_7"`
	GnrlEaParam8 int64  `json:"GNRL_EA_PARAM_8"`
}

//	Publish event.
//
//	Publishes event with Administration Server as publisher
//
//	Parameters:
//	- params (interface{}) with fields:
//	|- wstrEventType	(string) event type
//	|- pEventBody	(params) event body, content depends on event type
//	|- tmBirthTime	(datetime) time when event was published
func (ts *EventNotificationsApi) PublishEvent(ctx context.Context, params EventNotificationParams) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/EventNotificationsApi.PublishEvent", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}
