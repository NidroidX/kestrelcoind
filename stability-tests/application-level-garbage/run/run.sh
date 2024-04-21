#!/bin/bash
rm -rf /tmp/kestrelcoind-temp

kestrelcoind --devnet --appdir=/tmp/kestrelcoind-temp --profile=6061 --loglevel=debug &
kestrelcoinD_PID=$!
kestrelcoinD_KILLED=0
function killkestrelcoindIfNotKilled() {
    if [ $kestrelcoinD_KILLED -eq 0 ]; then
      kill $kestrelcoinD_PID
    fi
}
trap "killkestrelcoindIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:22611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $kestrelcoinD_PID

wait $kestrelcoinD_PID
kestrelcoinD_KILLED=1
kestrelcoinD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "kestrelcoind exit code: $kestrelcoinD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $kestrelcoinD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
