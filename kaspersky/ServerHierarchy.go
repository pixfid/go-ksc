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
	"log"
	"net/http"
)

type ServerHierarchy struct {
	client *Client
}

func (sh *ServerHierarchy) GetServerInfo(ctx context.Context, lServer int) (PxgValStr, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"lServer": %d,
	"pFields": 
			[
				"KLSRVH_SRV_ID",
				"KLSRVH_SRV_INST_ID",
				"KLSRVH_SRV_ADDR"
			]
	}`, lServer))

	request, err := http.NewRequest("POST", sh.client.Server+"/api/v1.0/ServerHierarchy.GetServerInfo", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = sh.client.Do(ctx, request, nil, false)
	return PxgValStr{}, nil
}

//------
func (sh *ServerHierarchy) GetChildServers(ctx context.Context, nGroupId int64) (PxgRetError, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nGroupId": %d
	}`, nGroupId))

	request, err := http.NewRequest("POST", sh.client.Server+"/api/v1.0/ServerHierarchy.GetChildServers", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	_, _ = sh.client.Do(ctx, request, nil, false)
	//TODO FIX THIS SHIT!!!
	return PxgRetError{}, err
}
