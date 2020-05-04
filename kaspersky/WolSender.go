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
	"encoding/json"
	"log"
	"net/http"
)

//WolSender

type WolSender struct {
	client *Client
}

type WOLSignalParams struct {
	SzwHostID string `json:"szwHostId,omitempty"`
}

//Sends Wake-On-LAN signal to host.
//
//The goal of this call is to wake target host up.
//This is done by sending WOL packets by server, some Update Agent Versions assigned to target host and some nAgent Versions which are likely to be located in same subnet where target host is located.
//Besides server will wake up Connection gateway assigned to host as well.
//Target WOL packets are sent to broadcast address 255.255.255.255, direct host IP and subnet-directed broadcast (like '10.11.12.255'). WOL packets sent to ports 7 and 9.
func (ah *WolSender) SendWolSignal(ctx context.Context, wolsp WOLSignalParams) {

	postData, _ := json.Marshal(wolsp)

	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/WolSender.SendWolSignal", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = ah.client.Do(ctx, request, nil)
}
