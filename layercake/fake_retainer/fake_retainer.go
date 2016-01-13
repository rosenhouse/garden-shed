// This file was generated by counterfeiter
package fake_retainer

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/layercake"
	"github.com/pivotal-golang/lager"
)

type FakeRetainer struct {
	RetainStub        func(log lager.Logger, id layercake.ID)
	retainMutex       sync.RWMutex
	retainArgsForCall []struct {
		log lager.Logger
		id  layercake.ID
	}
}

func (fake *FakeRetainer) Retain(log lager.Logger, id layercake.ID) {
	fake.retainMutex.Lock()
	fake.retainArgsForCall = append(fake.retainArgsForCall, struct {
		log lager.Logger
		id  layercake.ID
	}{log, id})
	fake.retainMutex.Unlock()
	if fake.RetainStub != nil {
		fake.RetainStub(log, id)
	}
}

func (fake *FakeRetainer) RetainCallCount() int {
	fake.retainMutex.RLock()
	defer fake.retainMutex.RUnlock()
	return len(fake.retainArgsForCall)
}

func (fake *FakeRetainer) RetainArgsForCall(i int) (lager.Logger, layercake.ID) {
	fake.retainMutex.RLock()
	defer fake.retainMutex.RUnlock()
	return fake.retainArgsForCall[i].log, fake.retainArgsForCall[i].id
}

var _ layercake.Retainer = new(FakeRetainer)
