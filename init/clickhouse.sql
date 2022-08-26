CREATE database "pr-test-3" ENGINE = "Atomic";


CREATE TABLE `pr-test-3`.kafka (cmd String, username String, timestamp Int64)
    ENGINE = Kafka
    SETTINGS kafka_broker_list = '192.168.3.78:9092',
             kafka_topic_list = 'pr-test-3',
             kafka_group_name = 'group1',
             kafka_format = 'JSONEachRow',
             kafka_max_block_size = 1048576;
	
	
CREATE TABLE `pr-test-3`.view (cmd String, username String, timestamp Int64)
    ENGINE = MergeTree()
    ORDER BY cmd;
	
	
CREATE MATERIALIZED VIEW `pr-test-3`.consumer TO `pr-test-3`.view AS
SELECT
	FROM_UNIXTIME(timestamp) as "timestamp",
	cmd,
	username
FROM
	`pr-test-3`.kafka;