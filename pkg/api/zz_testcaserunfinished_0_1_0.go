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
	// TestCaseRunFinished event type v0.1.0
	TestCaseRunFinishedEventTypeV0_1_0 CDEventType = CDEventType{
		Subject:   "testcaserun",
		Predicate: "finished",
		Version:   "0.1.0",
	}
)

type TestCaseRunFinishedSubjectContentV0_1_0 struct {
	Environment *Reference `json:"environment"`

	Outcome string `json:"outcome"`

	Reason string `json:"reason,omitempty"`

	Severity string `json:"severity,omitempty"`

	TestCase *TestCaseRunFinishedSubjectContentTestCaseV0_1_0 `json:"testCase,omitempty"`

	TestSuiteRun *Reference `json:"testSuiteRun,omitempty"`
}

type TestCaseRunFinishedSubjectV0_1_0 struct {
	SubjectBase
	Content TestCaseRunFinishedSubjectContentV0_1_0 `json:"content"`
}

func (sc TestCaseRunFinishedSubjectV0_1_0) GetSubjectType() SubjectType {
	return "testCaseRun"
}

type TestCaseRunFinishedEventV0_1_0 struct {
	Context Context                          `json:"context"`
	Subject TestCaseRunFinishedSubjectV0_1_0 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e TestCaseRunFinishedEventV0_1_0) GetType() CDEventType {
	return TestCaseRunFinishedEventTypeV0_1_0
}

func (e TestCaseRunFinishedEventV0_1_0) GetVersion() string {
	return e.Context.GetVersion()
}

func (e TestCaseRunFinishedEventV0_1_0) GetId() string {
	return e.Context.Id
}

func (e TestCaseRunFinishedEventV0_1_0) GetSource() string {
	return e.Context.Source
}

func (e TestCaseRunFinishedEventV0_1_0) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TestCaseRunFinishedEventV0_1_0) GetSubjectId() string {
	return e.Subject.Id
}

func (e TestCaseRunFinishedEventV0_1_0) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TestCaseRunFinishedEventV0_1_0) GetSubject() Subject {
	return e.Subject
}

func (e TestCaseRunFinishedEventV0_1_0) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e TestCaseRunFinishedEventV0_1_0) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e TestCaseRunFinishedEventV0_1_0) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e TestCaseRunFinishedEventV0_1_0) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *TestCaseRunFinishedEventV0_1_0) SetId(id string) {
	e.Context.Id = id
}

func (e *TestCaseRunFinishedEventV0_1_0) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TestCaseRunFinishedEventV0_1_0) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TestCaseRunFinishedEventV0_1_0) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TestCaseRunFinishedEventV0_1_0) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *TestCaseRunFinishedEventV0_1_0) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e TestCaseRunFinishedEventV0_1_0) GetSchema() (string, *jsonschema.Schema, error) {
	eType := e.GetType()
	return CompiledSchemas.GetBySpecSubjectPredicate("0.3.0", eType.Subject, eType.Predicate, eType.Custom)
}

func (e TestCaseRunFinishedEventV0_1_0) GetSubjectContent() interface{} {
	return e.Subject.Content
}

// Set subject custom fields

func (e *TestCaseRunFinishedEventV0_1_0) SetSubjectEnvironment(environment *Reference) {
	e.Subject.Content.Environment = environment
}

func (e *TestCaseRunFinishedEventV0_1_0) SetSubjectOutcome(outcome string) {
	e.Subject.Content.Outcome = outcome
}

func (e *TestCaseRunFinishedEventV0_1_0) SetSubjectReason(reason string) {
	e.Subject.Content.Reason = reason
}

func (e *TestCaseRunFinishedEventV0_1_0) SetSubjectSeverity(severity string) {
	e.Subject.Content.Severity = severity
}

func (e *TestCaseRunFinishedEventV0_1_0) SetSubjectTestCase(testCase *TestCaseRunFinishedSubjectContentTestCaseV0_1_0) {
	e.Subject.Content.TestCase = testCase
}

func (e *TestCaseRunFinishedEventV0_1_0) SetSubjectTestSuiteRun(testSuiteRun *Reference) {
	e.Subject.Content.TestSuiteRun = testSuiteRun
}

// New creates a new TestCaseRunFinishedEventV0_1_0
func NewTestCaseRunFinishedEventV0_1_0(specVersion string) (*TestCaseRunFinishedEventV0_1_0, error) {
	e := &TestCaseRunFinishedEventV0_1_0{
		Context: Context{
			Type:    TestCaseRunFinishedEventTypeV0_1_0,
			Version: specVersion,
		},
		Subject: TestCaseRunFinishedSubjectV0_1_0{
			SubjectBase: SubjectBase{
				Type: "testCaseRun",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// TestCaseRunFinishedSubjectContentTestCaseV0_1_0 holds the content of a TestCase field in the content
type TestCaseRunFinishedSubjectContentTestCaseV0_1_0 struct {
	Id string `json:"id"`

	Name string `json:"name,omitempty"`

	Type string `json:"type,omitempty"`

	Uri string `json:"uri,omitempty"`

	Version string `json:"version,omitempty"`
}
