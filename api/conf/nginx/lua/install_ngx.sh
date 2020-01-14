yum install readline-devel pcre-devel openssl-devel
echo '[start install pcre!]'
cd /opt/downloads/
wget http://sourceforge.net/projects/pcre/files/pcre/8.32/pcre-8.32.tar.gz/download
tar zxvf pcre-8.32.tar.gz
cd pcre-8.32
./configure
make
make install
cd /lib64
ln -s libpcre.so.0.0.1 libpcre.so.1
ldconfig

wget http://openresty.org/download/ngx_openresty-1.2.8.6.tar.gz
tar zxf  ngx_openresty-1.2.8.6.tar.gz
cd ngx_openresty-1.2.8
./configure --prefix=/opt/nginx --with-luajit --with-pcre=/opt/downloads/pcre-8.30  --with-pcre-jit