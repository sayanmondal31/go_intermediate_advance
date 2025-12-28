Dependencies installed
-
- golang.org/x/net/http2

Note
-
1. Enable TLS to secure communication between client and server
2. EOF in network is protocol error 

SSL
-
<code>
openssl genpkey -algorithm RSA -out server.key -pkeyopt rsa_keygen_bits:2048
</code>
setup server.crt file 


<code>
openssl req -new -x509 -key server.key -out server.crt -days 365
</code>


