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

//ConEvents Class Reference
//
//Interface to server events.
//
//This interface allow user to subscribe on server events and retrieve them.
//
//Public Member Functions
type ConEvents service

//ConEvents.Retrieve
//Retrieve.
//
//Use this method to retrieve events.
//
//	Parameters:
//	- pEvents (array) Events array. Each element of array is params with attributes:
//	+----------------------------+---------------+----------------------------+
//	|         Attribute          |     Type      |        Description         |
//	+----------------------------+---------------+----------------------------+
//	| KLSRV_CON_EVENT_TYPE       | string   | Event type                 |
//	| KLSRV_CON_EVENT_PARAMS     | paramParams   | Event params               |
//	| KLSRV_CON_EVENT_BIRTH_TIME | paramDateTime | UTC time of creation       |
//	| KLSRV_CON_EVENT_LIFE_TIME  | int64      | Event life time in seconds |
//	+----------------------------+---------------+----------------------------+
//
//	Returns:
//	- (bool) true if all events retrieved, otherwise - false
//
//	Exceptions:
//	- STDE_NOTFOUND	- subscription was not found
//	- STDE_UNAVAIL	- period of polling is too small.
//	Please use recommended period of polling.
func (ce *ConEvents) Retrieve(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.Retrieve",
		nil)

	if err != nil {
		return nil, err
	}

	raw, err := ce.client.Do(ctx, request, nil)
	return raw, err
}

//SubscribeEventParam struct
type SubscribeEventParam struct {
	WstrEvent string `json:"wstrEvent"`
}

//SubscribeEventResponse struct
type SubscribeEventResponse struct {
	NPeriod   int64 `json:"nPeriod"`
	PxgRetVal int64 `json:"PxgRetVal"`
}

//ConEvents.Subscribe
//Subscribe on event.
//
//Use this method to subscribe on events. Method returns period of polling.
//You should use it between retrieve calls.
//Also attribute pFilter allow you to cut off unnecessary events.
//
//	Parameters:
//	- wstrEvent (string) event type
//	- pFilter (params) event filter
//	- nPeriod (int64) new value of polling period, milliseconds
//
//	Returns:
//	- (int64) subscription id
//
//	Exceptions:
//	- STDE_TOOBIG	too many subscriptions for one session.
//	+---------------------------------+-------------------------------------------------------------------------+
//	|     Name of the event type      |                               Description                               |
//	+---------------------------------+-------------------------------------------------------------------------+
//	| KLPRCI_TaskState                | Task execution state also "task_new_state" attribute.                   |
//	| KLEVP_GroupTaskSyncState        | Task synchonization state changed; see also "task_new_state" attribute. |
//	| GNRL_EV_SUSPICIOUS_OBJECT_FOUND | Suspicious object found.                                                |
//	| GNRL_EV_VIRUS_FOUND             | Virus found.                                                            |
//	| GNRL_EV_OBJECT_CURED            | Object was cured.                                                       |
//	| GNRL_EV_OBJECT_DELETED          | Object was deleted.                                                     |
//	| GNRL_EV_OBJECT_REPORTED         | Object was reported.                                                    |
//	| GNRL_EV_PASSWD_ARCHIVE_FOUND    | Password protected archive was found.                                   |
//	| GNRL_EV_OBJECT_QUARANTINED      | Object was out into quarantine.                                         |
//	| GNRL_EV_OBJECT_NOTCURED         | Object wasn't cured.                                                    |
//	+---------------------------------+-------------------------------------------------------------------------+
//
//	type SubscribeEventParam struct { <- Subscribe interface example
//		WstrEvent string `json:"wstrEvent"`
//	}
//
//	type SubscribeEventResponse struct { <- Response
//		NPeriod   int64 `json:"nPeriod"`
//		PxgRetVal int64 `json:"PxgRetVal"`
//	}
func (ce *ConEvents) Subscribe(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.Subscribe", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := ce.client.Do(ctx, request, nil)
	return raw, err
}

//ConEvents.UnSubscribe
//UnSubscribe from event.
//
//Use this method to unsubscribe from an event.
//
//	- Parameters:
//	nSubsId (int64) subscription id (see method ConEvents.Subscribe)
//
//	Exceptions:
//	- STDE_NOTFOUND	subscription was not found
func (ce *ConEvents) UnSubscribe(ctx context.Context, nSubsId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nSubsId": %d}`, nSubsId))
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.UnSubscribe", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := ce.client.Do(ctx, request, nil)
	return raw, err
}
