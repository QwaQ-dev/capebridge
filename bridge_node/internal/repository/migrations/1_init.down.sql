DROP INDEX IF EXISTS idx_bridge_target_chain;
DROP INDEX IF EXISTS idx_bridge_source_chain;
DROP INDEX IF EXISTS idx_bridge_nonce;
DROP INDEX IF EXISTS idx_bridge_block;
DROP INDEX IF EXISTS idx_bridge_status;


DROP TABLE IF EXISTS bridge_events;

DROP TABLE IF EXISTS indexer_state;


DROP TYPE IF EXISTS bridge_status;