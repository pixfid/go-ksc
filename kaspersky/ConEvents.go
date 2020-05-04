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
	"fmt"
	"log"
	"net/http"
)

//	ConEvents Class Reference
//
//	Interface to server events.
//
//	This interface allow user to subscribe on server events and retrieve them.
//
//	Public Member Functions
type ConEvents struct {
	client *Client
}

//	Retrieve.
//
//	Use this method to retrieve events.
//
//	Parameters:
//	- pEvents	[out] (array) Events array. Each element of array is params with attributes:
//	+----------------------------+---------------+----------------------------+
//	|         Attribute          |     Type      |        Description         |
//	+----------------------------+---------------+----------------------------+
//	| KLSRV_CON_EVENT_TYPE       | paramString   | Event type                 |
//	| KLSRV_CON_EVENT_PARAMS     | paramParams   | Event params               |
//	| KLSRV_CON_EVENT_BIRTH_TIME | paramDateTime | UTC time of creation       |
//	| KLSRV_CON_EVENT_LIFE_TIME  | paramInt      | Event life time in seconds |
//	+----------------------------+---------------+----------------------------+
//
//	Returns:
//	- (boolean) true if all events retrieved, otherwise - false
//
//	Exceptions:
//	- STDE_NOTFOUND	- subscription was not found
//	- STDE_UNAVAIL	- period of polling is too small.
//	Please use recommended period of polling.
func (ce *ConEvents) Retrieve(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.Retrieve",
		nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := ce.client.Do(ctx, request, nil)
	return raw, err
}

//	Subscribe on event.
//
//	Use this method to subscribe on events. Method returns period of polling. You should use it between retrieve calls. Also attribute pFilter allow you to cut off unnecessary events.
//
//	Parameters:
//	- wstrEvent	[in] (string) event type
//	- pFilter	[in] (params) event filter
//	- nPeriod	[out] (int64) new value of polling period, milliseconds
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
func (ce *ConEvents) Subscribe(ctx context.Context, v interface{}) ([]byte, error) {
	postData, _ := json.Marshal(v)
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.Subscribe", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := ce.client.Do(ctx, request, nil)
	return raw, err
}

//	UnSubscribe from event.
//
//	Use this method to unsubscribe from an event.
//
//	- Parameters:
//	nSubsId	[in] (int64) subscription id (see method ConEvents::Subscribe)
//
//	Exceptions:
//	- STDE_NOTFOUND	subscription was not found
func (ce *ConEvents) UnSubscribe(ctx context.Context, nSubsId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nSubsId": %d}`, nSubsId))
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/ConEvents.UnSubscribe", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := ce.client.Do(ctx, request, nil)
	return raw, err
}
