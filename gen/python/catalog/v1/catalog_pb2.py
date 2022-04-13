# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: catalog/v1/catalog.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import field_behavior_pb2 as google_dot_api_dot_field__behavior__pb2
from google.api import resource_pb2 as google_dot_api_dot_resource__pb2
from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18\x63\x61talog/v1/catalog.proto\x12\ncatalog.v1\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a google/protobuf/field_mask.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x17validate/validate.proto\"\xc5\x02\n\x07\x43\x61talog\x12\x17\n\x04name\x18\x01 \x01(\tB\x03\xe0\x41\x03R\x04name\x12\x15\n\x03uid\x18\x02 \x01(\tB\x03\xe0\x41\x03R\x03uid\x12-\n\x0c\x64isplay_name\x18\x03 \x01(\tB\n\xfa\x42\x07r\x05 \x01(\x80\x02R\x0b\x64isplayName\x12\x43\n\x0b\x63reate_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x06\xe0\x41\x05\xe0\x41\x03R\ncreateTime\x12\x43\n\x0bupdate_time\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x06\xe0\x41\x05\xe0\x41\x03R\nupdateTime:Q\xea\x41N\n$catalog-api-toolkit.apis.com/Catalog\x12&companies/{company}/catalogs/{catalog}\"\x90\x01\n\x14\x43reateCatalogRequest\x12\x44\n\x06parent\x18\x01 \x01(\tB,\xe0\x41\x02\xfa\x41&\x12$catalog-api-toolkit.apis.com/CatalogR\x06parent\x12\x32\n\x07\x63\x61talog\x18\x02 \x01(\x0b\x32\x13.catalog.v1.CatalogB\x03\xe0\x41\x02R\x07\x63\x61talog\"U\n\x11GetCatalogRequest\x12@\n\x04name\x18\x01 \x01(\tB,\xe0\x41\x02\xfa\x41&\n$catalog-api-toolkit.apis.com/CatalogR\x04name\"\xca\x01\n\x13ListCatalogsRequest\x12\x44\n\x06parent\x18\x01 \x01(\tB,\xe0\x41\x02\xfa\x41&\x12$catalog-api-toolkit.apis.com/CatalogR\x06parent\x12\x1b\n\tpage_size\x18\x02 \x01(\x05R\x08pageSize\x12\x1d\n\npage_token\x18\x03 \x01(\tR\tpageToken\x12\x16\n\x06\x66ilter\x18\x04 \x01(\tR\x06\x66ilter\x12\x19\n\x08order_by\x18\x05 \x01(\tR\x07orderBy\"o\n\x14ListCatalogsResponse\x12/\n\x08\x63\x61talogs\x18\x01 \x03(\x0b\x32\x13.catalog.v1.CatalogR\x08\x63\x61talogs\x12&\n\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\x87\x01\n\x14UpdateCatalogRequest\x12\x32\n\x07\x63\x61talog\x18\x01 \x01(\x0b\x32\x13.catalog.v1.CatalogB\x03\xe0\x41\x02R\x07\x63\x61talog\x12;\n\x0bupdate_mask\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.FieldMaskR\nupdateMask\"n\n\x14\x44\x65leteCatalogRequest\x12@\n\x04name\x18\x01 \x01(\tB,\xe0\x41\x02\xfa\x41&\x12$catalog-api-toolkit.apis.com/CatalogR\x04name\x12\x14\n\x05\x66orce\x18\x02 \x01(\x08R\x05\x66orceB9Z7github.com/tyrocca/catalog-api-toolkit/proto/catalog/v1b\x06proto3')



_CATALOG = DESCRIPTOR.message_types_by_name['Catalog']
_CREATECATALOGREQUEST = DESCRIPTOR.message_types_by_name['CreateCatalogRequest']
_GETCATALOGREQUEST = DESCRIPTOR.message_types_by_name['GetCatalogRequest']
_LISTCATALOGSREQUEST = DESCRIPTOR.message_types_by_name['ListCatalogsRequest']
_LISTCATALOGSRESPONSE = DESCRIPTOR.message_types_by_name['ListCatalogsResponse']
_UPDATECATALOGREQUEST = DESCRIPTOR.message_types_by_name['UpdateCatalogRequest']
_DELETECATALOGREQUEST = DESCRIPTOR.message_types_by_name['DeleteCatalogRequest']
Catalog = _reflection.GeneratedProtocolMessageType('Catalog', (_message.Message,), {
  'DESCRIPTOR' : _CATALOG,
  '__module__' : 'catalog.v1.catalog_pb2'
  # @@protoc_insertion_point(class_scope:catalog.v1.Catalog)
  })
_sym_db.RegisterMessage(Catalog)

CreateCatalogRequest = _reflection.GeneratedProtocolMessageType('CreateCatalogRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATECATALOGREQUEST,
  '__module__' : 'catalog.v1.catalog_pb2'
  # @@protoc_insertion_point(class_scope:catalog.v1.CreateCatalogRequest)
  })
_sym_db.RegisterMessage(CreateCatalogRequest)

GetCatalogRequest = _reflection.GeneratedProtocolMessageType('GetCatalogRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETCATALOGREQUEST,
  '__module__' : 'catalog.v1.catalog_pb2'
  # @@protoc_insertion_point(class_scope:catalog.v1.GetCatalogRequest)
  })
_sym_db.RegisterMessage(GetCatalogRequest)

ListCatalogsRequest = _reflection.GeneratedProtocolMessageType('ListCatalogsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTCATALOGSREQUEST,
  '__module__' : 'catalog.v1.catalog_pb2'
  # @@protoc_insertion_point(class_scope:catalog.v1.ListCatalogsRequest)
  })
_sym_db.RegisterMessage(ListCatalogsRequest)

ListCatalogsResponse = _reflection.GeneratedProtocolMessageType('ListCatalogsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTCATALOGSRESPONSE,
  '__module__' : 'catalog.v1.catalog_pb2'
  # @@protoc_insertion_point(class_scope:catalog.v1.ListCatalogsResponse)
  })
_sym_db.RegisterMessage(ListCatalogsResponse)

UpdateCatalogRequest = _reflection.GeneratedProtocolMessageType('UpdateCatalogRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATECATALOGREQUEST,
  '__module__' : 'catalog.v1.catalog_pb2'
  # @@protoc_insertion_point(class_scope:catalog.v1.UpdateCatalogRequest)
  })
_sym_db.RegisterMessage(UpdateCatalogRequest)

DeleteCatalogRequest = _reflection.GeneratedProtocolMessageType('DeleteCatalogRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETECATALOGREQUEST,
  '__module__' : 'catalog.v1.catalog_pb2'
  # @@protoc_insertion_point(class_scope:catalog.v1.DeleteCatalogRequest)
  })
_sym_db.RegisterMessage(DeleteCatalogRequest)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z7github.com/tyrocca/catalog-api-toolkit/proto/catalog/v1'
  _CATALOG.fields_by_name['name']._options = None
  _CATALOG.fields_by_name['name']._serialized_options = b'\340A\003'
  _CATALOG.fields_by_name['uid']._options = None
  _CATALOG.fields_by_name['uid']._serialized_options = b'\340A\003'
  _CATALOG.fields_by_name['display_name']._options = None
  _CATALOG.fields_by_name['display_name']._serialized_options = b'\372B\007r\005 \001(\200\002'
  _CATALOG.fields_by_name['create_time']._options = None
  _CATALOG.fields_by_name['create_time']._serialized_options = b'\340A\005\340A\003'
  _CATALOG.fields_by_name['update_time']._options = None
  _CATALOG.fields_by_name['update_time']._serialized_options = b'\340A\005\340A\003'
  _CATALOG._options = None
  _CATALOG._serialized_options = b'\352AN\n$catalog-api-toolkit.apis.com/Catalog\022&companies/{company}/catalogs/{catalog}'
  _CREATECATALOGREQUEST.fields_by_name['parent']._options = None
  _CREATECATALOGREQUEST.fields_by_name['parent']._serialized_options = b'\340A\002\372A&\022$catalog-api-toolkit.apis.com/Catalog'
  _CREATECATALOGREQUEST.fields_by_name['catalog']._options = None
  _CREATECATALOGREQUEST.fields_by_name['catalog']._serialized_options = b'\340A\002'
  _GETCATALOGREQUEST.fields_by_name['name']._options = None
  _GETCATALOGREQUEST.fields_by_name['name']._serialized_options = b'\340A\002\372A&\n$catalog-api-toolkit.apis.com/Catalog'
  _LISTCATALOGSREQUEST.fields_by_name['parent']._options = None
  _LISTCATALOGSREQUEST.fields_by_name['parent']._serialized_options = b'\340A\002\372A&\022$catalog-api-toolkit.apis.com/Catalog'
  _UPDATECATALOGREQUEST.fields_by_name['catalog']._options = None
  _UPDATECATALOGREQUEST.fields_by_name['catalog']._serialized_options = b'\340A\002'
  _DELETECATALOGREQUEST.fields_by_name['name']._options = None
  _DELETECATALOGREQUEST.fields_by_name['name']._serialized_options = b'\340A\002\372A&\022$catalog-api-toolkit.apis.com/Catalog'
  _CATALOG._serialized_start=193
  _CATALOG._serialized_end=518
  _CREATECATALOGREQUEST._serialized_start=521
  _CREATECATALOGREQUEST._serialized_end=665
  _GETCATALOGREQUEST._serialized_start=667
  _GETCATALOGREQUEST._serialized_end=752
  _LISTCATALOGSREQUEST._serialized_start=755
  _LISTCATALOGSREQUEST._serialized_end=957
  _LISTCATALOGSRESPONSE._serialized_start=959
  _LISTCATALOGSRESPONSE._serialized_end=1070
  _UPDATECATALOGREQUEST._serialized_start=1073
  _UPDATECATALOGREQUEST._serialized_end=1208
  _DELETECATALOGREQUEST._serialized_start=1210
  _DELETECATALOGREQUEST._serialized_end=1320
# @@protoc_insertion_point(module_scope)