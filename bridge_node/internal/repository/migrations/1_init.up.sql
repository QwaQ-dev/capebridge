CREATE TYPE bridge_status AS ENUM (
    'detected',
    'confirmed',
    'relaying',
    'relayed',
    'failed',
    'voted'
);

CREATE TABLE IF NOT EXISTS bridge_events (
    id BIGSERIAL PRIMARY KEY,

    source_chain TEXT NOT NULL,
    target_chain TEXT NOT NULL,

    tx_hash TEXT NOT NULL,
    log_index INTEGER NOT NULL,

    block_number BIGINT NOT NULL,
    block_hash TEXT NOT NULL DEFAULT '',

    sender TEXT NOT NULL,
    receiver TEXT NOT NULL,

    amount NUMERIC(78,0) NOT NULL,
    nonce BIGINT NULL,

    status bridge_status NOT NULL DEFAULT 'detected',

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE(tx_hash, log_index),
    UNIQUE(source_chain, nonce)
);
 

CREATE TABLE IF NOT EXISTS indexer_state (
    chain TEXT PRIMARY KEY,
    last_block BIGINT NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE INDEX IF NOT EXISTS idx_bridge_status
ON bridge_events(status);

CREATE INDEX IF NOT EXISTS idx_bridge_block
ON bridge_events(block_number);

CREATE INDEX IF NOT EXISTS idx_bridge_nonce
ON bridge_events(nonce);

CREATE INDEX IF NOT EXISTS idx_bridge_source_chain
ON bridge_events(source_chain);

CREATE INDEX IF NOT EXISTS idx_bridge_target_chain
ON bridge_events(target_chain);