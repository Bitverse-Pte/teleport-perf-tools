#!/bin/bash

for ((i=1; i<100; i++))
do
  curl -s 172.17.0.1:26657/commit?height=$i | jq "{time: .result.signed_header.header.time}"
done
