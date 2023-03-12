FROM node:18.15-alpine as build

# Set the working directory
WORKDIR /angolar

# Add the source code to app
COPY ./angolar .

# Install all the dependencies
RUN npm install --loglevel verbose

# Overwriting Docker config file
COPY ./docker/config/properties.service.ts ./src/app/services

# Generate the build of the application
RUN npm run build


# Stage 2: Serve app with nginx server

# Use official nginx image as the base image
FROM nginx:latest

# Copy the build output to replace the default nginx contents.
COPY --from=build /angolar/dist/angolar /usr/share/nginx/html

# Copy nginx.conf file for custom Nginx config
COPY ./docker/nginx/nginx.conf /etc/nginx/conf.d
COPY ./docker/nginx/self-signed.conf /etc/nginx/snippets

# WARNING: environment dependency
# Copying SSL certificates to enable HTTPS: need to be changed for your usage
COPY $HOME/Documents/certs/certs/nginx.crt /etc/ssl/certs
COPY $HOME/Documents/certs/private/nginx.key /etc/ssl/private
COPY $HOME/Documents/certs/pem/dhparam.pem /etc/nginx

# Expose port 80
EXPOSE 80
EXPOSE 443
