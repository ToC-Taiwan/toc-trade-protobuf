#!/bin/bash

# find all *.py files
for file in $(find ./src/python/toc_trade_pb -name "*.py"); do
    sed -i '' 's/from forwarder import/from toc_trade_pb.forwarder import/g' $file
    sed -i '' 's/from app import/from toc_trade_pb.app import/g' $file
done
