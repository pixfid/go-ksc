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
	"context"

	"net/http"
)

//	SmsSenders Class Reference
//
//	Configure mobile devices as SMS senders..
//
//	List of all members.
type SmsSenders service

//	checks if there is a device allowed to send SMS
//
//	Returns:
//	- (bool) true if server has devices allowed to send SMS, false otherwise
func (ss *SmsSenders) HasAllowedSenders(ctx context.Context) (*PxgValBool, []byte, error) {
	request, err := http.NewRequest("POST", ss.client.Server+"/api/v1.0/SmsSenders.Clear", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := ss.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//TODO func (ss *SmsSenders) AllowSenders(ctx context.Context, params interface{}) ([]byte, error) {
