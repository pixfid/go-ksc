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

// SsRevisionGetNames Product backend integration.
//
// This interface allows to set up and remove integration with backend product.
// Product backend integration is virtual server specific.
// If integration is done on main server, then it applies to all virtual servers except ones with their own integration.
// Service console must be deployed before this interface can be used. When product integration is set,
// KSC server automatically refreshes integration token until integration is removed.
type SsRevisionGetNames service

type IntegrationToken struct {
	WstrProdName    string `json:"wstrProdName"`
	WstrProdVersion string `json:"wstrProdVersion"`
	NTokenTtl       int    `json:"nTokenTtl"`
	PRefreshToken   struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"pRefreshToken"`
}

// SetIntegrationToken Set up integration with product backend.
func (srn *SsRevisionGetNames) SetIntegrationToken(ctx context.Context, params IntegrationToken) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", srn.client.Server+"/api/v1.0/SsRevisionGetNames.SetIntegrationToken", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := srn.client.Request(ctx, request, nil)

	if srn.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}

// DeleteIntegrationToken Remove integration with product backend.
//
// Before removing integration KSC server calls backend command to delete all user tokens issued within this integration.
func (srn *SsRevisionGetNames) DeleteIntegrationToken(ctx context.Context, wstrProdName, wstrProdVersion string) error {
	postData := []byte(fmt.Sprintf(`{"wstrProdName": "%s", "wstrProdVersion": "%s"}`, wstrProdName, wstrProdVersion))
	request, err := http.NewRequest("POST", srn.client.Server+"/api/v1.0/SsRevisionGetNames.DeleteIntegrationToken",
		bytes.NewBuffer(postData))

	if err != nil {
		return err
	}

	raw, err := srn.client.Request(ctx, request, nil)

	if srn.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}
