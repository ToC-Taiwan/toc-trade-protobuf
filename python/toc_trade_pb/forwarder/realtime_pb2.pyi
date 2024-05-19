from google.protobuf import empty_pb2 as _empty_pb2
from forwarder import entity_pb2 as _entity_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SnapshotResponse(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[SnapshotMessage]
    def __init__(self, data: _Optional[_Iterable[_Union[SnapshotMessage, _Mapping]]] = ...) -> None: ...

class YahooFinancePrice(_message.Message):
    __slots__ = ("price", "last")
    PRICE_FIELD_NUMBER: _ClassVar[int]
    LAST_FIELD_NUMBER: _ClassVar[int]
    price: float
    last: float
    def __init__(self, price: _Optional[float] = ..., last: _Optional[float] = ...) -> None: ...

class VolumeRankRequest(_message.Message):
    __slots__ = ("count", "date")
    COUNT_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    count: int
    date: str
    def __init__(self, count: _Optional[int] = ..., date: _Optional[str] = ...) -> None: ...

class SnapshotMessage(_message.Message):
    __slots__ = ("ts", "code", "exchange", "open", "high", "low", "close", "tick_type", "change_price", "change_rate", "change_type", "average_price", "volume", "total_volume", "amount", "total_amount", "yesterday_volume", "buy_price", "buy_volume", "sell_price", "sell_volume", "volume_ratio")
    TS_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    EXCHANGE_FIELD_NUMBER: _ClassVar[int]
    OPEN_FIELD_NUMBER: _ClassVar[int]
    HIGH_FIELD_NUMBER: _ClassVar[int]
    LOW_FIELD_NUMBER: _ClassVar[int]
    CLOSE_FIELD_NUMBER: _ClassVar[int]
    TICK_TYPE_FIELD_NUMBER: _ClassVar[int]
    CHANGE_PRICE_FIELD_NUMBER: _ClassVar[int]
    CHANGE_RATE_FIELD_NUMBER: _ClassVar[int]
    CHANGE_TYPE_FIELD_NUMBER: _ClassVar[int]
    AVERAGE_PRICE_FIELD_NUMBER: _ClassVar[int]
    VOLUME_FIELD_NUMBER: _ClassVar[int]
    TOTAL_VOLUME_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    TOTAL_AMOUNT_FIELD_NUMBER: _ClassVar[int]
    YESTERDAY_VOLUME_FIELD_NUMBER: _ClassVar[int]
    BUY_PRICE_FIELD_NUMBER: _ClassVar[int]
    BUY_VOLUME_FIELD_NUMBER: _ClassVar[int]
    SELL_PRICE_FIELD_NUMBER: _ClassVar[int]
    SELL_VOLUME_FIELD_NUMBER: _ClassVar[int]
    VOLUME_RATIO_FIELD_NUMBER: _ClassVar[int]
    ts: int
    code: str
    exchange: str
    open: float
    high: float
    low: float
    close: float
    tick_type: str
    change_price: float
    change_rate: float
    change_type: str
    average_price: float
    volume: int
    total_volume: int
    amount: int
    total_amount: int
    yesterday_volume: float
    buy_price: float
    buy_volume: float
    sell_price: float
    sell_volume: int
    volume_ratio: float
    def __init__(self, ts: _Optional[int] = ..., code: _Optional[str] = ..., exchange: _Optional[str] = ..., open: _Optional[float] = ..., high: _Optional[float] = ..., low: _Optional[float] = ..., close: _Optional[float] = ..., tick_type: _Optional[str] = ..., change_price: _Optional[float] = ..., change_rate: _Optional[float] = ..., change_type: _Optional[str] = ..., average_price: _Optional[float] = ..., volume: _Optional[int] = ..., total_volume: _Optional[int] = ..., amount: _Optional[int] = ..., total_amount: _Optional[int] = ..., yesterday_volume: _Optional[float] = ..., buy_price: _Optional[float] = ..., buy_volume: _Optional[float] = ..., sell_price: _Optional[float] = ..., sell_volume: _Optional[int] = ..., volume_ratio: _Optional[float] = ...) -> None: ...

class StockVolumeRankResponse(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[StockVolumeRankMessage]
    def __init__(self, data: _Optional[_Iterable[_Union[StockVolumeRankMessage, _Mapping]]] = ...) -> None: ...

class StockVolumeRankMessage(_message.Message):
    __slots__ = ("date", "code", "name", "ts", "open", "high", "low", "close", "price_range", "tick_type", "change_price", "change_type", "average_price", "volume", "total_volume", "amount", "total_amount", "yesterday_volume", "volume_ratio", "buy_price", "buy_volume", "sell_price", "sell_volume", "bid_orders", "bid_volumes", "ask_orders", "ask_volumes")
    DATE_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    TS_FIELD_NUMBER: _ClassVar[int]
    OPEN_FIELD_NUMBER: _ClassVar[int]
    HIGH_FIELD_NUMBER: _ClassVar[int]
    LOW_FIELD_NUMBER: _ClassVar[int]
    CLOSE_FIELD_NUMBER: _ClassVar[int]
    PRICE_RANGE_FIELD_NUMBER: _ClassVar[int]
    TICK_TYPE_FIELD_NUMBER: _ClassVar[int]
    CHANGE_PRICE_FIELD_NUMBER: _ClassVar[int]
    CHANGE_TYPE_FIELD_NUMBER: _ClassVar[int]
    AVERAGE_PRICE_FIELD_NUMBER: _ClassVar[int]
    VOLUME_FIELD_NUMBER: _ClassVar[int]
    TOTAL_VOLUME_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    TOTAL_AMOUNT_FIELD_NUMBER: _ClassVar[int]
    YESTERDAY_VOLUME_FIELD_NUMBER: _ClassVar[int]
    VOLUME_RATIO_FIELD_NUMBER: _ClassVar[int]
    BUY_PRICE_FIELD_NUMBER: _ClassVar[int]
    BUY_VOLUME_FIELD_NUMBER: _ClassVar[int]
    SELL_PRICE_FIELD_NUMBER: _ClassVar[int]
    SELL_VOLUME_FIELD_NUMBER: _ClassVar[int]
    BID_ORDERS_FIELD_NUMBER: _ClassVar[int]
    BID_VOLUMES_FIELD_NUMBER: _ClassVar[int]
    ASK_ORDERS_FIELD_NUMBER: _ClassVar[int]
    ASK_VOLUMES_FIELD_NUMBER: _ClassVar[int]
    date: str
    code: str
    name: str
    ts: int
    open: float
    high: float
    low: float
    close: float
    price_range: float
    tick_type: int
    change_price: float
    change_type: int
    average_price: float
    volume: int
    total_volume: int
    amount: int
    total_amount: int
    yesterday_volume: int
    volume_ratio: float
    buy_price: float
    buy_volume: int
    sell_price: float
    sell_volume: int
    bid_orders: int
    bid_volumes: int
    ask_orders: int
    ask_volumes: int
    def __init__(self, date: _Optional[str] = ..., code: _Optional[str] = ..., name: _Optional[str] = ..., ts: _Optional[int] = ..., open: _Optional[float] = ..., high: _Optional[float] = ..., low: _Optional[float] = ..., close: _Optional[float] = ..., price_range: _Optional[float] = ..., tick_type: _Optional[int] = ..., change_price: _Optional[float] = ..., change_type: _Optional[int] = ..., average_price: _Optional[float] = ..., volume: _Optional[int] = ..., total_volume: _Optional[int] = ..., amount: _Optional[int] = ..., total_amount: _Optional[int] = ..., yesterday_volume: _Optional[int] = ..., volume_ratio: _Optional[float] = ..., buy_price: _Optional[float] = ..., buy_volume: _Optional[int] = ..., sell_price: _Optional[float] = ..., sell_volume: _Optional[int] = ..., bid_orders: _Optional[int] = ..., bid_volumes: _Optional[int] = ..., ask_orders: _Optional[int] = ..., ask_volumes: _Optional[int] = ...) -> None: ...
