from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class FuturePosition(_message.Message):
    __slots__ = ("code", "direction", "quantity", "price", "last_price", "pnl")
    CODE_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    LAST_PRICE_FIELD_NUMBER: _ClassVar[int]
    PNL_FIELD_NUMBER: _ClassVar[int]
    code: str
    direction: str
    quantity: int
    price: float
    last_price: float
    pnl: float
    def __init__(self, code: _Optional[str] = ..., direction: _Optional[str] = ..., quantity: _Optional[int] = ..., price: _Optional[float] = ..., last_price: _Optional[float] = ..., pnl: _Optional[float] = ...) -> None: ...

class FuturePositionArr(_message.Message):
    __slots__ = ("position_arr",)
    POSITION_ARR_FIELD_NUMBER: _ClassVar[int]
    position_arr: _containers.RepeatedCompositeFieldContainer[FuturePosition]
    def __init__(self, position_arr: _Optional[_Iterable[_Union[FuturePosition, _Mapping]]] = ...) -> None: ...

class StockPosition(_message.Message):
    __slots__ = ("id", "code", "direction", "quantity", "price", "last_price", "pnl", "yd_quantity", "cond", "margin_purchase_amount", "collateral", "short_sale_margin", "interest", "detail_arr")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    LAST_PRICE_FIELD_NUMBER: _ClassVar[int]
    PNL_FIELD_NUMBER: _ClassVar[int]
    YD_QUANTITY_FIELD_NUMBER: _ClassVar[int]
    COND_FIELD_NUMBER: _ClassVar[int]
    MARGIN_PURCHASE_AMOUNT_FIELD_NUMBER: _ClassVar[int]
    COLLATERAL_FIELD_NUMBER: _ClassVar[int]
    SHORT_SALE_MARGIN_FIELD_NUMBER: _ClassVar[int]
    INTEREST_FIELD_NUMBER: _ClassVar[int]
    DETAIL_ARR_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    direction: str
    quantity: int
    price: float
    last_price: float
    pnl: float
    yd_quantity: int
    cond: str
    margin_purchase_amount: int
    collateral: int
    short_sale_margin: int
    interest: int
    detail_arr: _containers.RepeatedCompositeFieldContainer[StockPositionDetail]
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., direction: _Optional[str] = ..., quantity: _Optional[int] = ..., price: _Optional[float] = ..., last_price: _Optional[float] = ..., pnl: _Optional[float] = ..., yd_quantity: _Optional[int] = ..., cond: _Optional[str] = ..., margin_purchase_amount: _Optional[int] = ..., collateral: _Optional[int] = ..., short_sale_margin: _Optional[int] = ..., interest: _Optional[int] = ..., detail_arr: _Optional[_Iterable[_Union[StockPositionDetail, _Mapping]]] = ...) -> None: ...

class StockPositionArr(_message.Message):
    __slots__ = ("position_arr",)
    POSITION_ARR_FIELD_NUMBER: _ClassVar[int]
    position_arr: _containers.RepeatedCompositeFieldContainer[StockPosition]
    def __init__(self, position_arr: _Optional[_Iterable[_Union[StockPosition, _Mapping]]] = ...) -> None: ...

class StockPositionDetail(_message.Message):
    __slots__ = ("date", "code", "quantity", "price", "last_price", "dseq", "direction", "pnl", "currency", "fee")
    DATE_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    LAST_PRICE_FIELD_NUMBER: _ClassVar[int]
    DSEQ_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    PNL_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_FIELD_NUMBER: _ClassVar[int]
    FEE_FIELD_NUMBER: _ClassVar[int]
    date: str
    code: str
    quantity: int
    price: float
    last_price: float
    dseq: str
    direction: str
    pnl: float
    currency: str
    fee: float
    def __init__(self, date: _Optional[str] = ..., code: _Optional[str] = ..., quantity: _Optional[int] = ..., price: _Optional[float] = ..., last_price: _Optional[float] = ..., dseq: _Optional[str] = ..., direction: _Optional[str] = ..., pnl: _Optional[float] = ..., currency: _Optional[str] = ..., fee: _Optional[float] = ...) -> None: ...

class StockOrderDetail(_message.Message):
    __slots__ = ("stock_num", "price", "quantity", "simulate")
    STOCK_NUM_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    SIMULATE_FIELD_NUMBER: _ClassVar[int]
    stock_num: str
    price: float
    quantity: int
    simulate: bool
    def __init__(self, stock_num: _Optional[str] = ..., price: _Optional[float] = ..., quantity: _Optional[int] = ..., simulate: bool = ...) -> None: ...

class OddStockOrderDetail(_message.Message):
    __slots__ = ("stock_num", "price", "share")
    STOCK_NUM_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    SHARE_FIELD_NUMBER: _ClassVar[int]
    stock_num: str
    price: float
    share: int
    def __init__(self, stock_num: _Optional[str] = ..., price: _Optional[float] = ..., share: _Optional[int] = ...) -> None: ...

class FutureOrderDetail(_message.Message):
    __slots__ = ("code", "price", "quantity", "simulate")
    CODE_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    SIMULATE_FIELD_NUMBER: _ClassVar[int]
    code: str
    price: float
    quantity: int
    simulate: bool
    def __init__(self, code: _Optional[str] = ..., price: _Optional[float] = ..., quantity: _Optional[int] = ..., simulate: bool = ...) -> None: ...

class OptionOrderDetail(_message.Message):
    __slots__ = ("code", "price", "quantity", "simulate")
    CODE_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    SIMULATE_FIELD_NUMBER: _ClassVar[int]
    code: str
    price: float
    quantity: int
    simulate: bool
    def __init__(self, code: _Optional[str] = ..., price: _Optional[float] = ..., quantity: _Optional[int] = ..., simulate: bool = ...) -> None: ...

class TradeResult(_message.Message):
    __slots__ = ("order_id", "status", "error")
    ORDER_ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    order_id: str
    status: str
    error: str
    def __init__(self, order_id: _Optional[str] = ..., status: _Optional[str] = ..., error: _Optional[str] = ...) -> None: ...

class OrderID(_message.Message):
    __slots__ = ("order_id", "simulate")
    ORDER_ID_FIELD_NUMBER: _ClassVar[int]
    SIMULATE_FIELD_NUMBER: _ClassVar[int]
    order_id: str
    simulate: bool
    def __init__(self, order_id: _Optional[str] = ..., simulate: bool = ...) -> None: ...

class Margin(_message.Message):
    __slots__ = ("status", "yesterday_balance", "today_balance", "deposit_withdrawal", "fee", "tax", "initial_margin", "maintenance_margin", "margin_call", "risk_indicator", "royalty_revenue_expenditure", "equity", "equity_amount", "option_openbuy_market_value", "option_opensell_market_value", "option_open_position", "option_settle_profitloss", "future_open_position", "today_future_open_position", "future_settle_profitloss", "available_margin", "plus_margin", "plus_margin_indicator", "security_collateral_amount", "order_margin_premium", "collateral_amount")
    STATUS_FIELD_NUMBER: _ClassVar[int]
    YESTERDAY_BALANCE_FIELD_NUMBER: _ClassVar[int]
    TODAY_BALANCE_FIELD_NUMBER: _ClassVar[int]
    DEPOSIT_WITHDRAWAL_FIELD_NUMBER: _ClassVar[int]
    FEE_FIELD_NUMBER: _ClassVar[int]
    TAX_FIELD_NUMBER: _ClassVar[int]
    INITIAL_MARGIN_FIELD_NUMBER: _ClassVar[int]
    MAINTENANCE_MARGIN_FIELD_NUMBER: _ClassVar[int]
    MARGIN_CALL_FIELD_NUMBER: _ClassVar[int]
    RISK_INDICATOR_FIELD_NUMBER: _ClassVar[int]
    ROYALTY_REVENUE_EXPENDITURE_FIELD_NUMBER: _ClassVar[int]
    EQUITY_FIELD_NUMBER: _ClassVar[int]
    EQUITY_AMOUNT_FIELD_NUMBER: _ClassVar[int]
    OPTION_OPENBUY_MARKET_VALUE_FIELD_NUMBER: _ClassVar[int]
    OPTION_OPENSELL_MARKET_VALUE_FIELD_NUMBER: _ClassVar[int]
    OPTION_OPEN_POSITION_FIELD_NUMBER: _ClassVar[int]
    OPTION_SETTLE_PROFITLOSS_FIELD_NUMBER: _ClassVar[int]
    FUTURE_OPEN_POSITION_FIELD_NUMBER: _ClassVar[int]
    TODAY_FUTURE_OPEN_POSITION_FIELD_NUMBER: _ClassVar[int]
    FUTURE_SETTLE_PROFITLOSS_FIELD_NUMBER: _ClassVar[int]
    AVAILABLE_MARGIN_FIELD_NUMBER: _ClassVar[int]
    PLUS_MARGIN_FIELD_NUMBER: _ClassVar[int]
    PLUS_MARGIN_INDICATOR_FIELD_NUMBER: _ClassVar[int]
    SECURITY_COLLATERAL_AMOUNT_FIELD_NUMBER: _ClassVar[int]
    ORDER_MARGIN_PREMIUM_FIELD_NUMBER: _ClassVar[int]
    COLLATERAL_AMOUNT_FIELD_NUMBER: _ClassVar[int]
    status: str
    yesterday_balance: float
    today_balance: float
    deposit_withdrawal: float
    fee: float
    tax: float
    initial_margin: float
    maintenance_margin: float
    margin_call: float
    risk_indicator: float
    royalty_revenue_expenditure: float
    equity: float
    equity_amount: float
    option_openbuy_market_value: float
    option_opensell_market_value: float
    option_open_position: float
    option_settle_profitloss: float
    future_open_position: float
    today_future_open_position: float
    future_settle_profitloss: float
    available_margin: float
    plus_margin: float
    plus_margin_indicator: float
    security_collateral_amount: float
    order_margin_premium: float
    collateral_amount: float
    def __init__(self, status: _Optional[str] = ..., yesterday_balance: _Optional[float] = ..., today_balance: _Optional[float] = ..., deposit_withdrawal: _Optional[float] = ..., fee: _Optional[float] = ..., tax: _Optional[float] = ..., initial_margin: _Optional[float] = ..., maintenance_margin: _Optional[float] = ..., margin_call: _Optional[float] = ..., risk_indicator: _Optional[float] = ..., royalty_revenue_expenditure: _Optional[float] = ..., equity: _Optional[float] = ..., equity_amount: _Optional[float] = ..., option_openbuy_market_value: _Optional[float] = ..., option_opensell_market_value: _Optional[float] = ..., option_open_position: _Optional[float] = ..., option_settle_profitloss: _Optional[float] = ..., future_open_position: _Optional[float] = ..., today_future_open_position: _Optional[float] = ..., future_settle_profitloss: _Optional[float] = ..., available_margin: _Optional[float] = ..., plus_margin: _Optional[float] = ..., plus_margin_indicator: _Optional[float] = ..., security_collateral_amount: _Optional[float] = ..., order_margin_premium: _Optional[float] = ..., collateral_amount: _Optional[float] = ...) -> None: ...

class AccountBalance(_message.Message):
    __slots__ = ("date", "balance")
    DATE_FIELD_NUMBER: _ClassVar[int]
    BALANCE_FIELD_NUMBER: _ClassVar[int]
    date: str
    balance: float
    def __init__(self, date: _Optional[str] = ..., balance: _Optional[float] = ...) -> None: ...

class Settlement(_message.Message):
    __slots__ = ("date", "amount")
    DATE_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    date: str
    amount: float
    def __init__(self, date: _Optional[str] = ..., amount: _Optional[float] = ...) -> None: ...

class SettlementList(_message.Message):
    __slots__ = ("settlement",)
    SETTLEMENT_FIELD_NUMBER: _ClassVar[int]
    settlement: _containers.RepeatedCompositeFieldContainer[Settlement]
    def __init__(self, settlement: _Optional[_Iterable[_Union[Settlement, _Mapping]]] = ...) -> None: ...
