// Copyright 2018 The containerd Authors.
// Copyright 2019 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cli defines the command line interface for the V2 shim.
package cli

import (
	containerdshim "github.com/containerd/containerd/runtime/v2/shim"

	shim "gvisor.dev/gvisor/pkg/shim/v1"
)

// Main is the main entrypoint.
func Main() {
	containerdshim.Run("io.containerd.runsc.v1", shim.New)
}
