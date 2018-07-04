package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/asn1"
	"encoding/hex"
	"fmt"

	"math/big"

	"github.com/facebookgo/inject"
	"github.com/studyzy/gotest/consensus"
	_ "github.com/studyzy/gotest/core"
	"github.com/studyzy/gotest/dag"
)

type Address [32]byte

func (a Address) Hex() string { return fmt.Sprintf("0x%x", a) }
func main() {
	fmt.Println("Hello World!")

	var dag = new(dag.Dag)
	var g inject.Graph
	var m consensus.Mediator
	g.Provide(
		&inject.Object{Value: &m},
		&inject.Object{Value: dag})
	g.Populate()

	fmt.Println(m.ReadSystemConfig("Count"))
	SignAndVerify()
}
func SignAndVerify() {
	pubkeyCurve := elliptic.P256() //see http://golang.org/pkg/crypto/elliptic/#P256

	//privatekey := new(ecdsa.PrivateKey)
	priv, _ := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

	hashed := []byte("testing")
	r, s, err := ecdsa.Sign(rand.Reader, priv, hashed)
	if err != nil {
		fmt.Errorf("%s: error signing: %s", err)
		return
	}
	signBytes := PointsToDER(r, s)
	fmt.Printf("Signature:{%x}\n", signBytes)

	pr, ps := PointsFromDER(decodeHex(signBytes))
	if !ecdsa.Verify(&priv.PublicKey, hashed, pr, ps) {
		fmt.Errorf("%s: Verify failed", hashed)
	} else {
		fmt.Println("Verify signature.")
	}

	hashed[0] ^= 0xff
	if ecdsa.Verify(&priv.PublicKey, hashed, r, s) {
		fmt.Errorf("%s: Verify always works!", hashed)
	}
}

// Convert an ECDSA signature (points R and S) to a byte array using ASN.1 DER encoding.
// This is a port of Bitcore's Key.rs2DER method.
func PointsToDER(r, s *big.Int) []byte {
	// Ensure MSB doesn't break big endian encoding in DER sigs
	prefixPoint := func(b []byte) []byte {
		if len(b) == 0 {
			b = []byte{0x00}
		}
		if b[0]&0x80 != 0 {
			paddedBytes := make([]byte, len(b)+1)
			copy(paddedBytes[1:], b)
			b = paddedBytes
		}
		return b
	}

	rb := prefixPoint(r.Bytes())
	sb := prefixPoint(s.Bytes())

	// DER encoding:
	// 0x30 + z + 0x02 + len(rb) + rb + 0x02 + len(sb) + sb
	length := 2 + len(rb) + 2 + len(sb)

	der := append([]byte{0x30, byte(length), 0x02, byte(len(rb))}, rb...)
	der = append(der, 0x02, byte(len(sb)))
	der = append(der, sb...)

	encoded := make([]byte, hex.EncodedLen(len(der)))
	hex.Encode(encoded, der)

	return encoded
}

// Get the X and Y points from a DER encoded signature
// Sometimes demarshalling using Golang's DEC to struct unmarshalling fails; this extracts R and S from the bytes
// manually to prevent crashing.
// This should NOT be a hex encoded byte array
func PointsFromDER(der []byte) (R, S *big.Int) {
	R, S = &big.Int{}, &big.Int{}

	data := asn1.RawValue{}
	if _, err := asn1.Unmarshal(der, &data); err != nil {
		panic(err.Error())
	}

	// The format of our DER string is 0x02 + rlen + r + 0x02 + slen + s
	rLen := data.Bytes[1] // The entire length of R + offset of 2 for 0x02 and rlen
	r := data.Bytes[2 : rLen+2]
	// Ignore the next 0x02 and slen bytes and just take the start of S to the end of the byte array
	s := data.Bytes[rLen+4:]

	R.SetBytes(r)
	S.SetBytes(s)

	return
}
func decodeHex(data []byte) []byte {
	var decoded = make([]byte, hex.DecodedLen(len(data)))
	hex.Decode(decoded, data)
	return decoded
}
