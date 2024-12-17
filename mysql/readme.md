### Notes

```sh
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql
mysql -h 127.0.0.1 -P 3306 -u root -p

create database recordings;
use recordings;
source /path/to/recordings.sql
## drivers https://go.dev/wiki/SQLDrivers
## export DBUSER=root; export DBPASS=root

```