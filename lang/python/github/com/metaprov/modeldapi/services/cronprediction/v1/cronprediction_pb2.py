# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/metaprov/modeldapi/services/cronprediction/v1/cronprediction.proto
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
  name='github.com/metaprov/modeldapi/services/cronprediction/v1/cronprediction.proto',
  package='github.com.metaprov.modeldapi.services.cronprediction.v1',
  syntax='proto3',
  serialized_options=b'Z8github.com/metaprov/modeldapi/services/cronprediction/v1',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\nMgithub.com/metaprov/modeldapi/services/cronprediction/v1/cronprediction.proto\x12\x38github.com.metaprov.modeldapi.services.cronprediction.v1\x1a\x1cgoogle/api/annotations.proto\x1aIgithub.com/metaprov/modeldapi/pkg/apis/inference/v1alpha1/generated.proto\"\xde\x01\n\x1aListCronPredictionsRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12p\n\x06labels\x18\x03 \x03(\x0b\x32`.github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsRequest.LabelsEntry\x1a-\n\x0bLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"{\n\x1bListCronPredictionsResponse\x12\\\n\x05items\x18\x01 \x01(\x0b\x32M.github.com.metaprov.modeldapi.pkg.apis.inference.v1alpha1.CronPredictionList\"\x1e\n\x1c\x43reateCronPredictionResponse\"v\n\x1b\x43reateCronPredictionRequest\x12W\n\x04item\x18\x01 \x01(\x0b\x32I.github.com.metaprov.modeldapi.pkg.apis.inference.v1alpha1.CronPrediction\"v\n\x1bUpdateCronPredictionRequest\x12W\n\x04item\x18\x01 \x01(\x0b\x32I.github.com.metaprov.modeldapi.pkg.apis.inference.v1alpha1.CronPrediction\"\x1e\n\x1cUpdateCronPredictionResponse\";\n\x18GetCronPredictionRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"\x82\x01\n\x19GetCronPredictionResponse\x12W\n\x04item\x18\x01 \x01(\x0b\x32I.github.com.metaprov.modeldapi.pkg.apis.inference.v1alpha1.CronPrediction\x12\x0c\n\x04yaml\x18\x02 \x01(\t\">\n\x1b\x44\x65leteCronPredictionRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"\x1e\n\x1c\x44\x65leteCronPredictionResponse\"s\n\x18RunCronPredictionRequest\x12W\n\x04item\x18\x01 \x01(\x0b\x32I.github.com.metaprov.modeldapi.pkg.apis.inference.v1alpha1.CronPrediction\"\x1b\n\x19RunCronPredictionResponse2\x9a\x0b\n\x15\x43ronPredictionService\x12\xdf\x01\n\x13ListCronPredictions\x12T.github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsRequest\x1aU.github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/v1/cronpredictions\x12\xe5\x01\n\x14\x43reateCronPrediction\x12U.github.com.metaprov.modeldapi.services.cronprediction.v1.CreateCronPredictionRequest\x1aV.github.com.metaprov.modeldapi.services.cronprediction.v1.CreateCronPredictionResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/v1/cronpredictions:\x01*\x12\xe0\x01\n\x11GetCronPrediction\x12R.github.com.metaprov.modeldapi.services.cronprediction.v1.GetCronPredictionRequest\x1aS.github.com.metaprov.modeldapi.services.cronprediction.v1.GetCronPredictionResponse\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/cronpredictions/{name}\x12\x84\x02\n\x14UpdateCronPrediction\x12U.github.com.metaprov.modeldapi.services.cronprediction.v1.UpdateCronPredictionRequest\x1aV.github.com.metaprov.modeldapi.services.cronprediction.v1.UpdateCronPredictionResponse\"=\x82\xd3\xe4\x93\x02\x37\x1a\x32/v1/cronpredictions/{cronprediction.metadata.name}:\x01*\x12\xe9\x01\n\x14\x44\x65leteCronPrediction\x12U.github.com.metaprov.modeldapi.services.cronprediction.v1.DeleteCronPredictionRequest\x1aV.github.com.metaprov.modeldapi.services.cronprediction.v1.DeleteCronPredictionResponse\"\"\x82\xd3\xe4\x93\x02\x1c*\x1a/v1/cronpredictions/{name}\x12\xe0\x01\n\rRunPrediction\x12R.github.com.metaprov.modeldapi.services.cronprediction.v1.RunCronPredictionRequest\x1aS.github.com.metaprov.modeldapi.services.cronprediction.v1.RunCronPredictionResponse\"&\x82\xd3\xe4\x93\x02 \"\x1e/v1/cronpredictions/{name}:runB:Z8github.com/metaprov/modeldapi/services/cronprediction/v1b\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2.DESCRIPTOR,])




_LISTCRONPREDICTIONSREQUEST_LABELSENTRY = _descriptor.Descriptor(
  name='LabelsEntry',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsRequest.LabelsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsRequest.LabelsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsRequest.LabelsEntry.value', index=1,
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
  serialized_start=422,
  serialized_end=467,
)

_LISTCRONPREDICTIONSREQUEST = _descriptor.Descriptor(
  name='ListCronPredictionsRequest',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsRequest.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='labels', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsRequest.labels', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_LISTCRONPREDICTIONSREQUEST_LABELSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=245,
  serialized_end=467,
)


_LISTCRONPREDICTIONSRESPONSE = _descriptor.Descriptor(
  name='ListCronPredictionsResponse',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='items', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsResponse.items', index=0,
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
  serialized_start=469,
  serialized_end=592,
)


_CREATECRONPREDICTIONRESPONSE = _descriptor.Descriptor(
  name='CreateCronPredictionResponse',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.CreateCronPredictionResponse',
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
  serialized_start=594,
  serialized_end=624,
)


_CREATECRONPREDICTIONREQUEST = _descriptor.Descriptor(
  name='CreateCronPredictionRequest',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.CreateCronPredictionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='item', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.CreateCronPredictionRequest.item', index=0,
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
  serialized_start=626,
  serialized_end=744,
)


_UPDATECRONPREDICTIONREQUEST = _descriptor.Descriptor(
  name='UpdateCronPredictionRequest',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.UpdateCronPredictionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='item', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.UpdateCronPredictionRequest.item', index=0,
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
  serialized_start=746,
  serialized_end=864,
)


_UPDATECRONPREDICTIONRESPONSE = _descriptor.Descriptor(
  name='UpdateCronPredictionResponse',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.UpdateCronPredictionResponse',
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
  serialized_start=866,
  serialized_end=896,
)


_GETCRONPREDICTIONREQUEST = _descriptor.Descriptor(
  name='GetCronPredictionRequest',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.GetCronPredictionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.GetCronPredictionRequest.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.GetCronPredictionRequest.name', index=1,
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
  serialized_start=898,
  serialized_end=957,
)


_GETCRONPREDICTIONRESPONSE = _descriptor.Descriptor(
  name='GetCronPredictionResponse',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.GetCronPredictionResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='item', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.GetCronPredictionResponse.item', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='yaml', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.GetCronPredictionResponse.yaml', index=1,
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
  serialized_start=960,
  serialized_end=1090,
)


_DELETECRONPREDICTIONREQUEST = _descriptor.Descriptor(
  name='DeleteCronPredictionRequest',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.DeleteCronPredictionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.DeleteCronPredictionRequest.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.DeleteCronPredictionRequest.name', index=1,
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
  serialized_start=1092,
  serialized_end=1154,
)


_DELETECRONPREDICTIONRESPONSE = _descriptor.Descriptor(
  name='DeleteCronPredictionResponse',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.DeleteCronPredictionResponse',
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
  serialized_start=1156,
  serialized_end=1186,
)


_RUNCRONPREDICTIONREQUEST = _descriptor.Descriptor(
  name='RunCronPredictionRequest',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.RunCronPredictionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='item', full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.RunCronPredictionRequest.item', index=0,
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
  serialized_start=1188,
  serialized_end=1303,
)


_RUNCRONPREDICTIONRESPONSE = _descriptor.Descriptor(
  name='RunCronPredictionResponse',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.RunCronPredictionResponse',
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
  serialized_start=1305,
  serialized_end=1332,
)

_LISTCRONPREDICTIONSREQUEST_LABELSENTRY.containing_type = _LISTCRONPREDICTIONSREQUEST
_LISTCRONPREDICTIONSREQUEST.fields_by_name['labels'].message_type = _LISTCRONPREDICTIONSREQUEST_LABELSENTRY
_LISTCRONPREDICTIONSRESPONSE.fields_by_name['items'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2._CRONPREDICTIONLIST
_CREATECRONPREDICTIONREQUEST.fields_by_name['item'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2._CRONPREDICTION
_UPDATECRONPREDICTIONREQUEST.fields_by_name['item'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2._CRONPREDICTION
_GETCRONPREDICTIONRESPONSE.fields_by_name['item'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2._CRONPREDICTION
_RUNCRONPREDICTIONREQUEST.fields_by_name['item'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_inference_dot_v1alpha1_dot_generated__pb2._CRONPREDICTION
DESCRIPTOR.message_types_by_name['ListCronPredictionsRequest'] = _LISTCRONPREDICTIONSREQUEST
DESCRIPTOR.message_types_by_name['ListCronPredictionsResponse'] = _LISTCRONPREDICTIONSRESPONSE
DESCRIPTOR.message_types_by_name['CreateCronPredictionResponse'] = _CREATECRONPREDICTIONRESPONSE
DESCRIPTOR.message_types_by_name['CreateCronPredictionRequest'] = _CREATECRONPREDICTIONREQUEST
DESCRIPTOR.message_types_by_name['UpdateCronPredictionRequest'] = _UPDATECRONPREDICTIONREQUEST
DESCRIPTOR.message_types_by_name['UpdateCronPredictionResponse'] = _UPDATECRONPREDICTIONRESPONSE
DESCRIPTOR.message_types_by_name['GetCronPredictionRequest'] = _GETCRONPREDICTIONREQUEST
DESCRIPTOR.message_types_by_name['GetCronPredictionResponse'] = _GETCRONPREDICTIONRESPONSE
DESCRIPTOR.message_types_by_name['DeleteCronPredictionRequest'] = _DELETECRONPREDICTIONREQUEST
DESCRIPTOR.message_types_by_name['DeleteCronPredictionResponse'] = _DELETECRONPREDICTIONRESPONSE
DESCRIPTOR.message_types_by_name['RunCronPredictionRequest'] = _RUNCRONPREDICTIONREQUEST
DESCRIPTOR.message_types_by_name['RunCronPredictionResponse'] = _RUNCRONPREDICTIONRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListCronPredictionsRequest = _reflection.GeneratedProtocolMessageType('ListCronPredictionsRequest', (_message.Message,), {

  'LabelsEntry' : _reflection.GeneratedProtocolMessageType('LabelsEntry', (_message.Message,), {
    'DESCRIPTOR' : _LISTCRONPREDICTIONSREQUEST_LABELSENTRY,
    '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
    # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsRequest.LabelsEntry)
    })
  ,
  'DESCRIPTOR' : _LISTCRONPREDICTIONSREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsRequest)
  })
_sym_db.RegisterMessage(ListCronPredictionsRequest)
_sym_db.RegisterMessage(ListCronPredictionsRequest.LabelsEntry)

ListCronPredictionsResponse = _reflection.GeneratedProtocolMessageType('ListCronPredictionsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTCRONPREDICTIONSRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.ListCronPredictionsResponse)
  })
_sym_db.RegisterMessage(ListCronPredictionsResponse)

CreateCronPredictionResponse = _reflection.GeneratedProtocolMessageType('CreateCronPredictionResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATECRONPREDICTIONRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.CreateCronPredictionResponse)
  })
_sym_db.RegisterMessage(CreateCronPredictionResponse)

CreateCronPredictionRequest = _reflection.GeneratedProtocolMessageType('CreateCronPredictionRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATECRONPREDICTIONREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.CreateCronPredictionRequest)
  })
_sym_db.RegisterMessage(CreateCronPredictionRequest)

UpdateCronPredictionRequest = _reflection.GeneratedProtocolMessageType('UpdateCronPredictionRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATECRONPREDICTIONREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.UpdateCronPredictionRequest)
  })
_sym_db.RegisterMessage(UpdateCronPredictionRequest)

UpdateCronPredictionResponse = _reflection.GeneratedProtocolMessageType('UpdateCronPredictionResponse', (_message.Message,), {
  'DESCRIPTOR' : _UPDATECRONPREDICTIONRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.UpdateCronPredictionResponse)
  })
_sym_db.RegisterMessage(UpdateCronPredictionResponse)

GetCronPredictionRequest = _reflection.GeneratedProtocolMessageType('GetCronPredictionRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETCRONPREDICTIONREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.GetCronPredictionRequest)
  })
_sym_db.RegisterMessage(GetCronPredictionRequest)

GetCronPredictionResponse = _reflection.GeneratedProtocolMessageType('GetCronPredictionResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETCRONPREDICTIONRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.GetCronPredictionResponse)
  })
_sym_db.RegisterMessage(GetCronPredictionResponse)

DeleteCronPredictionRequest = _reflection.GeneratedProtocolMessageType('DeleteCronPredictionRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETECRONPREDICTIONREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.DeleteCronPredictionRequest)
  })
_sym_db.RegisterMessage(DeleteCronPredictionRequest)

DeleteCronPredictionResponse = _reflection.GeneratedProtocolMessageType('DeleteCronPredictionResponse', (_message.Message,), {
  'DESCRIPTOR' : _DELETECRONPREDICTIONRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.DeleteCronPredictionResponse)
  })
_sym_db.RegisterMessage(DeleteCronPredictionResponse)

RunCronPredictionRequest = _reflection.GeneratedProtocolMessageType('RunCronPredictionRequest', (_message.Message,), {
  'DESCRIPTOR' : _RUNCRONPREDICTIONREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.RunCronPredictionRequest)
  })
_sym_db.RegisterMessage(RunCronPredictionRequest)

RunCronPredictionResponse = _reflection.GeneratedProtocolMessageType('RunCronPredictionResponse', (_message.Message,), {
  'DESCRIPTOR' : _RUNCRONPREDICTIONRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.cronprediction.v1.cronprediction_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeldapi.services.cronprediction.v1.RunCronPredictionResponse)
  })
_sym_db.RegisterMessage(RunCronPredictionResponse)


DESCRIPTOR._options = None
_LISTCRONPREDICTIONSREQUEST_LABELSENTRY._options = None

_CRONPREDICTIONSERVICE = _descriptor.ServiceDescriptor(
  name='CronPredictionService',
  full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1335,
  serialized_end=2769,
  methods=[
  _descriptor.MethodDescriptor(
    name='ListCronPredictions',
    full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService.ListCronPredictions',
    index=0,
    containing_service=None,
    input_type=_LISTCRONPREDICTIONSREQUEST,
    output_type=_LISTCRONPREDICTIONSRESPONSE,
    serialized_options=b'\202\323\344\223\002\025\022\023/v1/cronpredictions',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateCronPrediction',
    full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService.CreateCronPrediction',
    index=1,
    containing_service=None,
    input_type=_CREATECRONPREDICTIONREQUEST,
    output_type=_CREATECRONPREDICTIONRESPONSE,
    serialized_options=b'\202\323\344\223\002\030\"\023/v1/cronpredictions:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetCronPrediction',
    full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService.GetCronPrediction',
    index=2,
    containing_service=None,
    input_type=_GETCRONPREDICTIONREQUEST,
    output_type=_GETCRONPREDICTIONRESPONSE,
    serialized_options=b'\202\323\344\223\002\034\022\032/v1/cronpredictions/{name}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateCronPrediction',
    full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService.UpdateCronPrediction',
    index=3,
    containing_service=None,
    input_type=_UPDATECRONPREDICTIONREQUEST,
    output_type=_UPDATECRONPREDICTIONRESPONSE,
    serialized_options=b'\202\323\344\223\0027\0322/v1/cronpredictions/{cronprediction.metadata.name}:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteCronPrediction',
    full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService.DeleteCronPrediction',
    index=4,
    containing_service=None,
    input_type=_DELETECRONPREDICTIONREQUEST,
    output_type=_DELETECRONPREDICTIONRESPONSE,
    serialized_options=b'\202\323\344\223\002\034*\032/v1/cronpredictions/{name}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='RunPrediction',
    full_name='github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService.RunPrediction',
    index=5,
    containing_service=None,
    input_type=_RUNCRONPREDICTIONREQUEST,
    output_type=_RUNCRONPREDICTIONRESPONSE,
    serialized_options=b'\202\323\344\223\002 \"\036/v1/cronpredictions/{name}:run',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_CRONPREDICTIONSERVICE)

DESCRIPTOR.services_by_name['CronPredictionService'] = _CRONPREDICTIONSERVICE

# @@protoc_insertion_point(module_scope)
