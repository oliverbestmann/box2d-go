#!/usr/bin/env bash

set -e -u -o pipefail

rm -rf box2d-c
git clone --rev v3.2.0 https://github.com/erincatto/box2d.git box2d-c

git -C box2d-c apply < box2d-3.2.0.patch
