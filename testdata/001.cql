CREATE KEYSPACE test_keyspace_2 WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
CREATE ROLE IF NOT EXISTS test_role_2 WITH PASSWORD = 'test2' AND LOGIN = true;
GRANT ALL ON KEYSPACE test_keyspace_2 TO test_role_2;
