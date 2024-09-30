import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'

import { AppMeta } from '../models/AppMeta';
import { BrowserMeta } from '../models/BrowserMeta';
import { Event } from '../models/Event';
import { Exception } from '../models/Exception';
import { ExceptionTrace } from '../models/ExceptionTrace';
import { Log } from '../models/Log';
import { LogTrace } from '../models/LogTrace';
import { Measurement } from '../models/Measurement';
import { PageMeta } from '../models/PageMeta';
import { Payload } from '../models/Payload';
import { PayloadMeta } from '../models/PayloadMeta';
import { PayloadMetaView } from '../models/PayloadMetaView';
import { PayloadTraces } from '../models/PayloadTraces';
import { PayloadTracesResourceSpansInner } from '../models/PayloadTracesResourceSpansInner';
import { PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner } from '../models/PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner';
import { PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerInstrumentationLibrary } from '../models/PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerInstrumentationLibrary';
import { PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner } from '../models/PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner';
import { PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerAttributes } from '../models/PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerAttributes';
import { PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerStatus } from '../models/PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerStatus';
import { PayloadTracesResourceSpansInnerResource } from '../models/PayloadTracesResourceSpansInnerResource';
import { PayloadTracesResourceSpansInnerResourceAttributes } from '../models/PayloadTracesResourceSpansInnerResourceAttributes';
import { SdkMeta } from '../models/SdkMeta';
import { SessionMeta } from '../models/SessionMeta';
import { SpanEvent } from '../models/SpanEvent';
import { StackFrame } from '../models/StackFrame';
import { UserMeta } from '../models/UserMeta';
import { ObservableDefaultApi } from './ObservableAPI';

import { DefaultApiRequestFactory, DefaultApiResponseProcessor} from "../apis/DefaultApi";
export class PromiseDefaultApi {
    private api: ObservableDefaultApi

    public constructor(
        configuration: Configuration,
        requestFactory?: DefaultApiRequestFactory,
        responseProcessor?: DefaultApiResponseProcessor
    ) {
        this.api = new ObservableDefaultApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * Submit a Faro payload
     * @param appKey The application key is provided by your endpoint provider. 
     * @param payload Optional description in *Markdown*
     */
    public collectAppKeyPostWithHttpInfo(appKey: string, payload: Payload, _options?: Configuration): Promise<HttpInfo<void>> {
        const result = this.api.collectAppKeyPostWithHttpInfo(appKey, payload, _options);
        return result.toPromise();
    }

    /**
     * Submit a Faro payload
     * @param appKey The application key is provided by your endpoint provider. 
     * @param payload Optional description in *Markdown*
     */
    public collectAppKeyPost(appKey: string, payload: Payload, _options?: Configuration): Promise<void> {
        const result = this.api.collectAppKeyPost(appKey, payload, _options);
        return result.toPromise();
    }


}


