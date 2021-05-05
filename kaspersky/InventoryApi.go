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
	"log"
	"net/http"
)

// InventoryAPI service for working with Software Inventory subsystem.
//
// Service allows to get information about software applications that's are installed on client hosts
// and modify some settings for Software Inventory subsystem.
//
// To get additional information you also can use SrvView (InvSrvViewName)
type InventoryAPI service

type HostProducts struct {
	PxgRetVal struct {
		GNRLEAPARAM1 []struct {
			Type  string `json:"type"`
			Value struct {
				ARPRegKey          string `json:"ARPRegKey"`
				CleanerProductName string `json:"CleanerProductName"`
				Comments           string `json:"Comments"`
				DisplayName        string `json:"DisplayName"`
				DisplayVersion     string `json:"DisplayVersion"`
				HelpLink           string `json:"HelpLink"`
				HelpTelephone      string `json:"HelpTelephone"`
				InstallDate        string `json:"InstallDate"`
				InstallDir         string `json:"InstallDir"`
				InstanceID         struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"InstanceID"`
				LangId               int    `json:"LangId"`
				PackageCode          string `json:"PackageCode"`
				ProductID            string `json:"ProductID"`
				Publisher            string `json:"Publisher"`
				QuietUninstallString string `json:"QuietUninstallString"`
				UninstallString      string `json:"UninstallString"`
				VapmBuild            struct {
					Type  string `json:"type"`
					Value int    `json:"value"`
				} `json:"VapmBuild"`
				BIsMsi bool `json:"bIsMsi"`
			} `json:"value"`
		} `json:"GNRL_EA_PARAM_1"`
	} `json:"PxgRetVal"`
}

// GetHostInvProducts Acquire all software applications.
func (ia *InventoryAPI) GetHostInvProducts(ctx context.Context, szwHostID string) (*HostProducts, error) {
	postData := []byte(fmt.Sprintf(`{"szwHostId": "%s"}`, szwHostID))
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetHostInvProducts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	result := new(HostProducts)
	raw, err := ia.client.Do(ctx, request, &result)

	if ia.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

// GetHostInvPatches Acquire software application updates which are installed on specified host.
func (ia *InventoryAPI) GetHostInvPatches(ctx context.Context, szwHostID string) (*InvPatches, error) {
	postData := []byte(fmt.Sprintf(`{"szwHostId": "%s"}`, szwHostID))
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetHostInvPatches", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	result := new(InvPatches)
	raw, err := ia.client.Do(ctx, request, &result)

	if ia.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

type InvPatches struct {
	PxgRetVal struct {
		GNRLEAPARAM1 []struct {
			Type  string `json:"type"`
			Value struct {
				Classification       string `json:"Classification"`
				Comments             string `json:"Comments"`
				DisplayName          string `json:"DisplayName"`
				DisplayVersion       string `json:"DisplayVersion"`
				HelpLink             string `json:"HelpLink"`
				HelpTelephone        string `json:"HelpTelephone"`
				InstallDate          string `json:"InstallDate"`
				MoreInfoURL          string `json:"MoreInfoURL"`
				ParentID             string `json:"ParentID"`
				PatchID              string `json:"PatchID"`
				Publisher            string `json:"Publisher"`
				QuietUninstallString string `json:"QuietUninstallString"`
				UninstallString      string `json:"UninstallString"`
				BIsMsi               bool   `json:"bIsMsi"`
			} `json:"value"`
		} `json:"GNRL_EA_PARAM_1"`
	} `json:"PxgRetVal"`
}

// GetInvPatchesList Acquire all software application updates.
func (ia *InventoryAPI) GetInvPatchesList(ctx context.Context, params Null) (*InvPatches, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetInvPatchesList", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	result := new(InvPatches)
	raw, err := ia.client.Do(ctx, request, &result)

	if ia.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

type InvProducts struct {
	PxgRetVal struct {
		GNRLEAPARAM1 []struct {
			Type  string `json:"type"`
			Value struct {
				ARPRegKey            string `json:"ARPRegKey"`
				CleanerProductName   string `json:"CleanerProductName"`
				Comments             string `json:"Comments"`
				DisplayName          string `json:"DisplayName"`
				DisplayVersion       string `json:"DisplayVersion"`
				HelpLink             string `json:"HelpLink"`
				HelpTelephone        string `json:"HelpTelephone"`
				InstallDate          string `json:"InstallDate"`
				InstallDir           string `json:"InstallDir"`
				LangId               int    `json:"LangId"`
				PackageCode          string `json:"PackageCode"`
				ProductID            string `json:"ProductID"`
				Publisher            string `json:"Publisher"`
				QuietUninstallString string `json:"QuietUninstallString"`
				UninstallString      string `json:"UninstallString"`
				VapmBuild            struct {
					Type  string `json:"type"`
					Value int    `json:"value"`
				} `json:"VapmBuild"`
				BIsMsi bool `json:"bIsMsi"`
			} `json:"value"`
		} `json:"GNRL_EA_PARAM_1"`
	} `json:"PxgRetVal"`
}

// GetInvProductsList Acquire all software applications.
func (ia *InventoryAPI) GetInvProductsList(ctx context.Context, params Null) (*InvProducts, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetInvProductsList", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	result := new(InvProducts)
	raw, err := ia.client.Do(ctx, request, &result)

	if ia.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

// DeleteUninstalledApps Remove from database info about software applications which aren't installed on any host.
func (ia *InventoryAPI) DeleteUninstalledApps(ctx context.Context) error {
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.DeleteUninstalledApps", nil)
	if err != nil {
		return err
	}

	raw, err := ia.client.Do(ctx, request, nil)

	if ia.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

// GetSrvCompetitorIniFileInfoList Acquire info about all cleaner ini-files of specified type from SC-server.
// Returns info about cleaner ini-files of specified type from SC-server.
// These files are used to detect and uninstall applications which incompatible with KasperskyLab antivirus applications
func (ia *InventoryAPI) GetSrvCompetitorIniFileInfoList(ctx context.Context, wstrType string) (*PxgValCIFIL, error) {
	postData := []byte(fmt.Sprintf(`{"wstrType": "%s"}`, wstrType))
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetSrvCompetitorIniFileInfoList", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	result := new(PxgValCIFIL)
	raw, err := ia.client.Do(ctx, request, &result)

	if ia.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

// GetObservedApps Acquire list of observed applications.
func (ia *InventoryAPI) GetObservedApps(ctx context.Context, params Null) (*PxgValArrayOfString, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetObservedApps", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	result := new(PxgValArrayOfString)
	raw, err := ia.client.Do(ctx, request, &result)

	if ia.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

// ObservedAppsParams struct using in InventoryApi.SetObservedApps
type ObservedAppsParams struct {
	// PAppIDS collection of (string) application string Id.
	PAppIDS []string `json:"pAppIds"`
	// PParams reserved. (params)
	PParams Null `json:"pParams"`
}

// SetObservedApps Set list of observed applications.
func (ia *InventoryAPI) SetObservedApps(ctx context.Context, params ObservedAppsParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.SetObservedApps", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, nil)

	if ia.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}
