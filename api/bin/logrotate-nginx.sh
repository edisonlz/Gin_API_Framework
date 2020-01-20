#!/bin/bash
# set file path
NGINX_ACCESS_LOG=/data/logs/nginx/access/log.pipe
NGINX_ERROR_LOG=/data/logs/nginx/error/log
NGINX_STATIS_LOG=/data/logs/nginx/statis/log
# rename log
mv $NGINX_ACCESS_LOG $NGINX_ACCESS_LOG.`date -d yesterday +%Y%m%d`
mv $NGINX_ERROR_LOG $NGINX_ERROR_LOG.`date -d yesterday +%Y%m%d`
mv $NGINX_STATIS_LOG $NGINX_STATIS_LOG.`date -d yesterday +%Y%m%d`
touch $NGINX_ACCESS_LOG
touch $NGINX_STATIS_LOG
/etc/init.d/syslog-ng restart
# restart nginx
#[ ! -f /opt/nginx/logs/nginx.pid ] || kill -USR1 $(cat /opt/nginx/logs/nginx.pid)
/etc/init.d/nginx reload

