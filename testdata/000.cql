CREATE KEYSPACE test_keyspace_1 WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
CREATE ROLE IF NOT EXISTS test_role_1 WITH PASSWORD = 'test1' AND LOGIN = true;
GRANT ALL ON KEYSPACE test_keyspace_1 TO test_role_1;

