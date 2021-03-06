// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type FakeMetadata struct {
	UsesServiceNetworkStub        func() bool
	usesServiceNetworkMutex       sync.RWMutex
	usesServiceNetworkArgsForCall []struct {
	}
	usesServiceNetworkReturns struct {
		result1 bool
	}
	usesServiceNetworkReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMetadata) UsesServiceNetwork() bool {
	fake.usesServiceNetworkMutex.Lock()
	ret, specificReturn := fake.usesServiceNetworkReturnsOnCall[len(fake.usesServiceNetworkArgsForCall)]
	fake.usesServiceNetworkArgsForCall = append(fake.usesServiceNetworkArgsForCall, struct {
	}{})
	fake.recordInvocation("UsesServiceNetwork", []interface{}{})
	fake.usesServiceNetworkMutex.Unlock()
	if fake.UsesServiceNetworkStub != nil {
		return fake.UsesServiceNetworkStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.usesServiceNetworkReturns
	return fakeReturns.result1
}

func (fake *FakeMetadata) UsesServiceNetworkCallCount() int {
	fake.usesServiceNetworkMutex.RLock()
	defer fake.usesServiceNetworkMutex.RUnlock()
	return len(fake.usesServiceNetworkArgsForCall)
}

func (fake *FakeMetadata) UsesServiceNetworkCalls(stub func() bool) {
	fake.usesServiceNetworkMutex.Lock()
	defer fake.usesServiceNetworkMutex.Unlock()
	fake.UsesServiceNetworkStub = stub
}

func (fake *FakeMetadata) UsesServiceNetworkReturns(result1 bool) {
	fake.usesServiceNetworkMutex.Lock()
	defer fake.usesServiceNetworkMutex.Unlock()
	fake.UsesServiceNetworkStub = nil
	fake.usesServiceNetworkReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMetadata) UsesServiceNetworkReturnsOnCall(i int, result1 bool) {
	fake.usesServiceNetworkMutex.Lock()
	defer fake.usesServiceNetworkMutex.Unlock()
	fake.UsesServiceNetworkStub = nil
	if fake.usesServiceNetworkReturnsOnCall == nil {
		fake.usesServiceNetworkReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.usesServiceNetworkReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMetadata) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.usesServiceNetworkMutex.RLock()
	defer fake.usesServiceNetworkMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMetadata) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
