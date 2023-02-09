#base image
FROM node:16

# create app directory
WORKDIR /abhishek

COPY package*.json /abhishek/
RUN npm install
COPY . /abhishek/


EXPOSE 8080
CMD [ "node", "index.js" ]
