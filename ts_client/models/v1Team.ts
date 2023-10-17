/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { v1Bid } from './v1Bid';
import type { v1Hero } from './v1Hero';
import type { v1Industry } from './v1Industry';
import type { v1RoleCategory } from './v1RoleCategory';

export type v1Team = {
    id: string;
    creatorId?: string;
    title?: string;
    description?: string;
    highlight?: string;
    industry?: v1Industry;
    image?: string;
    timeEstimate?: string;
    bidders?: v1Bid;
    members?: v1Hero;
    roleCategories?: v1RoleCategory;
    createdAt?: string;
};

