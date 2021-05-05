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
	"log"
	"net/http"
)

// SrvSsRevision service provide access to virtual server settings storage revisions.
type SrvSsRevision service

// SsRevisionOpen Open specified version of virtual server settings storage.
func (ssr *SrvSsRevision) SsRevisionOpen(ctx context.Context, nVServer, nRevision int64, szwType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nVServer": %d, "nRevision": %d, "szwType": "%s" }`, nVServer, nRevision, szwType))
	request, err := http.NewRequest("POST", ssr.client.Server+"/api/v1.0/SrvSsRevision.SsRevision_Open", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ssr.client.Do(ctx, request, nil)
	return raw, err
}

// SsRevisionClose Close settings storage opened by SrvSsRevision.SsRevisionOpen
func (ssr *SrvSsRevision) SsRevisionClose(ctx context.Context, szwType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{
		"szwType": "%s"
	}`, szwType))
	request, err := http.NewRequest("POST", ssr.client.Server+"/api/v1.0/SrvSsRevision.SsRevision_Close", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ssr.client.Do(ctx, request, nil)
	return raw, err
}

func (ssr *SrvSsRevision) SsRevisionGetNames(ctx context.Context, szwId, product, version string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwId": "%s", "product": ""%s", "version": "%s" }`, szwId, product, version))
	request, err := http.NewRequest("POST", ssr.client.Server+"/api/v1.0/SrvSsRevision.SsRevisionGetNames", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ssr.client.Do(ctx, request, nil)

	if err != nil {
		return nil, err
	}

	if ssr.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}
