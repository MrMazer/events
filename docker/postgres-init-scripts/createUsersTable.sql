CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    status INT NOT NULL
);



CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    organizer_id INT REFERENCES users(id),
    name_event VARCHAR(255) NOT NULL,
    shape VARCHAR(255) NOT NULL,
    place VARCHAR(255) NOT NULL,
    begin_time TIMESTAMP NOT NULL,
    duration VARCHAR(255) NOT NULL
);

CREATE TABLE participation (
                               id SERIAL PRIMARY KEY,
                               user_id INT NOT NULL,
                               event_id INT NOT NULL,
                               FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                               FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);

CREATE TABLE event_feedback (
    id SERIAL PRIMARY KEY,
    event_id INT NOT NULL,
    user_id INT NOT NULL,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE event_media (
    id SERIAL PRIMARY KEY,
    event_id INT NOT NULL,
    media_type VARCHAR(50), -- например, фото, видео
    media_url TEXT NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);

CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    capacity INT,
    contact_person VARCHAR(255),
    phone_number VARCHAR(20)
);

CREATE TABLE logs (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    action_type VARCHAR(255) NOT NULL, -- например, создание события, обновление информации
    description TEXT, -- более детальное описание действия
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ip_address VARCHAR(45), -- IP-адрес пользователя
    user_agent TEXT, -- информация о браузере и устройстве
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);





