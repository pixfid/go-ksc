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

//	SmsSenders Class Reference
//
//	Configure mobile devices as SMS senders..
//
//	List of all members.
type SmsSenders service

//	checks if there is a device allowed to send SMS
//
//	Returns:
//	- (bool) true if server has devices allowed to send SMS, false otherwise
func (ss *SmsSenders) HasAllowedSenders(ctx context.Context) (*PxgValBool, []byte, error) {
	request, err := http.NewRequest("POST", ss.client.Server+"/api/v1.0/SmsSenders.Clear", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := ss.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//TODO func (ss *SmsSenders) AllowSenders(ctx context.Context, params interface{}) ([]byte, error) {
