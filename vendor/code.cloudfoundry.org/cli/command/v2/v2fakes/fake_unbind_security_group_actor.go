// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeUnbindSecurityGroupActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	UnbindSecurityGroupByNameAndSpaceStub        func(securityGroupName string, spaceGUID string, lifecycle ccv2.SecurityGroupLifecycle) (v2action.Warnings, error)
	unbindSecurityGroupByNameAndSpaceMutex       sync.RWMutex
	unbindSecurityGroupByNameAndSpaceArgsForCall []struct {
		securityGroupName string
		spaceGUID         string
		lifecycle         ccv2.SecurityGroupLifecycle
	}
	unbindSecurityGroupByNameAndSpaceReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	unbindSecurityGroupByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	UnbindSecurityGroupByNameOrganizationNameAndSpaceNameStub        func(securityGroupName string, orgName string, spaceName string, lifecycle ccv2.SecurityGroupLifecycle) (v2action.Warnings, error)
	unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex       sync.RWMutex
	unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall []struct {
		securityGroupName string
		orgName           string
		spaceName         string
		lifecycle         ccv2.SecurityGroupLifecycle
	}
	unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnbindSecurityGroupActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeUnbindSecurityGroupActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeUnbindSecurityGroupActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeUnbindSecurityGroupActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameAndSpace(securityGroupName string, spaceGUID string, lifecycle ccv2.SecurityGroupLifecycle) (v2action.Warnings, error) {
	fake.unbindSecurityGroupByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.unbindSecurityGroupByNameAndSpaceReturnsOnCall[len(fake.unbindSecurityGroupByNameAndSpaceArgsForCall)]
	fake.unbindSecurityGroupByNameAndSpaceArgsForCall = append(fake.unbindSecurityGroupByNameAndSpaceArgsForCall, struct {
		securityGroupName string
		spaceGUID         string
		lifecycle         ccv2.SecurityGroupLifecycle
	}{securityGroupName, spaceGUID, lifecycle})
	fake.recordInvocation("UnbindSecurityGroupByNameAndSpace", []interface{}{securityGroupName, spaceGUID, lifecycle})
	fake.unbindSecurityGroupByNameAndSpaceMutex.Unlock()
	if fake.UnbindSecurityGroupByNameAndSpaceStub != nil {
		return fake.UnbindSecurityGroupByNameAndSpaceStub(securityGroupName, spaceGUID, lifecycle)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.unbindSecurityGroupByNameAndSpaceReturns.result1, fake.unbindSecurityGroupByNameAndSpaceReturns.result2
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameAndSpaceCallCount() int {
	fake.unbindSecurityGroupByNameAndSpaceMutex.RLock()
	defer fake.unbindSecurityGroupByNameAndSpaceMutex.RUnlock()
	return len(fake.unbindSecurityGroupByNameAndSpaceArgsForCall)
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameAndSpaceArgsForCall(i int) (string, string, ccv2.SecurityGroupLifecycle) {
	fake.unbindSecurityGroupByNameAndSpaceMutex.RLock()
	defer fake.unbindSecurityGroupByNameAndSpaceMutex.RUnlock()
	return fake.unbindSecurityGroupByNameAndSpaceArgsForCall[i].securityGroupName, fake.unbindSecurityGroupByNameAndSpaceArgsForCall[i].spaceGUID, fake.unbindSecurityGroupByNameAndSpaceArgsForCall[i].lifecycle
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameAndSpaceReturns(result1 v2action.Warnings, result2 error) {
	fake.UnbindSecurityGroupByNameAndSpaceStub = nil
	fake.unbindSecurityGroupByNameAndSpaceReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameAndSpaceReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.UnbindSecurityGroupByNameAndSpaceStub = nil
	if fake.unbindSecurityGroupByNameAndSpaceReturnsOnCall == nil {
		fake.unbindSecurityGroupByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.unbindSecurityGroupByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameOrganizationNameAndSpaceName(securityGroupName string, orgName string, spaceName string, lifecycle ccv2.SecurityGroupLifecycle) (v2action.Warnings, error) {
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.Lock()
	ret, specificReturn := fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall[len(fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall)]
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall = append(fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall, struct {
		securityGroupName string
		orgName           string
		spaceName         string
		lifecycle         ccv2.SecurityGroupLifecycle
	}{securityGroupName, orgName, spaceName, lifecycle})
	fake.recordInvocation("UnbindSecurityGroupByNameOrganizationNameAndSpaceName", []interface{}{securityGroupName, orgName, spaceName, lifecycle})
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.Unlock()
	if fake.UnbindSecurityGroupByNameOrganizationNameAndSpaceNameStub != nil {
		return fake.UnbindSecurityGroupByNameOrganizationNameAndSpaceNameStub(securityGroupName, orgName, spaceName, lifecycle)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturns.result1, fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturns.result2
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameOrganizationNameAndSpaceNameCallCount() int {
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RLock()
	defer fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RUnlock()
	return len(fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall)
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall(i int) (string, string, string, ccv2.SecurityGroupLifecycle) {
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RLock()
	defer fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RUnlock()
	return fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall[i].securityGroupName, fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall[i].orgName, fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall[i].spaceName, fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall[i].lifecycle
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameOrganizationNameAndSpaceNameReturns(result1 v2action.Warnings, result2 error) {
	fake.UnbindSecurityGroupByNameOrganizationNameAndSpaceNameStub = nil
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.UnbindSecurityGroupByNameOrganizationNameAndSpaceNameStub = nil
	if fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall == nil {
		fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnbindSecurityGroupActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.unbindSecurityGroupByNameAndSpaceMutex.RLock()
	defer fake.unbindSecurityGroupByNameAndSpaceMutex.RUnlock()
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RLock()
	defer fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnbindSecurityGroupActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.UnbindSecurityGroupActor = new(FakeUnbindSecurityGroupActor)
