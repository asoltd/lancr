/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { protobufAny } from './protobufAny';
import type { v1FirebaseInfo } from './v1FirebaseInfo';

export type v1Token = {
    authTime?: string;
    issuer?: string;
    audience?: string;
    expires?: string;
    issuedAt?: string;
    subject?: string;
    uid?: string;
    firebaseInfo?: v1FirebaseInfo;
    claims?: Record<string, protobufAny>;
};

