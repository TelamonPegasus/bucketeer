# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/feature/target.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1aproto/feature/target.proto\x12\x11\x62ucketeer.feature\"*\n\x06Target\x12\x11\n\tvariation\x18\x01 \x01(\t\x12\r\n\x05users\x18\x02 \x03(\tB1Z/github.com/bucketeer-io/bucketeer/proto/featureb\x06proto3')



_TARGET = DESCRIPTOR.message_types_by_name['Target']
Target = _reflection.GeneratedProtocolMessageType('Target', (_message.Message,), {
  'DESCRIPTOR' : _TARGET,
  '__module__' : 'proto.feature.target_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.feature.Target)
  })
_sym_db.RegisterMessage(Target)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z/github.com/bucketeer-io/bucketeer/proto/feature'
  _TARGET._serialized_start=49
  _TARGET._serialized_end=91
# @@protoc_insertion_point(module_scope)