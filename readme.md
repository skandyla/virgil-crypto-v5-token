# virgil-crypto-v5-token
generating tokens for VirgilSecurity api

### requirements
- golang 1.10 installed
- virgil crypto and sdk see: https://developer.virgilsecurity.com/docs/go/how-to/setup/v5/install-sdk

Centos7 dependenses (for using prebuilded crypto_lib):
```
yum install gcc libstdc++-static
```

### install

```
go get -u github.com/skandyla/virgil-crypto-v5-token
```


### usage
```
virgil-crypto-v5-token -h
virgil-crypto-v5-token -config token.conf
```

example of token.conf:
```
privateKeyStr example
appId example
appPubKeyId example
identity example
searchCard example
```