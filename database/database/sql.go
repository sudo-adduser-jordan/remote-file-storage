package database

// pgx v5

const CREATE_USER_TABLE = `
	CREATE TABLE IF NOT EXISTS Users (
		UserID SERIAL,
		Username TEXT NOT NULL UNIQUE,
		Password TEXT NOT NULL UNIQUE,
		PRIMARY KEY(UserID)
	);
`

const DROP_TABLE = `
	DROP TABLE Users;
`

const INSERT_USER = `
	INSERT INTO Users (Username, Password)
	VALUES ($1, $2); 
`

const SELECT_USER = `
	SELECT * 
	FROM users
	WHERE username = $1
	ORDER BY username ASC
	LIMIT 1
`

const UPDATE_USER = `
	UPDATE Users
	SET Username = $1, Password = $2
	WHERE Username = $3; 
`

const DELETE_USER = `
	DELETE FROM Users 
	WHERE Username = $1;
`
