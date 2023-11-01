# CRONTAB 설정
# 0 1 * * * sh /home/retina/retina_rapi_gin/log.sh

DATE_TIME=`date +'%y%m%d_%H%M%S'`

cp /home/retina/retina_rapi_gin/nohup.out /home/retina/retina_rapi_gin/nohup.out.${DATE_TIME}
> /home/retina/retina_rapi_gin/nohup.out