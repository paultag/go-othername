go-othername
============

[![GoDoc](https://godoc.org/pault.ag/go/othername?status.svg)](https://godoc.org/pault.ag/go/othername)

Parse and export non-standard x.509
[Subject Alternative Name](https://en.wikipedia.org/wiki/Subject_Alternative_Name)
Other Names, as defined by `4.2.1.7` of RFC 3280.

Microsoft UPN
-------------

The Universal Principal Name is a Microsoft specific (but generally useful!)
Other Name type that allows the Certificate to map directly to a login to
a computer. This is widely deployed in conjuction with
[FIPS 201](https://en.wikipedia.org/wiki/FIPS_201) style smartcards to
enable computer logins.

```go

cert, err := x509.ParseCertificates(certDER)
...
names, err := othername.UPNs(cert)
...
```
