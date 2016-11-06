package lib

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
    "io/ioutil"
    "path/filepath"
	"fmt"
	"log"
    "net"
	"math/big"
	"os"
	"time"
)

var pathSep = string(filepath.Separator)

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}

func GenerateKeyPair(host string) (string, string) {
	var priv interface{}
	var err error

    priv, err = rsa.GenerateKey(rand.Reader, 4096)

	if err != nil {
		log.Fatalf("failed to generate private key: %s", err)
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %s", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization:       []string{"Acme Co"},
		},
		NotBefore:              time.Now(),
		NotAfter:               time.Now().Add(24 * 365 * time.Hour),

		KeyUsage:               x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:            []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid:  true,
        IPAddresses:            []net.IP{net.ParseIP(host)},
        IsCA:                   true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}

    certOut := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
    log.Print("certOut Encoded!\n")

    keyOut := pem.EncodeToMemory(pemBlockForKey(priv))
    log.Print("keyOut Encoded!\n")

    return string(certOut), string(keyOut)
}

func WriteToFileIfNotExists(cert, key string) error {
    if FilesExist() {
        return nil
    }

    certErr := ioutil.WriteFile("." + pathSep + "certs" + pathSep + "cert.pem", []byte(cert), 0600)

    if certErr != nil {
        return certErr
    }

    keyErr := ioutil.WriteFile("." + pathSep + "certs" + pathSep + "key.pem", []byte(key), 0600)

    if(keyErr != nil) {
        return keyErr
    }

    return nil
}

func FilesExist() bool {
    if _, err := os.Stat("." + pathSep + "certs"); os.IsNotExist(err) {
        os.Mkdir("." + pathSep + "certs", 0600)
        return false
    }

    if _, err := os.Stat("." + pathSep + "certs" + pathSep + "cert.pem"); os.IsNotExist(err) {
        return false
    }

    if _, err := os.Stat("." + pathSep + "certs" + pathSep + "key.pem"); os.IsNotExist(err) {
        return false
    }

    return true
}
