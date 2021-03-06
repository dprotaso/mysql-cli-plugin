// Code generated by counterfeiter. DO NOT EDIT.
package migratefakes

import (
	"sync"
)

type FakeUnpacker struct {
	UnpackStub        func(destDir string) error
	unpackMutex       sync.RWMutex
	unpackArgsForCall []struct {
		destDir string
	}
	unpackReturns struct {
		result1 error
	}
	unpackReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnpacker) Unpack(destDir string) error {
	fake.unpackMutex.Lock()
	ret, specificReturn := fake.unpackReturnsOnCall[len(fake.unpackArgsForCall)]
	fake.unpackArgsForCall = append(fake.unpackArgsForCall, struct {
		destDir string
	}{destDir})
	fake.recordInvocation("Unpack", []interface{}{destDir})
	fake.unpackMutex.Unlock()
	if fake.UnpackStub != nil {
		return fake.UnpackStub(destDir)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unpackReturns.result1
}

func (fake *FakeUnpacker) UnpackCallCount() int {
	fake.unpackMutex.RLock()
	defer fake.unpackMutex.RUnlock()
	return len(fake.unpackArgsForCall)
}

func (fake *FakeUnpacker) UnpackArgsForCall(i int) string {
	fake.unpackMutex.RLock()
	defer fake.unpackMutex.RUnlock()
	return fake.unpackArgsForCall[i].destDir
}

func (fake *FakeUnpacker) UnpackReturns(result1 error) {
	fake.UnpackStub = nil
	fake.unpackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUnpacker) UnpackReturnsOnCall(i int, result1 error) {
	fake.UnpackStub = nil
	if fake.unpackReturnsOnCall == nil {
		fake.unpackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unpackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUnpacker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.unpackMutex.RLock()
	defer fake.unpackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnpacker) recordInvocation(key string, args []interface{}) {
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
