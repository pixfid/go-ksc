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

// OAuth2 product managing.
type OAuth2 service

//TODO ApproveNewClients
//TODO ApproveNewResServers
//TODO ConfirmClientsRegistration
//TODO ConfirmResServersRegistration
//TODO DeclineClientsRegistration
//TODO DeclineNewClients
//TODO DeclineNewResServers
//TODO DeclineResServersRegistration
//TODO DeleteClients
//TODO DeleteResServers

// GetClients Retrieves OAuth2 clients information.
//	nFilterByState	(int64) filters the result by client state (see OAuth2 items states),
//	if set to -1 the result will not be filtered
func (oa *OAuth2) GetClients(ctx context.Context, nFilterByState int64) (*IssuanceSetting, error) {
	postData := []byte(fmt.Sprintf(`{"nFilterByState": %d}`, nFilterByState))
	request, err := http.NewRequest("POST", oa.client.Server+"/api/v1.0/OAuth2.GetClients",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	issuanceSetting := new(IssuanceSetting)
	raw, err := oa.client.Request(ctx, request, &issuanceSetting)

	if oa.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return issuanceSetting, err
}

// GetClientsToRegistration Retrieves clients for registration them on the OAuth2 server.
func (oa *OAuth2) GetClientsToRegistration(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", oa.client.Server+"/api/v1.0/OAuth2.GetClientsToRegistration", nil)
	if err != nil {
		return nil, err
	}

	raw, err := oa.client.Request(ctx, request, nil)

	if oa.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}

// GetNewClients Retrieves new OAuth2 clients for admin approval.
func (oa *OAuth2) GetNewClients(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", oa.client.Server+"/api/v1.0/OAuth2.GetNewClients", nil)
	if err != nil {
		return nil, err
	}

	raw, err := oa.client.Request(ctx, request, nil)

	if oa.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}

// GetNewResServers Retrieves new OAuth2 resource servers for admin approval.
func (oa *OAuth2) GetNewResServers(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", oa.client.Server+"/api/v1.0/OAuth2.GetNewResServers", nil)
	if err != nil {
		return nil, err
	}

	raw, err := oa.client.Request(ctx, request, nil)

	if oa.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}

// GetResServers Retrieves OAuth2 resource servers information.
//	nFilterByState	(int64) filters the result by resource server state (see OAuth2 items states),
//	if set to -1 the result will not be filtered
func (oa *OAuth2) GetResServers(ctx context.Context, nFilterByState int64) (*IssuanceSetting, error) {
	postData := []byte(fmt.Sprintf(`{"nFilterByState": %d}`, nFilterByState))
	request, err := http.NewRequest("POST", oa.client.Server+"/api/v1.0/OAuth2.GetResServers",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	issuanceSetting := new(IssuanceSetting)
	raw, err := oa.client.Request(ctx, request, &issuanceSetting)

	if oa.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return issuanceSetting, err
}

// GetResServersToRegistration Retrieves OAuth2 resource servers for registration them on the OAuth2 server.
func (oa *OAuth2) GetResServersToRegistration(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", oa.client.Server+"/api/v1.0/OAuth2.GetResServersToRegistration", nil)
	if err != nil {
		return nil, err
	}

	raw, err := oa.client.Request(ctx, request, nil)

	if oa.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}

//TODO UpdateClients
//TODO UpdateResServers
