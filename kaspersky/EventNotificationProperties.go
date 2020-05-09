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
	"fmt"
	"log"
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
		log.Fatal(err.Error())
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
		log.Fatal(err.Error())
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
		log.Fatal(err.Error())
	}

	raw, err := enp.client.Do(ctx, request, nil)
	return raw, err
}

//TODO func (enp *EventNotificationProperties) SetNotificationLimits(ctx context.Context) ([]byte, error)
//TODO func (enp *EventNotificationProperties) SetDefaultSettings(ctx context.Context) ([]byte, error)
