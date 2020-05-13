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

//	LicenseInfoSync Class Reference
//
//	Operating with licenses
//
//	List of all members.
type LicenseInfoSync service

//TODO AcquireKeysForProductOnHost
//TODO GetKeyDataForHost
//TODO IsLicForSaasValid2

//	Check whether the key's product id belongs to the Public Cloud product ids list.
//
//	Parameters:
//	- nProductId	(int64) Product ID, mandatory.
//
//	Returns:
//	- (bool) true if product id belongs to the Public Cloud product ids list, false otherwise.
//
//	Exceptions:
//	Throws	exception in case of error.
func (cp *LicenseInfoSync) IsPCloudKey(ctx context.Context, nProductId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nProductId": %d}`, nProductId))
	request, err := http.NewRequest("POST", cp.client.Server+"/api/v1.0/LicenseInfoSync.IsPCloudKey",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := cp.client.Do(ctx, request, nil)
	return raw, err
}

//	Force synchronization of subscription licenses' metadata.
//
//	Returns:
//	- (string) Request ID used to subscribe to the event that is triggered when operation is complete.
//
//	Remarks:
//	Check the operation state by calling AsyncActionStateChecker::CheckActionState periodically until it's finalized
//	or an exception is thrown.
//
//	Exceptions:
//	Throws	exception in case of error.
func (gs *LicenseInfoSync) SynchronizeLicInfo2(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", gs.client.Server+"/api/v1.0/LicenseInfoSync.SynchronizeLicInfo2",
		nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := gs.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//TODO TryToInstallLicForSaas2

//	Uninstall adm. server's license.
//
//	Parameters:
//	- bCurrent	(bool) true to install active license, otherwise uninstall the reserved one.
//
//	Exceptions:
//	Throws	exception in case of error.
func (cp *LicenseInfoSync) TryToUnistallLicense(ctx context.Context, bCurrent bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"bCurrent": %v}`, bCurrent))
	request, err := http.NewRequest("POST", cp.client.Server+"/api/v1.0/LicenseInfoSync.TryToUnistallLicense",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := cp.client.Do(ctx, request, nil)
	return raw, err
}
