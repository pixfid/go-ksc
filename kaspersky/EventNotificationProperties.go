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

// EventNotificationProperties service to working with Notification properties.
// Allows to setup SMTP server address, port and authentication parameters to send notifications via SMTP and/or SMTP SMS gateway, etc.
type EventNotificationProperties service

// DefaultSettings struct using in GetDefaultSettings
type DefaultSettings struct {
	DefaultSettingsVal *DefaultSettingsVal `json:"PxgRetVal,omitempty"`
}

// Events notification settings
type DefaultSettingsVal struct {
	// KlevpNdDaysToStoreEvent Number of days to store events in KSC server DB
	KlevpNdDaysToStoreEvent *int64 `json:"KLEVP_ND_DAYS_TO_STORE_EVENT,omitempty"`

	// KlevpNdEmail List of email recipients separated by comma
	KlevpNdEmail *string `json:"KLEVP_ND_EMAIL,omitempty"`

	// KlevpNdEmailEsmtpUser Account name to be used for authorization on SMTP server
	KlevpNdEmailEsmtpUser *string `json:"KLEVP_ND_EMAIL_ESMTP_USER,omitempty"`

	// KlevpNdEmailFrom Sender address
	KlevpNdEmailFrom *string `json:"KLEVP_ND_EMAIL_FROM,omitempty"`

	// KlevpNdEmailSubject Email subject
	KlevpNdEmailSubject *string `json:"KLEVP_ND_EMAIL_SUBJECT,omitempty"`

	// KlevpNdEvetnType Event type (e.g., "GNRL_EV_VIRUS_FOUND" or "KLPRCI_TaskState")
	KlevpNdEvetnType *string `json:"KLEVP_ND_EVETN_TYPE,omitempty"`

	// KlevpNdMessageTemplate Template of the message to be sent as email;
	// to see the lis of available tamplate substitutions, see Events templates
	//	╔══════════╦════════╦═══════════════════╗
	//	║   Name   ║  Type  ║    Description    ║
	//	╠══════════╬════════╬═══════════════════╣
	//	║ SEVERITY ║ string ║ Event severity    ║
	//	║ COMPUTER ║ string ║ Computer name     ║
	//	║ EVENT    ║ string ║ Event type        ║
	//	║ DESCR    ║ string ║ Event description ║
	//	╚══════════╩════════╩═══════════════════╝
	KlevpNdMessageTemplate *string `json:"KLEVP_ND_MESSAGE_TEMPLATE,omitempty"`

	// KlevpNdNetSend Obsolete parameter - must be empty or not present
	KlevpNdNetSend *string `json:"KLEVP_ND_NET_SEND,omitempty"`

	// KlevpNdResolveMX Use MX record lookup when email notification is enabled
	// (meaningful only in case if KLEVP_ND_USE_EMAIL is set to true);
	// When enabled, KLEVP_ND_SMTP_SERVER is interpreted as a domain name which is to be resolved
	KlevpNdResolveMX *bool `json:"KLEVP_ND_RESOLVE_MX,omitempty"`

	// KlevpNdScript Script to be run as event notification
	KlevpNdScript *string `json:"KLEVP_ND_SCRIPT,omitempty"`

	// KlevpNdSMSEmailEsmtpUser Account name to be used for authorization on SMTP SMS gateway
	KlevpNdSMSEmailEsmtpUser *string `json:"KLEVP_ND_SMS_EMAIL_ESMTP_USER,omitempty"`

	// KlevpNdSMSEmailFrom Sender address to be used for SMTP SMS gateway
	KlevpNdSMSEmailFrom *string `json:"KLEVP_ND_SMS_EMAIL_FROM,omitempty"`

	// KlevpNdSMSEmailSubject Email subject for SMTP messages to be sent to SMTP SMS gateway
	KlevpNdSMSEmailSubject *string `json:"KLEVP_ND_SMS_EMAIL_SUBJECT,omitempty"`

	// KlevpNdSMSEmailTo Recipient address to be used for SMTP SMS gateway
	KlevpNdSMSEmailTo *string `json:"KLEVP_ND_SMS_EMAIL_TO,omitempty"`

	// KlevpNdSMSLimit Limitation on the number of SMS notifications
	KlevpNdSMSLimit *int64 `json:"KLEVP_ND_SMS_LIMIT,omitempty"`

	// KlevpNdSMSRecipients SMS recipients list
	KlevpNdSMSRecipients *string `json:"KLEVP_ND_SMS_RECIPIENTS,omitempty"`

	// KlevpNdSMSServiceID Unsupported, must be empty
	KlevpNdSMSServiceID *string `json:"KLEVP_ND_SMS_SERVICE_ID,omitempty"`

	// KlevpNdSMSSMTPPort SMTP SMS gateway server port to be used for SMS notifications
	KlevpNdSMSSMTPPort *int64 `json:"KLEVP_ND_SMS_SMTP_PORT,omitempty"`

	// KlevpNdSMSSMTPServer SMTP SMS gateway server address to be used for SMS notifications
	KlevpNdSMSSMTPServer *string `json:"KLEVP_ND_SMS_SMTP_SERVER,omitempty"`

	// KlevpNdSMSTemplate SMS message template;
	//to see the list of available tamplate substitutions, see Events templates
	//	╔══════════╦════════╦═══════════════════╗
	//	║   Name   ║  Type  ║    Description    ║
	//	╠══════════╬════════╬═══════════════════╣
	//	║ SEVERITY ║ string ║ Event severity    ║
	//	║ COMPUTER ║ string ║ Computer name     ║
	//	║ EVENT    ║ string ║ Event type        ║
	//	║ DESCR    ║ string ║ Event description ║
	//	╚══════════╩════════╩═══════════════════╝
	KlevpNdSMSTemplate *string `json:"KLEVP_ND_SMS_TEMPLATE,omitempty"`

	// KlevpNdSMSType SMS notification type:
	//
	//	0 - Undefined
	//	1 - SMTP SMS gateway
	//	2 - SMS service
	KlevpNdSMSType *int64 `json:"KLEVP_ND_SMS_TYPE,omitempty"`

	// KlevpNdSMTPPort SMTP server port
	KlevpNdSMTPPort *int64 `json:"KLEVP_ND_SMTP_PORT,omitempty"`

	// KlevpNdSMTPServer SMTP server address to be used for email notifications
	KlevpNdSMTPServer *string `json:"KLEVP_ND_SMTP_SERVER,omitempty"`

	// KlevpNdStoreAtClientLog Store events in Kaspersky Event Log on client computer
	KlevpNdStoreAtClientLog *bool `json:"KLEVP_ND_STORE_AT_CLIENT_LOG,omitempty"`

	// KlevpNdStoreAtClientPres Obsolete parameter - must be set to false or not present
	KlevpNdStoreAtClientPres *bool `json:"KLEVP_ND_STORE_AT_CLIENT_PRES,omitempty"`

	// KlevpNdStoreAtServerLog Store events in Kaspersky Event Log on KSC server computer
	KlevpNdStoreAtServerLog *bool `json:"KLEVP_ND_STORE_AT_SERVER_LOG,omitempty"`

	// KlevpNdUseEmail Using emails for notifications; must be false for default settings
	KlevpNdUseEmail *bool `json:"KLEVP_ND_USE_EMAIL,omitempty"`

	// KlevpNdUseNetSend Obsolete parameter - must be set to false or not present
	KlevpNdUseNetSend *bool `json:"KLEVP_ND_USE_NET_SEND,omitempty"`

	// KlevpNdUseScript Notify on event by running the script. The script itself (KLEVP_ND_SCRIPT)
	// can be set in the same params, otherwise default parameters for server are used
	KlevpNdUseScript *bool `json:"KLEVP_ND_USE_SCRIPT,omitempty"`

	// KlevpNdUseSMS Using SMS for notifications; must be false for default settings
	KlevpNdUseSMS *bool `json:"KLEVP_ND_USE_SMS,omitempty"`

	// KlevpNdUseSNMP Notify on events by SNMP
	KlevpNdUseSNMP *bool `json:"KLEVP_ND_USE_SNMP,omitempty"`

	// KlevpNdUseSyslog Notify on events by exporting to SysLog
	KlevpNdUseSyslog *bool `json:"KLEVP_ND_USE_SYSLOG,omitempty"`

	//KLEVP_ND_BODY_FILTER interface ???
}

// GetDefaultSettings Reads the default notification settings. Reads the default notification settings, such as SMTP server properties, etc.
func (enp *EventNotificationProperties) GetDefaultSettings(ctx context.Context) (*DefaultSettings, []byte, error) {
	request, err := http.NewRequest("POST", enp.client.Server+"/api/v1.0/EventNotificationProperties.GetDefaultSettings", nil)
	if err != nil {
		return nil, nil, err
	}

	defaultSettings := new(DefaultSettings)
	raw, err := enp.client.Do(ctx, request, &defaultSettings)
	return defaultSettings, raw, err
}

// ENLimits struct
type ENLimits struct {
	NLimits *NLimits `json:"PxgRetVal,omitempty"`
}

// ENLimitsSettings struct
type ENLimitsParams struct {
	NLimits *NLimits `json:"pSettings,omitempty"`
}

// NLimits struct
type NLimits struct {
	KlevpMaxEventsToSendPerPeriod   *int64 `json:"KLEVP_MAX_EVENTS_TO_SEND_PER_PERIOD,omitempty"`
	KlevpMaxVirusEventsForOutbreak  *int64 `json:"KLEVP_MAX_VIRUS_EVENTS_FOR_OUTBREAK,omitempty"`
	KlevpMaxVirusEventsForOutbreakE *int64 `json:"KLEVP_MAX_VIRUS_EVENTS_FOR_OUTBREAK_E,omitempty"`
	KlevpMaxVirusEventsForOutbreakP *int64 `json:"KLEVP_MAX_VIRUS_EVENTS_FOR_OUTBREAK_P,omitempty"`
	KlevpTestPeriodToOutbreak       *int64 `json:"KLEVP_TEST_PERIOD_TO_OUTBREAK,omitempty"`
	KlevpTestPeriodToOutbreakE      *int64 `json:"KLEVP_TEST_PERIOD_TO_OUTBREAK_E,omitempty"`
	KlevpTestPeriodToOutbreakP      *int64 `json:"KLEVP_TEST_PERIOD_TO_OUTBREAK_P,omitempty"`
	KlevpTestPeriodToSendEvents     *int64 `json:"KLEVP_TEST_PERIOD_TO_SEND_EVENTS,omitempty"`
}

// GetNotificationLimits Reads the notification limits.
func (enp *EventNotificationProperties) GetNotificationLimits(ctx context.Context) (*ENLimits, []byte, error) {
	request, err := http.NewRequest("POST", enp.client.Server+"/api/v1.0/EventNotificationProperties.GetNotificationLimits", nil)
	if err != nil {
		return nil, nil, err
	}

	enLimits := new(ENLimits)
	raw, err := enp.client.Do(ctx, request, &enLimits)
	return enLimits, raw, err
}

// TestNotification Tests the notification settings.
// Allows to test the notification settings, such as SMTP server properties, etc.
// by sending a test notification using the provided notification settings.
func (enp *EventNotificationProperties) TestNotification(ctx context.Context, eType int, pSettings interface{}) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"eType": %d,"pSettings": %v	}`, eType, pSettings))
	request, err := http.NewRequest("POST", enp.client.Server+"/api/v1.0/EventNotificationProperties.TestNotification", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := enp.client.Do(ctx, request, nil)
	return raw, err
}

// SetNotificationLimits Sets up the notification limits. Allows to setup notification limits.
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

// SetDefaultSettings Sets up the default notification settings.
// Allows to setup the default notification settings, such as SMTP server properties, etc.
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
