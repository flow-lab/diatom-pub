#!/bin/bash

set -euo pipefail

# Generate a self-signed certificate for db.
# This is only for testing purposes.

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

PASSWORD=$(echo $RANDOM | md5sum | head -c 20)
openssl req -new -text -passout pass:${PASSWORD} -subj /CN=db -out ${DIR}/server.req -keyout ${DIR}/privkey.pem
openssl rsa -in ${DIR}/privkey.pem -passin pass:${PASSWORD} -out ${DIR}/server.key
openssl req -x509 -in ${DIR}/server.req -text -key ${DIR}/server.key -out ${DIR}/server.crt
chmod 600 ${DIR}/server.key

rm ${DIR}/privkey.pem
rm ${DIR}/server.req

# for debian
OS=`uname`

if [ "${OS}" == "Linux" ]; then
    if [ -f /etc/debian_version ] ; then
        echo "setting keys for debian";
        chown 999:999 ${DIR}/server.key
    fi
fi

chmod 600 ${DIR}/server.key