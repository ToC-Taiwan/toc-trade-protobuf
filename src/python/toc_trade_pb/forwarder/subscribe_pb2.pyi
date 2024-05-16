from google.protobuf import empty_pb2 as _empty_pb2
from forwarder import entity_pb2 as _entity_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class SubscribeResponse(_message.Message):
    __slots__ = ("fail_arr",)
    FAIL_ARR_FIELD_NUMBER: _ClassVar[int]
    fail_arr: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, fail_arr: _Optional[_Iterable[str]] = ...) -> None: ...
