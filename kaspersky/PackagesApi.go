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

//	PackagesApi Class Reference
//
//	Operating with packages..
//
//	List of all members.
type PackagesApi service

//Request user agreements related to user packages, registered on current VS.
func (pa *PackagesApi) GetUserAgreements(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetUserAgreements", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := pa.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

type ListOfPackages struct {
	Packages []Package `json:"PxgRetVal"`
}

type Package struct {
	Type  string        `json:"type"`
	Value PackageStruct `json:"value"`
}

type PackageStruct struct {
	KlpkgNpiCreationTime        KlpkgNpiTime `json:"KLPKG_NPI_CREATION_TIME"`
	KlpkgNpiModifTime           KlpkgNpiTime `json:"KLPKG_NPI_MODIF_TIME"`
	KlpkgNpiName                string       `json:"KLPKG_NPI_NAME"`
	KlpkgNpiPackagePath         string       `json:"KLPKG_NPI_PACKAGE_PATH"`
	KlpkgNpiPkgid               int64        `json:"KLPKG_NPI_PKGID"`
	KlpkgNpiProductDisplName    string       `json:"KLPKG_NPI_PRODUCT_DISPL_NAME"`
	KlpkgNpiProductDisplVersion string       `json:"KLPKG_NPI_PRODUCT_DISPL_VERSION"`
	KlpkgNpiProductName         string       `json:"KLPKG_NPI_PRODUCT_NAME"`
	KlpkgNpiProductVersion      string       `json:"KLPKG_NPI_PRODUCT_VERSION"`
	KlpkgNpiSize                KlpkgNpiSize `json:"KLPKG_NPI_SIZE"`
	KlpkgNpiSsDescr             string       `json:"KLPKG_NPI_SS_DESCR"`
}

type KlpkgNpiTime struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type KlpkgNpiSize struct {
	Type  string `json:"type"`
	Value int64  `json:"value"`
}

func (pa *PackagesApi) GetPackages(ctx context.Context) (*ListOfPackages, []byte, error) {
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetPackages", nil)
	if err != nil {
		return nil, nil, err
	}

	listOfPackages := new(ListOfPackages)
	raw, err := pa.client.Do(ctx, request, listOfPackages)
	return listOfPackages, raw, err
}

func (pa *PackagesApi) GetPackages2(ctx context.Context) (*ListOfPackages, []byte, error) {
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetPackages2", nil)
	if err != nil {
		return nil, nil, err
	}

	listOfPackages := new(ListOfPackages)
	raw, err := pa.client.Do(ctx, request, listOfPackages)
	return listOfPackages, raw, err
}

/*
func (pa *PackagesApi) GetPackageInfo2(ctx context.Context, nPackageId int) (SKlpkg, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nPackageId": %d
	}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server + "/api/v1.0/PackagesApi.GetPackageInfo2", bytes.NewBuffer(postData))
    if err != nil {
    return nil, err
  }

	raw, err := pa.client.Do(ctx, request)
	pxgRetVal, err := UnmarshalKlpkg(raw)
	return pxgRetVal, err
}

func (pa *PackagesApi) DeleteExecutablePkg(ctx context.Context, nPackageId int) (SKlpkg, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nPackageId": %d
	}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server + "/api/v1.0/PackagesApi.DeleteExecutablePkg", bytes.NewBuffer(postData))
    if err != nil {
    return nil, err
  }

	raw, err := pa.client.Do(ctx, request)
	pxgRetVal, err := UnmarshalKlpkg(raw)
	return pxgRetVal, err
}

//message":"Package #1 was not found.
func (pa *PackagesApi) RemovePackage2(ctx context.Context, nPackageId int) (PDResult, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nPackageId": %d
	}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server + "/api/v1.0/PackagesApi.RemovePackage2", bytes.NewBuffer(postData))
    if err != nil {
    return nil, err
  }

	raw, err := pa.client.Do(ctx, request)
	pxgRetVal, err := UnmarshalPDResult(raw)
	return pxgRetVal, err
}

func (pa *PackagesApi) RemovePackage(ctx context.Context, nPackageId int) (SKlpkg, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nPackageId": %d
	}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server + "/api/v1.0/PackagesApi.RemovePackage", bytes.NewBuffer(postData))
    if err != nil {
    return nil, err
  }

	raw, err := pa.client.Do(ctx, request)
	pxgRetVal, err := UnmarshalKlpkg(raw)
	return pxgRetVal, err
}

*/
