CREATE TABLE public.coordinates (
	id bigint NOT NULL,
	country varchar NOT NULL,
	city varchar NOT NULL,
	latitude float8 NOT NULL,
	longitude float8 NOT NULL,
	CONSTRAINT coordinates_pk PRIMARY KEY (id)
);

INSERT INTO coordinates (id, country, city, latitude, longitude) VALUES(1, 'USA', 'New York', 40.7128, 74.006);
INSERT INTO coordinates (id, country, city, latitude, longitude) VALUES(2, 'Russia', 'Moscow', 55.7558, 37.6173);
INSERT INTO coordinates (id, country, city, latitude, longitude) VALUES(3, 'South Africa', 'Cape Town', 33.9249, 18.4241);