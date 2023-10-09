/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { rpcStatus } from '../models/rpcStatus';
import type { v1CreateHeroRequest } from '../models/v1CreateHeroRequest';
import type { v1CreateHeroResponse } from '../models/v1CreateHeroResponse';
import type { v1DeleteHeroResponse } from '../models/v1DeleteHeroResponse';
import type { v1GetQuestCreatorResponse } from '../models/v1GetQuestCreatorResponse';
import type { v1GetTotalHeroesResponse } from '../models/v1GetTotalHeroesResponse';
import type { v1ListHeroesResponse } from '../models/v1ListHeroesResponse';
import type { v1ReadHeroResponse } from '../models/v1ReadHeroResponse';
import type { v1UpdateHeroRequest } from '../models/v1UpdateHeroRequest';
import type { v1UpdateHeroResponse } from '../models/v1UpdateHeroResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class HeroServiceService {

    /**
     * @param pageSize
     * @param pageToken
     * @returns v1ListHeroesResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public static heroServiceListHeroes(
        pageSize?: number,
        pageToken?: number,
    ): CancelablePromise<v1ListHeroesResponse | rpcStatus> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/v1/heroes',
            query: {
                'pageSize': pageSize,
                'pageToken': pageToken,
            },
        });
    }

    /**
     * @param body
     * @returns v1CreateHeroResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public static heroServiceCreateHero(
        body: v1CreateHeroRequest,
    ): CancelablePromise<v1CreateHeroResponse | rpcStatus> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/v1/heroes',
            body: body,
        });
    }

    /**
     * @param body
     * @returns v1UpdateHeroResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public static heroServiceUpdateHero(
        body: v1UpdateHeroRequest,
    ): CancelablePromise<v1UpdateHeroResponse | rpcStatus> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/v1/heroes',
            body: body,
        });
    }

    /**
     * @param id
     * @returns v1ReadHeroResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public static heroServiceReadHero(
        id: string,
    ): CancelablePromise<v1ReadHeroResponse | rpcStatus> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/v1/heroes/{id}',
            path: {
                'id': id,
            },
        });
    }

    /**
     * @param id
     * @returns v1DeleteHeroResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public static heroServiceDeleteHero(
        id: string,
    ): CancelablePromise<v1DeleteHeroResponse | rpcStatus> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/v1/heroes/{id}',
            path: {
                'id': id,
            },
        });
    }

    /**
     * @param questId
     * @returns v1GetQuestCreatorResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public static heroServiceGetQuestCreator(
        questId?: string,
    ): CancelablePromise<v1GetQuestCreatorResponse | rpcStatus> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/v1/heroes:getQuestCreator',
            query: {
                'questId': questId,
            },
        });
    }

    /**
     * @returns v1GetTotalHeroesResponse A successful response.
     * @returns rpcStatus An unexpected error response.
     * @throws ApiError
     */
    public static heroServiceGetTotalHeroes(): CancelablePromise<v1GetTotalHeroesResponse | rpcStatus> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/v1/heroes:getTotalHeroes',
        });
    }

}
