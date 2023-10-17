/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { lancrv1Location } from './lancrv1Location';
import type { v1Mentor } from './v1Mentor';
import type { v1Rating } from './v1Rating';
import type { v1WorkingHours } from './v1WorkingHours';

export type v1Apprentice = {
    id: string;
    heroId?: string;
    rate?: number;
    workingHours?: v1WorkingHours;
    mentors?: Array<v1Mentor>;
    favoriteTo?: string;
    workingDays?: Array<string>;
    ratings?: Array<v1Rating>;
    level?: number;
    location?: lancrv1Location;
};

