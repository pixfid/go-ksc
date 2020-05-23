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
func (lis *LicenseInfoSync) IsPCloudKey(ctx context.Context, nProductId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nProductId": %d}`, nProductId))
	request, err := http.NewRequest("POST", lis.client.Server+"/api/v1.0/LicenseInfoSync.IsPCloudKey",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lis.client.Do(ctx, request, nil)
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
func (lis *LicenseInfoSync) SynchronizeLicInfo2(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", lis.client.Server+"/api/v1.0/LicenseInfoSync.SynchronizeLicInfo2",
		nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := lis.client.Do(ctx, request, &pxgValStr)
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
func (lis *LicenseInfoSync) TryToUnistallLicense(ctx context.Context, bCurrent bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"bCurrent": %v}`, bCurrent))
	request, err := http.NewRequest("POST", lis.client.Server+"/api/v1.0/LicenseInfoSync.TryToUnistallLicense",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lis.client.Do(ctx, request, nil)
	return raw, err
}
