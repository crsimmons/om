// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type InstallationsService struct {
	TriggerStub        func(bool) (api.InstallationsServiceOutput, error)
	triggerMutex       sync.RWMutex
	triggerArgsForCall []struct {
		arg1 bool
	}
	triggerReturns struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	triggerReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	StatusStub        func(id int) (api.InstallationsServiceOutput, error)
	statusMutex       sync.RWMutex
	statusArgsForCall []struct {
		id int
	}
	statusReturns struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	statusReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	LogsStub        func(id int) (api.InstallationsServiceOutput, error)
	logsMutex       sync.RWMutex
	logsArgsForCall []struct {
		id int
	}
	logsReturns struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	logsReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	RunningInstallationStub        func() (api.InstallationsServiceOutput, error)
	runningInstallationMutex       sync.RWMutex
	runningInstallationArgsForCall []struct{}
	runningInstallationReturns     struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	runningInstallationReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	ListInstallationsStub        func() ([]api.InstallationsServiceOutput, error)
	listInstallationsMutex       sync.RWMutex
	listInstallationsArgsForCall []struct{}
	listInstallationsReturns     struct {
		result1 []api.InstallationsServiceOutput
		result2 error
	}
	listInstallationsReturnsOnCall map[int]struct {
		result1 []api.InstallationsServiceOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *InstallationsService) Trigger(arg1 bool) (api.InstallationsServiceOutput, error) {
	fake.triggerMutex.Lock()
	ret, specificReturn := fake.triggerReturnsOnCall[len(fake.triggerArgsForCall)]
	fake.triggerArgsForCall = append(fake.triggerArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("Trigger", []interface{}{arg1})
	fake.triggerMutex.Unlock()
	if fake.TriggerStub != nil {
		return fake.TriggerStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.triggerReturns.result1, fake.triggerReturns.result2
}

func (fake *InstallationsService) TriggerCallCount() int {
	fake.triggerMutex.RLock()
	defer fake.triggerMutex.RUnlock()
	return len(fake.triggerArgsForCall)
}

func (fake *InstallationsService) TriggerArgsForCall(i int) bool {
	fake.triggerMutex.RLock()
	defer fake.triggerMutex.RUnlock()
	return fake.triggerArgsForCall[i].arg1
}

func (fake *InstallationsService) TriggerReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.TriggerStub = nil
	fake.triggerReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *InstallationsService) TriggerReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.TriggerStub = nil
	if fake.triggerReturnsOnCall == nil {
		fake.triggerReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.triggerReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *InstallationsService) Status(id int) (api.InstallationsServiceOutput, error) {
	fake.statusMutex.Lock()
	ret, specificReturn := fake.statusReturnsOnCall[len(fake.statusArgsForCall)]
	fake.statusArgsForCall = append(fake.statusArgsForCall, struct {
		id int
	}{id})
	fake.recordInvocation("Status", []interface{}{id})
	fake.statusMutex.Unlock()
	if fake.StatusStub != nil {
		return fake.StatusStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.statusReturns.result1, fake.statusReturns.result2
}

func (fake *InstallationsService) StatusCallCount() int {
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	return len(fake.statusArgsForCall)
}

func (fake *InstallationsService) StatusArgsForCall(i int) int {
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	return fake.statusArgsForCall[i].id
}

func (fake *InstallationsService) StatusReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.StatusStub = nil
	fake.statusReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *InstallationsService) StatusReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.StatusStub = nil
	if fake.statusReturnsOnCall == nil {
		fake.statusReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.statusReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *InstallationsService) Logs(id int) (api.InstallationsServiceOutput, error) {
	fake.logsMutex.Lock()
	ret, specificReturn := fake.logsReturnsOnCall[len(fake.logsArgsForCall)]
	fake.logsArgsForCall = append(fake.logsArgsForCall, struct {
		id int
	}{id})
	fake.recordInvocation("Logs", []interface{}{id})
	fake.logsMutex.Unlock()
	if fake.LogsStub != nil {
		return fake.LogsStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.logsReturns.result1, fake.logsReturns.result2
}

func (fake *InstallationsService) LogsCallCount() int {
	fake.logsMutex.RLock()
	defer fake.logsMutex.RUnlock()
	return len(fake.logsArgsForCall)
}

func (fake *InstallationsService) LogsArgsForCall(i int) int {
	fake.logsMutex.RLock()
	defer fake.logsMutex.RUnlock()
	return fake.logsArgsForCall[i].id
}

func (fake *InstallationsService) LogsReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.LogsStub = nil
	fake.logsReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *InstallationsService) LogsReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.LogsStub = nil
	if fake.logsReturnsOnCall == nil {
		fake.logsReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.logsReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *InstallationsService) RunningInstallation() (api.InstallationsServiceOutput, error) {
	fake.runningInstallationMutex.Lock()
	ret, specificReturn := fake.runningInstallationReturnsOnCall[len(fake.runningInstallationArgsForCall)]
	fake.runningInstallationArgsForCall = append(fake.runningInstallationArgsForCall, struct{}{})
	fake.recordInvocation("RunningInstallation", []interface{}{})
	fake.runningInstallationMutex.Unlock()
	if fake.RunningInstallationStub != nil {
		return fake.RunningInstallationStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runningInstallationReturns.result1, fake.runningInstallationReturns.result2
}

func (fake *InstallationsService) RunningInstallationCallCount() int {
	fake.runningInstallationMutex.RLock()
	defer fake.runningInstallationMutex.RUnlock()
	return len(fake.runningInstallationArgsForCall)
}

func (fake *InstallationsService) RunningInstallationReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.RunningInstallationStub = nil
	fake.runningInstallationReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *InstallationsService) RunningInstallationReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.RunningInstallationStub = nil
	if fake.runningInstallationReturnsOnCall == nil {
		fake.runningInstallationReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.runningInstallationReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *InstallationsService) ListInstallations() ([]api.InstallationsServiceOutput, error) {
	fake.listInstallationsMutex.Lock()
	ret, specificReturn := fake.listInstallationsReturnsOnCall[len(fake.listInstallationsArgsForCall)]
	fake.listInstallationsArgsForCall = append(fake.listInstallationsArgsForCall, struct{}{})
	fake.recordInvocation("ListInstallations", []interface{}{})
	fake.listInstallationsMutex.Unlock()
	if fake.ListInstallationsStub != nil {
		return fake.ListInstallationsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listInstallationsReturns.result1, fake.listInstallationsReturns.result2
}

func (fake *InstallationsService) ListInstallationsCallCount() int {
	fake.listInstallationsMutex.RLock()
	defer fake.listInstallationsMutex.RUnlock()
	return len(fake.listInstallationsArgsForCall)
}

func (fake *InstallationsService) ListInstallationsReturns(result1 []api.InstallationsServiceOutput, result2 error) {
	fake.ListInstallationsStub = nil
	fake.listInstallationsReturns = struct {
		result1 []api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *InstallationsService) ListInstallationsReturnsOnCall(i int, result1 []api.InstallationsServiceOutput, result2 error) {
	fake.ListInstallationsStub = nil
	if fake.listInstallationsReturnsOnCall == nil {
		fake.listInstallationsReturnsOnCall = make(map[int]struct {
			result1 []api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.listInstallationsReturnsOnCall[i] = struct {
		result1 []api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *InstallationsService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.triggerMutex.RLock()
	defer fake.triggerMutex.RUnlock()
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	fake.logsMutex.RLock()
	defer fake.logsMutex.RUnlock()
	fake.runningInstallationMutex.RLock()
	defer fake.runningInstallationMutex.RUnlock()
	fake.listInstallationsMutex.RLock()
	defer fake.listInstallationsMutex.RUnlock()
	return fake.invocations
}

func (fake *InstallationsService) recordInvocation(key string, args []interface{}) {
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
