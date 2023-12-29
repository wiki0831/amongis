-- enable extensions 
CREATE EXTENSION postgis;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION postgis_sfcgal;
-- Performance related, you can turn it on of off and use 'explain analyze' on geometrydump.geom 
SET jit = off;
-- Create Table
CREATE TABLE player (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(100),
    role VARCHAR(50),
    room VARCHAR(50),
    status VARCHAR(50),
    location GEOMETRY(Point, 4326)
);
-- Create latest player view
CREATE VIEW latest_player_data AS
SELECT DISTINCT ON (name) *
FROM player
ORDER BY name,
    created_at DESC;