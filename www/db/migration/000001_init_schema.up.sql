CREATE TABLE building(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE lot(
    id SERIAL NOT NULL PRIMARY KEY,
    floor INTEGER NOT NULL,
    total_square NUMERIC(6,2) NOT NULL,
    local_square NUMERIC(6,2) NOT NULL,
    kitchen_square NUMERIC(6,2) NOT NULL,
    price INTEGER NOT NULL,
    lot_type VARCHAR(255) NOT NULL,
    room_type VARCHAR(255) NOT NULL
);

CREATE TABLE lotbuilding(
    id SERIAL NOT NULL PRIMARY KEY,
    idlot INTEGER NOT NULL REFERENCES lot(id),
    idbuilding INTEGER NOT NULL REFERENCES building(id)
);

CREATE TABLE project(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(2047) NOT NULL,
    address VARCHAR(255) NOT NULL
);

CREATE TABLE projectbuilding(
    id SERIAL NOT NULL PRIMARY KEY,
    idproject INTEGER NOT NULL REFERENCES project(id),
    idbuilding INTEGER NOT NULL REFERENCES building(id)
);