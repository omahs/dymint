// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	da "github.com/dymensionxyz/dymint/da"
	mock "github.com/stretchr/testify/mock"

	settlement "github.com/dymensionxyz/dymint/settlement"

	types "github.com/dymensionxyz/dymint/types"
)

// HubClient is an autogenerated mock type for the HubClient type
type HubClient struct {
	mock.Mock
}

// GetBatchAtIndex provides a mock function with given fields: rollappID, index
func (_m *HubClient) GetBatchAtIndex(rollappID string, index uint64) (*settlement.ResultRetrieveBatch, error) {
	ret := _m.Called(rollappID, index)

	var r0 *settlement.ResultRetrieveBatch
	if rf, ok := ret.Get(0).(func(string, uint64) *settlement.ResultRetrieveBatch); ok {
		r0 = rf(rollappID, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settlement.ResultRetrieveBatch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint64) error); ok {
		r1 = rf(rollappID, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBatch provides a mock function with given fields: rollappID
func (_m *HubClient) GetLatestBatch(rollappID string) (*settlement.ResultRetrieveBatch, error) {
	ret := _m.Called(rollappID)

	var r0 *settlement.ResultRetrieveBatch
	if rf, ok := ret.Get(0).(func(string) *settlement.ResultRetrieveBatch); ok {
		r0 = rf(rollappID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settlement.ResultRetrieveBatch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(rollappID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSequencers provides a mock function with given fields: rollappID
func (_m *HubClient) GetSequencers(rollappID string) ([]*types.Sequencer, error) {
	ret := _m.Called(rollappID)

	var r0 []*types.Sequencer
	if rf, ok := ret.Get(0).(func(string) []*types.Sequencer); ok {
		r0 = rf(rollappID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Sequencer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(rollappID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostBatch provides a mock function with given fields: batch, daClient, daResult
func (_m *HubClient) PostBatch(batch *types.Batch, daClient da.Client, daResult *da.ResultSubmitBatch) (settlement.PostBatchResp, error) {
	ret := _m.Called(batch, daClient, daResult)

	var r0 settlement.PostBatchResp
	if rf, ok := ret.Get(0).(func(*types.Batch, da.Client, *da.ResultSubmitBatch) settlement.PostBatchResp); ok {
		r0 = rf(batch, daClient, daResult)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settlement.PostBatchResp)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Batch, da.Client, *da.ResultSubmitBatch) error); ok {
		r1 = rf(batch, daClient, daResult)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *HubClient) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *HubClient) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewHubClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewHubClient creates a new instance of HubClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHubClient(t mockConstructorTestingTNewHubClient) *HubClient {
	mock := &HubClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
