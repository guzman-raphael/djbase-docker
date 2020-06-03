#!/bin/sh
. /opt/conda/etc/profile.d/conda.sh
conda activate
if ! [ $(id -u) = 0 ]; then
    export HOME=/home/$(whoami)
fi
export PATH=$(readlink -f "$HOME")/.local/bin:$PATH