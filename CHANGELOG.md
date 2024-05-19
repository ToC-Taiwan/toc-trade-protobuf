# CHANGELOG

## [v0.1.10](https://github.com/ToC-Taiwan/toc-machine-trading-fe/compare/v0.1.9...v0.1.10)

> 2024-05-20

## [v0.1.9](https://github.com/ToC-Taiwan/toc-machine-trading-fe/compare/v0.1.7...v0.1.9)

> 2024-05-20

### Code Refactoring (1)

* split layout by language

## [v0.1.7](https://github.com/ToC-Taiwan/toc-machine-trading-fe/compare/v0.1.4...v0.1.7)

> 2024-05-19

### Features (1)

* add dart pre-compiled files

## [v0.1.4](https://github.com/ToC-Taiwan/toc-machine-trading-fe/compare/v0.1.0...v0.1.4)

> 2024-05-18

### Features (1)

* add ts, modify python generated import

## v0.1.0

> 2024-05-16

### Code Refactoring (16)

* refactor WebSocket message types and import proto files
* refactor stock snapshot methods to return SnapshotResponse
* refactor trade service and update proto dependencies
* refactor and simplify data interface methods
* refactor `app.proto` to remove unnecessary imports and messages
* refactor 'PickRealList' to 'PickRealMap' in message and field
* improve code readability and testing robustness
* refactor `PickRealList` message location and structure
* refactor `OrderType` enum for stock order simplification
* adjust `OrderType` enum values in `mq.proto`
* refactor trade.proto and enhance tests and APIs
* refactor `BasicDataInterface` service in `basic.proto`
* refactor data interface and improve test accuracy
* refactor data interface and improve test accuracy
* refactor `BasicDataInterface` service functions
* refactor `Login` function and remove `LoginResponse` definition

### Features (19)

* upload pre-compiled go and python, adjust makefile
* integrate forwarder snapshot into WebSocket message
* "Integrate PickFuture message into app.proto"
* refactor `TradeInterface` service to consolidate cancel functions
* "Integrate StockPositionDetail message into trade.proto"
* "Introduce `combo_count` field to `WSFutureTick` message"
* refactor PickRealList with PickListType enum and map
* implement new protobuf definitions for stock data
* refine `OrderType` enum and enhance test precision
* enhance compatibility and functionality across systems
* integrate `OrderType` enum into `OrderStatus` message
* refine ShioajiUsage message and update testing parameters
* refactor data interface for BeatMessage streaming
* "Implement long connection in BasicDataInterface service"
* refactor authentication services in BasicDataInterface
* implement `LogOut` method in `BasicDataInterface` service
* implement Shioaji usage tracking in BasicDataInterface
* add support for odd stock orders in TradeInterface service
* add 'odd' field to StockNumArr message
