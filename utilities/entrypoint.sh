#!/bin/sh

#Fix UID/GID
/startup $(id -u) $(id -g)

#Install Python dependencies
if [ -f "$PIP_REQUIREMENTS" ]; then
    pip install --user -r $PIP_REQUIREMENTS
fi

#Install Conda dependencies
if [ -f "$CONDA_REQUIREMENTS" ]; then
    conda install -yc conda-forge --file $CONDA_REQUIREMENTS
fi

#Command
. /etc/profile.d/shell_intercept.sh
"$@"