from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class OrderType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TYPE_UNKNOWN: _ClassVar[OrderType]
    TYPE_STOCK_LOT: _ClassVar[OrderType]
    TYPE_STOCK_SHARE: _ClassVar[OrderType]
    TYPE_FUTURE: _ClassVar[OrderType]
TYPE_UNKNOWN: OrderType
TYPE_STOCK_LOT: OrderType
TYPE_STOCK_SHARE: OrderType
TYPE_FUTURE: OrderType

class EventMessage(_message.Message):
    __slots__ = ("resp_code", "event_code", "info", "event", "event_time")
    RESP_CODE_FIELD_NUMBER: _ClassVar[int]
    EVENT_CODE_FIELD_NUMBER: _ClassVar[int]
    INFO_FIELD_NUMBER: _ClassVar[int]
    EVENT_FIELD_NUMBER: _ClassVar[int]
    EVENT_TIME_FIELD_NUMBER: _ClassVar[int]
    resp_code: int
    event_code: int
    info: str
    event: str
    event_time: str
    def __init__(self, resp_code: _Optional[int] = ..., event_code: _Optional[int] = ..., info: _Optional[str] = ..., event: _Optional[str] = ..., event_time: _Optional[str] = ...) -> None: ...

class StockRealTimeTickMessage(_message.Message):
    __slots__ = ("code", "date_time", "open", "avg_price", "close", "high", "low", "amount", "total_amount", "volume", "total_volume", "tick_type", "chg_type", "price_chg", "pct_chg", "bid_side_total_vol", "ask_side_total_vol", "bid_side_total_cnt", "ask_side_total_cnt", "suspend", "simtrade")
    CODE_FIELD_NUMBER: _ClassVar[int]
    DATE_TIME_FIELD_NUMBER: _ClassVar[int]
    OPEN_FIELD_NUMBER: _ClassVar[int]
    AVG_PRICE_FIELD_NUMBER: _ClassVar[int]
    CLOSE_FIELD_NUMBER: _ClassVar[int]
    HIGH_FIELD_NUMBER: _ClassVar[int]
    LOW_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    TOTAL_AMOUNT_FIELD_NUMBER: _ClassVar[int]
    VOLUME_FIELD_NUMBER: _ClassVar[int]
    TOTAL_VOLUME_FIELD_NUMBER: _ClassVar[int]
    TICK_TYPE_FIELD_NUMBER: _ClassVar[int]
    CHG_TYPE_FIELD_NUMBER: _ClassVar[int]
    PRICE_CHG_FIELD_NUMBER: _ClassVar[int]
    PCT_CHG_FIELD_NUMBER: _ClassVar[int]
    BID_SIDE_TOTAL_VOL_FIELD_NUMBER: _ClassVar[int]
    ASK_SIDE_TOTAL_VOL_FIELD_NUMBER: _ClassVar[int]
    BID_SIDE_TOTAL_CNT_FIELD_NUMBER: _ClassVar[int]
    ASK_SIDE_TOTAL_CNT_FIELD_NUMBER: _ClassVar[int]
    SUSPEND_FIELD_NUMBER: _ClassVar[int]
    SIMTRADE_FIELD_NUMBER: _ClassVar[int]
    code: str
    date_time: str
    open: float
    avg_price: float
    close: float
    high: float
    low: float
    amount: float
    total_amount: float
    volume: int
    total_volume: int
    tick_type: int
    chg_type: int
    price_chg: float
    pct_chg: float
    bid_side_total_vol: int
    ask_side_total_vol: int
    bid_side_total_cnt: int
    ask_side_total_cnt: int
    suspend: bool
    simtrade: bool
    def __init__(self, code: _Optional[str] = ..., date_time: _Optional[str] = ..., open: _Optional[float] = ..., avg_price: _Optional[float] = ..., close: _Optional[float] = ..., high: _Optional[float] = ..., low: _Optional[float] = ..., amount: _Optional[float] = ..., total_amount: _Optional[float] = ..., volume: _Optional[int] = ..., total_volume: _Optional[int] = ..., tick_type: _Optional[int] = ..., chg_type: _Optional[int] = ..., price_chg: _Optional[float] = ..., pct_chg: _Optional[float] = ..., bid_side_total_vol: _Optional[int] = ..., ask_side_total_vol: _Optional[int] = ..., bid_side_total_cnt: _Optional[int] = ..., ask_side_total_cnt: _Optional[int] = ..., suspend: bool = ..., simtrade: bool = ...) -> None: ...

class StockRealTimeBidAskMessage(_message.Message):
    __slots__ = ("code", "date_time", "bid_price", "bid_volume", "diff_bid_vol", "ask_price", "ask_volume", "diff_ask_vol", "suspend", "simtrade")
    CODE_FIELD_NUMBER: _ClassVar[int]
    DATE_TIME_FIELD_NUMBER: _ClassVar[int]
    BID_PRICE_FIELD_NUMBER: _ClassVar[int]
    BID_VOLUME_FIELD_NUMBER: _ClassVar[int]
    DIFF_BID_VOL_FIELD_NUMBER: _ClassVar[int]
    ASK_PRICE_FIELD_NUMBER: _ClassVar[int]
    ASK_VOLUME_FIELD_NUMBER: _ClassVar[int]
    DIFF_ASK_VOL_FIELD_NUMBER: _ClassVar[int]
    SUSPEND_FIELD_NUMBER: _ClassVar[int]
    SIMTRADE_FIELD_NUMBER: _ClassVar[int]
    code: str
    date_time: str
    bid_price: _containers.RepeatedScalarFieldContainer[float]
    bid_volume: _containers.RepeatedScalarFieldContainer[int]
    diff_bid_vol: _containers.RepeatedScalarFieldContainer[int]
    ask_price: _containers.RepeatedScalarFieldContainer[float]
    ask_volume: _containers.RepeatedScalarFieldContainer[int]
    diff_ask_vol: _containers.RepeatedScalarFieldContainer[int]
    suspend: bool
    simtrade: bool
    def __init__(self, code: _Optional[str] = ..., date_time: _Optional[str] = ..., bid_price: _Optional[_Iterable[float]] = ..., bid_volume: _Optional[_Iterable[int]] = ..., diff_bid_vol: _Optional[_Iterable[int]] = ..., ask_price: _Optional[_Iterable[float]] = ..., ask_volume: _Optional[_Iterable[int]] = ..., diff_ask_vol: _Optional[_Iterable[int]] = ..., suspend: bool = ..., simtrade: bool = ...) -> None: ...

class FutureRealTimeTickMessage(_message.Message):
    __slots__ = ("code", "date_time", "open", "underlying_price", "bid_side_total_vol", "ask_side_total_vol", "avg_price", "close", "high", "low", "amount", "total_amount", "volume", "total_volume", "tick_type", "chg_type", "price_chg", "pct_chg", "simtrade")
    CODE_FIELD_NUMBER: _ClassVar[int]
    DATE_TIME_FIELD_NUMBER: _ClassVar[int]
    OPEN_FIELD_NUMBER: _ClassVar[int]
    UNDERLYING_PRICE_FIELD_NUMBER: _ClassVar[int]
    BID_SIDE_TOTAL_VOL_FIELD_NUMBER: _ClassVar[int]
    ASK_SIDE_TOTAL_VOL_FIELD_NUMBER: _ClassVar[int]
    AVG_PRICE_FIELD_NUMBER: _ClassVar[int]
    CLOSE_FIELD_NUMBER: _ClassVar[int]
    HIGH_FIELD_NUMBER: _ClassVar[int]
    LOW_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    TOTAL_AMOUNT_FIELD_NUMBER: _ClassVar[int]
    VOLUME_FIELD_NUMBER: _ClassVar[int]
    TOTAL_VOLUME_FIELD_NUMBER: _ClassVar[int]
    TICK_TYPE_FIELD_NUMBER: _ClassVar[int]
    CHG_TYPE_FIELD_NUMBER: _ClassVar[int]
    PRICE_CHG_FIELD_NUMBER: _ClassVar[int]
    PCT_CHG_FIELD_NUMBER: _ClassVar[int]
    SIMTRADE_FIELD_NUMBER: _ClassVar[int]
    code: str
    date_time: str
    open: float
    underlying_price: float
    bid_side_total_vol: int
    ask_side_total_vol: int
    avg_price: float
    close: float
    high: float
    low: float
    amount: float
    total_amount: float
    volume: int
    total_volume: int
    tick_type: int
    chg_type: int
    price_chg: float
    pct_chg: float
    simtrade: bool
    def __init__(self, code: _Optional[str] = ..., date_time: _Optional[str] = ..., open: _Optional[float] = ..., underlying_price: _Optional[float] = ..., bid_side_total_vol: _Optional[int] = ..., ask_side_total_vol: _Optional[int] = ..., avg_price: _Optional[float] = ..., close: _Optional[float] = ..., high: _Optional[float] = ..., low: _Optional[float] = ..., amount: _Optional[float] = ..., total_amount: _Optional[float] = ..., volume: _Optional[int] = ..., total_volume: _Optional[int] = ..., tick_type: _Optional[int] = ..., chg_type: _Optional[int] = ..., price_chg: _Optional[float] = ..., pct_chg: _Optional[float] = ..., simtrade: bool = ...) -> None: ...

class FutureRealTimeBidAskMessage(_message.Message):
    __slots__ = ("code", "date_time", "bid_total_vol", "ask_total_vol", "bid_price", "bid_volume", "diff_bid_vol", "ask_price", "ask_volume", "diff_ask_vol", "first_derived_bid_price", "first_derived_ask_price", "first_derived_bid_vol", "first_derived_ask_vol", "underlying_price", "simtrade")
    CODE_FIELD_NUMBER: _ClassVar[int]
    DATE_TIME_FIELD_NUMBER: _ClassVar[int]
    BID_TOTAL_VOL_FIELD_NUMBER: _ClassVar[int]
    ASK_TOTAL_VOL_FIELD_NUMBER: _ClassVar[int]
    BID_PRICE_FIELD_NUMBER: _ClassVar[int]
    BID_VOLUME_FIELD_NUMBER: _ClassVar[int]
    DIFF_BID_VOL_FIELD_NUMBER: _ClassVar[int]
    ASK_PRICE_FIELD_NUMBER: _ClassVar[int]
    ASK_VOLUME_FIELD_NUMBER: _ClassVar[int]
    DIFF_ASK_VOL_FIELD_NUMBER: _ClassVar[int]
    FIRST_DERIVED_BID_PRICE_FIELD_NUMBER: _ClassVar[int]
    FIRST_DERIVED_ASK_PRICE_FIELD_NUMBER: _ClassVar[int]
    FIRST_DERIVED_BID_VOL_FIELD_NUMBER: _ClassVar[int]
    FIRST_DERIVED_ASK_VOL_FIELD_NUMBER: _ClassVar[int]
    UNDERLYING_PRICE_FIELD_NUMBER: _ClassVar[int]
    SIMTRADE_FIELD_NUMBER: _ClassVar[int]
    code: str
    date_time: str
    bid_total_vol: int
    ask_total_vol: int
    bid_price: _containers.RepeatedScalarFieldContainer[float]
    bid_volume: _containers.RepeatedScalarFieldContainer[int]
    diff_bid_vol: _containers.RepeatedScalarFieldContainer[int]
    ask_price: _containers.RepeatedScalarFieldContainer[float]
    ask_volume: _containers.RepeatedScalarFieldContainer[int]
    diff_ask_vol: _containers.RepeatedScalarFieldContainer[int]
    first_derived_bid_price: float
    first_derived_ask_price: float
    first_derived_bid_vol: int
    first_derived_ask_vol: int
    underlying_price: float
    simtrade: bool
    def __init__(self, code: _Optional[str] = ..., date_time: _Optional[str] = ..., bid_total_vol: _Optional[int] = ..., ask_total_vol: _Optional[int] = ..., bid_price: _Optional[_Iterable[float]] = ..., bid_volume: _Optional[_Iterable[int]] = ..., diff_bid_vol: _Optional[_Iterable[int]] = ..., ask_price: _Optional[_Iterable[float]] = ..., ask_volume: _Optional[_Iterable[int]] = ..., diff_ask_vol: _Optional[_Iterable[int]] = ..., first_derived_bid_price: _Optional[float] = ..., first_derived_ask_price: _Optional[float] = ..., first_derived_bid_vol: _Optional[int] = ..., first_derived_ask_vol: _Optional[int] = ..., underlying_price: _Optional[float] = ..., simtrade: bool = ...) -> None: ...

class OrderStatusArr(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[OrderStatus]
    def __init__(self, data: _Optional[_Iterable[_Union[OrderStatus, _Mapping]]] = ...) -> None: ...

class OrderStatus(_message.Message):
    __slots__ = ("type", "status", "code", "action", "price", "quantity", "order_id", "order_time")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    ORDER_ID_FIELD_NUMBER: _ClassVar[int]
    ORDER_TIME_FIELD_NUMBER: _ClassVar[int]
    type: OrderType
    status: str
    code: str
    action: str
    price: float
    quantity: int
    order_id: str
    order_time: str
    def __init__(self, type: _Optional[_Union[OrderType, str]] = ..., status: _Optional[str] = ..., code: _Optional[str] = ..., action: _Optional[str] = ..., price: _Optional[float] = ..., quantity: _Optional[int] = ..., order_id: _Optional[str] = ..., order_time: _Optional[str] = ...) -> None: ...
