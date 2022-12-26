CREATE TABLE food (
    id serial,
    name text,
    primary key (id)
);
CREATE TABLE content (
    food_id integer,
    sequence integer,
    name text,
    primary key (food_id, sequence)
);
INSERT INTO food (name)
VALUES ('aa'),
    ('bb');
INSERT INTO content
VALUES (1, 1, 'aa1'),
    (1, 2, 'aa2'),
    (1, 3, 'aa3'),
    (2, 1, 'bb1'),
    (2, 2, 'bb2');