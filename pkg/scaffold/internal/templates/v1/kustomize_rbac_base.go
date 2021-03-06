/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/pkg/model/file"
)

var _ file.Template = &KustomizeRBAC{}

// KustomizeRBAC scaffolds the Kustomization file in rbac folder.
type KustomizeRBAC struct {
	file.Input
}

// GetInput implements input.Template
func (f *KustomizeRBAC) GetInput() (file.Input, error) {
	if f.Path == "" {
		f.Path = filepath.Join("config", "rbac", "kustomization.yaml")
	}
	f.TemplateBody = kustomizeRBACTemplate
	f.Input.IfExistsAction = file.Error
	return f.Input, nil
}

const kustomizeRBACTemplate = `resources:
- rbac_role.yaml
- rbac_role_binding.yaml
  # Comment the following 3 lines if you want to disable
  # the auth proxy (https://github.com/brancz/kube-rbac-proxy)
  # which protects your /metrics endpoint.
- auth_proxy_service.yaml
- auth_proxy_role.yaml
- auth_proxy_role_binding.yaml
`
