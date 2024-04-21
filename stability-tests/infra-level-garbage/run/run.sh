#!/bin/bash
rm -rf /tmp/kestrelcoind-temp

kestrelcoind --devnet --appdir=/tmp/kestrelcoind-temp --profile=6061 &
kestrelcoinD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:22611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $kestrelcoinD_PID

wait $kestrelcoinD_PID
kestrelcoinD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "kestrelcoind exit code: $kestrelcoinD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $kestrelcoinD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
