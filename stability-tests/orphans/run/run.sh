#!/bin/bash
rm -rf /tmp/kestrelcoind-temp

kestrelcoind --simnet --appdir=/tmp/kestrelcoind-temp --profile=6061 &
kestrelcoinD_PID=$!

sleep 1

orphans --simnet -alocalhost:22511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $kestrelcoinD_PID

wait $kestrelcoinD_PID
kestrelcoinD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "kestrelcoind exit code: $kestrelcoinD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $kestrelcoinD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
