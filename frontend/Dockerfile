FROM node:16.3.0-alpine

RUN apk add --no-cache \
    build-base \
    g++ \
    libpng \
    libpng-dev \
    jpeg-dev \
    pango-dev \
    cairo-dev \
    giflib-dev \
    python3 
RUN npm install -g http-server
WORKDIR ./frontend
COPY package*.json ./

RUN npm cache clean --force 
RUN npm install -g @vue/cli
RUN npm install
#RUN vue ui
COPY . .

CMD [ "http-server", "dist" ]