/*
 *
 * 	Copyright (C) 2020  <Semchenko Aleksandr>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.If not, see <http://www.gnu.org/licenses/>.
 * /
 */

package kaspersky

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

//	EventNotificationsApi Class Reference
//
//	Allows to publish event with Administration Server as publisher.
//
//	List of all members.
type EventNotificationsApi struct {
	client *Client
}

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
func (ts *EventNotificationsApi) PublishEvent(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)

	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/EventNotificationsApi.PublishEvent", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}
	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}
