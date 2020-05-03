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

func TestSession_CreateToken(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/Session.CreateToken", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w,
			`{"PxgRetVal":"5525d99d-2195-42d2-a1c4-ee5e227c16fd"}`,
		)
	})

	artifacts, raw, err := client.Session.CreateToken(context.Background())
	if err != nil {
		t.Errorf("Session.CreateToken returned error: %v", err)
	}

	println(string(raw))

	want := &PxgValStr{Str: "5525d99d-2195-42d2-a1c4-ee5e227c16fd"}
	if !reflect.DeepEqual(artifacts, want) {
		t.Errorf("Session.CreateToken returned %+v, want %+v", artifacts, want)
	}
}