/**
 * Faro Collector API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { LogTrace } from '../models/LogTrace';
import { HttpFile } from '../http/http';

export class Log {
    'message': string;
    'level': string;
    'context'?: { [key: string]: string; };
    'timestamp': Date;
    'trace'?: LogTrace;

    static readonly discriminator: string | undefined = undefined;

    static readonly mapping: {[index: string]: string} | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "message",
            "baseName": "message",
            "type": "string",
            "format": ""
        },
        {
            "name": "level",
            "baseName": "level",
            "type": "string",
            "format": ""
        },
        {
            "name": "context",
            "baseName": "context",
            "type": "{ [key: string]: string; }",
            "format": ""
        },
        {
            "name": "timestamp",
            "baseName": "timestamp",
            "type": "Date",
            "format": "date-time"
        },
        {
            "name": "trace",
            "baseName": "trace",
            "type": "LogTrace",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return Log.attributeTypeMap;
    }

    public constructor() {
    }
}