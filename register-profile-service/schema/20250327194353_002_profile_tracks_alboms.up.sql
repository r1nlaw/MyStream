CREATE TABLE profiles (
    user_id INT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    avatar_url TEXT,
    user_bio TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);


CREATE TABLE all_albums (
    id SERIAL PRIMARY KEY,
    album_name VARCHAR(50) NOT NULL,
    artist VARCHAR(20) NOT NULL,
    release_date DATE,
    cover_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE all_tracks (
    id SERIAL PRIMARY KEY,
    track_name VARCHAR(50) NOT NULL,
    artist VARCHAR(20) NOT NULL,
    album_id INT REFERENCES all_albums(id) ON DELETE SET NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_tracks (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    track_id INT REFERENCES all_tracks(id) ON DELETE CASCADE,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, track_id)
);

CREATE TABLE user_albums (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    album_id INT REFERENCES all_albums(id) ON DELETE CASCADE,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, album_id)
);

CREATE TABLE friends_requests (
    sender_id INT REFERENCES users(id) ON DELETE CASCADE,
    receiver_id INT REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'accepted', 'rejected')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (sender_id, receiver_id)
);