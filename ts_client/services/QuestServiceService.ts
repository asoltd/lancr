/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { rpcStatus } from '../models/rpcStatus';
import type { v1CreateQuestRequest } from '../models/v1CreateQuestRequest';
import type { v1CreateQuestResponse } from '../models/v1CreateQuestResponse';
import type { v1DeleteQuestResponse } from '../models/v1DeleteQuestResponse';
import type { v1GetTopQuestersResponse } from '../models/v1GetTopQuestersResponse';
import type { v1ListQuestsResponse } from '../models/v1ListQuestsResponse';
import type { v1ReadQuestResponse } from '../models/v1ReadQuestResponse';
import type { v1UpdateQuestRequest } from '../models/v1UpdateQuestRequest';
import type { v1UpdateQuestResponse } from '../models/v1UpdateQuestResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class QuestServiceService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * @param pageSize
     * @param pageToken
     * @returns v1ListQuestsResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public questServiceListQuests(
        pageSize?: number,
        pageToken?: number,
    ): CancelablePromise<v1ListQuestsResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/v1/quests',
            query: {
                'pageSize': pageSize,
                'pageToken': pageToken,
            },
        });
    }

    /**
     * @param body
     * @returns v1CreateQuestResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public questServiceCreateQuest(
        body: v1CreateQuestRequest,
    ): CancelablePromise<v1CreateQuestResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/v1/quests',
            body: body,
        });
    }

    /**
     * @param body
     * @returns v1UpdateQuestResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public questServiceUpdateQuest(
        body: v1UpdateQuestRequest,
    ): CancelablePromise<v1UpdateQuestResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/v1/quests',
            body: body,
        });
    }

    /**
     * @param id
     * @returns v1ReadQuestResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public questServiceReadQuest(
        id: string,
    ): CancelablePromise<v1ReadQuestResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/v1/quests/{id}',
            path: {
                'id': id,
            },
        });
    }

    /**
     * @param id
     * @returns v1DeleteQuestResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public questServiceDeleteQuest(
        id: string,
    ): CancelablePromise<v1DeleteQuestResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/v1/quests/{id}',
            path: {
                'id': id,
            },
        });
    }

    /**
     * @returns v1GetTopQuestersResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public questServiceGetTopQuesters(): CancelablePromise<v1GetTopQuestersResponse | rpcStatus> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/v1/quests:getTopQuesters',
        });
    }

}
