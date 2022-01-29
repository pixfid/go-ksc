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
	"net/http"
)

// Session management service.
type Session service

// CreateToken Creates session token.
//
// Creates session token for current security context.
//
// Those token can be used for logon purposes to Administaration Server for a short time (3 minutes by default).
func (s *Session) CreateToken(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", s.client.Server+"/api/v1.0/Session.CreateToken", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := s.client.Request(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// Ping Does nothing. May be used to effectively verify session validity.
// Session id to verify is passed in "X-KSC-Session" header.
func (s *Session) Ping(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", s.client.Server+"/api/v1.0/Session.Ping", nil)
	if err != nil {
		return nil, err
	}

	raw, err := s.client.Request(ctx, request, nil)
	return raw, err
}

// EndSession Terminate authentication session. After this call all requests within session will fail with 403 Forbidden status.
//
// If current session is bount do a gateway connection, such connection will be closed.
//
// Session id to terminate is passed in "X-KSC-Session" header.
func (s *Session) EndSession(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", s.client.Server+"/api/v1.0/Session.EndSession", nil)
	if err != nil {
		return nil, err
	}

	raw, err := s.client.Request(ctx, request, nil)
	return raw, err
}

// StartSession Method to create authenticated session.
// Authentication details should be provided in Authorization HTTP header.
func (s *Session) StartSession(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", s.client.Server+"/api/v1.0/Session.StartSession", nil)

	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Authorization", "KSCBasic user=\""+s.client.UserName+"\", pass=\""+s.client.Password+"\"")
	request.Header.Set("X-KSC-VServer", s.client.VServerName)

	pxgValStr := new(PxgValStr)
	raw, err := s.client.Request(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// CreateBlob Create blob with connection parameters for the klsctunnel utility,
// including target address and one-time authentication token. See Tunnels creation
func (s *Session) CreateBlob(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", s.client.Server+"/api/v1.0/Session.CreateBlob",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := s.client.Request(ctx, request, nil)
	return raw, err
}
