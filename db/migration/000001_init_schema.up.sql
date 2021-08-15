CREATE TABLE project(
	id SERIAL NOT NULL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description VARCHAR(2047),
	address VARCHAR(255)
	);
	
CREATE TABLE building(
	id INTEGER NOT NULL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	id_project integer NOT NULL REFERENCES project(id)
	);
	
CREATE TABLE section(
	id SERIAL NOT NULL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	id_building integer NOT NULL REFERENCES building(id)
	);
	
CREATE TABLE lot(
	id integer NOT NULL PRIMARY KEY,
	floor integer NOT NULL,
	total_square NUMERIC(6,2) NOT NULL,
	living_square NUMERIC(6,2) NOT NULL,
	kitchen_square NUMERIC(6,2) NOT NULL,
	price integer NOT NULL,
	lot_type VARCHAR(255) NOT NULL,
	room_type VARCHAR(255) NOT NULL,
	id_section integer NOT NULL references section(id)
	);