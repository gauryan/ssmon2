#!/bin/sh

CURDATE=`date +%Y%m%d_%H%M%S`
echo $CURDATE

RELEASE_COUNT=`git branch | grep release | wc -l`

if [ $RELEASE_COUNT -eq 0 ] ; then
    nohup ./retina-rapi-gin &
else
    nohup sudo ./retina-rapi-gin &
fi

echo "[PS LIST] ---------------------------"
ps -ef | grep retina-rapi-gin
echo "-------------------------------------"

echo "starting....."
echo ""