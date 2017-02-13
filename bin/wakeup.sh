#!/usr/bin/env bash

# 
# run wakeup-my-site
#
# last update: 2017.2.13
# 
# by hwj4477@gmail.com
#

URL="http://home.wjhong.pe.kr"

BIN_PATH="$HOME/go/src/wakeup-my-site"
export PATH=$BIN_PATH

exec wakeup-my-site $URL

