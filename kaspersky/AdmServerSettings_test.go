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
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAdmServerSettings_GetSharedFolder(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/AdmServerSettings.GetSharedFolder", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w,
			`{"PxgRetVal":"\\\\T278-vnoa-ksc\\klshare\\"}`,
		)
	})

	artifacts, raw, err := client.AdmServerSettings.GetSharedFolder(context.Background())
	if err != nil {
		t.Errorf("AdmServerSettings.GetSharedFolder returned error: %v", err)
	}

	println(string(raw))

	want := &PxgValStr{Str: "\\\\T278-vnoa-ksc\\klshare\\"}
	if !reflect.DeepEqual(artifacts, want) {
		t.Errorf("AdmServerSettings.GetSharedFolder returned %+v, want %+v", artifacts, want)
	}
}
