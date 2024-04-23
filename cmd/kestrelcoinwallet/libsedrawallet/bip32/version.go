package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// kestrelcoinMainnetPrivate is the version that is used for
// kestrelcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var kestrelcoinMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// kestrelcoinMainnetPublic is the version that is used for
// kestrelcoin mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var kestrelcoinMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// kestrelcoinTestnetPrivate is the version that is used for
// kestrelcoin testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var kestrelcoinTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// kestrelcoinTestnetPublic is the version that is used for
// kestrelcoin testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var kestrelcoinTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// kestrelcoinDevnetPrivate is the version that is used for
// kestrelcoin devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var kestrelcoinDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// kestrelcoinDevnetPublic is the version that is used for
// kestrelcoin devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var kestrelcoinDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// kestrelcoinSimnetPrivate is the version that is used for
// kestrelcoin simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var kestrelcoinSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// kestrelcoinSimnetPublic is the version that is used for
// kestrelcoin simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var kestrelcoinSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case kestrelcoinMainnetPrivate:
		return kestrelcoinMainnetPublic, nil
	case kestrelcoinTestnetPrivate:
		return kestrelcoinTestnetPublic, nil
	case kestrelcoinDevnetPrivate:
		return kestrelcoinDevnetPublic, nil
	case kestrelcoinSimnetPrivate:
		return kestrelcoinSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case kestrelcoinMainnetPrivate:
		return true
	case kestrelcoinTestnetPrivate:
		return true
	case kestrelcoinDevnetPrivate:
		return true
	case kestrelcoinSimnetPrivate:
		return true
	}

	return false
}
