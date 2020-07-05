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

// FilesAcceptor service to upload files to server.
type FilesAcceptor service

// CancelFileUpload Cancel file upload operation.
//
// After cancellation provided URL will not be valid anymore and any uploaded by the moment file chunks will be dropped.
// You should not call this method unless you want to break upload operation.
func (di *FilesAcceptor) CancelFileUpload(ctx context.Context, wstrFileId string) error {
	postData := []byte(fmt.Sprintf(`{"wstrFileId": "%s"}`, wstrFileId))
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/FilesAcceptor.CancelFileUpload", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = di.client.Do(ctx, request, nil)
	return err
}

// FileUploadData struct
type FileUploadData struct {
	// WstrFileID File identifier to be used in KSC API after file upload
	WstrFileID string `json:"wstrFileId,omitempty"`

	// WstrUploadURL relative URL on KSC server to upload file to.
	// URL becomes invalid when file is uploaded, or after server is restarted, or CancelFileUpload is called.
	WstrUploadURL string `json:"wstrUploadURL,omitempty"`
}

// InitiateFileUpload Prepare server to accept file.
//
// In order to upload file(s) to KSC server follow these steps:
//
// Call this function to get upload URL and file identifier which can be used in KSC API (e.g. uploading installation package)
//
// Upload file using provided URL via HTTP PUT calls
//
// Provide file identifier to KSC API to use file on server side
//
// Uploaded file should be used immediately after upload. Otherwise it may be dropped by timeout or during KSC server restart.
//
// Original file name is ignored, only returned file identifier is used to identify uploaded file.
//
// If you need to transfer to server a directory or several files put them into zip or tar.gz archive and use 'bIsArchive' flag.
//
// All path names inside archive must be in UTF-8 encoding.
func (di *FilesAcceptor) InitiateFileUpload(ctx context.Context, bIsArchive bool, qwFileSize int64) (*FileUploadData, error) {
	postData := []byte(fmt.Sprintf(`{"bIsArchive": %v, "qwFileSize": %d}`, bIsArchive, qwFileSize))
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/FilesAcceptor.InitiateFileUpload", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	fileUploadData := new(FileUploadData)
	_, err = di.client.Do(ctx, request, &fileUploadData)
	return fileUploadData, err
}
