// Code generated by counterfeiter. DO NOT EDIT.
package pluginfakes

import (
	"net/http"
	"sync"

	"code.cloudfoundry.org/cli/api/plugin"
)

type FakeConnectionWrapper struct {
	MakeStub        func(request *http.Request, passedResponse *plugin.Response, proxyReader plugin.ProxyReader) error
	makeMutex       sync.RWMutex
	makeArgsForCall []struct {
		request        *http.Request
		passedResponse *plugin.Response
		proxyReader    plugin.ProxyReader
	}
	makeReturns struct {
		result1 error
	}
	makeReturnsOnCall map[int]struct {
		result1 error
	}
	WrapStub        func(innerconnection plugin.Connection) plugin.Connection
	wrapMutex       sync.RWMutex
	wrapArgsForCall []struct {
		innerconnection plugin.Connection
	}
	wrapReturns struct {
		result1 plugin.Connection
	}
	wrapReturnsOnCall map[int]struct {
		result1 plugin.Connection
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConnectionWrapper) Make(request *http.Request, passedResponse *plugin.Response, proxyReader plugin.ProxyReader) error {
	fake.makeMutex.Lock()
	ret, specificReturn := fake.makeReturnsOnCall[len(fake.makeArgsForCall)]
	fake.makeArgsForCall = append(fake.makeArgsForCall, struct {
		request        *http.Request
		passedResponse *plugin.Response
		proxyReader    plugin.ProxyReader
	}{request, passedResponse, proxyReader})
	fake.recordInvocation("Make", []interface{}{request, passedResponse, proxyReader})
	fake.makeMutex.Unlock()
	if fake.MakeStub != nil {
		return fake.MakeStub(request, passedResponse, proxyReader)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.makeReturns.result1
}

func (fake *FakeConnectionWrapper) MakeCallCount() int {
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	return len(fake.makeArgsForCall)
}

func (fake *FakeConnectionWrapper) MakeArgsForCall(i int) (*http.Request, *plugin.Response, plugin.ProxyReader) {
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	return fake.makeArgsForCall[i].request, fake.makeArgsForCall[i].passedResponse, fake.makeArgsForCall[i].proxyReader
}

func (fake *FakeConnectionWrapper) MakeReturns(result1 error) {
	fake.MakeStub = nil
	fake.makeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnectionWrapper) MakeReturnsOnCall(i int, result1 error) {
	fake.MakeStub = nil
	if fake.makeReturnsOnCall == nil {
		fake.makeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.makeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnectionWrapper) Wrap(innerconnection plugin.Connection) plugin.Connection {
	fake.wrapMutex.Lock()
	ret, specificReturn := fake.wrapReturnsOnCall[len(fake.wrapArgsForCall)]
	fake.wrapArgsForCall = append(fake.wrapArgsForCall, struct {
		innerconnection plugin.Connection
	}{innerconnection})
	fake.recordInvocation("Wrap", []interface{}{innerconnection})
	fake.wrapMutex.Unlock()
	if fake.WrapStub != nil {
		return fake.WrapStub(innerconnection)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.wrapReturns.result1
}

func (fake *FakeConnectionWrapper) WrapCallCount() int {
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	return len(fake.wrapArgsForCall)
}

func (fake *FakeConnectionWrapper) WrapArgsForCall(i int) plugin.Connection {
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	return fake.wrapArgsForCall[i].innerconnection
}

func (fake *FakeConnectionWrapper) WrapReturns(result1 plugin.Connection) {
	fake.WrapStub = nil
	fake.wrapReturns = struct {
		result1 plugin.Connection
	}{result1}
}

func (fake *FakeConnectionWrapper) WrapReturnsOnCall(i int, result1 plugin.Connection) {
	fake.WrapStub = nil
	if fake.wrapReturnsOnCall == nil {
		fake.wrapReturnsOnCall = make(map[int]struct {
			result1 plugin.Connection
		})
	}
	fake.wrapReturnsOnCall[i] = struct {
		result1 plugin.Connection
	}{result1}
}

func (fake *FakeConnectionWrapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConnectionWrapper) recordInvocation(key string, args []interface{}) {
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

var _ plugin.ConnectionWrapper = new(FakeConnectionWrapper)
