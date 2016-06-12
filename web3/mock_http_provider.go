// Copyright (c) 2016, Alan Chen
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package web3

import (
	"fmt"
	"strings"

	"github.com/alanchchen/web3go/rpc"
  	"github.com/stretchr/testify/mock"
)

// MockHTTPProvider provides basic web3 interface
type MockHTTPProvider struct {
	mock mock.Mock
	rpc rpc.RPC
	apis map[string]func(rpc.Request) (rpc.Response, error)
}

// NewMockHTTPProvider creates a HTTP provider mock
func NewMockHTTPProvider() Provider {
	return &MockHTTPProvider{rpc: rpc.GetRPCMethod()}
}

// IsConnected ...
func (provider *MockHTTPProvider) IsConnected() bool {
	return true
}

// Send JSON RPC request through http client
func (provider *MockHTTPProvider) send(request rpc.Request) (response rpc.Response, err error) {
	m := request.Get("method")
	switch m.(type) {
	case string:
		method := m.(string)
		return provider.dispatchMethod(method, request)
	default:
		return nil, fmt.Errorf("Invalid method %v", m)
	}
}

func (provider *MockHTTPProvider) dispatchMethod(method string, request rpc.Request) (response rpc.Response, err error) {
	if index := strings.Index(method, "_"); index > 0 {
		if api, ok := provider.apis[method[:index]]; ok {
			return api(request)
		}
	}

	return nil, fmt.Errorf("Unrecognized method %s", method)
}

func (provider *MockHTTPProvider) getRPCMethod() rpc.RPC {
	return provider.rpc
}