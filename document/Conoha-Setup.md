## Conoha サーバーの環境構築手順

サーバーに入る。パスワードは１パスワード保存

```
chmod 600 stg-charalarm-api-2020-08-26.pem
ssh -i stg-charalarm-api-2020-08-26.pem root@150.95.137.238
```

`.ssh` か1passwordに秘密鍵を保存してあるのでこちらを参照。

```
ssh stg-charalarm-api
ssh prd-charalarm-api
```

### Set motd

間違えて別のサーバーに入っても気がつけるようにログイン時にメッセージを表示します。
`/etc/motd` 以下のメッセージを設置します。上下に空行を入れるとめた目が良いかもです。

```
 ____  _              _             
/ ___|| |_ __ _  __ _(_)_ __   __ _ 
\___ \| __/ _` |/ _` | | '_ \ / _` |
 ___) | || (_| | (_| | | | | | (_| |
|____/ \__\__,_|\__, |_|_| |_|\__, |
                |___/         |___/ 
```

```
 ____                _            _   _             
|  _ \ _ __ ___   __| |_   _  ___| |_(_) ___  _ __  
| |_) | '__/ _ \ / _` | | | |/ __| __| |/ _ \| '_ \ 
|  __/| | | (_) | (_| | |_| | (__| |_| | (_) | | | |
|_|   |_|  \___/ \__,_|\__,_|\___|\__|_|\___/|_| |_|
```                                                 


### Install Nginx

```
apt-get -y update
apt-get -y upgrade
apt -y install nginx
systemctl enable nginx
systemctl start nginx
systemctl restart nginx
systemctl stop nginx
```

`/etc/nginx/conf.d/charalarm.conf`  にファイルを設置し、nginx を再起動する

```
server {
    listen 80;
    listen [::]:80;
    server_name api.charalarm.com;
    return 301 https://api.charalarm.com$request_uri;
}

server {
    listen 443;
    server_name api.charalarm.com;
    ssl on;
    ssl_certificate      /etc/letsencrypt/live/api.charalarm.com/fullchain.pem;
    ssl_certificate_key  /etc/letsencrypt/live/api.charalarm.com/privkey.pem;
    location / {
            proxy_set_header Host $http_host;
            proxy_pass http://localhost:8080;
    }
}
```

```
server {
    listen 80;
    listen [::]:80;
    server_name api-stg.charalarm.com;
    return 301 https://api-stg.charalarm.com$request_uri;
}

server {
    listen 443;
    server_name api.charalarm.com;
    ssl on;
    ssl_certificate      /etc/letsencrypt/live/api-stg.charalarm.com/fullchain.pem;
    ssl_certificate_key  /etc/letsencrypt/live/api-stg.charalarm.com/privkey.pem;
    location / {
            proxy_set_header Host $http_host;
            proxy_pass http://localhost:8080;
    }
}
```


### Install Docker

以下の公式の記事を参照すること

[Install Docker Engine on Ubuntu](https://docs.docker.com/engine/install/ubuntu/)

[Install Docker Compose](https://docs.docker.com/compose/install/)


### Install AWS CLI

公式の記事を参照する

[Installing the AWS CLI version 2 on Linux](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-linux.html)

ECRへのアクセス件を持つIAMユーザーを設定

```
aws configure
```

デフォルトのリージョンは `ap-northeast-1`


### Install Certbot

```
apt -y install certbot
```




