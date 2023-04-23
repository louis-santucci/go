# Certificates folder

This folder is used for CA certificates when serving application with HTTPS.

You can add in this folder the following files:

- nginx.crt
- nginx.key (you will need to `sudo cp nginx.key certificate` because Docker won't copy the original file on containers)
- dhparam.pem