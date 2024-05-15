/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	energyCertificateContract := new(EnergyCertificateContract)
	energyCertificateContract.Info.Version = "0.0.1"
	energyCertificateContract.Info.Description = "My Smart Contract"
	energyCertificateContract.Info.License = new(metadata.LicenseMetadata)
	energyCertificateContract.Info.License.Name = "Apache-2.0"
	energyCertificateContract.Info.Contact = new(metadata.ContactMetadata)
	energyCertificateContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(energyCertificateContract)
	chaincode.Info.Title = "energyTradingBlockchain chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from EnergyCertificateContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
