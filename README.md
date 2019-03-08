# printNTLM

__On Windows:__
* Prints Net-NTLM hashes for the current user by starting a local webDAV server and authenticating
    to it
* Starts a persistent webDAV server the will capture hashes of remote computers authenticating to
    it

__On Linux:__
* Starts a persistent webDAV server the will capture hashes of remote computers authenticating to
    it

Importable as a library from `pkg` or run as a standalone binary with:

```sh
GOOS=windows go build
```
