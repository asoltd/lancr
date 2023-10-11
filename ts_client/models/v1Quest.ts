/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { lancrv1Tag } from './lancrv1Tag';
import type { v1Image } from './v1Image';

export type v1Quest = {
    id: string;
    creatorId?: string;
    reward?: number;
    title?: string;
    description?: string;
    tags?: Array<lancrv1Tag>;
    images?: Array<v1Image>;
    bidders?: Array<string>;
    status?: string;
    createdAt?: string;
    summary?: string;
    level?: number;
};

