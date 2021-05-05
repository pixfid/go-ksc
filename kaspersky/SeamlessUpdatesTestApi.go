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
	"context"
	"log"
	"net/http"
)

type SeamlessUpdatesTestApi service

type RequiredPlugins struct {
	Pxgretval []struct {
		Type  string `json:"type"`
		Value struct {
			Nrequiredpluginversion struct {
				Type  string `json:"type"`
				Value int64  `json:"value"`
			} `json:"nRequiredPluginVersion"`
			Wstrproductname    string `json:"wstrProductName"`
			Wstrproductversion string `json:"wstrProductVersion"`
		} `json:"value"`
	} `json:"PxgRetVal"`
}

func (suta *SeamlessUpdatesTestApi) GetRequiredPlugins(ctx context.Context) (*RequiredPlugins, error) {
	request, err := http.NewRequest("POST", suta.client.Server+"/api/v1.0/SeamlessUpdatesTestApi.GetRequiredPlugins", nil)
	if err != nil {
		return nil, err
	}

	result := new(RequiredPlugins)
	raw, err := suta.client.Do(ctx, request, &result)

	if suta.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

type VapmKlUpdatesToApprove struct {
	Pxgretval []struct {
		Type  string `json:"type"`
		Value struct {
			Bkleulastoaccept           bool `json:"bKlEulasToAccept"`
			Bklmobileeulastoaccept     bool `json:"bKlMobileEulasToAccept"`
			Bklpatchestoapprove        bool `json:"bKlPatchesToApprove"`
			Bksnagreementstoaccept     bool `json:"bKsnAgreementsToAccept"`
			Brevokedpatchesnotdeclined bool `json:"bRevokedPatchesNotDeclined"`
		} `json:"value"`
	} `json:"PxgRetVal"`
}

func (suta *SeamlessUpdatesTestApi) GetVapmKlUpdatesToApprove(ctx context.Context) (*VapmKlUpdatesToApprove, error) {
	request, err := http.NewRequest("POST", suta.client.Server+"/api/v1.0/SeamlessUpdatesTestApi.GetVapmKlUpdatesToApprove", nil)
	if err != nil {
		return nil, err
	}

	result := new(VapmKlUpdatesToApprove)
	raw, err := suta.client.Do(ctx, request, &result)

	if suta.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

func (suta *SeamlessUpdatesTestApi) CleanupSeamlessUpdates(ctx context.Context) (*LoggedInUsing2FA, error) {
	request, err := http.NewRequest("POST", suta.client.Server+"/api/v1.0/SeamlessUpdatesTestApi.CleanupSeamlessUpdates", nil)
	if err != nil {
		return nil, err
	}

	result := new(LoggedInUsing2FA)
	raw, err := suta.client.Do(ctx, request, &result)

	if suta.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}
