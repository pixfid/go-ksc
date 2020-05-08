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

//	TrafficManager Class Reference
//
//	Traffic manager interface.
//
//	Traffic manager allows to limit network traffic speed between
//	KSC server and Network agents or between servers within server hierarchy.
//	This interface manages restrictions used by traffic manager.
//
//	+-----------------------+--------+--------------------------------------------------------------------------------------------------------+
//	|         Param         |  Type  |                                              Description                                               |
//	+-----------------------+--------+--------------------------------------------------------------------------------------------------------+
//	| TRFM_RESTR_IP4_LOW    | string | Low border of IP addresses diapason                                                                    |
//	| TRFM_RESTR_IP4_HIGH   | string | High border of IP addresses diapason                                                                   |
//	| TRFM_RESTR_IP4_SUBNET | string | IP subnet                                                                                              |
//	| TRFM_RESTR_IP4_MASK   | string | IP subnet mask                                                                                         |
//	| TRFM_RESTR_FROM_HOUR  | int64  | time period start hour, 0-23                                                                           |
//	| TRFM_RESTR_FROM_MIN   | int64  | time period start minute, 0-59                                                                         |
//	| TRFM_RESTR_TO_HOUR    | int64  | time period end hour, 0-23                                                                             |
//	| TRFM_RESTR_TO_MIN     | int64  | time period end minute, 0-59                                                                           |
//	| TRFM_RESTR_TIME_LIMIT | int64  | limit for specified time, kilobytes per second                                                         |
//	| TRFM_RESTR_LIMIT      | int64  | limit for all other time, kilobytes per second                                                         |
//	| TRFM_RESTR_ID         | int64  | restriction id. This is output-only field and is ignored when passed to Add/Update restriction methods |
//	+-----------------------+--------+--------------------------------------------------------------------------------------------------------+
//
//	List of all members:
type TrafficManager struct {
	client *Client
}

//TrafficRestrictions struct
type TrafficRestrictions struct {
	TrafficPRestrictions TrafficPRestrictions `json:"pRestriction"`
}

type TrafficPRestrictions struct {
	TrfmRestrFromHour  int64   `json:"TRFM_RESTR_FROM_HOUR"`
	TrfmRestrFromMin   int64   `json:"TRFM_RESTR_FROM_MIN"`
	TrfmRestrIp4High   *string `json:"TRFM_RESTR_IP4_HIGH"`
	TrfmRestrIp4Low    *string `json:"TRFM_RESTR_IP4_LOW"`
	TrfmRestrIp4Mask   *string `json:"TRFM_RESTR_IP4_MASK,omitempty"`
	TrfmRestrIp4Subnet *string `json:"TRFM_RESTR_IP4_SUBNET,omitempty"`
	TrfmRestrLimit     int64   `json:"TRFM_RESTR_LIMIT"`
	TrfmRestrTimeLimit int64   `json:"TRFM_RESTR_TIME_LIMIT"`
	TrfmRestrToHour    int64   `json:"TRFM_RESTR_TO_HOUR"`
	TrfmRestrToMin     int64   `json:"TRFM_RESTR_TO_MIN"`
}

//
//	Parameters:
//	 [in]	pRestriction	Restriction definition, see Traffic restrictions for details and attributes meaning
//
//	Example of pRestriction:
//	{
//	  "pRestriction": {
//	    "TRFM_RESTR_FROM_HOUR": 0,
//	    "TRFM_RESTR_FROM_MIN": 0,
//	    "TRFM_RESTR_IP4_HIGH": "10.10.10.25",
//	    "TRFM_RESTR_IP4_LOW": "10.10.10.20",
//	    "TRFM_RESTR_LIMIT": 3000,
//	    "TRFM_RESTR_TIME_LIMIT": 3000,
//	    "TRFM_RESTR_TO_HOUR": 23,
//	    "TRFM_RESTR_TO_MIN": 59
//	  }
//	}
//
//	Returns:
//	- (int64) added restriction id
func (uc *TrafficManager) AddRestriction(ctx context.Context, params interface{}) (*PxgValInt, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/TrafficManager.AddRestriction", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	pxgValInt := new(PxgValInt)
	raw, err := uc.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Remove traffic restriction.
//
//	Parameters:
//	- nRestrictionId	(int64)	restriction to delete
func (uc *TrafficManager) DeleteRestriction(ctx context.Context, nRestrictionId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nRestrictionId": %d}`, nRestrictionId))
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/TrafficManager.DeleteRestriction", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

//	Returns all currently active restrictions list.
//
//	Parameters:
//	- [out]	pRestrictions	array of restrictions
//
//	Example Response:
//		{
//	  "pRestrictions" : [
//	    {
//	      "type" : "params",
//	      "value" : {
//	        "TRFM_RESTR_FROM_HOUR" : 0,
//	        "TRFM_RESTR_FROM_MIN" : 0,
//	        "TRFM_RESTR_ID" : 2,
//	        "TRFM_RESTR_IP4_HIGH" : "10.10.10.254",
//	        "TRFM_RESTR_IP4_LOW" : "10.10.10.2",
//	        "TRFM_RESTR_LIMIT" : 3000,
//	        "TRFM_RESTR_TIME_LIMIT" : 3000,
//	        "TRFM_RESTR_TO_HOUR" : 23,
//	        "TRFM_RESTR_TO_MIN" : 59
//	      }
//	    },
//	    {
//	      "type" : "params",
//	      "value" : {
//	        "TRFM_RESTR_FROM_HOUR" : 20,
//	        "TRFM_RESTR_FROM_MIN" : 0,
//	        "TRFM_RESTR_ID" : 3,
//	        "TRFM_RESTR_IP4_MASK" : "255.255.255.0",
//	        "TRFM_RESTR_IP4_SUBNET" : "10.10.10.0",
//	        "TRFM_RESTR_LIMIT" : 3001,
//	        "TRFM_RESTR_TIME_LIMIT" : 3001,
//	        "TRFM_RESTR_TO_HOUR" : 23,
//	        "TRFM_RESTR_TO_MIN" : 59
//	      }
//	    }
//	  ]
//	}
func (uc *TrafficManager) GetRestrictions(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/TrafficManager.GetRestrictions", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

//
//	Parameters:
//	- [in]	nRestrictionId	restriction to modify.
//
//	If restriction with such id does not exist then new restriction will be added and this parameters will be ignored.
//	- [in]	pRestriction	new restriction settings, see Traffic restrictions for details and attributes meaning
//
//	Example Update Restriction Params:
//	{
//			"nRestrictionId" : 4,
//			"pRestriction" : {
//				"TRFM_RESTR_FROM_HOUR" : 20,
//				"TRFM_RESTR_FROM_MIN" : 0,
//				"TRFM_RESTR_ID" : 3,
//				"TRFM_RESTR_IP4_MASK" : "255.255.255.0",
//				"TRFM_RESTR_IP4_SUBNET" : "10.10.10.0",
//				"TRFM_RESTR_LIMIT" : 3001,
//				"TRFM_RESTR_TIME_LIMIT" : 3001,
//				"TRFM_RESTR_TO_HOUR" : 23,
//				"TRFM_RESTR_TO_MIN" : 59
//			}
//	}
//
//	Returns:
//	- (int) modified restriction id. If restriction did not exist before call then newly created restriction id.
func (uc *TrafficManager) UpdateRestriction(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/TrafficManager.UpdateRestriction", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}
