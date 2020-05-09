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
	"log"
	"net/http"
)

type Config struct {
	Server   string
	Username string
	Password string
}

//-------------Client------------------
type Client struct {
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
	LicenseKeys                 *LicenseKeys
	LicensePolicy               *LicensePolicy
	Limits                      *Limits
	ListTags                    *ListTags
	Multitenancy                *Multitenancy
	NagCgwHelper                *NagCgwHelper
	NagHstCtl                   *NagHstCtl
	PackagesApi                 *PackagesApi
	Policy                      *Policy
	SecurityPolicy3             *SecurityPolicy3
	ServerHierarchy             *ServerHierarchy
	Session                     *Session
	SrvView                     *SrvView
	SsContents                  *SsContents
	Tasks                       *Tasks
	TrafficManager              *TrafficManager
	Updates                     *Updates
	UserDevicesApi              *UserDevicesApi
	UserName, Password, Server  string
	VServers                    *VServers
	VServers2                   *VServers2
	WolSender                   *WolSender
	client                      *http.Client
}

func New(cfg Config) *Client {

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	client := &Client{
		client:   httpClient,
		Server:   cfg.Server,
		UserName: cfg.Username,
		Password: cfg.Password,
	}

	client.AdHosts = &AdHosts{client: client}
	client.AdmServerSettings = &AdmServerSettings{client: client}
	client.AKPatches = &AKPatches{client: client}
	client.AppCtrlApi = &AppCtrlApi{client: client}
	client.AsyncActionStateChecker = &AsyncActionStateChecker{client: client}
	client.CertPoolCtrl = &CertPoolCtrl{client: client}
	client.CertPoolCtrl2 = &CertPoolCtrl2{client: client}
	client.CgwHelper = &CgwHelper{client: client}
	client.ChunkAccessor = &ChunkAccessor{client: client}
	client.ConEvents = &ConEvents{client: client}
	client.DatabaseInfo = &DatabaseInfo{client: client}
	client.DataProtectionApi = &DataProtectionApi{client: client}
	client.EventNotificationProperties = &EventNotificationProperties{client: client}
	client.EventNotificationsApi = &EventNotificationsApi{client: client}
	client.EventProcessing = &EventProcessing{client: client}
	client.EventProcessingFactory = &EventProcessingFactory{client: client}
	client.HostGroup = &HostGroup{client: client}
	client.HostMoveRules = &HostMoveRules{client: client}
	client.HostTagsRulesApi = &HostTagsRulesApi{client: client}
	client.HostTasks = &HostTasks{client: client}
	client.HWInvStorage = &HWInvStorage{client: client}
	client.GroupSyncIterator = &GroupSyncIterator{client: client}
	client.InventoryApi = &InventoryApi{client: client}
	client.LicenseKeys = &LicenseKeys{client: client}
	client.LicensePolicy = &LicensePolicy{client: client}
	client.Limits = &Limits{client: client}
	client.ListTags = &ListTags{client: client}
	client.Multitenancy = &Multitenancy{client: client}
	client.NagCgwHelper = &NagCgwHelper{client: client}
	client.NagHstCtl = &NagHstCtl{client: client}
	client.PackagesApi = &PackagesApi{client: client}
	client.Policy = &Policy{client: client}
	client.SecurityPolicy3 = &SecurityPolicy3{client: client}
	client.ServerHierarchy = &ServerHierarchy{client: client}
	client.Session = &Session{client: client}
	client.SrvView = &SrvView{client: client}
	client.SsContents = &SsContents{client: client}
	client.Tasks = &Tasks{client: client}
	client.TrafficManager = &TrafficManager{client: client}
	client.Updates = &Updates{client: client}
	client.UserDevicesApi = &UserDevicesApi{client: client}
	client.VServers = &VServers{client: client}
	client.VServers2 = &VServers2{client: client}
	client.WolSender = &WolSender{client: client}

	return client
}

func (c *Client) KSCAuth(ctx context.Context) {
	c.UserName = base64.StdEncoding.EncodeToString([]byte(c.UserName))
	c.Password = base64.StdEncoding.EncodeToString([]byte(c.Password))

	request, err := http.NewRequest("POST", c.Server+"/api/v1.0/login", nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	request.Header.Set("Authorization", "KSCBasic user=\""+c.UserName+"\", pass=\""+c.Password+"\"")
	request.Header.Set("X-KSC-VServer", "x")
	request.Header.Set("Content-Length", "2")

	_, err = c.Do(ctx, request, nil)

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
