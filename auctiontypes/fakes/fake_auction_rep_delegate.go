// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/auction/auctiontypes"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
)

type FakeAuctionRepDelegate struct {
	RemainingResourcesStub        func() (auctiontypes.Resources, error)
	remainingResourcesMutex       sync.RWMutex
	remainingResourcesArgsForCall []struct{}
	remainingResourcesReturns     struct {
		result1 auctiontypes.Resources
		result2 error
	}
	TotalResourcesStub        func() (auctiontypes.Resources, error)
	totalResourcesMutex       sync.RWMutex
	totalResourcesArgsForCall []struct{}
	totalResourcesReturns     struct {
		result1 auctiontypes.Resources
		result2 error
	}
	NumInstancesForProcessGuidStub        func(processGuid string) (int, error)
	numInstancesForProcessGuidMutex       sync.RWMutex
	numInstancesForProcessGuidArgsForCall []struct {
		processGuid string
	}
	numInstancesForProcessGuidReturns struct {
		result1 int
		result2 error
	}
	InstanceGuidsForProcessGuidAndIndexStub        func(processGuid string, index int) ([]string, error)
	instanceGuidsForProcessGuidAndIndexMutex       sync.RWMutex
	instanceGuidsForProcessGuidAndIndexArgsForCall []struct {
		processGuid string
		index       int
	}
	instanceGuidsForProcessGuidAndIndexReturns struct {
		result1 []string
		result2 error
	}
	ReserveStub        func(startAuctionInfo auctiontypes.StartAuctionInfo) error
	reserveMutex       sync.RWMutex
	reserveArgsForCall []struct {
		startAuctionInfo auctiontypes.StartAuctionInfo
	}
	reserveReturns struct {
		result1 error
	}
	ReleaseReservationStub        func(startAuctionInfo auctiontypes.StartAuctionInfo) error
	releaseReservationMutex       sync.RWMutex
	releaseReservationArgsForCall []struct {
		startAuctionInfo auctiontypes.StartAuctionInfo
	}
	releaseReservationReturns struct {
		result1 error
	}
	RunStub        func(startAuction models.LRPStartAuction) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		startAuction models.LRPStartAuction
	}
	runReturns struct {
		result1 error
	}
	StopStub        func(stopInstance models.StopLRPInstance) error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		stopInstance models.StopLRPInstance
	}
	stopReturns struct {
		result1 error
	}
}

func (fake *FakeAuctionRepDelegate) RemainingResources() (auctiontypes.Resources, error) {
	fake.remainingResourcesMutex.Lock()
	fake.remainingResourcesArgsForCall = append(fake.remainingResourcesArgsForCall, struct{}{})
	fake.remainingResourcesMutex.Unlock()
	if fake.RemainingResourcesStub != nil {
		return fake.RemainingResourcesStub()
	} else {
		return fake.remainingResourcesReturns.result1, fake.remainingResourcesReturns.result2
	}
}

func (fake *FakeAuctionRepDelegate) RemainingResourcesCallCount() int {
	fake.remainingResourcesMutex.RLock()
	defer fake.remainingResourcesMutex.RUnlock()
	return len(fake.remainingResourcesArgsForCall)
}

func (fake *FakeAuctionRepDelegate) RemainingResourcesReturns(result1 auctiontypes.Resources, result2 error) {
	fake.RemainingResourcesStub = nil
	fake.remainingResourcesReturns = struct {
		result1 auctiontypes.Resources
		result2 error
	}{result1, result2}
}

func (fake *FakeAuctionRepDelegate) TotalResources() (auctiontypes.Resources, error) {
	fake.totalResourcesMutex.Lock()
	fake.totalResourcesArgsForCall = append(fake.totalResourcesArgsForCall, struct{}{})
	fake.totalResourcesMutex.Unlock()
	if fake.TotalResourcesStub != nil {
		return fake.TotalResourcesStub()
	} else {
		return fake.totalResourcesReturns.result1, fake.totalResourcesReturns.result2
	}
}

func (fake *FakeAuctionRepDelegate) TotalResourcesCallCount() int {
	fake.totalResourcesMutex.RLock()
	defer fake.totalResourcesMutex.RUnlock()
	return len(fake.totalResourcesArgsForCall)
}

func (fake *FakeAuctionRepDelegate) TotalResourcesReturns(result1 auctiontypes.Resources, result2 error) {
	fake.TotalResourcesStub = nil
	fake.totalResourcesReturns = struct {
		result1 auctiontypes.Resources
		result2 error
	}{result1, result2}
}

func (fake *FakeAuctionRepDelegate) NumInstancesForProcessGuid(processGuid string) (int, error) {
	fake.numInstancesForProcessGuidMutex.Lock()
	fake.numInstancesForProcessGuidArgsForCall = append(fake.numInstancesForProcessGuidArgsForCall, struct {
		processGuid string
	}{processGuid})
	fake.numInstancesForProcessGuidMutex.Unlock()
	if fake.NumInstancesForProcessGuidStub != nil {
		return fake.NumInstancesForProcessGuidStub(processGuid)
	} else {
		return fake.numInstancesForProcessGuidReturns.result1, fake.numInstancesForProcessGuidReturns.result2
	}
}

func (fake *FakeAuctionRepDelegate) NumInstancesForProcessGuidCallCount() int {
	fake.numInstancesForProcessGuidMutex.RLock()
	defer fake.numInstancesForProcessGuidMutex.RUnlock()
	return len(fake.numInstancesForProcessGuidArgsForCall)
}

func (fake *FakeAuctionRepDelegate) NumInstancesForProcessGuidArgsForCall(i int) string {
	fake.numInstancesForProcessGuidMutex.RLock()
	defer fake.numInstancesForProcessGuidMutex.RUnlock()
	return fake.numInstancesForProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeAuctionRepDelegate) NumInstancesForProcessGuidReturns(result1 int, result2 error) {
	fake.NumInstancesForProcessGuidStub = nil
	fake.numInstancesForProcessGuidReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeAuctionRepDelegate) InstanceGuidsForProcessGuidAndIndex(processGuid string, index int) ([]string, error) {
	fake.instanceGuidsForProcessGuidAndIndexMutex.Lock()
	fake.instanceGuidsForProcessGuidAndIndexArgsForCall = append(fake.instanceGuidsForProcessGuidAndIndexArgsForCall, struct {
		processGuid string
		index       int
	}{processGuid, index})
	fake.instanceGuidsForProcessGuidAndIndexMutex.Unlock()
	if fake.InstanceGuidsForProcessGuidAndIndexStub != nil {
		return fake.InstanceGuidsForProcessGuidAndIndexStub(processGuid, index)
	} else {
		return fake.instanceGuidsForProcessGuidAndIndexReturns.result1, fake.instanceGuidsForProcessGuidAndIndexReturns.result2
	}
}

func (fake *FakeAuctionRepDelegate) InstanceGuidsForProcessGuidAndIndexCallCount() int {
	fake.instanceGuidsForProcessGuidAndIndexMutex.RLock()
	defer fake.instanceGuidsForProcessGuidAndIndexMutex.RUnlock()
	return len(fake.instanceGuidsForProcessGuidAndIndexArgsForCall)
}

func (fake *FakeAuctionRepDelegate) InstanceGuidsForProcessGuidAndIndexArgsForCall(i int) (string, int) {
	fake.instanceGuidsForProcessGuidAndIndexMutex.RLock()
	defer fake.instanceGuidsForProcessGuidAndIndexMutex.RUnlock()
	return fake.instanceGuidsForProcessGuidAndIndexArgsForCall[i].processGuid, fake.instanceGuidsForProcessGuidAndIndexArgsForCall[i].index
}

func (fake *FakeAuctionRepDelegate) InstanceGuidsForProcessGuidAndIndexReturns(result1 []string, result2 error) {
	fake.InstanceGuidsForProcessGuidAndIndexStub = nil
	fake.instanceGuidsForProcessGuidAndIndexReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeAuctionRepDelegate) Reserve(startAuctionInfo auctiontypes.StartAuctionInfo) error {
	fake.reserveMutex.Lock()
	fake.reserveArgsForCall = append(fake.reserveArgsForCall, struct {
		startAuctionInfo auctiontypes.StartAuctionInfo
	}{startAuctionInfo})
	fake.reserveMutex.Unlock()
	if fake.ReserveStub != nil {
		return fake.ReserveStub(startAuctionInfo)
	} else {
		return fake.reserveReturns.result1
	}
}

func (fake *FakeAuctionRepDelegate) ReserveCallCount() int {
	fake.reserveMutex.RLock()
	defer fake.reserveMutex.RUnlock()
	return len(fake.reserveArgsForCall)
}

func (fake *FakeAuctionRepDelegate) ReserveArgsForCall(i int) auctiontypes.StartAuctionInfo {
	fake.reserveMutex.RLock()
	defer fake.reserveMutex.RUnlock()
	return fake.reserveArgsForCall[i].startAuctionInfo
}

func (fake *FakeAuctionRepDelegate) ReserveReturns(result1 error) {
	fake.ReserveStub = nil
	fake.reserveReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuctionRepDelegate) ReleaseReservation(startAuctionInfo auctiontypes.StartAuctionInfo) error {
	fake.releaseReservationMutex.Lock()
	fake.releaseReservationArgsForCall = append(fake.releaseReservationArgsForCall, struct {
		startAuctionInfo auctiontypes.StartAuctionInfo
	}{startAuctionInfo})
	fake.releaseReservationMutex.Unlock()
	if fake.ReleaseReservationStub != nil {
		return fake.ReleaseReservationStub(startAuctionInfo)
	} else {
		return fake.releaseReservationReturns.result1
	}
}

func (fake *FakeAuctionRepDelegate) ReleaseReservationCallCount() int {
	fake.releaseReservationMutex.RLock()
	defer fake.releaseReservationMutex.RUnlock()
	return len(fake.releaseReservationArgsForCall)
}

func (fake *FakeAuctionRepDelegate) ReleaseReservationArgsForCall(i int) auctiontypes.StartAuctionInfo {
	fake.releaseReservationMutex.RLock()
	defer fake.releaseReservationMutex.RUnlock()
	return fake.releaseReservationArgsForCall[i].startAuctionInfo
}

func (fake *FakeAuctionRepDelegate) ReleaseReservationReturns(result1 error) {
	fake.ReleaseReservationStub = nil
	fake.releaseReservationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuctionRepDelegate) Run(startAuction models.LRPStartAuction) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		startAuction models.LRPStartAuction
	}{startAuction})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(startAuction)
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeAuctionRepDelegate) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeAuctionRepDelegate) RunArgsForCall(i int) models.LRPStartAuction {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].startAuction
}

func (fake *FakeAuctionRepDelegate) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuctionRepDelegate) Stop(stopInstance models.StopLRPInstance) error {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		stopInstance models.StopLRPInstance
	}{stopInstance})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub(stopInstance)
	} else {
		return fake.stopReturns.result1
	}
}

func (fake *FakeAuctionRepDelegate) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeAuctionRepDelegate) StopArgsForCall(i int) models.StopLRPInstance {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.stopArgsForCall[i].stopInstance
}

func (fake *FakeAuctionRepDelegate) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

var _ auctiontypes.AuctionRepDelegate = new(FakeAuctionRepDelegate)
