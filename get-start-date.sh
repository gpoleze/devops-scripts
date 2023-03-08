#!/bin/bash

export days_backwards="${1:-30}"

# get the day 20 days before today
OS=`uname`
if [ $OS == 'Darwin' ]; then
  argument="-v -${days_backwards}d"
fi
if [ $OS == 'Linux' ]; then
  argument="-d \"${days_backwards} day ago\""
fi

eval "date $argument +%F"