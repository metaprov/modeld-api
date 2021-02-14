import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../../../../../google/api/annotations_pb';
import * as github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb from '../../../../../github.com/metaprov/modeldapi/pkg/apis/inference/v1alpha1/generated_pb';


export class ListPredictorsRequest extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): ListPredictorsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListPredictorsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPredictorsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPredictorsRequest): ListPredictorsRequest.AsObject;
  static serializeBinaryToWriter(message: ListPredictorsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPredictorsRequest;
  static deserializeBinaryFromReader(message: ListPredictorsRequest, reader: jspb.BinaryReader): ListPredictorsRequest;
}

export namespace ListPredictorsRequest {
  export type AsObject = {
    namespace: string,
    labelsMap: Array<[string, string]>,
  }
}

export class ListPredictorsResponse extends jspb.Message {
  getItems(): github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.PredictorList | undefined;
  setItems(value?: github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.PredictorList): ListPredictorsResponse;
  hasItems(): boolean;
  clearItems(): ListPredictorsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPredictorsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPredictorsResponse): ListPredictorsResponse.AsObject;
  static serializeBinaryToWriter(message: ListPredictorsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPredictorsResponse;
  static deserializeBinaryFromReader(message: ListPredictorsResponse, reader: jspb.BinaryReader): ListPredictorsResponse;
}

export namespace ListPredictorsResponse {
  export type AsObject = {
    items?: github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.PredictorList.AsObject,
  }
}

export class CreatePredictorResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePredictorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePredictorResponse): CreatePredictorResponse.AsObject;
  static serializeBinaryToWriter(message: CreatePredictorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePredictorResponse;
  static deserializeBinaryFromReader(message: CreatePredictorResponse, reader: jspb.BinaryReader): CreatePredictorResponse;
}

export namespace CreatePredictorResponse {
  export type AsObject = {
  }
}

export class CreatePredictorRequest extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): CreatePredictorRequest;

  getName(): string;
  setName(value: string): CreatePredictorRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): CreatePredictorRequest;

  getSpec(): github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.PredictorSpec | undefined;
  setSpec(value?: github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.PredictorSpec): CreatePredictorRequest;
  hasSpec(): boolean;
  clearSpec(): CreatePredictorRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePredictorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePredictorRequest): CreatePredictorRequest.AsObject;
  static serializeBinaryToWriter(message: CreatePredictorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePredictorRequest;
  static deserializeBinaryFromReader(message: CreatePredictorRequest, reader: jspb.BinaryReader): CreatePredictorRequest;
}

export namespace CreatePredictorRequest {
  export type AsObject = {
    namespace: string,
    name: string,
    labelsMap: Array<[string, string]>,
    spec?: github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.PredictorSpec.AsObject,
  }
}

export class UpdatePredictorRequest extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): UpdatePredictorRequest;

  getName(): string;
  setName(value: string): UpdatePredictorRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): UpdatePredictorRequest;

  getSpec(): github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.PredictorSpec | undefined;
  setSpec(value?: github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.PredictorSpec): UpdatePredictorRequest;
  hasSpec(): boolean;
  clearSpec(): UpdatePredictorRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePredictorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePredictorRequest): UpdatePredictorRequest.AsObject;
  static serializeBinaryToWriter(message: UpdatePredictorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePredictorRequest;
  static deserializeBinaryFromReader(message: UpdatePredictorRequest, reader: jspb.BinaryReader): UpdatePredictorRequest;
}

export namespace UpdatePredictorRequest {
  export type AsObject = {
    namespace: string,
    name: string,
    labelsMap: Array<[string, string]>,
    spec?: github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.PredictorSpec.AsObject,
  }
}

export class UpdatePredictorResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePredictorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePredictorResponse): UpdatePredictorResponse.AsObject;
  static serializeBinaryToWriter(message: UpdatePredictorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePredictorResponse;
  static deserializeBinaryFromReader(message: UpdatePredictorResponse, reader: jspb.BinaryReader): UpdatePredictorResponse;
}

export namespace UpdatePredictorResponse {
  export type AsObject = {
  }
}

export class GetPredictorRequest extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): GetPredictorRequest;

  getName(): string;
  setName(value: string): GetPredictorRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPredictorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetPredictorRequest): GetPredictorRequest.AsObject;
  static serializeBinaryToWriter(message: GetPredictorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPredictorRequest;
  static deserializeBinaryFromReader(message: GetPredictorRequest, reader: jspb.BinaryReader): GetPredictorRequest;
}

export namespace GetPredictorRequest {
  export type AsObject = {
    namespace: string,
    name: string,
  }
}

export class GetPredictorResponse extends jspb.Message {
  getItem(): github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.Predictor | undefined;
  setItem(value?: github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.Predictor): GetPredictorResponse;
  hasItem(): boolean;
  clearItem(): GetPredictorResponse;

  getYaml(): string;
  setYaml(value: string): GetPredictorResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPredictorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetPredictorResponse): GetPredictorResponse.AsObject;
  static serializeBinaryToWriter(message: GetPredictorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPredictorResponse;
  static deserializeBinaryFromReader(message: GetPredictorResponse, reader: jspb.BinaryReader): GetPredictorResponse;
}

export namespace GetPredictorResponse {
  export type AsObject = {
    item?: github_com_metaprov_modeldapi_pkg_apis_inference_v1alpha1_generated_pb.Predictor.AsObject,
    yaml: string,
  }
}

export class DeletePredictorRequest extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): DeletePredictorRequest;

  getName(): string;
  setName(value: string): DeletePredictorRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeletePredictorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeletePredictorRequest): DeletePredictorRequest.AsObject;
  static serializeBinaryToWriter(message: DeletePredictorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeletePredictorRequest;
  static deserializeBinaryFromReader(message: DeletePredictorRequest, reader: jspb.BinaryReader): DeletePredictorRequest;
}

export namespace DeletePredictorRequest {
  export type AsObject = {
    namespace: string,
    name: string,
  }
}

export class DeletePredictorResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeletePredictorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeletePredictorResponse): DeletePredictorResponse.AsObject;
  static serializeBinaryToWriter(message: DeletePredictorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeletePredictorResponse;
  static deserializeBinaryFromReader(message: DeletePredictorResponse, reader: jspb.BinaryReader): DeletePredictorResponse;
}

export namespace DeletePredictorResponse {
  export type AsObject = {
  }
}

export class PredictOneRequest extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): PredictOneRequest;

  getName(): string;
  setName(value: string): PredictOneRequest;

  getFields(): string;
  setFields(value: string): PredictOneRequest;

  getValues(): string;
  setValues(value: string): PredictOneRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PredictOneRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PredictOneRequest): PredictOneRequest.AsObject;
  static serializeBinaryToWriter(message: PredictOneRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PredictOneRequest;
  static deserializeBinaryFromReader(message: PredictOneRequest, reader: jspb.BinaryReader): PredictOneRequest;
}

export namespace PredictOneRequest {
  export type AsObject = {
    namespace: string,
    name: string,
    fields: string,
    values: string,
  }
}

export class PredictOneResponse extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): PredictOneResponse;

  getName(): string;
  setName(value: string): PredictOneResponse;

  getLabel(): string;
  setLabel(value: string): PredictOneResponse;

  getScore(): number;
  setScore(value: number): PredictOneResponse;

  getProbaMap(): jspb.Map<string, number>;
  clearProbaMap(): PredictOneResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PredictOneResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PredictOneResponse): PredictOneResponse.AsObject;
  static serializeBinaryToWriter(message: PredictOneResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PredictOneResponse;
  static deserializeBinaryFromReader(message: PredictOneResponse, reader: jspb.BinaryReader): PredictOneResponse;
}

export namespace PredictOneResponse {
  export type AsObject = {
    namespace: string,
    name: string,
    label: string,
    score: number,
    probaMap: Array<[string, number]>,
  }
}
