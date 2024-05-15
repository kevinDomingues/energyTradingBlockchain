/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// EnergyCertificate stores a value
type EnergyCertificate struct {
	EnergyCertificateID   string `json:"energyCertificateId"`
	OwnerID               string `json:"ownerId"`
	ProducerID            string `json:"producerId"`
	EmissionDate          string `json:"emissionDate"`
	UsableMonth           int    `json:"usableMonth"`
	UsableYear            int    `json:"usableYear"`
	RegulatoryAuthorityID string `json:"regulatoryAuthorityID"`
}

type Transaction struct {
	TransactionID       string  `json:"transactionId"`
	CertificateTokenRef string  `json:"tokenRef"`
	FromUserID          string  `json:"fromUserId"`
	ToUserID            string  `json:"toUserId"`
	TransactionDate     string  `json:"transactionDate"`
	Price               float64 `json:"price"`
}
