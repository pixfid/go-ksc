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
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"

	"net/http"
)

type Config struct {
	Server   string
	Username string
	Password string
}

//-------------Client------------------
type Client struct {
	AdfsSso                     *AdfsSso
	AdHosts                     *AdHosts
	AdmServerSettings           *AdmServerSettings
	AppCtrlApi                  *AppCtrlApi
	AKPatches                   *AKPatches
	AsyncActionStateChecker     *AsyncActionStateChecker
	CertPoolCtrl                *CertPoolCtrl
	CertPoolCtrl2               *CertPoolCtrl2
	CgwHelper                   *CgwHelper
	ChunkAccessor               *ChunkAccessor
	ConEvents                   *ConEvents
	DatabaseInfo                *DatabaseInfo
	DataProtectionApi           *DataProtectionApi
	DpeKeyService               *DpeKeyService
	EventNotificationProperties *EventNotificationProperties
	EventNotificationsApi       *EventNotificationsApi
	EventProcessing             *EventProcessing
	EventProcessingFactory      *EventProcessingFactory
	HostGroup                   *HostGroup
	HostMoveRules               *HostMoveRules
	HostTagsRulesApi            *HostTagsRulesApi
	HostTasks                   *HostTasks
	HWInvStorage                *HWInvStorage
	GroupSyncIterator           *GroupSyncIterator
	InventoryApi                *InventoryApi
	InvLicenseProducts          *InvLicenseProducts
	KsnInternal                 *KsnInternal
	LicenseKeys                 *LicenseKeys
	LicensePolicy               *LicensePolicy
	Limits                      *Limits
	ListTags                    *ListTags
	Multitenancy                *Multitenancy
	NagCgwHelper                *NagCgwHelper
	NagGuiCalls                 *NagGuiCalls
	NagHstCtl                   *NagHstCtl
	NagRdu                      *NagRdu
	PackagesApi                 *PackagesApi
	Policy                      *Policy
	ReportManager               *ReportManager
	ScanDiapasons               *ScanDiapasons
	SecurityPolicy3             *SecurityPolicy3
	ServerHierarchy             *ServerHierarchy
	ServerTransportSettings     *ServerTransportSettings
	Session                     *Session
	SmsQueue                    *SmsQueue
	SmsSenders                  *SmsSenders
	SrvSsRevision               *SrvSsRevision
	SrvView                     *SrvView
	SsContents                  *SsContents
	Tasks                       *Tasks
	TrafficManager              *TrafficManager
	UaControl                   *UaControl
	Updates                     *Updates
	UserDevicesApi              *UserDevicesApi
	UserName, Password, Server  string
	VServers                    *VServers
	VServers2                   *VServers2
	WolSender                   *WolSender
	client                      *http.Client
	common                      service
}

type service struct {
	client *Client
}

func New(cfg Config) *Client {

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	c := &Client{
		client:   httpClient,
		Server:   cfg.Server,
		UserName: cfg.Username,
		Password: cfg.Password,
	}

	c.common.client = c
	c.AdfsSso = (*AdfsSso)(&c.common)
	c.DatabaseInfo = (*DatabaseInfo)(&c.common)
	c.AdHosts = (*AdHosts)(&c.common)
	c.AdmServerSettings = (*AdmServerSettings)(&c.common)
	c.AKPatches = (*AKPatches)(&c.common)
	c.AppCtrlApi = (*AppCtrlApi)(&c.common)
	c.AsyncActionStateChecker = (*AsyncActionStateChecker)(&c.common)
	c.CertPoolCtrl = (*CertPoolCtrl)(&c.common)
	c.CertPoolCtrl2 = (*CertPoolCtrl2)(&c.common)
	c.CgwHelper = (*CgwHelper)(&c.common)
	c.ChunkAccessor = (*ChunkAccessor)(&c.common)
	c.ConEvents = (*ConEvents)(&c.common)
	c.DatabaseInfo = (*DatabaseInfo)(&c.common)
	c.DataProtectionApi = (*DataProtectionApi)(&c.common)
	c.DpeKeyService = (*DpeKeyService)(&c.common)
	c.EventNotificationProperties = (*EventNotificationProperties)(&c.common)
	c.EventNotificationsApi = (*EventNotificationsApi)(&c.common)
	c.EventProcessing = (*EventProcessing)(&c.common)
	c.EventProcessingFactory = (*EventProcessingFactory)(&c.common)
	c.HostGroup = (*HostGroup)(&c.common)
	c.HostMoveRules = (*HostMoveRules)(&c.common)
	c.HostTagsRulesApi = (*HostTagsRulesApi)(&c.common)
	c.HostTasks = (*HostTasks)(&c.common)
	c.HWInvStorage = (*HWInvStorage)(&c.common)
	c.GroupSyncIterator = (*GroupSyncIterator)(&c.common)
	c.InventoryApi = (*InventoryApi)(&c.common)
	c.InvLicenseProducts = (*InvLicenseProducts)(&c.common)
	c.KsnInternal = (*KsnInternal)(&c.common)
	c.LicenseKeys = (*LicenseKeys)(&c.common)
	c.LicensePolicy = (*LicensePolicy)(&c.common)
	c.Limits = (*Limits)(&c.common)
	c.ListTags = (*ListTags)(&c.common)
	c.Multitenancy = (*Multitenancy)(&c.common)
	c.NagCgwHelper = (*NagCgwHelper)(&c.common)
	c.NagGuiCalls = (*NagGuiCalls)(&c.common)
	c.NagHstCtl = (*NagHstCtl)(&c.common)
	c.NagRdu = (*NagRdu)(&c.common)
	c.PackagesApi = (*PackagesApi)(&c.common)
	c.Policy = (*Policy)(&c.common)
	c.ReportManager = (*ReportManager)(&c.common)
	c.ScanDiapasons = (*ScanDiapasons)(&c.common)
	c.SecurityPolicy3 = (*SecurityPolicy3)(&c.common)
	c.ServerHierarchy = (*ServerHierarchy)(&c.common)
	c.ServerTransportSettings = (*ServerTransportSettings)(&c.common)
	c.Session = (*Session)(&c.common)
	c.SmsQueue = (*SmsQueue)(&c.common)
	c.SmsSenders = (*SmsSenders)(&c.common)
	c.SrvSsRevision = (*SrvSsRevision)(&c.common)
	c.SrvView = (*SrvView)(&c.common)
	c.SsContents = (*SsContents)(&c.common)
	c.Tasks = (*Tasks)(&c.common)
	c.TrafficManager = (*TrafficManager)(&c.common)
	c.UaControl = (*UaControl)(&c.common)
	c.Updates = (*Updates)(&c.common)
	c.UserDevicesApi = (*UserDevicesApi)(&c.common)
	c.VServers = (*VServers)(&c.common)
	c.VServers2 = (*VServers2)(&c.common)
	c.WolSender = (*WolSender)(&c.common)
	return c
}

func (c *Client) KSCAuth(ctx context.Context) error {
	c.UserName = base64.StdEncoding.EncodeToString([]byte(c.UserName))
	c.Password = base64.StdEncoding.EncodeToString([]byte(c.Password))

	request, err := http.NewRequest("POST", c.Server+"/api/v1.0/login", nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "KSCBasic user=\""+c.UserName+"\", pass=\""+c.Password+"\"")
	request.Header.Set("X-KSC-VServer", "x")
	request.Header.Set("Content-Length", "2")

	_, err = c.Do(ctx, request, nil)
	return err
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (dt []byte, err error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}

	req = withContext(ctx, req)

	var resp *http.Response

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err = c.client.Do(req)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	defer resp.Body.Close()

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
	default:
		reader = resp.Body
	}

	body, _ := ioutil.ReadAll(reader)

	if v != nil {
		decErr := json.Unmarshal(body, v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}

	return body, err
}
