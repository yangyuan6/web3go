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

package provider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yangyuan6/web3go/rpc"
)

// HTTPProvider provides basic web3 interface
type HTTPProvider struct {
	host string
	rpc  rpc.RPC
}

// NewHTTPProvider creates a HTTP provider
func NewHTTPProvider(host string, method rpc.RPC) Provider {
	if !strings.HasPrefix(host, "http://") {
		host = "http://" + host
	}
	if method == nil {
		method = rpc.GetDefaultMethod()
	}
	return &HTTPProvider{host: host, rpc: method}
}

// IsConnected ...
func (provider *HTTPProvider) IsConnected() bool {
	req := provider.rpc.NewRequest("net_listening")
	resp, err := provider.Send(req)
	if err != nil {
		return false
	}
	return resp.Get("result").(bool)
}

// Send JSON RPC request through http client
func (provider *HTTPProvider) Send(request rpc.Request) (response rpc.Response, err error) {
	contentType := provider.determineContentType()
	resp, err := http.Post(provider.host, contentType, strings.NewReader(request.String()))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response = provider.rpc.NewResponse(body)
	if response == nil {
		err = fmt.Errorf("Malformed response body, %s", string(body))
	}
	return response, err
}

func (provider *HTTPProvider) GetRPCMethod() rpc.RPC {
	return provider.rpc
}

func (provider *HTTPProvider) determineContentType() string {
	switch provider.rpc.Name() {
	case "jsonrpc":
		return "application/json"
	default:
		return "application/json"
	}
}
