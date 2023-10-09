/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { HeroExperience } from './HeroExperience';
import type { lancrv1Location } from './lancrv1Location';
import type { v1Image } from './v1Image';
import type { v1Name } from './v1Name';

export type v1Hero = {
    id: string;
    displayName?: string;
    visibility?: string;
    profilePicture?: v1Image;
    email?: string;
    phoneNumber?: string;
    name?: v1Name;
    location?: lancrv1Location;
    userName?: string;
    bio?: string;
    workType?: string;
    subWorkType?: string;
    twitter?: string;
    linkedin?: string;
    website?: string;
    experiences?: Array<HeroExperience>;
    rating?: number;
    xp?: string;
    region?: string;
    language?: string;
    level?: number;
};
