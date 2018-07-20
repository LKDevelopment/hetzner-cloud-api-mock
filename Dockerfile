FROM mhart/alpine-node:10
RUN mkdir /app
# Set the working directory to /app
WORKDIR /app

RUN apk add --no-cache bash git openssh python make gcc g++ curl

# Copy the current directory contents into the container at /app
COPY package.json /app
COPY api-mock-server.sh /app
RUN yarn

CMD ["./api-mock-server.sh"]

EXPOSE 8080
