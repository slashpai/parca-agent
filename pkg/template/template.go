// Copyright 2021 The Parca Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package template

import (
	_ "embed"
	"html/template"
	"time"

	"github.com/prometheus/prometheus/pkg/labels"
)

//go:embed statuspage.html
var StatusPageTemplateBytes []byte

var StatusPageTemplate = template.Must(template.New("statuspage").Parse(string(StatusPageTemplateBytes)))

//go:embed profileview.html
var ProfileViewTemplateBytes []byte

var ProfileViewTemplate = template.Must(template.New("profileview").Parse(string(ProfileViewTemplateBytes)))

type ActiveProfiler struct {
	Type         string
	Labels       labels.Labels
	LastTakenAgo time.Duration
	Error        error
	Link         string
}

type StatusPage struct {
	ActiveProfilers []ActiveProfiler
}
