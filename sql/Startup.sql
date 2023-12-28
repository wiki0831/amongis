CREATE EXTENSION postgis;
-- Enable sfcgal functions
CREATE EXTENSION postgis_sfcgal;
-- Performance related, you can turn it on of off and use 'explain analyze' on geometrydump.geom 
SET jit = off;