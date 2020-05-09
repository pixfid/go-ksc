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
	"net/http"
)

//	AdHosts Class Reference
//
//	Scanned active directory OU structure.
//
//	Allows to enumerate scanned active directory OU structure.
//
//	List of all members:
type AdHosts struct {
	client *Client
}

//FindAdGroupsParams struct
type FindAdGroupsParams struct {
	VecFieldsToReturn []string        `json:"vecFieldsToReturn,omitempty"`
	VecFieldsToOrder  []FieldsToOrder `json:"vecFieldsToOrder,empty"`
	POptions          POptions        `json:"pOptions,omitempty"`
	LMaxLifeTime      int64           `json:"lMaxLifeTime,omitempty"`
}

type POptions struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value"`
}

//	ADHostIterator struct
type ADHostIterator struct {
	WstrIterator string `json:"wstrIterator"`
	PxgRetVal    int64  `json:"PxgRetVal"`
}

//Enumerates AD groups.
//
//	- Parameters:
//	- Example:
//		VecFieldsToReturn: []string{
//			"adhst_adgroup_id",
//			"adhst_adgroup_distinguished_name",
//			"adhst_adgroup_name",
//			"adhst_adgroup_sam_name",
//		},
//		FieldsToOrder: []kaspersky.FieldsToOrder{
//			{
//				Type: "params",
//				Value: kaspersky.Value{
//					Asc:  true,
//					Name: "adhst_adgroup_id",
//				},
//			},
//		},
//		POptions: kaspersky.POptions{
//			Type:  "params",
//			Value: "adhst_id",
//		},
//		LMaxLifeTime: 100,
//	}
//	- Where:
//	- vecFieldsToReturn	(array) attributes to return, possible values
//	"adhst_adgroup_id" (int64) AD group id
//	"adhst_adgroup_distinguished_name" (string) AD group unique name (distinguished name)
//	"adhst_adgroup_name" (string) AD group short name (may be non-unique)
//	"adhst_adgroup_sam_name" (string) AD group NT4-compatible name (unique but may be absent)
//
//	- vecFieldsToOrder	(array) array of containers each of them containing two attributes :
//	"Name" (string) name of attributes used for sorting
//	"Asc" (paramBool) ascending if true descending otherwise
//
//	- pOptions	(params) options, possible attributes:
//	"adhst_id" - return groups for given ad host
//
//	- lMaxLifeTime	(int64) max result-set lifetime in seconds, not more than 7200
//	- [out]	wstrIterator(string) result-set ID, identifier of the server-side ordered collection.
//The result-set is destroyed and associated memory is freed in following cases:
//
//Passed lMaxLifeTime seconds after last access to the result-set (by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//Session to the Administration Server has been closed.
//ChunkAccessor.Release has been called.
func (ah *AdHosts) FindAdGroups(ctx context.Context, params FindAdGroupsParams) (*ADHostIterator, []byte, error) {
	postData, _ := json.Marshal(params)

	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.FindAdGroups", bytes.NewBuffer(postData))

	aDHostIterator := new(ADHostIterator)

	raw, err := ah.client.Do(ctx, request, &aDHostIterator)
	return aDHostIterator, raw, err
}

type ChildComputersParams struct {
	IDOU              int64    `json:"idOU"`
	VecFieldsToReturn []string `json:"vecFieldsToReturn"`
	LMaxLifeTime      int64    `json:"lMaxLifeTime"`
}

//ChildComputerParams struct
type ChildComputerParams struct {
	IDAdhst           int64    `json:"idAdhst,omitempty"`
	VecFieldsToReturn []string `json:"vecFieldsToReturn,omitempty"`
}

//AdHstIDParent struct contain AD host attributes.
type AdHstIDParent struct {
	PxgRetVal AdHstIDParentPxgRetVal `json:"PxgRetVal"`
}

type AdHstIDParentPxgRetVal struct {
	AdhstIDParent int64 `json:"adhst_idParent"`
}

//	Retrieves AD host attributes.
//
//	Parameters:
//	- idAdhst	(int64) host identifier (same as GetChildComputers attribute adhst_id)
//	- vecFieldsToReturn	(array) Array of propery names.
//	See List of host attributes and List of attributes of organization units for attribute names.
func (ah *AdHosts) GetChildComputer(ctx context.Context, params ChildComputerParams) (*AdHstIDParent,
	[]byte, error) {
	postData, _ := json.Marshal(params)

	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.GetChildComputer", bytes.NewBuffer(postData))

	adHstIdParent := new(AdHstIDParent)
	raw, err := ah.client.Do(ctx, request, &adHstIdParent)
	return adHstIdParent, raw, err
}

//Acquire list of hosts located in "Unassigned computers" for specified OU.
//
//Returns list of hosts located in "Unassigned computers" for specified organization unit.
//
//	- Parameters:
//	- Example:
//kaspersky.ChildComputersParams{
//		IDOU:              508,
//		VecFieldsToReturn: []string{
//			"adhst_idParent",
//			"adhst_idComputer",
//			"adhst_adgroups",
//			"KLHST_WKS_HOSTNAME",
//			"KLHST_WKS_LID",
//			"KLHST_WKS_GROUPID",
//			"KLHST_WKS_LAST_VISIBLE",
//			"KLHST_WKS_LAST_INFOUDATE",
//			"KLHST_WKS_STATUS",
//			"KLHST_WKS_LAST_UPDATE",
//			"KLHST_WKS_LAST_NAGENT_CONNECTED",
//			"KLHST_WKS_DN",
//			"KLHST_WKS_WINHOSTNAME",
//			"KLHST_WKS_DNSNAME",
//			"KLHST_WKS_WINDOMAIN",
//			"KLHST_WKS_WINDOMAIN_TYPE",
//			"KLHST_WKS_DNSDOMAIN",
//			"KLHST_WKS_CTYPE",
//			"KLHST_WKS_PTYPE",
//			"KLHST_WKS_OS_VER_MAJOR",
//			"KLHST_WKS_OS_VER_MINOR",
//			"KLHST_WKS_OSSP_VER_MAJOR",
//			"KLHST_WKS_OSSP_VER_MINOR",
//			"KLHST_WKS_CPU_ARCH",
//			"KLHST_WKS_IP_LONG",
//			"KLHST_WKS_STATUS_ID",
//			"KLHST_WKS_STATUS_MASK",
//			"KLHST_WKS_STATUS_HSDP"},
//
//		LMaxLifeTime:      100,
//	}
//
//	- Where:
//	idOU	(int64) id of organization unit
//	vecFieldsToReturn	(array) fields names to acquire, following names may be specified:
//	lMaxLifeTime	(int64) max result-set lifetime in seconds, not more than 7200
//
//Returns:
//	- (string) result-set ID, identifier of the server-side ordered collection.
//
//	The result-set is destroyed and associated memory is freed in following cases:
//	Passed lMaxLifeTime seconds after last access to the result-set (by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//	Session to the Administration Server has been closed.
//	ChunkAccessor.Release has been called.
func (ah *AdHosts) GetChildComputers(ctx context.Context, params ChildComputersParams) (*PxgValStr, []byte, error) {

	postData, _ := json.Marshal(params)

	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.GetChildComputers", bytes.NewBuffer(postData))

	pxgValStr := new(PxgValStr)

	raw, err := ah.client.Do(ctx, request, &pxgValStr)

	return pxgValStr, raw, err
}

//ChildOUParams struct
type ChildOUParams struct {
	IDOU         int64    `json:"idOU,omitempty"`
	PFields      []string `json:"pFields,omitempty"`
	LMaxLifeTime int64    `json:"lMaxLifeTime,omitempty"`
}

//Acquire list of child OUs for specified OU.
//
//Returns list of child organization units for specified organization unit
//
//Parameters:
//	- type ChildOUParams struct {
//	-   IDOU         int64
//	-   PFields      string
//	-   LMaxLifeTime int64
//-   }
//
//	- idOU	(int64) id of organization unit (or 0 to acquire root of hierarchy)
//	- pFields	([]string) fields names to acquire, following names may be specified:
//
//	- adhst_id
//	- adhst_idParent
//	- adhst_idComputer
//	- adhst_binOu
//
//	- lMaxLifeTime	(int64) max result-set lifetime in seconds, not more than 7200
//Returns:
//(string) result-set ID, identifier of the server-side ordered collection. The result-set is destroyed and associated memory is freed in following cases:
//Passed lMaxLifeTime seconds after last access to the result-set (by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//Session to the Administration Server has been closed.
//ChunkAccessor.Release has been called.
func (ah *AdHosts) GetChildOUs(ctx context.Context, params ChildOUParams) (*PxgValStr,
	[]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.GetChildOUs", bytes.NewBuffer(postData))
	pxgValStr := new(PxgValStr)
	raw, err := ah.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

type OUAttributesParams struct {
	IDOU    int64    `json:"idOU,omitempty"`
	PFields []string `json:"pFields"`
}

//OUAttributes struct
type OUAttributes struct {
	Attributes Attributes `json:"PxgRetVal"`
}

type Attributes struct {
	AdhstBinOu            AdhstBinOu `json:"adhst_binOu"`
	AdhstChildSubunitsNum int64      `json:"adhst_childSubunitsNum"`
	AdhstEnableAdScan     bool       `json:"adhst_enable_ad_scan"`
	AdhstHostsNum         int64      `json:"adhst_hostsNum"`
	AdhstID               int64      `json:"adhst_id"`
	AdhstIDComputer       string     `json:"adhst_idComputer"`
	AdhstIDParent         int64      `json:"adhst_idParent"`
}

type AdhstBinOu struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

//Returns attributes of specified OU
//
//	- Parameters: OUAttributesParams
// Examples of struct field tags and their meanings:
//   // idOU	(int64) id of organization unit
//   Field	IDOU    int64   `json:"idOU,omitempty"`
//
//   // pFields	([]string) fields names to acquire (see List of attributes of organization units for attribute names).
//   Field	PFields []string `json:"pFields"`
//
//	- Example:
//	kaspersky.OUAttributesParams{
//		IDOU:    10449,
//		PFields: []string{
//			"adhst_id",
//			"adhst_idParent",
//			"adhst_binOu",
//			"adhst_idComputer",
//			"adhst_hostsNum",
//			"adhst_childSubunitsNum",
//			"adhst_enable_ad_scan",
//		},
//	}
//	- Returns:
//   (params) object containing specified attributes
func (ah *AdHosts) GetOU(ctx context.Context, params OUAttributesParams) (*OUAttributes, []byte, error) {
	postData, _ := json.Marshal(params)

	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.GetOU", bytes.NewBuffer(postData))
	oUAttributes := new(OUAttributes)
	raw, err := ah.client.Do(ctx, request, &oUAttributes)
	return oUAttributes, raw, err
}

//	Updates OU properties.
//
//	Parameters:
//	- idOU	(int64) id of organization unit
//	- pData	(params) may contain following values:
//
//	"adhst_enable_ad_scan" (see Active Directory-specific attributes for organization units and computers.)
func (ah *AdHosts) UpdateOU(ctx context.Context, idOU int, params interface{}) ([]byte, error) {
	//postData, _ := json.Marshal(params)
	//TODO Find correct request
	postData := []byte(fmt.Sprintf(`
	{
		"idOU":   %d,
		"pData": {
			"type" : "params",
			"value" : {
				"adhst_enable_ad_scan" : 1 
			}
		}
	}`, idOU))

	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.UpdateOU", bytes.NewBuffer(postData))
	raw, err := ah.client.Do(ctx, request, nil)

	return raw, err
}
