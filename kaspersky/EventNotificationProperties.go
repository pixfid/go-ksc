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

//DefaultSettings struct using in GetDefaultSettings
type DefaultSettings struct {
	DefaultSettingsVal *DefaultSettingsVal `json:"PxgRetVal,omitempty"`
}

//Events notification settings
type DefaultSettingsVal struct {
	//Number of days to store events in KSC server DB
	KlevpNdDaysToStoreEvent int64 `json:"KLEVP_ND_DAYS_TO_STORE_EVENT,omitempty"`

	//List of email recipients separated by comma
	KlevpNdEmail string `json:"KLEVP_ND_EMAIL,omitempty"`

	//Account name to be used for authorization on SMTP server
	KlevpNdEmailEsmtpUser string `json:"KLEVP_ND_EMAIL_ESMTP_USER,omitempty"`

	//Sender address
	KlevpNdEmailFrom string `json:"KLEVP_ND_EMAIL_FROM,omitempty"`

	//Email subject
	KlevpNdEmailSubject string `json:"KLEVP_ND_EMAIL_SUBJECT,omitempty"`

	//Event type (e.g., "GNRL_EV_VIRUS_FOUND" or "KLPRCI_TaskState")
	KlevpNdEvetnType string `json:"KLEVP_ND_EVETN_TYPE,omitempty"`

	//Template of the message to be sent as email;
	//to see the lis of available tamplate substitutions, see Events templates
	//	╔══════════╦════════╦═══════════════════╗
	//	║   Name   ║  Type  ║    Description    ║
	//	╠══════════╬════════╬═══════════════════╣
	//	║ SEVERITY ║ string ║ Event severity    ║
	//	║ COMPUTER ║ string ║ Computer name     ║
	//	║ EVENT    ║ string ║ Event type        ║
	//	║ DESCR    ║ string ║ Event description ║
	//	╚══════════╩════════╩═══════════════════╝
	KlevpNdMessageTemplate string `json:"KLEVP_ND_MESSAGE_TEMPLATE,omitempty"`

	//Obsolete parameter - must be empty or not present
	KlevpNdNetSend string `json:"KLEVP_ND_NET_SEND,omitempty"`

	//Use MX record lookup when email notification is enabled
	//(meaningful only in case if KLEVP_ND_USE_EMAIL is set to true);
	//When enabled, KLEVP_ND_SMTP_SERVER is interpreted as a domain name which is to be resolved
	KlevpNdResolveMX bool `json:"KLEVP_ND_RESOLVE_MX,omitempty"`

	//Script to be run as event notification
	KlevpNdScript string `json:"KLEVP_ND_SCRIPT,omitempty"`

	//Account name to be used for authorization on SMTP SMS gateway
	KlevpNdSMSEmailEsmtpUser string `json:"KLEVP_ND_SMS_EMAIL_ESMTP_USER,omitempty"`

	//Sender address to be used for SMTP SMS gateway
	KlevpNdSMSEmailFrom string `json:"KLEVP_ND_SMS_EMAIL_FROM,omitempty"`

	//Email subject for SMTP messages to be sent to SMTP SMS gateway
	KlevpNdSMSEmailSubject string `json:"KLEVP_ND_SMS_EMAIL_SUBJECT,omitempty"`

	//Recipient address to be used for SMTP SMS gateway
	KlevpNdSMSEmailTo string `json:"KLEVP_ND_SMS_EMAIL_TO,omitempty"`

	//Limitation on the number of SMS notifications
	KlevpNdSMSLimit int64 `json:"KLEVP_ND_SMS_LIMIT,omitempty"`

	//SMS recipients list
	KlevpNdSMSRecipients string `json:"KLEVP_ND_SMS_RECIPIENTS,omitempty"`

	//Unsupported, must be empty
	KlevpNdSMSServiceID string `json:"KLEVP_ND_SMS_SERVICE_ID,omitempty"`

	//SMTP SMS gateway server port to be used for SMS notifications
	KlevpNdSMSSMTPPort int64 `json:"KLEVP_ND_SMS_SMTP_PORT,omitempty"`

	//SMTP SMS gateway server address to be used for SMS notifications
	KlevpNdSMSSMTPServer string `json:"KLEVP_ND_SMS_SMTP_SERVER,omitempty"`

	//SMS message template;
	//to see the list of available tamplate substitutions, see Events templates
	//	╔══════════╦════════╦═══════════════════╗
	//	║   Name   ║  Type  ║    Description    ║
	//	╠══════════╬════════╬═══════════════════╣
	//	║ SEVERITY ║ string ║ Event severity    ║
	//	║ COMPUTER ║ string ║ Computer name     ║
	//	║ EVENT    ║ string ║ Event type        ║
	//	║ DESCR    ║ string ║ Event description ║
	//	╚══════════╩════════╩═══════════════════╝
	KlevpNdSMSTemplate string `json:"KLEVP_ND_SMS_TEMPLATE,omitempty"`

	//SMS notification type:
	//
	//	0 - Undefined
	//	1 - SMTP SMS gateway
	//	2 - SMS service
	KlevpNdSMSType int64 `json:"KLEVP_ND_SMS_TYPE,omitempty"`

	//SMTP server port
	KlevpNdSMTPPort int64 `json:"KLEVP_ND_SMTP_PORT,omitempty"`

	//SMTP server address to be used for email notifications
	KlevpNdSMTPServer string `json:"KLEVP_ND_SMTP_SERVER,omitempty"`

	//Store events in Kaspersky Event Log on client computer
	KlevpNdStoreAtClientLog bool `json:"KLEVP_ND_STORE_AT_CLIENT_LOG,omitempty"`

	//Obsolete parameter - must be set to false or not present
	KlevpNdStoreAtClientPres bool `json:"KLEVP_ND_STORE_AT_CLIENT_PRES,omitempty"`

	//Store events in Kaspersky Event Log on KSC server computer
	KlevpNdStoreAtServerLog bool `json:"KLEVP_ND_STORE_AT_SERVER_LOG,omitempty"`

	//Using emails for notifications; must be false for default settings
	KlevpNdUseEmail bool `json:"KLEVP_ND_USE_EMAIL,omitempty"`

	//Obsolete parameter - must be set to false or not present
	KlevpNdUseNetSend bool `json:"KLEVP_ND_USE_NET_SEND,omitempty"`

	//Notify on event by running the script. The script itself (KLEVP_ND_SCRIPT)
	//can be set in the same params, otherwise default parameters for server are used
	KlevpNdUseScript bool `json:"KLEVP_ND_USE_SCRIPT,omitempty"`

	//Using SMS for notifications; must be false for default settings
	KlevpNdUseSMS bool `json:"KLEVP_ND_USE_SMS,omitempty"`

	//Notify on events by SNMP
	KlevpNdUseSNMP bool `json:"KLEVP_ND_USE_SNMP,omitempty"`

	//Notify on events by exporting to SysLog
	KlevpNdUseSyslog bool `json:"KLEVP_ND_USE_SYSLOG,omitempty"`

	//KLEVP_ND_BODY_FILTER interface ???
}

//	Reads the default notification settings.
//
//	Reads the default notification settings, such as SMTP server properties, etc.
//
//	Returns:
//	- (params) object containing the current notification settings
//	- (see Events notification settings).
func (enp *EventNotificationProperties) GetDefaultSettings(ctx context.Context) (*DefaultSettings, []byte, error) {
	request, err := http.NewRequest("POST", enp.client.Server+"/api/v1.0/EventNotificationProperties.GetDefaultSettings", nil)
	if err != nil {
		return nil, nil, err
	}

	defaultSettings := new(DefaultSettings)
	raw, err := enp.client.Do(ctx, request, &defaultSettings)
	return defaultSettings, raw, err
}

//ENLimits struct
type ENLimits struct {
	NLimits *NLimits `json:"PxgRetVal,omitempty"`
}

//ENLimitsSettings struct
type ENLimitsParams struct {
	NLimits *NLimits `json:"pSettings,omitempty"`
}

type NLimits struct {
	KlevpMaxEventsToSendPerPeriod   int64 `json:"KLEVP_MAX_EVENTS_TO_SEND_PER_PERIOD,omitempty"`
	KlevpMaxVirusEventsForOutbreak  int64 `json:"KLEVP_MAX_VIRUS_EVENTS_FOR_OUTBREAK,omitempty"`
	KlevpMaxVirusEventsForOutbreakE int64 `json:"KLEVP_MAX_VIRUS_EVENTS_FOR_OUTBREAK_E,omitempty"`
	KlevpMaxVirusEventsForOutbreakP int64 `json:"KLEVP_MAX_VIRUS_EVENTS_FOR_OUTBREAK_P,omitempty"`
	KlevpTestPeriodToOutbreak       int64 `json:"KLEVP_TEST_PERIOD_TO_OUTBREAK,omitempty"`
	KlevpTestPeriodToOutbreakE      int64 `json:"KLEVP_TEST_PERIOD_TO_OUTBREAK_E,omitempty"`
	KlevpTestPeriodToOutbreakP      int64 `json:"KLEVP_TEST_PERIOD_TO_OUTBREAK_P,omitempty"`
	KlevpTestPeriodToSendEvents     int64 `json:"KLEVP_TEST_PERIOD_TO_SEND_EVENTS,omitempty"`
}

//	Reads the notification limits.
//
//
//	Returns:
//	- (params) object containing the current notification limits (see Events notification settings).
func (enp *EventNotificationProperties) GetNotificationLimits(ctx context.Context) (*ENLimits, []byte, error) {
	request, err := http.NewRequest("POST", enp.client.Server+"/api/v1.0/EventNotificationProperties.GetNotificationLimits", nil)
	if err != nil {
		return nil, nil, err
	}

	enLimits := new(ENLimits)
	raw, err := enp.client.Do(ctx, request, &enLimits)
	return enLimits, raw, err
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
	postData := []byte(fmt.Sprintf(`{"eType": %d,"pSettings": %v	}`, eType, pSettings))
	request, err := http.NewRequest("POST", enp.client.Server+"/api/v1.0/EventNotificationProperties.TestNotification", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := enp.client.Do(ctx, request, nil)
	return raw, err
}

//	Sets up the notification limits.
//
//	Allows to setup notification limits.
//
//Parameters:
//	- pSettings	(params) object containing the notification limits to be set (see Events notification limits).
func (enp *EventNotificationProperties) SetNotificationLimits(ctx context.Context, params ENLimitsParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", enp.client.Server+"/api/v1.0/EventNotificationProperties.SetNotificationLimits", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := enp.client.Do(ctx, request, nil)
	return raw, err
}

//	Sets up the default notification settings.
//
//	Allows to setup the default notification settings, such as SMTP server properties, etc.
//
//	Parameters:
//	- pSettings	(params) object containing the notification settings to be set (see Events notification settings).
func (enp *EventNotificationProperties) SetDefaultSettings(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", enp.client.Server+"/api/v1.0/EventNotificationProperties.SetDefaultSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := enp.client.Do(ctx, request, nil)
	return raw, err
}
