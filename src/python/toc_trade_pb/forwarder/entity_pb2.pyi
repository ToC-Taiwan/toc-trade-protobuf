from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class StockNumArr(_message.Message):
    __slots__ = ("stock_num_arr", "odd")
    STOCK_NUM_ARR_FIELD_NUMBER: _ClassVar[int]
    ODD_FIELD_NUMBER: _ClassVar[int]
    stock_num_arr: _containers.RepeatedScalarFieldContainer[str]
    odd: bool
    def __init__(self, stock_num_arr: _Optional[_Iterable[str]] = ..., odd: bool = ...) -> None: ...

class FutureCodeArr(_message.Message):
    __slots__ = ("future_code_arr",)
    FUTURE_CODE_ARR_FIELD_NUMBER: _ClassVar[int]
    future_code_arr: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, future_code_arr: _Optional[_Iterable[str]] = ...) -> None: ...

class ErrorMessage(_message.Message):
    __slots__ = ("err",)
    ERR_FIELD_NUMBER: _ClassVar[int]
    err: str
    def __init__(self, err: _Optional[str] = ...) -> None: ...
