FROM node:latest

ENV PORT 3000

# Create app directory
RUN mkdir -p /app/web
WORKDIR /app/web

# Installing dependencies
COPY package*.json /app/web
RUN npm install

# Copying source files
COPY . /app/web

# Building app
EXPOSE 3000

# Running the app
CMD "npm" "run" "dev"