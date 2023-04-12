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

func init() {
	// Create events equal to examples in the spec
	examplesProduced = make(map[string]CDEvent)
	// ArtifactPackaged Event producer
	newArtifactPackaged, _ := NewArtifactPackagedEvent()
	setContext(newArtifactPackaged, testSubjectId)
	exampleArtifactPackagedEvent(newArtifactPackaged)
	examplesProduced[newArtifactPackaged.GetType().Short()] = newArtifactPackaged

	// ArtifactPublished Event producer
	newArtifactPublished, _ := NewArtifactPublishedEvent()
	setContext(newArtifactPublished, testSubjectId)
	exampleArtifactPublishedEvent(newArtifactPublished)
	examplesProduced[newArtifactPublished.GetType().Short()] = newArtifactPublished

	// ArtifactSigned Event producer
	newArtifactSigned, _ := NewArtifactSignedEvent()
	setContext(newArtifactSigned, testSubjectId)
	exampleArtifactSignedEvent(newArtifactSigned)
	examplesProduced[newArtifactSigned.GetType().Short()] = newArtifactSigned

	// BranchCreated Event producer
	newBranchCreated, _ := NewBranchCreatedEvent()
	setContext(newBranchCreated, testSubjectId)
	exampleBranchCreatedEvent(newBranchCreated)
	examplesProduced[newBranchCreated.GetType().Short()] = newBranchCreated

	// BranchDeleted Event producer
	newBranchDeleted, _ := NewBranchDeletedEvent()
	setContext(newBranchDeleted, testSubjectId)
	exampleBranchDeletedEvent(newBranchDeleted)
	examplesProduced[newBranchDeleted.GetType().Short()] = newBranchDeleted

	// BuildFinished Event producer
	newBuildFinished, _ := NewBuildFinishedEvent()
	setContext(newBuildFinished, testSubjectId)
	exampleBuildFinishedEvent(newBuildFinished)
	examplesProduced[newBuildFinished.GetType().Short()] = newBuildFinished

	// BuildQueued Event producer
	newBuildQueued, _ := NewBuildQueuedEvent()
	setContext(newBuildQueued, testSubjectId)
	exampleBuildQueuedEvent(newBuildQueued)
	examplesProduced[newBuildQueued.GetType().Short()] = newBuildQueued

	// BuildStarted Event producer
	newBuildStarted, _ := NewBuildStartedEvent()
	setContext(newBuildStarted, testSubjectId)
	exampleBuildStartedEvent(newBuildStarted)
	examplesProduced[newBuildStarted.GetType().Short()] = newBuildStarted

	// ChangeAbandoned Event producer
	newChangeAbandoned, _ := NewChangeAbandonedEvent()
	setContext(newChangeAbandoned, testSubjectId)
	exampleChangeAbandonedEvent(newChangeAbandoned)
	examplesProduced[newChangeAbandoned.GetType().Short()] = newChangeAbandoned

	// ChangeCreated Event producer
	newChangeCreated, _ := NewChangeCreatedEvent()
	setContext(newChangeCreated, testSubjectId)
	exampleChangeCreatedEvent(newChangeCreated)
	examplesProduced[newChangeCreated.GetType().Short()] = newChangeCreated

	// ChangeMerged Event producer
	newChangeMerged, _ := NewChangeMergedEvent()
	setContext(newChangeMerged, testSubjectId)
	exampleChangeMergedEvent(newChangeMerged)
	examplesProduced[newChangeMerged.GetType().Short()] = newChangeMerged

	// ChangeReviewed Event producer
	newChangeReviewed, _ := NewChangeReviewedEvent()
	setContext(newChangeReviewed, testSubjectId)
	exampleChangeReviewedEvent(newChangeReviewed)
	examplesProduced[newChangeReviewed.GetType().Short()] = newChangeReviewed

	// ChangeUpdated Event producer
	newChangeUpdated, _ := NewChangeUpdatedEvent()
	setContext(newChangeUpdated, testSubjectId)
	exampleChangeUpdatedEvent(newChangeUpdated)
	examplesProduced[newChangeUpdated.GetType().Short()] = newChangeUpdated

	// EnvironmentCreated Event producer
	newEnvironmentCreated, _ := NewEnvironmentCreatedEvent()
	setContext(newEnvironmentCreated, testSubjectId)
	exampleEnvironmentCreatedEvent(newEnvironmentCreated)
	examplesProduced[newEnvironmentCreated.GetType().Short()] = newEnvironmentCreated

	// EnvironmentDeleted Event producer
	newEnvironmentDeleted, _ := NewEnvironmentDeletedEvent()
	setContext(newEnvironmentDeleted, testSubjectId)
	exampleEnvironmentDeletedEvent(newEnvironmentDeleted)
	examplesProduced[newEnvironmentDeleted.GetType().Short()] = newEnvironmentDeleted

	// EnvironmentModified Event producer
	newEnvironmentModified, _ := NewEnvironmentModifiedEvent()
	setContext(newEnvironmentModified, testSubjectId)
	exampleEnvironmentModifiedEvent(newEnvironmentModified)
	examplesProduced[newEnvironmentModified.GetType().Short()] = newEnvironmentModified

	// IncidentDetected Event producer
	newIncidentDetected, _ := NewIncidentDetectedEvent()
	setContext(newIncidentDetected, testSubjectId)
	exampleIncidentDetectedEvent(newIncidentDetected)
	examplesProduced[newIncidentDetected.GetType().Short()] = newIncidentDetected

	// IncidentReported Event producer
	newIncidentReported, _ := NewIncidentReportedEvent()
	setContext(newIncidentReported, testSubjectId)
	exampleIncidentReportedEvent(newIncidentReported)
	examplesProduced[newIncidentReported.GetType().Short()] = newIncidentReported

	// IncidentResolved Event producer
	newIncidentResolved, _ := NewIncidentResolvedEvent()
	setContext(newIncidentResolved, testSubjectId)
	exampleIncidentResolvedEvent(newIncidentResolved)
	examplesProduced[newIncidentResolved.GetType().Short()] = newIncidentResolved

	// PipelineRunFinished Event producer
	newPipelineRunFinished, _ := NewPipelineRunFinishedEvent()
	setContext(newPipelineRunFinished, testSubjectId)
	examplePipelineRunFinishedEvent(newPipelineRunFinished)
	examplesProduced[newPipelineRunFinished.GetType().Short()] = newPipelineRunFinished

	// PipelineRunQueued Event producer
	newPipelineRunQueued, _ := NewPipelineRunQueuedEvent()
	setContext(newPipelineRunQueued, testSubjectId)
	examplePipelineRunQueuedEvent(newPipelineRunQueued)
	examplesProduced[newPipelineRunQueued.GetType().Short()] = newPipelineRunQueued

	// PipelineRunStarted Event producer
	newPipelineRunStarted, _ := NewPipelineRunStartedEvent()
	setContext(newPipelineRunStarted, testSubjectId)
	examplePipelineRunStartedEvent(newPipelineRunStarted)
	examplesProduced[newPipelineRunStarted.GetType().Short()] = newPipelineRunStarted

	// RepositoryCreated Event producer
	newRepositoryCreated, _ := NewRepositoryCreatedEvent()
	setContext(newRepositoryCreated, testSubjectId)
	exampleRepositoryCreatedEvent(newRepositoryCreated)
	examplesProduced[newRepositoryCreated.GetType().Short()] = newRepositoryCreated

	// RepositoryDeleted Event producer
	newRepositoryDeleted, _ := NewRepositoryDeletedEvent()
	setContext(newRepositoryDeleted, testSubjectId)
	exampleRepositoryDeletedEvent(newRepositoryDeleted)
	examplesProduced[newRepositoryDeleted.GetType().Short()] = newRepositoryDeleted

	// RepositoryModified Event producer
	newRepositoryModified, _ := NewRepositoryModifiedEvent()
	setContext(newRepositoryModified, testSubjectId)
	exampleRepositoryModifiedEvent(newRepositoryModified)
	examplesProduced[newRepositoryModified.GetType().Short()] = newRepositoryModified

	// ServiceDeployed Event producer
	newServiceDeployed, _ := NewServiceDeployedEvent()
	setContext(newServiceDeployed, testSubjectId)
	exampleServiceDeployedEvent(newServiceDeployed)
	examplesProduced[newServiceDeployed.GetType().Short()] = newServiceDeployed

	// ServicePublished Event producer
	newServicePublished, _ := NewServicePublishedEvent()
	setContext(newServicePublished, testSubjectId)
	exampleServicePublishedEvent(newServicePublished)
	examplesProduced[newServicePublished.GetType().Short()] = newServicePublished

	// ServiceRemoved Event producer
	newServiceRemoved, _ := NewServiceRemovedEvent()
	setContext(newServiceRemoved, testSubjectId)
	exampleServiceRemovedEvent(newServiceRemoved)
	examplesProduced[newServiceRemoved.GetType().Short()] = newServiceRemoved

	// ServiceRolledback Event producer
	newServiceRolledback, _ := NewServiceRolledbackEvent()
	setContext(newServiceRolledback, testSubjectId)
	exampleServiceRolledbackEvent(newServiceRolledback)
	examplesProduced[newServiceRolledback.GetType().Short()] = newServiceRolledback

	// ServiceUpgraded Event producer
	newServiceUpgraded, _ := NewServiceUpgradedEvent()
	setContext(newServiceUpgraded, testSubjectId)
	exampleServiceUpgradedEvent(newServiceUpgraded)
	examplesProduced[newServiceUpgraded.GetType().Short()] = newServiceUpgraded

	// TaskRunFinished Event producer
	newTaskRunFinished, _ := NewTaskRunFinishedEvent()
	setContext(newTaskRunFinished, testSubjectId)
	exampleTaskRunFinishedEvent(newTaskRunFinished)
	examplesProduced[newTaskRunFinished.GetType().Short()] = newTaskRunFinished

	// TaskRunStarted Event producer
	newTaskRunStarted, _ := NewTaskRunStartedEvent()
	setContext(newTaskRunStarted, testSubjectId)
	exampleTaskRunStartedEvent(newTaskRunStarted)
	examplesProduced[newTaskRunStarted.GetType().Short()] = newTaskRunStarted

	// TestCaseRunFinished Event producer
	newTestCaseRunFinished, _ := NewTestCaseRunFinishedEvent()
	setContext(newTestCaseRunFinished, testSubjectId)
	exampleTestCaseRunFinishedEvent(newTestCaseRunFinished)
	examplesProduced[newTestCaseRunFinished.GetType().Short()] = newTestCaseRunFinished

	// TestCaseRunQueued Event producer
	newTestCaseRunQueued, _ := NewTestCaseRunQueuedEvent()
	setContext(newTestCaseRunQueued, testSubjectId)
	exampleTestCaseRunQueuedEvent(newTestCaseRunQueued)
	examplesProduced[newTestCaseRunQueued.GetType().Short()] = newTestCaseRunQueued

	// TestCaseRunStarted Event producer
	newTestCaseRunStarted, _ := NewTestCaseRunStartedEvent()
	setContext(newTestCaseRunStarted, testSubjectId)
	exampleTestCaseRunStartedEvent(newTestCaseRunStarted)
	examplesProduced[newTestCaseRunStarted.GetType().Short()] = newTestCaseRunStarted

	// TestOutputPublished Event producer
	newTestOutputPublished, _ := NewTestOutputPublishedEvent()
	setContext(newTestOutputPublished, testSubjectId)
	exampleTestOutputPublishedEvent(newTestOutputPublished)
	examplesProduced[newTestOutputPublished.GetType().Short()] = newTestOutputPublished

	// TestSuiteRunFinished Event producer
	newTestSuiteRunFinished, _ := NewTestSuiteRunFinishedEvent()
	setContext(newTestSuiteRunFinished, testSubjectId)
	exampleTestSuiteRunFinishedEvent(newTestSuiteRunFinished)
	examplesProduced[newTestSuiteRunFinished.GetType().Short()] = newTestSuiteRunFinished

	// TestSuiteRunQueued Event producer
	newTestSuiteRunQueued, _ := NewTestSuiteRunQueuedEvent()
	setContext(newTestSuiteRunQueued, testSubjectId)
	exampleTestSuiteRunQueuedEvent(newTestSuiteRunQueued)
	examplesProduced[newTestSuiteRunQueued.GetType().Short()] = newTestSuiteRunQueued

	// TestSuiteRunStarted Event producer
	newTestSuiteRunStarted, _ := NewTestSuiteRunStartedEvent()
	setContext(newTestSuiteRunStarted, testSubjectId)
	exampleTestSuiteRunStartedEvent(newTestSuiteRunStarted)
	examplesProduced[newTestSuiteRunStarted.GetType().Short()] = newTestSuiteRunStarted

}