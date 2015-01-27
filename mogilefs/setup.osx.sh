#!/bin/sh

cpanm -f MogileFS::Server
cpanm MogileFS::Utils
cpanm DBD::mysql

mysql -u root -e 'create database if not exists mogilefs;'
mogdbsetup --yes --dbname=mogilefs --dbuser=root

mogstored -c ./mogilefs/mogstored.conf &
mogilefsd -c ./mogilefs/mogilefsd.conf &

mogadm --trackers=127.0.0.1:7001 host add localhost --ip=127.0.0.1 --port=7500 --status=alive
mogadm --trackers=127.0.0.1:7001 device add localhost 1

mogadm --trackers=127.0.0.1:7001 domain add go-mogilefs-client
mogadm --trackers=127.0.0.1:7001 class  add go-mogilefs-client test --mindevcount=1

mkdir -p mogilefs/var/mogdata/dev1

mogadm check

pkill -f mogilefsd
pkill -f mogstored
