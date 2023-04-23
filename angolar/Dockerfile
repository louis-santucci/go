FROM node:18.15-alpine as build

# Set the working directory
WORKDIR /angolar

# Add the source code to app
COPY ./angolar .

# Overwriting Docker config file
COPY ./docker/angolar/properties.service.template.ts ./src/app/services/
COPY ./docker/.env .

# Install envsubst
RUN apk update && apk add gettext

# Install all the dependencies
RUN npm install

RUN source .env && envsubst '$${IP}' < ./src/app/services/properties.service.template.ts > ./src/app/services/properties.service.ts

# Generate the build of the application
RUN npm run build


# Stage 2: Serve app with nginx server

# Use official nginx image as the base image
FROM nginx:alpine

# Expose port 80
EXPOSE 80
EXPOSE 443

# Copy the build output to replace the default nginx contents.
COPY --from=build /angolar/dist/angolar /usr/share/nginx/html

# Copy nginx.conf file for custom Nginx config
COPY ./docker/angolar/nginx.conf.template /etc/nginx/conf.d/
COPY ./docker/angolar/self-signed.conf /etc/nginx/snippets/
COPY ./docker/angolar/ssl-params.conf /etc/nginx/snippets/

# WARNING: environment dependency
# Copying SSL certificates to enable HTTPS: need to be changed for your usage
COPY ./certs/nginx.crt /etc/ssl/certs/
COPY ./certs/certificate /etc/ssl/private/nginx.key
COPY ./certs/dhparam.pem /etc/nginx/

COPY ./docker/.env .
RUN source .env && envsubst '$${IP}' < /etc/nginx/conf.d/nginx.conf.template > /etc/nginx/conf.d/nginx.conf
