
#!/bin/bash
if [ ! -d "/data/downloads" ];then
 mkdir /data/downloads
fi
echo '[start install g++!]'
yum install gcc-c++.x86_64

yum install gcc
yum install gcc-c++
yum install openssl-devel
yum install readline-devel
yum install mysql-devel

echo '[start install pcre!]'
cd /data/downloads/
wget -O pcre-8.32.tar.gz http://sourceforge.net/projects/pcre/files/pcre/8.32/pcre-8.32.tar.gz/download
tar zxvf pcre-8.32.tar.gz
cd pcre-8.32
./configure
make
make install

cd /lib64/
ln -s libpcre.so.0.0.1 libpcre.so.1
ldconfig

echo '[start install zlib!]'
cd /data/downloads/
wget http://prdownloads.sourceforge.net/libpng/zlib-1.2.7.tar.gz?download
tar zxvf zlib-1.2.7.tar.gz?download
cd  zlib-1.2.7
./configure --prefix=/usr/local
make 
make install




cd /data/downloads/ 
wget http://www.lua.org/ftp/lua-5.1.2.tar.gz
tar zxvf lua-5.1.2.tar.gz
cd lua-5.1.2
make linux
make install

cd /data/downloads/ 
wget http://luajit.org/download/LuaJIT-2.0.0.tar.gz
tar zxvf LuaJIT-2.0.0.tar.gz
cd LuaJIT-2.0.0
make
make install

cd /data/downloads/ 
wget https://codeload.github.com/simpl/ngx_devel_kit/tar.gz/v0.2.17rc2
tar zxvf v0.2.17rc2

cd /data/downloads/ 
wget https://codeload.github.com/openresty/lua-nginx-module/tar.gz/v0.7.8
tar zxvf v0.7.8

export LUAJIT_LIB=/usr/local/lib 
export LUAJIT_INC=/usr/local/include/luajit-2.0

#vim .bash_profile 
#lua问题
#http://blog.sina.com.cn/u/1155571747

num=$( cat /root/.bash_profile|grep -c LUAJIT_LIB ) 
if [ $num = 0 ];then 
   echo "OK-O-O-O write hosts"
   read i
   echo 'export LUAJIT_LIB=/usr/local/lib'>> /root/.bash_profile
   echo 'export LUAJIT_INC=/usr/local/include/luajit-2.0'>> /root/.bash_profile
fi 


echo '[start install nginx-1.0.4!]'
cd /data/downloads/
wget http://nginx.org/download/nginx-1.0.4.tar.gz
tar zxvf nginx-1.0.4.tar.gz
cd nginx-1.0.4
./configure --prefix=/data/nginx --add-module=/data/downloads/ngx_devel_kit-0.2.17rc2 --add-module=/data/downloads/lua-nginx-module-0.7.8   --with-pcre=/data/downloads/pcre-8.32  --with-http_stub_status_module --with-http_ssl_module --with-cc-opt="-Wno-unused-but-set-variable"
make
make install



echo "make logs dir"
mkdir -p /data/logs/nginx/access
mkdir -p /data/logs/nginx/error
mkdir -p /data/logs/nginx/statis
mkdir -p /data/logs/tornado
mkdir -p /data/run


echo "start nginx"
ldconfig
/data/nginx/sbin/nginx -c /data/nginx/conf/nginx.conf

