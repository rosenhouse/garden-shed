// This file was generated by counterfeiter
package aufsfakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/docker_drivers/aufs"
)

type FakeLoopMounter struct {
	MountFileStub        func(filePath, destPath string) error
	mountFileMutex       sync.RWMutex
	mountFileArgsForCall []struct {
		filePath string
		destPath string
	}
	mountFileReturns struct {
		result1 error
	}
	UnmountStub        func(path string) error
	unmountMutex       sync.RWMutex
	unmountArgsForCall []struct {
		path string
	}
	unmountReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLoopMounter) MountFile(filePath string, destPath string) error {
	fake.mountFileMutex.Lock()
	fake.mountFileArgsForCall = append(fake.mountFileArgsForCall, struct {
		filePath string
		destPath string
	}{filePath, destPath})
	fake.recordInvocation("MountFile", []interface{}{filePath, destPath})
	fake.mountFileMutex.Unlock()
	if fake.MountFileStub != nil {
		return fake.MountFileStub(filePath, destPath)
	} else {
		return fake.mountFileReturns.result1
	}
}

func (fake *FakeLoopMounter) MountFileCallCount() int {
	fake.mountFileMutex.RLock()
	defer fake.mountFileMutex.RUnlock()
	return len(fake.mountFileArgsForCall)
}

func (fake *FakeLoopMounter) MountFileArgsForCall(i int) (string, string) {
	fake.mountFileMutex.RLock()
	defer fake.mountFileMutex.RUnlock()
	return fake.mountFileArgsForCall[i].filePath, fake.mountFileArgsForCall[i].destPath
}

func (fake *FakeLoopMounter) MountFileReturns(result1 error) {
	fake.MountFileStub = nil
	fake.mountFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLoopMounter) Unmount(path string) error {
	fake.unmountMutex.Lock()
	fake.unmountArgsForCall = append(fake.unmountArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Unmount", []interface{}{path})
	fake.unmountMutex.Unlock()
	if fake.UnmountStub != nil {
		return fake.UnmountStub(path)
	} else {
		return fake.unmountReturns.result1
	}
}

func (fake *FakeLoopMounter) UnmountCallCount() int {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return len(fake.unmountArgsForCall)
}

func (fake *FakeLoopMounter) UnmountArgsForCall(i int) string {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return fake.unmountArgsForCall[i].path
}

func (fake *FakeLoopMounter) UnmountReturns(result1 error) {
	fake.UnmountStub = nil
	fake.unmountReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLoopMounter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mountFileMutex.RLock()
	defer fake.mountFileMutex.RUnlock()
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLoopMounter) recordInvocation(key string, args []interface{}) {
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

var _ aufs.LoopMounter = new(FakeLoopMounter)
