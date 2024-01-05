// Copyright Istio Authors
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

package echo

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jacob-delgado/integ-testing-framework/framework/components/namespace"
)

func TestParseConfigs(t *testing.T) {
	cfgs, err := ParseConfigs([]byte(`
- Service: "foo"
  Namespace: "bar"
  DeployAsVM: true
  VMDistro: "Centos8"
`))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(cfgs, []Config{{
		Service:    "foo",
		Namespace:  namespace.Static("bar"),
		DeployAsVM: true,
		VMDistro:   "Centos8",
	}}); diff != "" {
		t.Fatal(diff)
	}
}
