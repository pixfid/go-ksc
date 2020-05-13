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
	"fmt"
	"net/http"
)

//	KsnInternal Class Reference
//
//	Interface for working with KsnProxy subsystem. More...
//
//	This interface allow you to set KPSN settings, check license and check connection.
//
//	List of all members.
type KsnInternal service

//	Check connection with KSN cloud (or KPSN)
//
//	Returns:
//	- (bool) Returns true if connection checked. Otherwise - false.
//
//	Exceptions:
//	- KLSTD.STDE_NOACCESS	- Access denied
//	- KLSTD.STDE_NOTPERM - KsnProxy is disabled,
//	- KLPRCP.ERR_CANT_CONNECT - Can not connect to KSN.
func (sd *KsnInternal) CheckKsnConnection(ctx context.Context) (*PxgValBool, []byte, error) {
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/KsnInternal.CheckKsnConnection", nil)

	pxgValBool := new(PxgValBool)
	raw, err := sd.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//	Get all KPSN eula.
//
//	Return:
//	- pNKsnEulas (array) Array of eula.
//	See Format of KPSN eula params.
//
//	Exceptions:
//	- KLSTD.STDE_NOTPERM	- Can't call on virtual server,
//	- KLSTD.STDE_NOACCESS - Access denied.
func (sd *KsnInternal) GetNKsnEulas(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/KsnInternal.GetNKsnEulas", nil)

	raw, err := sd.client.Do(ctx, request, nil)
	return raw, err
}

// KsnSettings is returned by GetSettings.
type KsnSettings struct {
	PSettings *PSettings `json:"pSettings,omitempty"`
}

type PSettings struct {
	KsnproxyEnable                 *bool  `json:"KSNPROXY_ENABLE,omitempty"`
	KsnproxyEnablePatchManStat     *bool  `json:"KSNPROXY_ENABLE_PATCH_MAN_STAT,omitempty"`
	KsnproxyKsnType                *int64 `json:"KSNPROXY_KSN_TYPE,omitempty"`
	KsnproxyTCPPort                *int64 `json:"KSNPROXY_TCP_PORT,omitempty"`
	KsnproxyUDPPort                *int64 `json:"KSNPROXY_UDP_PORT,omitempty"`
	KsnproxyUseMasterKsnproxyAsKsn *bool  `json:"KSNPROXY_USE_MASTER_KSNPROXY_AS_KSN,omitempty"`
}

//	Returns settings of KsnProxy. May be used on virtual server.
//
//	Parameters:
//	- pSettings (params) Section KSNPROXY_SETTINGS.
//	See Section KSNPROXY_SETTINGS attributes.
func (sd *KsnInternal) GetSettings(ctx context.Context) (*KsnSettings, []byte, error) {
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/KsnInternal.GetSettings", nil)

	ksnSettings := new(KsnSettings)
	raw, err := sd.client.Do(ctx, request, &ksnSettings)
	return ksnSettings, raw, err
}

//	Check possibility to send statistics.
//
//	Returns:
//	- (bool) Returns true when possible to send statistics.
//	Otherwise - false.
func (sd *KsnInternal) NeedToSendStatistics(ctx context.Context) (*PxgValBool, []byte, error) {
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/KsnInternal.NeedToSendStatistics", nil)

	pxgValBool := new(PxgValBool)
	raw, err := sd.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//	Get KPSN eula.
//
//	Parameters:
//	- wstrNKsnLoc	(string) Localization ('ru', 'en')
//
//	Return:
//	- pEula	(params) Params with EULA text and localization. See Format of KPSN eula params.
//	Only NKsnEula and NKsnEulaLoc present.
//
//	Exceptions:
//	- KLSTD.STDE_NOTPERM	- Can't call on virtual server
//	- KLSTD.STDE_NOACCESS - Access denied,
//	- KLSTD.STDE_NOTFOUND - Eula with specified localization not found.
func (sd *KsnInternal) GetNKsnEula(ctx context.Context, wstrNKsnLoc string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrNKsnLoc": "%s"}`, wstrNKsnLoc))
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/KsnInternal.GetNKsnEula", bytes.NewBuffer(postData))

	raw, err := sd.client.Do(ctx, request, nil)
	return raw, err
}
