FROM mhart/alpine-node:latest

# Set the working directory to /app
WORKDIR /app

RUN apk add --no-cache bash git openssh python make gcc g++ curl

# Copy the current directory contents into the container at /app
COPY package*.json ./
RUN npm install --production --quiet --silent

COPY . .

CMD ["./api-mock-server.sh"]

EXPOSE 8080