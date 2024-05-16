from forwarder import history_pb2 as _history_pb2
from forwarder import realtime_pb2 as _realtime_pb2
from forwarder import basic_pb2 as _basic_pb2
from forwarder import mq_pb2 as _mq_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class PickListType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TYPE_ADD: _ClassVar[PickListType]
    TYPE_REMOVE: _ClassVar[PickListType]
TYPE_ADD: PickListType
TYPE_REMOVE: PickListType

class PickRealMap(_message.Message):
    __slots__ = ("pick_map",)
    class PickMapEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: PickListType
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[PickListType, str]] = ...) -> None: ...
    PICK_MAP_FIELD_NUMBER: _ClassVar[int]
    pick_map: _containers.ScalarMap[str, PickListType]
    def __init__(self, pick_map: _Optional[_Mapping[str, PickListType]] = ...) -> None: ...

class PickFuture(_message.Message):
    __slots__ = ("code",)
    CODE_FIELD_NUMBER: _ClassVar[int]
    code: str
    def __init__(self, code: _Optional[str] = ...) -> None: ...

class WSMessage(_message.Message):
    __slots__ = ("future_detail", "future_tick", "history_kbar", "snapshot", "trade_index")
    FUTURE_DETAIL_FIELD_NUMBER: _ClassVar[int]
    FUTURE_TICK_FIELD_NUMBER: _ClassVar[int]
    HISTORY_KBAR_FIELD_NUMBER: _ClassVar[int]
    SNAPSHOT_FIELD_NUMBER: _ClassVar[int]
    TRADE_INDEX_FIELD_NUMBER: _ClassVar[int]
    future_detail: _basic_pb2.FutureDetailMessage
    future_tick: _mq_pb2.FutureRealTimeTickMessage
    history_kbar: _history_pb2.HistoryKbarResponse
    snapshot: _realtime_pb2.SnapshotMessage
    trade_index: TradeIndex
    def __init__(self, future_detail: _Optional[_Union[_basic_pb2.FutureDetailMessage, _Mapping]] = ..., future_tick: _Optional[_Union[_mq_pb2.FutureRealTimeTickMessage, _Mapping]] = ..., history_kbar: _Optional[_Union[_history_pb2.HistoryKbarResponse, _Mapping]] = ..., snapshot: _Optional[_Union[_realtime_pb2.SnapshotMessage, _Mapping]] = ..., trade_index: _Optional[_Union[TradeIndex, _Mapping]] = ...) -> None: ...

class TradeIndex(_message.Message):
    __slots__ = ("tse", "otc", "nasdaq", "nf")
    TSE_FIELD_NUMBER: _ClassVar[int]
    OTC_FIELD_NUMBER: _ClassVar[int]
    NASDAQ_FIELD_NUMBER: _ClassVar[int]
    NF_FIELD_NUMBER: _ClassVar[int]
    tse: IndexStatus
    otc: IndexStatus
    nasdaq: IndexStatus
    nf: IndexStatus
    def __init__(self, tse: _Optional[_Union[IndexStatus, _Mapping]] = ..., otc: _Optional[_Union[IndexStatus, _Mapping]] = ..., nasdaq: _Optional[_Union[IndexStatus, _Mapping]] = ..., nf: _Optional[_Union[IndexStatus, _Mapping]] = ...) -> None: ...

class IndexStatus(_message.Message):
    __slots__ = ("break_count", "price_chg")
    BREAK_COUNT_FIELD_NUMBER: _ClassVar[int]
    PRICE_CHG_FIELD_NUMBER: _ClassVar[int]
    break_count: int
    price_chg: float
    def __init__(self, break_count: _Optional[int] = ..., price_chg: _Optional[float] = ...) -> None: ...
