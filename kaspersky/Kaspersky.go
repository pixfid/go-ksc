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
	Server      string
	UserName    string
	Password    string
	VServerName string
	XKscSession bool
}

//-------------Client------------------
type Client struct {
	AdfsSso                                                   *AdfsSso
	AdHosts                                                   *AdHosts
	AdmServerSettings                                         *AdmServerSettings
	AdSecManager                                              *AdSecManager
	AppCtrlApi                                                *AppCtrlApi
	AKPatches                                                 *AKPatches
	AsyncActionStateChecker                                   *AsyncActionStateChecker
	CertPoolCtrl                                              *CertPoolCtrl
	CertPoolCtrl2                                             *CertPoolCtrl2
	CgwHelper                                                 *CgwHelper
	ChunkAccessor                                             *ChunkAccessor
	CloudAccess                                               *CloudAccess
	ConEvents                                                 *ConEvents
	DatabaseInfo                                              *DatabaseInfo
	DataProtectionApi                                         *DataProtectionApi
	DpeKeyService                                             *DpeKeyService
	EventNotificationProperties                               *EventNotificationProperties
	EventNotificationsApi                                     *EventNotificationsApi
	EventProcessing                                           *EventProcessing
	EventProcessingFactory                                    *EventProcessingFactory
	ExtAud                                                    *ExtAud
	FileCategorizer2                                          *FileCategorizer2
	FilesAcceptor                                             *FilesAcceptor
	GatewayConnection                                         *GatewayConnection
	GroupSync                                                 *GroupSync
	HostGroup                                                 *HostGroup
	HostMoveRules                                             *HostMoveRules
	HostTagsApi                                               *HostTagsApi
	HostTagsRulesApi                                          *HostTagsRulesApi
	HostTasks                                                 *HostTasks
	HstAccessControl                                          *HstAccessControl
	HWInvStorage                                              *HWInvStorage
	GroupSyncIterator                                         *GroupSyncIterator
	GroupTaskControlApi                                       *GroupTaskControlApi
	InventoryApi                                              *InventoryApi
	InvLicenseProducts                                        *InvLicenseProducts
	IWebSrvSettings                                           *IWebSrvSettings
	IWebUsersSrv                                              *IWebUsersSrv
	IWebUsersSrv2                                             *IWebUsersSrv2
	KeyService                                                *KeyService
	KeyService2                                               *KeyService2
	KillChain                                                 *KillChain
	KLEVerControl                                             *KLEVerControl
	KsnInternal                                               *KsnInternal
	LicenseInfoSync                                           *LicenseInfoSync
	LicenseKeys                                               *LicenseKeys
	LicensePolicy                                             *LicensePolicy
	Limits                                                    *Limits
	ListTags                                                  *ListTags
	MigrationData                                             *MigrationData
	Multitenancy                                              *Multitenancy
	NagCgwHelper                                              *NagCgwHelper
	NagGuiCalls                                               *NagGuiCalls
	NagHstCtl                                                 *NagHstCtl
	NagNetworkListApi                                         *NagNetworkListApi
	NagRdu                                                    *NagRdu
	NagRemoteScreen                                           *NagRemoteScreen
	NetUtils                                                  *NetUtils
	NlaDefinedNetworks                                        *NlaDefinedNetworks
	OsVersion                                                 *OsVersion
	PackagesApi                                               *PackagesApi
	PatchParameters                                           *PatchParameters
	Policy                                                    *Policy
	QueriesStorage                                            *QueriesStorage
	QBTNetworkListApi                                         *QBTNetworkListApi
	ReportManager                                             *ReportManager
	RetrFiles                                                 *RetrFiles
	ScanDiapasons                                             *ScanDiapasons
	SecurityPolicy                                            *SecurityPolicy
	SecurityPolicy3                                           *SecurityPolicy3
	ServerHierarchy                                           *ServerHierarchy
	ServerTransportSettings                                   *ServerTransportSettings
	Session                                                   *Session
	SmsQueue                                                  *SmsQueue
	SmsSenders                                                *SmsSenders
	SrvCloud                                                  *SrvCloud
	SrvSsRevision                                             *SrvSsRevision
	SrvView                                                   *SrvView
	SsContents                                                *SsContents
	SubnetMasks                                               *SubnetMasks
	Tasks                                                     *Tasks
	TrafficManager                                            *TrafficManager
	UaControl                                                 *UaControl
	Updates                                                   *Updates
	UpdComps                                                  *UpdComps
	UserDevicesApi                                            *UserDevicesApi
	VapmControlApi                                            *VapmControlApi
	UserName, Password, Server, VServerName, XKscSessionToken string
	XKscSession                                               bool
	VServers                                                  *VServers
	VServers2                                                 *VServers2
	WolSender                                                 *WolSender
	client                                                    *http.Client
	common                                                    service
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
		client:      httpClient,
		Server:      cfg.Server,
		UserName:    cfg.UserName,
		Password:    cfg.Password,
		VServerName: cfg.VServerName,
		XKscSession: cfg.XKscSession,
	}

	c.common.client = c
	c.AdfsSso = (*AdfsSso)(&c.common)
	c.DatabaseInfo = (*DatabaseInfo)(&c.common)
	c.AdHosts = (*AdHosts)(&c.common)
	c.AdmServerSettings = (*AdmServerSettings)(&c.common)
	c.AdSecManager = (*AdSecManager)(&c.common)
	c.AKPatches = (*AKPatches)(&c.common)
	c.AppCtrlApi = (*AppCtrlApi)(&c.common)
	c.AsyncActionStateChecker = (*AsyncActionStateChecker)(&c.common)
	c.CertPoolCtrl = (*CertPoolCtrl)(&c.common)
	c.CertPoolCtrl2 = (*CertPoolCtrl2)(&c.common)
	c.CgwHelper = (*CgwHelper)(&c.common)
	c.ChunkAccessor = (*ChunkAccessor)(&c.common)
	c.CloudAccess = (*CloudAccess)(&c.common)
	c.ConEvents = (*ConEvents)(&c.common)
	c.DatabaseInfo = (*DatabaseInfo)(&c.common)
	c.DataProtectionApi = (*DataProtectionApi)(&c.common)
	c.DpeKeyService = (*DpeKeyService)(&c.common)
	c.EventNotificationProperties = (*EventNotificationProperties)(&c.common)
	c.EventNotificationsApi = (*EventNotificationsApi)(&c.common)
	c.EventProcessing = (*EventProcessing)(&c.common)
	c.EventProcessingFactory = (*EventProcessingFactory)(&c.common)
	c.ExtAud = (*ExtAud)(&c.common)
	c.FileCategorizer2 = (*FileCategorizer2)(&c.common)
	c.FilesAcceptor = (*FilesAcceptor)(&c.common)
	c.GatewayConnection = (*GatewayConnection)(&c.common)
	c.GroupSync = (*GroupSync)(&c.common)
	c.HostGroup = (*HostGroup)(&c.common)
	c.HostMoveRules = (*HostMoveRules)(&c.common)
	c.HostTagsApi = (*HostTagsApi)(&c.common)
	c.HostTagsRulesApi = (*HostTagsRulesApi)(&c.common)
	c.HostTasks = (*HostTasks)(&c.common)
	c.HstAccessControl = (*HstAccessControl)(&c.common)
	c.HWInvStorage = (*HWInvStorage)(&c.common)
	c.GroupSyncIterator = (*GroupSyncIterator)(&c.common)
	c.GroupTaskControlApi = (*GroupTaskControlApi)(&c.common)
	c.InventoryApi = (*InventoryApi)(&c.common)
	c.InvLicenseProducts = (*InvLicenseProducts)(&c.common)
	c.IWebSrvSettings = (*IWebSrvSettings)(&c.common)
	c.IWebUsersSrv = (*IWebUsersSrv)(&c.common)
	c.IWebUsersSrv2 = (*IWebUsersSrv2)(&c.common)
	c.KeyService = (*KeyService)(&c.common)
	c.KeyService2 = (*KeyService2)(&c.common)
	c.KillChain = (*KillChain)(&c.common)
	c.KLEVerControl = (*KLEVerControl)(&c.common)
	c.KsnInternal = (*KsnInternal)(&c.common)
	c.LicenseInfoSync = (*LicenseInfoSync)(&c.common)
	c.LicenseKeys = (*LicenseKeys)(&c.common)
	c.LicensePolicy = (*LicensePolicy)(&c.common)
	c.Limits = (*Limits)(&c.common)
	c.ListTags = (*ListTags)(&c.common)
	c.MigrationData = (*MigrationData)(&c.common)
	c.Multitenancy = (*Multitenancy)(&c.common)
	c.NagCgwHelper = (*NagCgwHelper)(&c.common)
	c.NagGuiCalls = (*NagGuiCalls)(&c.common)
	c.NagHstCtl = (*NagHstCtl)(&c.common)
	c.NagNetworkListApi = (*NagNetworkListApi)(&c.common)
	c.NagRdu = (*NagRdu)(&c.common)
	c.NagRemoteScreen = (*NagRemoteScreen)(&c.common)
	c.NetUtils = (*NetUtils)(&c.common)
	c.NlaDefinedNetworks = (*NlaDefinedNetworks)(&c.common)
	c.OsVersion = (*OsVersion)(&c.common)
	c.PackagesApi = (*PackagesApi)(&c.common)
	c.PatchParameters = (*PatchParameters)(&c.common)
	c.Policy = (*Policy)(&c.common)
	c.QueriesStorage = (*QueriesStorage)(&c.common)
	c.QBTNetworkListApi = (*QBTNetworkListApi)(&c.common)
	c.ReportManager = (*ReportManager)(&c.common)
	c.RetrFiles = (*RetrFiles)(&c.common)
	c.ScanDiapasons = (*ScanDiapasons)(&c.common)
	c.SecurityPolicy = (*SecurityPolicy)(&c.common)
	c.SecurityPolicy3 = (*SecurityPolicy3)(&c.common)
	c.ServerHierarchy = (*ServerHierarchy)(&c.common)
	c.ServerTransportSettings = (*ServerTransportSettings)(&c.common)
	c.Session = (*Session)(&c.common)
	c.SmsQueue = (*SmsQueue)(&c.common)
	c.SmsSenders = (*SmsSenders)(&c.common)
	c.SrvCloud = (*SrvCloud)(&c.common)
	c.SrvSsRevision = (*SrvSsRevision)(&c.common)
	c.SrvView = (*SrvView)(&c.common)
	c.SsContents = (*SsContents)(&c.common)
	c.SubnetMasks = (*SubnetMasks)(&c.common)
	c.Tasks = (*Tasks)(&c.common)
	c.TrafficManager = (*TrafficManager)(&c.common)
	c.UaControl = (*UaControl)(&c.common)
	c.Updates = (*Updates)(&c.common)
	c.UpdComps = (*UpdComps)(&c.common)
	c.UserDevicesApi = (*UserDevicesApi)(&c.common)
	c.VapmControlApi = (*VapmControlApi)(&c.common)
	c.VServers = (*VServers)(&c.common)
	c.VServers2 = (*VServers2)(&c.common)
	c.WolSender = (*WolSender)(&c.common)
	return c
}

func (c *Client) basicAuth(ctx context.Context) error {

	request, err := http.NewRequest("POST", c.Server+"/api/v1.0/login", nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "KSCBasic user=\""+c.UserName+"\", pass=\""+c.Password+"\"")
	request.Header.Set("X-KSC-VServer", c.VServerName)

	_, err = c.Do(ctx, request, nil)
	return err
}

func (c *Client) xkscSession(ctx context.Context) error {
	s, _, e := c.Session.StartSession(ctx)

	if s != nil {
		c.XKscSessionToken = s.Str
	}

	return e
}

func (c *Client) KSCAuth(ctx context.Context) error {

	c.UserName = base64.StdEncoding.EncodeToString([]byte(c.UserName))
	c.Password = base64.StdEncoding.EncodeToString([]byte(c.Password))

	if len(c.VServerName) != 0 {
		c.VServerName = base64.StdEncoding.EncodeToString([]byte(c.VServerName))
	} else {
		c.VServerName = "x"
	}

	if c.XKscSession {
		return c.xkscSession(ctx)
	} else {
		return c.basicAuth(ctx)
	}
}

func (c *Client) KSCGWAuth(ctx context.Context, kscgw string) error {
	request, err := http.NewRequest("POST", c.Server+"/api/v1.0/login", nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "KSCGW "+kscgw)

	_, err = c.Do(ctx, request, nil)
	return err
}

func (c *Client) KSCWTAuth(ctx context.Context, kscwt string) error {
	request, err := http.NewRequest("POST", c.Server+"/api/v1.0/login", nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "KSCWT "+kscwt)

	_, err = c.Do(ctx, request, nil)
	return err
}

func (c *Client) KSCTAuth(ctx context.Context, ksct string) error {
	request, err := http.NewRequest("POST", c.Server+"/api/v1.0/login", nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "KSCT "+ksct)

	_, err = c.Do(ctx, request, nil)
	return err
}

func (c *Client) Do(ctx context.Context, req *http.Request, out interface{}) (dt []byte, err error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}

	req = withContext(ctx, req)

	var resp *http.Response

	if c.XKscSession && c.XKscSessionToken != "" {
		req.Header.Set("X-KSC-Session", c.XKscSessionToken)
	}

	req.Header.Set("User-Agent", "go-ksc")
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

	body, err := ioutil.ReadAll(reader)

	if err != nil {
		return body, err
	}

	err = CheckResponse(&body)

	if err != nil {
		return body, err
	}

	if out != nil {
		decErr := json.Unmarshal(body, out)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}

	return body, err
}

//CheckResponse check KSC Response error
func CheckResponse(body *[]byte) (err error) {

	pre := new(PxgRetError)

	decErr := json.Unmarshal(*body, &pre)

	if decErr == io.EOF {
		decErr = nil // ignore EOF errors caused by empty response body
	}
	if decErr != nil {
		err = decErr
	}

	if pre.Error != nil {
		err = pre.Error
	}

	return err
}
