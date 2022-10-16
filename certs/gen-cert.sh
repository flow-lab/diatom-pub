#!/bin/bash

set -euo pipefail

PASS=apud123

openssl req -new -text -passout pass:${PASS} -subj /CN=localhost -out server.req -keyout privkey.pem
openssl rsa -in privkey.pem -passin pass:${PASS} -out server.key
openssl req -x509 -in server.req -text -key server.key -out server.crt
chmod 600 server.key

rm privkey.pem
rm server.req

# for debian
OS=`uname`

if [ "${OS}" == "Linux" ]; then
    if [ -f /etc/debian_version ] ; then
        echo "setting keys for debian";
        chown 999:999 server.key
    fi
fi

chmod 600 server.key