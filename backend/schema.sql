-- schema.sql

-- Create user table
CREATE TABLE user (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
);

-- Create director table
CREATE TABLE director (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Create star table
CREATE TABLE star (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Create movie table
CREATE TABLE movie (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    genre VARCHAR(50),
    duration INT,
    director_id INT REFERENCES director(id),
    featured_image VARCHAR(255),
    thumbnail VARCHAR(255)
);

-- Create schedule table
CREATE TABLE schedule (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movie(id),
    showtime TIMESTAMP NOT NULL
);

-- Create ticket table
CREATE TABLE ticket (
    id SERIAL PRIMARY KEY,
    schedule_id INT REFERENCES schedule(id),
    user_id INT REFERENCES user(id),
    purchase_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create rating table
CREATE TABLE rating (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movie(id),
    user_id INT REFERENCES user(id),
    rating INT CHECK (rating >= 1 AND rating <= 5)
);

-- Create movie_star table
CREATE TABLE movie_star (
    movie_id INT REFERENCES movie(id),
    star_id INT REFERENCES star(id),
    PRIMARY KEY (movie_id, star_id)
);

-- Create saved_movie table
CREATE TABLE saved_movie (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES user(id),
    movie_id INT REFERENCES movie(id),
    saved_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
