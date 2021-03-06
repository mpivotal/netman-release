// This file was generated by counterfeiter
package fakes

import "sync"

type CCClient struct {
	GetAllAppGUIDsStub        func(token string) (map[string]interface{}, error)
	getAllAppGUIDsMutex       sync.RWMutex
	getAllAppGUIDsArgsForCall []struct {
		token string
	}
	getAllAppGUIDsReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CCClient) GetAllAppGUIDs(token string) (map[string]interface{}, error) {
	fake.getAllAppGUIDsMutex.Lock()
	fake.getAllAppGUIDsArgsForCall = append(fake.getAllAppGUIDsArgsForCall, struct {
		token string
	}{token})
	fake.recordInvocation("GetAllAppGUIDs", []interface{}{token})
	fake.getAllAppGUIDsMutex.Unlock()
	if fake.GetAllAppGUIDsStub != nil {
		return fake.GetAllAppGUIDsStub(token)
	} else {
		return fake.getAllAppGUIDsReturns.result1, fake.getAllAppGUIDsReturns.result2
	}
}

func (fake *CCClient) GetAllAppGUIDsCallCount() int {
	fake.getAllAppGUIDsMutex.RLock()
	defer fake.getAllAppGUIDsMutex.RUnlock()
	return len(fake.getAllAppGUIDsArgsForCall)
}

func (fake *CCClient) GetAllAppGUIDsArgsForCall(i int) string {
	fake.getAllAppGUIDsMutex.RLock()
	defer fake.getAllAppGUIDsMutex.RUnlock()
	return fake.getAllAppGUIDsArgsForCall[i].token
}

func (fake *CCClient) GetAllAppGUIDsReturns(result1 map[string]interface{}, result2 error) {
	fake.GetAllAppGUIDsStub = nil
	fake.getAllAppGUIDsReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *CCClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAllAppGUIDsMutex.RLock()
	defer fake.getAllAppGUIDsMutex.RUnlock()
	return fake.invocations
}

func (fake *CCClient) recordInvocation(key string, args []interface{}) {
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
