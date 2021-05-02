# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/metaprov/modeldapi/services/predictor/v1/predictor.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from github.com.metaprov.modeldapi.pkg.apis.inference.v1alpha1 import generated_pb2 as github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='github.com/metaprov/modeldapi/services/predictor/v1/predictor.proto',
  package='github.com.metaprov.modeldapi.services.predictor.v1',
  syntax='proto3',
  serialized_options=b'Z3github.com/metaprov/modeldapi/services/predictor/v1',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\nCgithub.com/metaprov/modeldapi/services/predictor/v1/predictor.proto\x12\x33github.com.metaprov.modeldapi.services.predictor.v1\x1a\x1cgoogle/api/annotations.proto\x1aIgithub.com/metaprov/modeldapi/pkg/apis/inference/v1alpha1/generated.proto\"\xc1\x01\n\x15ListPredictorsRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x66\n\x06labels\x18\x02 \x03(\x0b\x32V.github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsRequest.LabelsEntry\x1a-\n\x0bLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"q\n\x16ListPredictorsResponse\x12W\n\x05items\x18\x01 \x01(\x0b\x32H.github.com.metaprov.modeldapi.pkg.apis.inference.v1alpha1.PredictorList\"\x19\n\x17\x43reatePredictorResponse\"l\n\x16\x43reatePredictorRequest\x12R\n\x04item\x18\x01 \x01(\x0b\x32\x44.github.com.metaprov.modeldapi.pkg.apis.inference.v1alpha1.Predictor\"l\n\x16UpdatePredictorRequest\x12R\n\x04item\x18\x01 \x01(\x0b\x32\x44.github.com.metaprov.modeldapi.pkg.apis.inference.v1alpha1.Predictor\"\x19\n\x17UpdatePredictorResponse\"6\n\x13GetPredictorRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"x\n\x14GetPredictorResponse\x12R\n\x04item\x18\x01 \x01(\x0b\x32\x44.github.com.metaprov.modeldapi.pkg.apis.inference.v1alpha1.Predictor\x12\x0c\n\x04yaml\x18\x02 \x01(\t\"9\n\x16\x44\x65letePredictorRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"\x19\n\x17\x44\x65letePredictorResponse\"2\n\x0fRollbackRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"\x12\n\x10RollbackResponse\"T\n\x11PredictOneRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06\x66ields\x18\x03 \x01(\t\x12\x0e\n\x06values\x18\x04 \x01(\t\"\xe4\x01\n\x12PredictOneResponse\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05label\x18\x03 \x01(\t\x12\r\n\x05score\x18\x04 \x01(\x02\x12\x61\n\x05proba\x18\x05 \x03(\x0b\x32R.github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse.ProbaEntry\x1a,\n\nProbaEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\x02:\x02\x38\x01\x32\xb9\x0b\n\x10PredictorService\x12\xc1\x01\n\x0eListPredictors\x12J.github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsRequest\x1aK.github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/v1/predictors\x12\xc7\x01\n\x0f\x43reatePredictor\x12K.github.com.metaprov.modeldapi.services.predictor.v1.CreatePredictorRequest\x1aL.github.com.metaprov.modeldapi.services.predictor.v1.CreatePredictorResponse\"\x19\x82\xd3\xe4\x93\x02\x13\"\x0e/v1/predictors:\x01*\x12\xc2\x01\n\x0cGetPredictor\x12H.github.com.metaprov.modeldapi.services.predictor.v1.GetPredictorRequest\x1aI.github.com.metaprov.modeldapi.services.predictor.v1.GetPredictorResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/v1/predictors/{name}\x12\xe1\x01\n\x0fUpdatePredictor\x12K.github.com.metaprov.modeldapi.services.predictor.v1.UpdatePredictorRequest\x1aL.github.com.metaprov.modeldapi.services.predictor.v1.UpdatePredictorResponse\"3\x82\xd3\xe4\x93\x02-\x1a(/v1/predictors/{predictor.metadata.name}:\x01*\x12\xd5\x01\n\x08RollBack\x12\x44.github.com.metaprov.modeldapi.services.predictor.v1.RollbackRequest\x1a\x45.github.com.metaprov.modeldapi.services.predictor.v1.RollbackResponse\"<\x82\xd3\xe4\x93\x02\x36\"1/v1/predictors/{predictor.metadata.name}:rollback:\x01*\x12\xcb\x01\n\x0f\x44\x65letePredictor\x12K.github.com.metaprov.modeldapi.services.predictor.v1.DeletePredictorRequest\x1aL.github.com.metaprov.modeldapi.services.predictor.v1.DeletePredictorResponse\"\x1d\x82\xd3\xe4\x93\x02\x17*\x15/v1/predictors/{name}\x12\xc7\x01\n\nPredictOne\x12\x46.github.com.metaprov.modeldapi.services.predictor.v1.PredictOneRequest\x1aG.github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse\"(\x82\xd3\xe4\x93\x02\"\" /v1/predictors/{name}:predictoneB5Z3github.com/metaprov/modeldapi/services/predictor/v1b\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2.DESCRIPTOR,])




_LISTPREDICTORSREQUEST_LABELSENTRY = _descriptor.Descriptor(
  name='LabelsEntry',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsRequest.LabelsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsRequest.LabelsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsRequest.LabelsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=378,
  serialized_end=423,
)

_LISTPREDICTORSREQUEST = _descriptor.Descriptor(
  name='ListPredictorsRequest',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsRequest.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='labels', full_name='github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsRequest.labels', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_LISTPREDICTORSREQUEST_LABELSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=230,
  serialized_end=423,
)


_LISTPREDICTORSRESPONSE = _descriptor.Descriptor(
  name='ListPredictorsResponse',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='items', full_name='github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsResponse.items', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=425,
  serialized_end=538,
)


_CREATEPREDICTORRESPONSE = _descriptor.Descriptor(
  name='CreatePredictorResponse',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.CreatePredictorResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=540,
  serialized_end=565,
)


_CREATEPREDICTORREQUEST = _descriptor.Descriptor(
  name='CreatePredictorRequest',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.CreatePredictorRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='item', full_name='github.com.metaprov.modeldapi.services.predictor.v1.CreatePredictorRequest.item', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=567,
  serialized_end=675,
)


_UPDATEPREDICTORREQUEST = _descriptor.Descriptor(
  name='UpdatePredictorRequest',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.UpdatePredictorRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='item', full_name='github.com.metaprov.modeldapi.services.predictor.v1.UpdatePredictorRequest.item', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=677,
  serialized_end=785,
)


_UPDATEPREDICTORRESPONSE = _descriptor.Descriptor(
  name='UpdatePredictorResponse',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.UpdatePredictorResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=787,
  serialized_end=812,
)


_GETPREDICTORREQUEST = _descriptor.Descriptor(
  name='GetPredictorRequest',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.GetPredictorRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeldapi.services.predictor.v1.GetPredictorRequest.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='github.com.metaprov.modeldapi.services.predictor.v1.GetPredictorRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=814,
  serialized_end=868,
)


_GETPREDICTORRESPONSE = _descriptor.Descriptor(
  name='GetPredictorResponse',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.GetPredictorResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='item', full_name='github.com.metaprov.modeldapi.services.predictor.v1.GetPredictorResponse.item', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='yaml', full_name='github.com.metaprov.modeldapi.services.predictor.v1.GetPredictorResponse.yaml', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=870,
  serialized_end=990,
)


_DELETEPREDICTORREQUEST = _descriptor.Descriptor(
  name='DeletePredictorRequest',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.DeletePredictorRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeldapi.services.predictor.v1.DeletePredictorRequest.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='github.com.metaprov.modeldapi.services.predictor.v1.DeletePredictorRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=992,
  serialized_end=1049,
)


_DELETEPREDICTORRESPONSE = _descriptor.Descriptor(
  name='DeletePredictorResponse',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.DeletePredictorResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1051,
  serialized_end=1076,
)


_ROLLBACKREQUEST = _descriptor.Descriptor(
  name='RollbackRequest',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.RollbackRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeldapi.services.predictor.v1.RollbackRequest.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='github.com.metaprov.modeldapi.services.predictor.v1.RollbackRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1078,
  serialized_end=1128,
)


_ROLLBACKRESPONSE = _descriptor.Descriptor(
  name='RollbackResponse',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.RollbackResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1130,
  serialized_end=1148,
)


_PREDICTONEREQUEST = _descriptor.Descriptor(
  name='PredictOneRequest',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneRequest.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='fields', full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneRequest.fields', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='values', full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneRequest.values', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1150,
  serialized_end=1234,
)


_PREDICTONERESPONSE_PROBAENTRY = _descriptor.Descriptor(
  name='ProbaEntry',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse.ProbaEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse.ProbaEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse.ProbaEntry.value', index=1,
      number=2, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1421,
  serialized_end=1465,
)

_PREDICTONERESPONSE = _descriptor.Descriptor(
  name='PredictOneResponse',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='label', full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse.label', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='score', full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse.score', index=3,
      number=4, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='proba', full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse.proba', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_PREDICTONERESPONSE_PROBAENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1237,
  serialized_end=1465,
)

_LISTPREDICTORSREQUEST_LABELSENTRY.containing_type = _LISTPREDICTORSREQUEST
_LISTPREDICTORSREQUEST.fields_by_name['labels'].message_type = _LISTPREDICTORSREQUEST_LABELSENTRY
_LISTPREDICTORSRESPONSE.fields_by_name['items'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2._PREDICTORLIST
_CREATEPREDICTORREQUEST.fields_by_name['item'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2._PREDICTOR
_UPDATEPREDICTORREQUEST.fields_by_name['item'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2._PREDICTOR
_GETPREDICTORRESPONSE.fields_by_name['item'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2._PREDICTOR
_PREDICTONERESPONSE_PROBAENTRY.containing_type = _PREDICTONERESPONSE
_PREDICTONERESPONSE.fields_by_name['proba'].message_type = _PREDICTONERESPONSE_PROBAENTRY
DESCRIPTOR.message_types_by_name['ListPredictorsRequest'] = _LISTPREDICTORSREQUEST
DESCRIPTOR.message_types_by_name['ListPredictorsResponse'] = _LISTPREDICTORSRESPONSE
DESCRIPTOR.message_types_by_name['CreatePredictorResponse'] = _CREATEPREDICTORRESPONSE
DESCRIPTOR.message_types_by_name['CreatePredictorRequest'] = _CREATEPREDICTORREQUEST
DESCRIPTOR.message_types_by_name['UpdatePredictorRequest'] = _UPDATEPREDICTORREQUEST
DESCRIPTOR.message_types_by_name['UpdatePredictorResponse'] = _UPDATEPREDICTORRESPONSE
DESCRIPTOR.message_types_by_name['GetPredictorRequest'] = _GETPREDICTORREQUEST
DESCRIPTOR.message_types_by_name['GetPredictorResponse'] = _GETPREDICTORRESPONSE
DESCRIPTOR.message_types_by_name['DeletePredictorRequest'] = _DELETEPREDICTORREQUEST
DESCRIPTOR.message_types_by_name['DeletePredictorResponse'] = _DELETEPREDICTORRESPONSE
DESCRIPTOR.message_types_by_name['RollbackRequest'] = _ROLLBACKREQUEST
DESCRIPTOR.message_types_by_name['RollbackResponse'] = _ROLLBACKRESPONSE
DESCRIPTOR.message_types_by_name['PredictOneRequest'] = _PREDICTONEREQUEST
DESCRIPTOR.message_types_by_name['PredictOneResponse'] = _PREDICTONERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListPredictorsRequest = _reflection.GeneratedProtocolMessageType('ListPredictorsRequest', (_message.Message,), {

  'LabelsEntry' : _reflection.GeneratedProtocolMessageType('LabelsEntry', (_message.Message,), {
    'DESCRIPTOR' : _LISTPREDICTORSREQUEST_LABELSENTRY,
    '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
    # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsRequest.LabelsEntry)
    })
  ,
  'DESCRIPTOR' : _LISTPREDICTORSREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsRequest)
  })
_sym_db.RegisterMessage(ListPredictorsRequest)
_sym_db.RegisterMessage(ListPredictorsRequest.LabelsEntry)

ListPredictorsResponse = _reflection.GeneratedProtocolMessageType('ListPredictorsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTPREDICTORSRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.ListPredictorsResponse)
  })
_sym_db.RegisterMessage(ListPredictorsResponse)

CreatePredictorResponse = _reflection.GeneratedProtocolMessageType('CreatePredictorResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPREDICTORRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.CreatePredictorResponse)
  })
_sym_db.RegisterMessage(CreatePredictorResponse)

CreatePredictorRequest = _reflection.GeneratedProtocolMessageType('CreatePredictorRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPREDICTORREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.CreatePredictorRequest)
  })
_sym_db.RegisterMessage(CreatePredictorRequest)

UpdatePredictorRequest = _reflection.GeneratedProtocolMessageType('UpdatePredictorRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEPREDICTORREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.UpdatePredictorRequest)
  })
_sym_db.RegisterMessage(UpdatePredictorRequest)

UpdatePredictorResponse = _reflection.GeneratedProtocolMessageType('UpdatePredictorResponse', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEPREDICTORRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.UpdatePredictorResponse)
  })
_sym_db.RegisterMessage(UpdatePredictorResponse)

GetPredictorRequest = _reflection.GeneratedProtocolMessageType('GetPredictorRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETPREDICTORREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.GetPredictorRequest)
  })
_sym_db.RegisterMessage(GetPredictorRequest)

GetPredictorResponse = _reflection.GeneratedProtocolMessageType('GetPredictorResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETPREDICTORRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.GetPredictorResponse)
  })
_sym_db.RegisterMessage(GetPredictorResponse)

DeletePredictorRequest = _reflection.GeneratedProtocolMessageType('DeletePredictorRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEPREDICTORREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.DeletePredictorRequest)
  })
_sym_db.RegisterMessage(DeletePredictorRequest)

DeletePredictorResponse = _reflection.GeneratedProtocolMessageType('DeletePredictorResponse', (_message.Message,), {
  'DESCRIPTOR' : _DELETEPREDICTORRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.DeletePredictorResponse)
  })
_sym_db.RegisterMessage(DeletePredictorResponse)

RollbackRequest = _reflection.GeneratedProtocolMessageType('RollbackRequest', (_message.Message,), {
  'DESCRIPTOR' : _ROLLBACKREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.RollbackRequest)
  })
_sym_db.RegisterMessage(RollbackRequest)

RollbackResponse = _reflection.GeneratedProtocolMessageType('RollbackResponse', (_message.Message,), {
  'DESCRIPTOR' : _ROLLBACKRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.RollbackResponse)
  })
_sym_db.RegisterMessage(RollbackResponse)

PredictOneRequest = _reflection.GeneratedProtocolMessageType('PredictOneRequest', (_message.Message,), {
  'DESCRIPTOR' : _PREDICTONEREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.PredictOneRequest)
  })
_sym_db.RegisterMessage(PredictOneRequest)

PredictOneResponse = _reflection.GeneratedProtocolMessageType('PredictOneResponse', (_message.Message,), {

  'ProbaEntry' : _reflection.GeneratedProtocolMessageType('ProbaEntry', (_message.Message,), {
    'DESCRIPTOR' : _PREDICTONERESPONSE_PROBAENTRY,
    '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
    # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse.ProbaEntry)
    })
  ,
  'DESCRIPTOR' : _PREDICTONERESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.predictor.v1.predictor_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.predictor.v1.PredictOneResponse)
  })
_sym_db.RegisterMessage(PredictOneResponse)
_sym_db.RegisterMessage(PredictOneResponse.ProbaEntry)


DESCRIPTOR._options = None
_LISTPREDICTORSREQUEST_LABELSENTRY._options = None
_PREDICTONERESPONSE_PROBAENTRY._options = None

_PREDICTORSERVICE = _descriptor.ServiceDescriptor(
  name='PredictorService',
  full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictorService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1468,
  serialized_end=2933,
  methods=[
  _descriptor.MethodDescriptor(
    name='ListPredictors',
    full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictorService.ListPredictors',
    index=0,
    containing_service=None,
    input_type=_LISTPREDICTORSREQUEST,
    output_type=_LISTPREDICTORSRESPONSE,
    serialized_options=b'\202\323\344\223\002\020\022\016/v1/predictors',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreatePredictor',
    full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictorService.CreatePredictor',
    index=1,
    containing_service=None,
    input_type=_CREATEPREDICTORREQUEST,
    output_type=_CREATEPREDICTORRESPONSE,
    serialized_options=b'\202\323\344\223\002\023\"\016/v1/predictors:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetPredictor',
    full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictorService.GetPredictor',
    index=2,
    containing_service=None,
    input_type=_GETPREDICTORREQUEST,
    output_type=_GETPREDICTORRESPONSE,
    serialized_options=b'\202\323\344\223\002\027\022\025/v1/predictors/{name}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='UpdatePredictor',
    full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictorService.UpdatePredictor',
    index=3,
    containing_service=None,
    input_type=_UPDATEPREDICTORREQUEST,
    output_type=_UPDATEPREDICTORRESPONSE,
    serialized_options=b'\202\323\344\223\002-\032(/v1/predictors/{predictor.metadata.name}:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='RollBack',
    full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictorService.RollBack',
    index=4,
    containing_service=None,
    input_type=_ROLLBACKREQUEST,
    output_type=_ROLLBACKRESPONSE,
    serialized_options=b'\202\323\344\223\0026\"1/v1/predictors/{predictor.metadata.name}:rollback:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DeletePredictor',
    full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictorService.DeletePredictor',
    index=5,
    containing_service=None,
    input_type=_DELETEPREDICTORREQUEST,
    output_type=_DELETEPREDICTORRESPONSE,
    serialized_options=b'\202\323\344\223\002\027*\025/v1/predictors/{name}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='PredictOne',
    full_name='github.com.metaprov.modeldapi.services.predictor.v1.PredictorService.PredictOne',
    index=6,
    containing_service=None,
    input_type=_PREDICTONEREQUEST,
    output_type=_PREDICTONERESPONSE,
    serialized_options=b'\202\323\344\223\002\"\" /v1/predictors/{name}:predictone',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_PREDICTORSERVICE)

DESCRIPTOR.services_by_name['PredictorService'] = _PREDICTORSERVICE

# @@protoc_insertion_point(module_scope)
