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
	"log"
	"net/http"
)

// ServiceNwcCommandProvider Service console command provider.
type ServiceNwcCommandProvider service

// GetCommandPayload Get command payload.
func (srn *ServiceNwcCommandProvider) GetCommandPayload(ctx context.Context, wstrCommandId string) error {
	postData := []byte(fmt.Sprintf(`{"wstrCommandId": "%s"}`, wstrCommandId))
	request, err := http.NewRequest("POST", srn.client.Server+"/api/v1.0/ServiceNwcCommandProvider.GetCommandPayload",
		bytes.NewBuffer(postData))

	if err != nil {
		return err
	}

	//TODO check response
	raw, err := srn.client.Request(ctx, request, nil)

	if srn.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

//TODO SetCommandError
//TODO SetCommandResult
