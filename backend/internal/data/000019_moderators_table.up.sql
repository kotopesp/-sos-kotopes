CREATE TABLE IF NOT EXISTS
    moderators
(
    user_id INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
);