FROM mhart/alpine-node:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app

COPY . .
RUN apk add --no-cache bash git openssh python make gcc g++ curl
RUN npm install --production --quiet --silent
CMD ["./api-mock-server.sh"]

EXPOSE 8080