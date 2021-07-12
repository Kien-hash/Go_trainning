CREATE TABLE IF NOT EXISTS Users
(
  id INT(6) UNSIGNED AUTO_INCREMENT ,
  username VARCHAR(30) NOT NULL unique key,
  passwords VARCHAR(30) NOT NULL,
  firstname VARCHAR(30) NOT NULL,
  lastname VARCHAR(30) NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS Rooms (
id INT(6) UNSIGNED AUTO_INCREMENT ,
UsersID VARCHAR(30) NOT NULL unique key,
PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS Message (
id INT(6) UNSIGNED AUTO_INCREMENT ,
RoomID INT(6) UNSIGNED,
SenderID INT(6) UNSIGNED,
MsgContent VARCHAR(100),
PRIMARY KEY (id),
FOREIGN KEY (RoomID) REFERENCES Rooms(id),
foreign key (SenderID) REFERENCES Users(id)
);

INSERT INTO Users (username, passwords, firstname, lastname) VALUES ("admin", "3026972", "_", "admin _");
