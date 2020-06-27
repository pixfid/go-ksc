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

// IWebUsersSrv service to operating with emails
type IWebUsersSrv service

// SendEmail Send an email to multiple recipients.
//
// How to embeed QR-code pictures into html body ("QrCodes" array):
//
// Set "EmbedQrCodes" value to true if you want to embed QR-code pictures into html rather then sending them as attachments.
//
// Pictures will be embedded into html using pattern:
//
// '<div id="qrcode<arrayIndex>"><img src="cid:<"QrCodePicture">"></div>' which will pe appended to the end of "MailBody"
// (in respect of the order they appear in "QrCodes" array)
//
// For example, embedding of the first qr-code picture from "QrCodes" array would look like:
// <div id="qrcode0"><img src="cid:example.png"></div>
//
// You can also css-style that div using its id inside of your html message body.
func (iwus *IWebUsersSrv) SendEmail(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", iwus.client.Server+"/api/v1.0/IWebUsersSrv.SendEmail",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := iwus.client.Do(ctx, request, nil)
	return raw, err
}
