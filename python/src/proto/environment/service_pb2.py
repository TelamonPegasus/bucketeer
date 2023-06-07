# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/environment/service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from proto.environment import environment_pb2 as proto_dot_environment_dot_environment__pb2
from proto.environment import project_pb2 as proto_dot_environment_dot_project__pb2
from proto.environment import command_pb2 as proto_dot_environment_dot_command__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fproto/environment/service.proto\x12\x15\x62ucketeer.environment\x1a\x1egoogle/protobuf/wrappers.proto\x1a#proto/environment/environment.proto\x1a\x1fproto/environment/project.proto\x1a\x1fproto/environment/command.proto\"#\n\x15GetEnvironmentRequest\x12\n\n\x02id\x18\x01 \x01(\t\"Q\n\x16GetEnvironmentResponse\x12\x37\n\x0b\x65nvironment\x18\x01 \x01(\x0b\x32\".bucketeer.environment.Environment\"5\n GetEnvironmentByNamespaceRequest\x12\x11\n\tnamespace\x18\x01 \x01(\t\"\\\n!GetEnvironmentByNamespaceResponse\x12\x37\n\x0b\x65nvironment\x18\x01 \x01(\x0b\x32\".bucketeer.environment.Environment\"\xef\x02\n\x17ListEnvironmentsRequest\x12\x11\n\tpage_size\x18\x01 \x01(\x03\x12\x0e\n\x06\x63ursor\x18\x02 \x01(\t\x12\x12\n\nproject_id\x18\x03 \x01(\t\x12H\n\x08order_by\x18\x04 \x01(\x0e\x32\x36.bucketeer.environment.ListEnvironmentsRequest.OrderBy\x12V\n\x0forder_direction\x18\x05 \x01(\x0e\x32=.bucketeer.environment.ListEnvironmentsRequest.OrderDirection\x12\x16\n\x0esearch_keyword\x18\x06 \x01(\t\">\n\x07OrderBy\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x00\x12\x06\n\x02ID\x10\x01\x12\x0e\n\nCREATED_AT\x10\x02\x12\x0e\n\nUPDATED_AT\x10\x03\"#\n\x0eOrderDirection\x12\x07\n\x03\x41SC\x10\x00\x12\x08\n\x04\x44\x45SC\x10\x01\"y\n\x18ListEnvironmentsResponse\x12\x38\n\x0c\x65nvironments\x18\x01 \x03(\x0b\x32\".bucketeer.environment.Environment\x12\x0e\n\x06\x63ursor\x18\x02 \x01(\t\x12\x13\n\x0btotal_count\x18\x03 \x01(\x03\"\\\n\x18\x43reateEnvironmentRequest\x12@\n\x07\x63ommand\x18\x01 \x01(\x0b\x32/.bucketeer.environment.CreateEnvironmentCommand\"\x1b\n\x19\x43reateEnvironmentResponse\"\xd3\x01\n\x18UpdateEnvironmentRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12K\n\x0erename_command\x18\x02 \x01(\x0b\x32/.bucketeer.environment.RenameEnvironmentCommandB\x02\x18\x01\x12^\n\x1a\x63hange_description_command\x18\x03 \x01(\x0b\x32:.bucketeer.environment.ChangeDescriptionEnvironmentCommand\"\x1b\n\x19UpdateEnvironmentResponse\"h\n\x18\x44\x65leteEnvironmentRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12@\n\x07\x63ommand\x18\x02 \x01(\x0b\x32/.bucketeer.environment.DeleteEnvironmentCommand\"\x1b\n\x19\x44\x65leteEnvironmentResponse\"\x1f\n\x11GetProjectRequest\x12\n\n\x02id\x18\x01 \x01(\t\"E\n\x12GetProjectResponse\x12/\n\x07project\x18\x01 \x01(\x0b\x32\x1e.bucketeer.environment.Project\"\xfd\x02\n\x13ListProjectsRequest\x12\x11\n\tpage_size\x18\x01 \x01(\x03\x12\x0e\n\x06\x63ursor\x18\x02 \x01(\t\x12\x44\n\x08order_by\x18\x03 \x01(\x0e\x32\x32.bucketeer.environment.ListProjectsRequest.OrderBy\x12R\n\x0forder_direction\x18\x04 \x01(\x0e\x32\x39.bucketeer.environment.ListProjectsRequest.OrderDirection\x12\x16\n\x0esearch_keyword\x18\x05 \x01(\t\x12,\n\x08\x64isabled\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.BoolValue\">\n\x07OrderBy\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x00\x12\x06\n\x02ID\x10\x01\x12\x0e\n\nCREATED_AT\x10\x02\x12\x0e\n\nUPDATED_AT\x10\x03\"#\n\x0eOrderDirection\x12\x07\n\x03\x41SC\x10\x00\x12\x08\n\x04\x44\x45SC\x10\x01\"m\n\x14ListProjectsResponse\x12\x30\n\x08projects\x18\x01 \x03(\x0b\x32\x1e.bucketeer.environment.Project\x12\x0e\n\x06\x63ursor\x18\x02 \x01(\t\x12\x13\n\x0btotal_count\x18\x03 \x01(\x03\"T\n\x14\x43reateProjectRequest\x12<\n\x07\x63ommand\x18\x01 \x01(\x0b\x32+.bucketeer.environment.CreateProjectCommand\"\x17\n\x15\x43reateProjectResponse\"^\n\x19\x43reateTrialProjectRequest\x12\x41\n\x07\x63ommand\x18\x01 \x01(\x0b\x32\x30.bucketeer.environment.CreateTrialProjectCommand\"\x1c\n\x1a\x43reateTrialProjectResponse\"~\n\x14UpdateProjectRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12Z\n\x1a\x63hange_description_command\x18\x02 \x01(\x0b\x32\x36.bucketeer.environment.ChangeDescriptionProjectCommand\"\x17\n\x15UpdateProjectResponse\"`\n\x14\x45nableProjectRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12<\n\x07\x63ommand\x18\x02 \x01(\x0b\x32+.bucketeer.environment.EnableProjectCommand\"\x17\n\x15\x45nableProjectResponse\"b\n\x15\x44isableProjectRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12=\n\x07\x63ommand\x18\x02 \x01(\x0b\x32,.bucketeer.environment.DisableProjectCommand\"\x18\n\x16\x44isableProjectResponse\"l\n\x1a\x43onvertTrialProjectRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x42\n\x07\x63ommand\x18\x02 \x01(\x0b\x32\x31.bucketeer.environment.ConvertTrialProjectCommand\"\x1d\n\x1b\x43onvertTrialProjectResponse2\x85\r\n\x12\x45nvironmentService\x12o\n\x0eGetEnvironment\x12,.bucketeer.environment.GetEnvironmentRequest\x1a-.bucketeer.environment.GetEnvironmentResponse\"\x00\x12\x90\x01\n\x19GetEnvironmentByNamespace\x12\x37.bucketeer.environment.GetEnvironmentByNamespaceRequest\x1a\x38.bucketeer.environment.GetEnvironmentByNamespaceResponse\"\x00\x12u\n\x10ListEnvironments\x12..bucketeer.environment.ListEnvironmentsRequest\x1a/.bucketeer.environment.ListEnvironmentsResponse\"\x00\x12x\n\x11\x43reateEnvironment\x12/.bucketeer.environment.CreateEnvironmentRequest\x1a\x30.bucketeer.environment.CreateEnvironmentResponse\"\x00\x12x\n\x11UpdateEnvironment\x12/.bucketeer.environment.UpdateEnvironmentRequest\x1a\x30.bucketeer.environment.UpdateEnvironmentResponse\"\x00\x12x\n\x11\x44\x65leteEnvironment\x12/.bucketeer.environment.DeleteEnvironmentRequest\x1a\x30.bucketeer.environment.DeleteEnvironmentResponse\"\x00\x12\x63\n\nGetProject\x12(.bucketeer.environment.GetProjectRequest\x1a).bucketeer.environment.GetProjectResponse\"\x00\x12i\n\x0cListProjects\x12*.bucketeer.environment.ListProjectsRequest\x1a+.bucketeer.environment.ListProjectsResponse\"\x00\x12l\n\rCreateProject\x12+.bucketeer.environment.CreateProjectRequest\x1a,.bucketeer.environment.CreateProjectResponse\"\x00\x12{\n\x12\x43reateTrialProject\x12\x30.bucketeer.environment.CreateTrialProjectRequest\x1a\x31.bucketeer.environment.CreateTrialProjectResponse\"\x00\x12l\n\rUpdateProject\x12+.bucketeer.environment.UpdateProjectRequest\x1a,.bucketeer.environment.UpdateProjectResponse\"\x00\x12l\n\rEnableProject\x12+.bucketeer.environment.EnableProjectRequest\x1a,.bucketeer.environment.EnableProjectResponse\"\x00\x12o\n\x0e\x44isableProject\x12,.bucketeer.environment.DisableProjectRequest\x1a-.bucketeer.environment.DisableProjectResponse\"\x00\x12~\n\x13\x43onvertTrialProject\x12\x31.bucketeer.environment.ConvertTrialProjectRequest\x1a\x32.bucketeer.environment.ConvertTrialProjectResponse\"\x00\x42\x35Z3github.com/bucketeer-io/bucketeer/proto/environmentb\x06proto3')



_GETENVIRONMENTREQUEST = DESCRIPTOR.message_types_by_name['GetEnvironmentRequest']
_GETENVIRONMENTRESPONSE = DESCRIPTOR.message_types_by_name['GetEnvironmentResponse']
_GETENVIRONMENTBYNAMESPACEREQUEST = DESCRIPTOR.message_types_by_name['GetEnvironmentByNamespaceRequest']
_GETENVIRONMENTBYNAMESPACERESPONSE = DESCRIPTOR.message_types_by_name['GetEnvironmentByNamespaceResponse']
_LISTENVIRONMENTSREQUEST = DESCRIPTOR.message_types_by_name['ListEnvironmentsRequest']
_LISTENVIRONMENTSRESPONSE = DESCRIPTOR.message_types_by_name['ListEnvironmentsResponse']
_CREATEENVIRONMENTREQUEST = DESCRIPTOR.message_types_by_name['CreateEnvironmentRequest']
_CREATEENVIRONMENTRESPONSE = DESCRIPTOR.message_types_by_name['CreateEnvironmentResponse']
_UPDATEENVIRONMENTREQUEST = DESCRIPTOR.message_types_by_name['UpdateEnvironmentRequest']
_UPDATEENVIRONMENTRESPONSE = DESCRIPTOR.message_types_by_name['UpdateEnvironmentResponse']
_DELETEENVIRONMENTREQUEST = DESCRIPTOR.message_types_by_name['DeleteEnvironmentRequest']
_DELETEENVIRONMENTRESPONSE = DESCRIPTOR.message_types_by_name['DeleteEnvironmentResponse']
_GETPROJECTREQUEST = DESCRIPTOR.message_types_by_name['GetProjectRequest']
_GETPROJECTRESPONSE = DESCRIPTOR.message_types_by_name['GetProjectResponse']
_LISTPROJECTSREQUEST = DESCRIPTOR.message_types_by_name['ListProjectsRequest']
_LISTPROJECTSRESPONSE = DESCRIPTOR.message_types_by_name['ListProjectsResponse']
_CREATEPROJECTREQUEST = DESCRIPTOR.message_types_by_name['CreateProjectRequest']
_CREATEPROJECTRESPONSE = DESCRIPTOR.message_types_by_name['CreateProjectResponse']
_CREATETRIALPROJECTREQUEST = DESCRIPTOR.message_types_by_name['CreateTrialProjectRequest']
_CREATETRIALPROJECTRESPONSE = DESCRIPTOR.message_types_by_name['CreateTrialProjectResponse']
_UPDATEPROJECTREQUEST = DESCRIPTOR.message_types_by_name['UpdateProjectRequest']
_UPDATEPROJECTRESPONSE = DESCRIPTOR.message_types_by_name['UpdateProjectResponse']
_ENABLEPROJECTREQUEST = DESCRIPTOR.message_types_by_name['EnableProjectRequest']
_ENABLEPROJECTRESPONSE = DESCRIPTOR.message_types_by_name['EnableProjectResponse']
_DISABLEPROJECTREQUEST = DESCRIPTOR.message_types_by_name['DisableProjectRequest']
_DISABLEPROJECTRESPONSE = DESCRIPTOR.message_types_by_name['DisableProjectResponse']
_CONVERTTRIALPROJECTREQUEST = DESCRIPTOR.message_types_by_name['ConvertTrialProjectRequest']
_CONVERTTRIALPROJECTRESPONSE = DESCRIPTOR.message_types_by_name['ConvertTrialProjectResponse']
_LISTENVIRONMENTSREQUEST_ORDERBY = _LISTENVIRONMENTSREQUEST.enum_types_by_name['OrderBy']
_LISTENVIRONMENTSREQUEST_ORDERDIRECTION = _LISTENVIRONMENTSREQUEST.enum_types_by_name['OrderDirection']
_LISTPROJECTSREQUEST_ORDERBY = _LISTPROJECTSREQUEST.enum_types_by_name['OrderBy']
_LISTPROJECTSREQUEST_ORDERDIRECTION = _LISTPROJECTSREQUEST.enum_types_by_name['OrderDirection']
GetEnvironmentRequest = _reflection.GeneratedProtocolMessageType('GetEnvironmentRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETENVIRONMENTREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.GetEnvironmentRequest)
  })
_sym_db.RegisterMessage(GetEnvironmentRequest)

GetEnvironmentResponse = _reflection.GeneratedProtocolMessageType('GetEnvironmentResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETENVIRONMENTRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.GetEnvironmentResponse)
  })
_sym_db.RegisterMessage(GetEnvironmentResponse)

GetEnvironmentByNamespaceRequest = _reflection.GeneratedProtocolMessageType('GetEnvironmentByNamespaceRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETENVIRONMENTBYNAMESPACEREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.GetEnvironmentByNamespaceRequest)
  })
_sym_db.RegisterMessage(GetEnvironmentByNamespaceRequest)

GetEnvironmentByNamespaceResponse = _reflection.GeneratedProtocolMessageType('GetEnvironmentByNamespaceResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETENVIRONMENTBYNAMESPACERESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.GetEnvironmentByNamespaceResponse)
  })
_sym_db.RegisterMessage(GetEnvironmentByNamespaceResponse)

ListEnvironmentsRequest = _reflection.GeneratedProtocolMessageType('ListEnvironmentsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTENVIRONMENTSREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.ListEnvironmentsRequest)
  })
_sym_db.RegisterMessage(ListEnvironmentsRequest)

ListEnvironmentsResponse = _reflection.GeneratedProtocolMessageType('ListEnvironmentsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTENVIRONMENTSRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.ListEnvironmentsResponse)
  })
_sym_db.RegisterMessage(ListEnvironmentsResponse)

CreateEnvironmentRequest = _reflection.GeneratedProtocolMessageType('CreateEnvironmentRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEENVIRONMENTREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.CreateEnvironmentRequest)
  })
_sym_db.RegisterMessage(CreateEnvironmentRequest)

CreateEnvironmentResponse = _reflection.GeneratedProtocolMessageType('CreateEnvironmentResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEENVIRONMENTRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.CreateEnvironmentResponse)
  })
_sym_db.RegisterMessage(CreateEnvironmentResponse)

UpdateEnvironmentRequest = _reflection.GeneratedProtocolMessageType('UpdateEnvironmentRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEENVIRONMENTREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.UpdateEnvironmentRequest)
  })
_sym_db.RegisterMessage(UpdateEnvironmentRequest)

UpdateEnvironmentResponse = _reflection.GeneratedProtocolMessageType('UpdateEnvironmentResponse', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEENVIRONMENTRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.UpdateEnvironmentResponse)
  })
_sym_db.RegisterMessage(UpdateEnvironmentResponse)

DeleteEnvironmentRequest = _reflection.GeneratedProtocolMessageType('DeleteEnvironmentRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEENVIRONMENTREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.DeleteEnvironmentRequest)
  })
_sym_db.RegisterMessage(DeleteEnvironmentRequest)

DeleteEnvironmentResponse = _reflection.GeneratedProtocolMessageType('DeleteEnvironmentResponse', (_message.Message,), {
  'DESCRIPTOR' : _DELETEENVIRONMENTRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.DeleteEnvironmentResponse)
  })
_sym_db.RegisterMessage(DeleteEnvironmentResponse)

GetProjectRequest = _reflection.GeneratedProtocolMessageType('GetProjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETPROJECTREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.GetProjectRequest)
  })
_sym_db.RegisterMessage(GetProjectRequest)

GetProjectResponse = _reflection.GeneratedProtocolMessageType('GetProjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETPROJECTRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.GetProjectResponse)
  })
_sym_db.RegisterMessage(GetProjectResponse)

ListProjectsRequest = _reflection.GeneratedProtocolMessageType('ListProjectsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTPROJECTSREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.ListProjectsRequest)
  })
_sym_db.RegisterMessage(ListProjectsRequest)

ListProjectsResponse = _reflection.GeneratedProtocolMessageType('ListProjectsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTPROJECTSRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.ListProjectsResponse)
  })
_sym_db.RegisterMessage(ListProjectsResponse)

CreateProjectRequest = _reflection.GeneratedProtocolMessageType('CreateProjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPROJECTREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.CreateProjectRequest)
  })
_sym_db.RegisterMessage(CreateProjectRequest)

CreateProjectResponse = _reflection.GeneratedProtocolMessageType('CreateProjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPROJECTRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.CreateProjectResponse)
  })
_sym_db.RegisterMessage(CreateProjectResponse)

CreateTrialProjectRequest = _reflection.GeneratedProtocolMessageType('CreateTrialProjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATETRIALPROJECTREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.CreateTrialProjectRequest)
  })
_sym_db.RegisterMessage(CreateTrialProjectRequest)

CreateTrialProjectResponse = _reflection.GeneratedProtocolMessageType('CreateTrialProjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATETRIALPROJECTRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.CreateTrialProjectResponse)
  })
_sym_db.RegisterMessage(CreateTrialProjectResponse)

UpdateProjectRequest = _reflection.GeneratedProtocolMessageType('UpdateProjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEPROJECTREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.UpdateProjectRequest)
  })
_sym_db.RegisterMessage(UpdateProjectRequest)

UpdateProjectResponse = _reflection.GeneratedProtocolMessageType('UpdateProjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEPROJECTRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.UpdateProjectResponse)
  })
_sym_db.RegisterMessage(UpdateProjectResponse)

EnableProjectRequest = _reflection.GeneratedProtocolMessageType('EnableProjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _ENABLEPROJECTREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.EnableProjectRequest)
  })
_sym_db.RegisterMessage(EnableProjectRequest)

EnableProjectResponse = _reflection.GeneratedProtocolMessageType('EnableProjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _ENABLEPROJECTRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.EnableProjectResponse)
  })
_sym_db.RegisterMessage(EnableProjectResponse)

DisableProjectRequest = _reflection.GeneratedProtocolMessageType('DisableProjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _DISABLEPROJECTREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.DisableProjectRequest)
  })
_sym_db.RegisterMessage(DisableProjectRequest)

DisableProjectResponse = _reflection.GeneratedProtocolMessageType('DisableProjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _DISABLEPROJECTRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.DisableProjectResponse)
  })
_sym_db.RegisterMessage(DisableProjectResponse)

ConvertTrialProjectRequest = _reflection.GeneratedProtocolMessageType('ConvertTrialProjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _CONVERTTRIALPROJECTREQUEST,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.ConvertTrialProjectRequest)
  })
_sym_db.RegisterMessage(ConvertTrialProjectRequest)

ConvertTrialProjectResponse = _reflection.GeneratedProtocolMessageType('ConvertTrialProjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _CONVERTTRIALPROJECTRESPONSE,
  '__module__' : 'proto.environment.service_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.environment.ConvertTrialProjectResponse)
  })
_sym_db.RegisterMessage(ConvertTrialProjectResponse)

_ENVIRONMENTSERVICE = DESCRIPTOR.services_by_name['EnvironmentService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z3github.com/bucketeer-io/bucketeer/proto/environment'
  _UPDATEENVIRONMENTREQUEST.fields_by_name['rename_command']._options = None
  _UPDATEENVIRONMENTREQUEST.fields_by_name['rename_command']._serialized_options = b'\030\001'
  _GETENVIRONMENTREQUEST._serialized_start=193
  _GETENVIRONMENTREQUEST._serialized_end=228
  _GETENVIRONMENTRESPONSE._serialized_start=230
  _GETENVIRONMENTRESPONSE._serialized_end=311
  _GETENVIRONMENTBYNAMESPACEREQUEST._serialized_start=313
  _GETENVIRONMENTBYNAMESPACEREQUEST._serialized_end=366
  _GETENVIRONMENTBYNAMESPACERESPONSE._serialized_start=368
  _GETENVIRONMENTBYNAMESPACERESPONSE._serialized_end=460
  _LISTENVIRONMENTSREQUEST._serialized_start=463
  _LISTENVIRONMENTSREQUEST._serialized_end=830
  _LISTENVIRONMENTSREQUEST_ORDERBY._serialized_start=731
  _LISTENVIRONMENTSREQUEST_ORDERBY._serialized_end=793
  _LISTENVIRONMENTSREQUEST_ORDERDIRECTION._serialized_start=795
  _LISTENVIRONMENTSREQUEST_ORDERDIRECTION._serialized_end=830
  _LISTENVIRONMENTSRESPONSE._serialized_start=832
  _LISTENVIRONMENTSRESPONSE._serialized_end=953
  _CREATEENVIRONMENTREQUEST._serialized_start=955
  _CREATEENVIRONMENTREQUEST._serialized_end=1047
  _CREATEENVIRONMENTRESPONSE._serialized_start=1049
  _CREATEENVIRONMENTRESPONSE._serialized_end=1076
  _UPDATEENVIRONMENTREQUEST._serialized_start=1079
  _UPDATEENVIRONMENTREQUEST._serialized_end=1290
  _UPDATEENVIRONMENTRESPONSE._serialized_start=1292
  _UPDATEENVIRONMENTRESPONSE._serialized_end=1319
  _DELETEENVIRONMENTREQUEST._serialized_start=1321
  _DELETEENVIRONMENTREQUEST._serialized_end=1425
  _DELETEENVIRONMENTRESPONSE._serialized_start=1427
  _DELETEENVIRONMENTRESPONSE._serialized_end=1454
  _GETPROJECTREQUEST._serialized_start=1456
  _GETPROJECTREQUEST._serialized_end=1487
  _GETPROJECTRESPONSE._serialized_start=1489
  _GETPROJECTRESPONSE._serialized_end=1558
  _LISTPROJECTSREQUEST._serialized_start=1561
  _LISTPROJECTSREQUEST._serialized_end=1942
  _LISTPROJECTSREQUEST_ORDERBY._serialized_start=731
  _LISTPROJECTSREQUEST_ORDERBY._serialized_end=793
  _LISTPROJECTSREQUEST_ORDERDIRECTION._serialized_start=795
  _LISTPROJECTSREQUEST_ORDERDIRECTION._serialized_end=830
  _LISTPROJECTSRESPONSE._serialized_start=1944
  _LISTPROJECTSRESPONSE._serialized_end=2053
  _CREATEPROJECTREQUEST._serialized_start=2055
  _CREATEPROJECTREQUEST._serialized_end=2139
  _CREATEPROJECTRESPONSE._serialized_start=2141
  _CREATEPROJECTRESPONSE._serialized_end=2164
  _CREATETRIALPROJECTREQUEST._serialized_start=2166
  _CREATETRIALPROJECTREQUEST._serialized_end=2260
  _CREATETRIALPROJECTRESPONSE._serialized_start=2262
  _CREATETRIALPROJECTRESPONSE._serialized_end=2290
  _UPDATEPROJECTREQUEST._serialized_start=2292
  _UPDATEPROJECTREQUEST._serialized_end=2418
  _UPDATEPROJECTRESPONSE._serialized_start=2420
  _UPDATEPROJECTRESPONSE._serialized_end=2443
  _ENABLEPROJECTREQUEST._serialized_start=2445
  _ENABLEPROJECTREQUEST._serialized_end=2541
  _ENABLEPROJECTRESPONSE._serialized_start=2543
  _ENABLEPROJECTRESPONSE._serialized_end=2566
  _DISABLEPROJECTREQUEST._serialized_start=2568
  _DISABLEPROJECTREQUEST._serialized_end=2666
  _DISABLEPROJECTRESPONSE._serialized_start=2668
  _DISABLEPROJECTRESPONSE._serialized_end=2692
  _CONVERTTRIALPROJECTREQUEST._serialized_start=2694
  _CONVERTTRIALPROJECTREQUEST._serialized_end=2802
  _CONVERTTRIALPROJECTRESPONSE._serialized_start=2804
  _CONVERTTRIALPROJECTRESPONSE._serialized_end=2833
  _ENVIRONMENTSERVICE._serialized_start=2836
  _ENVIRONMENTSERVICE._serialized_end=4505
# @@protoc_insertion_point(module_scope)
