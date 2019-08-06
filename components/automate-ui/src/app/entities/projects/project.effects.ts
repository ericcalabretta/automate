import { Injectable } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { Store } from '@ngrx/store';
import { Actions, Effect, ofType } from '@ngrx/effects';
import { interval as observableInterval, of as observableOf } from 'rxjs';
import { catchError, mergeMap, map, filter, switchMap, withLatestFrom } from 'rxjs/operators';
import { identity } from 'lodash/fp';

import { NgrxStateAtom } from 'app/ngrx.reducers';
import { HttpStatus } from 'app/types/types';
import { CreateNotification } from 'app/entities/notifications/notification.actions';
import { Type } from 'app/entities/notifications/notification.model';
import { iamMajorVersion, iamMinorVersion } from 'app/entities/policies/policy.selectors';
import { ProjectRequests } from './project.requests';

import {
  GetProjectsSuccess,
  GetProjectsSuccessPayload,
  GetProjectsFailure,
  GetProject,
  GetProjectSuccess,
  GetProjectFailure,
  CreateProject,
  CreateProjectSuccess,
  CreateProjectFailure,
  DeleteProject,
  DeleteProjectSuccess,
  DeleteProjectFailure,
  UpdateProject,
  UpdateProjectFailure,
  UpdateProjectSuccess,
  ApplyRulesStart,
  ApplyRulesStartSuccess,
  ApplyRulesStartFailure,
  ApplyRulesStop,
  ApplyRulesStopSuccess,
  ApplyRulesStopFailure,
  GetApplyRulesStatus,
  GetApplyRulesStatusSuccess,
  GetApplyRulesStatusFailure,
  ProjectSuccessPayload,
  ProjectActionTypes
} from './project.actions';

const POLLING_INTERVAL_IN_SECONDS = 5;

@Injectable()
export class ProjectEffects {
  constructor(
    private actions$: Actions,
    private requests: ProjectRequests,
    private store: Store<NgrxStateAtom>
  ) { }

  @Effect()
  getProjects$ = this.actions$.pipe(
      ofType(ProjectActionTypes.GET_ALL),
      mergeMap(() =>
        this.requests.getProjects().pipe(
          map((resp: GetProjectsSuccessPayload) => new GetProjectsSuccess(resp)),
          catchError((error: HttpErrorResponse) => observableOf(new GetProjectsFailure(error))))));

  @Effect()
  getProjectsFailure$ = this.actions$.pipe(
      ofType(ProjectActionTypes.GET_ALL_FAILURE),
      map(({ payload }: GetProjectsFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not get projects: ${msg || payload.error}`
        });
      }));

  @Effect()
  getProject$ = this.actions$.pipe(
      ofType(ProjectActionTypes.GET),
      mergeMap(({ payload: { id } }: GetProject) =>
        this.requests.getProject(id).pipe(
          map((resp: ProjectSuccessPayload) => new GetProjectSuccess(resp)),
          catchError((error: HttpErrorResponse) => observableOf(new GetProjectFailure(error))))));

 @Effect()
  getProjectFailure$ = this.actions$.pipe(
      ofType(ProjectActionTypes.GET_FAILURE),
      map(({ payload }: GetProjectFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not get project: ${msg || payload.error}`
        });
      }));

  @Effect()
  createProject$ = this.actions$.pipe(
      ofType(ProjectActionTypes.CREATE),
      mergeMap(({ payload: { id, name } }: CreateProject) =>
      this.requests.createProject(id, name).pipe(
        map((resp: ProjectSuccessPayload) => new CreateProjectSuccess(resp)),
        catchError((error: HttpErrorResponse) => observableOf(new CreateProjectFailure(error))))));

  @Effect()
  createProjectSuccess$ = this.actions$.pipe(
      ofType(ProjectActionTypes.CREATE_SUCCESS),
      map(({ payload: { project } }: CreateProjectSuccess) => new CreateNotification({
      type: Type.info,
      message: `Created project ${project.id}`
    })));

  @Effect()
  createProjectFailure$ = this.actions$.pipe(
    ofType(ProjectActionTypes.CREATE_FAILURE),
    filter(({ payload }: CreateProjectFailure) => payload.status !== HttpStatus.CONFLICT),
    map(({ payload }: CreateProjectFailure) => new CreateNotification({
        type: Type.error,
        message: `Could not create project: ${payload.error.error || payload}`
      })));

  @Effect()
  deleteProject$ = this.actions$.pipe(
      ofType(ProjectActionTypes.DELETE),
      mergeMap(({ payload: { id } }: DeleteProject) =>
        this.requests.deleteProject(id).pipe(
          map(() => new DeleteProjectSuccess({id})),
          catchError((error: HttpErrorResponse) =>
            observableOf(new DeleteProjectFailure(error))))));

  @Effect()
  deleteProjectSuccess$ = this.actions$.pipe(
      ofType(ProjectActionTypes.DELETE_SUCCESS),
      map(({ payload: { id } }: DeleteProjectSuccess) => {
        return new CreateNotification({
          type: Type.info,
          message: `Deleted project ${id}.`
        });
      }));

  @Effect()
  deleteProjectFailure$ = this.actions$.pipe(
      ofType(ProjectActionTypes.DELETE_FAILURE),
      map(({ payload }: DeleteProjectFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not delete project: ${msg || payload.error}`
        });
      }));

  @Effect()
  updateProject$ = this.actions$.pipe(
      ofType(ProjectActionTypes.UPDATE),
      mergeMap(({ payload: { id, name } }: UpdateProject) =>
        this.requests.updateProject(id, name).pipe(
          map((resp: ProjectSuccessPayload) => new UpdateProjectSuccess(resp)),
          catchError((error: HttpErrorResponse) =>
            observableOf(new UpdateProjectFailure(error))))));

  @Effect()
  updateProjectFailure$ = this.actions$.pipe(
      ofType(ProjectActionTypes.UPDATE_FAILURE),
      map(({ payload }: UpdateProjectFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not update project: ${msg || payload.error}`
        });
      }));

  @Effect()
  applyRulesStart$ = this.actions$.pipe(
    ofType<ApplyRulesStart>(ProjectActionTypes.APPLY_RULES_START),
    mergeMap(() =>
      this.requests.applyRulesStart().pipe(
        switchMap(() => [
          new ApplyRulesStartSuccess(),
          new GetApplyRulesStatus()
        ]),
        catchError((error: HttpErrorResponse) =>
          observableOf(new ApplyRulesStartFailure(error))))));

  @Effect()
  applyRulesStop$ = this.actions$.pipe(
    ofType<ApplyRulesStop>(ProjectActionTypes.APPLY_RULES_STOP),
    mergeMap(() =>
      this.requests.applyRulesStop().pipe(
        switchMap(() => [
          new ApplyRulesStopSuccess(),
          new GetApplyRulesStatus()
        ]),
        catchError((error: HttpErrorResponse) => observableOf(new ApplyRulesStopFailure(error))))));

  @Effect()
  getApplyRulesStatus$ = this.actions$.pipe(
    ofType<GetApplyRulesStatus>(ProjectActionTypes.GET_APPLY_RULES_STATUS),
    mergeMap(() =>
      this.requests.getApplyRulesStatus().pipe(
        map((resp) => new GetApplyRulesStatusSuccess(resp)),
        catchError((error: HttpErrorResponse) =>
          observableOf(new GetApplyRulesStatusFailure(error))))));

  @Effect()
  getLatestApplyRulesStatus$ = observableInterval(1000 * POLLING_INTERVAL_IN_SECONDS).pipe(
    withLatestFrom(this.store.select(iamMajorVersion).pipe(filter(identity))),
    withLatestFrom(this.store.select(iamMinorVersion).pipe(filter(identity))),
    filter(([[_, major], minor]) => {
      return major === 'v2' && minor === 'v1';
    }),
    switchMap(() =>
      this.requests.getApplyRulesStatus().pipe(
        map((resp) => new GetApplyRulesStatusSuccess(resp)),
        catchError((error: HttpErrorResponse) =>
          observableOf(new GetApplyRulesStatusFailure(error))))));
}

