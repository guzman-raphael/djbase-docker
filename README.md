# DJBase

A minimal base docker image with DataJoint dependencies installed.

# Features

- Creates a internal system user `dja`. By default, utilizes `dja:anaconda` when creating containers.
- Image avoids using root to preserve possible privilege escalation vulnerabilities on host. As a means to perform installs while under `dja:anaconda`, you may utilize any combination of the following and trigger it via `/entrypoint.sh [command to run once completed]` e.g. `/entrypoint.sh echo done`.
  - Debian dependencies: Create a `\n` delimited file containing the system dependencies at `/tmp/apt_requirements.txt`. This can be created manually within container/image or mounted in.
  - Alpine dependencies: Create a `\n` delimited file containing the system dependencies at `/tmp/apk_requirements.txt`. This can be created manually within container/image or mounted in.
  - Conda dependencies: Create a `\n` delimited file containing the conda dependencies at `/tmp/conda_requirements.txt`. This can be created manually within container/image or mounted in.
  - Pip dependencies: Create a `\n` delimited file containing the pip dependencies at `/tmp/pip_requirements.txt`. This can be created manually within container/image or mounted in.
- When mounting volumes (avoid mounting files with `root` permissions only!), accessing files within image via `dja:anaconda` can result in permission denied errors. To avoid this, simply add a `user` spec in `docker-compose.yml` with the appropriate UID e.g. `user: 1000:anaconda`. Running `/entrypoint.sh` will then trigger a reassociation of `dja`'s UID to allow permissions to access mounted files. Note that `entrypoint.sh` is automatically invoked when starting containers.
- Adds `conda-forge` channel
- Installs `datajoint` dependencies w/o actually installing `datajoint`
- Applies image compresssion

# Launch locally

```shell
docker-compose -f dist/alpine/docker-compose.yml --env-file config/.env up --build
```

OR

```shell
docker-compose -f dist/debian/docker-compose.yml --env-file config/.env up --build
```

# Notes

https://hub.docker.com/r/datajoint/djbase