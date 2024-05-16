from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class StockNumArrWithDate(_message.Message):
    __slots__ = ("stock_num_arr", "date")
    STOCK_NUM_ARR_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    stock_num_arr: _containers.RepeatedScalarFieldContainer[str]
    date: str
    def __init__(self, stock_num_arr: _Optional[_Iterable[str]] = ..., date: _Optional[str] = ...) -> None: ...

class FutureCodeArrWithDate(_message.Message):
    __slots__ = ("future_code_arr", "date")
    FUTURE_CODE_ARR_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    future_code_arr: _containers.RepeatedScalarFieldContainer[str]
    date: str
    def __init__(self, future_code_arr: _Optional[_Iterable[str]] = ..., date: _Optional[str] = ...) -> None: ...

class StockNumArrWithDateArr(_message.Message):
    __slots__ = ("stock_num_arr", "date_arr")
    STOCK_NUM_ARR_FIELD_NUMBER: _ClassVar[int]
    DATE_ARR_FIELD_NUMBER: _ClassVar[int]
    stock_num_arr: _containers.RepeatedScalarFieldContainer[str]
    date_arr: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, stock_num_arr: _Optional[_Iterable[str]] = ..., date_arr: _Optional[_Iterable[str]] = ...) -> None: ...

class HistoryTickResponse(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[HistoryTickMessage]
    def __init__(self, data: _Optional[_Iterable[_Union[HistoryTickMessage, _Mapping]]] = ...) -> None: ...

class HistoryKbarResponse(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[HistoryKbarMessage]
    def __init__(self, data: _Optional[_Iterable[_Union[HistoryKbarMessage, _Mapping]]] = ...) -> None: ...

class HistoryCloseResponse(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[HistoryCloseMessage]
    def __init__(self, data: _Optional[_Iterable[_Union[HistoryCloseMessage, _Mapping]]] = ...) -> None: ...

class HistoryTickMessage(_message.Message):
    __slots__ = ("ts", "close", "volume", "bid_price", "bid_volume", "ask_price", "ask_volume", "tick_type", "code")
    TS_FIELD_NUMBER: _ClassVar[int]
    CLOSE_FIELD_NUMBER: _ClassVar[int]
    VOLUME_FIELD_NUMBER: _ClassVar[int]
    BID_PRICE_FIELD_NUMBER: _ClassVar[int]
    BID_VOLUME_FIELD_NUMBER: _ClassVar[int]
    ASK_PRICE_FIELD_NUMBER: _ClassVar[int]
    ASK_VOLUME_FIELD_NUMBER: _ClassVar[int]
    TICK_TYPE_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    ts: int
    close: float
    volume: int
    bid_price: float
    bid_volume: int
    ask_price: float
    ask_volume: int
    tick_type: int
    code: str
    def __init__(self, ts: _Optional[int] = ..., close: _Optional[float] = ..., volume: _Optional[int] = ..., bid_price: _Optional[float] = ..., bid_volume: _Optional[int] = ..., ask_price: _Optional[float] = ..., ask_volume: _Optional[int] = ..., tick_type: _Optional[int] = ..., code: _Optional[str] = ...) -> None: ...

class HistoryKbarMessage(_message.Message):
    __slots__ = ("ts", "close", "open", "high", "low", "volume", "code")
    TS_FIELD_NUMBER: _ClassVar[int]
    CLOSE_FIELD_NUMBER: _ClassVar[int]
    OPEN_FIELD_NUMBER: _ClassVar[int]
    HIGH_FIELD_NUMBER: _ClassVar[int]
    LOW_FIELD_NUMBER: _ClassVar[int]
    VOLUME_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    ts: int
    close: float
    open: float
    high: float
    low: float
    volume: int
    code: str
    def __init__(self, ts: _Optional[int] = ..., close: _Optional[float] = ..., open: _Optional[float] = ..., high: _Optional[float] = ..., low: _Optional[float] = ..., volume: _Optional[int] = ..., code: _Optional[str] = ...) -> None: ...

class HistoryCloseMessage(_message.Message):
    __slots__ = ("date", "code", "close")
    DATE_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    CLOSE_FIELD_NUMBER: _ClassVar[int]
    date: str
    code: str
    close: float
    def __init__(self, date: _Optional[str] = ..., code: _Optional[str] = ..., close: _Optional[float] = ...) -> None: ...

class Date(_message.Message):
    __slots__ = ("date",)
    DATE_FIELD_NUMBER: _ClassVar[int]
    date: str
    def __init__(self, date: _Optional[str] = ...) -> None: ...
