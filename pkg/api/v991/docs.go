// Code generated by tools/generator. DO NOT EDIT.

//go:build testonly

/*
Copyright 2024 The CDEvents Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

// Package v03 contains method to create events that belong to the
// CDEvents specification v99.1.x

package v991

import "github.com/cdevents/sdk-go/pkg/api"

var SpecVersion = "99.1.0"

type FooSubjectBarPredicateEvent = api.FooSubjectBarPredicateEventV2_2_3
type FooSubjectBarPredicateSubject = api.FooSubjectBarPredicateSubjectV2_2_3

// FooSubjectFooSubjectSubjectContentObjectFieldV2_2_3 holds the content of a ObjectField field in the content
type FooSubjectBarPredicateSubjectContentObjectField = api.FooSubjectBarPredicateSubjectContentObjectFieldV2_2_3

func NewFooSubjectBarPredicateEvent() (*FooSubjectBarPredicateEvent, error) {
	return api.NewFooSubjectBarPredicateEventV2_2_3(SpecVersion)
}

var FooSubjectBarPredicateEventType = api.FooSubjectBarPredicateEventTypeV2_2_3

// NewFromJsonBytes builds a new CDEventReader from a JSON string as []bytes
// This works by unmarshalling the context first, extracting the event type and using
// that to unmarshal the rest of the event into the correct object.
// It assumes the context can be unmarshalled in a `Context` object.
func NewFromJsonBytes(event []byte) (api.CDEventV04, error) {
	return api.NewFromJsonBytesContext[api.CDEventV04](event, CDEventsByUnversionedTypes)
}

// Build a new CDEventReader from a JSON string
func NewFromJsonString(event string) (api.CDEventV04, error) {
	return NewFromJsonBytes([]byte(event))
}