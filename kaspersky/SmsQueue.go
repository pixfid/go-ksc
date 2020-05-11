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

//	SmsQueue Class Reference
//
//	Manage SMS message queue. More...
//
//	List of all members.
type SmsQueue service

type SQParams struct {
	WstrBody    *string  `json:"wstrBody,omitempty"`
	PRecipients []string `json:"pRecipients"`
}

//	Enqueue message into SMS queue for one or more recipients.
//
//	Parameters:
//	- wstrBody	(string) - message text
//	- pRecipients	(array) of paramString phone numbers
//
//	Returns:
//(int64) Request Id
func (sq *SmsQueue) Enqueue(ctx context.Context, params SQParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sq.client.Server+"/api/v1.0/SmsQueue.Enqueue", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sq.client.Do(ctx, request, nil)
	return raw, err
}

//	Clear sms queue.
//
//	All requests which have not been sent to any device,
//	or marked as failed will be removed.
func (sq *SmsQueue) Clear(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sq.client.Server+"/api/v1.0/SmsQueue.Clear", nil)
	if err != nil {
		return nil, err
	}

	raw, err := sq.client.Do(ctx, request, nil)
	return raw, err
}

//	SQCParams struct
type SQCParams struct {
	NRequestID  *int64   `json:"nRequestId,omitempty"`
	PRecipients []string `json:"pRecipients"`
}

//	Cancel request nRequestId to recipients pRecipients.
//
//	If request has not yet been sent to any device, or is marked as failed, it will be removed from SMS queue.
//	nRequestId must be result of one of Enqueue()
//	calls or value from c_szwRequestId field of c_szwSmsQueueSrvViewName server view.
//
//	Parameters:
//	- nRequestId	(int64) Request id to cancel. Value 0 means all requests to pRecipients will be removed.
//	- pRecipients	(array) Recipients to cancel request for.
//
//	Value null means request will be removed for all recipients
func (sq *SmsQueue) Cancel(ctx context.Context, params SQCParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sq.client.Server+"/api/v1.0/SmsQueue.Cancel", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sq.client.Do(ctx, request, nil)
	return raw, err
}
