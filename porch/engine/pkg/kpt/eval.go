// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kpt

import (
	"context"
	"fmt"
	"io"

	kptfilev1 "github.com/GoogleContainerTools/kpt/pkg/api/kptfile/v1"
	"github.com/GoogleContainerTools/kpt/pkg/fn"
	"github.com/GoogleContainerTools/kpt/porch/engine/pkg/internal"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/kio"
)

func NewPlaceholderFunctionRuntime() fn.FunctionRuntime {
	return &runtime{}
}

type runtime struct {
}

var _ fn.FunctionRuntime = &runtime{}

func (e *runtime) GetRunner(ctx context.Context, fn *kptfilev1.Function) (fn.FunctionRunner, error) {
	processor := internal.FindProcessor(fn.Image)
	if processor == nil {
		return nil, fmt.Errorf("unsupported kpt function %q", fn.Image)
	}

	return &runner{
		ctx:       ctx,
		fn:        *fn,
		processor: processor,
	}, nil
}

type runner struct {
	ctx       context.Context
	fn        kptfilev1.Function
	processor framework.ResourceListProcessorFunc
}

var _ fn.FunctionRunner = &runner{}

func (fr *runner) Run(r io.Reader, w io.Writer) error {
	rw := &kio.ByteReadWriter{
		Reader:                r,
		Writer:                w,
		KeepReaderAnnotations: true,
	}

	return framework.Execute(fr.processor, rw)
}
