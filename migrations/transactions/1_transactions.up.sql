CREATE TABLE IF NOT EXISTS transactions (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  userId INTEGER NOT NULL,
  t_date INTEGER NOT NULL,
  title VARCHAR(255) NOT NULL,
  t_description VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL,
  amount INTEGER NOT NULL,
);