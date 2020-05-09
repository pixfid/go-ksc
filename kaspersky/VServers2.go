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
	"bytes"
	"context"
	"fmt"
	"net/http"
)

//	VServers2 Class Reference
//	Virtual servers processing. More...
//
//	List of all members:
type VServers2 service

//Acquire info on virtual server.
//
//Returns info about the specified virtual server
//
//Parameters:
//	- lVsId	(int64) virtual server id
//Returns:
//	- (params) a container, see Virtual server statistic.
func (vs *VServers2) GetVServerStatistic(ctx context.Context, lVsId int) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"lVsId": %d
	}`, lVsId))

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers2.GetVServerStatistic", bytes.NewBuffer(postData))

	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}
