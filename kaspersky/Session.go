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
	"context"
	"net/http"
)

//	Session Class Reference
//
//	Session management interface..
//
//	It allows to create session token for current security context.
//
//	List of all members.
type Session service

//Creates session token.
//
//Creates session token for current security context.
//Those token can be used for logon purposes to Administaration Server for a short time (3 minutes by default).
//
//	Returns:
//	- Session token. (data.PxgValStr)
func (s *Session) CreateToken(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", s.client.Server+"/api/v1.0/Session.CreateToken", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := s.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//	Does nothing. May be used to effectively verify session validity.
//	Session id to verify is passed in "X-KSC-Session" header.
//
//	See also:
//	Authenticated session
func (s *Session) Ping(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", s.client.Server+"/api/v1.0/Session.Ping", nil)
	if err != nil {
		return nil, err
	}

	raw, err := s.client.Do(ctx, request, nil)
	return raw, err
}

//	Terminate authentication session.
//	After this call all requests within session will fail with 403 Forbidden status.
//	If current session is bount do a gateway connection, such connection will be closed.
//
//	Session id to terminate is passed in "X-KSC-Session" header.
//
//	Returns:
//	- Session token.
//
//	See also:
//
//	Authenticated session
func (s *Session) EndSession(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", s.client.Server+"/api/v1.0/Session.EndSession", nil)
	if err != nil {
		return nil, err
	}

	raw, err := s.client.Do(ctx, request, nil)
	return raw, err
}

//	Method to create authenticated session.
//	Authentication details should be provided in Authorization HTTP header.
//
//	Returns:
//	- Session token.
//
//	See also:
//
//	Authenticated session
func (s *Session) StartSession(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", s.client.Server+"/api/v1.0/Session.StartSession", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := s.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//TODO func (s *Session) CreateBlob(ctx context.Context, params interface{}) ([]byte, error) {return nil, nil}
