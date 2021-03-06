// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	v1api "github.com/m3o/services/v1api/proto"
	"github.com/micro/micro/v3/service/client"
)

type FakeV1Service struct {
	BlockKeyStub        func(context.Context, *v1api.BlockKeyRequest, ...client.CallOption) (*v1api.BlockKeyResponse, error)
	blockKeyMutex       sync.RWMutex
	blockKeyArgsForCall []struct {
		arg1 context.Context
		arg2 *v1api.BlockKeyRequest
		arg3 []client.CallOption
	}
	blockKeyReturns struct {
		result1 *v1api.BlockKeyResponse
		result2 error
	}
	blockKeyReturnsOnCall map[int]struct {
		result1 *v1api.BlockKeyResponse
		result2 error
	}
	DisableAPIStub        func(context.Context, *v1api.DisableAPIRequest, ...client.CallOption) (*v1api.DisableAPIResponse, error)
	disableAPIMutex       sync.RWMutex
	disableAPIArgsForCall []struct {
		arg1 context.Context
		arg2 *v1api.DisableAPIRequest
		arg3 []client.CallOption
	}
	disableAPIReturns struct {
		result1 *v1api.DisableAPIResponse
		result2 error
	}
	disableAPIReturnsOnCall map[int]struct {
		result1 *v1api.DisableAPIResponse
		result2 error
	}
	EnableAPIStub        func(context.Context, *v1api.EnableAPIRequest, ...client.CallOption) (*v1api.EnableAPIResponse, error)
	enableAPIMutex       sync.RWMutex
	enableAPIArgsForCall []struct {
		arg1 context.Context
		arg2 *v1api.EnableAPIRequest
		arg3 []client.CallOption
	}
	enableAPIReturns struct {
		result1 *v1api.EnableAPIResponse
		result2 error
	}
	enableAPIReturnsOnCall map[int]struct {
		result1 *v1api.EnableAPIResponse
		result2 error
	}
	GenerateKeyStub        func(context.Context, *v1api.GenerateKeyRequest, ...client.CallOption) (*v1api.GenerateKeyResponse, error)
	generateKeyMutex       sync.RWMutex
	generateKeyArgsForCall []struct {
		arg1 context.Context
		arg2 *v1api.GenerateKeyRequest
		arg3 []client.CallOption
	}
	generateKeyReturns struct {
		result1 *v1api.GenerateKeyResponse
		result2 error
	}
	generateKeyReturnsOnCall map[int]struct {
		result1 *v1api.GenerateKeyResponse
		result2 error
	}
	ListAPIsStub        func(context.Context, *v1api.ListAPIsRequest, ...client.CallOption) (*v1api.ListAPIsResponse, error)
	listAPIsMutex       sync.RWMutex
	listAPIsArgsForCall []struct {
		arg1 context.Context
		arg2 *v1api.ListAPIsRequest
		arg3 []client.CallOption
	}
	listAPIsReturns struct {
		result1 *v1api.ListAPIsResponse
		result2 error
	}
	listAPIsReturnsOnCall map[int]struct {
		result1 *v1api.ListAPIsResponse
		result2 error
	}
	ListKeysStub        func(context.Context, *v1api.ListRequest, ...client.CallOption) (*v1api.ListResponse, error)
	listKeysMutex       sync.RWMutex
	listKeysArgsForCall []struct {
		arg1 context.Context
		arg2 *v1api.ListRequest
		arg3 []client.CallOption
	}
	listKeysReturns struct {
		result1 *v1api.ListResponse
		result2 error
	}
	listKeysReturnsOnCall map[int]struct {
		result1 *v1api.ListResponse
		result2 error
	}
	RevokeKeyStub        func(context.Context, *v1api.RevokeRequest, ...client.CallOption) (*v1api.RevokeResponse, error)
	revokeKeyMutex       sync.RWMutex
	revokeKeyArgsForCall []struct {
		arg1 context.Context
		arg2 *v1api.RevokeRequest
		arg3 []client.CallOption
	}
	revokeKeyReturns struct {
		result1 *v1api.RevokeResponse
		result2 error
	}
	revokeKeyReturnsOnCall map[int]struct {
		result1 *v1api.RevokeResponse
		result2 error
	}
	UnblockKeyStub        func(context.Context, *v1api.UnblockKeyRequest, ...client.CallOption) (*v1api.UnblockKeyResponse, error)
	unblockKeyMutex       sync.RWMutex
	unblockKeyArgsForCall []struct {
		arg1 context.Context
		arg2 *v1api.UnblockKeyRequest
		arg3 []client.CallOption
	}
	unblockKeyReturns struct {
		result1 *v1api.UnblockKeyResponse
		result2 error
	}
	unblockKeyReturnsOnCall map[int]struct {
		result1 *v1api.UnblockKeyResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV1Service) BlockKey(arg1 context.Context, arg2 *v1api.BlockKeyRequest, arg3 ...client.CallOption) (*v1api.BlockKeyResponse, error) {
	fake.blockKeyMutex.Lock()
	ret, specificReturn := fake.blockKeyReturnsOnCall[len(fake.blockKeyArgsForCall)]
	fake.blockKeyArgsForCall = append(fake.blockKeyArgsForCall, struct {
		arg1 context.Context
		arg2 *v1api.BlockKeyRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.BlockKeyStub
	fakeReturns := fake.blockKeyReturns
	fake.recordInvocation("BlockKey", []interface{}{arg1, arg2, arg3})
	fake.blockKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV1Service) BlockKeyCallCount() int {
	fake.blockKeyMutex.RLock()
	defer fake.blockKeyMutex.RUnlock()
	return len(fake.blockKeyArgsForCall)
}

func (fake *FakeV1Service) BlockKeyCalls(stub func(context.Context, *v1api.BlockKeyRequest, ...client.CallOption) (*v1api.BlockKeyResponse, error)) {
	fake.blockKeyMutex.Lock()
	defer fake.blockKeyMutex.Unlock()
	fake.BlockKeyStub = stub
}

func (fake *FakeV1Service) BlockKeyArgsForCall(i int) (context.Context, *v1api.BlockKeyRequest, []client.CallOption) {
	fake.blockKeyMutex.RLock()
	defer fake.blockKeyMutex.RUnlock()
	argsForCall := fake.blockKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV1Service) BlockKeyReturns(result1 *v1api.BlockKeyResponse, result2 error) {
	fake.blockKeyMutex.Lock()
	defer fake.blockKeyMutex.Unlock()
	fake.BlockKeyStub = nil
	fake.blockKeyReturns = struct {
		result1 *v1api.BlockKeyResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) BlockKeyReturnsOnCall(i int, result1 *v1api.BlockKeyResponse, result2 error) {
	fake.blockKeyMutex.Lock()
	defer fake.blockKeyMutex.Unlock()
	fake.BlockKeyStub = nil
	if fake.blockKeyReturnsOnCall == nil {
		fake.blockKeyReturnsOnCall = make(map[int]struct {
			result1 *v1api.BlockKeyResponse
			result2 error
		})
	}
	fake.blockKeyReturnsOnCall[i] = struct {
		result1 *v1api.BlockKeyResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) DisableAPI(arg1 context.Context, arg2 *v1api.DisableAPIRequest, arg3 ...client.CallOption) (*v1api.DisableAPIResponse, error) {
	fake.disableAPIMutex.Lock()
	ret, specificReturn := fake.disableAPIReturnsOnCall[len(fake.disableAPIArgsForCall)]
	fake.disableAPIArgsForCall = append(fake.disableAPIArgsForCall, struct {
		arg1 context.Context
		arg2 *v1api.DisableAPIRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.DisableAPIStub
	fakeReturns := fake.disableAPIReturns
	fake.recordInvocation("DisableAPI", []interface{}{arg1, arg2, arg3})
	fake.disableAPIMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV1Service) DisableAPICallCount() int {
	fake.disableAPIMutex.RLock()
	defer fake.disableAPIMutex.RUnlock()
	return len(fake.disableAPIArgsForCall)
}

func (fake *FakeV1Service) DisableAPICalls(stub func(context.Context, *v1api.DisableAPIRequest, ...client.CallOption) (*v1api.DisableAPIResponse, error)) {
	fake.disableAPIMutex.Lock()
	defer fake.disableAPIMutex.Unlock()
	fake.DisableAPIStub = stub
}

func (fake *FakeV1Service) DisableAPIArgsForCall(i int) (context.Context, *v1api.DisableAPIRequest, []client.CallOption) {
	fake.disableAPIMutex.RLock()
	defer fake.disableAPIMutex.RUnlock()
	argsForCall := fake.disableAPIArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV1Service) DisableAPIReturns(result1 *v1api.DisableAPIResponse, result2 error) {
	fake.disableAPIMutex.Lock()
	defer fake.disableAPIMutex.Unlock()
	fake.DisableAPIStub = nil
	fake.disableAPIReturns = struct {
		result1 *v1api.DisableAPIResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) DisableAPIReturnsOnCall(i int, result1 *v1api.DisableAPIResponse, result2 error) {
	fake.disableAPIMutex.Lock()
	defer fake.disableAPIMutex.Unlock()
	fake.DisableAPIStub = nil
	if fake.disableAPIReturnsOnCall == nil {
		fake.disableAPIReturnsOnCall = make(map[int]struct {
			result1 *v1api.DisableAPIResponse
			result2 error
		})
	}
	fake.disableAPIReturnsOnCall[i] = struct {
		result1 *v1api.DisableAPIResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) EnableAPI(arg1 context.Context, arg2 *v1api.EnableAPIRequest, arg3 ...client.CallOption) (*v1api.EnableAPIResponse, error) {
	fake.enableAPIMutex.Lock()
	ret, specificReturn := fake.enableAPIReturnsOnCall[len(fake.enableAPIArgsForCall)]
	fake.enableAPIArgsForCall = append(fake.enableAPIArgsForCall, struct {
		arg1 context.Context
		arg2 *v1api.EnableAPIRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.EnableAPIStub
	fakeReturns := fake.enableAPIReturns
	fake.recordInvocation("EnableAPI", []interface{}{arg1, arg2, arg3})
	fake.enableAPIMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV1Service) EnableAPICallCount() int {
	fake.enableAPIMutex.RLock()
	defer fake.enableAPIMutex.RUnlock()
	return len(fake.enableAPIArgsForCall)
}

func (fake *FakeV1Service) EnableAPICalls(stub func(context.Context, *v1api.EnableAPIRequest, ...client.CallOption) (*v1api.EnableAPIResponse, error)) {
	fake.enableAPIMutex.Lock()
	defer fake.enableAPIMutex.Unlock()
	fake.EnableAPIStub = stub
}

func (fake *FakeV1Service) EnableAPIArgsForCall(i int) (context.Context, *v1api.EnableAPIRequest, []client.CallOption) {
	fake.enableAPIMutex.RLock()
	defer fake.enableAPIMutex.RUnlock()
	argsForCall := fake.enableAPIArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV1Service) EnableAPIReturns(result1 *v1api.EnableAPIResponse, result2 error) {
	fake.enableAPIMutex.Lock()
	defer fake.enableAPIMutex.Unlock()
	fake.EnableAPIStub = nil
	fake.enableAPIReturns = struct {
		result1 *v1api.EnableAPIResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) EnableAPIReturnsOnCall(i int, result1 *v1api.EnableAPIResponse, result2 error) {
	fake.enableAPIMutex.Lock()
	defer fake.enableAPIMutex.Unlock()
	fake.EnableAPIStub = nil
	if fake.enableAPIReturnsOnCall == nil {
		fake.enableAPIReturnsOnCall = make(map[int]struct {
			result1 *v1api.EnableAPIResponse
			result2 error
		})
	}
	fake.enableAPIReturnsOnCall[i] = struct {
		result1 *v1api.EnableAPIResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) GenerateKey(arg1 context.Context, arg2 *v1api.GenerateKeyRequest, arg3 ...client.CallOption) (*v1api.GenerateKeyResponse, error) {
	fake.generateKeyMutex.Lock()
	ret, specificReturn := fake.generateKeyReturnsOnCall[len(fake.generateKeyArgsForCall)]
	fake.generateKeyArgsForCall = append(fake.generateKeyArgsForCall, struct {
		arg1 context.Context
		arg2 *v1api.GenerateKeyRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.GenerateKeyStub
	fakeReturns := fake.generateKeyReturns
	fake.recordInvocation("GenerateKey", []interface{}{arg1, arg2, arg3})
	fake.generateKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV1Service) GenerateKeyCallCount() int {
	fake.generateKeyMutex.RLock()
	defer fake.generateKeyMutex.RUnlock()
	return len(fake.generateKeyArgsForCall)
}

func (fake *FakeV1Service) GenerateKeyCalls(stub func(context.Context, *v1api.GenerateKeyRequest, ...client.CallOption) (*v1api.GenerateKeyResponse, error)) {
	fake.generateKeyMutex.Lock()
	defer fake.generateKeyMutex.Unlock()
	fake.GenerateKeyStub = stub
}

func (fake *FakeV1Service) GenerateKeyArgsForCall(i int) (context.Context, *v1api.GenerateKeyRequest, []client.CallOption) {
	fake.generateKeyMutex.RLock()
	defer fake.generateKeyMutex.RUnlock()
	argsForCall := fake.generateKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV1Service) GenerateKeyReturns(result1 *v1api.GenerateKeyResponse, result2 error) {
	fake.generateKeyMutex.Lock()
	defer fake.generateKeyMutex.Unlock()
	fake.GenerateKeyStub = nil
	fake.generateKeyReturns = struct {
		result1 *v1api.GenerateKeyResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) GenerateKeyReturnsOnCall(i int, result1 *v1api.GenerateKeyResponse, result2 error) {
	fake.generateKeyMutex.Lock()
	defer fake.generateKeyMutex.Unlock()
	fake.GenerateKeyStub = nil
	if fake.generateKeyReturnsOnCall == nil {
		fake.generateKeyReturnsOnCall = make(map[int]struct {
			result1 *v1api.GenerateKeyResponse
			result2 error
		})
	}
	fake.generateKeyReturnsOnCall[i] = struct {
		result1 *v1api.GenerateKeyResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) ListAPIs(arg1 context.Context, arg2 *v1api.ListAPIsRequest, arg3 ...client.CallOption) (*v1api.ListAPIsResponse, error) {
	fake.listAPIsMutex.Lock()
	ret, specificReturn := fake.listAPIsReturnsOnCall[len(fake.listAPIsArgsForCall)]
	fake.listAPIsArgsForCall = append(fake.listAPIsArgsForCall, struct {
		arg1 context.Context
		arg2 *v1api.ListAPIsRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.ListAPIsStub
	fakeReturns := fake.listAPIsReturns
	fake.recordInvocation("ListAPIs", []interface{}{arg1, arg2, arg3})
	fake.listAPIsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV1Service) ListAPIsCallCount() int {
	fake.listAPIsMutex.RLock()
	defer fake.listAPIsMutex.RUnlock()
	return len(fake.listAPIsArgsForCall)
}

func (fake *FakeV1Service) ListAPIsCalls(stub func(context.Context, *v1api.ListAPIsRequest, ...client.CallOption) (*v1api.ListAPIsResponse, error)) {
	fake.listAPIsMutex.Lock()
	defer fake.listAPIsMutex.Unlock()
	fake.ListAPIsStub = stub
}

func (fake *FakeV1Service) ListAPIsArgsForCall(i int) (context.Context, *v1api.ListAPIsRequest, []client.CallOption) {
	fake.listAPIsMutex.RLock()
	defer fake.listAPIsMutex.RUnlock()
	argsForCall := fake.listAPIsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV1Service) ListAPIsReturns(result1 *v1api.ListAPIsResponse, result2 error) {
	fake.listAPIsMutex.Lock()
	defer fake.listAPIsMutex.Unlock()
	fake.ListAPIsStub = nil
	fake.listAPIsReturns = struct {
		result1 *v1api.ListAPIsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) ListAPIsReturnsOnCall(i int, result1 *v1api.ListAPIsResponse, result2 error) {
	fake.listAPIsMutex.Lock()
	defer fake.listAPIsMutex.Unlock()
	fake.ListAPIsStub = nil
	if fake.listAPIsReturnsOnCall == nil {
		fake.listAPIsReturnsOnCall = make(map[int]struct {
			result1 *v1api.ListAPIsResponse
			result2 error
		})
	}
	fake.listAPIsReturnsOnCall[i] = struct {
		result1 *v1api.ListAPIsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) ListKeys(arg1 context.Context, arg2 *v1api.ListRequest, arg3 ...client.CallOption) (*v1api.ListResponse, error) {
	fake.listKeysMutex.Lock()
	ret, specificReturn := fake.listKeysReturnsOnCall[len(fake.listKeysArgsForCall)]
	fake.listKeysArgsForCall = append(fake.listKeysArgsForCall, struct {
		arg1 context.Context
		arg2 *v1api.ListRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.ListKeysStub
	fakeReturns := fake.listKeysReturns
	fake.recordInvocation("ListKeys", []interface{}{arg1, arg2, arg3})
	fake.listKeysMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV1Service) ListKeysCallCount() int {
	fake.listKeysMutex.RLock()
	defer fake.listKeysMutex.RUnlock()
	return len(fake.listKeysArgsForCall)
}

func (fake *FakeV1Service) ListKeysCalls(stub func(context.Context, *v1api.ListRequest, ...client.CallOption) (*v1api.ListResponse, error)) {
	fake.listKeysMutex.Lock()
	defer fake.listKeysMutex.Unlock()
	fake.ListKeysStub = stub
}

func (fake *FakeV1Service) ListKeysArgsForCall(i int) (context.Context, *v1api.ListRequest, []client.CallOption) {
	fake.listKeysMutex.RLock()
	defer fake.listKeysMutex.RUnlock()
	argsForCall := fake.listKeysArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV1Service) ListKeysReturns(result1 *v1api.ListResponse, result2 error) {
	fake.listKeysMutex.Lock()
	defer fake.listKeysMutex.Unlock()
	fake.ListKeysStub = nil
	fake.listKeysReturns = struct {
		result1 *v1api.ListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) ListKeysReturnsOnCall(i int, result1 *v1api.ListResponse, result2 error) {
	fake.listKeysMutex.Lock()
	defer fake.listKeysMutex.Unlock()
	fake.ListKeysStub = nil
	if fake.listKeysReturnsOnCall == nil {
		fake.listKeysReturnsOnCall = make(map[int]struct {
			result1 *v1api.ListResponse
			result2 error
		})
	}
	fake.listKeysReturnsOnCall[i] = struct {
		result1 *v1api.ListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) RevokeKey(arg1 context.Context, arg2 *v1api.RevokeRequest, arg3 ...client.CallOption) (*v1api.RevokeResponse, error) {
	fake.revokeKeyMutex.Lock()
	ret, specificReturn := fake.revokeKeyReturnsOnCall[len(fake.revokeKeyArgsForCall)]
	fake.revokeKeyArgsForCall = append(fake.revokeKeyArgsForCall, struct {
		arg1 context.Context
		arg2 *v1api.RevokeRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.RevokeKeyStub
	fakeReturns := fake.revokeKeyReturns
	fake.recordInvocation("RevokeKey", []interface{}{arg1, arg2, arg3})
	fake.revokeKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV1Service) RevokeKeyCallCount() int {
	fake.revokeKeyMutex.RLock()
	defer fake.revokeKeyMutex.RUnlock()
	return len(fake.revokeKeyArgsForCall)
}

func (fake *FakeV1Service) RevokeKeyCalls(stub func(context.Context, *v1api.RevokeRequest, ...client.CallOption) (*v1api.RevokeResponse, error)) {
	fake.revokeKeyMutex.Lock()
	defer fake.revokeKeyMutex.Unlock()
	fake.RevokeKeyStub = stub
}

func (fake *FakeV1Service) RevokeKeyArgsForCall(i int) (context.Context, *v1api.RevokeRequest, []client.CallOption) {
	fake.revokeKeyMutex.RLock()
	defer fake.revokeKeyMutex.RUnlock()
	argsForCall := fake.revokeKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV1Service) RevokeKeyReturns(result1 *v1api.RevokeResponse, result2 error) {
	fake.revokeKeyMutex.Lock()
	defer fake.revokeKeyMutex.Unlock()
	fake.RevokeKeyStub = nil
	fake.revokeKeyReturns = struct {
		result1 *v1api.RevokeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) RevokeKeyReturnsOnCall(i int, result1 *v1api.RevokeResponse, result2 error) {
	fake.revokeKeyMutex.Lock()
	defer fake.revokeKeyMutex.Unlock()
	fake.RevokeKeyStub = nil
	if fake.revokeKeyReturnsOnCall == nil {
		fake.revokeKeyReturnsOnCall = make(map[int]struct {
			result1 *v1api.RevokeResponse
			result2 error
		})
	}
	fake.revokeKeyReturnsOnCall[i] = struct {
		result1 *v1api.RevokeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) UnblockKey(arg1 context.Context, arg2 *v1api.UnblockKeyRequest, arg3 ...client.CallOption) (*v1api.UnblockKeyResponse, error) {
	fake.unblockKeyMutex.Lock()
	ret, specificReturn := fake.unblockKeyReturnsOnCall[len(fake.unblockKeyArgsForCall)]
	fake.unblockKeyArgsForCall = append(fake.unblockKeyArgsForCall, struct {
		arg1 context.Context
		arg2 *v1api.UnblockKeyRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.UnblockKeyStub
	fakeReturns := fake.unblockKeyReturns
	fake.recordInvocation("UnblockKey", []interface{}{arg1, arg2, arg3})
	fake.unblockKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV1Service) UnblockKeyCallCount() int {
	fake.unblockKeyMutex.RLock()
	defer fake.unblockKeyMutex.RUnlock()
	return len(fake.unblockKeyArgsForCall)
}

func (fake *FakeV1Service) UnblockKeyCalls(stub func(context.Context, *v1api.UnblockKeyRequest, ...client.CallOption) (*v1api.UnblockKeyResponse, error)) {
	fake.unblockKeyMutex.Lock()
	defer fake.unblockKeyMutex.Unlock()
	fake.UnblockKeyStub = stub
}

func (fake *FakeV1Service) UnblockKeyArgsForCall(i int) (context.Context, *v1api.UnblockKeyRequest, []client.CallOption) {
	fake.unblockKeyMutex.RLock()
	defer fake.unblockKeyMutex.RUnlock()
	argsForCall := fake.unblockKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV1Service) UnblockKeyReturns(result1 *v1api.UnblockKeyResponse, result2 error) {
	fake.unblockKeyMutex.Lock()
	defer fake.unblockKeyMutex.Unlock()
	fake.UnblockKeyStub = nil
	fake.unblockKeyReturns = struct {
		result1 *v1api.UnblockKeyResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) UnblockKeyReturnsOnCall(i int, result1 *v1api.UnblockKeyResponse, result2 error) {
	fake.unblockKeyMutex.Lock()
	defer fake.unblockKeyMutex.Unlock()
	fake.UnblockKeyStub = nil
	if fake.unblockKeyReturnsOnCall == nil {
		fake.unblockKeyReturnsOnCall = make(map[int]struct {
			result1 *v1api.UnblockKeyResponse
			result2 error
		})
	}
	fake.unblockKeyReturnsOnCall[i] = struct {
		result1 *v1api.UnblockKeyResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeV1Service) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blockKeyMutex.RLock()
	defer fake.blockKeyMutex.RUnlock()
	fake.disableAPIMutex.RLock()
	defer fake.disableAPIMutex.RUnlock()
	fake.enableAPIMutex.RLock()
	defer fake.enableAPIMutex.RUnlock()
	fake.generateKeyMutex.RLock()
	defer fake.generateKeyMutex.RUnlock()
	fake.listAPIsMutex.RLock()
	defer fake.listAPIsMutex.RUnlock()
	fake.listKeysMutex.RLock()
	defer fake.listKeysMutex.RUnlock()
	fake.revokeKeyMutex.RLock()
	defer fake.revokeKeyMutex.RUnlock()
	fake.unblockKeyMutex.RLock()
	defer fake.unblockKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV1Service) recordInvocation(key string, args []interface{}) {
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

var _ v1api.V1Service = new(FakeV1Service)
