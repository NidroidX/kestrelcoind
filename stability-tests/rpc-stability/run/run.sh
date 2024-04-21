#!/bin/bash
rm -rf /tmp/kestrelcoind-temp

kestrelcoind --devnet --appdir=/tmp/kestrelcoind-temp --profile=6061 --loglevel=debug &
kestrelcoinD_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $kestrelcoinD_PID

wait $kestrelcoinD_PID
kestrelcoinD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "kestrelcoind exit code: $kestrelcoinD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $kestrelcoinD_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
