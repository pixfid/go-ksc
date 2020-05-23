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

	"net/http"
)

//	WolSender Class Reference
//
//	Wake-On-LAN signal sender.
//
//	List of all members.
type WolSender service

//	Sends Wake-On-LAN signal to host.
//
//	The goal of this call is to wake target host up.
//	This is done by sending WOL packets by server,
//	some Update Agent Versions assigned to target host and some nAgent Versions which are likely
//	to be located in same subnet where target host is located.
//	Besides server will wake up Connection gateway assigned to host as well.
//	Target WOL packets are sent to broadcast address 255.255.255.255,
//	direct host IP and subnet-directed broadcast (like '10.11.12.255').
//	WOL packets sent to ports 7 and 9.
func (ws *WolSender) SendWolSignal(ctx context.Context, szwHostId string) error {
	postData := []byte(fmt.Sprintf(`{"szwHostId":"%s"}`, szwHostId))
	request, err := http.NewRequest("POST", ws.client.Server+"/api/v1.0/WolSender.SendWolSignal",
		bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = ws.client.Do(ctx, request, nil)
	return err
}
