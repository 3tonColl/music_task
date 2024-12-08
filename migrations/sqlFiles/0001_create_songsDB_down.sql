CREATE TABLE IF NOT EXISTS songs
(
    id           SERIAL PRIMARY KEY,
    band_name      VARCHAR(255) NOT NULL,
    song_name        VARCHAR(255) NOT NULL,
    release_date VARCHAR(10),
    lyrics         TEXT,
    link         VARCHAR(255)
)