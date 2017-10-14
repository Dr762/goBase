package security

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/blowfish"
	"log"
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"os"
	"encoding/gob"
	"encoding/pem"
	"crypto/x509"
	"time"
	"math/big"
	"crypto/x509/pkix"
)

func Md5Hash(str string) {

	hash := md5.New()
	bts := []byte(str + "\n")
	hash.Write(bts)
	hashVal := hash.Sum(nil)
	hashSize := hash.Size()

	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashVal[n])<<24 +
			uint32(hashVal[n]+1)<<16 +
			uint32(hashVal[n+2])<<8 +
			uint32(hashVal[n+3])

		fmt.Printf("%x", val)
	}

	fmt.Println()
}

//need to set right buffer sizes
//looks like it can't decrypt strings with len>8
func BlowFish(str string) {
	key := []byte("super secret and long key")

	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	src := []byte(str + "\n\n\n\n\n\n\n")
	var enc [256]byte

	cipher.Encrypt(enc[0:], src)
	fmt.Println("Encoded", enc)

	var decrypt [8] byte
	cipher.Decrypt(decrypt[0:], enc[0:])

	result := bytes.NewBuffer(nil)
	result.Write(decrypt[0:8])

	fmt.Println(string(result.Bytes()))
}

func GenRsaKey() {
	reader := rand.Reader
	bitSize := 512
	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())

	publicKey := key.PublicKey
	fmt.Println("Public key modulus", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)

	saveGobKey("private.key", key)
	saveGobKey("public.key", publicKey)
	savePemKey("private.pem", key)

}

func LoadRsaKey() {
	var key rsa.PrivateKey
	loadKey("private.key", &key)
	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())

	var publicKey rsa.PublicKey
	loadKey("public.key", &publicKey)
	fmt.Println("Public key modulus", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)

}

func GenX509() {
	random := rand.Reader

	var key rsa.PrivateKey
	loadKey("private.key", &key)

	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000)

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   "borscht.com",
			Organization: []string{"Borscht Systems AG"},
		},
		NotBefore:             now,
		NotAfter:              then,
		SubjectKeyId:          []byte{1, 2, 3, 4},
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames:              []string{"borscht.com", "localhost"},
	}

	derBytes, err := x509.CreateCertificate(random, &template, &template, &key.PublicKey, &key)
	if err != nil {
		log.Fatal(err)
	}

	certCerFile, err := os.Create("certificate.cer")
	if err != nil {
		log.Fatal(err)
	}

	certCerFile.Write(derBytes)
	certCerFile.Close()

	certPemFile, err := os.Create("certificate.pem")
	if err != nil {
		log.Fatal(err)
	}

	pem.Encode(certPemFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certPemFile.Close()

	keyPemFile, err := os.Create("private.pem")
	if err != nil {
		log.Fatal(err)
	}

	pem.Encode(keyPemFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(&key)})
	keyPemFile.Close()

}

func LoadX509() {
	certCerFile, err := os.Open("certificate.cer")
	if err != nil {
		log.Fatal(err)
	}

	derBytes := make([]byte, 1000)

	count, err := certCerFile.Read(derBytes)
	if err != nil {
		log.Fatal(err)
	}

	certCerFile.Close()


	// trim the bytes to actual length in call
	cert, err := x509.ParseCertificate(derBytes[0:count])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name %s\n", cert.Subject.CommonName)
	fmt.Printf("Not before %s\n", cert.NotBefore.String())
	fmt.Printf("Not after %s\n", cert.NotAfter.String())
}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	if err != nil {
		log.Fatal(err)
	}

	outFile.Close()
}

func savePemKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var privateKey = &pem.Block{Type: "RSA Private Key", Bytes: x509.MarshalPKCS1PrivateKey(key)}

	pem.Encode(outFile, privateKey)
	outFile.Close()
}

func loadKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	if err != nil {
		log.Fatal(err)
	}

	inFile.Close()
}
