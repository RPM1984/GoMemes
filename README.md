# GoMemes
[![Build Status](https://dev.azure.com/rpmir1/GoMemes/_apis/build/status/RPM1984.GoMemes)](https://dev.azure.com/rpmir1/GoMemes/_build/latest?definitionId=1)

# What is this?
A simple Go web app that generates random memes, using the [imflip](https://api.imgflip.com/) API.

# Building
> docker build . -t rpm1984/gomemes

or pull from [DockerHub](https://hub.docker.com/r/rpm1984/gomemes/):
> docker pull rpm1984/gomemes

# Running
> docker run -p 3030:3001 -it rpm1984/gomemes