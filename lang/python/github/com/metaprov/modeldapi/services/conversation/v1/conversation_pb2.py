# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/metaprov/modeldapi/services/conversation/v1/conversation.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from github.com.metaprov.modeldapi.pkg.apis.team.v1alpha1 import generated_pb2 as github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_team_dot_v1alpha1_dot_generated__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='github.com/metaprov/modeldapi/services/conversation/v1/conversation.proto',
  package='github.com.metaprov.modeld.services.conversation.v1',
  syntax='proto3',
  serialized_options=b'Z6github.com/metaprov/modeldapi/services/conversation/v1',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\nIgithub.com/metaprov/modeldapi/services/conversation/v1/conversation.proto\x12\x33github.com.metaprov.modeld.services.conversation.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x44github.com/metaprov/modeldapi/pkg/apis/team/v1alpha1/generated.proto\x1a\x1cgoogle/api/annotations.proto\"\xc7\x01\n\x11\x43onversationQuery\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x62\n\x06labels\x18\x03 \x03(\x0b\x32R.github.com.metaprov.modeld.services.conversation.v1.ConversationQuery.LabelsEntry\x1a-\n\x0bLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x16\n\x14\x43onversationResponse\"\xcf\x02\n\x19\x43onversationCreateRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12j\n\x06labels\x18\x03 \x03(\x0b\x32Z.github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest.LabelsEntry\x12T\n\x04spec\x18\x05 \x01(\x0b\x32\x46.github.com.metaprov.modeldapi.pkg.apis.team.v1alpha1.ConversationSpec\x12\x10\n\x08password\x18\x06 \x01(\t\x12\x0e\n\x06upsert\x18\x07 \x01(\x08\x1a-\n\x0bLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xad\x02\n\x19\x43onversationUpdateRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12j\n\x06labels\x18\x03 \x03(\x0b\x32Z.github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest.LabelsEntry\x12T\n\x04spec\x18\x05 \x01(\x0b\x32\x46.github.com.metaprov.modeldapi.pkg.apis.team.v1alpha1.ConversationSpec\x1a-\n\x0bLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"y\n\x17\x43onversationGetResponse\x12P\n\x04item\x18\x01 \x01(\x0b\x32\x42.github.com.metaprov.modeldapi.pkg.apis.team.v1alpha1.Conversation\x12\x0c\n\x04yaml\x18\x02 \x01(\t2\xf2\x06\n\x13\x43onversationService\x12\xbb\x01\n\x04List\x12\x46.github.com.metaprov.modeld.services.conversation.v1.ConversationQuery\x1a\x46.github.com.metaprov.modeldapi.pkg.apis.team.v1alpha1.ConversationList\"#\x82\xd3\xe4\x93\x02\x1d\x12\x1b/api/v1alpha1/conversations\x12\xa3\x01\n\x06\x43reate\x12N.github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest\x1a\x16.google.protobuf.Empty\"1\x82\xd3\xe4\x93\x02+\"\x1b/api/v1alpha1/conversations:\x0c\x63onversation\x12\xc7\x01\n\x03Get\x12\x46.github.com.metaprov.modeld.services.conversation.v1.ConversationQuery\x1aL.github.com.metaprov.modeld.services.conversation.v1.ConversationGetResponse\"*\x82\xd3\xe4\x93\x02$\x12\"/api/v1alpha1/conversations/{name}\x12\xc0\x01\n\x06Update\x12N.github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest\x1a\x16.google.protobuf.Empty\"N\x82\xd3\xe4\x93\x02H\x1a\x38/api/v1alpha1/conversations/{conversation.metadata.name}:\x0c\x63onversation\x12j\n\x06\x44\x65lete\x12\x46.github.com.metaprov.modeld.services.conversation.v1.ConversationQuery\x1a\x16.google.protobuf.Empty\"\x00\x42\x38Z6github.com/metaprov/modeldapi/services/conversation/v1b\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_team_dot_v1alpha1_dot_generated__pb2.DESCRIPTOR,google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




_CONVERSATIONQUERY_LABELSENTRY = _descriptor.Descriptor(
  name='LabelsEntry',
  full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationQuery.LabelsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationQuery.LabelsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationQuery.LabelsEntry.value', index=1,
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
  serialized_start=414,
  serialized_end=459,
)

_CONVERSATIONQUERY = _descriptor.Descriptor(
  name='ConversationQuery',
  full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationQuery.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationQuery.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='labels', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationQuery.labels', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_CONVERSATIONQUERY_LABELSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=260,
  serialized_end=459,
)


_CONVERSATIONRESPONSE = _descriptor.Descriptor(
  name='ConversationResponse',
  full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationResponse',
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
  serialized_start=461,
  serialized_end=483,
)


_CONVERSATIONCREATEREQUEST_LABELSENTRY = _descriptor.Descriptor(
  name='LabelsEntry',
  full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest.LabelsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest.LabelsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest.LabelsEntry.value', index=1,
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
  serialized_start=414,
  serialized_end=459,
)

_CONVERSATIONCREATEREQUEST = _descriptor.Descriptor(
  name='ConversationCreateRequest',
  full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='labels', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest.labels', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='spec', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest.spec', index=3,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='password', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest.password', index=4,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='upsert', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest.upsert', index=5,
      number=7, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_CONVERSATIONCREATEREQUEST_LABELSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=486,
  serialized_end=821,
)


_CONVERSATIONUPDATEREQUEST_LABELSENTRY = _descriptor.Descriptor(
  name='LabelsEntry',
  full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest.LabelsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest.LabelsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest.LabelsEntry.value', index=1,
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
  serialized_start=414,
  serialized_end=459,
)

_CONVERSATIONUPDATEREQUEST = _descriptor.Descriptor(
  name='ConversationUpdateRequest',
  full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='labels', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest.labels', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='spec', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest.spec', index=3,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_CONVERSATIONUPDATEREQUEST_LABELSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=824,
  serialized_end=1125,
)


_CONVERSATIONGETRESPONSE = _descriptor.Descriptor(
  name='ConversationGetResponse',
  full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationGetResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='item', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationGetResponse.item', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='yaml', full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationGetResponse.yaml', index=1,
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
  serialized_start=1127,
  serialized_end=1248,
)

_CONVERSATIONQUERY_LABELSENTRY.containing_type = _CONVERSATIONQUERY
_CONVERSATIONQUERY.fields_by_name['labels'].message_type = _CONVERSATIONQUERY_LABELSENTRY
_CONVERSATIONCREATEREQUEST_LABELSENTRY.containing_type = _CONVERSATIONCREATEREQUEST
_CONVERSATIONCREATEREQUEST.fields_by_name['labels'].message_type = _CONVERSATIONCREATEREQUEST_LABELSENTRY
_CONVERSATIONCREATEREQUEST.fields_by_name['spec'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_team_dot_v1alpha1_dot_generated__pb2._CONVERSATIONSPEC
_CONVERSATIONUPDATEREQUEST_LABELSENTRY.containing_type = _CONVERSATIONUPDATEREQUEST
_CONVERSATIONUPDATEREQUEST.fields_by_name['labels'].message_type = _CONVERSATIONUPDATEREQUEST_LABELSENTRY
_CONVERSATIONUPDATEREQUEST.fields_by_name['spec'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_team_dot_v1alpha1_dot_generated__pb2._CONVERSATIONSPEC
_CONVERSATIONGETRESPONSE.fields_by_name['item'].message_type = github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_team_dot_v1alpha1_dot_generated__pb2._CONVERSATION
DESCRIPTOR.message_types_by_name['ConversationQuery'] = _CONVERSATIONQUERY
DESCRIPTOR.message_types_by_name['ConversationResponse'] = _CONVERSATIONRESPONSE
DESCRIPTOR.message_types_by_name['ConversationCreateRequest'] = _CONVERSATIONCREATEREQUEST
DESCRIPTOR.message_types_by_name['ConversationUpdateRequest'] = _CONVERSATIONUPDATEREQUEST
DESCRIPTOR.message_types_by_name['ConversationGetResponse'] = _CONVERSATIONGETRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ConversationQuery = _reflection.GeneratedProtocolMessageType('ConversationQuery', (_message.Message,), {

  'LabelsEntry' : _reflection.GeneratedProtocolMessageType('LabelsEntry', (_message.Message,), {
    'DESCRIPTOR' : _CONVERSATIONQUERY_LABELSENTRY,
    '__module__' : 'github.com.metaprov.modeldapi.services.conversation.v1.conversation_pb2'
    # @@protoc_insertion_point(class_scope:github.com.metaprov.modeld.services.conversation.v1.ConversationQuery.LabelsEntry)
    })
  ,
  'DESCRIPTOR' : _CONVERSATIONQUERY,
  '__module__' : 'github.com.metaprov.modeldapi.services.conversation.v1.conversation_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeld.services.conversation.v1.ConversationQuery)
  })
_sym_db.RegisterMessage(ConversationQuery)
_sym_db.RegisterMessage(ConversationQuery.LabelsEntry)

ConversationResponse = _reflection.GeneratedProtocolMessageType('ConversationResponse', (_message.Message,), {
  'DESCRIPTOR' : _CONVERSATIONRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.conversation.v1.conversation_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeld.services.conversation.v1.ConversationResponse)
  })
_sym_db.RegisterMessage(ConversationResponse)

ConversationCreateRequest = _reflection.GeneratedProtocolMessageType('ConversationCreateRequest', (_message.Message,), {

  'LabelsEntry' : _reflection.GeneratedProtocolMessageType('LabelsEntry', (_message.Message,), {
    'DESCRIPTOR' : _CONVERSATIONCREATEREQUEST_LABELSENTRY,
    '__module__' : 'github.com.metaprov.modeldapi.services.conversation.v1.conversation_pb2'
    # @@protoc_insertion_point(class_scope:github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest.LabelsEntry)
    })
  ,
  'DESCRIPTOR' : _CONVERSATIONCREATEREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.conversation.v1.conversation_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeld.services.conversation.v1.ConversationCreateRequest)
  })
_sym_db.RegisterMessage(ConversationCreateRequest)
_sym_db.RegisterMessage(ConversationCreateRequest.LabelsEntry)

ConversationUpdateRequest = _reflection.GeneratedProtocolMessageType('ConversationUpdateRequest', (_message.Message,), {

  'LabelsEntry' : _reflection.GeneratedProtocolMessageType('LabelsEntry', (_message.Message,), {
    'DESCRIPTOR' : _CONVERSATIONUPDATEREQUEST_LABELSENTRY,
    '__module__' : 'github.com.metaprov.modeldapi.services.conversation.v1.conversation_pb2'
    # @@protoc_insertion_point(class_scope:github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest.LabelsEntry)
    })
  ,
  'DESCRIPTOR' : _CONVERSATIONUPDATEREQUEST,
  '__module__' : 'github.com.metaprov.modeldapi.services.conversation.v1.conversation_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeld.services.conversation.v1.ConversationUpdateRequest)
  })
_sym_db.RegisterMessage(ConversationUpdateRequest)
_sym_db.RegisterMessage(ConversationUpdateRequest.LabelsEntry)

ConversationGetResponse = _reflection.GeneratedProtocolMessageType('ConversationGetResponse', (_message.Message,), {
  'DESCRIPTOR' : _CONVERSATIONGETRESPONSE,
  '__module__' : 'github.com.metaprov.modeldapi.services.conversation.v1.conversation_pb2'
  # @@protoc_insertion_point(class_scope:github.com.metaprov.modeld.services.conversation.v1.ConversationGetResponse)
  })
_sym_db.RegisterMessage(ConversationGetResponse)


DESCRIPTOR._options = None
_CONVERSATIONQUERY_LABELSENTRY._options = None
_CONVERSATIONCREATEREQUEST_LABELSENTRY._options = None
_CONVERSATIONUPDATEREQUEST_LABELSENTRY._options = None

_CONVERSATIONSERVICE = _descriptor.ServiceDescriptor(
  name='ConversationService',
  full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1251,
  serialized_end=2133,
  methods=[
  _descriptor.MethodDescriptor(
    name='List',
    full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationService.List',
    index=0,
    containing_service=None,
    input_type=_CONVERSATIONQUERY,
    output_type=github_dot_com_dot_metaprov_dot_modeldapi_dot_pkg_dot_apis_dot_team_dot_v1alpha1_dot_generated__pb2._CONVERSATIONLIST,
    serialized_options=b'\202\323\344\223\002\035\022\033/api/v1alpha1/conversations',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Create',
    full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationService.Create',
    index=1,
    containing_service=None,
    input_type=_CONVERSATIONCREATEREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=b'\202\323\344\223\002+\"\033/api/v1alpha1/conversations:\014conversation',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Get',
    full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationService.Get',
    index=2,
    containing_service=None,
    input_type=_CONVERSATIONQUERY,
    output_type=_CONVERSATIONGETRESPONSE,
    serialized_options=b'\202\323\344\223\002$\022\"/api/v1alpha1/conversations/{name}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Update',
    full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationService.Update',
    index=3,
    containing_service=None,
    input_type=_CONVERSATIONUPDATEREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=b'\202\323\344\223\002H\0328/api/v1alpha1/conversations/{conversation.metadata.name}:\014conversation',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Delete',
    full_name='github.com.metaprov.modeld.services.conversation.v1.ConversationService.Delete',
    index=4,
    containing_service=None,
    input_type=_CONVERSATIONQUERY,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_CONVERSATIONSERVICE)

DESCRIPTOR.services_by_name['ConversationService'] = _CONVERSATIONSERVICE

# @@protoc_insertion_point(module_scope)