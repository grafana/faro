import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'
import { Observable, of, from } from '../rxjsStub';
import {mergeMap, map} from  '../rxjsStub';
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

import { DefaultApiRequestFactory, DefaultApiResponseProcessor} from "../apis/DefaultApi";
export class ObservableDefaultApi {
    private requestFactory: DefaultApiRequestFactory;
    private responseProcessor: DefaultApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: DefaultApiRequestFactory,
        responseProcessor?: DefaultApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new DefaultApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new DefaultApiResponseProcessor();
    }

    /**
     * Submit a Faro payload
     * @param appKey The application key is provided by your endpoint provider. 
     * @param payload Optional description in *Markdown*
     */
    public collectAppKeyPostWithHttpInfo(appKey: string, payload: Payload, _options?: Configuration): Observable<HttpInfo<void>> {
        const requestContextPromise = this.requestFactory.collectAppKeyPost(appKey, payload, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (const middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (const middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.collectAppKeyPostWithHttpInfo(rsp)));
            }));
    }

    /**
     * Submit a Faro payload
     * @param appKey The application key is provided by your endpoint provider. 
     * @param payload Optional description in *Markdown*
     */
    public collectAppKeyPost(appKey: string, payload: Payload, _options?: Configuration): Observable<void> {
        return this.collectAppKeyPostWithHttpInfo(appKey, payload, _options).pipe(map((apiResponse: HttpInfo<void>) => apiResponse.data));
    }

}
