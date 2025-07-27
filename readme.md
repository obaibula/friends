# Simple example of neo4j usage

I chose Neo4j because friend relationships are naturally a graph. This allows efficient querying of multi-level connections,
recommendations, and cycles, things that would require expensive joins or recursive CTEs in SQL.

1. To access a neo4j query browser go to `localhost:7474`
2. To generate grpc stabs use buf generate
3. Just fire up a server and restgw reverse restful proxy, openapi spec is in [openapi spec](./proto/friend.swagger.json). This is just a test project to leverage my graph db and grpc knoweledge.
