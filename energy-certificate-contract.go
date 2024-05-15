/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// EnergyCertificateContract contract for managing CRUD for EnergyCertificate
type EnergyCertificateContract struct {
	contractapi.Contract
}

// EnergyCertificateExists returns true when asset with given ID exists in world state
func (c *EnergyCertificateContract) EnergyCertificateExists(ctx contractapi.TransactionContextInterface, tokenRef string) (bool, error) {
	data, err := ctx.GetStub().GetState(tokenRef)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateEnergyCertificate creates a new instance of EnergyCertificate
func (c *EnergyCertificateContract) CreateEnergyCertificate(ctx contractapi.TransactionContextInterface, ownerId string, producerId string, emissionDate string, usableMonth int, usableYear int, regulatoryAuthorityID string) (string, error) {
	energyCertificate := EnergyCertificate{
		EnergyCertificateID:   ctx.GetStub().GetTxID(),
		OwnerID:               ownerId,
		ProducerID:            producerId,
		EmissionDate:          emissionDate,
		UsableMonth:           usableMonth,
		UsableYear:            usableYear,
		RegulatoryAuthorityID: regulatoryAuthorityID,
	}

	bytes, err := json.Marshal(energyCertificate)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(energyCertificate.EnergyCertificateID, bytes)
	if err != nil {
		return "", fmt.Errorf("failed to create energy certificate: %v", err)
	}

	return fmt.Sprintf("%s created successfully", energyCertificate.EnergyCertificateID), nil
}

// ReadEnergyCertificate retrieves an instance of EnergyCertificate from the world state
func (c *EnergyCertificateContract) ReadEnergyCertificate(ctx contractapi.TransactionContextInterface, energyCertificateID string) (*EnergyCertificate, error) {
	exists, err := c.EnergyCertificateExists(ctx, energyCertificateID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", energyCertificateID)
	}

	bytes, _ := ctx.GetStub().GetState(energyCertificateID)

	energyCertificate := new(EnergyCertificate)

	err = json.Unmarshal(bytes, energyCertificate)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type EnergyCertificate")
	}

	return energyCertificate, nil
}

// UpdateEnergyCertificate retrieves an instance of EnergyCertificate from the world state and updates its value
func (c *EnergyCertificateContract) UpdateEnergyCertificate(ctx contractapi.TransactionContextInterface, energyCertificateID string, newUsableMonth int, newUsableYear int) error {
	certificateJSON, err := ctx.GetStub().GetState(energyCertificateID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if certificateJSON == nil {
		return fmt.Errorf("the asset %s does not exist", energyCertificateID)
	}

	var energyCertificate EnergyCertificate
	err = json.Unmarshal(certificateJSON, &energyCertificate)
	if err != nil {
		return err
	}

	updatedCertificate, err := json.Marshal(energyCertificate)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(energyCertificateID, updatedCertificate)
}

// DeleteEnergyCertificate deletes an instance of EnergyCertificate from the world state
func (c *EnergyCertificateContract) DeleteEnergyCertificate(ctx contractapi.TransactionContextInterface, energyCertificateID string) error {
	exists, err := c.EnergyCertificateExists(ctx, energyCertificateID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", energyCertificateID)
	}

	return ctx.GetStub().DelState(energyCertificateID)
}

func (c *EnergyCertificateContract) TransferEnergyCertificate(ctx contractapi.TransactionContextInterface, energyCertificateID string, newOwnerId string, price float64) (string, error) {
	certificateJSON, err := ctx.GetStub().GetState(energyCertificateID)
	if err != nil {
		return "", fmt.Errorf("failed to read certificate: %v", err)
	}
	if certificateJSON == nil {
		return "", fmt.Errorf("certificate does not exist: %s", energyCertificateID)
	}

	var energyCertificate EnergyCertificate
	err = json.Unmarshal(certificateJSON, &energyCertificate)
	if err != nil {
		return "", err
	}

	oldOwnerId := energyCertificate.OwnerID
	energyCertificate.OwnerID = newOwnerId

	updatedCertificateJSON, err := json.Marshal(energyCertificate)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(energyCertificateID, updatedCertificateJSON)
	if err != nil {
		return "", err
	}

	transaction := Transaction{
		TransactionID:       ctx.GetStub().GetTxID(),
		CertificateTokenRef: energyCertificate.EnergyCertificateID,
		FromUserID:          oldOwnerId,
		ToUserID:            newOwnerId,
		TransactionDate:     time.Now().Format(time.RFC3339),
		Price:               price,
	}

	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return "", err
	}

	transactionKey := fmt.Sprintf("TRANSACTION_%s", transaction.TransactionID)

	err = ctx.GetStub().PutState(transactionKey, transactionJSON)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s transaction created successfully", transactionKey), nil
}

func (c *EnergyCertificateContract) QueryEnergyCertificate(ctx contractapi.TransactionContextInterface, energyCertificateID string) (*EnergyCertificate, error) {
	certificateJSON, err := ctx.GetStub().GetState(energyCertificateID)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate: %v", err)
	}
	if certificateJSON == nil {
		return nil, fmt.Errorf("certificate does not exist: %s", energyCertificateID)
	}

	var certificate EnergyCertificate
	err = json.Unmarshal(certificateJSON, &certificate)
	if err != nil {
		return nil, err
	}

	return &certificate, nil
}

func (c *EnergyCertificateContract) QueryTransaction(ctx contractapi.TransactionContextInterface, transactionKey string) (*Transaction, error) {
	transactionJSON, err := ctx.GetStub().GetState(transactionKey)
	if err != nil {
		return nil, fmt.Errorf("failed to read transaction: %v", err)
	}
	if transactionJSON == nil {
		return nil, fmt.Errorf("transaction does not exist: %s", transactionKey)
	}

	var transaction Transaction
	err = json.Unmarshal(transactionJSON, &transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (c *EnergyCertificateContract) GetCertificatesByOwnerID(ctx contractapi.TransactionContextInterface, ownerId string) ([]EnergyCertificate, error) {
	queryString := fmt.Sprintf(`{"selector":{"ownerId":"%s"}}`, ownerId)
	return c.getCertificatesByQueryString(ctx, queryString)
}

func (c *EnergyCertificateContract) GetCertificatesByProducerID(ctx contractapi.TransactionContextInterface, producerId string) ([]EnergyCertificate, error) {
	queryString := fmt.Sprintf(`{"selector":{"producerId":"%s"}}`, producerId)
	return c.getCertificatesByQueryString(ctx, queryString)
}

func (c *EnergyCertificateContract) GetTransactionsByFromUserID(ctx contractapi.TransactionContextInterface, fromUserId string) ([]Transaction, error) {
	queryString := fmt.Sprintf(`{"selector":{"fromUserId":"%s"}}`, fromUserId)
	return c.getTransactionsByQueryString(ctx, queryString)
}

func (c *EnergyCertificateContract) GetTransactionsByToUserID(ctx contractapi.TransactionContextInterface, toUserId string) ([]Transaction, error) {
	queryString := fmt.Sprintf(`{"selector":{"toUserId":"%s"}}`, toUserId)
	return c.getTransactionsByQueryString(ctx, queryString)
}

func (c *EnergyCertificateContract) GetTransactionsByCertificateID(ctx contractapi.TransactionContextInterface, tokenRef string) ([]Transaction, error) {
	queryString := fmt.Sprintf(`{"selector":{"tokenRef":"%s"}}`, tokenRef)
	return c.getTransactionsByQueryString(ctx, queryString)
}

func (c *EnergyCertificateContract) getCertificatesByQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]EnergyCertificate, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var certificates []EnergyCertificate
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var certificate EnergyCertificate
		err = json.Unmarshal(queryResponse.Value, &certificate)
		if err != nil {
			return nil, err
		}
		certificates = append(certificates, certificate)
	}

	return certificates, nil
}

func (c *EnergyCertificateContract) getTransactionsByQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]Transaction, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var transactions []Transaction
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var transaction Transaction
		err = json.Unmarshal(queryResponse.Value, &transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
