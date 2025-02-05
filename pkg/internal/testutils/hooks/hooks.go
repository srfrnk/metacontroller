/*
Copyright 2021 Metacontroller authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hooks

import (
	"fmt"
	"metacontroller/pkg/apis/metacontroller/v1alpha1"
	"metacontroller/pkg/hooks"
	"reflect"
)

// NewHookExecutorStub creates new HookExecutorStub which returns given response
func NewHookExecutorStub(response interface{}) hooks.HookExecutor {
	return &hookExecutorStub{
		enabled:  true,
		response: response,
	}
}

// HookExecutorStub is HookExecutor stub to return any given response
type hookExecutorStub struct {
	enabled  bool
	response interface{}
}

func (h *hookExecutorStub) IsEnabled() bool {
	return true
}

func (h *hookExecutorStub) Execute(request interface{}, response interface{}) error {
	val := reflect.ValueOf(response)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf(`panic("not a pointer")`)
	}

	val = val.Elem()

	newVal := reflect.Indirect(reflect.ValueOf(h.response))

	if !val.Type().AssignableTo(newVal.Type()) {
		return fmt.Errorf(`panic("mismatched types")`)
	}

	val.Set(newVal)
	return nil
}

func (h hookExecutorStub) Close() {}

type NilCustomizableController struct {
}

func (cc *NilCustomizableController) GetCustomizeHook() *v1alpha1.Hook {
	return nil
}

type FakeCustomizableController struct {
}

func (cc *FakeCustomizableController) GetCustomizeHook() *v1alpha1.Hook {
	url := "fake"
	return &v1alpha1.Hook{
		Webhook: &v1alpha1.Webhook{
			URL: &url,
		},
	}
}
