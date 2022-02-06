# pull official base image
FROM node:13.12.0-alpine

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY web/package.json ./
COPY web/package-lock.json ./
COPY web/tsconfig.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent
RUN npm install -g typescript

ADD web ./

EXPOSE 3000

CMD ["npm", "start"]