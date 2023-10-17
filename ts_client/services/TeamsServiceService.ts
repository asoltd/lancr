/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { rpcStatus } from '../models/rpcStatus';
import type { v1CreateTeamRequest } from '../models/v1CreateTeamRequest';
import type { v1CreateTeamResponse } from '../models/v1CreateTeamResponse';
import type { v1DeleteTeamResponse } from '../models/v1DeleteTeamResponse';
import type { v1ListTeamsResponse } from '../models/v1ListTeamsResponse';
import type { v1ReadTeamResponse } from '../models/v1ReadTeamResponse';
import type { v1UpdateTeamRequest } from '../models/v1UpdateTeamRequest';
import type { v1UpdateTeamResponse } from '../models/v1UpdateTeamResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class TeamsServiceService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * @param pageSize
     * @param pageToken
     * @returns v1ListTeamsResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public teamsServiceListTeams(
        pageSize?: number,
        pageToken?: number,
    ): CancelablePromise<v1ListTeamsResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/v1/teams',
            query: {
                'pageSize': pageSize,
                'pageToken': pageToken,
            },
        });
    }

    /**
     * @param body
     * @returns v1CreateTeamResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public teamsServiceCreateTeam(
        body: v1CreateTeamRequest,
    ): CancelablePromise<v1CreateTeamResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/v1/teams',
            body: body,
        });
    }

    /**
     * @param body
     * @returns v1UpdateTeamResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public teamsServiceUpdateTeam(
        body: v1UpdateTeamRequest,
    ): CancelablePromise<v1UpdateTeamResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/v1/teams',
            body: body,
        });
    }

    /**
     * @param id
     * @returns v1ReadTeamResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public teamsServiceReadTeam(
        id: string,
    ): CancelablePromise<v1ReadTeamResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/v1/teams/{id}',
            path: {
                'id': id,
            },
        });
    }

    /**
     * @param id
     * @returns v1DeleteTeamResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public teamsServiceDeleteTeam(
        id: string,
    ): CancelablePromise<v1DeleteTeamResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/v1/teams/{id}',
            path: {
                'id': id,
            },
        });
    }

}
