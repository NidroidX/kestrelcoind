package util_test

import (
	"fmt"
	"github.com/NidroidX/kestrelcoind/util/difficulty"
	"math"
	"math/big"

	"github.com/NidroidX/kestrelcoind/util"
)

func ExampleAmount() {

	a := util.Amount(0)
	fmt.Println("Zero Sium:", a)

	a = util.Amount(1e8)
	fmt.Println("100,000,000 Sium:", a)

	a = util.Amount(1e5)
	fmt.Println("100,000 Sium:", a)
	// Output:
	// Zero Sium: 0 SDR
	// 100,000,000 Sium: 1 SDR
	// 100,000 Sium: 0.001 SDR
}

func ExampleNewAmount() {
	amountOne, err := util.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := util.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := util.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := util.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 SDR
	// 0.01234567 SDR
	// 0 SDR
	// invalid kestrelcoin amount
}

func ExampleAmount_unitConversions() {
	amount := util.Amount(44433324206900)

	fmt.Println("Sium to kSDR:", amount.Format(util.AmountKiloSDR))
	fmt.Println("Sium to SDR:", amount)
	fmt.Println("Sium to MilliSDR:", amount.Format(util.AmountMilliSDR))
	fmt.Println("Sium to MicroSDR:", amount.Format(util.AmountMicroSDR))
	fmt.Println("Sium to Sium:", amount.Format(util.AmountSium))

	// Output:
	// Sium to kSDR: 444.333242069 kSDR
	// Sium to SDR: 444333.242069 SDR
	// Sium to MilliSDR: 444333222.111 mSDR
	// Sium to MicroSDR: 444333242069 μSDR
	// Sium to Sium: 44433324206900 Sium
}

// This example demonstrates how to convert the compact "bits" in a block header
// which represent the target difficulty to a big integer and display it using
// the typical hex notation.
func ExampleCompactToBig() {
	bits := uint32(419465580)
	targetDifficulty := difficulty.CompactToBig(bits)

	// Display it in hex.
	fmt.Printf("%064x\n", targetDifficulty.Bytes())

	// Output:
	// 0000000000000000896c00000000000000000000000000000000000000000000
}

// This example demonstrates how to convert a target difficulty into the compact
// "bits" in a block header which represent that target difficulty .
func ExampleBigToCompact() {
	// Convert the target difficulty from block 300000 in the bitcoin
	// main chain to compact form.
	t := "0000000000000000896c00000000000000000000000000000000000000000000"
	targetDifficulty, success := new(big.Int).SetString(t, 16)
	if !success {
		fmt.Println("invalid target difficulty")
		return
	}
	bits := difficulty.BigToCompact(targetDifficulty)

	fmt.Println(bits)

	// Output:
	// 419465580
}
