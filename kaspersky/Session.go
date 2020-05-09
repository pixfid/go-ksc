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
//	Session management interface. More...
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

	pxgValStr := new(PxgValStr)

	raw, err := s.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}
