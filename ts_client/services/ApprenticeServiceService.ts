/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { rpcStatus } from '../models/rpcStatus';
import type { v1CreateApprenticeRequest } from '../models/v1CreateApprenticeRequest';
import type { v1CreateApprenticeResponse } from '../models/v1CreateApprenticeResponse';
import type { v1DeleteApprenticeResponse } from '../models/v1DeleteApprenticeResponse';
import type { v1ListApprenticesResponse } from '../models/v1ListApprenticesResponse';
import type { v1ReadApprenticeResponse } from '../models/v1ReadApprenticeResponse';
import type { v1UpdateApprenticeRequest } from '../models/v1UpdateApprenticeRequest';
import type { v1UpdateApprenticeResponse } from '../models/v1UpdateApprenticeResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class ApprenticeServiceService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * @param pageSize
     * @param pageToken
     * @returns v1ListApprenticesResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public apprenticeServiceListApprentices(
        pageSize?: number,
        pageToken?: number,
    ): CancelablePromise<v1ListApprenticesResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/v1/apprentices',
            query: {
                'pageSize': pageSize,
                'pageToken': pageToken,
            },
        });
    }

    /**
     * @param body
     * @returns v1CreateApprenticeResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public apprenticeServiceCreateApprentice(
        body: v1CreateApprenticeRequest,
    ): CancelablePromise<v1CreateApprenticeResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/v1/apprentices',
            body: body,
        });
    }

    /**
     * @param body
     * @returns v1UpdateApprenticeResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public apprenticeServiceUpdateApprentice(
        body: v1UpdateApprenticeRequest,
    ): CancelablePromise<v1UpdateApprenticeResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/v1/apprentices',
            body: body,
        });
    }

    /**
     * @param id
     * @returns v1ReadApprenticeResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public apprenticeServiceReadApprentice(
        id: string,
    ): CancelablePromise<v1ReadApprenticeResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/v1/apprentices/{id}',
            path: {
                'id': id,
            },
        });
    }

    /**
     * @param id
     * @returns v1DeleteApprenticeResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public apprenticeServiceDeleteApprentice(
        id: string,
    ): CancelablePromise<v1DeleteApprenticeResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/v1/apprentices/{id}',
            path: {
                'id': id,
            },
        });
    }

}
