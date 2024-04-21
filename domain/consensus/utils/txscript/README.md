txscript
========

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://godoc.org/github.com/NidroidX/kestrelcoind/txscript?status.png)](http://godoc.org/github.com/NidroidX/kestrelcoind/txscript)

Package txscript implements the kestrelcoin transaction script language. There is
a comprehensive test suite.

## kestrelcoin Scripts

kestrelcoin provides a stack-based, FORTH-like language for the scripts in
the kestrelcoin transactions. This language is not turing complete
although it is still fairly powerful. 

## Examples

* [Standard Pay-to-pubkey Script](http://godoc.org/github.com/NidroidX/kestrelcoind/txscript#example-PayToAddrScript)  
  Demonstrates creating a script which pays to a kestrelcoin address. It also
  prints the created script hex and uses the DisasmString function to display
  the disassembled script.

* [Extracting Details from Standard Scripts](http://godoc.org/github.com/NidroidX/kestrelcoind/txscript#example-ExtractPkScriptAddrs)  
  Demonstrates extracting information from a standard public key script.
