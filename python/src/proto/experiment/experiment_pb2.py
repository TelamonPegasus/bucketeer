# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/experiment/experiment.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.feature import variation_pb2 as proto_dot_feature_dot_variation__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!proto/experiment/experiment.proto\x12\x14\x62ucketeer.experiment\x1a\x1dproto/feature/variation.proto\"\x8e\x04\n\nExperiment\x12\n\n\x02id\x18\x01 \x01(\t\x12\x13\n\x07goal_id\x18\x02 \x01(\tB\x02\x18\x01\x12\x12\n\nfeature_id\x18\x03 \x01(\t\x12\x17\n\x0f\x66\x65\x61ture_version\x18\x04 \x01(\x05\x12\x30\n\nvariations\x18\x05 \x03(\x0b\x32\x1c.bucketeer.feature.Variation\x12\x10\n\x08start_at\x18\x06 \x01(\x03\x12\x0f\n\x07stop_at\x18\x07 \x01(\x03\x12\x13\n\x07stopped\x18\x08 \x01(\x08\x42\x02\x18\x01\x12\x16\n\nstopped_at\x18\t \x01(\x03\x42\x02\x30\x01\x12\x12\n\ncreated_at\x18\n \x01(\x03\x12\x12\n\nupdated_at\x18\x0b \x01(\x03\x12\x0f\n\x07\x64\x65leted\x18\x0c \x01(\x08\x12\x10\n\x08goal_ids\x18\r \x03(\t\x12\x0c\n\x04name\x18\x0e \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x0f \x01(\t\x12\x19\n\x11\x62\x61se_variation_id\x18\x10 \x01(\t\x12\x37\n\x06status\x18\x12 \x01(\x0e\x32\'.bucketeer.experiment.Experiment.Status\x12\x12\n\nmaintainer\x18\x13 \x01(\t\x12\x10\n\x08\x61rchived\x18\x14 \x01(\x08\"B\n\x06Status\x12\x0b\n\x07WAITING\x10\x00\x12\x0b\n\x07RUNNING\x10\x01\x12\x0b\n\x07STOPPED\x10\x02\x12\x11\n\rFORCE_STOPPED\x10\x03J\x04\x08\x11\x10\x12\"D\n\x0b\x45xperiments\x12\x35\n\x0b\x65xperiments\x18\x01 \x03(\x0b\x32 .bucketeer.experiment.ExperimentB4Z2github.com/bucketeer-io/bucketeer/proto/experimentb\x06proto3')



_EXPERIMENT = DESCRIPTOR.message_types_by_name['Experiment']
_EXPERIMENTS = DESCRIPTOR.message_types_by_name['Experiments']
_EXPERIMENT_STATUS = _EXPERIMENT.enum_types_by_name['Status']
Experiment = _reflection.GeneratedProtocolMessageType('Experiment', (_message.Message,), {
  'DESCRIPTOR' : _EXPERIMENT,
  '__module__' : 'proto.experiment.experiment_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.experiment.Experiment)
  })
_sym_db.RegisterMessage(Experiment)

Experiments = _reflection.GeneratedProtocolMessageType('Experiments', (_message.Message,), {
  'DESCRIPTOR' : _EXPERIMENTS,
  '__module__' : 'proto.experiment.experiment_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.experiment.Experiments)
  })
_sym_db.RegisterMessage(Experiments)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z2github.com/bucketeer-io/bucketeer/proto/experiment'
  _EXPERIMENT.fields_by_name['goal_id']._options = None
  _EXPERIMENT.fields_by_name['goal_id']._serialized_options = b'\030\001'
  _EXPERIMENT.fields_by_name['stopped']._options = None
  _EXPERIMENT.fields_by_name['stopped']._serialized_options = b'\030\001'
  _EXPERIMENT.fields_by_name['stopped_at']._options = None
  _EXPERIMENT.fields_by_name['stopped_at']._serialized_options = b'0\001'
  _EXPERIMENT._serialized_start=91
  _EXPERIMENT._serialized_end=617
  _EXPERIMENT_STATUS._serialized_start=545
  _EXPERIMENT_STATUS._serialized_end=611
  _EXPERIMENTS._serialized_start=619
  _EXPERIMENTS._serialized_end=687
# @@protoc_insertion_point(module_scope)