// Copyright 2019 Google LLC
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

package yaml

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var Filters = map[string]func() Filter{
	"AnnotationClearer": func() Filter { return &AnnotationClearer{} },
	"AnnotationGetter":  func() Filter { return &AnnotationGetter{} },
	"AnnotationSetter":  func() Filter { return &AnnotationSetter{} },
	"ElementAppender":   func() Filter { return &ElementAppender{} },
	"ElementMatcher":    func() Filter { return &ElementMatcher{} },
	"FieldClearer":      func() Filter { return &FieldClearer{} },
	"FieldMatcher":      func() Filter { return &FieldMatcher{} },
	"FieldSetter":       func() Filter { return &FieldSetter{} },
	"HasFilter":         func() Filter { return &HasFilter{} },
	"PathGetter":        func() Filter { return &PathGetter{} },
	"Parser":            func() Filter { return &Parser{} },
	"PrefixFilter":      func() Filter { return &PrefixFilter{} },
	"SedFilter":         func() Filter { return &SedFilter{} },
	"SuffixFilter":      func() Filter { return &SuffixFilter{} },
	"TeeFilter":         func() Filter { return &TeeFilter{} },
}

// YFilter wraps the Filter interface so it can be unmarshalled into a struct.
type YFilter struct {
	Filter
}

func (y YFilter) MarshalYAML() (interface{}, error) {
	return y.Filter, nil
}

func (y *YFilter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	meta := &ResourceMeta{}
	if err := unmarshal(meta); err != nil {
		return err
	}
	if filter, found := Filters[meta.Kind]; !found {
		var knownFilters []string
		for k := range Filters {
			knownFilters = append(knownFilters, k)
		}
		sort.Strings(knownFilters)
		return fmt.Errorf("unsupported Filter Kind %s:  may be one of: [%s]",
			meta.Kind, strings.Join(knownFilters, ","))
	} else {
		y.Filter = filter()
	}
	if err := unmarshal(y.Filter); err != nil {
		return err
	}
	return nil
}

type YFilters []YFilter

func (y YFilters) Filters() []Filter {
	var f []Filter
	for i := range y {
		f = append(f, y[i].Filter)
	}
	return f
}

type HasFilter struct {
	Kind string `yaml:"kind"`

	// Filters are the set of Filters run by TeeFilter.
	Filters YFilters `yaml:"pipeline,omitempty"`
}

func (t HasFilter) Filter(rn *RNode) (*RNode, error) {
	v, err := rn.Pipe(t.Filters.Filters()...)
	if v == nil || err != nil {
		return nil, err
	}
	// return the original input if the pipeline resolves to true
	return rn, err
}

type SedFilter struct {
	Kind string `yaml:"kind"`

	StringMatch string `yaml:"stringMatch"`
	RegexMatch  string `yaml:"regexMatch"`
	Replace     string `yaml:"replace"`
	Count       int    `yaml:"count"`
}

func (s SedFilter) Filter(object *RNode) (*RNode, error) {
	if s.Count == 0 {
		s.Count = -1
	}
	if s.StringMatch != "" {
		object.value.Value = strings.Replace(object.value.Value, s.StringMatch, s.Replace, s.Count)
	} else if s.RegexMatch != "" {
		r, err := regexp.Compile(s.RegexMatch)
		if err != nil {
			return nil, fmt.Errorf("SedFilter RegexMatch does not compile: %v", err)
		}
		object.value.Value = r.ReplaceAllString(object.value.Value, s.Replace)
	} else {
		return nil, fmt.Errorf("SedFilter missing StringMatch and RegexMatch")
	}
	return object, nil
}

type PrefixFilter struct {
	Kind string `yaml:"kind"`

	Value string `yaml:"value"`
}

func (s PrefixFilter) Filter(object *RNode) (*RNode, error) {
	if !strings.HasPrefix(object.value.Value, s.Value) {
		object.value.Value = s.Value + object.value.Value
	}
	return object, nil
}

type SuffixFilter struct {
	Kind string `yaml:"kind"`

	Value string `yaml:"value"`
}

func (s SuffixFilter) Filter(object *RNode) (*RNode, error) {
	if !strings.HasSuffix(object.value.Value, s.Value) {
		object.value.Value = object.value.Value + s.Value
	}
	return object, nil
}
