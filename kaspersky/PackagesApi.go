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
	"net/http"
)

//	PackagesApi Class Reference
//
//	Operating with packages. More...
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
