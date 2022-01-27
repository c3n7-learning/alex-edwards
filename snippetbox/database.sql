CREATE TABLE snippets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMP NOT NULL,
    expires TIMESTAMP NOT NULL
);

INSERT INTO
    snippets (title, content, created, expires)
VALUES
    (
        'An old silent pond',
        'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP + INTERVAL '1 year'
    );

INSERT INTO
    snippets (title, content, created, expires)
VALUES
    (
        'Over the wintry forest',
        'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Never to be seen again.',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP + INTERVAL '1 year'
    );

INSERT INTO
    snippets (title, content, created, expires)
VALUES
    (
        'First autumn morning',
        'First autumn morning\nthe mirror I stare into\nshows my father\'' s face.\ n \ nIt was uncanny,
        like a mirror.',
CURRENT_TIMESTAMP,
CURRENT_TIMESTAMP + INTERVAL ' 1 year '
);