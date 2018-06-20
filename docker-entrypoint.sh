#! /bin/sh

function term()
{
    kill -15 $child
    wait $child
}

trap term SIGTERM

exec "`sgmock --port 9001 --key ${SGMOCK_KEY}`" &

child=$!
wait $child