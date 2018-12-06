FROM node

RUN mkdir /app

COPY src/server /app
COPY src/public /app/public

COPY package.json /app/package.json

WORKDIR /app

RUN npm install

CMD [ "node", "/app/index.js" ]