FROM node:14.17.0

ENV PORT 3000

# Create app directory
WORKDIR /app

# Installing dependencies
COPY package*.json /app
RUN npm install

# Copying source files
COPY . /app

# Building app
EXPOSE 3000

# Running the app
CMD "npm" "run" "dev"