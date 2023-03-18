FROM node:16.19.0

WORKDIR /front_app

COPY package.json ./
COPY package-lock.json ./

RUN npm install

EXPOSE 3000 

COPY . .

CMD ["npm", "start"]