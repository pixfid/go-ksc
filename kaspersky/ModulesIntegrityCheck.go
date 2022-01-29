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

// ModulesIntegrityCheck Interface for modules integrity check.
type ModulesIntegrityCheck service

type IntegrityCheckInfo struct {
	PxgRetVal struct {
		IsInvoked  bool   `json:"IsInvoked"`
		ResultCode int    `json:"ResultCode"`
		ResultText string `json:"ResultText"`
	} `json:"PxgRetVal"`
}

// GetIntegrityCheckInfo Returns integrity check info.
func (mic *ModulesIntegrityCheck) GetIntegrityCheckInfo(ctx context.Context) (*IntegrityCheckInfo, error) {
	request, err := http.NewRequest("POST", mic.client.Server+"/api/v1.0/ModulesIntegrityCheck.GetIntegrityCheckInfo", nil)
	if err != nil {
		return nil, err
	}

	integrityCheckInfo := new(IntegrityCheckInfo)
	raw, err := mic.client.Request(ctx, request, &integrityCheckInfo)

	if mic.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return integrityCheckInfo, err
}
