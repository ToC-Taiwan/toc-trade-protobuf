# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: forwarder/trade.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'forwarder/trade.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15\x66orwarder/trade.proto\x12\tforwarder\x1a\x1bgoogle/protobuf/empty.proto\"s\n\x0e\x46uturePosition\x12\x0c\n\x04\x63ode\x18\x01 \x01(\t\x12\x11\n\tdirection\x18\x02 \x01(\t\x12\x10\n\x08quantity\x18\x03 \x01(\x05\x12\r\n\x05price\x18\x04 \x01(\x01\x12\x12\n\nlast_price\x18\x05 \x01(\x01\x12\x0b\n\x03pnl\x18\x06 \x01(\x01\"D\n\x11\x46uturePositionArr\x12/\n\x0cposition_arr\x18\x01 \x03(\x0b\x32\x19.forwarder.FuturePosition\"\xb6\x02\n\rStockPosition\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04\x63ode\x18\x02 \x01(\t\x12\x11\n\tdirection\x18\x03 \x01(\t\x12\x10\n\x08quantity\x18\x04 \x01(\x05\x12\r\n\x05price\x18\x05 \x01(\x01\x12\x12\n\nlast_price\x18\x06 \x01(\x01\x12\x0b\n\x03pnl\x18\x07 \x01(\x01\x12\x13\n\x0byd_quantity\x18\x08 \x01(\x05\x12\x0c\n\x04\x63ond\x18\t \x01(\t\x12\x1e\n\x16margin_purchase_amount\x18\n \x01(\x05\x12\x12\n\ncollateral\x18\x0b \x01(\x05\x12\x19\n\x11short_sale_margin\x18\x0c \x01(\x05\x12\x10\n\x08interest\x18\r \x01(\x05\x12\x32\n\ndetail_arr\x18\x0e \x03(\x0b\x32\x1e.forwarder.StockPositionDetail\"B\n\x10StockPositionArr\x12.\n\x0cposition_arr\x18\x01 \x03(\x0b\x32\x18.forwarder.StockPosition\"\xb3\x01\n\x13StockPositionDetail\x12\x0c\n\x04\x64\x61te\x18\x01 \x01(\t\x12\x0c\n\x04\x63ode\x18\x02 \x01(\t\x12\x10\n\x08quantity\x18\x03 \x01(\x05\x12\r\n\x05price\x18\x04 \x01(\x01\x12\x12\n\nlast_price\x18\x05 \x01(\x01\x12\x0c\n\x04\x64seq\x18\x06 \x01(\t\x12\x11\n\tdirection\x18\x07 \x01(\t\x12\x0b\n\x03pnl\x18\x08 \x01(\x01\x12\x10\n\x08\x63urrency\x18\t \x01(\t\x12\x0b\n\x03\x66\x65\x65\x18\n \x01(\x01\"X\n\x10StockOrderDetail\x12\x11\n\tstock_num\x18\x01 \x01(\t\x12\r\n\x05price\x18\x02 \x01(\x01\x12\x10\n\x08quantity\x18\x03 \x01(\x03\x12\x10\n\x08simulate\x18\x04 \x01(\x08\"F\n\x13OddStockOrderDetail\x12\x11\n\tstock_num\x18\x01 \x01(\t\x12\r\n\x05price\x18\x02 \x01(\x01\x12\r\n\x05share\x18\x03 \x01(\x03\"T\n\x11\x46utureOrderDetail\x12\x0c\n\x04\x63ode\x18\x01 \x01(\t\x12\r\n\x05price\x18\x02 \x01(\x01\x12\x10\n\x08quantity\x18\x03 \x01(\x03\x12\x10\n\x08simulate\x18\x04 \x01(\x08\"T\n\x11OptionOrderDetail\x12\x0c\n\x04\x63ode\x18\x01 \x01(\t\x12\r\n\x05price\x18\x02 \x01(\x01\x12\x10\n\x08quantity\x18\x03 \x01(\x03\x12\x10\n\x08simulate\x18\x04 \x01(\x08\">\n\x0bTradeResult\x12\x10\n\x08order_id\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\t\x12\r\n\x05\x65rror\x18\x03 \x01(\t\"-\n\x07OrderID\x12\x10\n\x08order_id\x18\x01 \x01(\t\x12\x10\n\x08simulate\x18\x02 \x01(\x08\"\xc7\x05\n\x06Margin\x12\x0e\n\x06status\x18\x01 \x01(\t\x12\x19\n\x11yesterday_balance\x18\x02 \x01(\x01\x12\x15\n\rtoday_balance\x18\x03 \x01(\x01\x12\x1a\n\x12\x64\x65posit_withdrawal\x18\x04 \x01(\x01\x12\x0b\n\x03\x66\x65\x65\x18\x05 \x01(\x01\x12\x0b\n\x03tax\x18\x06 \x01(\x01\x12\x16\n\x0einitial_margin\x18\x07 \x01(\x01\x12\x1a\n\x12maintenance_margin\x18\x08 \x01(\x01\x12\x13\n\x0bmargin_call\x18\t \x01(\x01\x12\x16\n\x0erisk_indicator\x18\n \x01(\x01\x12#\n\x1broyalty_revenue_expenditure\x18\x0b \x01(\x01\x12\x0e\n\x06\x65quity\x18\x0c \x01(\x01\x12\x15\n\requity_amount\x18\r \x01(\x01\x12#\n\x1boption_openbuy_market_value\x18\x0e \x01(\x01\x12$\n\x1coption_opensell_market_value\x18\x0f \x01(\x01\x12\x1c\n\x14option_open_position\x18\x10 \x01(\x01\x12 \n\x18option_settle_profitloss\x18\x11 \x01(\x01\x12\x1c\n\x14\x66uture_open_position\x18\x12 \x01(\x01\x12\"\n\x1atoday_future_open_position\x18\x13 \x01(\x01\x12 \n\x18\x66uture_settle_profitloss\x18\x14 \x01(\x01\x12\x18\n\x10\x61vailable_margin\x18\x15 \x01(\x01\x12\x13\n\x0bplus_margin\x18\x16 \x01(\x01\x12\x1d\n\x15plus_margin_indicator\x18\x17 \x01(\x01\x12\"\n\x1asecurity_collateral_amount\x18\x18 \x01(\x01\x12\x1c\n\x14order_margin_premium\x18\x19 \x01(\x01\x12\x19\n\x11\x63ollateral_amount\x18\x1a \x01(\x01\"/\n\x0e\x41\x63\x63ountBalance\x12\x0c\n\x04\x64\x61te\x18\x01 \x01(\t\x12\x0f\n\x07\x62\x61lance\x18\x02 \x01(\x01\"*\n\nSettlement\x12\x0c\n\x04\x64\x61te\x18\x01 \x01(\t\x12\x0e\n\x06\x61mount\x18\x02 \x01(\x01\";\n\x0eSettlementList\x12)\n\nsettlement\x18\x01 \x03(\x0b\x32\x15.forwarder.Settlement2\xd9\n\n\x0eTradeInterface\x12;\n\x0b\x43\x61ncelOrder\x12\x12.forwarder.OrderID\x1a\x16.forwarder.TradeResult\"\x00\x12\x41\n\x08\x42uyStock\x12\x1b.forwarder.StockOrderDetail\x1a\x16.forwarder.TradeResult\"\x00\x12\x42\n\tSellStock\x12\x1b.forwarder.StockOrderDetail\x1a\x16.forwarder.TradeResult\"\x00\x12G\n\x0b\x42uyOddStock\x12\x1e.forwarder.OddStockOrderDetail\x1a\x16.forwarder.TradeResult\"\x00\x12H\n\x0cSellOddStock\x12\x1e.forwarder.OddStockOrderDetail\x1a\x16.forwarder.TradeResult\"\x00\x12G\n\x0eSellFirstStock\x12\x1b.forwarder.StockOrderDetail\x1a\x16.forwarder.TradeResult\"\x00\x12\x43\n\tBuyFuture\x12\x1c.forwarder.FutureOrderDetail\x1a\x16.forwarder.TradeResult\"\x00\x12\x44\n\nSellFuture\x12\x1c.forwarder.FutureOrderDetail\x1a\x16.forwarder.TradeResult\"\x00\x12I\n\x0fSellFirstFuture\x12\x1c.forwarder.FutureOrderDetail\x1a\x16.forwarder.TradeResult\"\x00\x12\x43\n\tBuyOption\x12\x1c.forwarder.OptionOrderDetail\x1a\x16.forwarder.TradeResult\"\x00\x12\x44\n\nSellOption\x12\x1c.forwarder.OptionOrderDetail\x1a\x16.forwarder.TradeResult\"\x00\x12I\n\x0fSellFirstOption\x12\x1c.forwarder.OptionOrderDetail\x1a\x16.forwarder.TradeResult\"\x00\x12J\n\x16GetLocalOrderStatusArr\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\"\x00\x12M\n\x19GetSimulateOrderStatusArr\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\"\x00\x12K\n\x11GetFuturePosition\x12\x16.google.protobuf.Empty\x1a\x1c.forwarder.FuturePositionArr\"\x00\x12I\n\x10GetStockPosition\x12\x16.google.protobuf.Empty\x1a\x1b.forwarder.StockPositionArr\"\x00\x12\x44\n\rGetSettlement\x12\x16.google.protobuf.Empty\x1a\x19.forwarder.SettlementList\"\x00\x12H\n\x11GetAccountBalance\x12\x16.google.protobuf.Empty\x1a\x19.forwarder.AccountBalance\"\x00\x12\x38\n\tGetMargin\x12\x16.google.protobuf.Empty\x1a\x11.forwarder.Margin\"\x00\x42\x06Z\x04./pbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'forwarder.trade_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\004./pb'
  _globals['_FUTUREPOSITION']._serialized_start=65
  _globals['_FUTUREPOSITION']._serialized_end=180
  _globals['_FUTUREPOSITIONARR']._serialized_start=182
  _globals['_FUTUREPOSITIONARR']._serialized_end=250
  _globals['_STOCKPOSITION']._serialized_start=253
  _globals['_STOCKPOSITION']._serialized_end=563
  _globals['_STOCKPOSITIONARR']._serialized_start=565
  _globals['_STOCKPOSITIONARR']._serialized_end=631
  _globals['_STOCKPOSITIONDETAIL']._serialized_start=634
  _globals['_STOCKPOSITIONDETAIL']._serialized_end=813
  _globals['_STOCKORDERDETAIL']._serialized_start=815
  _globals['_STOCKORDERDETAIL']._serialized_end=903
  _globals['_ODDSTOCKORDERDETAIL']._serialized_start=905
  _globals['_ODDSTOCKORDERDETAIL']._serialized_end=975
  _globals['_FUTUREORDERDETAIL']._serialized_start=977
  _globals['_FUTUREORDERDETAIL']._serialized_end=1061
  _globals['_OPTIONORDERDETAIL']._serialized_start=1063
  _globals['_OPTIONORDERDETAIL']._serialized_end=1147
  _globals['_TRADERESULT']._serialized_start=1149
  _globals['_TRADERESULT']._serialized_end=1211
  _globals['_ORDERID']._serialized_start=1213
  _globals['_ORDERID']._serialized_end=1258
  _globals['_MARGIN']._serialized_start=1261
  _globals['_MARGIN']._serialized_end=1972
  _globals['_ACCOUNTBALANCE']._serialized_start=1974
  _globals['_ACCOUNTBALANCE']._serialized_end=2021
  _globals['_SETTLEMENT']._serialized_start=2023
  _globals['_SETTLEMENT']._serialized_end=2065
  _globals['_SETTLEMENTLIST']._serialized_start=2067
  _globals['_SETTLEMENTLIST']._serialized_end=2126
  _globals['_TRADEINTERFACE']._serialized_start=2129
  _globals['_TRADEINTERFACE']._serialized_end=3498
# @@protoc_insertion_point(module_scope)
