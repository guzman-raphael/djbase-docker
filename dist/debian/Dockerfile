ARG CONDA_VER
ARG PY_VER

#Temp Image to create exec to allow UID/GID to be updated on boot
FROM golang:alpine3.11 as go_tmp
COPY ./utilities/startup.go /startup.go
RUN cd / && go build startup.go

from raphaelguzman/miniconda3:${CONDA_VER}-py${PY_VER}-debian as py_build
# # datajoint admin (dja) mapped to docker host user
user root
RUN export uid=3000 gid=3000 && \
    mkdir -p /home/dja && \
    mkdir /src && \
    echo "dja:x:${uid}:anaconda:Developer,,,:/home/dja:/bin/bash" >> /etc/passwd && \
    # echo "dja:x:${uid}:" >> /etc/group && \
    chown ${uid}:${gid} -R /home/dja && \
    chown ${uid}:${gid} -R /src && \
    chown -R ${uid}:anaconda /opt/conda && \
    chmod -R g+w /opt/conda

ENV PYTHON_USER dja
LABEL maintainerName="Raphael Guzman" \
      maintainerEmail="raphael@vathes.com" \
      maintainerCompany="DataJoint"

USER dja:anaconda
COPY ./utilities/shell_intercept.sh /etc/profile.d/

SHELL ["/bin/bash", "-lc"]

ARG PY_VER
ARG CONDA_VER
RUN \
    conda config --add channels conda-forge && \
    conda install -yc conda-forge conda=$CONDA_VER datajoint --only-deps && \
    conda config --set auto_update_conda False && \
    conda clean -ya && \
    export MATPLOT_DIR=$(ls /opt/conda/pkgs | grep matplotlib-base) && \
    if [ ! -z "$MATPLOT_DIR" ]; then \
      rm /opt/conda/pkgs/$MATPLOT_DIR/lib/libtk.so; \
      rm /opt/conda/pkgs/$MATPLOT_DIR/lib/libtcl.so; \
      ln -s /opt/conda/lib/libtk.so  /opt/conda/pkgs/$MATPLOT_DIR/lib/libtk.so; \
      ln -s /opt/conda/lib/libtcl.so /opt/conda/pkgs/$MATPLOT_DIR/lib/libtcl.so; \
    fi && \ 
    find /opt/conda -user 3000 -exec chmod g+w "{}" \;

user dja:anaconda
ENV HOME /home/dja
ENV LANG C.UTF-8
ENV APK_REQUIREMENTS /apk_requirements.txt
ENV PIP_REQUIREMENTS /pip_requirements.txt
ENV CONDA_REQUIREMENTS /conda_requirements.txt
COPY ./utilities/entrypoint.sh /entrypoint.sh
COPY --from=go_tmp /startup /startup
user root
RUN \
  chmod +x /entrypoint.sh && \
  chmod 4755 /startup
ENTRYPOINT ["/entrypoint.sh"]
WORKDIR /src
VOLUME /src
VOLUME /tmp/.X11-unix
EXPOSE 8888
CMD ["/bin/bash", "-l"]


# #Squashed Final Image
FROM scratch
COPY --from=py_build / /
ENV PYTHON_USER dja
RUN chmod 4755 /startup && /startup 3000 anaconda
LABEL maintainerName="Raphael Guzman" \
      maintainerEmail="raphael@vathes.com" \
      maintainerCompany="DataJoint"
USER dja:anaconda
ENV HOME /home/dja
ENV LANG C.UTF-8
ENV APK_REQUIREMENTS /apk_requirements.txt
ENV PIP_REQUIREMENTS /pip_requirements.txt
ENV CONDA_REQUIREMENTS /conda_requirements.txt
ENV PATH "/home/dja/.local/bin:/opt/conda/bin:$PATH"
ENTRYPOINT ["/entrypoint.sh"]
WORKDIR /src
VOLUME /src
VOLUME /tmp/.X11-unix
EXPOSE 8888
CMD ["/bin/bash", "-l"]