// Copyright © 2020 Attestant Limited.
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

package tekuhttp_test

import (
	"context"
	"os"
	"testing"

	"github.com/attestantio/go-eth2-client/tekuhttp"
	"github.com/stretchr/testify/require"
)

func TestValidators(t *testing.T) {
	tests := []struct {
		name    string
		stateID string
	}{
		{
			name:    "Good",
			stateID: "genesis",
		},
	}

	service, err := tekuhttp.New(context.Background(),
		tekuhttp.WithAddress(os.Getenv("TEKUHTTP_ADDRESS")),
		tekuhttp.WithTimeout(timeout),
	)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			validators, err := service.Validators(context.Background(), test.stateID, nil)
			require.NoError(t, err)
			require.NotNil(t, validators)
		})
	}
}
