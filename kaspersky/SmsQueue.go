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

// SmsQueue service to manage SMS message queue.
type SmsQueue service

type SQParams struct {
	WstrBody    string   `json:"wstrBody,omitempty"`
	PRecipients []string `json:"pRecipients"`
}

// Enqueue message into SMS queue for one or more recipients.
func (sq *SmsQueue) Enqueue(ctx context.Context, params SQParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sq.client.Server+"/api/v1.0/SmsQueue.Enqueue", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sq.client.Request(ctx, request, nil)
	return raw, err
}

// Clear sms queue. All requests which have not been sent to any device, or marked as failed will be removed.
func (sq *SmsQueue) Clear(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sq.client.Server+"/api/v1.0/SmsQueue.Clear", nil)
	if err != nil {
		return nil, err
	}

	raw, err := sq.client.Request(ctx, request, nil)
	return raw, err
}

// SQCParams struct
type SQCParams struct {
	NRequestID  int64    `json:"nRequestId,omitempty"`
	PRecipients []string `json:"pRecipients"`
}

// Cancel request nRequestId to recipients pRecipients.
func (sq *SmsQueue) Cancel(ctx context.Context, params SQCParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sq.client.Server+"/api/v1.0/SmsQueue.Cancel", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sq.client.Request(ctx, request, nil)
	return raw, err
}
