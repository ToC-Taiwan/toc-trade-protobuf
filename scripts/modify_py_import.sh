#!/bin/bash

SEARCH_DIR=$1

if [ -z "$SEARCH_DIR" ]; then
    echo "Usage: $0 <search_dir>"
    exit 1
fi

# find all *.py files
for file in $(find $SEARCH_DIR -name "*.py"); do
    sed -i '' 's/from forwarder import/from toc_trade_pb.forwarder import/g' $file
    sed -i '' 's/from app import/from toc_trade_pb.app import/g' $file
done
