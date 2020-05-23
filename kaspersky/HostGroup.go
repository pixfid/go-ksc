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

//	HostGroup Class Reference
//
//	Hosts and management groups processing..
//
//	List of all members.
type HostGroup service

//	Add a new domain to the database.
//
//	Parameters:
//	- strDomain	(string) domain name
//	- nType	(int64) domain type:
//	- 0 - Windows NT domain
//	- 1 - Windows work group
//
//	Exceptions:
//	- STDE_EXIST	domain with the specified name already exists.
func (hg *HostGroup) AddDomain(ctx context.Context, strDomain string, nType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strDomain": "%s", "nType" : %d
	}`, strDomain, nType))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddDomain", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	Create new administration group.
//
//	Creates new group with the specified attributes and returns its Id. If such group already exists returns Id of existing group.
//
//	Parameters:
//	- pInfo	(params) container with group attributes. May contain following attributes (see List of group attributes):
//	- "name"
//	- "parentId"
//
//	Returns:
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
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Add hosts from specified group to synchronization.
//
//	Parameters:
//	- nGroupId	(int64) group id
//	- strSSType	(string) setting storage identity (empty string means synchronization of all setting storages)
//	Return:
//	- strActionGuid	(string) id of asynchronous operation, to get status use AsyncActionStateChecker.
//	CheckActionState, lStateCode "1" means OK and "0" means fail
func (hg *HostGroup) AddGroupHostsForSync(ctx context.Context, nGroupId int64, strSSType string) (*WActionGUID, []byte,
	error) {
	postData := []byte(fmt.Sprintf(` {"nGroupId": %d , "strSSType": "%s" }`, nGroupId, strSSType))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddGroupHostsForSync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := hg.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

//	Create new host record.
//
//	Parameters:
//	- pInfo	(params) container with host attributes. Must contain following attributes (see List of host attributes).
//	|- "KLHST_WKS_DN"
//	|- "KLHST_WKS_GROUPID"
//	|- "KLHST_WKS_WINDOMAIN", may be empty string only for non-windows hosts
//	|- "KLHST_WKS_WINHOSTNAME", may be empty string only for non-windows hosts
//	|- "KLHST_WKS_DNSDOMAIN"
//	|- "KLHST_WKS_DNSNAME"
//
//	Example params for Create New Host:
//
//	type NewHostParams struct {
//		PInfo PInfo `json:"pInfo"`
//	}
//
//	type PInfo struct {
//		KlhstWksDN          string `json:"KLHST_WKS_DN"`
//		KlhstWksGroupid     int64  `json:"KLHST_WKS_GROUPID"`
//		KlhstWksWindomain   string `json:"KLHST_WKS_WINDOMAIN"`
//		KlhstWksWinhostname string `json:"KLHST_WKS_WINHOSTNAME"`
//		KlhstWksDnsdomain   string `json:"KLHST_WKS_DNSDOMAIN"`
//		KlhstWksDnsname     string `json:"KLHST_WKS_DNSNAME"`
//	}
//	Returns:
//	- (string) unique server-generated string
func (hg *HostGroup) AddHost(ctx context.Context, params interface{}) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//TODO AddHostsForSync

//	AddIncidentsParams struct
type AddIncidentsParams struct {
	PData PData `json:"pData"`
}

//	PData struct
type PData struct {
	KlincdtSeverity  *int64    `json:"KLINCDT_SEVERITY"`
	KlincdtAdded     *DateTime `json:"KLINCDT_ADDED"`
	KlincdtBody      *string   `json:"KLINCDT_BODY"`
	KlhstWksHostname *string   `json:"KLHST_WKS_HOSTNAME"`
	KlhstuserID      *int64    `json:"KLHSTUSER_ID"`
}

type DateTime struct {
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

func (hg *HostGroup) AddIncident(ctx context.Context, params AddIncidentsParams) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddIncident", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//	Removes a domain from the database.
//
//	Parameters:
//	- strDomain	(string) domain name
func (hg *HostGroup) DelDomain(ctx context.Context, strDomain string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strDomain": "%s"
	}`, strDomain))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.DelDomain", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//Delete incident.
//
//Parameters:
//	- nId (int64)	incident id
func (hg *HostGroup) DeleteIncident(ctx context.Context, nId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nId": %d}`, nId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.DeleteIncident", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

type HGParams struct {
	WstrFilter        string          `json:"wstrFilter"`
	VecFieldsToReturn []string        `json:"vecFieldsToReturn"`
	VecFieldsToOrder  []FieldsToOrder `json:"vecFieldsToOrder"`
	PParams           PParams         `json:"pParams"`
	LMaxLifeTime      int64           `json:"lMaxLifeTime"`
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
func (hg *HostGroup) FindGroups(ctx context.Context, params HGParams) (*Accessor, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindGroups", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

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
func (hg *HostGroup) FindHosts(ctx context.Context, params HGParams) (*Accessor, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHosts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

//	Find host asynchronously by filter string.
//
//	Finds hosts asynchronously that satisfy conditions from filter string wstrFilter,
//	and creates a server-side collection of found hosts. Search is performed over the hierarchy
//
//	Parameters:
//	- wstrFilter	(string) filter string, contains a condition over host attributes, see also Search filter syntax.
//	- vecFieldsToReturn	([]string) array of host attribute names to return.
//	See List of host attributes for attribute names
//	- vecFieldsToOrder	([]string) array of containers each of them containing two attributes :
//		|- "Name" (string) name of attribute used for sorting
//		|- "Asc" (paramBool) ascending if true descending otherwise
//	- pParams	(params) extra options. Possible attributes are listed below
//(see details in Extra search attributes for hosts and administration groups):
//		|- KLSRVH_SLAVE_REC_DEPTH
//		|- KLGRP_FIND_FROM_CUR_VS_ONLY
//	- lMaxLifeTime	(int64) max result-set lifetime in seconds, not more than 7200
//
//	Return:
//	- data.RequestID	(string) identity of asynchronous operation,
//
//	to get status use AsyncActionStateChecker.CheckActionState, lStateCode "1" means OK and "0" means fail
//
//	to get accessor id call HostGroup.FindHostsAsyncGetAccessor
//
//	to cancel operation call HostGroup.FindHostsAsyncCancel
func (hg *HostGroup) FindHostsAsync(ctx context.Context, params HGParams) (*RequestID, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHostsAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	requestID := new(RequestID)
	raw, err := hg.client.Do(ctx, request, &requestID)
	return requestID, raw, err
}

//	Cancel FindHostsAsync operation.
//
//	Cancels asynchronous operation HostGroup.FindHostsAsync
//
//	Parameters:
//	- strRequestId	(string) identity of asynchronous operation
func (hg *HostGroup) FindHostsAsyncCancel(ctx context.Context, strRequestId string) {
	postData := []byte(fmt.Sprintf(`
	{
	"strRequestId": "%s"
	}`, strRequestId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHostsAsyncCancel", bytes.NewBuffer(postData))
	if err != nil {
		return
	}

	_, err = hg.client.Do(ctx, request, nil)
}

//	Get result of FindHostsAsync operation.
//
//	Gets result of asynchronous operation HostGroup.FindHostsAsync
//
//	Parameters:
//	- strRequestId	(string) identity of asynchronous operation
//	- strAccessor	(string) result-set ID, identifier of the server-side ordered collection of found hosts. The result-set is destroyed and associated memory is freed in following cases:
//	Passed lMaxLifeTime seconds after last access to the result-set (by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//	Session to the Administration Server has been closed.
//	ChunkAccessor.Release has been called.
//	Return:
//	- pFailedSlavesInfo	(params) information about slave servers the search for which failed due to various reasons, contains array KLGRP_FAILED_SLAVES_PARAMS of params which have attributes:
//		|- KLSRVH_SRV_ID - Slave server id (int64)
//		|- KLSRVH_SRV_DN - Slave server display name (string)
//	- (int64) number of found hosts
func (hg *HostGroup) FindHostsAsyncGetAccessor(ctx context.Context, strRequestId string) (*AsyncAccessor,
	[]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId" : "%s" }`, strRequestId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHostsAsyncGetAccessor", bytes.NewBuffer(postData))

	asyncAccessor := new(AsyncAccessor)
	raw, err := hg.client.Do(ctx, request, &asyncAccessor)
	return asyncAccessor, raw, err
}

type FindIncidentsParams struct {
	StrFilter       *string          `json:"strFilter,omitempty"`
	PFieldsToReturn *[]string        `json:"pFieldsToReturn,omitempty"`
	PFieldsToOrder  *[]FieldsToOrder `json:"pFieldsToOrder,omitempty"`
	LMaxLifeTime    *int64           `json:"lMaxLifeTime,omitempty"`
}

//	Find incident by filter string.
//
//	Finds incidents that satisfy conditions from filter string strFilter.
//
//	Parameters:
//	- strFilter	(string) incident filtering expression (see Search filter syntax). See the list of incident attributes that can be used in this expression in Remarks section below
//	- pFieldsToReturn	([]string) array of incident attribute names to return. See List of incident attributes for attribute names
//	- pFieldsToOrder	([]string) array of containers each of them containing two attributes:
//		|- "Name" of type String, name of attribute used for ordering (see Remarks below)
//		|- "Asc" of type bool, ascending if true descending otherwise
//	- lMaxLifeTime	(int64) max lifetime of accessor (sec)
//
//	Example request params struct:
//	{
//	  "strFilter": "(&(!KLINCDT_ID = 0))",
//	  "pFieldsToReturn": [
//	    "KLINCDT_ID",
//	    "KLINCDT_BODY",
//	    "KLHST_WKS_HOSTNAME",
//	    "KLINCDT_SEVERITY"
//	  ],
//	  "pFieldsToOrder": [
//	    {
//	      "type": "params",
//	      "value": {
//	        "Name": "KLINCDT_ID",
//	        "Asc": true
//	      }
//	    }
//	  ],
//	  "lMaxLifeTime": 120
//	}
//
//	Return:
//	- strAccessor	(string) result-set ID, identifier of the server-side ordered collection of found incidents.
//	The result-set is destroyed and associated memory is freed in following cases:
//	Passed lMaxLifeTime seconds after last access to the result-set (by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//	Session to the Administration Server has been closed.
//	ChunkAccessor.Release has been called.
//	- (int64) number of records found
//
//Remarks:
//	Attributes can be used in filter string (strFilter):
//
//	|-"KLINCDT_ID"
//	|-"KLINCDT_SEVERITY"
//	|-"KLINCDT_ADDED"
//	|-"KLINCDT_IS_HANDLED"
//	|-"KLINCDT_BODY"
//	|-"KLHST_WKS_HOSTNAME"
//
//	Attributes can be used for ordering (pFields2Order):
//
//	|-"KLINCDT_ID"
//	|-"KLINCDT_SEVERITY"
//	|-"KLINCDT_ADDED"
//	|-"KLINCDT_IS_HANDLED"
//
//	Attributes can NOT be used for ordering (pFields2Order):
//
//	|-"KLINCDT_BODY"
//	|-"KLHST_WKS_HOSTNAME"
//	|-"GNRL_EXTRA_PARAMS"
func (hg *HostGroup) FindIncidents(ctx context.Context, params FindIncidentsParams) (*Accessor, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindIncidents", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

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

//	Finds existing users.
//
//	Finds users that satisfy conditions from filter string strFilter.
//
//	Parameters:
//	- strFilter	(string) filter string, see Search filter syntax
//	- pFieldsToReturn	(array) array of user's attribute names to return. See List of user's attributes
//	- pFieldsToOrder	(array) array of containers each of them containing two attributes:
//		|- "Name" of type String, name of attribute used for sorting
//		|- "Asc" of type bool, ascending if true descending otherwise
//	- lMaxLifeTime	(int64) max lifetime of accessor (sec)
//
//	Return:
//	- strAccessor	(string) result-set ID, identifier of the server-side ordered collection of found users.
//	The result-set is destroyed and associated memory is freed in following cases:
//	Passed lMaxLifeTime seconds after last access to the result-set (by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//	Session to the Administration Server has been closed.
//	ChunkAccessor.Release has been called.
//	- (int64) number of records found
func (hg *HostGroup) FindUsers(ctx context.Context, params UHGParams) (*Accessor, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindUsers", bytes.NewBuffer(postData))

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

//Returns all hotfixes installed in the network.
//
//Returns:
//	- (params) contains following attributes:
//	- KLHST_HF_PRODUCTS - hotfix products (paramArray|paramParams)
//	- KLHST_WKS_PRODUCT_NAME - product name (string)
//	- KLHST_WKS_PRODUCT_VERSION - product version (string)
//	- KLHST_WKS_PRODUCT_ID - productname and version divided by slash (string)
//	- KLHST_HF_PRODID - hotfix product id (int64)
//	- KLHST_HF_DATA - hotfix data (paramArray|paramParams)
//	- KLHST_HF_ID - hotfix id (string)
//	- KLHST_HF_DN - hotfix display name (string)
//	- KLHST_HF_PRODID - hotfix product id (int64)
func (hg *HostGroup) GetAllHostfixes(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetAllHostfixes", nil)
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	ProductComponents is returned by GetComponentsForProductOnHost
type ProductComponents struct {
	ProductComponentsArray []ProductComponentsArray `json:"PxgRetVal"`
}

type ProductComponentsArray struct {
	Type             *string           `json:"type,omitempty"`
	ProductComponent *ProductComponent `json:"value,omitempty"`
}

type ProductComponent struct {
	KlhstPrcstComponentDN      *string                     `json:"KLHST_PRCST_COMPONENT_DN,omitempty"`
	KlhstPrcstComponentID      *KlhstPrcstComponentID      `json:"KLHST_PRCST_COMPONENT_ID,omitempty"`
	KlhstPrcstComponentStatus  *int64                      `json:"KLHST_PRCST_COMPONENT_STATUS,omitempty"`
	KlhstPrcstComponentVersion *KlhstPrcstComponentVersion `json:"KLHST_PRCST_COMPONENT_VERSION,omitempty"`
}

type KlhstPrcstComponentID struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type KlhstPrcstComponentVersion struct {
	Type  *string `json:"type,omitempty"`
	Value *int64  `json:"value,omitempty"`
}

//	Return array of product components for specified host and product.
//
//	Parameters:
//	- strHostName	(string) host name
//	- strProductName	(string) product name
//	- strProductVersion	(string) product version
//
//	Returns:
//	- (array) each item of array is container (params) with information about component,
//	see List of product component attributes
func (hg *HostGroup) GetComponentsForProductOnHost(ctx context.Context, strHostName, strProductName,
	strProductVersion string) (*ProductComponents, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostName": "%s","strProductName": "%s","strProductVersion": "%s"}`,
		strHostName, strProductName, strProductVersion))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetComponentsForProductOnHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	var productComponents *ProductComponents
	raw, err := hg.client.Do(ctx, request, &productComponents)
	return productComponents, raw, err
}

//	Return a list of workstation names in the domain.
//
//	The information is obtained from the domain controller. This call returns the full list of workstations in the domain, even if the workstation is now turned off.
//
//	Parameters:
//	- domain	(string) domain name.
//	Returns:
//	- (array) array of hosts in domain, each item is container which contains following attributes:
//	- KLHST_WKS_HOSTNAME (string) host name (GUID-like identifier)
//	- KLHST_WKS_WINHOSTNAME (string) host windows (NetBIOS) name
//	- KLHST_WKS_STATUS (int64) host state:
//	- 0x00000001 - The computer is online ('visible')
//	- 0x00000002 - The computer is added into the administration group
//	- 0x00000004 - The computer has Network Agent Version installed
//	- 0x00000008 - Network Agent Version is working
//	- 0x00000010 - The computer has real time protection (RTOP)
//	- 0x00000020 - The computer has been temporarily switched into this server as a result of NLA profile switching
//	- 0x00000040 - The computer is a part of the cluster or a cluster array
//	- 0x00000080 - appliance
//	Deprecated:
//	- This method is deprecated, use either HostGroup.FindHostsAsync or HostGroup.FindHosts instead.
func (hg *HostGroup) GetDomainHosts(ctx context.Context, domain string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"domain": "%s"}`, domain))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetDomainHosts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	List of Windows domain in the network.
//
//	Returns:
//	- (array) array of domains, each item is container which contains following attributes:
//		|- KLHST_WKS_WINDOMAIN (string) domain name
//		|- KLHST_WKS_WINDOMAIN_TYPE (int64) domain type:
//
//			0 - Windows NT domain
//			1 - Windows work group
func (hg *HostGroup) GetDomains(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetDomains", nil)
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire administration group id by its name and id of parent group.
//
//	Returns administration group id by id of parent and name.
//
//	Parameters:
//	- nParent	(int64) Id of parent group
//	- strName	(string) name of group
//
//	Returns:
//	- (int64) id of group found and -1 if no group was found.
func (hg *HostGroup) GetGroupId(ctx context.Context, nParent int64, strName string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nParent": %d, "strName": "%s"}`, nParent, strName))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetGroupId", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Acquire administration group attributes.
//
//	Returns attributes of the specified administration group.
//
//	Parameters:
//	- nGroupId	(int64) Id of existing group
//
//	Returns:
//	- (params) group attributes (List of group attributes for attribute names)
//
//Deprecated: Use HostGroup.GetGroupInfoEx instead
func (hg *HostGroup) GetGroupInfo(ctx context.Context, nGroupId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nGroupId": %d}`, nGroupId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetGroupInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire administration group attributes.
//
//	Returns required attributes of the specified administration group.
//
//	Parameters:
//	- nGroupId	(int64) Id of existing group
//	- pArrAttributes	([]string) Array of up to 100 strings. Each entry is an attrbute name (see List of group attributes).
//
// Example request params:
//
//	type GroupInfoExParams struct {
//		NGroupID       int64    `json:"nGroupId"`
//		PArrAttributes []string `json:"pArrAttributes"`
//	}
//	Returns:
//	- (params) group attributes (List of group attributes for attribute names)
//
//	Remark: not working on KSC 10
func (hg *HostGroup) GetGroupInfoEx(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetGroupInfoEx", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	ProductFixes struct
type ProductFixes struct {
	Fixes []Fixes `json:"PxgRetVal"`
}

type Fixes struct {
	Type       string     `json:"type"`
	FixesValue FixesValue `json:"value"`
}

type FixesValue struct {
	KlhstHFDN string `json:"KLHST_HF_DN"`
	KlhstHFID string `json:"KLHST_HF_ID"`
}

//	Return array of hotfixes for specified host and product.
//
//	Array is ordered according hotfix installation order.
//
//	Parameters:
//	- strHostName	(string) host name
//	- strProductName	(string) product name
//	- strProductVersion	(string) product version
//
//	Returns:
//	- (array) each item of array is container (params) with following attributes:
//		|- KLHST_HF_ID - hotfix id (string)
//		|- KLHST_HF_DN - hotfix display name (string)
//
//	Example response:
//	{
//	  "PxgRetVal" : [
//	    {
//	      "type" : "params",
//	      "value" : {
//	        "KLHST_HF_DN" : "a",
//	        "KLHST_HF_ID" : "{F1FE7235-5744-49D8-8BF6-A55A345383E2}"
//	      }
//	    }
//	  ]
//	}
func (hg *HostGroup) GetHostfixesForProductOnHost(ctx context.Context, strHostName, strProductName,
	strProductVersion string) (*ProductFixes, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostName": "%s","strProductName": "%s","strProductVersion": "%s"}`,
		strHostName, strProductName, strProductVersion))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetHostfixesForProductOnHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	productFixes := new(ProductFixes)
	raw, err := hg.client.Do(ctx, request, &productFixes)
	return productFixes, raw, err
}

//	Acquire specified host attributes.
//
//	Returns specified attributes of given host
//
//	Parameters:
//	- strHostName	(wstring) host name, a unique server-generated string (see KLHST_WKS_HOSTNAME attribute).
//	It is NOT the same as computer network name (DNS-, FQDN-, NetBIOS-name)
//	- pFields2Return	(array) array of names of host attributes to return.
//	See List of host attributes for attribute names
//
//	Example request params struct:
//
//	type HostInfoParams struct {
//		StrHostName    string   `json:"strHostName"`
//		PFields2Return []string `json:"pFields2Return"`
//	}
//
//	Returns:
//	- (params) container with host attributes. See List of host attributes.
func (hg *HostGroup) GetHostInfo(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetHostInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	Return information about installed products on the host.
//
//	Parameters:
//	- strHostName	(wstring) host name (GUID-like identifier)
//
//	Returns:
//	- (params) contains containers with names of products that contain containers with verions of product, i.e:
//		|- <Product> (paramParams)
//		|- <Version> (paramParams)
//			|- InstallTime
//			|- InstallationId
//			|- DisplayName
//			|- BaseRecords
//			|- ConnDisplayName
//			|- ConnProdVersion
//			|- ConnectorComponentName
//
//	See also:
//	Product info attributes
//	Local settings and policy format for some products
func (hg *HostGroup) GetHostProducts(ctx context.Context, strHostName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostName": "%s"}`, strHostName))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetHostProducts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	Return server specific identity to acquire and manage host tasks.
//
//	Parameters:
//	- strHostName	(string) hostid
//
//	Returns:
//	- (string) server object ID to acquire and manage host tasks, used to HostTasks
func (hg *HostGroup) GetHostTasks(ctx context.Context, hostId string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostName": "%s"}`, hostId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetHostTasks", bytes.NewBuffer(postData))
	pxgValStr := new(PxgValStr)

	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

type InstanceStatisticsParams struct {
	VecFilterFields []string `json:"vecFilterFields"` // <- can be empty, but not nil
}

//	Acquire Server statistics info.
//
//	Parameters:
//	- vecFilterFields	(array) Array of filtered attributes
//		|- "KLSRV_ST_ALL_CONS_CNT"	Number of all connections and number of all nAgent Version connections
//		|- "KLSRV_ST_CTLNGT_CONS_CNT"	Number of controlled nAgent Version connections
//		|- "KLSRV_ST_NETWORK_DOMAIN_SCANNED"	Currently scanned domain name
//		|- "KLSRV_ST_VIRT_SERVER_COUNT"	Virtual servers count
//		|- "KLSRV_ST_TOTAL_HOSTS_COUNT"	Total active hosts count
//		|- "KLSRV_ST_VIRT_SERVERS_DETAILS"	Array of active hosts count on virtual server
//		|- "KLSRV_ST_CON_EVENTS"	Container with ConEvents statistics
//
//	Returns:
//	- filtered statistic
//
//	Remark:
//	- not working on KSC 10
func (hg *HostGroup) GetInstanceStatistics(ctx context.Context, params InstanceStatisticsParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetInstanceStatistics", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

type StaticInfoParams struct {
	PValues []string `json:"pValues"`
}

// Return server run-time info.
//
//	Parameters:
//	- pValues	(array) string array with names of requested values, possible values:
//		|- KLADMSRV_SSS_PORT - server port (int64)
//		|- KLADMSRV_SSS_ID - server id (string)
//		|- KLADMSRV_VS_LICDISABLED - licensing for the VS is disabled (paramBool)
//		|- KLADMSRV_SAAS_BLOCKED - adding new virtual servers is blocked due to expired/absent/blacklisted license (paramBool)
//		|- KLADMSRV_SAAS_EXPIRED_DAYS_TO_WORK - adding new virtual servers will be blocked in c_szwIfSaasExpiredDaysToWork days (int64)
//		|- KLADMSRV_SAAS_OVERUSE - number of VS created is more specified in the license (paramBool)
//		|- KLADMSRV_IF_WAIK_INSTALLED - true if WAIK is installed (paramBool)
//
//	Returns:
//	- (params) requsted values
func (hg *HostGroup) GetRunTimeInfo(ctx context.Context, params StaticInfoParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetRunTimeInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	Return server static info.
//
//	Parameters:
//	- pValues	(array) string array with names of requested values, possible values are listed below.
//		|- KLADMSRV_SERVER_CERT - server certificate (paramBinary)
//		|- KLADMSRV_SERVER_KEY - server key (paramBinary)
//		|- InstancePort - instance port (int64)
//		|- KLADMSRV_SERVER_ADDRESSSES - array of server addresses that can be used by clients to connect to the administration server (paramArray|string)
//		|- KLADMSRV_SERVER_UNDER_SYSTEM_ACCOUNT - Name of account used by Administration Server (string)
//		|- KLADMSRV_OLA_ROOTCER_ACTUAL - Array of allowed OLA certificates, paramArray, each entry is a certificate as 'DER-encoded binary X.509' of type paramBinary
//		|- KLADMSRV_OLA_ROOTCER_REVOKED - Array of disallowed (revoked) OLA certificates, paramArray, each entry is a certificate as 'DER-encoded binary X.509' of type paramBinary
//
//	If pValues is NULL then described below values will be returned:
//
//	- KLADMSRV_IS_VIRTUAL - true if server is virtual (paramBool)
//	- KLADMSRV_VSID - VS id (int64)
//	- KLADMSRV_GRP_ROOT - id of group groups (int64)
//	- KLADMSRV_GRP_UNASSIGNED - id of group unassigned (int64)
//	- KLADMSRV_GRP_SUPER - id of group super (int64)
//	- KLADMSRV_SERVER_VERSION_ID - server version id (int64)
//	- KLADMSRV_B2B_CLOUD_MODE - if the server installed in the 'B2B Cloud' mode (paramBool)
//	- KLADMSRV_PCLOUD_MODE - if public cloud support mode is turned on (paramBool)
//	- KLADMSRV_PRODUCT_FULL_VERSION - server product full version (string)
//	- KLADMSRV_SERVER_HOSTNAME - server host name (string)
//	- KLADMSRV_PRODUCT_VERSION - server product version (string)
//	- KLADMSRV_PRODUCT_NAME - server product name (string)
//	- KLADMSRV_FORCE_SYNC_SUPPORTED - force sync supported (paramBool)
//	- KLADMSRV_MAINTENANCE_SUPPORTED - if maintenance task must is supported (paramBool)
//	- KLADMSRV_NAgent Version_RUNNING - true if nAgent Version is running (paramBool)
//	- KLADMSRV_NAC_IS_BEING_USED - true if NAC is used (paramBool)
//	- KLADMSRV_SPL_PPC_ENABLED - If password policy compliance for SPL users is enabled (paramBool)
//	- KLSRV_NETSIZE - network size (int64)
//	- KLADMSRV_USERID - id of the user account; NULL if OS user account is used (int64)
//	- KLADMSRV_NEED_UNC_PATH - if UNC path must be specified into backup task settings (paramBool)
//	- KLADMSRV_EV_EV_SIZE - average size of a single event, Kb (paramDouble)
//
//	Returns:
//	- (params) requsted values
func (hg *HostGroup) GetStaticInfo(ctx context.Context, params StaticInfoParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetStaticInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire administration group subgroups tree.
//
//	Parameters:
//	- nGroupId	(int64) Id of existing group
//	- nDepth	(int64) depth of subgroups tree, 0 means all grandchildren tree with no limits
//
//	Returns:
//	- (array) array of containers paramParams, each of them contains up to three attributes:
//	"id" (subgroup id), "name" (subgroup name) and "groups" (similar recursive array), may be NULL.
func (hg *HostGroup) GetSubgroups(ctx context.Context, nGroupId int64, nDepth int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nParent": %d, "nDepth": %d }`, nGroupId, nDepth))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetSubgroups", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	Id of predefined root group "Managed computers".
//
//	Returns:
//	- (data.PxgValInt) id of predefined root group "Managed computers"
func (hg *HostGroup) GroupIdGroups(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GroupIdGroups", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Id of predefined group "Master server".
//
//	Returns:
//	- (data.PxgValInt) id of predefined group "Master server" ("Super")
func (hg *HostGroup) GroupIdSuper(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GroupIdSuper", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Id of predefined group "Unassigned computers".
//
//	Returns:
//	- (data.PxgValInt) id of predefined group "Unassigned computers"
func (hg *HostGroup) GroupIdUnassigned(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GroupIdUnassigned", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Move hosts from group to group.
//
//	Moves hosts from root of source group to root of destination group. Operation is asynchronous.
//
//	Parameters:
//	- nSrcGroupId	(int64) id of source group
//	- nDstGroupId	(int64) id of destionation group
//
// Return:
//	- strActionGuid	(data.WActionGUID) id of asynchronous operation,
//	to get status use AsyncActionStateChecker.CheckActionState, lStateCode "1" means OK and "0" means fail
func (hg *HostGroup) MoveHostsFromGroupToGroup(ctx context.Context, nSrcGroupId int64,
	nDstGroupId int64) (*WActionGUID, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nSrcGroupId": %d, "nDstGroupId": %d}`, nSrcGroupId, nDstGroupId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.MoveHostsFromGroupToGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}
	wActionGUID := new(WActionGUID)
	raw, err := hg.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

type HostsToGroupParams struct {
	NGroup     int64    `json:"nGroup"`
	PHostNames []string `json:"pHostNames"`
}

//Move multiple hosts into specified administration group.
//
//Parameters:
//	- nGroup	(int64) id of destination group
//	- pHostNames	([]string) array of host names
func (hg *HostGroup) MoveHostsToGroup(ctx context.Context, params HostsToGroupParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.MoveHostsToGroup",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	Delete administration group.
//
//	Parameters:
//	- nGroup	(int64) Id of existing group to delete
//	- nFlags	(int64) flags. May have following value
//		|- 1 (default value) group is deleted only if it is empty, "empty" means it doesn't contain subgroups, hosts,
//		policies, tasks, slave servers
//		|- 2 delete group with subgroups, policies and tasks
//		|- 3 delete group with subgroups, hosts, policies and tasks
//
//Return:
//	- strActionGuid	(data.PxgValStr) id of asynchronous operation, to get status use AsyncActionStateChecker.CheckActionState,
//	lStateCode "1" means OK and "0" means fail
func (hg *HostGroup) RemoveGroup(ctx context.Context, nGroup, nFlags int64) (*WActionGUID, []byte, error) {
	postData := []byte(fmt.Sprintf(`{ "nGroup": %d, "nFlags": %d }`, nGroup, nFlags))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RemoveGroup", bytes.NewBuffer(postData))
	wActionGUID := new(WActionGUID)

	raw, err := hg.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

//	Removes host record.
//
//	Parameters:
//	- strHostName	(string) host name, a unique server-generated string (see KLHST_WKS_HOSTNAME attribute).
//	It is NOT the same as computer network name (DNS-, FQDN-, NetBIOS-name)
func (hg *HostGroup) RemoveHost(ctx context.Context, strHostName string) {
	postData := []byte(fmt.Sprintf(`{ "strHostName": "%s" }`, strHostName))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RemoveHost", bytes.NewBuffer(postData))
	if err != nil {
		return
	}

	_, err = hg.client.Do(ctx, request, nil)
}

//	Remove multiple hosts.
//
//	Removes multiple hosts. Function behavior depends on bForceDestroy flag.
//	If bForceDestroy is true then hosts records are deleted.
//	If bForceDestroy is false hosts records will be deleted only for hosts located in group "Unassigned computers"
//	or its subgroups, others will be moved into corresponding subgroups of group "Unassigned computers".
//
//	Parameters:
//	- pHostNames	([]string) array of host names
//	- bForceDestroy	(bool) whether to force deleting hosts records
//
//	Example request params struct:
//
//	type RemoveHostsParams struct {
//		PHostNames    []string `json:"pHostNames"`
//		BForceDestroy bool     `json:"bForceDestroy"`
//	}
//
func (hg *HostGroup) RemoveHosts(ctx context.Context, params struct{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RemoveHosts", bytes.NewBuffer(postData))

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
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

//	Moves hosts into a group by name or ip-address.
//
//	If the entered name corresponds to the ip-address format,
//	then the server tries to find in the database a host with the indicated ip-address.
//	Otherwise, the server tries to interpret the host as follows:
//
//	- Name (KLHST_WKS_HOSTNAME)
//	- Display name (KLHST_WKS_DN)
//	- NetBIOS name (KLHST_WKS_WINHOSTNAME)
//	- DNS name (KLHST_WKS_DNSNAME)
//
//	Parameters:
//	- pInfo	(params) the input container must contain variables:
//	- KLHST_WKS_ANYNAME (paramArray) array of strings with host names
//	- KLHST_WKS_GROUPID (int64) identifier of the group to which the designated hosts are to be placed
//
//	Return:
//	- pResults	(params) the output container will contain variables:
//	- KLHST_WKS_RESULTS (paramArray) array of bool values,
//	If the i-th element of this array is false,
//	then the i-th host of the input array KLHST_WKS_ANYNAME
//	could not be placed in the group (could not resolve name).
//
//Example:
//
//{
//	"pInfo": {
//		"KLHST_WKS_ANYNAME" : ["ip", "KLHST_WKS_HOSTNAME", "KLHST_WKS_DN", "KLHST_WKS_WINHOSTNAME", "KLHST_WKS_DNSNAME" ],
//		"KLHST_WKS_GROUPID" : 1 //GroupID
//		}
//	}
func (hg *HostGroup) ResolveAndMoveToGroup(ctx context.Context, params PInfoRaM) (*KlhstWksResults, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.ResolveAndMoveToGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	klhstWksResults := new(KlhstWksResults)
	raw, err := hg.client.Do(ctx, request, &klhstWksResults)
	return klhstWksResults, raw, err
}

//	Restarts specified network scanning type.
//
//	Parameters:
//	- nType	(int64) IN network scanning type:
//		|- 1 - AD
//		|- 2 - Ms network fast scanning
//		|- 3 - Ms network full scanning
//		|- 4 - Ip diapasons scanning
func (hg *HostGroup) RestartNetworkScanning(ctx context.Context, nType int64) (*PxgRetError, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nType": %d
	}`, nType))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RestartNetworkScanning", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgRetError := new(PxgRetError)
	raw, err := hg.client.Do(ctx, request, &pxgRetError)
	return pxgRetError, raw, err
}

//TODO SetLocInfo
//TODO SS_CreateSection
//TODO SS_DeleteSection

//	Get section names from host settings storage.
//
//	If product is empty then names will contain all product names. If product is not empty and version is empty then names will contain all versions for the specified product name. If product is not empty and version is not empty then names will contain all sections for the specified product and version.
//
//	Parameters:
//	- strHostName	(string) host name (unique server-generated string)
//	- strType	(string) type of storage (for example: "SS_SETTINGS")
//	- strProduct	(string) product name string, non-empty string, not longer than 31 character,
//	and cannot contain characters /\:*?"<>.
//	- strVersion	(string) version string, non-empty string, not longer than 31 character,
//	and cannot contain characters /\:*?"<>.
//
//	Example request params struct:
//
//	type SSParams struct {
//		StrHostName string `json:"strHostName"`
//		StrType     string `json:"strType"`
//		StrProduct  string `json:"strProduct"`
//		StrVersion  string `json:"strVersion"`
//		StrSection  string `json:"strSection"`
//	}
//
//	Returns:
//	- (array) array of strings with section names
//
//	See also:
//	Local settings and policy format for some products
func (hg *HostGroup) SS_GetNames(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.SS_GetNames", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//	Read data from host settings storage.
//
//	Parameters:
//	- params (SSParams)
//	|- strHostName	(string) host name (unique server-generated string)
//	|- strType	(string) type of storage, one of following:
//			|- specify "SS_SETTINGS" to access local settings (see also Local settings and policy format for some products)
//			|- specify "SS_PRODINFO" to access host system properties (see also Contents of host SS_PRODINFO storage)
//	|- strProduct	(string) product name string, non-empty string, not longer than 31 character,
//	and cannot contain characters /\:*?"<>.
//	|- strVersion	(string) version string, non-empty string, not longer than 31 character,
//	and cannot contain characters /\:*?"<>.
//	|- strSection	(string) section name string, non-empty string, not longer than 31 character,
//	and cannot contain characters /\:*?"<>.
//
//	Example request params struct:
//
//	type SSParams struct {
//		StrHostName string `json:"strHostName"`
//		StrType     string `json:"strType"`
//		StrProduct  string `json:"strProduct"`
//		StrVersion  string `json:"strVersion"`
//		StrSection  string `json:"strSection"`
//	}
//	Returns:
//	- (params) host settings
//
//	See also:
//	- Local settings and policy format for some products
//	- Contents of host SS_PRODINFO storage
func (hg *HostGroup) SS_Read(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.SS_Read", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//TODO SS_Write

//	Change attributes of existing administration group.
//
//	Parameters:
//	- nGroup	(int) id of the group
//	- pInfo	(params) container with group attributes.
//	May contain non-readonly attributes from the List of group attributes
func (hg *HostGroup) UpdateGroup(ctx context.Context, params interface{}) (*PxgValStr, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.UpdateGroup", bytes.NewBuffer(postData))
	pxgValStr := new(PxgValStr)

	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//	Change attributes of existing administration group.
//
//	Parameters:
//	- nGroup	(int) id of the group
//	- pInfo	(params) container with group attributes.
//	May contain non-readonly attributes from the List of group attributes
func (hg *HostGroup) UpdateHost(ctx context.Context, v interface{}) (*Accessor, []byte, error) {
	data, _ := json.Marshal(v)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.UpdateHost", bytes.NewBuffer(data))

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

//TODO UpdateHostsMultiple
//TODO UpdateIncident

//Zero virus count for hosts in group and all subgroups.
//
//Parameters:
//nParent	(int64) Id of group to start from
//[out]	strActionGuid	(string) id of asynchronous operation,
//to get status use AsyncActionStateChecker.CheckActionState, lStateCode "1" means OK and "0" means fail
func (hg *HostGroup) ZeroVirusCountForGroup(ctx context.Context, nParent int64) (*WActionGUID, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nParent": %d}`, nParent))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.ZeroVirusCountForGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := hg.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

// Zero virus count for specified hosts.
//
//	Parameters:
//	- pHostNames	(array) array of host names
//
//	Example request params struct:
//
//	type ZeroVirusCountForHostsParams struct {
//		PHostNames []string `json:"pHostNames"`
//	}
//
//	Return:
//		- strActionGuid	(string) id of asynchronous operation,
//
//	to get status use AsyncActionStateChecker.CheckActionState,
//	lStateCode "1" means OK and "0" means fail
func (hg *HostGroup) ZeroVirusCountForHosts(ctx context.Context, params interface{}) (*WActionGUID, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.ZeroVirusCountForHosts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := hg.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}
