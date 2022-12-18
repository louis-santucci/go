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

## TODOs

- First, the goal is to implement a REST API allowing to CRUD redirections.
    - [ ] Create all project architecture
    - [ ] Finalize all CRUD endpoints
    - [ ] Secure API by using JWT
    - [ ] Enable HTTPS
- Then, we need to implement a front-end application so that when typing go/{redirection}, we will be either redirected to the good redirection, else we will be invited to create a new redirection for the entered redirection.
    - [ ] Choose front-end framework
    - [ ] Login/Logout component
    - [ ] Create form and component to Add/Update redirections