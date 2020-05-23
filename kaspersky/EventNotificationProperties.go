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
	"fmt"
	"net/http"
)

//	EventNotificationProperties Class Reference
//
//	Notification properties.
//
//	Detailed Description
//
//	Notification properties.
//
//	Allows to setup SMTP server address, port and authentication parameters
//	to send notifications via SMTP and/or SMTP SMS gateway, etc.
//
//	List of all members.
type EventNotificationProperties service

//	Reads the default notification settings.
//
//	Reads the default notification settings, such as SMTP server properties, etc.
//
//	Returns:
//	- (params) object containing the current notification settings
//	- (see Events notification settings).
func (enp *EventNotificationProperties) GetDefaultSettings(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", enp.client.Server+"/api/v1.0/EventNotificationProperties.GetDefaultSettings", nil)
	if err != nil {
		return nil, err
	}

	raw, err := enp.client.Do(ctx, request, nil)
	return raw, err
}

//	Reads the notification limits.
//
//
//	Returns:
//	- (params) object containing the current notification limits (see Events notification settings).
func (enp *EventNotificationProperties) GetNotificationLimits(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", enp.client.Server+"/api/v1.0/EventNotificationProperties.GetNotificationLimits", nil)
	if err != nil {
		return nil, err
	}

	raw, err := enp.client.Do(ctx, request, nil)
	return raw, err
}

//	Tests the notification settings.
//
//	Allows to test the notification settings, such as SMTP server properties,
//	etc. by sending a test notification using the provided notification settings.
//
//	Parameters:
//	eType	(EventNT) type of the notification to be tested (see Events notification types).
//
//	- Undefined EventNT = 0 //Undefined
//	- EMAIL     EventNT = 1 //EMAIL
//	- SMS       EventNT = 8 //SMS
//	- pSettings	(params) object containing the notification settings to be tested (see Events notification settings).
func (enp *EventNotificationProperties) TestNotification(ctx context.Context, eType int,
	pSettings interface{}) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
		"eType": %d,
		"pSettings": %v
	}`, eType, pSettings))
	request, err := http.NewRequest("POST", enp.client.Server+"/api/v1.0/EventNotificationProperties.TestNotification", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := enp.client.Do(ctx, request, nil)
	return raw, err
}

//TODO func (enp *EventNotificationProperties) SetNotificationLimits(ctx context.Context) ([]byte, error)
//TODO func (enp *EventNotificationProperties) SetDefaultSettings(ctx context.Context) ([]byte, error)
