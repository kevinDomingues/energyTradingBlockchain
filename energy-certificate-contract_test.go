/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	// "encoding/json"
	// "errors"
	// "fmt"
	// "testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const getStateError = "world state get error"

type MockStub struct {
	shim.ChaincodeStubInterface
	mock.Mock
}

func (ms *MockStub) GetState(key string) ([]byte, error) {
	args := ms.Called(key)

	return args.Get(0).([]byte), args.Error(1)
}

func (ms *MockStub) PutState(key string, value []byte) error {
	args := ms.Called(key, value)

	return args.Error(0)
}

func (ms *MockStub) DelState(key string) error {
	args := ms.Called(key)

	return args.Error(0)
}

type MockContext struct {
	contractapi.TransactionContextInterface
	mock.Mock
}

func (mc *MockContext) GetStub() shim.ChaincodeStubInterface {
	args := mc.Called()

	return args.Get(0).(*MockStub)
}

// func configureStub() (*MockContext, *MockStub) {
// 	var nilBytes []byte

// 	testEnergyCertificate := new(EnergyCertificate)
// 	testEnergyCertificate.Value = "set value"
// 	energyCertificateBytes, _ := json.Marshal(testEnergyCertificate)

// 	ms := new(MockStub)
// 	ms.On("GetState", "statebad").Return(nilBytes, errors.New(getStateError))
// 	ms.On("GetState", "missingkey").Return(nilBytes, nil)
// 	ms.On("GetState", "existingkey").Return([]byte("some value"), nil)
// 	ms.On("GetState", "energyCertificatekey").Return(energyCertificateBytes, nil)
// 	ms.On("PutState", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)
// 	ms.On("DelState", mock.AnythingOfType("string")).Return(nil)

// 	mc := new(MockContext)
// 	mc.On("GetStub").Return(ms)

// 	return mc, ms
// }

// func TestEnergyCertificateExists(t *testing.T) {
// 	var exists bool
// 	var err error

// 	ctx, _ := configureStub()
// 	c := new(EnergyCertificateContract)

// 	exists, err = c.EnergyCertificateExists(ctx, "statebad")
// 	assert.EqualError(t, err, getStateError)
// 	assert.False(t, exists, "should return false on error")

// 	exists, err = c.EnergyCertificateExists(ctx, "missingkey")
// 	assert.Nil(t, err, "should not return error when can read from world state but no value for key")
// 	assert.False(t, exists, "should return false when no value for key in world state")

// 	exists, err = c.EnergyCertificateExists(ctx, "existingkey")
// 	assert.Nil(t, err, "should not return error when can read from world state and value exists for key")
// 	assert.True(t, exists, "should return true when value for key in world state")
// }

// func TestCreateEnergyCertificate(t *testing.T) {
// 	var err error

// 	ctx, stub := configureStub()
// 	c := new(EnergyCertificateContract)

// 	err = c.CreateEnergyCertificate(ctx, "statebad", "some value")
// 	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

// 	err = c.CreateEnergyCertificate(ctx, "existingkey", "some value")
// 	assert.EqualError(t, err, "The asset existingkey already exists", "should error when exists returns true")

// 	err = c.CreateEnergyCertificate(ctx, "missingkey", "some value")
// 	stub.AssertCalled(t, "PutState", "missingkey", []byte("{\"value\":\"some value\"}"))
// }

// func TestReadEnergyCertificate(t *testing.T) {
// 	var energyCertificate *EnergyCertificate
// 	var err error

// 	ctx, _ := configureStub()
// 	c := new(EnergyCertificateContract)

// 	energyCertificate, err = c.ReadEnergyCertificate(ctx, "statebad")
// 	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when reading")
// 	assert.Nil(t, energyCertificate, "should not return EnergyCertificate when exists errors when reading")

// 	energyCertificate, err = c.ReadEnergyCertificate(ctx, "missingkey")
// 	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when reading")
// 	assert.Nil(t, energyCertificate, "should not return EnergyCertificate when key does not exist in world state when reading")

// 	energyCertificate, err = c.ReadEnergyCertificate(ctx, "existingkey")
// 	assert.EqualError(t, err, "Could not unmarshal world state data to type EnergyCertificate", "should error when data in key is not EnergyCertificate")
// 	assert.Nil(t, energyCertificate, "should not return EnergyCertificate when data in key is not of type EnergyCertificate")

// 	energyCertificate, err = c.ReadEnergyCertificate(ctx, "energyCertificatekey")
// 	expectedEnergyCertificate := new(EnergyCertificate)
// 	expectedEnergyCertificate.Value = "set value"
// 	assert.Nil(t, err, "should not return error when EnergyCertificate exists in world state when reading")
// 	assert.Equal(t, expectedEnergyCertificate, energyCertificate, "should return deserialized EnergyCertificate from world state")
// }

// func TestUpdateEnergyCertificate(t *testing.T) {
// 	var err error

// 	ctx, stub := configureStub()
// 	c := new(EnergyCertificateContract)

// 	err = c.UpdateEnergyCertificate(ctx, "statebad", "new value")
// 	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

// 	err = c.UpdateEnergyCertificate(ctx, "missingkey", "new value")
// 	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

// 	err = c.UpdateEnergyCertificate(ctx, "energyCertificatekey", "new value")
// 	expectedEnergyCertificate := new(EnergyCertificate)
// 	expectedEnergyCertificate.Value = "new value"
// 	expectedEnergyCertificateBytes, _ := json.Marshal(expectedEnergyCertificate)
// 	assert.Nil(t, err, "should not return error when EnergyCertificate exists in world state when updating")
// 	stub.AssertCalled(t, "PutState", "energyCertificatekey", expectedEnergyCertificateBytes)
// }

// func TestDeleteEnergyCertificate(t *testing.T) {
// 	var err error

// 	ctx, stub := configureStub()
// 	c := new(EnergyCertificateContract)

// 	err = c.DeleteEnergyCertificate(ctx, "statebad")
// 	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

// 	err = c.DeleteEnergyCertificate(ctx, "missingkey")
// 	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when deleting")

// 	err = c.DeleteEnergyCertificate(ctx, "energyCertificatekey")
// 	assert.Nil(t, err, "should not return error when EnergyCertificate exists in world state when deleting")
// 	stub.AssertCalled(t, "DelState", "energyCertificatekey")
// }
