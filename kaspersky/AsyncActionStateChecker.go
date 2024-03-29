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

// AsyncActionStateChecker service to monitor state of async action
//
// It is supposed that client of the AsyncActionStateChecker has added async action and this action has identifier wstrActionGuid.
//
// After that client should wait while CheckActionState will return bFinalized for this action or should cancel this action.
//
// If the count of finalized actions is reach some limit then client connection will be closed automatically.
type AsyncActionStateChecker service

type ActionStateResult struct {
	BFinalized         bool        `json:"bFinalized"`
	BSuccededFinalized bool        `json:"bSuccededFinalized"`
	LStateCode         int64       `json:"lStateCode"`
	PStateData         *PStateData `json:"pStateData,omitempty"`
	LNextCheckDelay    int64       `json:"lNextCheckDelay"`
}

type PStateData struct {
	KlblagErrorCode    int64  `json:"KLBLAG_ERROR_CODE,omitempty"`
	KlblagErrorFname   string `json:"KLBLAG_ERROR_FNAME,omitempty"`
	KlblagErrorLnumber int64  `json:"KLBLAG_ERROR_LNUMBER,omitempty"`
	KlblagErrorModule  string `json:"KLBLAG_ERROR_MODULE,omitempty"`
	KlblagErrorMsg     string `json:"KLBLAG_ERROR_MSG,omitempty"`
	KlblagErrorSubcode int64  `json:"KLBLAG_ERROR_SUBCODE,omitempty"`
	//
	KlrptOutputChart  string             `json:"KLRPT_OUTPUT_CHART,omitempty"`
	KlrptOutputFile   string             `json:"KLRPT_OUTPUT_FILE,omitempty"`
	KlrptOutputFormat *KlrptOutputFormat `json:"KLRPT_OUTPUT_FORMAT,omitempty"`
	KlrptOutputLogo   string             `json:"KLRPT_OUTPUT_LOGO,omitempty"`
	//
	LastActionResult string `json:"LastActionResult,omitempty"`
	//
	EventLogs        []string          `json:"EventLogs"`
	HostDN           string            `json:"HostDN,omitempty"`
	InstallationLogs *InstallationLogs `json:"InstallationLogs,omitempty"`
	Products         *Products         `json:"Products,omitempty"`
	WuaLogs          []string          `json:"WuaLogs"`
	WuaLogsWin10     []string          `json:"WuaLogs_Win10"`
}

// CheckActionState Check status of the async action.
//
// Precondition:
// Action with identifier wstrActionGuid has been added on the same connection where this method is call
// if there was previous call of this method and it return lNextCheckDelay
// then there have passed not less than lNextCheckDelay milliseconds from this previous call
//
// Postcondition:
//
// if returns bFinalized==true then this action has been removed, and wstrActionGuid is not valid any more.
// Otherwise in lNextCheckDelay it should be returned delay in msec to Request next call of the CheckActionState
func (ac *AsyncActionStateChecker) CheckActionState(ctx context.Context, wstrActionGuid string) (*ActionStateResult, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrActionGuid": "%s"}`, wstrActionGuid))
	request, err := http.NewRequest("POST", ac.client.Server+"/api/v1.0/AsyncActionStateChecker.CheckActionState", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	aSResult := new(ActionStateResult)
	raw, err := ac.client.Request(ctx, request, &aSResult)
	return aSResult, raw, err
}
