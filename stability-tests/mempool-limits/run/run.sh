#!/bin/bash

APPDIR=/tmp/kestrelcoind-temp
kestrelcoinD_RPC_PORT=29587

rm -rf "${APPDIR}"

kestrelcoind --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${kestrelcoinD_RPC_PORT}" --profile=6061 &
kestrelcoinD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${kestrelcoinD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $kestrelcoinD_PID

wait $kestrelcoinD_PID
kestrelcoinD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "kestrelcoind exit code: $kestrelcoinD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $kestrelcoinD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
