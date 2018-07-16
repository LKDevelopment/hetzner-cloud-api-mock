FROM mhart/alpine-node:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app

COPY . .
RUN apk add --no-cache bash git openssh python make gcc g++
RUN npm install --production --quiet --silent
CMD ["./node_modules/.bin/drakov","-f docs.md","-p 8080","--disableCORS","--stealthmode","--public"]

EXPOSE 8080