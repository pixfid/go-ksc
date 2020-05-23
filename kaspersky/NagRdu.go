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
