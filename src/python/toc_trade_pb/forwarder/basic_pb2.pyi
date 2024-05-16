from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class BeatMessage(_message.Message):
    __slots__ = ("message", "error")
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    message: str
    error: str
    def __init__(self, message: _Optional[str] = ..., error: _Optional[str] = ...) -> None: ...

class ShioajiUsage(_message.Message):
    __slots__ = ("connections", "bytes", "limit_bytes", "remaining_bytes")
    CONNECTIONS_FIELD_NUMBER: _ClassVar[int]
    BYTES_FIELD_NUMBER: _ClassVar[int]
    LIMIT_BYTES_FIELD_NUMBER: _ClassVar[int]
    REMAINING_BYTES_FIELD_NUMBER: _ClassVar[int]
    connections: int
    bytes: int
    limit_bytes: int
    remaining_bytes: int
    def __init__(self, connections: _Optional[int] = ..., bytes: _Optional[int] = ..., limit_bytes: _Optional[int] = ..., remaining_bytes: _Optional[int] = ...) -> None: ...

class StockDetailResponse(_message.Message):
    __slots__ = ("stock",)
    STOCK_FIELD_NUMBER: _ClassVar[int]
    stock: _containers.RepeatedCompositeFieldContainer[StockDetailMessage]
    def __init__(self, stock: _Optional[_Iterable[_Union[StockDetailMessage, _Mapping]]] = ...) -> None: ...

class FutureDetailResponse(_message.Message):
    __slots__ = ("future",)
    FUTURE_FIELD_NUMBER: _ClassVar[int]
    future: _containers.RepeatedCompositeFieldContainer[FutureDetailMessage]
    def __init__(self, future: _Optional[_Iterable[_Union[FutureDetailMessage, _Mapping]]] = ...) -> None: ...

class OptionDetailResponse(_message.Message):
    __slots__ = ("option",)
    OPTION_FIELD_NUMBER: _ClassVar[int]
    option: _containers.RepeatedCompositeFieldContainer[OptionDetailMessage]
    def __init__(self, option: _Optional[_Iterable[_Union[OptionDetailMessage, _Mapping]]] = ...) -> None: ...

class StockDetailMessage(_message.Message):
    __slots__ = ("exchange", "category", "code", "name", "reference", "update_date", "day_trade")
    EXCHANGE_FIELD_NUMBER: _ClassVar[int]
    CATEGORY_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    REFERENCE_FIELD_NUMBER: _ClassVar[int]
    UPDATE_DATE_FIELD_NUMBER: _ClassVar[int]
    DAY_TRADE_FIELD_NUMBER: _ClassVar[int]
    exchange: str
    category: str
    code: str
    name: str
    reference: float
    update_date: str
    day_trade: str
    def __init__(self, exchange: _Optional[str] = ..., category: _Optional[str] = ..., code: _Optional[str] = ..., name: _Optional[str] = ..., reference: _Optional[float] = ..., update_date: _Optional[str] = ..., day_trade: _Optional[str] = ...) -> None: ...

class FutureDetailMessage(_message.Message):
    __slots__ = ("code", "symbol", "name", "category", "delivery_month", "delivery_date", "underlying_kind", "unit", "limit_up", "limit_down", "reference", "update_date")
    CODE_FIELD_NUMBER: _ClassVar[int]
    SYMBOL_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CATEGORY_FIELD_NUMBER: _ClassVar[int]
    DELIVERY_MONTH_FIELD_NUMBER: _ClassVar[int]
    DELIVERY_DATE_FIELD_NUMBER: _ClassVar[int]
    UNDERLYING_KIND_FIELD_NUMBER: _ClassVar[int]
    UNIT_FIELD_NUMBER: _ClassVar[int]
    LIMIT_UP_FIELD_NUMBER: _ClassVar[int]
    LIMIT_DOWN_FIELD_NUMBER: _ClassVar[int]
    REFERENCE_FIELD_NUMBER: _ClassVar[int]
    UPDATE_DATE_FIELD_NUMBER: _ClassVar[int]
    code: str
    symbol: str
    name: str
    category: str
    delivery_month: str
    delivery_date: str
    underlying_kind: str
    unit: int
    limit_up: float
    limit_down: float
    reference: float
    update_date: str
    def __init__(self, code: _Optional[str] = ..., symbol: _Optional[str] = ..., name: _Optional[str] = ..., category: _Optional[str] = ..., delivery_month: _Optional[str] = ..., delivery_date: _Optional[str] = ..., underlying_kind: _Optional[str] = ..., unit: _Optional[int] = ..., limit_up: _Optional[float] = ..., limit_down: _Optional[float] = ..., reference: _Optional[float] = ..., update_date: _Optional[str] = ...) -> None: ...

class OptionDetailMessage(_message.Message):
    __slots__ = ("code", "symbol", "name", "category", "delivery_month", "delivery_date", "strike_price", "option_right", "underlying_kind", "unit", "limit_up", "limit_down", "reference", "update_date")
    CODE_FIELD_NUMBER: _ClassVar[int]
    SYMBOL_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CATEGORY_FIELD_NUMBER: _ClassVar[int]
    DELIVERY_MONTH_FIELD_NUMBER: _ClassVar[int]
    DELIVERY_DATE_FIELD_NUMBER: _ClassVar[int]
    STRIKE_PRICE_FIELD_NUMBER: _ClassVar[int]
    OPTION_RIGHT_FIELD_NUMBER: _ClassVar[int]
    UNDERLYING_KIND_FIELD_NUMBER: _ClassVar[int]
    UNIT_FIELD_NUMBER: _ClassVar[int]
    LIMIT_UP_FIELD_NUMBER: _ClassVar[int]
    LIMIT_DOWN_FIELD_NUMBER: _ClassVar[int]
    REFERENCE_FIELD_NUMBER: _ClassVar[int]
    UPDATE_DATE_FIELD_NUMBER: _ClassVar[int]
    code: str
    symbol: str
    name: str
    category: str
    delivery_month: str
    delivery_date: str
    strike_price: float
    option_right: str
    underlying_kind: str
    unit: int
    limit_up: float
    limit_down: float
    reference: float
    update_date: str
    def __init__(self, code: _Optional[str] = ..., symbol: _Optional[str] = ..., name: _Optional[str] = ..., category: _Optional[str] = ..., delivery_month: _Optional[str] = ..., delivery_date: _Optional[str] = ..., strike_price: _Optional[float] = ..., option_right: _Optional[str] = ..., underlying_kind: _Optional[str] = ..., unit: _Optional[int] = ..., limit_up: _Optional[float] = ..., limit_down: _Optional[float] = ..., reference: _Optional[float] = ..., update_date: _Optional[str] = ...) -> None: ...
