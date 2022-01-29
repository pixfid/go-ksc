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
	"fmt"
	"io"
	"io/ioutil"

	"net/http"
)

type Config struct {
	Server             string
	UserName           string
	Password           string
	Domain             string
	InternalUser       bool
	VServerName        string
	XKscSession        bool
	InsecureSkipVerify bool
	Debug              bool
}

// KscClient -------------Client------------------
type KscClient struct {
	AdfsSso                                                           *AdfsSso
	AdHosts                                                           *AdHosts
	AdmServerSettings                                                 *AdmServerSettings
	AdSecManager                                                      *AdSecManager
	AppCtrlAPI                                                        *AppCtrlApi
	AKPatches                                                         *AKPatches
	AsyncActionStateChecker                                           *AsyncActionStateChecker
	CertPoolCtrl                                                      *CertPoolCtrl
	CertPoolCtrl2                                                     *CertPoolCtrl2
	CertUtils                                                         *CertUtils
	CgwHelper                                                         *CgwHelper
	ChunkAccessor                                                     *ChunkAccessor
	CloudAccess                                                       *CloudAccess
	ConEvents                                                         *ConEvents
	DatabaseInfo                                                      *DatabaseInfo
	DataProtectionAPI                                                 *DataProtectionApi
	DpeKeyService                                                     *DpeKeyService
	EventNotificationProperties                                       *EventNotificationProperties
	EventNotificationsAPI                                             *EventNotificationsApi
	EventProcessing                                                   *EventProcessing
	EventProcessingFactory                                            *EventProcessingFactory
	ExtAud                                                            *ExtAud
	ExtTenant                                                         *ExtTenant
	FileCategorizer2                                                  *FileCategorizer2
	FilesAcceptor                                                     *FilesAcceptor
	GatewayConnection                                                 *GatewayConnection
	Gcm                                                               *Gcm
	GroupSync                                                         *GroupSync
	HostGroup                                                         *HostGroup
	HostMoveRules                                                     *HostMoveRules
	HostTagsAPI                                                       *HostTagsApi
	HostTagsRulesAPI                                                  *HostTagsRulesApi
	HostTasks                                                         *HostTasks
	HstAccessControl                                                  *HstAccessControl
	HWInvStorage                                                      *HWInvStorage
	GroupSyncIterator                                                 *GroupSyncIterator
	GroupTaskControlAPI                                               *GroupTaskControlApi
	GuiContext                                                        *GuiContext
	InventoryAPI                                                      *InventoryAPI
	InvLicenseProducts                                                *InvLicenseProducts
	IWebSrvSettings                                                   *IWebSrvSettings
	IWebUsersSrv                                                      *IWebUsersSrv
	IWebUsersSrv2                                                     *IWebUsersSrv2
	KeyService                                                        *KeyService
	KeyService2                                                       *KeyService2
	KillChain                                                         *KillChain
	KLEVerControl                                                     *KLEVerControl
	KsnInternal                                                       *KsnInternal
	LicenseInfoSync                                                   *LicenseInfoSync
	LicenseKeys                                                       *LicenseKeys
	LicensePolicy                                                     *LicensePolicy
	Limits                                                            *Limits
	ListTags                                                          *ListTags
	MfaCache                                                          *MfaCache
	MdmCertCtrlApi                                                    *MdmCertCtrlApi
	MfaCacheInner                                                     *MfaCacheInner
	MfaCacheInnerTest                                                 *MfaCacheInnerTest
	MigrationData                                                     *MigrationData
	ModulesIntegrityCheck                                             *ModulesIntegrityCheck
	Multitenancy                                                      *Multitenancy
	NagCgwHelper                                                      *NagCgwHelper
	NagGuiCalls                                                       *NagGuiCalls
	NagHstCtl                                                         *NagHstCtl
	NagNetworkListAPI                                                 *NagNetworkListApi
	NagRdu                                                            *NagRdu
	NagRemoteScreen                                                   *NagRemoteScreen
	NetUtils                                                          *NetUtils
	NlaDefinedNetworks                                                *NlaDefinedNetworks
	OAuth2                                                            *OAuth2
	OsVersion                                                         *OsVersion
	PackagesAPI                                                       *PackagesApi
	PatchParameters                                                   *PatchParameters
	PLCDevAPI                                                         *PLCDevApi
	PluginData                                                        *PluginData
	PluginDataStorage                                                 *PluginDataStorage
	Policy                                                            *Policy
	PolicyProfiles                                                    *PolicyProfiles
	ProductBackendIntegration                                         *ProductBackendIntegration
	ProductUserTokenIssuer                                            *ProductUserTokenIssuer
	QueriesStorage                                                    *QueriesStorage
	QBTNetworkListAPI                                                 *QBTNetworkListApi
	ReportManager                                                     *ReportManager
	RetrFiles                                                         *RetrFiles
	ScanDiapasons                                                     *ScanDiapasons
	SeamlessUpdatesTestAPI                                            *SeamlessUpdatesTestApi
	SecurityPolicy                                                    *SecurityPolicy
	SecurityPolicy3                                                   *SecurityPolicy3
	ServerHierarchy                                                   *ServerHierarchy
	ServerTransportSettings                                           *ServerTransportSettings
	ServiceNwcCommandProvider                                         *ServiceNwcCommandProvider
	ServiceNwcDeployment                                              *ServiceNwcDeployment
	Session                                                           *Session
	SiemExport                                                        *SiemExport
	SmsQueue                                                          *SmsQueue
	SmsSenders                                                        *SmsSenders
	SpamEvents                                                        *SpamEvents
	SrvCloud                                                          *SrvCloud
	SrvCloudStat                                                      *SrvCloudStat
	SrvIpmNewsAndStatistics                                           *SrvIpmNewsAndStatistics
	SrvRi                                                             *SrvRi
	SrvSsRevision                                                     *SrvSsRevision
	SrvView                                                           *SrvView
	SsContents                                                        *SsContents
	SsRevisionGetNames                                                *SsRevisionGetNames
	SubnetMasks                                                       *SubnetMasks
	Tasks                                                             *Tasks
	TotpGlobalSettings                                                *TotpGlobalSettings
	TotpRegistration                                                  *TotpRegistration
	TotpUserSettings                                                  *TotpUserSettings
	TrafficManager                                                    *TrafficManager
	UaControl                                                         *UaControl
	Updates                                                           *Updates
	UpdComps                                                          *UpdComps
	UserDevicesAPI                                                    *UserDevicesApi
	VapmControlAPI                                                    *VapmControlApi
	UserName, Password, Server, VServerName, XKscSessionToken, Domain string
	XKscSession, InsecureSkipVerify, InternalUser                     bool
	VServers                                                          *VServers
	VServers2                                                         *VServers2
	WolSender                                                         *WolSender
	client                                                            *http.Client
	common                                                            service
	Debug                                                             bool
}

type service struct {
	client *KscClient
}

func NewKscClient(cfg Config) *KscClient {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.InsecureSkipVerify},
		},
	}

	ksc := &KscClient{
		client:       httpClient,
		Server:       cfg.Server,
		UserName:     cfg.UserName,
		Password:     cfg.Password,
		Domain:       cfg.Domain,
		InternalUser: cfg.InternalUser,
		VServerName:  cfg.VServerName,
		XKscSession:  cfg.XKscSession,
		Debug:        cfg.Debug,
	}

	ksc.common.client = ksc
	ksc.AdfsSso = (*AdfsSso)(&ksc.common)
	ksc.DatabaseInfo = (*DatabaseInfo)(&ksc.common)
	ksc.AdHosts = (*AdHosts)(&ksc.common)
	ksc.AdmServerSettings = (*AdmServerSettings)(&ksc.common)
	ksc.AdSecManager = (*AdSecManager)(&ksc.common)
	ksc.AKPatches = (*AKPatches)(&ksc.common)
	ksc.AppCtrlAPI = (*AppCtrlApi)(&ksc.common)
	ksc.AsyncActionStateChecker = (*AsyncActionStateChecker)(&ksc.common)
	ksc.CertPoolCtrl = (*CertPoolCtrl)(&ksc.common)
	ksc.CertPoolCtrl2 = (*CertPoolCtrl2)(&ksc.common)
	ksc.CertUtils = (*CertUtils)(&ksc.common)
	ksc.CgwHelper = (*CgwHelper)(&ksc.common)
	ksc.ChunkAccessor = (*ChunkAccessor)(&ksc.common)
	ksc.CloudAccess = (*CloudAccess)(&ksc.common)
	ksc.ConEvents = (*ConEvents)(&ksc.common)
	ksc.DatabaseInfo = (*DatabaseInfo)(&ksc.common)
	ksc.DataProtectionAPI = (*DataProtectionApi)(&ksc.common)
	ksc.DpeKeyService = (*DpeKeyService)(&ksc.common)
	ksc.EventNotificationProperties = (*EventNotificationProperties)(&ksc.common)
	ksc.EventNotificationsAPI = (*EventNotificationsApi)(&ksc.common)
	ksc.EventProcessing = (*EventProcessing)(&ksc.common)
	ksc.EventProcessingFactory = (*EventProcessingFactory)(&ksc.common)
	ksc.ExtAud = (*ExtAud)(&ksc.common)
	ksc.ExtTenant = (*ExtTenant)(&ksc.common)
	ksc.FileCategorizer2 = (*FileCategorizer2)(&ksc.common)
	ksc.FilesAcceptor = (*FilesAcceptor)(&ksc.common)
	ksc.GatewayConnection = (*GatewayConnection)(&ksc.common)
	ksc.Gcm = (*Gcm)(&ksc.common)
	ksc.GroupSync = (*GroupSync)(&ksc.common)
	ksc.HostGroup = (*HostGroup)(&ksc.common)
	ksc.HostMoveRules = (*HostMoveRules)(&ksc.common)
	ksc.HostTagsAPI = (*HostTagsApi)(&ksc.common)
	ksc.HostTagsRulesAPI = (*HostTagsRulesApi)(&ksc.common)
	ksc.HostTasks = (*HostTasks)(&ksc.common)
	ksc.HstAccessControl = (*HstAccessControl)(&ksc.common)
	ksc.HWInvStorage = (*HWInvStorage)(&ksc.common)
	ksc.GroupSyncIterator = (*GroupSyncIterator)(&ksc.common)
	ksc.GroupTaskControlAPI = (*GroupTaskControlApi)(&ksc.common)
	ksc.GuiContext = (*GuiContext)(&ksc.common)
	ksc.InventoryAPI = (*InventoryAPI)(&ksc.common)
	ksc.InvLicenseProducts = (*InvLicenseProducts)(&ksc.common)
	ksc.IWebSrvSettings = (*IWebSrvSettings)(&ksc.common)
	ksc.IWebUsersSrv = (*IWebUsersSrv)(&ksc.common)
	ksc.IWebUsersSrv2 = (*IWebUsersSrv2)(&ksc.common)
	ksc.KeyService = (*KeyService)(&ksc.common)
	ksc.KeyService2 = (*KeyService2)(&ksc.common)
	ksc.KillChain = (*KillChain)(&ksc.common)
	ksc.KLEVerControl = (*KLEVerControl)(&ksc.common)
	ksc.KsnInternal = (*KsnInternal)(&ksc.common)
	ksc.LicenseInfoSync = (*LicenseInfoSync)(&ksc.common)
	ksc.LicenseKeys = (*LicenseKeys)(&ksc.common)
	ksc.LicensePolicy = (*LicensePolicy)(&ksc.common)
	ksc.Limits = (*Limits)(&ksc.common)
	ksc.ListTags = (*ListTags)(&ksc.common)
	ksc.MfaCache = (*MfaCache)(&ksc.common)
	ksc.MdmCertCtrlApi = (*MdmCertCtrlApi)(&ksc.common)
	ksc.MfaCacheInner = (*MfaCacheInner)(&ksc.common)
	ksc.MfaCacheInnerTest = (*MfaCacheInnerTest)(&ksc.common)
	ksc.MigrationData = (*MigrationData)(&ksc.common)
	ksc.ModulesIntegrityCheck = (*ModulesIntegrityCheck)(&ksc.common)
	ksc.Multitenancy = (*Multitenancy)(&ksc.common)
	ksc.NagCgwHelper = (*NagCgwHelper)(&ksc.common)
	ksc.NagGuiCalls = (*NagGuiCalls)(&ksc.common)
	ksc.NagHstCtl = (*NagHstCtl)(&ksc.common)
	ksc.NagNetworkListAPI = (*NagNetworkListApi)(&ksc.common)
	ksc.NagRdu = (*NagRdu)(&ksc.common)
	ksc.NagRemoteScreen = (*NagRemoteScreen)(&ksc.common)
	ksc.NetUtils = (*NetUtils)(&ksc.common)
	ksc.NlaDefinedNetworks = (*NlaDefinedNetworks)(&ksc.common)
	ksc.OAuth2 = (*OAuth2)(&ksc.common)
	ksc.OsVersion = (*OsVersion)(&ksc.common)
	ksc.PackagesAPI = (*PackagesApi)(&ksc.common)
	ksc.PatchParameters = (*PatchParameters)(&ksc.common)
	ksc.PLCDevAPI = (*PLCDevApi)(&ksc.common)
	ksc.PluginData = (*PluginData)(&ksc.common)
	ksc.PluginDataStorage = (*PluginDataStorage)(&ksc.common)
	ksc.Policy = (*Policy)(&ksc.common)
	ksc.PolicyProfiles = (*PolicyProfiles)(&ksc.common)
	ksc.ProductBackendIntegration = (*ProductBackendIntegration)(&ksc.common)
	ksc.ProductUserTokenIssuer = (*ProductUserTokenIssuer)(&ksc.common)
	ksc.QueriesStorage = (*QueriesStorage)(&ksc.common)
	ksc.QBTNetworkListAPI = (*QBTNetworkListApi)(&ksc.common)
	ksc.ReportManager = (*ReportManager)(&ksc.common)
	ksc.RetrFiles = (*RetrFiles)(&ksc.common)
	ksc.ScanDiapasons = (*ScanDiapasons)(&ksc.common)
	ksc.SeamlessUpdatesTestAPI = (*SeamlessUpdatesTestApi)(&ksc.common)
	ksc.SecurityPolicy = (*SecurityPolicy)(&ksc.common)
	ksc.SecurityPolicy3 = (*SecurityPolicy3)(&ksc.common)
	ksc.ServerHierarchy = (*ServerHierarchy)(&ksc.common)
	ksc.ServerTransportSettings = (*ServerTransportSettings)(&ksc.common)
	ksc.ServiceNwcCommandProvider = (*ServiceNwcCommandProvider)(&ksc.common)
	ksc.ServiceNwcDeployment = (*ServiceNwcDeployment)(&ksc.common)
	ksc.Session = (*Session)(&ksc.common)
	ksc.SiemExport = (*SiemExport)(&ksc.common)
	ksc.SmsQueue = (*SmsQueue)(&ksc.common)
	ksc.SmsSenders = (*SmsSenders)(&ksc.common)
	ksc.SpamEvents = (*SpamEvents)(&ksc.common)
	ksc.SrvCloud = (*SrvCloud)(&ksc.common)
	ksc.SrvCloudStat = (*SrvCloudStat)(&ksc.common)
	ksc.SrvIpmNewsAndStatistics = (*SrvIpmNewsAndStatistics)(&ksc.common)
	ksc.SrvRi = (*SrvRi)(&ksc.common)
	ksc.SrvSsRevision = (*SrvSsRevision)(&ksc.common)
	ksc.SrvView = (*SrvView)(&ksc.common)
	ksc.SsContents = (*SsContents)(&ksc.common)
	ksc.SsRevisionGetNames = (*SsRevisionGetNames)(&ksc.common)
	ksc.SubnetMasks = (*SubnetMasks)(&ksc.common)
	ksc.Tasks = (*Tasks)(&ksc.common)
	ksc.TotpGlobalSettings = (*TotpGlobalSettings)(&ksc.common)
	ksc.TotpRegistration = (*TotpRegistration)(&ksc.common)
	ksc.TotpUserSettings = (*TotpUserSettings)(&ksc.common)
	ksc.TrafficManager = (*TrafficManager)(&ksc.common)
	ksc.UaControl = (*UaControl)(&ksc.common)
	ksc.Updates = (*Updates)(&ksc.common)
	ksc.UpdComps = (*UpdComps)(&ksc.common)
	ksc.UserDevicesAPI = (*UserDevicesApi)(&ksc.common)
	ksc.VapmControlAPI = (*VapmControlApi)(&ksc.common)
	ksc.VServers = (*VServers)(&ksc.common)
	ksc.VServers2 = (*VServers2)(&ksc.common)
	ksc.WolSender = (*WolSender)(&ksc.common)
	return ksc
}

func (ksc *KscClient) kscAuth(ctx context.Context) error {
	request, err := http.NewRequest("POST", ksc.Server+"/api/v1.0/login", nil)
	if err != nil {
		return err
	}

	authorization := fmt.Sprintf(`KSCBasic user="%s", pass="%s", domain="%s", internal=%v`,
		ksc.UserName, ksc.Password, ksc.Domain, ksc.InternalUser)
	request.Header.Set("Authorization", authorization)
	request.Header.Set("X-KSC-VServer", ksc.VServerName)

	_, err = ksc.Request(ctx, request, nil)
	return err
}

func (ksc *KscClient) xkscSession(ctx context.Context) error {
	s, _, e := ksc.Session.StartSession(ctx)

	if s != nil {
		ksc.XKscSessionToken = s.Str
	}

	return e
}

func (ksc *KscClient) basicAuth(ctx context.Context) error {
	ksc.UserName = base64.StdEncoding.EncodeToString([]byte(ksc.UserName))
	ksc.Password = base64.StdEncoding.EncodeToString([]byte(ksc.Password))

	if len(ksc.VServerName) != 0 {
		ksc.VServerName = base64.StdEncoding.EncodeToString([]byte(ksc.VServerName))
	} else {
		ksc.VServerName = "x"
	}

	if ksc.XKscSession {
		return ksc.xkscSession(ctx)
	} else {
		return ksc.kscAuth(ctx)
	}
}

func (ksc *KscClient) kscGwAuth(ctx context.Context, token string) error {
	request, err := http.NewRequest("POST", ksc.Server+"/api/v1.0/login", nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "KSCGW "+token)

	_, err = ksc.Request(ctx, request, nil)
	return err
}

func (ksc *KscClient) kscWTAuth(ctx context.Context, kscwt string) error {
	request, err := http.NewRequest("POST", ksc.Server+"/api/v1.0/login", nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "KSCWT "+kscwt)

	_, err = ksc.Request(ctx, request, nil)
	return err
}

func (ksc *KscClient) kscTAuth(ctx context.Context, token string) error {
	request, err := http.NewRequest("POST", ksc.Server+"/api/v1.0/login", nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "KSCT "+token)

	_, err = ksc.Request(ctx, request, nil)
	return err
}

func (ksc *KscClient) Request(ctx context.Context, request *http.Request, out interface{}) (dt []byte, err error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}

	request = withContext(ctx, request)

	var response *http.Response

	if ksc.XKscSession && ksc.XKscSessionToken != "" {
		request.Header.Set("X-KSC-Session", ksc.XKscSessionToken)
	}

	request.Header.Set("User-Agent", "go-ksc")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept-Encoding", "gzip")

	response, err = ksc.client.Do(request)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	defer response.Body.Close()

	var reader io.ReadCloser

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
	default:
		reader = response.Body
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

type AuthType int

const (
	BasicAuth    AuthType = 0
	TokenAuth    AuthType = 1
	WebTokenAuth AuthType = 2
	GatewayAuth  AuthType = 3
)

func (ksc *KscClient) Login(ctx context.Context, authType AuthType, token string) error {
	switch authType {
	case BasicAuth:
		return ksc.basicAuth(ctx)
	case TokenAuth:
		return ksc.kscTAuth(ctx, token)
	case WebTokenAuth:
		return ksc.kscWTAuth(ctx, token)
	case GatewayAuth:
		return ksc.kscGwAuth(ctx, token)
	default:
		return ksc.basicAuth(ctx)
	}
}

// CheckResponse check KSC Response error
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
