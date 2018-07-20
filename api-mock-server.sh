#!/usr/bin/env bash
echo "First get the latest version of the api!\n"
curl https://docs.hetzner.cloud/docs.apib --output docs.apib
echo "Now we start the server"
./node_modules/.bin/drakov -f docs.apib -p 8080 --disableCORS --stealthmode --public