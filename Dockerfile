FROM node:latest
RUN apt-get update
RUN apt-get -y install golang sudo
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
WORKDIR /run
CMD ["yarn", "run", "test"]