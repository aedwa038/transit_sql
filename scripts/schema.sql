CREATE TABLE public.agency
(
agency_id VARCHAR (255) NOT NULL,
agency_name VARCHAR (255) NOT NULL,
agency_url VARCHAR (255) NOT NULL,
agency_timezone VARCHAR (255) NOT NULL,
agency_lang VARCHAR (255) NOT NULL,
agency_phone VARCHAR (255)
);

CREATE TABLE public.calendar_dates
(
service_id INT NOT NULL,
date INT NOT NULL,
exception_type INT NOT NULL
);

CREATE TABLE public.routes
(
route_id INT NOT NULL,
agency_id VARCHAR (255) NOT NULL,
route_short_name VARCHAR (255),
route_long_name VARCHAR (255) NOT NULL,
route_type INT NOT NULL,
route_url VARCHAR (255),
route_color VARCHAR (255) NOT NULL
);

CREATE TABLE public.shapes
(
shape_id INT NOT NULL,
shape_pt_lat DECIMAL(10,6) NOT NULL,
shape_pt_lon DECIMAL(10,6) NOT NULL,
shape_pt_sequence INT NOT NULL,
shape_dist_traveled real
);

CREATE TABLE public.stop_times
(
trip_id INT NOT NULL,
arrival_time VARCHAR (255) NOT NULL,
departure_time VARCHAR (255) NOT NULL,
stop_id INT NOT NULL,
stop_sequence INT NOT NULL,
pickup_type INT NOT NULL,
drop_off_type INT NOT NULL,
shape_dist_traveled real
);

CREATE TABLE public.stops
(
stop_id INT NOT NULL,
stop_code INT,
stop_name VARCHAR (255) NOT NULL,
stop_desc VARCHAR (255),
stop_lat DECIMAL(10,6) NOT NULL,
stop_lon DECIMAL(10,6) NOT NULL,
zone_id INT NOT NULL
);

CREATE TABLE public.trips
(
route_id INT NOT NULL,
service_id INT NOT NULL,
trip_id INT NOT NULL,
trip_headsign VARCHAR (255) NOT NULL,
direction_id INT NOT NULL,
block_id VARCHAR (255) NOT NULL,
shape_id INT NOT NULL
);

