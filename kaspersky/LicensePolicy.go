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

// LicensePolicy service by default, all functionalities are not available (IsLimitedMode returns true - functionality is limited).
// Functionality can be enabled in various ways.
// As a rule, functionality is enabled by certain parameters of license used.
// "Pay-per-Use Licensing (Paid AMI)" functionality is enabled if KSC is installed on a virtual machine with certain parameters in cloud infrastructure.
type LicensePolicy service

// GetFreeLicenseCount Get number of free licenses for functionality.
func (lp *LicensePolicy) GetFreeLicenseCount(ctx context.Context, nFunctionality int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nFunctionality": %d}`, nFunctionality))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.GetFreeLicenseCount", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lp.client.Request(ctx, request, nil)
	return raw, err
}

// GetTotalLicenseCount Get total number of licenses for functionality.
func (lp *LicensePolicy) GetTotalLicenseCount(ctx context.Context, nFunctionality int64) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nFunctionality": %d}`, nFunctionality))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.GetTotalLicenseCount", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := lp.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// IsLimitedMode Check a functionality in restricted mode.
func (lp *LicensePolicy) IsLimitedMode(ctx context.Context, nFunctionality int64) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nFunctionality": %d}`, nFunctionality))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.IsLimitedMode", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := lp.client.Request(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// SetLimitedModeTest Enable or disable restricted mode for functionality.
func (lp *LicensePolicy) SetLimitedModeTest(ctx context.Context, bLimited bool, eFunctionality int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"bLimited": %v, "eFunctionality": %d}`, bLimited, eFunctionality))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.SetLimitedModeTest", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lp.client.Request(ctx, request, nil)
	return raw, err
}

// SetTotalLicenseCountTest Set total number of licenses for functionality in restricted mode.
func (lp *LicensePolicy) SetTotalLicenseCountTest(ctx context.Context, eFunctionality, nCount int64) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"eFunctionality": %d, "nCount": %d}`, eFunctionality, nCount))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.SetTotalLicenseCountTest", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lp.client.Request(ctx, request, nil)
	return raw, err
}

// SetUsedLicenseCountTest Set number of used licenses for functionality in restricted mode.
func (lp *LicensePolicy) SetUsedLicenseCountTest(ctx context.Context, eFunctionality, nCount int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"eFunctionality": %d, "nCount": %d}`, eFunctionality, nCount))
	request, err := http.NewRequest("POST", lp.client.Server+"/api/v1.0/LicensePolicy.SetUsedLicenseCountTest", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lp.client.Request(ctx, request, nil)
	return raw, err
}
