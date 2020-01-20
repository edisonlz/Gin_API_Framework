upstream make_app_api {
    server 127.0.0.1:10000 fail_timeout=0;
    server 127.0.0.1:10001 fail_timeout=0;
    server 127.0.0.1:10002 fail_timeout=0;
    server 127.0.0.1:10003 fail_timeout=0;
}


limit_req_zone $binary_remote_addr zone=one:100m rate=3r/s;

server {
    listen 80;
    server_name api.fastor.com;
    charset utf-8;

    set $x_remote_addr $proxy_add_x_forwarded_for;

    if ($x_remote_addr = "") {
        set $x_remote_addr $remote_addr;
    }

  
    location ~ \.(gif|jpg|png)$ {
        access_log off;
        proxy_pass          http://make_app_api;
        proxy_connect_timeout 3;
        proxy_send_timeout 3;
        proxy_read_timeout 3;
        proxy_redirect      default;
        proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header    X-Real-IP $remote_addr;
        proxy_set_header    Host $http_host;
        proxy_set_header    Range $http_range;
    }

    location ~ \.css$ {
        access_log off;
        proxy_pass          http://make_app_api;
        proxy_connect_timeout 3;
        proxy_send_timeout 3;
        proxy_read_timeout 3;
        proxy_redirect      default;
        proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header    X-Real-IP $remote_addr;
        proxy_set_header    Host $http_host;
        proxy_set_header    Range $http_range;
    }



    location ~ /doc {
        access_log off;
        proxy_pass          http://make_app_api;
        proxy_connect_timeout 3;
        proxy_send_timeout 3;
        proxy_read_timeout 3;
        proxy_redirect      default;
        proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header    X-Real-IP $remote_addr;
        proxy_set_header    Host $http_host;
        proxy_set_header    Range $http_range;
    }

    location ~ /sign  {
        access_log off;
        proxy_pass          http://make_app_api;
        proxy_connect_timeout 3;
        proxy_send_timeout 3;
        proxy_read_timeout 3;
        proxy_redirect      default;
        proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header    X-Real-IP $remote_addr;
        proxy_set_header    Host $http_host;
        proxy_set_header    Range $http_range;
    }

    location / {

        add_header Access-Control-Allow-Origin *;
        add_header Access-Control-Allow-Headers X-Requested-With;
        add_header Access-Control-Allow-Methsods GET,POST,OPTIONS;

        proxy_pass          http://make_app_api;
        proxy_connect_timeout 3;
        proxy_send_timeout 3;
        proxy_read_timeout 3;
        proxy_redirect      default;
        proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header    X-Real-IP $remote_addr;
        proxy_set_header    Host $http_host;
        proxy_set_header    Range $http_range;
    }
}




