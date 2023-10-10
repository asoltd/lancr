/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { v1WorkingHours } from './v1WorkingHours';

export type v1Apprentice = {
    id?: string;
    rate?: number;
    workingHours?: v1WorkingHours;
    mentor?: string;
    favoriteTo?: string;
    workingDays?: Array<string>;
};

