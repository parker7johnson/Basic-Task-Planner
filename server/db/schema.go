package schema

var Schema = `
create table task (
	id integer primary key,
	name string,
	deadline datetime,
	status string
);
`
