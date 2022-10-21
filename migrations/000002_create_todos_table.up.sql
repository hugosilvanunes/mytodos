CREATE TABLE IF NOT EXISTS todos(
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  title VARCHAR (300) NOT NULL,
  content VARCHAR (300),
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);
