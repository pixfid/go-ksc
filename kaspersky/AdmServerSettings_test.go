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

