# DJBase

A minimal base docker image with DataJoint dependencies installed.

# Launch locally


`docker-compose -f dist/alpine/docker-compose.yml --env-file config/.env up --build`
OR
`docker-compose -f dist/debian/docker-compose.yml --env-file config/.env up --build`


# Notes

https://hub.docker.com/r/datajoint/djbase