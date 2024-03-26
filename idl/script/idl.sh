#!/bin/bash

cd "$(dirname "$0")" || exit 1
cd ..

kitex -module github.com/Mutezebra/tiktok -gen-path ../kitex_gen base.thrift
kitex -module github.com/Mutezebra/tiktok -gen-path ../kitex_gen user.thrift
kitex -module github.com/Mutezebra/tiktok -gen-path ../kitex_gen video.thrift

exit 0
