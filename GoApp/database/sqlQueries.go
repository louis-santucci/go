package database

const migrateQuery string = `
	create table if not exists redirections(
		id integer primary key autoincrement ,
		shortcut text not null unique,
		redirect_url text not null,
		views integer not null,
		created_at text not null
   );
	`

const createQuery string = `insert into redirections(shortcut, redirect_url, views, created_at) values (?,?,?,?)`

const selectAllQuery string = `select * from redirections`

const selectByIdQuery string = `select * from redirections where id = ?`

const selectByRedirectionQuery string = `select * from redirections where redirect_url = ?`

const updateQuery string = `update redirections set shortcut = ?, redirect_url = ?, views = ?, created_at = ? where id = ?`

const deleteQuery string = `delete from redirections where id = ?`
