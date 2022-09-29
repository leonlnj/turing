-- Create type for Router and Route protocol
CREATE TYPE RouteProtocol as ENUM ('HTTP', 'GRPC');
CREATE TYPE RouterProtocol as ENUM ('HTTP', 'UPI_V1');

-- Add Router protocol
ALTER TABLE router_versions ADD protocol RouterProtocol NOT NULL DEFAULT 'HTTP';

-- Add protocol to individual routes of router_versions.routes
WITH individual_route AS (
    SELECT id, jsonb_array_elements(routes) || jsonb '{"protocol":"HTTP"}' as route FROM router_versions
),
updated_route AS (
    SELECT id, json_agg(route) as updated_routes from individual_route
    GROUP BY individual_route.id
)
UPDATE router_versions
SET routes = updated_routes
FROM updated_route;