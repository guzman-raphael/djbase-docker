Documentation for the DataJoint's DJBase Image
##############################################

| A minimal base docker image with `DataJoint Python <https://github.com/datajoint/datajoint-python>`_ dependencies installed.
| For more details, have a look at `prebuilt images <https://hub.docker.com/r/datajoint/djbase>`_, `source <https://github.com/datajoint/djbase-docker>`_, and `documentation <https://datajoint.github.io/djbase-docker>`_.

.. toctree::
   :maxdepth: 2
   :caption: Contents:

Launch Locally
**************

Debian
======
.. code-block:: shell

   docker-compose -f dist/debian/docker-compose.yaml --env-file config/.env up --build

Alpine
======
.. code-block:: shell

   docker-compose -f dist/alpine/docker-compose.yaml --env-file config/.env up --build

Features
********

- Creates a internal system user ``dja``. By default, utilizes ``dja:anaconda`` when creating containers.
- Image avoids using ``root`` to preserve possible privilege escalation vulnerabilities on docker host. As a means to perform installs while under ``dja:anaconda``, you may utilize any combination of the following and trigger it via ``/entrypoint.sh [command to run once completed]`` e.g. ``/entrypoint.sh echo done``.

  - Debian dependencies: Create a ``\n`` delimited file containing the system dependencies at ``/tmp/apt_requirements.txt``. This can be created manually within container/image or mounted in.
  - Alpine dependencies: Create a ``\n`` delimited file containing the system dependencies at ``/tmp/apk_requirements.txt``. This can be created manually within container/image or mounted in.
  - Conda dependencies: Create a ``\n`` delimited file containing the conda dependencies at ``/tmp/conda_requirements.txt``. This can be created manually within container/image or mounted in.
  - Pip dependencies: Create a ``\n`` delimited file containing the pip dependencies at ``/tmp/pip_requirements.txt``. This can be created manually within container/image or mounted in.

- When mounting volumes (avoid mounting files with ``root`` permissions only!), accessing files within image via ``dja:anaconda`` can result in permission denied errors. To avoid this, simply add a ``user`` spec in ``docker-compose.yaml`` with the appropriate ``HOST_UID`` e.g. ``user: 1000:anaconda``. Running ``/entrypoint.sh`` will then trigger a reassociation of ``dja``'s UID to allow permissions to access mounted files. Note that ``entrypoint.sh`` is automatically invoked when starting containers. If you are utilizing the included reference ``docker-compose.yaml`` files, you may simply set the ``HOST_UID`` environment value when building or starting the container.
- Installs ``datajoint`` dependencies w/o actually installing ``datajoint``.
- Applies image compresssion.

Base Image
**********

Build is a child of `datajoint/miniconda3 <https://github.com/datajoint/miniconda3-docker>`_.