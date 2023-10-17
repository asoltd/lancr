/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { lancrv1Location } from './lancrv1Location';
import type { v1Experience } from './v1Experience';
import type { v1Name } from './v1Name';

export type v1Hero = {
    id: string;
    firebaseId?: string;
    name?: v1Name;
    visibility?: string;
    profilePicture?: string;
    email?: string;
    phoneNumber?: string;
    location?: lancrv1Location;
    userName?: string;
    bio?: string;
    workType?: string;
    subWorkType?: string;
    twitter?: string;
    linkedin?: string;
    website?: string;
    experience?: Array<v1Experience>;
    rating?: number;
    xp?: string;
    region?: string;
    language?: string;
    level?: number;
};

