FROM node:alpine as build

# Set the working directory
WORKDIR /angolar

# Add the source code to app
COPY ./angolar /angolar

# Install all the dependencies
RUN npm install

# Generate the build of the application
RUN npm run build


# Stage 2: Serve app with nginx server

# Use official nginx image as the base image
FROM nginx:latest

# Copy the build output to replace the default nginx contents.
COPY --from=build /angolar/dist/angolar /usr/share/nginx/html

# Copy nginx.conf file for custom Nginx config
COPY ./nginx/nginx.conf /etc/nginx/conf.d
COPY ./nginx/self-signed.conf /etc/nginx/snippets

# WARNING: environment dependency
# Copying SSL certificates to enable HTTPS: need to be changed for your usage
COPY ~/Documents/certs/nginx.crt /etc/ssl/certs
COPY ~/Documents/certs/nginx.key /etc/ssl/private
COPY ~/Documents/certs/dhparam.pem /etc/nginx

# Expose port 80
EXPOSE 80
EXPOSE 443
