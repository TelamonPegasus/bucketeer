# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/account/api_key.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bproto/account/api_key.proto\x12\x11\x62ucketeer.account\"\xa8\x01\n\x06\x41PIKey\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12,\n\x04role\x18\x03 \x01(\x0e\x32\x1e.bucketeer.account.APIKey.Role\x12\x10\n\x08\x64isabled\x18\x04 \x01(\x08\x12\x12\n\ncreated_at\x18\x05 \x01(\x03\x12\x12\n\nupdated_at\x18\x06 \x01(\x03\"\x1c\n\x04Role\x12\x07\n\x03SDK\x10\x00\x12\x0b\n\x07SERVICE\x10\x01\"\x90\x01\n\x11\x45nvironmentAPIKey\x12\x1d\n\x15\x65nvironment_namespace\x18\x01 \x01(\t\x12*\n\x07\x61pi_key\x18\x02 \x01(\x0b\x32\x19.bucketeer.account.APIKey\x12\x1c\n\x14\x65nvironment_disabled\x18\x03 \x01(\x08\x12\x12\n\nproject_id\x18\x04 \x01(\tB1Z/github.com/bucketeer-io/bucketeer/proto/accountb\x06proto3')



_APIKEY = DESCRIPTOR.message_types_by_name['APIKey']
_ENVIRONMENTAPIKEY = DESCRIPTOR.message_types_by_name['EnvironmentAPIKey']
_APIKEY_ROLE = _APIKEY.enum_types_by_name['Role']
APIKey = _reflection.GeneratedProtocolMessageType('APIKey', (_message.Message,), {
  'DESCRIPTOR' : _APIKEY,
  '__module__' : 'proto.account.api_key_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.account.APIKey)
  })
_sym_db.RegisterMessage(APIKey)

EnvironmentAPIKey = _reflection.GeneratedProtocolMessageType('EnvironmentAPIKey', (_message.Message,), {
  'DESCRIPTOR' : _ENVIRONMENTAPIKEY,
  '__module__' : 'proto.account.api_key_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.account.EnvironmentAPIKey)
  })
_sym_db.RegisterMessage(EnvironmentAPIKey)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z/github.com/bucketeer-io/bucketeer/proto/account'
  _APIKEY._serialized_start=51
  _APIKEY._serialized_end=219
  _APIKEY_ROLE._serialized_start=191
  _APIKEY_ROLE._serialized_end=219
  _ENVIRONMENTAPIKEY._serialized_start=222
  _ENVIRONMENTAPIKEY._serialized_end=366
# @@protoc_insertion_point(module_scope)
