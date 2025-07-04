// Code generated by tools/generator. DO NOT EDIT.

/*
Copyright 2023 The CDEvents Authors

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

package api

import (
	"time"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
)

var (
	// RepositoryModified event type v0.1.1
	RepositoryModifiedEventTypeV0_1_1 CDEventType = CDEventType{
		Subject:   "repository",
		Predicate: "modified",
		Version:   "0.1.1",
	}
)

type RepositoryModifiedSubjectContentV0_1_1 struct {
	Name string `json:"name,omitempty"`

	Owner string `json:"owner,omitempty"`

	Url string `json:"url,omitempty"`

	ViewUrl string `json:"viewUrl,omitempty"`
}

type RepositoryModifiedSubjectV0_1_1 struct {
	SubjectBase
	Content RepositoryModifiedSubjectContentV0_1_1 `json:"content"`
}

func (sc RepositoryModifiedSubjectV0_1_1) GetSubjectType() SubjectType {
	return "repository"
}

type RepositoryModifiedEventV0_1_1 struct {
	Context Context                         `json:"context"`
	Subject RepositoryModifiedSubjectV0_1_1 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e RepositoryModifiedEventV0_1_1) GetType() CDEventType {
	return RepositoryModifiedEventTypeV0_1_1
}

func (e RepositoryModifiedEventV0_1_1) GetVersion() string {
	return e.Context.GetVersion()
}

func (e RepositoryModifiedEventV0_1_1) GetId() string {
	return e.Context.Id
}

func (e RepositoryModifiedEventV0_1_1) GetSource() string {
	return e.Context.Source
}

func (e RepositoryModifiedEventV0_1_1) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e RepositoryModifiedEventV0_1_1) GetSubjectId() string {
	return e.Subject.Id
}

func (e RepositoryModifiedEventV0_1_1) GetSubjectSource() string {
	return e.Subject.Source
}

func (e RepositoryModifiedEventV0_1_1) GetSubject() Subject {
	return e.Subject
}

func (e RepositoryModifiedEventV0_1_1) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryModifiedEventV0_1_1) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e RepositoryModifiedEventV0_1_1) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryModifiedEventV0_1_1) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *RepositoryModifiedEventV0_1_1) SetId(id string) {
	e.Context.Id = id
}

func (e *RepositoryModifiedEventV0_1_1) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *RepositoryModifiedEventV0_1_1) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *RepositoryModifiedEventV0_1_1) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *RepositoryModifiedEventV0_1_1) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *RepositoryModifiedEventV0_1_1) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e RepositoryModifiedEventV0_1_1) GetSchema() (string, *jsonschema.Schema, error) {
	eType := e.GetType()
	return CompiledSchemas.GetBySpecSubjectPredicate("0.3.0", eType.Subject, eType.Predicate, eType.Custom)
}

func (e RepositoryModifiedEventV0_1_1) GetSubjectContent() interface{} {
	return e.Subject.Content
}

// Set subject custom fields

func (e *RepositoryModifiedEventV0_1_1) SetSubjectName(name string) {
	e.Subject.Content.Name = name
}

func (e *RepositoryModifiedEventV0_1_1) SetSubjectOwner(owner string) {
	e.Subject.Content.Owner = owner
}

func (e *RepositoryModifiedEventV0_1_1) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func (e *RepositoryModifiedEventV0_1_1) SetSubjectViewUrl(viewUrl string) {
	e.Subject.Content.ViewUrl = viewUrl
}

// New creates a new RepositoryModifiedEventV0_1_1
func NewRepositoryModifiedEventV0_1_1(specVersion string) (*RepositoryModifiedEventV0_1_1, error) {
	e := &RepositoryModifiedEventV0_1_1{
		Context: Context{
			Type:    RepositoryModifiedEventTypeV0_1_1,
			Version: specVersion,
		},
		Subject: RepositoryModifiedSubjectV0_1_1{
			SubjectBase: SubjectBase{
				Type: "repository",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
