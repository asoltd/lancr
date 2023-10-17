/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { v1Apprentice } from './v1Apprentice';

export type v1Bid = {
    id?: string;
    bidderId?: string;
    createdAt?: string;
    updatedAt?: string;
    rate?: number;
    amount?: number;
    currency?: string;
    timeRequired?: string;
    workingTime?: string;
    questId?: string;
    apprentice?: v1Apprentice;
    apprenticeRate?: number;
    apprenticeCut?: number;
    totalEarnings?: number;
    status?: string;
    totalBidValue?: number;
};

