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

// SrvCloudStat Service for working with statistics of using clouds.
type SrvCloudStat service

// CloudWizardStarted Notify that the Cloud Wizard has been started.
func (scs *SrvCloudStat) CloudWizardStarted(ctx context.Context) error {
	request, err := http.NewRequest("POST", scs.client.Server+"/api/v1.0/SrvCloudStat.CloudWizardStarted", nil)
	if err != nil {
		return err
	}

	raw, err := scs.client.Do(ctx, request, nil)

	if scs.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

// CloudWizardCompleted Notify that the Cloud Wizard has been completed.
// Method shouldn't be called if user cancelled the wizard.
func (scs *SrvCloudStat) CloudWizardCompleted(ctx context.Context, bErrorHappen bool) error {
	postData := []byte(fmt.Sprintf(`{"bErrorHappen": %v}`, bErrorHappen))
	request, err := http.NewRequest("POST", scs.client.Server+"/api/v1.0/SrvCloudStat.CloudWizardCompleted",
		bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	raw, err := scs.client.Do(ctx, request, nil)

	if scs.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}
