package main

import (
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"pault.ag/go/othername"
)

func main() {
	for _, el := range os.Args[1:] {
		fd, err := os.Open(el)
		if err != nil {
			panic(err)
		}
		defer fd.Close()
		bytes, err := ioutil.ReadAll(fd)
		if err != nil {
			panic(err)
		}
		block, _ := pem.Decode(bytes)
		_ = block
		cert, err := x509.ParseCertificate(bytes)
		if err != nil {
			continue
		}

		fmt.Printf("Serial: %x\n", cert.SerialNumber)
		fmt.Printf(" CN: %s\n", cert.Subject.CommonName)

		onames, err := othername.All(cert)
		if err != nil {
			panic(err)
		}
		for _, oname := range onames.Find(
			asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 6, 6},
		) {
			fasc, err := oname.FASCN()
			if err != nil {
				panic(err)
			}
			fmt.Printf(" FASC: %d\n", len(fasc))
		}

		names, err := othername.UPNs(cert)
		if err != nil {
			panic(err)
		}
		for _, name := range names {
			fmt.Printf(" UPN: %s\n", name)
		}

		for _, email := range cert.EmailAddresses {
			fmt.Printf(" Email: %s\n", email)
		}

		fmt.Printf("\n")
	}
}
