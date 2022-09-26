## MySQL

MySQLのアクセス方法や初期構築の手順などを記します。

## MySQL へログイン

Terraform で家の IPアドレスからMySQLにアクセスできるようにします。
ローカルのMySQLクライアントからDBにアクセスします。
お金がないので本番DBと検証DBは分けられていないです。

```
# Root
mysql --host prd-charalarm-mysql-database.cgm0celpahsu.ap-northeast-1.rds.amazonaws.com --port 3306 -u charalarm -p

# PRD
mysql --host prd-charalarm-mysql-database.cgm0celpahsu.ap-northeast-1.rds.amazonaws.com --port 3306 -u prd_charalarm -p

# STG
mysql --host prd-charalarm-mysql-database.cgm0celpahsu.ap-northeast-1.rds.amazonaws.com --port 3306 -u stg_charalarm -p
```


## データベース構築

データーベース構築。# 検証のDBが本番のDBに間借りする形になるので、検証DBにはプレフィックスを付与。

```
$ create database charalarm character set utf8mb4;
$ create database stg_charalarm character set utf8mb4;
```

Docker を使った方法
ローカルに MySQLクライアントを入れていますが、Dockerの MySQLクライアントを使ってもアクセスできます。 ただ、DockerのMySQLを使うと日本語が打てなくなってしまいました。良い感じでやる方法があれば教えてください！

```
# STG
docker run -it --rm -v mysql:/etc/mysql mysql mysql --host prd-charalarm-mysql-database.cgm0celpahsu.ap-northeast-1.rds.amazonaws.com --port 3306 -u stg_charalarm -p
```

データベース破棄

```
$ drop database charalarm;
$ drop database stg_charalarm;
```

ユーザー作成

```
# Prd User
$ create user prd_charalarm identified by 'PRD_MYSQL_PASSWORD';
$ create user prd_charalarm_read_only identified by 'PRD_MYSQL_READONLY_PASSWORD';

# Stg User
$ create user stg_charalarm identified by 'STG_MYSQL_PASSWORD';
$ create user stg_charalarm_read_only identified by 'STG_MYSQL_READONLY_PASSWORD';
```


権限付与

```
# Prd
$ grant all on charalarm.* TO prd_charalarm@'%';
$ grant select on charalarm.* TO prd_charalarm_read_only@'%';

# Stg
$ grant all on stg_charalarm.* TO stg_charalarm@'%';
$ grant select on stg_charalarm.* TO stg_charalarm_read_only@'%';
```

ユーザー一覧

```
SELECT * FROM mysql.user;
SELECT User FROM mysql.user;
```

ユーザー削除

```
DROP USER stg_charalarm;
DROP USER prd_charalarm;
```

接続確認

```
docker run -it --rm mysql mysql --host prd-mysql-database.czypddgoq5rs.ap-northeast-1.rds.amazonaws.com --port 3306 -u stg_tomoni -p
```

言語設定確認

```
show variables like 'char%';
```
