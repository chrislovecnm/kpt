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

package kptlib

import (
	"context"
	"io"

	"github.com/GoogleContainerTools/kpt/internal/printer"
)

type Printer = printer.Printer

func WithPrinterContext(ctx context.Context, pr printer.Printer) context.Context {
	return printer.WithContext(ctx, pr)
}

func NewPrinter(stdout, stderr io.Writer) Printer {
	return printer.New(stdout, stderr)
}