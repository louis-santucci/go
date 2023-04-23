# GO-APP

Check https://go.dev/doc/tutorial/


## Goal

This is an application allowing users to create redirections, and running them by typing go/{redirection} to go through them.

It is an internal existing project for Société Générale Corporate & Investment Banking, and Société Générale users.

As a challenge, the goal is to realize the same objectives, in a different language, and from scratch

For this project, the selected language is the Go for the back-end part.

## Prerequisities

- Go 1.19
- A GNU Compiler Collection (GCC) (Some dependencies use C compiled codes)
- Swag (if you want to contribute and then add some endpoints on Swagger-UI)
- Node.js & npm

## How-to build the project

- You need them to go to goApp folder and then run `go install .`
- Concerning the Angolar application, you can go to angolar folder and then run `npm install` to install all dependencies of the front-end project.

## Deployment on server

To deploy the application with Docker:

- First you need to create an environment variable $IP with the IP of your server so that it is added to the nginx configuration file for the Angolar application, and the overall Nginx reverse proxy instance (for example: `export IP="192.168.1.117"`)
- You need to generate some CA certificates to run the application with HTTPS. There are many tutorials on the web, but you can use this one to generate nginx.crt file, nginx.key and dhparam.pem file ([Tutorial link](https://mpolinowski.github.io/docs/DevOps/NGINX/2020-08-27--nginx-docker-ssl-certs-self-signed/2020-08-27)). You need to create at the root of the project a folder named certs. You can edit the folder or even the file names directly in the Dockerfiles of Goapi & Angolar applications.
- Then you need to run in the project root
  `docker compose build` and then `docker compose up` to run the application.

## Swagger

The project contains a Swagger configuration built with [swag](https://github.com/swaggo/swag)
To get swag, you need to enter: `go install github.com/swaggo/swag/cmd/swag@latest`

## Middlewares

In this project, JWT Authentication is necessary to log in and then use the project.
Moreover, there is another custom middleware used to check if the request is coming from the frontend angolar application, to prevent usages from the exterior, and then limit certain features of the project only to the frontend usage

## TODOs

- First, the goal is to implement a REST API allowing to CRUD redirections.
    - [x] Create all project architecture
    - [x] Finalize all CRUD endpoints
    - [x] Secure API by using JWT
    - [ ] Enable HTTPS
    - [ ] Deploy on docker-compose
    - 
- Then, we need to implement a front-end application so that when typing go/{redirection}, we will be either redirected to the good redirection, else we will be invited to create a new redirection for the entered redirection.
    - [x] Choose front-end framework: ANGULAR
    - [x] Login/Logout component
    - [x] Create form and component to Add/Update redirections
    - [x] Add history
    - [ ] User handling
