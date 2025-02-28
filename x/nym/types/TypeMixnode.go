// Copyright 2020 Nym Technologies SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Mixnode models a Nym mixnode, which shuffles Sphinx packets together inside itself to provide network privacy for users.
type Mixnode struct {
	Creator    sdk.AccAddress `json:"creator" yaml:"creator"`
	ID         string         `json:"id" yaml:"id"`
	PubKey     string         `json:"pubKey" yaml:"pubKey"`
	Layer      int32          `json:"layer" yaml:"layer"`
	Version    string         `json:"version" yaml:"version"`
	Host       string         `json:"host" yaml:"host"`
	Location   string         `json:"location" yaml:"location"`
	Reputation int32          `json:"reputation" yaml:"reputation"`
}
