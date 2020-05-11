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

//	LicensePolicy Class Reference
//
//	List of all members.
type LicensePolicy service

//
//	Get number of free licenses for functionality.
//
//	Parameters:
//	- nFunctionality	(int64) functionality
//
//	Returns:
//	- (int64) number of free licenses
//
//	Exceptions:
//	- Throws	exception in case of error.
//
//	+---------------+----------------------------------------------------------+
//	| Functionality |                       Description                        |
//	+---------------+----------------------------------------------------------+
//	|             1 | system managment                                         |
//	|             2 | mobile device managment                                  |
//	|             4 | PCLOUD is available via KL license                       |
//	|             5 | PCLOUD is available via Pay-per-use (Paid AMI) licensing |
//	+---------------+----------------------------------------------------------+
func (lp *LicensePolicy) GetFreeLicenseCount(ctx context.Context, nFunctionality int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nFunctionality": %d}`, nFunctionality))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.GetFreeLicenseCount", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lp.client.Do(ctx, request, nil)
	return raw, err
}

//	Get total number of licenses for functionality.
//
//	Parameters:
//	- nFunctionality	(int64) functionality
//
//	Returns:
//	 -(int64) total number of licenses
//
//	Exceptions:
//	- Throws	exception in case of error.
//
//	+---------------+----------------------------------------------------------+
//	| Functionality |                       Description                        |
//	+---------------+----------------------------------------------------------+
//	|             1 | system managment                                         |
//	|             2 | mobile device managment                                  |
//	|             4 | PCLOUD is available via KL license                       |
//	|             5 | PCLOUD is available via Pay-per-use (Paid AMI) licensing |
//	+---------------+----------------------------------------------------------+
func (lp *LicensePolicy) GetTotalLicenseCount(ctx context.Context, nFunctionality int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nFunctionality": %d}`, nFunctionality))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.GetTotalLicenseCount", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lp.client.Do(ctx, request, nil)
	return raw, err
}

//	Check a functionality in restricted mode.
//
//	Parameters:
//	 - nFunctionality	(unsignedInt) functionality for checking
//
//	Returns:
//	- (bool) true if a functionality is in restricted mode, otherwise false
//
//	Exceptions:
//	- Throws	exception in case of error.
//
//	+---------------+----------------------------------------------------------+
//	| Functionality |                       Description                        |
//	+---------------+----------------------------------------------------------+
//	|             1 | system managment                                         |
//	|             2 | mobile device managment                                  |
//	|             4 | PCLOUD is available via KL license                       |
//	|             5 | PCLOUD is available via Pay-per-use (Paid AMI) licensing |
//	+---------------+----------------------------------------------------------+
func (lp *LicensePolicy) IsLimitedMode(ctx context.Context, nFunctionality int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nFunctionality": %d}`, nFunctionality))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.IsLimitedMode", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lp.client.Do(ctx, request, nil)
	return raw, err
}

//	Enable or disable restricted mode for functionality.
//
//	Parameters:
//	- bLimited	(bool) true enable restricted mode for functionality, false for disable
//	- eFunctionality	(int) functionality
//
//	Exceptions:
//	- Throws	exception in case of error.
//
//	+---------------+----------------------------------------------------------+
//	| Functionality |                       Description                        |
//	+---------------+----------------------------------------------------------+
//	|             1 | system managment                                         |
//	|             2 | mobile device managment                                  |
//	|             4 | PCLOUD is available via KL license                       |
//	|             5 | PCLOUD is available via Pay-per-use (Paid AMI) licensing |
//	+---------------+----------------------------------------------------------+
func (lp *LicensePolicy) SetLimitedModeTest(ctx context.Context, bLimited bool, eFunctionality int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"bLimited": %v, "eFunctionality": %d}`, bLimited, eFunctionality))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.SetLimitedModeTest", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lp.client.Do(ctx, request, nil)
	return raw, err
}

//	Set total number of licenses for functionality in restricted mode.
//
//	Parameters:
//	- eFunctionality	(int) functionality
//	- nCount	(unsignedInt) total number of licenses
//
//	Exceptions:
//	- Throws	exception in case of error.
//
//	+---------------+----------------------------------------------------------+
//	| Functionality |                       Description                        |
//	+---------------+----------------------------------------------------------+
//	|             1 | system managment                                         |
//	|             2 | mobile device managment                                  |
//	|             4 | PCLOUD is available via KL license                       |
//	|             5 | PCLOUD is available via Pay-per-use (Paid AMI) licensing |
//	+---------------+----------------------------------------------------------+
func (lp *LicensePolicy) SetTotalLicenseCountTest(ctx context.Context, eFunctionality, nCount int64) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"eFunctionality": %d, "nCount": %d}`, eFunctionality, nCount))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.SetTotalLicenseCountTest", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lp.client.Do(ctx, request, nil)
	return raw, err
}

//	Set number of used licenses for functionality in restricted mode.
//
//	Parameters:
//	- eFunctionality	(int) functionality
//	- nCount	(unsignedInt) number of used licenses
//
//	Exceptions:
//	- Throws	exception in case of error.
//
//	+---------------+----------------------------------------------------------+
//	| Functionality |                       Description                        |
//	+---------------+----------------------------------------------------------+
//	|             1 | system managment                                         |
//	|             2 | mobile device managment                                  |
//	|             4 | PCLOUD is available via KL license                       |
//	|             5 | PCLOUD is available via Pay-per-use (Paid AMI) licensing |
//	+---------------+----------------------------------------------------------+
func (lp *LicensePolicy) SetUsedLicenseCountTest(ctx context.Context, eFunctionality, nCount int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"eFunctionality": %d, "nCount": %d}`, eFunctionality, nCount))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.SetUsedLicenseCountTest", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lp.client.Do(ctx, request, nil)
	return raw, err
}
