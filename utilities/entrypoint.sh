#!/bin/sh

#Set default permission of new files
umask u+rwx,g+rwx,o-rwx

#Fix UID/GID
/startup -user=dja -new_uid=$(id -u) -new_gid=$(id -g)

#Enable conda paths
. /etc/profile.d/shell_intercept.sh

#Install Conda dependencies
if [ -f "$CONDA_REQUIREMENTS" ]; then
    conda install -yc conda-forge --file $CONDA_REQUIREMENTS
fi

#Install Python dependencies
if [ -f "$PIP_REQUIREMENTS" ]; then
    pip install -r $PIP_REQUIREMENTS --upgrade --no-cache-dir
fi

#Run command
"$@"