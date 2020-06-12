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

//	QueriesStorage Class Reference
//
//	QueriesStorage interface.
//
//	List of all members.
type QueriesStorage service

//	New query registration.
//
//	Creates a new query and stores it for the current user
//	(associated with the connection to the Administration Server).
func (qs *QueriesStorage) AddQuery(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", qs.client.Server+"/api/v1.0/QueriesStorage.AddQuery",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := qs.client.Do(ctx, request, nil)
	return raw, err
}

//	Delete query.
//
//	Deletes the query with the specified ID.
//
//	Parameters:
//	- nId	(int64) query id
func (qs *QueriesStorage) DeleteQuery(ctx context.Context, nId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nId": %d}`, nId))
	request, err := http.NewRequest("POST", qs.client.Server+"/api/v1.0/QueriesStorage.DeleteQuery", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := qs.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire array of queries params of given type.
//
//	Returns array of IDs
//	and data of all queries of given type defined for the current user
//	(associated with the connection to the Administration Server).
//
//	Parameters:
//	- eType	(int) queries type QueryType
//
//	Returns:
//	- (array) array object, each element is a KLPAR::ParamsPtr object query data,
//	see Query storage data format
func (qs *QueriesStorage) GetQueries(ctx context.Context, eType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"eType": %d}`, eType))
	request, err := http.NewRequest("POST", qs.client.Server+"/api/v1.0/QueriesStorage.GetQueries", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := qs.client.Do(ctx, request, nil)
	return raw, err
}

//QueryParams struct
type QueryParams struct {
	QueryParamVal *QueryParamVal `json:"PxgRetVal,omitempty"`
}

type QueryParamVal struct {
	KlqrsQueryData *KlqrsQueryData `json:"KLQRS_QUERY_DATA,omitempty"`
	KlqrsQueryGUID string          `json:"KLQRS_QUERY_GUID,omitempty"`
	KlqrsQueryType int64           `json:"KLQRS_QUERY_TYPE,omitempty"`
}

type KlqrsQueryData struct {
	Type  string               `json:"type,omitempty"`
	Value *KLQRSQUERYDATAValue `json:"value,omitempty"`
}

type KLQRSQUERYDATAValue struct {
	Name                   string        `json:"Name,omitempty"`
	PredefinedID           string        `json:"PredefinedID,omitempty"`
	NetInfoPageSettings    *PageSettings `json:"NetInfoPageSettings,omitempty"`
	NetInfoExPageSettings  *PageSettings `json:"NetInfoExPageSettings,omitempty"`
	ProtectionPageSettings *PageSettings `json:"ProtectionPageSettings,omitempty"`
	Query                  string        `json:"Query,omitempty"`
}

type PageSettings struct {
	Type  string             `json:"type,omitempty"`
	Value *PageSettingsValue `json:"value,omitempty"`
}

type PageSettingsValue struct {
	Query          string `json:"Query,omitempty"`
	StatusID       int64  `json:"StatusId,omitempty"`
	FoundLastNDays int64  `json:"FoundLastNDays,omitempty"`
}

//	Acquire query param by id.
//
//	Returns data of the query with the specified ID.
//
//	Parameters:
//	- nId	(int64) query id
//
//	Returns:
//	- (params) object containing variables:
//	|- KLQRS_QUERY_DATA, KLPAR::ParamsPtr object, see Query storage data format
//	|- KLQRS_QUERY_TYPE, int64, value from enum QueryType
//	|- KLQRS_QUERY_GUID (string), string unique identifier of the query
func (qs *QueriesStorage) GetQuery(ctx context.Context, nId int64) (*QueryParams, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nId": %d}`, nId))
	request, err := http.NewRequest("POST", qs.client.Server+"/api/v1.0/QueriesStorage.GetQuery", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	queryParams := new(QueryParams)
	raw, err := qs.client.Do(ctx, request, &queryParams)
	return queryParams, raw, err
}

//	Acquire array of queries id given type.
//
//	Returns array of IDs of all queries of given type defined for the current user
//	(associated with the connection to the Administration Server).
//
//	Parameters:
//	- eType	(int64) queries type QueryType
//
//	Returns:
//	- (array) query ids. Collection of integers
func (qs *QueriesStorage) GetQueryIds(ctx context.Context, eType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"eType": %d}`, eType))
	request, err := http.NewRequest("POST", qs.client.Server+"/api/v1.0/QueriesStorage.GetQueryIds", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := qs.client.Do(ctx, request, nil)
	return raw, err
}

//	Query params update.
//
//	Clears current contents of the query data
//	and stores new contents specified in pParData
func (qs *QueriesStorage) UpdateQuery(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", qs.client.Server+"/api/v1.0/QueriesStorage.UpdateQuery",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := qs.client.Do(ctx, request, nil)
	return raw, err
}
