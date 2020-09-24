#!/bin/bash

while :
do
    ret=$(xdotool getmouselocation)
    x=$(echo $ret | sed 's/x:\([0-9]*\).*/\1/')
    y=$(echo $ret | sed 's/.*y:\([0-9]*\).*/\1/')
    echo "y:$y, x:$x"
    sleep 0.5
done


exit 0
