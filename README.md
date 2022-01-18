### Go AES Encryptor 

This is a POC for go encryption with AES in GCM mode.

use via command line `go run gocryptor.go`: 
 - `-f` The path to the file you want to encrypt/decrypt
 - `-d` Decrypt the file
 - `-e` Encrypt the file
 - `-p` The key file used to encrypt/decrypt the file (16/24/32 bytes)
 - `-i` The identification vector(iv/nonce) (should change for every encryption) - Base64 Expected
 - `-t` The tag - Base64 Expected
 - `-a` The additional authentication data string

### Encrypt
```bash
go run gocryptor.go -e -f ./myfile.txt -p ./base64key.key -i jQm0NcAOSkiAVFc9qGpMfA== -a AA_BB_CC_DD_EE
```
This will give you the tag save it, you will need it below. 

### Decrypt
```bash
go run gocryptor.go -d -f ./myfile.encrypted -p ./base64key.key -i jQm0NcAOSkiAVFc9qGpMfA== -a AA_BB_CC_DD_EE -t Z7s5GCSrSBcIoj+FvFNprw==
```

## Resources
This was my first encryption project so I'll leave below some ressources I used to create this project for legacy.

- https://pkg.go.dev/crypto/cipher
- https://cs.opensource.google/go/go/+/master:src/crypto/aes/aes_gcm.go;l=70;drc=master;bpv=1;bpt=1
- https://github.com/dvsekhvalnov/jose2go/blob/master/aes_gcm.go
- https://crypto.stackexchange.com/questions/41601/aes-gcm-recommended-iv-size-why-12-bytes
- https://stackoverflow.com/questions/47382035/unable-to-decrypt-after-aes-gcm-base64-in-go
- https://github.com/dvsekhvalnov/jose2go/blob/master/aes_gcm.go
- https://stackoverflow.com/questions/68040875/aes-256-gcm-encryption-from-ruby-decryption-with-golang
- https://cloud.google.com/kms/docs/additional-authenticated-data
- https://stackoverflow.com/a/68353192
- https://github.com/golang/go/issues/42470
- https://github.com/parkerdouglass/fcrypt/blob/master/main.go
- https://gist.github.com/cannium/c167a19030f2a3c6adbb5a5174bea3ff
- https://stackoverflow.com/questions/39347206/how-to-encrypt-files-with-aes256-gcm-in-golang
- https://github.com/SimonWaldherr/golang-examples/blob/master/advanced/aesgcm.go
- https://eli.thegreenplace.net/2019/aes-encryption-of-files-in-go/
- https://github.com/rfjakob/gocryptfs/issues/222#issuecomment-379539825
- https://github.com/Xeoncross/go-aesctr-with-hmac/blob/master/crypt.go
- https://github.com/gtank/cryptopasta/blob/master/encrypt.go
