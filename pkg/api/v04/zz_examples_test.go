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

package v04_test

import (
	"github.com/cdevents/sdk-go/pkg/api"
	specv04 "github.com/cdevents/sdk-go/pkg/api/v04"
)

func init() {
	// Create events equal to examples in the spec
	examplesProduced = make(map[string]api.CDEventV04)

	// Set up test links
	tags := api.Tags{
		"foo1": "bar",
		"foo2": "bar",
	}
	reference := api.EventReference{
		ContextId: testContextId,
	}
	elr := api.NewEmbeddedLinkRelation()
	elr.SetTags(tags)
	elr.SetLinkKind("TRIGGER")
	elr.SetTarget(reference)
	elp := api.NewEmbeddedLinkPath()
	elp.SetTags(tags)
	elp.SetFrom(reference)
	ele := api.NewEmbeddedLinkEnd()
	ele.SetTags(tags)
	ele.SetFrom(reference)
	testLinks = api.EmbeddedLinksArray{
		elr, elp, ele,
	}

	// Create events for test
	// ArtifactDeleted Event producer
	newArtifactDeleted, _ := specv04.NewArtifactDeletedEvent()
	setContext(newArtifactDeleted, testSubjectId)
	setContextV04(newArtifactDeleted, true, true)
	exampleArtifactDeletedEvent(newArtifactDeleted)
	examplesProduced[newArtifactDeleted.GetType().Short()] = newArtifactDeleted

	// ArtifactDownloaded Event producer
	newArtifactDownloaded, _ := specv04.NewArtifactDownloadedEvent()
	setContext(newArtifactDownloaded, testSubjectId)
	setContextV04(newArtifactDownloaded, true, true)
	exampleArtifactDownloadedEvent(newArtifactDownloaded)
	examplesProduced[newArtifactDownloaded.GetType().Short()] = newArtifactDownloaded

	// ArtifactPackaged Event producer
	newArtifactPackaged, _ := specv04.NewArtifactPackagedEvent()
	setContext(newArtifactPackaged, testSubjectId)
	setContextV04(newArtifactPackaged, true, true)
	exampleArtifactPackagedEvent(newArtifactPackaged)
	examplesProduced[newArtifactPackaged.GetType().Short()] = newArtifactPackaged

	// ArtifactPublished Event producer
	newArtifactPublished, _ := specv04.NewArtifactPublishedEvent()
	setContext(newArtifactPublished, testSubjectId)
	setContextV04(newArtifactPublished, true, true)
	exampleArtifactPublishedEvent(newArtifactPublished)
	examplesProduced[newArtifactPublished.GetType().Short()] = newArtifactPublished

	// ArtifactSigned Event producer
	newArtifactSigned, _ := specv04.NewArtifactSignedEvent()
	setContext(newArtifactSigned, testSubjectId)
	setContextV04(newArtifactSigned, true, true)
	exampleArtifactSignedEvent(newArtifactSigned)
	examplesProduced[newArtifactSigned.GetType().Short()] = newArtifactSigned

	// BranchCreated Event producer
	newBranchCreated, _ := specv04.NewBranchCreatedEvent()
	setContext(newBranchCreated, testSubjectId)
	setContextV04(newBranchCreated, true, true)
	exampleBranchCreatedEvent(newBranchCreated)
	examplesProduced[newBranchCreated.GetType().Short()] = newBranchCreated

	// BranchDeleted Event producer
	newBranchDeleted, _ := specv04.NewBranchDeletedEvent()
	setContext(newBranchDeleted, testSubjectId)
	setContextV04(newBranchDeleted, true, true)
	exampleBranchDeletedEvent(newBranchDeleted)
	examplesProduced[newBranchDeleted.GetType().Short()] = newBranchDeleted

	// BuildFinished Event producer
	newBuildFinished, _ := specv04.NewBuildFinishedEvent()
	setContext(newBuildFinished, testSubjectId)
	setContextV04(newBuildFinished, true, true)
	exampleBuildFinishedEvent(newBuildFinished)
	examplesProduced[newBuildFinished.GetType().Short()] = newBuildFinished

	// BuildQueued Event producer
	newBuildQueued, _ := specv04.NewBuildQueuedEvent()
	setContext(newBuildQueued, testSubjectId)
	setContextV04(newBuildQueued, true, true)
	exampleBuildQueuedEvent(newBuildQueued)
	examplesProduced[newBuildQueued.GetType().Short()] = newBuildQueued

	// BuildStarted Event producer
	newBuildStarted, _ := specv04.NewBuildStartedEvent()
	setContext(newBuildStarted, testSubjectId)
	setContextV04(newBuildStarted, true, true)
	exampleBuildStartedEvent(newBuildStarted)
	examplesProduced[newBuildStarted.GetType().Short()] = newBuildStarted

	// ChangeAbandoned Event producer
	newChangeAbandoned, _ := specv04.NewChangeAbandonedEvent()
	setContext(newChangeAbandoned, testSubjectId)
	setContextV04(newChangeAbandoned, true, true)
	exampleChangeAbandonedEvent(newChangeAbandoned)
	examplesProduced[newChangeAbandoned.GetType().Short()] = newChangeAbandoned

	// ChangeCreated Event producer
	newChangeCreated, _ := specv04.NewChangeCreatedEvent()
	setContext(newChangeCreated, testSubjectId)
	setContextV04(newChangeCreated, true, true)
	exampleChangeCreatedEvent(newChangeCreated)
	examplesProduced[newChangeCreated.GetType().Short()] = newChangeCreated

	// ChangeMerged Event producer
	newChangeMerged, _ := specv04.NewChangeMergedEvent()
	setContext(newChangeMerged, testSubjectId)
	setContextV04(newChangeMerged, true, true)
	exampleChangeMergedEvent(newChangeMerged)
	examplesProduced[newChangeMerged.GetType().Short()] = newChangeMerged

	// ChangeReviewed Event producer
	newChangeReviewed, _ := specv04.NewChangeReviewedEvent()
	setContext(newChangeReviewed, testSubjectId)
	setContextV04(newChangeReviewed, true, true)
	exampleChangeReviewedEvent(newChangeReviewed)
	examplesProduced[newChangeReviewed.GetType().Short()] = newChangeReviewed

	// ChangeUpdated Event producer
	newChangeUpdated, _ := specv04.NewChangeUpdatedEvent()
	setContext(newChangeUpdated, testSubjectId)
	setContextV04(newChangeUpdated, true, true)
	exampleChangeUpdatedEvent(newChangeUpdated)
	examplesProduced[newChangeUpdated.GetType().Short()] = newChangeUpdated

	// EnvironmentCreated Event producer
	newEnvironmentCreated, _ := specv04.NewEnvironmentCreatedEvent()
	setContext(newEnvironmentCreated, testSubjectId)
	setContextV04(newEnvironmentCreated, true, true)
	exampleEnvironmentCreatedEvent(newEnvironmentCreated)
	examplesProduced[newEnvironmentCreated.GetType().Short()] = newEnvironmentCreated

	// EnvironmentDeleted Event producer
	newEnvironmentDeleted, _ := specv04.NewEnvironmentDeletedEvent()
	setContext(newEnvironmentDeleted, testSubjectId)
	setContextV04(newEnvironmentDeleted, true, true)
	exampleEnvironmentDeletedEvent(newEnvironmentDeleted)
	examplesProduced[newEnvironmentDeleted.GetType().Short()] = newEnvironmentDeleted

	// EnvironmentModified Event producer
	newEnvironmentModified, _ := specv04.NewEnvironmentModifiedEvent()
	setContext(newEnvironmentModified, testSubjectId)
	setContextV04(newEnvironmentModified, true, true)
	exampleEnvironmentModifiedEvent(newEnvironmentModified)
	examplesProduced[newEnvironmentModified.GetType().Short()] = newEnvironmentModified

	// IncidentDetected Event producer
	newIncidentDetected, _ := specv04.NewIncidentDetectedEvent()
	setContext(newIncidentDetected, testSubjectId)
	setContextV04(newIncidentDetected, true, true)
	exampleIncidentDetectedEvent(newIncidentDetected)
	examplesProduced[newIncidentDetected.GetType().Short()] = newIncidentDetected

	// IncidentReported Event producer
	newIncidentReported, _ := specv04.NewIncidentReportedEvent()
	setContext(newIncidentReported, testSubjectId)
	setContextV04(newIncidentReported, true, true)
	exampleIncidentReportedEvent(newIncidentReported)
	examplesProduced[newIncidentReported.GetType().Short()] = newIncidentReported

	// IncidentResolved Event producer
	newIncidentResolved, _ := specv04.NewIncidentResolvedEvent()
	setContext(newIncidentResolved, testSubjectId)
	setContextV04(newIncidentResolved, true, true)
	exampleIncidentResolvedEvent(newIncidentResolved)
	examplesProduced[newIncidentResolved.GetType().Short()] = newIncidentResolved

	// PipelineRunFinished Event producer
	newPipelineRunFinished, _ := specv04.NewPipelineRunFinishedEvent()
	setContext(newPipelineRunFinished, testSubjectId)
	setContextV04(newPipelineRunFinished, true, true)
	examplePipelineRunFinishedEvent(newPipelineRunFinished)
	examplesProduced[newPipelineRunFinished.GetType().Short()] = newPipelineRunFinished

	// PipelineRunQueued Event producer
	newPipelineRunQueued, _ := specv04.NewPipelineRunQueuedEvent()
	setContext(newPipelineRunQueued, testSubjectId)
	setContextV04(newPipelineRunQueued, true, true)
	examplePipelineRunQueuedEvent(newPipelineRunQueued)
	examplesProduced[newPipelineRunQueued.GetType().Short()] = newPipelineRunQueued

	// PipelineRunStarted Event producer
	newPipelineRunStarted, _ := specv04.NewPipelineRunStartedEvent()
	setContext(newPipelineRunStarted, testSubjectId)
	setContextV04(newPipelineRunStarted, true, true)
	examplePipelineRunStartedEvent(newPipelineRunStarted)
	examplesProduced[newPipelineRunStarted.GetType().Short()] = newPipelineRunStarted

	// RepositoryCreated Event producer
	newRepositoryCreated, _ := specv04.NewRepositoryCreatedEvent()
	setContext(newRepositoryCreated, testSubjectId)
	setContextV04(newRepositoryCreated, true, true)
	exampleRepositoryCreatedEvent(newRepositoryCreated)
	examplesProduced[newRepositoryCreated.GetType().Short()] = newRepositoryCreated

	// RepositoryDeleted Event producer
	newRepositoryDeleted, _ := specv04.NewRepositoryDeletedEvent()
	setContext(newRepositoryDeleted, testSubjectId)
	setContextV04(newRepositoryDeleted, true, true)
	exampleRepositoryDeletedEvent(newRepositoryDeleted)
	examplesProduced[newRepositoryDeleted.GetType().Short()] = newRepositoryDeleted

	// RepositoryModified Event producer
	newRepositoryModified, _ := specv04.NewRepositoryModifiedEvent()
	setContext(newRepositoryModified, testSubjectId)
	setContextV04(newRepositoryModified, true, true)
	exampleRepositoryModifiedEvent(newRepositoryModified)
	examplesProduced[newRepositoryModified.GetType().Short()] = newRepositoryModified

	// ServiceDeployed Event producer
	newServiceDeployed, _ := specv04.NewServiceDeployedEvent()
	setContext(newServiceDeployed, testSubjectId)
	setContextV04(newServiceDeployed, true, true)
	exampleServiceDeployedEvent(newServiceDeployed)
	examplesProduced[newServiceDeployed.GetType().Short()] = newServiceDeployed

	// ServicePublished Event producer
	newServicePublished, _ := specv04.NewServicePublishedEvent()
	setContext(newServicePublished, testSubjectId)
	setContextV04(newServicePublished, true, true)
	exampleServicePublishedEvent(newServicePublished)
	examplesProduced[newServicePublished.GetType().Short()] = newServicePublished

	// ServiceRemoved Event producer
	newServiceRemoved, _ := specv04.NewServiceRemovedEvent()
	setContext(newServiceRemoved, testSubjectId)
	setContextV04(newServiceRemoved, true, true)
	exampleServiceRemovedEvent(newServiceRemoved)
	examplesProduced[newServiceRemoved.GetType().Short()] = newServiceRemoved

	// ServiceRolledback Event producer
	newServiceRolledback, _ := specv04.NewServiceRolledbackEvent()
	setContext(newServiceRolledback, testSubjectId)
	setContextV04(newServiceRolledback, true, true)
	exampleServiceRolledbackEvent(newServiceRolledback)
	examplesProduced[newServiceRolledback.GetType().Short()] = newServiceRolledback

	// ServiceUpgraded Event producer
	newServiceUpgraded, _ := specv04.NewServiceUpgradedEvent()
	setContext(newServiceUpgraded, testSubjectId)
	setContextV04(newServiceUpgraded, true, true)
	exampleServiceUpgradedEvent(newServiceUpgraded)
	examplesProduced[newServiceUpgraded.GetType().Short()] = newServiceUpgraded

	// TaskRunFinished Event producer
	newTaskRunFinished, _ := specv04.NewTaskRunFinishedEvent()
	setContext(newTaskRunFinished, testSubjectId)
	setContextV04(newTaskRunFinished, true, true)
	exampleTaskRunFinishedEvent(newTaskRunFinished)
	examplesProduced[newTaskRunFinished.GetType().Short()] = newTaskRunFinished

	// TaskRunStarted Event producer
	newTaskRunStarted, _ := specv04.NewTaskRunStartedEvent()
	setContext(newTaskRunStarted, testSubjectId)
	setContextV04(newTaskRunStarted, true, true)
	exampleTaskRunStartedEvent(newTaskRunStarted)
	examplesProduced[newTaskRunStarted.GetType().Short()] = newTaskRunStarted

	// TestCaseRunFinished Event producer
	newTestCaseRunFinished, _ := specv04.NewTestCaseRunFinishedEvent()
	setContext(newTestCaseRunFinished, testSubjectId)
	setContextV04(newTestCaseRunFinished, true, true)
	exampleTestCaseRunFinishedEvent(newTestCaseRunFinished)
	examplesProduced[newTestCaseRunFinished.GetType().Short()] = newTestCaseRunFinished

	// TestCaseRunQueued Event producer
	newTestCaseRunQueued, _ := specv04.NewTestCaseRunQueuedEvent()
	setContext(newTestCaseRunQueued, testSubjectId)
	setContextV04(newTestCaseRunQueued, true, true)
	exampleTestCaseRunQueuedEvent(newTestCaseRunQueued)
	examplesProduced[newTestCaseRunQueued.GetType().Short()] = newTestCaseRunQueued

	// TestCaseRunSkipped Event producer
	newTestCaseRunSkipped, _ := specv04.NewTestCaseRunSkippedEvent()
	setContext(newTestCaseRunSkipped, testSubjectId)
	setContextV04(newTestCaseRunSkipped, true, true)
	exampleTestCaseRunSkippedEvent(newTestCaseRunSkipped)
	examplesProduced[newTestCaseRunSkipped.GetType().Short()] = newTestCaseRunSkipped

	// TestCaseRunStarted Event producer
	newTestCaseRunStarted, _ := specv04.NewTestCaseRunStartedEvent()
	setContext(newTestCaseRunStarted, testSubjectId)
	setContextV04(newTestCaseRunStarted, true, true)
	exampleTestCaseRunStartedEvent(newTestCaseRunStarted)
	examplesProduced[newTestCaseRunStarted.GetType().Short()] = newTestCaseRunStarted

	// TestOutputPublished Event producer
	newTestOutputPublished, _ := specv04.NewTestOutputPublishedEvent()
	setContext(newTestOutputPublished, testSubjectId)
	setContextV04(newTestOutputPublished, true, true)
	exampleTestOutputPublishedEvent(newTestOutputPublished)
	examplesProduced[newTestOutputPublished.GetType().Short()] = newTestOutputPublished

	// TestSuiteRunFinished Event producer
	newTestSuiteRunFinished, _ := specv04.NewTestSuiteRunFinishedEvent()
	setContext(newTestSuiteRunFinished, testSubjectId)
	setContextV04(newTestSuiteRunFinished, true, true)
	exampleTestSuiteRunFinishedEvent(newTestSuiteRunFinished)
	examplesProduced[newTestSuiteRunFinished.GetType().Short()] = newTestSuiteRunFinished

	// TestSuiteRunQueued Event producer
	newTestSuiteRunQueued, _ := specv04.NewTestSuiteRunQueuedEvent()
	setContext(newTestSuiteRunQueued, testSubjectId)
	setContextV04(newTestSuiteRunQueued, true, true)
	exampleTestSuiteRunQueuedEvent(newTestSuiteRunQueued)
	examplesProduced[newTestSuiteRunQueued.GetType().Short()] = newTestSuiteRunQueued

	// TestSuiteRunStarted Event producer
	newTestSuiteRunStarted, _ := specv04.NewTestSuiteRunStartedEvent()
	setContext(newTestSuiteRunStarted, testSubjectId)
	setContextV04(newTestSuiteRunStarted, true, true)
	exampleTestSuiteRunStartedEvent(newTestSuiteRunStarted)
	examplesProduced[newTestSuiteRunStarted.GetType().Short()] = newTestSuiteRunStarted

	// TicketClosed Event producer
	newTicketClosed, _ := specv04.NewTicketClosedEvent()
	setContext(newTicketClosed, testSubjectId)
	setContextV04(newTicketClosed, true, true)
	exampleTicketClosedEvent(newTicketClosed)
	examplesProduced[newTicketClosed.GetType().Short()] = newTicketClosed

	// TicketCreated Event producer
	newTicketCreated, _ := specv04.NewTicketCreatedEvent()
	setContext(newTicketCreated, testSubjectId)
	setContextV04(newTicketCreated, true, true)
	exampleTicketCreatedEvent(newTicketCreated)
	examplesProduced[newTicketCreated.GetType().Short()] = newTicketCreated

	// TicketUpdated Event producer
	newTicketUpdated, _ := specv04.NewTicketUpdatedEvent()
	setContext(newTicketUpdated, testSubjectId)
	setContextV04(newTicketUpdated, true, true)
	exampleTicketUpdatedEvent(newTicketUpdated)
	examplesProduced[newTicketUpdated.GetType().Short()] = newTicketUpdated

	// CustomType Event producer
	newCustomType, _ := specv04.NewCustomTypeEvent()
	setContext(newCustomType, testSubjectId)
	setContextV04(newCustomType, true, true)
	exampleCustomTypeEvent(newCustomType)
	examplesProduced[newCustomType.GetType().Short()] = newCustomType

}
