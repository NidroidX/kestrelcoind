package common

import (
	"fmt"
	"github.com/NidroidX/kestrelcoind/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunkestrelcoindForTesting runs kestrelcoind for testing purposes
func RunkestrelcoindForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	kestrelcoindRunCommand, err := StartCmd("kestrelcoinD",
		"kestrelcoind",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("kestrelcoind started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := kestrelcoindRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("kestrelcoind closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := kestrelcoindRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("kestrelcoind stopped")
	}
}
