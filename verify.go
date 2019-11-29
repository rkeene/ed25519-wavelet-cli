package main

import ed25519 "github.com/perlin-network/noise/edwards25519"
import "encoding/hex"
import "fmt"
import "os"

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "usage: verify <pubKeyHex> <dataHex> <signatureHex>\n")
		os.Exit(3)
	}

	pubKeyStr    := []byte(os.Args[1])
	dataStr      := []byte(os.Args[2])
	signatureStr := []byte(os.Args[3])

	pubKeyBin    := make([]byte, hex.DecodedLen(len(pubKeyStr)))
	dataBin      := make([]byte, hex.DecodedLen(len(dataStr)))
	signatureBin := make([]byte, hex.DecodedLen(len(signatureStr)))

	if _, err := hex.Decode(pubKeyBin, pubKeyStr); err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode public key: %v\n", err)
		os.Exit(2)
	}
	if _, err := hex.Decode(dataBin, dataStr); err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode data: %v\n", err)
		os.Exit(2)
	}
	if _, err := hex.Decode(signatureBin, signatureStr); err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode signature: %v\n", err)
		os.Exit(2)
	}

	var pubKey [ed25519.SizePublicKey]byte
	var signature [ed25519.SizeSignature]byte

	copy(pubKey[:], pubKeyBin[:])
	copy(signature[:], signatureBin[:])

	valid := ed25519.Verify(pubKey, dataBin, signature)
	if valid {
		fmt.Println("SIGNATURE VALID")
		os.Exit(0)
	} else {
		fmt.Println("SIGNATURE INVALID")
		os.Exit(1)
	}

	return
}
