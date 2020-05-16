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

//	NagRdu Class Reference
//
//	Remote diagnostics on host..
//
//	This interface is implemented at Network Agent side,
//	so use gateway connection to connect Network Agent and call interface methods.
//
//	List of all members.
type NagRdu service

//	Acquire current host state
//
//	Returns:
//	- current host state
//
//	Exceptions:
//	- throws	exception in case of error
func (nr *NagRdu) GetCurrentHostState(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.GetCurrentHostState", nil)

	raw, err := nr.client.Do(ctx, request, nil)
	return raw, err
}

//	Get URL-path for later upload file to host
//
//	Returns:
//	- URL-path for uploading file to host and execute it later using ExecuteFileAsync
//
//	Exceptions:
//	- throws	exception in case of error
//
//	See also:
//	Some typical resources path prefixes
func (nr *NagRdu) GetUrlToUploadFileToHost(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.GetUrlToUploadFileToHost", nil)

	raw, err := nr.client.Do(ctx, request, nil)
	return raw, err
}
