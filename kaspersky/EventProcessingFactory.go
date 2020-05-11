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
	"net/http"
)

//	EventProcessingFactory Class Reference
//
//	Interface to create event processing iterators
//
//	List of all members:
type EventProcessingFactory service

//EventPFP struct
type EventPFP struct {
	PFilter           *PFilter    `json:"pFilter"`
	VecFieldsToOrder  interface{} `json:"vecFieldsToOrder"`
	VecFieldsToReturn []string    `json:"vecFieldsToReturn"`
	LifetimeSEC       int64       `json:"lifetimeSec"`
}

//	Create event processing iterator.
//
//	Parameters:
//	- params (EventPFP)
//
//	Example:
//	val, _, _ := client.EventProcessingFactory.CreateEventProcessing(ctx, kaspersky.EventPFP{
//		VecFieldsToOrder:  nil,
//		VecFieldsToReturn: []string{
//			"GNRL_EA_SEVERITY",
//			"product_name",
//			"hostname",
//			"task_display_name",
//			"event_type_display_name",
//			"event_type",
//			"body",
//		},
//		LifetimeSEC:       120,
//	})
//
//	Return:
//	- strIteratorId	(string) result-set ID, identifier of the server-side ordered collection of found data records.
//	The result-set is destroyed and associated memory is freed in following cases:
//
//	Passed lifetimeSec seconds after last access to the result-set (by methods EventProcessing. GetRecordCount
//
//	and EventProcessing.GetRecordRange).
//
//	Session to the Administration Server has been closed.
//
//	EventProcessing.ReleaseIterator has been called.
func (epf *EventProcessingFactory) CreateEventProcessing(ctx context.Context, params interface{}) (*StrIteratorId,
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

//	Create event processing iterator with filter.
//
//	Parameters:
//	- params (EventPFP)
//
//	Example:
//	val, _, _ := client.EventProcessingFactory.CreateEventProcessing(ctx, kaspersky.EventPFP{
//		VecFieldsToOrder:  nil,
//		PFilter: kaspersky.PFilter{
//			ProductName: "ess",
//			GnrlEaSeverity: 1,
//		},
//		VecFieldsToReturn: []string{
//			"GNRL_EA_SEVERITY",
//			"product_name",
//			"hostname",
//			"task_display_name",
//			"event_type_display_name",
//			"event_type",
//			"body",
//		},
//		LifetimeSEC:       120,
//	})
//
//Return:
//	- strIteratorId	(string) result-set ID,
//	identifier of the server-side ordered collection of found data records.
//
//	The result-set is destroyed and associated memory is freed in following cases:
//
//	Passed lifetimeSec seconds after last access to the result-set
//	(by methods EventProcessing.GetRecordCount and EventProcessing.GetRecordRange).
//
//	Session to the Administration Server has been closed.
//
//	EventProcessing.ReleaseIterator has been called.
func (epf *EventProcessingFactory) CreateEventProcessing2(ctx context.Context, params interface{}) (*StrIteratorId,
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

//EventPFH struct
type EventPFH struct {
	StrDomainName     string           `json:"strDomainName"`
	StrHostName       string           `json:"strHostName"`
	StrProduct        string           `json:"strProduct"`
	StrVersion        string           `json:"strVersion"`
	PFilter           *PFilter         `json:"pFilter"`
	VecFieldsToReturn []string         `json:"vecFieldsToReturn"`
	VecFieldsToOrder  *[]FieldsToOrder `json:"vecFieldsToOrder"`
	LifetimeSEC       int64            `json:"lifetimeSec"`
}

//	Create event processing iterator for host.
//
//	Parameters:
// params (EventPFH)
//
//	Example:
//	val, raw, _ := client.EventProcessingFactory.CreateEventProcessingForHost2(ctx, kaspersky.EventPFH{
//		StrDomainName: "domain.ru",
//		StrHostName:   "169b91af-bba5-480f-9f67-2ecb4800be78",
//		StrProduct:    "1093",
//		StrVersion:    "1.0.0.0",
//		VecFieldsToReturn: []string{
//			"product_name",
//			"product_version",
//			"product_displ_version",
//			"task_display_name",
//			"group_id",
//			"event_type",
//			"event_type_display_name",
//			"GNRL_EA_SEVERITY",
//			"GNRL_EA_DESCRIPTION",
//		},
//		VecFieldsToOrder: &[]kaspersky.FieldsToOrder{{
//			Type: "params",
//			OrderValue: kaspersky.OrderValue{
//				Name: "event_type",
//				Asc:  true,
//			},
//		}},
//		LifetimeSEC: 120,
//	})
//
//	Return:
//	- strIteratorId	(string) result-set ID,
//	identifier of the server-side ordered collection of found data records.
//	The result-set is destroyed and associated memory is freed in following cases:
//	Passed lifetimeSec seconds after last access to the result-set
//	(by methods EventProcessing.GetRecordCount and EventProcessing.GetRecordRange).
//	Session to the Administration Server has been closed.
//	EventProcessing.ReleaseIterator has been called.
func (epf *EventProcessingFactory) CreateEventProcessingForHost(ctx context.Context, params interface{}) (*StrIteratorId,
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

//	Create event processing iterator with filter for host.
//
//	Parameters:
// params (EventPFH)
//
//	Example:
//	val, raw, _ := client.EventProcessingFactory.CreateEventProcessingForHost2(ctx, kaspersky.EventPFH{
//		StrDomainName: "domain.ru",
//		StrHostName:   "169b91af-bba5-480f-9f67-2ecb4800be78",
//		StrProduct:    "1093",
//		StrVersion:    "1.0.0.0",
//		VecFieldsToReturn: []string{
//			"product_name",
//			"product_version",
//			"product_displ_version",
//			"task_display_name",
//			"group_id",
//			"event_type",
//			"event_type_display_name",
//			"GNRL_EA_SEVERITY",
//			"GNRL_EA_DESCRIPTION",
//		},
//		PFilter: &kaspersky.PFilter{
//			GnrlEaSeverity: 1,
//		},
//		VecFieldsToOrder: &[]kaspersky.FieldsToOrder{{
//			Type: "params",
//			OrderValue: kaspersky.OrderValue{
//				Name: "event_type",
//				Asc:  true,
//			},
//		}},
//		LifetimeSEC: 120,
//	})
//
//	Return:
//	- strIteratorId	(string) result-set ID, identifier of the server-side ordered collection of found data records.
//	The result-set is destroyed and associated memory is freed in following cases:
//	Passed lifetimeSec seconds after last access to the result-set
//	(by methods EventProcessing.GetRecordCount and EventProcessing.GetRecordRange).
//	Session to the Administration Server has been closed.
//	EventProcessing.ReleaseIterator has been called.
func (epf *EventProcessingFactory) CreateEventProcessingForHost2(ctx context.Context, params interface{}) (*StrIteratorId,
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
