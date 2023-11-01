#!/bin/sh

CURDATE=`date +%Y%m%d_%H%M%S`
echo $CURDATE

RELEASE_COUNT=`git branch | grep release | wc -l`

if [ $RELEASE_COUNT -eq 0 ] ; then
    kill -9 `ps -ef | grep retina-rapi-gin | grep -v grep | awk '{print $2}'`
else
    sudo kill -9 `ps -ef | grep retina-rapi-gin | grep -v grep | awk '{print $2}'`
fi

echo "[PS LIST] ---------------------------"
ps -ef | grep retina-rapi-gin
echo "-------------------------------------"

echo "end"
echo ""