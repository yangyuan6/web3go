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
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	"github.com/yangyuan6/web3go/common"
	"github.com/yangyuan6/web3go/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EthTestSuite struct {
	suite.Suite
	web3 *Web3
	eth  Eth
}

func (suite *EthTestSuite) Test_ProcotolVersion() {
	eth := suite.eth
	result, err := eth.ProtocolVersion()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.NotEqual(suite.T(), "", result, "version is empty")
}

func (suite *EthTestSuite) Test_Syncing() {
	eth := suite.eth
	status, err := eth.Syncing()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.Exactly(suite.T(), false, status.Result, "should be false")
}

func (suite *EthTestSuite) Test_Coinbase() {
	eth := suite.eth
	address, err := eth.Coinbase()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(), "0x407d73d8a49eeb85d32cf465507dd71d507100c1", address.String(), "should be equal")
}

func (suite *EthTestSuite) Test_Mining() {
	eth := suite.eth
	mining, err := eth.Mining()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(), true, mining, "should be equal")
}

func (suite *EthTestSuite) Test_HashRate() {
	eth := suite.eth
	hashrate, err := eth.HashRate()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(), 0x38a, hashrate, "Should be equal")
}

func (suite *EthTestSuite) Test_GasPrice() {
	eth := suite.eth
	price, err := eth.GasPrice()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(), big.NewInt(0x09184e72a000), price, "Should be equal")
}

func (suite *EthTestSuite) Test_Accounts() {
	eth := suite.eth
	accounts, err := eth.Accounts()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(), []common.Address{
		common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")),
		common.NewAddress(common.HexToBytes("0x407d73d8a49ee783afd32cf465507dd71d507100")),
	}, accounts, "Should be equal")
}

func (suite *EthTestSuite) Test_BlockNumber() {
	eth := suite.eth
	blockNumber, err := eth.BlockNumber()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(), big.NewInt(0x4b7), blockNumber, "Should be equal")
}

func (suite *EthTestSuite) Test_GetBalance() {
	eth := suite.eth
	balance, err := eth.GetBalance(common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")), "latest")
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		big.NewInt(0x0234c8a3397aab58),
		balance,
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetStorageAt() {
	eth := suite.eth
	storage, err := eth.GetStorageAt(common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")), 0, "latest")
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		0x03,
		storage,
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetTransactionCount() {
	eth := suite.eth
	transactionCount, err := eth.GetTransactionCount(common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")), "latest")
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		big.NewInt(0x1),
		transactionCount,
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetBlockTransactionCountByHash() {
	eth := suite.eth
	transactionCount, err := eth.GetBlockTransactionCountByHash(common.NewHash(common.HexToBytes("0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238")))
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		big.NewInt(0xb),
		transactionCount,
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetBlockTransactionCountByNumber() {
	eth := suite.eth
	transactionCount, err := eth.GetBlockTransactionCountByNumber("latest")
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		big.NewInt(0xa),
		transactionCount,
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetUncleCountByBlockHash() {
	eth := suite.eth
	uncleCount, err := eth.GetUncleCountByBlockHash(common.NewHash(common.HexToBytes("0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238")))
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		big.NewInt(0x1),
		uncleCount,
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetUncleCountByBlockNumber() {
	eth := suite.eth
	uncleCount, err := eth.GetUncleCountByBlockNumber("latest")
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		big.NewInt(0x1),
		uncleCount,
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetCode() {
	eth := suite.eth
	code, err := eth.GetCode(common.NewAddress(common.HexToBytes("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b")), "0x2")
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		common.HexToBytes("0x600160008035811a818181146012578301005b601b6001356025565b8060005260206000f25b600060078202905091905056"),
		code,
		"Should be equal")
}

func (suite *EthTestSuite) Test_Sign() {
	eth := suite.eth
	signedData, err := eth.Sign(common.NewAddress(common.HexToBytes("0xd1ade25ccd3d550a7eb532ac759cac7be09c2719")), []byte("Schoolbus"))
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		common.HexToBytes("0x2ac19db245478a06032e69cdbd2b54e648b78431d0a47bd1fbab18f79f820ba407466e37adbe9e84541cab97ab7d290f4a64a5825c876d22109f3bf813254e8601"),
		signedData,
		"Should be equal")
}

func (suite *EthTestSuite) Test_SendTransaction() {
	eth := suite.eth
	req := &common.TransactionRequest{
		From:     common.NewAddress(common.HexToBytes("0xb60e8dd61c5d32be8058bb8eb970870f07233155")),
		To:       common.NewAddress(common.HexToBytes("0xd46e8dd67c5d32be8058bb8eb970870f07244567")),
		Gas:      big.NewInt(0x76c0),
		GasPrice: big.NewInt(0x9184e72a000),
		Value:    big.NewInt(0x9184e72a),
		Data:     common.HexToBytes("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"),
	}
	tx, err := eth.SendTransaction(req)
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		tx,
		"Should be equal")
}

func (suite *EthTestSuite) Test_SendRawTransaction() {
	eth := suite.eth
	tx, err := eth.SendRawTransaction(common.HexToBytes("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"))
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		tx,
		"Should be equal")
}

func (suite *EthTestSuite) Test_Call() {
	eth := suite.eth
	req := &common.TransactionRequest{
		From:     common.NewAddress(common.HexToBytes("0xb60e8dd61c5d32be8058bb8eb970870f07233155")),
		To:       common.NewAddress(common.HexToBytes("0xd46e8dd67c5d32be8058bb8eb970870f07244567")),
		Gas:      big.NewInt(0x76c0),
		GasPrice: big.NewInt(0x9184e72a000),
		Value:    big.NewInt(0x9184e72a),
		Data:     common.HexToBytes("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"),
	}
	result, err := eth.Call(req, "latest")
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		common.HexToBytes("0x"),
		result,
		"Should be equal")
}

func (suite *EthTestSuite) Test_EstimateGas() {
	eth := suite.eth
	req := &common.TransactionRequest{
		From:     common.NewAddress(common.HexToBytes("0xb60e8dd61c5d32be8058bb8eb970870f07233155")),
		To:       common.NewAddress(common.HexToBytes("0xd46e8dd67c5d32be8058bb8eb970870f07244567")),
		Gas:      big.NewInt(0x76c0),
		GasPrice: big.NewInt(0x9184e72a000),
		Value:    big.NewInt(0x9184e72a),
		Data:     common.HexToBytes("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"),
	}
	gas, err := eth.EstimateGas(req, "latest")
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		big.NewInt(0x5208),
		gas,
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetBlockByHash() {
	eth := suite.eth
	block := &common.Block{
		Number:          big.NewInt(0x1b4),
		Hash:            common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		ParentHash:      common.NewHash(common.HexToBytes("0x9646252be9520f6e71339a8df9c55e4d7619deeb018d2a3f2d21fc165dde5eb5")),
		Nonce:           common.NewHash(common.HexToBytes("0xe04d296d2460cfb8472af2c5fd05b5a214109c25688d3704aed5484f9a7792f2")),
		Sha3Uncles:      common.NewHash(common.HexToBytes("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")),
		Bloom:           common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		TransactionRoot: common.NewHash(common.HexToBytes("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")),
		StateRoot:       common.NewHash(common.HexToBytes("0xd5855eb08b3387c0af375e9cdb6acfc05eb8f519e419b874b6ff2ffda7ed1dff")),
		Miner:           common.NewAddress(common.HexToBytes("0x4e65fda2159562a496f9f3522f89122a3088497a")),
		Difficulty:      big.NewInt(0x027f07),
		TotalDifficulty: big.NewInt(0x027f07),
		ExtraData:       common.NewHash(common.HexToBytes("0x0000000000000000000000000000000000000000000000000000000000000000")),
		Size:            big.NewInt(0x027f07),
		GasLimit:        big.NewInt(0x9f759),
		GasUsed:         big.NewInt(0x9f759),
		Timestamp:       big.NewInt(0x54e34e8e),
		Transactions:    []common.Hash{},
		Uncles:          []common.Hash{},
	}
	returnedBlock, err := eth.GetBlockByHash(common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")), true)
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		block, returnedBlock, "Should be equal")
}

func (suite *EthTestSuite) Test_GetBlockByNumber() {
	eth := suite.eth
	block := &common.Block{
		Number:          big.NewInt(0x1b4),
		Hash:            common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		ParentHash:      common.NewHash(common.HexToBytes("0x9646252be9520f6e71339a8df9c55e4d7619deeb018d2a3f2d21fc165dde5eb5")),
		Nonce:           common.NewHash(common.HexToBytes("0xe04d296d2460cfb8472af2c5fd05b5a214109c25688d3704aed5484f9a7792f2")),
		Sha3Uncles:      common.NewHash(common.HexToBytes("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")),
		Bloom:           common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		TransactionRoot: common.NewHash(common.HexToBytes("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")),
		StateRoot:       common.NewHash(common.HexToBytes("0xd5855eb08b3387c0af375e9cdb6acfc05eb8f519e419b874b6ff2ffda7ed1dff")),
		Miner:           common.NewAddress(common.HexToBytes("0x4e65fda2159562a496f9f3522f89122a3088497a")),
		Difficulty:      big.NewInt(0x027f07),
		TotalDifficulty: big.NewInt(0x027f07),
		ExtraData:       common.NewHash(common.HexToBytes("0x0000000000000000000000000000000000000000000000000000000000000000")),
		Size:            big.NewInt(0x027f07),
		GasLimit:        big.NewInt(0x9f759),
		GasUsed:         big.NewInt(0x9f759),
		Timestamp:       big.NewInt(0x54e34e8e),
		Transactions:    []common.Hash{},
		Uncles:          []common.Hash{},
	}
	returnedBlock, err := eth.GetBlockByNumber("0x1b4", true)
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		block, returnedBlock, "Should be equal")
}

func (suite *EthTestSuite) Test_GetTransactionByHash() {
	eth := suite.eth
	tx := &common.Transaction{
		Hash:             common.NewHash(common.HexToBytes("0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b")),
		Nonce:            common.NewHash(common.HexToBytes("0x")),
		BlockHash:        common.NewHash(common.HexToBytes("0xbeab0aa2411b7ab17f30a99d3cb9c6ef2fc5426d6ad6fd9e2a26a6aed1d1055b")),
		BlockNumber:      big.NewInt(0x15df),
		TransactionIndex: 0x1,
		From:             common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")),
		To:               common.NewAddress(common.HexToBytes("0x85h43d8a49eeb85d32cf465507dd71d507100c1")),
		Value:            big.NewInt(0x7f110),
		Gas:              big.NewInt(0x7f110),
		GasPrice:         big.NewInt(0x09184e72a000),
		Data:             common.HexToBytes("0x603880600c6000396000f300603880600c6000396000f3603880600c6000396000f360"),
	}
	returnedTx, err := eth.GetTransactionByHash(common.NewHash(common.HexToBytes("0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238")))
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		tx, returnedTx, "Should be equal")
}

func (suite *EthTestSuite) Test_GetTransactionByHashAndIndex() {
	eth := suite.eth
	tx := &common.Transaction{
		Hash:             common.NewHash(common.HexToBytes("0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b")),
		Nonce:            common.NewHash(common.HexToBytes("0x")),
		BlockHash:        common.NewHash(common.HexToBytes("0xbeab0aa2411b7ab17f30a99d3cb9c6ef2fc5426d6ad6fd9e2a26a6aed1d1055b")),
		BlockNumber:      big.NewInt(0x15df),
		TransactionIndex: 0x1,
		From:             common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")),
		To:               common.NewAddress(common.HexToBytes("0x85h43d8a49eeb85d32cf465507dd71d507100c1")),
		Value:            big.NewInt(0x7f110),
		Gas:              big.NewInt(0x7f110),
		GasPrice:         big.NewInt(0x09184e72a000),
		Data:             common.HexToBytes("0x603880600c6000396000f300603880600c6000396000f3603880600c6000396000f360"),
	}
	returnedTx, err := eth.GetTransactionByBlockHashAndIndex(common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")), 0)
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		tx, returnedTx, "Should be equal")
}

func (suite *EthTestSuite) Test_GetTransactionByNumberAndIndex() {
	eth := suite.eth
	tx := &common.Transaction{
		Hash:             common.NewHash(common.HexToBytes("0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b")),
		Nonce:            common.NewHash(common.HexToBytes("0x")),
		BlockHash:        common.NewHash(common.HexToBytes("0xbeab0aa2411b7ab17f30a99d3cb9c6ef2fc5426d6ad6fd9e2a26a6aed1d1055b")),
		BlockNumber:      big.NewInt(0x15df),
		TransactionIndex: 0x1,
		From:             common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")),
		To:               common.NewAddress(common.HexToBytes("0x85h43d8a49eeb85d32cf465507dd71d507100c1")),
		Value:            big.NewInt(0x7f110),
		Gas:              big.NewInt(0x7f110),
		GasPrice:         big.NewInt(0x09184e72a000),
		Data:             common.HexToBytes("0x603880600c6000396000f300603880600c6000396000f3603880600c6000396000f360"),
	}
	returnedTx, err := eth.GetTransactionByBlockNumberAndIndex("0x29c", 0)
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		tx, returnedTx, "Should be equal")
}

func (suite *EthTestSuite) Test_GetTransactionReceipt() {
	eth := suite.eth
	receipt := &common.TransactionReceipt{
		Hash:              common.NewHash(common.HexToBytes("0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238")),
		TransactionIndex:  0x1,
		BlockNumber:       big.NewInt(0xb),
		BlockHash:         common.NewHash(common.HexToBytes("0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b")),
		CumulativeGasUsed: big.NewInt(0x33bc),
		GasUsed:           big.NewInt(0x4dc),
		ContractAddress:   common.NewAddress(common.HexToBytes("0xb60e8dd61c5d32be8058bb8eb970870f07233155")),
		Logs:              []common.Log{},
	}
	returnReceipt, err := eth.GetTransactionReceipt(common.NewHash(common.HexToBytes("0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238")))
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		receipt, returnReceipt, "Should be equal")
}

func (suite *EthTestSuite) Test_GetUncleByBlockHashAndIndex() {
	eth := suite.eth
	block := &common.Block{
		Number:          big.NewInt(0x1b4),
		Hash:            common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		ParentHash:      common.NewHash(common.HexToBytes("0x9646252be9520f6e71339a8df9c55e4d7619deeb018d2a3f2d21fc165dde5eb5")),
		Nonce:           common.NewHash(common.HexToBytes("0xe04d296d2460cfb8472af2c5fd05b5a214109c25688d3704aed5484f9a7792f2")),
		Sha3Uncles:      common.NewHash(common.HexToBytes("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")),
		Bloom:           common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		TransactionRoot: common.NewHash(common.HexToBytes("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")),
		StateRoot:       common.NewHash(common.HexToBytes("0xd5855eb08b3387c0af375e9cdb6acfc05eb8f519e419b874b6ff2ffda7ed1dff")),
		Miner:           common.NewAddress(common.HexToBytes("0x4e65fda2159562a496f9f3522f89122a3088497a")),
		Difficulty:      big.NewInt(0x027f07),
		TotalDifficulty: big.NewInt(0x027f07),
		ExtraData:       common.NewHash(common.HexToBytes("0x0000000000000000000000000000000000000000000000000000000000000000")),
		Size:            big.NewInt(0x027f07),
		GasLimit:        big.NewInt(0x9f759),
		GasUsed:         big.NewInt(0x9f759),
		Timestamp:       big.NewInt(0x54e34e8e),
		Transactions:    []common.Hash{},
		Uncles:          []common.Hash{},
	}
	returnedBlock, err := eth.GetUncleByBlockHashAndIndex(common.NewHash(common.HexToBytes("0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b")), 0)
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		block, returnedBlock, "Should be equal")
}

func (suite *EthTestSuite) Test_GetUncleByBlockNumberAndIndex() {
	eth := suite.eth
	block := &common.Block{
		Number:          big.NewInt(0x1b4),
		Hash:            common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		ParentHash:      common.NewHash(common.HexToBytes("0x9646252be9520f6e71339a8df9c55e4d7619deeb018d2a3f2d21fc165dde5eb5")),
		Nonce:           common.NewHash(common.HexToBytes("0xe04d296d2460cfb8472af2c5fd05b5a214109c25688d3704aed5484f9a7792f2")),
		Sha3Uncles:      common.NewHash(common.HexToBytes("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")),
		Bloom:           common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		TransactionRoot: common.NewHash(common.HexToBytes("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")),
		StateRoot:       common.NewHash(common.HexToBytes("0xd5855eb08b3387c0af375e9cdb6acfc05eb8f519e419b874b6ff2ffda7ed1dff")),
		Miner:           common.NewAddress(common.HexToBytes("0x4e65fda2159562a496f9f3522f89122a3088497a")),
		Difficulty:      big.NewInt(0x027f07),
		TotalDifficulty: big.NewInt(0x027f07),
		ExtraData:       common.NewHash(common.HexToBytes("0x0000000000000000000000000000000000000000000000000000000000000000")),
		Size:            big.NewInt(0x027f07),
		GasLimit:        big.NewInt(0x9f759),
		GasUsed:         big.NewInt(0x9f759),
		Timestamp:       big.NewInt(0x54e34e8e),
		Transactions:    []common.Hash{},
		Uncles:          []common.Hash{},
	}
	returnedBlock, err := eth.GetUncleByBlockNumberAndIndex("0x29c", 0)
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		block, returnedBlock, "Should be equal")
}

func (suite *EthTestSuite) Test_GetCompilers() {
	eth := suite.eth
	compilers := []string{"solidity", "lll", "serpent"}
	returnedCompilers, err := eth.GetCompilers()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(),
		compilers, returnedCompilers, "Should be equal")
}

func (suite *EthTestSuite) Test_NewFilter() {
	eth := suite.eth
	option := &FilterOption{}
	filter, err := eth.NewFilter(option)
	assert.NoError(suite.T(), err, "Should be no error")
	if assert.NotNil(suite.T(), filter, "Should be equal") {
		assert.EqualValues(suite.T(),
			1, filter.ID(), "Should be equal")
	}
}

func (suite *EthTestSuite) Test_NewBlockFilter() {
	eth := suite.eth
	filter, err := eth.NewBlockFilter()
	assert.NoError(suite.T(), err, "Should be no error")
	if assert.NotNil(suite.T(), filter, "Should be equal") {
		assert.EqualValues(suite.T(),
			1, filter.ID(), "Should be equal")
	}
}

func (suite *EthTestSuite) Test_NewPendingTransactionFilter() {
	eth := suite.eth
	filter, err := eth.NewPendingTransactionFilter()
	assert.NoError(suite.T(), err, "Should be no error")
	if assert.NotNil(suite.T(), filter, "Should be equal") {
		assert.EqualValues(suite.T(),
			1, filter.ID(), "Should be equal")
	}
}

func (suite *EthTestSuite) Test_UninstallFilter() {
	eth := suite.eth
	option := &FilterOption{}
	filter, err := eth.NewFilter(option)
	ok, err := eth.UninstallFilter(filter)
	assert.NoError(suite.T(), err, "Should be no error")
	assert.True(suite.T(), ok, "Should be true")
}

func (suite *EthTestSuite) Test_GetFilterChanges() {
	eth := suite.eth
	option := &FilterOption{}
	filter, err := eth.NewFilter(option)
	logs := []common.Log{
		{
			LogIndex:         0x1,
			BlockNumber:      big.NewInt(0x1b4),
			BlockHash:        common.NewHash(common.HexToBytes("0x8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
			TransactionHash:  common.NewHash(common.HexToBytes("0xdf829c5a142f1fccd7d8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcf")),
			TransactionIndex: 0,
			Address:          common.NewAddress(common.HexToBytes("0x16c5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
			Data:             []byte("0000000000000000000000000000000000000000000000000000000000000000"),
			Topics: common.Topics{
				{
					Data: common.HexToBytes("0x59ebeb90bc63057b6515673c3ecf9438e5058bca0f92585014eced636878c9a5"),
				},
			},
		},
	}
	returnedLogs, err := eth.GetFilterChanges(filter)
	if assert.NoError(suite.T(), err, "Should be no error") {
		for i, l := range returnedLogs {
			log := common.Log{}
			rawBytes, err := json.Marshal(l)
			assert.NoError(suite.T(), err, "Should be no error")
			err = json.Unmarshal(rawBytes, &log)
			assert.NoError(suite.T(), err, "Should be no error")
			assert.EqualValues(suite.T(), logs[i], log, "Should be equal")
		}
	}
}

func (suite *EthTestSuite) Test_GetFilterLogs() {
	eth := suite.eth
	option := &FilterOption{}
	filter, err := eth.NewFilter(option)
	logs := []common.Log{
		{
			LogIndex:         0x1,
			BlockNumber:      big.NewInt(0x1b4),
			BlockHash:        common.NewHash(common.HexToBytes("0x8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
			TransactionHash:  common.NewHash(common.HexToBytes("0xdf829c5a142f1fccd7d8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcf")),
			TransactionIndex: 0,
			Address:          common.NewAddress(common.HexToBytes("0x16c5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
			Data:             []byte("0000000000000000000000000000000000000000000000000000000000000000"),
			Topics: common.Topics{
				{
					Data: common.HexToBytes("0x59ebeb90bc63057b6515673c3ecf9438e5058bca0f92585014eced636878c9a5"),
				},
			},
		},
	}
	returnedLogs, err := eth.GetFilterLogs(filter)
	if assert.NoError(suite.T(), err, "Should be no error") {
		for i, l := range returnedLogs {
			log := common.Log{}
			rawBytes, err := json.Marshal(l)
			assert.NoError(suite.T(), err, "Should be no error")
			err = json.Unmarshal(rawBytes, &log)
			assert.NoError(suite.T(), err, "Should be no error")
			assert.EqualValues(suite.T(), logs[i], log, "Should be equal")
		}
	}
}

func (suite *EthTestSuite) Test_GetLogs() {
	eth := suite.eth
	option := &FilterOption{}
	filter, err := eth.NewFilter(option)
	logs := []common.Log{
		{
			LogIndex:         0x1,
			BlockNumber:      big.NewInt(0x1b4),
			BlockHash:        common.NewHash(common.HexToBytes("0x8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
			TransactionHash:  common.NewHash(common.HexToBytes("0xdf829c5a142f1fccd7d8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcf")),
			TransactionIndex: 0,
			Address:          common.NewAddress(common.HexToBytes("0x16c5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
			Data:             []byte("0000000000000000000000000000000000000000000000000000000000000000"),
			Topics: common.Topics{
				{
					Data: common.HexToBytes("0x59ebeb90bc63057b6515673c3ecf9438e5058bca0f92585014eced636878c9a5"),
				},
			},
		},
	}
	returnedLogs, err := eth.GetLogs(filter)
	if assert.NoError(suite.T(), err, "Should be no error") {
		for i, l := range returnedLogs {
			log := common.Log{}
			rawBytes, err := json.Marshal(l)
			assert.NoError(suite.T(), err, "Should be no error")
			err = json.Unmarshal(rawBytes, &log)
			assert.NoError(suite.T(), err, "Should be no error")
			assert.EqualValues(suite.T(), logs[i], log, "Should be equal")
		}
	}
}

func (suite *EthTestSuite) Test_GetWork() {
	eth := suite.eth
	works := []string{
		"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		"0x5EED00000000000000000000000000005EED0000000000000000000000000000",
		"0xd1ff1c01710000000000000000000000d1ff1c01710000000000000000000000"}
	header, seed, boundary, err := eth.GetWork()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(), strings.ToLower(works[0]), header.String(), "Should be equal")
	assert.EqualValues(suite.T(), strings.ToLower(works[1]), seed.String(), "Should be equal")
	assert.EqualValues(suite.T(), strings.ToLower(works[2]), boundary.String(), "Should be equal")
}

func (suite *EthTestSuite) Test_SubmitWork() {
	eth := suite.eth
	header := common.NewHash(common.HexToBytes("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"))
	mixDigest := common.NewHash(common.HexToBytes("0xD1FE5700000000000000000000000000D1FE5700000000000000000000000000"))
	result, err := eth.SubmitWork(0, header, mixDigest)
	assert.NoError(suite.T(), err, "Should be no error")
	assert.True(suite.T(), result, "Should be true")
}

func (suite *EthTestSuite) SetupTest() {
	suite.web3 = NewWeb3(test.NewMockHTTPProvider())
	suite.eth = suite.web3.Eth
}

func Test_EthTestSuite(t *testing.T) {
	suite.Run(t, new(EthTestSuite))
}
