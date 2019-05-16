#!/usr/bin/env bash
git clone  -b stable/3.x --recursive git://github.com/apiaryio/drafter.git
cd drafter
./configure
make install