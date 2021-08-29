CREATE TABLE groups (
    group_id SERIAL PRIMARY KEY,
    gname VARCHAR(64)
);
INSERT INTO groups(gname) 
    VALUES ('Employee'),('Customers'),('Managers');
CREATE TABLE phonebook (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    phone VARCHAR(64),
    group_id int,
    FOREIGN KEY (group_id) REFERENCES groups(group_id)
);