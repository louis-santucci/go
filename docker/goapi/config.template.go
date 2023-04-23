package config

const Db_filename = "db/go-db.db"
const Jwt_secret = "" //FIXME TO be changed for security purposes
// Angolar_secret is a token to prevent API to be used by something other
// than the frontend application. It only allows the Angolar application
// to request the API for certain endpoints, such like the incrementation of the views,
// which is a special feature for the Angolar application
const Angolar_secret = "" //FIXME TO be changed for security purposes
const TLS_enabled = true
const HOST = "${IP}"
