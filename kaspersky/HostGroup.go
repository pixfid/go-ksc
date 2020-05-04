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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HostGroup struct {
	client *Client
}

type HGParams struct {
	WstrFilter        string             `json:"wstrFilter"`
	VecFieldsToReturn []string           `json:"vecFieldsToReturn"`
	VecFieldsToOrder  []VecFieldsToOrder `json:"vecFieldsToOrder"`
	PParams           PParams            `json:"pParams"`
	LMaxLifeTime      int64              `json:"lMaxLifeTime"`
}

type PParams struct {
	KlsrvhSlaveRecDepth    int64 `json:"KLSRVH_SLAVE_REC_DEPTH"`
	KlgrpFindFromCurVsOnly bool  `json:"KLGRP_FIND_FROM_CUR_VS_ONLY"`
}

//Find groups by filter string.
//
//Finds groups that satisfy conditions from filter pParams, and creates a server-side collection of found groups.
//Search is performed over the hierarchy
//
//Parameters:
//	- pParams data.HGParams
//
//Session to the Administration Server has been closed.
//ChunkAccessor.Release has been called.
//
//Returns:
//	- (int64) number of found groups
func (hg *HostGroup) FindGroups(ctx context.Context, pParams HGParams) (*Accessor, []byte, error) {
	postData, _ := json.Marshal(pParams)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindGroups", bytes.NewBuffer(postData))

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

//Find host by filter string.
//
//Finds hosts that satisfy conditions from filter string wstrFilter, and creates a server-side collection of found hosts. Search is performed over the hierarchy
//
//Parameters:
//	- wstrFilter	(string) filter string, contains a condition over host attributes, see also Search filter syntax.
//	- vecFieldsToReturn	([]string) array of host attribute names to return. See List of host attributes for attribute names
//	- vecFieldsToOrder	([]string) array of containers each of them containing two attributes :
//	- "Name" (string) name of attribute used for sorting
//	- "Asc" (string) ascending if true descending otherwise
//	- pParams	(params) extra options. Possible attributes are listed below (see details in Extra search attributes for hosts and administration groups):
//	- KLSRVH_SLAVE_REC_DEPTH
//	- KLGRP_FIND_FROM_CUR_VS_ONLY
//	- lMaxLifeTime	(int64) max result-set lifetime in seconds, not more than 7200
//returns strAccessor	(string, error) result-set ID, identifier of the server-side ordered collection of found hosts. The result-set is destroyed and associated memory is freed in following cases:
//Passed lMaxLifeTime seconds after last access to the result-set (by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//Session to the Administration Server has been closed.
//ChunkAccessor.Release has been called.
//
//Returns:
//(int64) number of found hosts
func (hg *HostGroup) FindHosts(ctx context.Context, hgParams HGParams) (*Accessor, []byte, error) {

	postData, _ := json.Marshal(hgParams)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHosts", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

func (hg *HostGroup) UpdateHost(ctx context.Context, v interface{}) (*Accessor, []byte, error) {
	data, _ := json.Marshal(v)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.UpdateHost", bytes.NewBuffer(data))

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

type UHGParams struct {
	StrFilter       string   `json:"strFilter"`
	PFieldsToReturn []string `json:"pFieldsToReturn"`
	PParams         PParams  `json:"pParams"`
	LMaxLifeTime    int64    `json:"lMaxLifeTime"`
}

//Finds existing users.
//
//Finds users that satisfy conditions from filter string strFilter.
//
//Parameters:
//	- strFilter	(string) filter string, see Search filter syntax
//	- pFieldsToReturn	(array) array of user's attribute names to return. See List of user's attributes
//	- pFieldsToOrder	(array) array of containers each of them containing two attributes:
//		"Name" of type String, name of attribute used for sorting
//		"Asc" of type Boolean, ascending if true descending otherwise
//	- lMaxLifeTime	(int64) max lifetime of accessor (sec)
//	- [out]	strAccessor	(string) result-set ID, identifier of the server-side ordered collection of found users.
//	The result-set is destroyed and associated memory is freed in following cases:
//	Passed lMaxLifeTime seconds after last access to the result-set (by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//	Session to the Administration Server has been closed.
//	ChunkAccessor.Release has been called.
//Returns:
//	- (int64) number of records found
func (hg *HostGroup) FindUsers(ctx context.Context, param UHGParams) (*Accessor, []byte, error) {
	postData, _ := json.Marshal(param)

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindUsers", bytes.NewBuffer(postData))

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

//Get result of FindHostsAsync operation.
//
//Gets result of asynchronous operation HostGroup.FindHostsAsync
//
//Parameters:
//	- strRequestId	(string) identity of asynchronous operation
//	- [out]	strAccessor	(string) result-set ID, identifier of the server-side ordered collection of found hosts. The result-set is destroyed and associated memory is freed in following cases:
//Passed lMaxLifeTime seconds after last access to the result-set (by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//Session to the Administration Server has been closed.
//ChunkAccessor.Release has been called.
//	- [out]	pFailedSlavesInfo	(params) information about slave servers the search for which failed due to various reasons, contains array KLGRP_FAILED_SLAVES_PARAMS of params which have attributes:
//	- KLSRVH_SRV_ID - Slave server id (paramInt)
//	- KLSRVH_SRV_DN - Slave server display name (paramString)
//Returns:
//	- (int64) number of found hosts
func (hg *HostGroup) FindHostsAsyncGetAccessor(ctx context.Context, strRequestId string) (*AsyncAccessor,
	[]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId" : "%s" }`, strRequestId))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHostsAsyncGetAccessor", bytes.NewBuffer(postData))

	asyncAccessor := new(AsyncAccessor)
	raw, err := hg.client.Do(ctx, request, &asyncAccessor)
	return asyncAccessor, raw, err
}

//Find host asynchronously by filter string.
//
//Finds hosts asynchronously that satisfy conditions from filter string wstrFilter,
//and creates a server-side collection of found hosts. Search is performed over the hierarchy
//
//Parameters:
//	- wstrFilter	(string) filter string, contains a condition over host attributes, see also Search filter syntax.
//	- vecFieldsToReturn	([]string) array of host attribute names to return.
//	See List of host attributes for attribute names
//	- vecFieldsToOrder	([]string) array of containers each of them containing two attributes :
//		"Name" (paramString) name of attribute used for sorting
//		"Asc" (paramBool) ascending if true descending otherwise
//	- pParams	(params) extra options. Possible attributes are listed below
//(see details in Extra search attributes for hosts and administration groups):
//	- KLSRVH_SLAVE_REC_DEPTH
//	- KLGRP_FIND_FROM_CUR_VS_ONLY
//	- lMaxLifeTime	(int64) max result-set lifetime in seconds, not more than 7200
//return:
//	- data.RequestID	(string) identity of asynchronous operation,
//
//	to get status use AsyncActionStateChecker.CheckActionState, lStateCode "1" means OK and "0" means fail
//
//	to get accessor id call HostGroup.FindHostsAsyncGetAccessor
//
//	to cancel operation call HostGroup.FindHostsAsyncCancel
func (hg *HostGroup) FindHostsAsync(ctx context.Context, pParam HGParams) (*RequestID, []byte, error) {
	postData, _ := json.Marshal(pParam)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHostsAsync", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	requestID := new(RequestID)
	raw, err := hg.client.Do(ctx, request, &requestID)
	return requestID, raw, err
}

//Cancel FindHostsAsync operation.
//
//Cancels asynchronous operation HostGroup.FindHostsAsync
//
//Parameters:
//	- strRequestId	(string) identity of asynchronous operation
func (hg *HostGroup) FindHostsAsyncCancel(ctx context.Context, strRequestId string) {

	postData := []byte(fmt.Sprintf(`
	{
	"strRequestId": "%s"
	}`, strRequestId))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHostsAsyncCancel", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = hg.client.Do(ctx, request, nil)
}

type AVHostProduct struct {
	AVProduct AVProduct `json:"PxgRetVal"`
}

type AVProduct struct {
	Klng *Klng  `json:"1103"`
	Kesl *KesHP `json:"kesl,omitempty"`
	Kes  *KesHP `json:"KES,omitempty"`
	Ess  *EssHP `json:"ESS,omitempty"`
}

type EssHP struct {
	Value ESSHPValue `json:"value"`
}

type ESSHPValue struct {
	The2100 The2100 `json:"2.1.0.0"`
}

type The2100 struct {
	Value The2100Value `json:"value"`
}

type The2100Value struct {
	BaseDate                     InstallTime `json:"BaseDate,omitempty"`
	BaseInstallDate              InstallTime `json:"BaseInstallDate,omitempty"`
	BaseRecords                  int64       `json:"BaseRecords"`
	ConnDisplayName              string      `json:"ConnDisplayName"`
	ConnProdVersion              string      `json:"ConnProdVersion"`
	ConnectorComponentName       string      `json:"ConnectorComponentName"`
	ConnectorFlags               int64       `json:"ConnectorFlags"`
	ConnectorInstanceID          string      `json:"ConnectorInstanceId"`
	ConnectorVersion             string      `json:"ConnectorVersion"`
	CustomName                   string      `json:"CustomName"`
	DataFolder                   string      `json:"DataFolder,omitempty"`
	DisplayName                  string      `json:"DisplayName"`
	FileName                     string      `json:"FileName"`
	FilePath                     string      `json:"FilePath"`
	InstallTime                  InstallTime `json:"InstallTime"`
	InstallationID               string      `json:"InstallationId"`
	KlconnappinstMustLoadOutside bool        `json:"KLCONNAPPINST_MUST_LOAD_OUTSIDE"`
	LastUpdateTime               InstallTime `json:"LastUpdateTime,omitempty"`
	ModuleType                   int64       `json:"ModuleType"`
	ProdVersion                  string      `json:"ProdVersion"`
	Tasks                        []string    `json:"Tasks"`
	TasksComplemented            []string    `json:"TasksComplemented"`
}

type InstallTime struct {
	Value string `json:"value"`
}

type KesHP struct {
	Value map[string]The2100 `json:"value"`
}

type Klng struct {
	Value The1103HP_Value `json:"value"`
}

type The1103HP_Value struct {
	The1000 The1000 `json:"1.0.0.0"`
}

type The1000 struct {
	Value The1000_Value `json:"value"`
}

type The1000_Value struct {
	BaseRecords         int64       `json:"BaseRecords"`
	DataFolder          string      `json:"DataFolder"`
	DisplayName         string      `json:"DisplayName"`
	FileName            string      `json:"FileName"`
	FilePath            string      `json:"FilePath"`
	Folder              string      `json:"Folder"`
	InstallTime         InstallTime `json:"InstallTime"`
	InstallationID      string      `json:"InstallationId"`
	KlnagInstallationID string      `json:"KLNAG_INSTALLATION_ID"`
	LastUpdateTime      InstallTime `json:"LastUpdateTime"`
	LOCID               string      `json:"LocID"`
	ModuleType          int64       `json:"ModuleType"`
	ProdVersion         string      `json:"ProdVersion"`
	BaseFolder          string      `json:"BaseFolder,omitempty"`
	UpdateRoot          bool        `json:"UpdateRoot,omitempty"`
}

//Return information about installed products on the host.
//
//Parameters:
//	- strHostName	(string) host name (GUID-like identifier)
//
//Returns:
//	- (data.AVHostProduct) contains containers with names of products that contain containers with verions of product, i.e:
//	- <Product> (paramParams)
//	- <Version> (paramParams)
//	- InstallTime
//	- InstallationId
//	- DisplayName
//	- BaseRecords
//	- ConnDisplayName
//	- ConnProdVersion
//	- ConnectorComponentName
func (hg *HostGroup) GetHostProducts(ctx context.Context, strHostName string) (*AVHostProduct, []byte, error) {

	postData := []byte(fmt.Sprintf(`
	{
	"strHostName": "%s"
	}`, strHostName))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetHostProducts", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	aVHostProduct := new(AVHostProduct)
	raw, err := hg.client.Do(ctx, request, &aVHostProduct)
	return aVHostProduct, raw, err
}

//InstanceStatistics struct
type InstanceStatistics struct {
	PxgRetVal *PxgRetVal `json:"PxgRetVal,omitempty"`
}

type PxgRetVal struct {
	KlasyncactCounters              *KlasyncactCounters              `json:"KLASYNCACT_COUNTERS,omitempty"`
	KlftsStGetChunkRequestsRejected *KlftsStGetChunkRequestsRejected `json:"KLFTS_ST_GET_CHUNK_REQUESTS_REJECTED,omitempty"`
	KlftsStGetChunkRequestsTotal    *KlftsStGetChunkRequestsRejected `json:"KLFTS_ST_GET_CHUNK_REQUESTS_TOTAL,omitempty"`
	KlftsStTransmittedSize          *KlftsStGetChunkRequestsRejected `json:"KLFTS_ST_TRANSMITTED_SIZE,omitempty"`
	KlnlstCounters                  *KlnlstCounters                  `json:"KLNLST_COUNTERS,omitempty"`
	KlsrvStAdScanPercent            *int64                           `json:"KLSRV_ST_AD_SCAN_PERCENT,omitempty"`
	KlsrvStAllConsCnt               *int64                           `json:"KLSRV_ST_ALL_CONS_CNT,omitempty"`
	KlsrvStConEvents                *KlsrvStConEvents                `json:"KLSRV_ST_CON_EVENTS,omitempty"`
	KlsrvStCPUKernel                *KlftsStGetChunkRequestsRejected `json:"KLSRV_ST_CPU_KERNEL,omitempty"`
	KlsrvStCPUUser                  *KlftsStGetChunkRequestsRejected `json:"KLSRV_ST_CPU_USER,omitempty"`
	KlsrvStCtlngtConsCnt            *int64                           `json:"KLSRV_ST_CTLNGT_CONS_CNT,omitempty"`
	KlsrvStDpnsScanPercent          *int64                           `json:"KLSRV_ST_DPNS_SCAN_PERCENT,omitempty"`
	KlsrvStFullScanPercent          *int64                           `json:"KLSRV_ST_FULL_SCAN_PERCENT,omitempty"`
	KlsrvStGUICallsCnt              *int64                           `json:"KLSRV_ST_GUI_CALLS_CNT,omitempty"`
	KlsrvStLastAdscanTime           *KlsrvSt                         `json:"KLSRV_ST_LAST_ADSCAN_TIME,omitempty"`
	KlsrvStLastDpnsscanTime         *KlsrvSt                         `json:"KLSRV_ST_LAST_DPNSSCAN_TIME,omitempty"`
	KlsrvStLastFastnetscanTime      *KlsrvSt                         `json:"KLSRV_ST_LAST_FASTNETSCAN_TIME,omitempty"`
	KlsrvStLastFullnetscanTime      *KlsrvSt                         `json:"KLSRV_ST_LAST_FULLNETSCAN_TIME,omitempty"`
	KlsrvStNagConsCnt               *int64                           `json:"KLSRV_ST_NAG_CONS_CNT,omitempty"`
	KlsrvStNetworkDomainScanned     *string                          `json:"KLSRV_ST_NETWORK_DOMAIN_SCANNED,omitempty"`
	KlsrvStNetworkScanned           *bool                            `json:"KLSRV_ST_NETWORK_SCANNED,omitempty"`
	KlsrvStNetworkScanPercent       *int64                           `json:"KLSRV_ST_NETWORK_SCAN_PERCENT,omitempty"`
	KlsrvStPingCnt                  *int64                           `json:"KLSRV_ST_PING_CNT,omitempty"`
	KlsrvStPingErrorCnt             *int64                           `json:"KLSRV_ST_PING_ERROR_CNT,omitempty"`
	KlsrvStPingJnCnt                *int64                           `json:"KLSRV_ST_PING_JN_CNT,omitempty"`
	KlsrvStPingRejectedCnt          *int64                           `json:"KLSRV_ST_PING_REJECTED_CNT,omitempty"`
	KlsrvStSyncCnt                  *int64                           `json:"KLSRV_ST_SYNC_CNT,omitempty"`
	KlsrvStSyncJnCnt                *int64                           `json:"KLSRV_ST_SYNC_JN_CNT,omitempty"`
	KlsrvStSyncQueueSize            *int64                           `json:"KLSRV_ST_SYNC_QUEUE_SIZE,omitempty"`
	KlsrvStSyncRealCnt              *int64                           `json:"KLSRV_ST_SYNC_REAL_CNT,omitempty"`
	KlsrvStSyncSucCnt               *int64                           `json:"KLSRV_ST_SYNC_SUC_CNT,omitempty"`
	KlsrvStTotalHostsCount          *int64                           `json:"KLSRV_ST_TOTAL_HOSTS_COUNT,omitempty"`
	KlsrvStVirtServersDetails       []KlsrvStVirtServersDetail       `json:"KLSRV_ST_VIRT_SERVERS_DETAILS"`
	KlsrvStVirtServerCount          *int64                           `json:"KLSRV_ST_VIRT_SERVER_COUNT,omitempty"`
	KltrStAcceptsFailed             *KlftsStGetChunkRequestsRejected `json:"KLTR_ST_ACCEPTS_FAILED,omitempty"`
	KltrStAcceptsTotal              *KlftsStGetChunkRequestsRejected `json:"KLTR_ST_ACCEPTS_TOTAL,omitempty"`
	KltrStReceivedBytes             *KlftsStGetChunkRequestsRejected `json:"KLTR_ST_RECEIVED_BYTES,omitempty"`
	KltrStSentBytes                 *KlftsStGetChunkRequestsRejected `json:"KLTR_ST_SENT_BYTES,omitempty"`
}

type KlasyncactCounters struct {
	Type  *string                  `json:"type,omitempty"`
	Value *KLASYNCACTCOUNTERSValue `json:"value,omitempty"`
}

type KLASYNCACTCOUNTERSValue struct {
	KlasyncactActionCount                   *int64                           `json:"KLASYNCACT_ACTION_COUNT,omitempty"`
	KlasyncactCheckLimitViolationCount      *int64                           `json:"KLASYNCACT_CHECK_LIMIT_VIOLATION_COUNT,omitempty"`
	KlasyncactConnectionCount               *int64                           `json:"KLASYNCACT_CONNECTION_COUNT,omitempty"`
	KlasyncactDurationMsec                  *KlftsStGetChunkRequestsRejected `json:"KLASYNCACT_DURATION_MSEC,omitempty"`
	KlasyncactFinalizedActionCount          *int64                           `json:"KLASYNCACT_FINALIZED_ACTION_COUNT,omitempty"`
	KlasyncactFinalizedActionViolationCount *int64                           `json:"KLASYNCACT_FINALIZED_ACTION_VIOLATION_COUNT,omitempty"`
}

type KlftsStGetChunkRequestsRejected struct {
	Type  *string  `json:"type,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

type KlnlstCounters struct {
	Type  *string              `json:"type,omitempty"`
	Value *KLNLSTCOUNTERSValue `json:"value,omitempty"`
}

type KLNLSTCOUNTERSValue struct {
	KlnlstCountersListsArray []PurpleKLNLSTCOUNTERSLISTSARRAY `json:"KLNLST_COUNTERS_LISTS_ARRAY"`
	KlnlstCountersTransqueue *KlnlstCountersTransqueue        `json:"KLNLST_COUNTERS_TRANSQUEUE,omitempty"`
}

type PurpleKLNLSTCOUNTERSLISTSARRAY struct {
	Type  *string `json:"type,omitempty"`
	Value *PValue `json:"value,omitempty"`
}

type PValue struct {
	KlnlstCountersDeleteallCnt                    *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_DELETEALL_CNT,omitempty"`
	KlnlstCountersDeleteditemCnt                  *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_DELETEDITEM_CNT,omitempty"`
	KlnlstCountersListName                        *string                          `json:"KLNLST_COUNTERS_LIST_NAME,omitempty"`
	KlnlstCountersUpdateditemCnt                  *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_UPDATEDITEM_CNT,omitempty"`
	KlnlstCountersUpdateditemOpCnt                *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_UPDATEDITEM_OP_CNT,omitempty"`
	KlnlstCountersdeleteditemMeandurationMsec     *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERSDELETEDITEM_MEANDURATION_MSEC,omitempty"`
	KlnlstCountersDeleteditemLongopCnt            *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_DELETEDITEM_LONGOP_CNT,omitempty"`
	KlnlstCountersDeleteditemNlstLongopCnt        *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_DELETEDITEM_NLST_LONGOP_CNT,omitempty"`
	KlnlstCountersDeleteditemNlstMeandurationMsec *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_DELETEDITEM_NLST_MEANDURATION_MSEC,omitempty"`
	KlnlstCountersTransactionLongopCnt            *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_TRANSACTION_LONGOP_CNT,omitempty"`
	KlnlstCountersTransactionMeandurationMsec     *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_TRANSACTION_MEANDURATION_MSEC,omitempty"`
	KlnlstCountersUpdateditemLongopCnt            *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_UPDATEDITEM_LONGOP_CNT,omitempty"`
	KlnlstCountersUpdateditemMeandurationMsec     *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_UPDATEDITEM_MEANDURATION_MSEC,omitempty"`
	KlnlstCountersUpdateditemNlstLongopCnt        *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_UPDATEDITEM_NLST_LONGOP_CNT,omitempty"`
	KlnlstCountersUpdateditemNlstMeandurationMsec *KlftsStGetChunkRequestsRejected `json:"KLNLST_COUNTERS_UPDATEDITEM_NLST_MEANDURATION_MSEC,omitempty"`
}

type KlnlstCountersTransqueue struct {
	Value *KLNLSTCOUNTERSTRANSQUEUEValue `json:"value,omitempty"`
}

type KLNLSTCOUNTERSTRANSQUEUEValue struct {
	KlnlstCountersListsArray         []FluffyKLNLSTCOUNTERSLISTSARRAY `json:"KLNLST_COUNTERS_LISTS_ARRAY"`
	KlnlstCountersTransqueueTransCnt *int64                           `json:"KLNLST_COUNTERS_TRANSQUEUE_TRANS_CNT,omitempty"`
}

type FluffyKLNLSTCOUNTERSLISTSARRAY struct {
	Value *FValue `json:"value,omitempty"`
}

type FValue struct {
	KlnlstCountersListName                 *string `json:"KLNLST_COUNTERS_LIST_NAME,omitempty"`
	KlnlstCountersTransqueueListLimited    *bool   `json:"KLNLST_COUNTERS_TRANSQUEUE_LIST_LIMITED,omitempty"`
	KlnlstCountersTransqueueListPending    *bool   `json:"KLNLST_COUNTERS_TRANSQUEUE_LIST_PENDING,omitempty"`
	KlnlstCountersTransqueueListPrelimited *bool   `json:"KLNLST_COUNTERS_TRANSQUEUE_LIST_PRELIMITED,omitempty"`
	KlnlstCountersTransqueueListTransCnt   *int64  `json:"KLNLST_COUNTERS_TRANSQUEUE_LIST_TRANS_CNT,omitempty"`
}

type KlsrvStConEvents struct {
	Value *KLSRVSTCONEVENTSValue `json:"value,omitempty"`
}

type KLSRVSTCONEVENTSValue struct {
	KlsrvStConEvLostEventsCount  *int64 `json:"KLSRV_ST_CON_EV_LOST_EVENTS_COUNT,omitempty"`
	KlsrvStConEvSessionCount     *int64 `json:"KLSRV_ST_CON_EV_SESSION_COUNT,omitempty"`
	KlsrvStConEvSubsCount        *int64 `json:"KLSRV_ST_CON_EV_SUBS_COUNT,omitempty"`
	KlsrvStConEvUnavailRetrCount *int64 `json:"KLSRV_ST_CON_EV_UNAVAIL_RETR_COUNT,omitempty"`
	KlsrvStConEvUnsubsRetrCount  *int64 `json:"KLSRV_ST_CON_EV_UNSUBS_RETR_COUNT,omitempty"`
}

type KlsrvSt struct {
	Value *string `json:"value,omitempty"`
}

type KlsrvStVirtServersDetail struct {
	Value *KLSRVSTVIRTSERVERSDETAILValue `json:"value,omitempty"`
}

type KLSRVSTVIRTSERVERSDETAILValue struct {
	KlsrvStVirtServerID           *KlftsStGetChunkRequestsRejected `json:"KLSRV_ST_VIRT_SERVER_ID,omitempty"`
	KlsrvStVirtServerName         *string                          `json:"KLSRV_ST_VIRT_SERVER_NAME,omitempty"`
	KlsrvStVserverCOMLicExists    *bool                            `json:"KLSRV_ST_VSERVER_COM_LIC_EXISTS,omitempty"`
	KlsrvStVserverHostCount       *int64                           `json:"KLSRV_ST_VSERVER_HOST_COUNT,omitempty"`
	KlsrvStVserverIosmdmDevCount  *int64                           `json:"KLSRV_ST_VSERVER_IOSMDM_DEV_COUNT,omitempty"`
	KlsrvStVserverKesmobDevCount  *int64                           `json:"KLSRV_ST_VSERVER_KESMOB_DEV_COUNT,omitempty"`
	KlsrvStVserverLastCOMLicExpir *KlsrvSt                         `json:"KLSRV_ST_VSERVER_LAST_COM_LIC_EXPIR,omitempty"`
	KlsrvStVserverRealHostCount   *int64                           `json:"KLSRV_ST_VSERVER_REAL_HOST_COUNT,omitempty"`
	KlsrvStVserverTrialLicExists  *bool                            `json:"KLSRV_ST_VSERVER_TRIAL_LIC_EXISTS,omitempty"`
}

//Acquire Server statistics info.
//
//Parameters:
//	-	 vecFilterFields	(array) Array of filtered attributes
//	- "KLSRV_ST_ALL_CONS_CNT"	Number of all connections and number of all nAgent Version connections
//	- "KLSRV_ST_CTLNGT_CONS_CNT"	Number of controlled nAgent Version connections
//	- "KLSRV_ST_NETWORK_DOMAIN_SCANNED"	Currently scanned domain name
//	- "KLSRV_ST_VIRT_SERVER_COUNT"	Virtual servers count
//	- "KLSRV_ST_TOTAL_HOSTS_COUNT"	Total active hosts count
//	- "KLSRV_ST_VIRT_SERVERS_DETAILS"	Array of active hosts count on virtual server
//	- "KLSRV_ST_CON_EVENTS"	Container with ConEvents statistics
//
//Returns:
//filtered statistic
//
//Remark:
//-	not working on KSC 10
func (hg *HostGroup) GetInstanceStatistics(ctx context.Context) (*InstanceStatistics, []byte, error) {

	postData := []byte(fmt.Sprintf(`
	{
		"vecFilterFields": 
		[
			"KLSRV_ST_ALL_CONS_CNT",
			"KLSRV_ST_NAG_CONS_CNT",
			"KLSRV_ST_CTLNGT_CONS_CNT",
			"KLSRV_ST_PING_CNT",
			"KLSRV_ST_PING_REJECTED_CNT",
			"KLSRV_ST_PING_ERROR_CNT",
			"KLSRV_ST_PING_JN_CNT",
			"KLSRV_ST_SYNC_CNT",
			"KLSRV_ST_SYNC_REAL_CNT",
			"KLSRV_ST_SYNC_SUC_CNT",
			"KLSRV_ST_SYNC_JN_CNT",
			"KLSRV_ST_NETWORK_SCANNED",
			"KLSRV_ST_LAST_FULLNETSCAN_TIME",
			"KLSRV_ST_LAST_FASTNETSCAN_TIME",
			"KLSRV_ST_LAST_DPNSSCAN_TIME",
			"KLSRV_ST_LAST_ADSCAN_TIME",
			"KLSRV_ST_FULL_SCAN_PERCENT",
			"KLSRV_ST_NETWORK_SCAN_PERCENT",
			"KLSRV_ST_DPNS_SCAN_PERCENT",
			"KLSRV_ST_AD_SCAN_PERCENT",
			"KLSRV_ST_NETWORK_DOMAIN_SCANNED",
			"KLSRV_ST_SYNC_QUEUE_SIZE",
			"KLNLST_COUNTERS",
			"KLASYNCACT_COUNTERS",
			"KLSRV_ST_CPU_USER",
			"KLSRV_ST_CPU_KERNEL",
			"KLSRV_ST_GUI_CALLS_CNT",
			"KLSRV_ST_VIRT_SERVER_COUNT",
			"KLSRV_ST_TOTAL_HOSTS_COUNT",
			"KLSRV_ST_VIRT_SERVERS_DETAILS",
			"KLTR_ST_ACCEPTS_TOTAL",
			"KLTR_ST_ACCEPTS_FAILED",
			"KLTR_ST_SENT_BYTES",
			"KLTR_ST_RECEIVED_BYTES",
			"KLFTS_ST_GET_CHUNK_REQUESTS_TOTAL",
			"KLFTS_ST_GET_CHUNK_REQUESTS_REJECTED",
			"KLFTS_ST_TRANSMITTED_SIZE",
			"KLSRV_ST_CON_EVENTS"
		]}`))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetInstanceStatistics", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	instanceStatistics := new(InstanceStatistics)
	raw, err := hg.client.Do(ctx, request, &instanceStatistics)

	return instanceStatistics, raw, err
}

//RunTimeInfo struct
type RunTimeInfo struct {
	PxgRetVal RTInfo `json:"PxgRetVal"`
}

type RTInfo struct {
	KladmsrvIfWaikInstalled bool   `json:"KLADMSRV_IF_WAIK_INSTALLED"`
	KladmsrvSaasBlocked     bool   `json:"KLADMSRV_SAAS_BLOCKED"`
	KladmsrvSaasOveruse     bool   `json:"KLADMSRV_SAAS_OVERUSE"`
	KladmsrvSssID           string `json:"KLADMSRV_SSS_ID"`
	KladmsrvSssPort         int64  `json:"KLADMSRV_SSS_PORT"`
	KladmsrvVsLicdisabled   bool   `json:"KLADMSRV_VS_LICDISABLED"`
}

// Return server run-time info.
//
//Parameters:
//	- pValues	(array) string array with names of requested values, possible values:
//	- KLADMSRV_SSS_PORT - server port (paramInt)
//	- KLADMSRV_SSS_ID - server id (paramString)
//	- KLADMSRV_VS_LICDISABLED - licensing for the VS is disabled (paramBool)
//	- KLADMSRV_SAAS_BLOCKED - adding new virtual servers is blocked due to expired/absent/blacklisted license (paramBool)
//	- KLADMSRV_SAAS_EXPIRED_DAYS_TO_WORK - adding new virtual servers will be blocked in c_szwIfSaasExpiredDaysToWork days (paramInt)
//	- KLADMSRV_SAAS_OVERUSE - number of VS created is more specified in the license (paramBool)
//	- KLADMSRV_IF_WAIK_INSTALLED - true if WAIK is installed (paramBool)
//
//Returns:
//	- (params) requsted values
func (hg *HostGroup) GetRunTimeInfo(ctx context.Context, pValues []string) (*RunTimeInfo, []byte, error) {
	vFtr, _ := json.Marshal(pValues)
	postData := []byte(fmt.Sprintf(`
	{
		"pValues": %s
	}`, string(vFtr)))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetRunTimeInfo", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	runTimeInfo := new(RunTimeInfo)
	raw, err := hg.client.Do(ctx, request, &runTimeInfo)
	return runTimeInfo, raw, err
}

//HostInfoParams struct
type HostInfoParams struct {
	StrHostName    string   `json:"strHostName"`
	PFields2Return []string `json:"pFields2Return"`
}

func (hg *HostGroup) GetHostInfo(ctx context.Context, hip HostInfoParams, hi interface{}) ([]byte, error) {

	postData, _ := json.Marshal(hip)

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetHostInfo", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hg.client.Do(ctx, request, &hi)
	return raw, err
}

type PxgHostGroupDomains struct {
	PxgRetVal []PxgRetValHGD `json:"PxgRetVal"`
}

type PxgRetValHGD struct {
	Type  string   `json:"type"`
	Value ValueHGD `json:"value"`
}

type ValueHGD struct {
	KlhstWksWindomain     string                `json:"KLHST_WKS_WINDOMAIN"`
	KlhstWksWindomainType KlhstWksWindomainType `json:"KLHST_WKS_WINDOMAIN_TYPE"`
}

type KlhstWksWindomainType int64

const (
	WindowsDomain    KlhstWksWindomainType = 0
	WindowsWorkGroup KlhstWksWindomainType = 1
)

//List of Windows domain in the network.
//
//Returns:
//	- (array) array of domains, each item is container which contains following attributes:
//	- KLHST_WKS_WINDOMAIN (paramString) domain name
//	- KLHST_WKS_WINDOMAIN_TYPE (paramInt) domain type:
//
//Type:
//	- 0 Windows NT domain
//	- 1 Windows work group
func (hg *HostGroup) GetDomains(ctx context.Context) (*PxgHostGroupDomains, []byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetDomains", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	pxgHostGroupDomains := new(PxgHostGroupDomains)
	raw, err := hg.client.Do(ctx, request, &pxgHostGroupDomains)
	return pxgHostGroupDomains, raw, err
}

//Id of predefined root group "Managed computers".
//
//Returns:
//	- (data.PxgValInt) id of predefined root group "Managed computers"
func (hg *HostGroup) GroupIdGroups(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GroupIdGroups", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//Id of predefined group "Master server".
//
//Returns:
//	- (data.PxgValInt) id of predefined group "Master server" ("Super")
func (hg *HostGroup) GroupIdSuper(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GroupIdSuper", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//Id of predefined group "Unassigned computers".
//
//Returns:
//	- (data.PxgValInt) id of predefined group "Unassigned computers"
func (hg *HostGroup) GroupIdUnassigned(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GroupIdUnassigned", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//Create new administration group.
//
//Creates new group with the specified attributes and returns its Id. If such group already exists returns Id of existing group.
//
//Parameters:
//	- pInfo	(params) container with group attributes. May contain following attributes (see List of group attributes):
//	- "name"
//	- "parentId"
//
//Returns:
//	- (int64) id of created group (or of existing one)
func (hg *HostGroup) AddGroup(ctx context.Context, name string, parentId int) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
    	"pInfo": {
        	"name": "%s",
        	"parentId": %d
    	}
	}`, name, parentId))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddGroup", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//StaticInfoParams struct
type StaticInfoParams struct {
	PValues []string `json:"pValues"`
}

//StaticsInfo struct
type StaticsInfo struct {
	StaticInfo StaticInfo `json:"PxgRetVal"`
}

type StaticInfo struct {
	KladmsrvB2BCloudMode         bool             `json:"KLADMSRV_B2B_CLOUD_MODE"`
	KladmsrvEvEvSize             KladmsrvEvEvSize `json:"KLADMSRV_EV_EV_SIZE"`
	KladmsrvForceSyncSupported   bool             `json:"KLADMSRV_FORCE_SYNC_SUPPORTED"`
	KladmsrvGrpRoot              int64            `json:"KLADMSRV_GRP_ROOT"`
	KladmsrvGrpSuper             int64            `json:"KLADMSRV_GRP_SUPER"`
	KladmsrvGrpUnassigned        int64            `json:"KLADMSRV_GRP_UNASSIGNED"`
	KladmsrvIsVirtual            bool             `json:"KLADMSRV_IS_VIRTUAL"`
	KladmsrvMaintenanceSupported bool             `json:"KLADMSRV_MAINTENANCE_SUPPORTED"`
	KladmsrvNacIsBeingUsed       bool             `json:"KLADMSRV_NAC_IS_BEING_USED"`
	KladmsrvNagentRunning        bool             `json:"KLADMSRV_NAGENT_RUNNING"`
	KladmsrvNeedUNCPath          bool             `json:"KLADMSRV_NEED_UNC_PATH"`
	KladmsrvPcloudMode           bool             `json:"KLADMSRV_PCLOUD_MODE"`
	KladmsrvProductFullVersion   string           `json:"KLADMSRV_PRODUCT_FULL_VERSION"`
	KladmsrvProductName          string           `json:"KLADMSRV_PRODUCT_NAME"`
	KladmsrvProductVersion       string           `json:"KLADMSRV_PRODUCT_VERSION"`
	KladmsrvServerHostname       string           `json:"KLADMSRV_SERVER_HOSTNAME"`
	KladmsrvServerVersionID      int64            `json:"KLADMSRV_SERVER_VERSION_ID"`
	KladmsrvSplPPCEnabled        bool             `json:"KLADMSRV_SPL_PPC_ENABLED"`
	KladmsrvUserid               int64            `json:"KLADMSRV_USERID"`
	KladmsrvVsid                 int64            `json:"KLADMSRV_VSID"`
	KladmsrvVsuid                string           `json:"KLADMSRV_VSUID"`
	KlsrvNetsize                 int64            `json:"KLSRV_NETSIZE"`
}

type KladmsrvEvEvSize struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

//Return server static info.
//
//Parameters:
//	- pValues	(array) string array with names of requested values, possible values are listed below.
//	- KLADMSRV_SERVER_CERT - server certificate (paramBinary)
//	- KLADMSRV_SERVER_KEY - server key (paramBinary)
//	- InstancePort - instance port (paramInt)
//	- KLADMSRV_SERVER_ADDRESSSES - array of server addresses that can be used by clients to connect to the administration server (paramArray|paramString)
//	- KLADMSRV_SERVER_UNDER_SYSTEM_ACCOUNT - Name of account used by Administration Server (paramString)
//	- KLADMSRV_OLA_ROOTCER_ACTUAL - Array of allowed OLA certificates, paramArray, each entry is a certificate as 'DER-encoded binary X.509' of type paramBinary
//	- KLADMSRV_OLA_ROOTCER_REVOKED - Array of disallowed (revoked) OLA certificates, paramArray, each entry is a certificate as 'DER-encoded binary X.509' of type paramBinary
//
//If pValues is NULL then described below values will be returned.
//
//	- KLADMSRV_IS_VIRTUAL - true if server is virtual (paramBool)
//	- KLADMSRV_VSID - VS id (paramInt)
//	- KLADMSRV_GRP_ROOT - id of group groups (paramInt)
//	- KLADMSRV_GRP_UNASSIGNED - id of group unassigned (paramInt)
//	- KLADMSRV_GRP_SUPER - id of group super (paramInt)
//	- KLADMSRV_SERVER_VERSION_ID - server version id (paramInt)
//	- KLADMSRV_B2B_CLOUD_MODE - if the server installed in the 'B2B Cloud' mode (paramBool)
//	- KLADMSRV_PCLOUD_MODE - if public cloud support mode is turned on (paramBool)
//	- KLADMSRV_PRODUCT_FULL_VERSION - server product full version (paramString)
//	- KLADMSRV_SERVER_HOSTNAME - server host name (paramString)
//	- KLADMSRV_PRODUCT_VERSION - server product version (paramString)
//	- KLADMSRV_PRODUCT_NAME - server product name (paramString)
//	- KLADMSRV_FORCE_SYNC_SUPPORTED - force sync supported (paramBool)
//	- KLADMSRV_MAINTENANCE_SUPPORTED - if maintenance task must is supported (paramBool)
//	- KLADMSRV_NAgent Version_RUNNING - true if nAgent Version is running (paramBool)
//	- KLADMSRV_NAC_IS_BEING_USED - true if NAC is used (paramBool)
//	- KLADMSRV_SPL_PPC_ENABLED - If password policy compliance for SPL users is enabled (paramBool)
//	- KLSRV_NETSIZE - network size (paramInt)
//	- KLADMSRV_USERID - id of the user account; NULL if OS user account is used (paramInt)
//	- KLADMSRV_NEED_UNC_PATH - if UNC path must be specified into backup task settings (paramBool)
//	- KLADMSRV_EV_EV_SIZE - average size of a single event, Kb (paramDouble)
//
//Returns:
//	- (params) requsted values
func (hg *HostGroup) GetStaticInfo(ctx context.Context, sip StaticInfoParams) (*StaticsInfo, []byte, error) {

	postData, _ := json.Marshal(sip)

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetStaticInfo", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	staticsInfo := new(StaticsInfo)
	raw, err := hg.client.Do(ctx, request, &staticsInfo)
	return staticsInfo, raw, err
}

func (hg *HostGroup) GetHostTasks(ctx context.Context, hostId string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strHostName": "%s"
	}`, hostId))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetHostTasks", bytes.NewBuffer(postData))
	pxgValStr := new(PxgValStr)

	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//PxgHostFixes struct
type PxgHostFixes struct {
	PxgRetVal PxgHostFixesRetVal `json:"PxgRetVal"`
}

type PxgHostFixesRetVal struct {
	KlhstHFData     []KlhstHFDatum   `json:"KLHST_HF_DATA"`
	KlhstHFProducts []KlhstHFProduct `json:"KLHST_HF_PRODUCTS"`
}

type KlhstHFDatum struct {
	Type  string            `json:"type"`
	Value KLHSTHFDATUMValue `json:"value"`
}

type KLHSTHFDATUMValue struct {
	KlhstHFDN     string `json:"KLHST_HF_DN"`
	KlhstHFID     string `json:"KLHST_HF_ID"`
	KlhstHFProdid int64  `json:"KLHST_HF_PRODID"`
}

type KlhstHFProduct struct {
	Type  string              `json:"type"`
	Value KLHSTHFPRODUCTValue `json:"value"`
}

type KLHSTHFPRODUCTValue struct {
	KlhstHFProdid          int64  `json:"KLHST_HF_PRODID"`
	KlhstWksProductID      string `json:"KLHST_WKS_PRODUCT_ID"`
	KlhstWksProductName    string `json:"KLHST_WKS_PRODUCT_NAME"`
	KlhstWksProductVersion string `json:"KLHST_WKS_PRODUCT_VERSION"`
}

//Returns all hotfixes installed in the network.
//
//Returns:
//	- (params) contains following attributes:
//	- KLHST_HF_PRODUCTS - hotfix products (paramArray|paramParams)
//	- KLHST_WKS_PRODUCT_NAME - product name (paramString)
//	- KLHST_WKS_PRODUCT_VERSION - product version (paramString)
//	- KLHST_WKS_PRODUCT_ID - productname and version divided by slash (paramString)
//	- KLHST_HF_PRODID - hotfix product id (paramInt)
//	- KLHST_HF_DATA - hotfix data (paramArray|paramParams)
//	- KLHST_HF_ID - hotfix id (paramString)
//	- KLHST_HF_DN - hotfix display name (paramString)
//	- KLHST_HF_PRODID - hotfix product id (paramInt)
func (hg *HostGroup) GetAllHostfixes(ctx context.Context) (*PxgHostFixes, []byte, error) {

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetAllHostfixes", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	pxgHostFixes := new(PxgHostFixes)
	raw, err := hg.client.Do(ctx, request, &pxgHostFixes)
	return pxgHostFixes, raw, err
}

//Acquire administration group attributes.
//
//Returns attributes of the specified administration group.
//
//Parameters:
//	- nGroupId	(int64) Id of existing group
//
//Returns:
//	- (params) group attributes (List of group attributes for attribute names)
//
//Deprecated: Use HostGroup.GetGroupInfoEx instead
func (hg *HostGroup) GetGroupInfo(ctx context.Context, nGroupId int64, result interface{}) error {
	postData := []byte(fmt.Sprintf(`
	{
	"nGroupId": %d
	}`, nGroupId))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetGroupInfo", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hg.client.Do(ctx, request, &result)

	dec := json.NewDecoder(bytes.NewReader(raw))

	if err = dec.Decode(result); err != nil {
		return err
	}
	return err
}

//Acquire administration group attributes.
//
//Returns required attributes of the specified administration group.
//
//Parameters:
//	- nGroupId	(int64) Id of existing group
//	- pArrAttributes	([]string) Array of up to 100 strings. Each entry is an attrbute name (see List of group attributes).
//Returns:
//	- (params) group attributes (List of group attributes for attribute names)
//
//	- Remark: not working on KSC 10
func (hg *HostGroup) GetGroupInfoEx(ctx context.Context, nGroupId int64, pArrAttributes []string) ([]byte, error) {
	vFtr, _ := json.Marshal(pArrAttributes)

	postData := []byte(fmt.Sprintf(`{
	"nGroupId" : %d, 
	"pArrAttributes" : %s
	}`, nGroupId, string(vFtr)))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetGroupInfoEx", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//Move hosts from group to group.
//
//Moves hosts from root of source group to root of destination group. Operation is asynchronous.
//
//Parameters:
//	- nSrcGroupId	(int64) id of source group
//	- nDstGroupId	(int64) id of destionation group
//
//
//[out]	strActionGuid	(data.PxgValStr) id of asynchronous operation,
//to get status use AsyncActionStateChecker.CheckActionState, lStateCode "1" means OK and "0" means fail
func (hg *HostGroup) MoveHostsFromGroupToGroup(ctx context.Context, nSrcGroupId int64,
	nDstGroupId int64) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nSrcGroupId": %d, "nDstGroupId": %d
	}`, nSrcGroupId, nDstGroupId))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.MoveHostsFromGroupToGroup", bytes.NewBuffer(postData))
	pxgValStr := new(PxgValStr)

	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//KlhstWksResults struct
type KlhstWksResults struct {
	PResults PResults `json:"pResults"`
}

type PResults struct {
	KlhstWksResults []bool `json:"KLHST_WKS_RESULTS"`
}

//PInfoRaM struct
type PInfoRaM struct {
	PInfo PInfo `json:"pInfo"`
}

type PInfo struct {
	KlhstWksAnyname []string `json:"KLHST_WKS_ANYNAME"`
	KlhstWksGroupid int64    `json:"KLHST_WKS_GROUPID"`
}

//Moves hosts into a group by name or ip-address.
//
//If the entered name corresponds to the ip-address format, then the server tries to find in the database a host with the indicated ip-address. Otherwise, the server tries to interpret the host as follows:
//
//Name (KLHST_WKS_HOSTNAME)
//Display name (KLHST_WKS_DN)
//NetBIOS name (KLHST_WKS_WINHOSTNAME)
//DNS name (KLHST_WKS_DNSNAME)
//
//Parameters:
//	- pInfo	(params) the input container must contain variables:
//	- KLHST_WKS_ANYNAME (paramArray) array of strings with host names
//	- KLHST_WKS_GROUPID (paramInt) identifier of the group to which the designated hosts are to be placed
//
//[out]	pResults	(params) the output container will contain variables:
//	- KLHST_WKS_RESULTS (paramArray) array of boolean values,
//If the i-th element of this array is false,
//then the i-th host of the input array KLHST_WKS_ANYNAME could not be placed in the group (could not resolve name).
//
//Example:
//
//{
//	"pInfo": {
//		"KLHST_WKS_ANYNAME" : ["ip", "KLHST_WKS_HOSTNAME", "KLHST_WKS_DN", "KLHST_WKS_WINHOSTNAME", "KLHST_WKS_DNSNAME" ],
//		"KLHST_WKS_GROUPID" : 1 //GroupID
//		}
//	}
func (hg *HostGroup) ResolveAndMoveToGroup(ctx context.Context, pInfo PInfoRaM) (*KlhstWksResults, []byte, error) {
	postData, _ := json.Marshal(pInfo)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.ResolveAndMoveToGroup", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}
	klhstWksResults := new(KlhstWksResults)
	raw, err := hg.client.Do(ctx, request, &klhstWksResults)
	return klhstWksResults, raw, err
}

//Removes host record.
//
//Parameters:
//	- strHostName	(string) host name, a unique server-generated string (see KLHST_WKS_HOSTNAME attribute).
//It is NOT the same as computer network name (DNS-, FQDN-, NetBIOS-name)
func (hg *HostGroup) RemoveHost(ctx context.Context, strHostName string) {
	postData := []byte(fmt.Sprintf(`
	{
	"strHostName": "%s"
	}`, strHostName))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RemoveHost", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = hg.client.Do(ctx, request, nil)
}

//Remove multiple hosts.
//
//Removes multiple hosts. Function behavior depends on bForceDestroy flag. If bForceDestroy is true then hosts records are deleted. If bForceDestroy is false hosts records will be deleted only for hosts located in group "Unassigned computers" or its subgroups, others will be moved into corresponding subgroups of group "Unassigned computers".
//
//Parameters:
//	- pHostNames	([]string) array of host names
//	- bForceDestroy	(boolean) whether to force deleting hosts records
func (hg *HostGroup) RemoveHosts(ctx context.Context, pHostNames []string, bForceDestroy bool) (*PxgValStr,
	[]byte, error) {
	vFtr, _ := json.Marshal(pHostNames)
	postData := []byte(fmt.Sprintf(`
	{
	"pHostNames": %s,
	"bForceDestroy" : %v
	}`, string(vFtr), bForceDestroy))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RemoveHosts", bytes.NewBuffer(postData))
	pxgValStr := new(PxgValStr)

	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//Add a new domain to the database.
//
//Parameters:
//	- strDomain	(wstring) domain name
//	- nType	(int64) domain type:
//	- 0 - Windows NT domain
//	- 1 - Windows work group
//
//Exceptions:
//	- STDE_EXIST	domain with the specified name already exists.
func (hg *HostGroup) AddDomain(ctx context.Context, strDomain string, nType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strDomain": "%s", "nType" : %d
	}`, strDomain, nType))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddDomain", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//Removes a domain from the database.
//
//Parameters:
//	- strDomain	(string) domain name
func (hg *HostGroup) DelDomain(ctx context.Context, strDomain string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strDomain": "%s"
	}`, strDomain))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.DelDomain", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//Delete incident.
//
//Parameters:
//	- nId (int64)	incident id
func (hg *HostGroup) DeleteIncident(ctx context.Context, nId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
		"nId": %d
	}`, nId))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.DeleteIncident", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//Return a list of workstation names in the domain.
//
//The information is obtained from the domain controller. This call returns the full list of workstations in the domain, even if the workstation is now turned off.
//
//Attention:
//	- This method is deprecated, use either HostGroup.FindHostsAsync or HostGroup.FindHosts instead.
//Parameters:
//	- domain	(string) domain name.
//Returns:
//	- (array) array of hosts in domain, each item is container which contains following attributes:
//	- KLHST_WKS_HOSTNAME (paramString) host name (GUID-like identifier)
//	- KLHST_WKS_WINHOSTNAME (paramString) host windows (NetBIOS) name
//	- KLHST_WKS_STATUS (paramInt) host state:
//	- 0x00000001 - The computer is online ('visible')
//	- 0x00000002 - The computer is added into the administration group
//	- 0x00000004 - The computer has Network Agent Version installed
//	- 0x00000008 - Network Agent Version is working
//	- 0x00000010 - The computer has real time protection (RTOP)
//	- 0x00000020 - The computer has been temporarily switched into this server as a result of NLA profile switching
//	- 0x00000040 - The computer is a part of the cluster or a cluster array
//	- 0x00000080 - appliance
func (hg *HostGroup) GetDomainHosts(ctx context.Context, domain string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"domain": "%s"
	}`, domain))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetDomainHosts", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//Find incident by filter string.
//
//Finds incidents that satisfy conditions from filter string strFilter.
//
//Parameters:
//	- strFilter	(string) incident filtering expression (see Search filter syntax). See the list of incident attributes that can be used in this expression in Remarks section below
//	- pFieldsToReturn	([]string) array of incident attribute names to return. See List of incident attributes for attribute names
//	- pFieldsToOrder	([]string) array of containers each of them containing two attributes:
//"Name" of type String, name of attribute used for ordering (see Remarks below)
//"Asc" of type Boolean, ascending if true descending otherwise
//	- lMaxLifeTime	(int64) max lifetime of accessor (sec)
//[out]	strAccessor	(string) result-set ID, identifier of the server-side ordered collection of found incidents. The result-set is destroyed and associated memory is freed in following cases:
//Passed lMaxLifeTime seconds after last access to the result-set (by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//Session to the Administration Server has been closed.
//ChunkAccessor.Release has been called.
//
//Returns:
//	- (int64) number of records found
//
//Remarks:
//	- Attributes can be used in filter string (strFilter):
//
//"KLINCDT_ID"
//"KLINCDT_SEVERITY"
//"KLINCDT_ADDED"
//"KLINCDT_IS_HANDLED"
//"KLINCDT_BODY"
//"KLHST_WKS_HOSTNAME"
//Attributes can be used for ordering (pFields2Order):
//
//"KLINCDT_ID"
//"KLINCDT_SEVERITY"
//"KLINCDT_ADDED"
//"KLINCDT_IS_HANDLED"
//Attributes can NOT be used for ordering (pFields2Order):
//
//"KLINCDT_BODY"
//"KLHST_WKS_HOSTNAME"
//"GNRL_EXTRA_PARAMS"
func (hg *HostGroup) FindIncidents(ctx context.Context, strFilter string, pFieldsToReturn []string,
	lMaxLifeTime int64) (*Accessor, []byte, error) {
	vFtr, _ := json.Marshal(pFieldsToReturn)

	postData := []byte(fmt.Sprintf(`{
	"strFilter" : "%s", 
	"pFieldsToReturn" : %s, 
	"pFieldsToOrder" : [],
	"lMaxLifeTime" : %d
	}`, strFilter, string(vFtr), lMaxLifeTime))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindIncidents", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

//Acquire administration group id by its name and id of parent group.
//
//Returns administration group id by id of parent and name.
//
//Parameters:
//nParent	(int64) Id of parent group
//strName	(string) name of group
//Returns:
//(int64) id of group found and -1 if no group was found.
func (hg *HostGroup) GetGroupId(ctx context.Context, nParent int64, strName string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nParent": %d,
	"strName": "%s"
	}`, nParent, strName))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetGroupId", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//SubGroupsList struct
type SubGroupsList struct {
	PxgRetVal []Groups `json:"PxgRetVal"`
}

type FluffyValue struct {
	Groups                []Groups `json:"groups"`
	GrpPartOfAdViewByRule bool     `json:"grp_part_of_ad_view_by_rule"`
	ID                    int64    `json:"id"`
	Name                  string   `json:"name"`
}

type FluffyGroup struct {
	Type  string      `json:"type"`
	Value FluffyValue `json:"value"`
}

type PurpleValue struct {
	Groups                []FluffyGroup `json:"groups"`
	GrpPartOfAdViewByRule bool          `json:"grp_part_of_ad_view_by_rule"`
	ID                    int64         `json:"id"`
	Name                  string        `json:"name"`
}

type PurpleGroup struct {
	Type  string      `json:"type"`
	Value PurpleValue `json:"value"`
}

type PxgRetValValue struct {
	Groups                []PurpleGroup `json:"groups"`
	GrpPartOfAdViewByRule bool          `json:"grp_part_of_ad_view_by_rule"`
	ID                    int64         `json:"id"`
	Name                  string        `json:"name"`
}

type Groups struct {
	Type  string         `json:"type"`
	Value PxgRetValValue `json:"value"`
}

//Acquire administration group subgroups tree.
//
//Parameters:
//	- nGroupId	(int64) Id of existing group
//	- nDepth	(int64) depth of subgroups tree, 0 means all grandchildren tree with no limits
//
//Returns:
//	- (array) array of containers paramParams, each of them contains up to three attributes:
//"id" (subgroup id), "name" (subgroup name) and "groups" (similar recursive array), may be NULL.
func (hg *HostGroup) GetSubgroups(ctx context.Context, nGroupId int64, nDepth int64) (*SubGroupsList, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nParent": %d,
	"nDepth": %d
	}`, nGroupId, nDepth))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetSubgroups", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}
	subGroupsList := new(SubGroupsList)
	raw, err := hg.client.Do(ctx, request, &subGroupsList)
	return subGroupsList, raw, err
}

//Move multiple hosts into specified administration group.
//
//Parameters:
//	- nGroup	(int64) id of destination group
//	- pHostNames	([]string) array of host names
func (hg *HostGroup) MoveHostsToGroup(ctx context.Context, nGroup int64, pHostNames []string) (*PxgValStr, []byte,
	error) {
	vFtr, _ := json.Marshal(pHostNames)
	postData := []byte(fmt.Sprintf(`
	{
	"nType": %d,
	"pHostNames": %s
	}`, nGroup, string(vFtr)))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.MoveHostsToGroup", bytes.NewBuffer(postData))
	pxgValStr := new(PxgValStr)

	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//Delete administration group.
//
//Parameters:
//	- nGroup	(int64) Id of existing group to delete
//	- nFlags	(int64) flags. May have following value
//
//	- 1 (default value) group is deleted only if it is empty, "empty" means it doesn't contain subgroups, hosts, policies, tasks, slave servers
//	- 2 delete group with subgroups, policies and tasks
//	- 3 delete group with subgroups, hosts, policies and tasks
//
//Return:
//	- strActionGuid	(data.PxgValStr) id of asynchronous operation, to get status use AsyncActionStateChecker.CheckActionState,
//lStateCode "1" means OK and "0" means fail
func (hg *HostGroup) RemoveGroup(ctx context.Context, nGroup, nFlags int64) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nGroup": %d,
	"nFlags": %d
	}`, nGroup, nFlags))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RemoveGroup", bytes.NewBuffer(postData))
	pxgValStr := new(PxgValStr)

	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//todo
func (hg *HostGroup) UpdateGroup(ctx context.Context, nGroup int64, pInfo interface{}) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nGroup": %d
	}`, nGroup))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.UpdateGroup", bytes.NewBuffer(postData))
	pxgValStr := new(PxgValStr)

	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//Restarts specified network scanning type.
//
//Parameters:
//	- nType	(int64) IN network scanning type:
//	- 1 - AD
//	- 2 - Ms network fast scanning
//	- 3 - Ms network full scanning
//	- 4 - Ip diapasons scanning
func (hg *HostGroup) RestartNetworkScanning(ctx context.Context, nType int64) (*PxgRetError, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nType": %d
	}`, nType))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RestartNetworkScanning", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	pxgRetError := new(PxgRetError)
	raw, err := hg.client.Do(ctx, request, &pxgRetError)
	return pxgRetError, raw, err
}

//Zero virus count for hosts in group and all subgroups.
//
//Parameters:
//nParent	(int64) Id of group to start from
//[out]	strActionGuid	(string) id of asynchronous operation,
//to get status use AsyncActionStateChecker.CheckActionState, lStateCode "1" means OK and "0" means fail
func (hg *HostGroup) ZeroVirusCountForGroup(ctx context.Context, nParent int64) (*RequestID, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nParent": %d
	}`, nParent))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.ZeroVirusCountForGroup", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	requestID := new(RequestID)
	raw, err := hg.client.Do(ctx, request, &requestID)

	return requestID, raw, err
}

// ZeroVirusCountForHostsParams struct
type ZeroVirusCountForHostsParams struct {
	NParent []string `json:"nParent"`
}

// Zero virus count for specified hosts.
//
//Parameters:
//pHostNames	(array) array of host names
//[out]	strActionGuid	(wstring) id of asynchronous operation,
//to get status use AsyncActionStateChecker.CheckActionState, lStateCode "1" means OK and "0" means fail
func (hg *HostGroup) ZeroVirusCountForHosts(ctx context.Context, zvcfhp ZeroVirusCountForHostsParams) (*RequestID,
	[]byte, error) {

	postData, _ := json.Marshal(zvcfhp)

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.ZeroVirusCountForHosts", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}
	requestID := new(RequestID)
	raw, err := hg.client.Do(ctx, request, &requestID)

	return requestID, raw, err
}

func (hg *HostGroup) SS_Read(ctx context.Context) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
   		"strHostName":"879e5ad6-2e6b-4e68-8ba5-146fc45c4f82",
   		"strType":"SS_SETTINGS",
   		"strProduct":"1103",
   		"strVersion":"11.0.0.29",
		"strSection": "86"
	}`))

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.SS_Read", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}
