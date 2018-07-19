#!/usr/bin/env bash
echo "First get the latest version of the api!\n"
curl https://docs.hetzner.cloud/docs.md --output docs.md
echo "Now we start the server"
./node_modules/.bin/drakov -f docs.md -p 8080 --disableCORS --stealthmode --public