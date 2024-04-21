#!/bin/bash
rm -rf /tmp/kestrelcoind-temp

NUM_CLIENTS=128
kestrelcoind --devnet --appdir=/tmp/kestrelcoind-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
kestrelcoinD_PID=$!
kestrelcoinD_KILLED=0
function killkestrelcoindIfNotKilled() {
  if [ $kestrelcoinD_KILLED -eq 0 ]; then
    kill $kestrelcoinD_PID
  fi
}
trap "killkestrelcoindIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $kestrelcoinD_PID

wait $kestrelcoinD_PID
kestrelcoinD_EXIT_CODE=$?
kestrelcoinD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "kestrelcoind exit code: $kestrelcoinD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $kestrelcoinD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
