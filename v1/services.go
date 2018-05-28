package main

import (
  "github.com/dgrijalva/jwt-go"
  "io/ioutil"
  "os"
  "crypto/rsa"
  "crypto/rand"
  "fmt"
  "encoding/pem"
  "crypto/x509"
)

type Authentication struct {
  *jwt.StandardClaims
  UserProfile
}

const (
  privKeyPath = "app.rsa"
  pubKeyPath  = "app.rsa.pub"
)

var (
  verifyKey *rsa.PublicKey
  signKey   *rsa.PrivateKey
)

// read the key files before starting http handlers
func LoadKeys() {
  if _, err := os.Stat(privKeyPath); os.IsNotExist(err) {
    GenerateKeys()
  }

  signBytes, err := ioutil.ReadFile(privKeyPath)
  checkError(err)

  signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
  checkError(err)

  verifyBytes, err := ioutil.ReadFile(pubKeyPath)
  checkError(err)

  verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
  checkError(err)
}

func GenerateKeys(){
  reader := rand.Reader
  bitSize := 2048

  key, err := rsa.GenerateKey(reader, bitSize)
  checkError(err)

  SavePrivatePEMKey(key)
  SavePublicPEMKey(key.PublicKey)
}

func SavePrivatePEMKey(key *rsa.PrivateKey) {
  outFile, err := os.Create(privKeyPath)
  checkError(err)
  defer outFile.Close()

  var privateKey = &pem.Block{
    Type:  "RSA PRIVATE KEY",
    Bytes: x509.MarshalPKCS1PrivateKey(key),
  }

  err = pem.Encode(outFile, privateKey)
  checkError(err)
}

func SavePublicPEMKey(pubkey rsa.PublicKey){
  asn1Bytes, err := x509.MarshalPKIXPublicKey(&pubkey)
  checkError(err)

  var pemkey = &pem.Block{
    Type:  "PUBLIC KEY",
    Bytes: asn1Bytes,
  }

  pemfile, err := os.Create(pubKeyPath)
  checkError(err)
  defer pemfile.Close()

  err = pem.Encode(pemfile, pemkey)
  checkError(err)
}

func checkError(err error) {
  if err != nil {
    fmt.Println("Fatal error ", err.Error())
    os.Exit(1)
  }
}