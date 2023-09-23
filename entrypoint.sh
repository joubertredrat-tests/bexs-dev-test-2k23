#!/bin/sh

/go/src/app/wait-for db.ms.local:3306 -- echo "mariadb is up"
/go/src/app/app migrate
/go/src/app/app api
