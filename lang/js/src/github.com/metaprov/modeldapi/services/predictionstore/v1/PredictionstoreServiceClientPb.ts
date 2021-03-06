/**
 * @fileoverview gRPC-Web generated client stub for github.com.metaprov.modeldapi.services.predictionstore.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as github_com_metaprov_modeldapi_services_predictionstore_v1_predictionstore_pb from '../../../../../../github.com/metaprov/modeldapi/services/predictionstore/v1/predictionstore_pb';


export class PredictionStoreServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoIngest = new grpcWeb.AbstractClientBase.MethodInfo(
    github_com_metaprov_modeldapi_services_predictionstore_v1_predictionstore_pb.PredictionIngestResponse,
    (request: github_com_metaprov_modeldapi_services_predictionstore_v1_predictionstore_pb.PredictionsIngestRequest) => {
      return request.serializeBinary();
    },
    github_com_metaprov_modeldapi_services_predictionstore_v1_predictionstore_pb.PredictionIngestResponse.deserializeBinary
  );

  ingest(
    request: github_com_metaprov_modeldapi_services_predictionstore_v1_predictionstore_pb.PredictionsIngestRequest,
    metadata: grpcWeb.Metadata | null): Promise<github_com_metaprov_modeldapi_services_predictionstore_v1_predictionstore_pb.PredictionIngestResponse>;

  ingest(
    request: github_com_metaprov_modeldapi_services_predictionstore_v1_predictionstore_pb.PredictionsIngestRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: github_com_metaprov_modeldapi_services_predictionstore_v1_predictionstore_pb.PredictionIngestResponse) => void): grpcWeb.ClientReadableStream<github_com_metaprov_modeldapi_services_predictionstore_v1_predictionstore_pb.PredictionIngestResponse>;

  ingest(
    request: github_com_metaprov_modeldapi_services_predictionstore_v1_predictionstore_pb.PredictionsIngestRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: github_com_metaprov_modeldapi_services_predictionstore_v1_predictionstore_pb.PredictionIngestResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/github.com.metaprov.modeldapi.services.predictionstore.v1.PredictionStoreService/Ingest',
        request,
        metadata || {},
        this.methodInfoIngest,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/github.com.metaprov.modeldapi.services.predictionstore.v1.PredictionStoreService/Ingest',
    request,
    metadata || {},
    this.methodInfoIngest);
  }

}

