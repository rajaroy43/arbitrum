/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package state

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type State interface {
	SendMessageToVM(msg protocol.Message)
	GetCore() *core.Core
	GetConfig() *core.Config
}

type ChannelState interface {
	State
	ChannelUpdateTime(uint64, bridge.Bridge) (ChannelState, error)
	ChannelUpdateState(ethbridge.Event, uint64, bridge.Bridge) (ChannelState, challenge.State, error)
}

type ChainState interface {
	ChannelState
	ChainUpdateTime(uint64, bridge.ArbVMBridge) (ChainState, error)
	ChainUpdateState(ethbridge.Event, uint64, bridge.ArbVMBridge) (ChainState, challenge.State, error)
}