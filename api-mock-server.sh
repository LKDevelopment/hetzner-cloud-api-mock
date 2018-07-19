#!/usr/bin/env bash
echo "First get the latest version of the api!\n"
curl https://docs.hetzner.cloud/latest_docs.apib --output latest_docs.apib
echo "Now we start the server"
./node_modules/.bin/drakov -f latest_docs.apib -p 8080 --disableCORS --stealthmode --public