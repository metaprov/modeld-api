import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../../../../../../google/api/annotations_pb';
import * as github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb from '../../../../../../github.com/metaprov/modeldapi/pkg/apis/data/v1alpha1/generated_pb';
import * as github_com_metaprov_modeldapi_services_common_v1_common_pb from '../../../../../../github.com/metaprov/modeldapi/services/common/v1/common_pb';


export class ListDataProductVersionsRequest extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): ListDataProductVersionsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListDataProductVersionsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDataProductVersionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListDataProductVersionsRequest): ListDataProductVersionsRequest.AsObject;
  static serializeBinaryToWriter(message: ListDataProductVersionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDataProductVersionsRequest;
  static deserializeBinaryFromReader(message: ListDataProductVersionsRequest, reader: jspb.BinaryReader): ListDataProductVersionsRequest;
}

export namespace ListDataProductVersionsRequest {
  export type AsObject = {
    namespace: string,
    labelsMap: Array<[string, string]>,
  }
}

export class ListDataProductVersionsResponse extends jspb.Message {
  getItems(): github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersionList | undefined;
  setItems(value?: github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersionList): ListDataProductVersionsResponse;
  hasItems(): boolean;
  clearItems(): ListDataProductVersionsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDataProductVersionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListDataProductVersionsResponse): ListDataProductVersionsResponse.AsObject;
  static serializeBinaryToWriter(message: ListDataProductVersionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDataProductVersionsResponse;
  static deserializeBinaryFromReader(message: ListDataProductVersionsResponse, reader: jspb.BinaryReader): ListDataProductVersionsResponse;
}

export namespace ListDataProductVersionsResponse {
  export type AsObject = {
    items?: github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersionList.AsObject,
  }
}

export class CreateDataProductVersionRequest extends jspb.Message {
  getItem(): github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersion | undefined;
  setItem(value?: github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersion): CreateDataProductVersionRequest;
  hasItem(): boolean;
  clearItem(): CreateDataProductVersionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateDataProductVersionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateDataProductVersionRequest): CreateDataProductVersionRequest.AsObject;
  static serializeBinaryToWriter(message: CreateDataProductVersionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateDataProductVersionRequest;
  static deserializeBinaryFromReader(message: CreateDataProductVersionRequest, reader: jspb.BinaryReader): CreateDataProductVersionRequest;
}

export namespace CreateDataProductVersionRequest {
  export type AsObject = {
    item?: github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersion.AsObject,
  }
}

export class CreateDataProductVersionResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateDataProductVersionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateDataProductVersionResponse): CreateDataProductVersionResponse.AsObject;
  static serializeBinaryToWriter(message: CreateDataProductVersionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateDataProductVersionResponse;
  static deserializeBinaryFromReader(message: CreateDataProductVersionResponse, reader: jspb.BinaryReader): CreateDataProductVersionResponse;
}

export namespace CreateDataProductVersionResponse {
  export type AsObject = {
  }
}

export class UpdateDataProductVersionRequest extends jspb.Message {
  getItem(): github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersion | undefined;
  setItem(value?: github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersion): UpdateDataProductVersionRequest;
  hasItem(): boolean;
  clearItem(): UpdateDataProductVersionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateDataProductVersionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateDataProductVersionRequest): UpdateDataProductVersionRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateDataProductVersionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateDataProductVersionRequest;
  static deserializeBinaryFromReader(message: UpdateDataProductVersionRequest, reader: jspb.BinaryReader): UpdateDataProductVersionRequest;
}

export namespace UpdateDataProductVersionRequest {
  export type AsObject = {
    item?: github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersion.AsObject,
  }
}

export class UpdateDataProductVersionResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateDataProductVersionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateDataProductVersionResponse): UpdateDataProductVersionResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateDataProductVersionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateDataProductVersionResponse;
  static deserializeBinaryFromReader(message: UpdateDataProductVersionResponse, reader: jspb.BinaryReader): UpdateDataProductVersionResponse;
}

export namespace UpdateDataProductVersionResponse {
  export type AsObject = {
  }
}

export class GetDataProductVersionRequest extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): GetDataProductVersionRequest;

  getName(): string;
  setName(value: string): GetDataProductVersionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDataProductVersionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDataProductVersionRequest): GetDataProductVersionRequest.AsObject;
  static serializeBinaryToWriter(message: GetDataProductVersionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDataProductVersionRequest;
  static deserializeBinaryFromReader(message: GetDataProductVersionRequest, reader: jspb.BinaryReader): GetDataProductVersionRequest;
}

export namespace GetDataProductVersionRequest {
  export type AsObject = {
    namespace: string,
    name: string,
  }
}

export class GetDataProductVersionResponse extends jspb.Message {
  getItem(): github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersion | undefined;
  setItem(value?: github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersion): GetDataProductVersionResponse;
  hasItem(): boolean;
  clearItem(): GetDataProductVersionResponse;

  getYaml(): string;
  setYaml(value: string): GetDataProductVersionResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDataProductVersionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDataProductVersionResponse): GetDataProductVersionResponse.AsObject;
  static serializeBinaryToWriter(message: GetDataProductVersionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDataProductVersionResponse;
  static deserializeBinaryFromReader(message: GetDataProductVersionResponse, reader: jspb.BinaryReader): GetDataProductVersionResponse;
}

export namespace GetDataProductVersionResponse {
  export type AsObject = {
    item?: github_com_metaprov_modeldapi_pkg_apis_data_v1alpha1_generated_pb.DataProductVersion.AsObject,
    yaml: string,
  }
}

export class DeleteDataProductVersionRequest extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): DeleteDataProductVersionRequest;

  getName(): string;
  setName(value: string): DeleteDataProductVersionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteDataProductVersionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteDataProductVersionRequest): DeleteDataProductVersionRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteDataProductVersionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteDataProductVersionRequest;
  static deserializeBinaryFromReader(message: DeleteDataProductVersionRequest, reader: jspb.BinaryReader): DeleteDataProductVersionRequest;
}

export namespace DeleteDataProductVersionRequest {
  export type AsObject = {
    namespace: string,
    name: string,
  }
}

export class DeleteDataProductVersionResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteDataProductVersionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteDataProductVersionResponse): DeleteDataProductVersionResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteDataProductVersionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteDataProductVersionResponse;
  static deserializeBinaryFromReader(message: DeleteDataProductVersionResponse, reader: jspb.BinaryReader): DeleteDataProductVersionResponse;
}

export namespace DeleteDataProductVersionResponse {
  export type AsObject = {
  }
}

