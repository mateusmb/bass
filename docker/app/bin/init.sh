#!/usr/bin/env bash
echo 'Running migrations...'
/bass/migrate up > /dev/null 2>&1 &

echo 'Deleting mysql-client...'
apk del mysql-client

echo 'Start appliation...'
/bass/app