
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'uint256') THEN
        CREATE DOMAIN UINT256 AS NUMERIC
            CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
    ELSE
        ALTER DOMAIN UINT256 DROP CONSTRAINT uint256_check;
        ALTER DOMAIN UINT256 ADD
            CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
    END IF;
END $$;


CREATE TABLE IF NOT EXISTS l1_block_headers (
    hash        VARCHAR PRIMARY KEY,
    parent_hash VARCHAR   UNIQUE,
    number      UINT256  DEFAULT 0 ,
    timestamp   INTEGER  DEFAULT 0 UNIQUE CHECK (timestamp > 0),
    rlp_bytes VARCHAR
);
CREATE INDEX IF NOT EXISTS l1_block_headers_timestamp ON l1_block_headers(timestamp);
CREATE INDEX IF NOT EXISTS l1_block_headers_number ON l1_block_headers(number);

CREATE TABLE IF NOT EXISTS l2_block_headers (
    hash        VARCHAR PRIMARY KEY,
    parent_hash VARCHAR   UNIQUE,
    number      UINT256  DEFAULT 0   UNIQUE,
    timestamp   INTEGER  DEFAULT 0,
    rlp_bytes VARCHAR
);
CREATE INDEX IF NOT EXISTS l2_block_headers_timestamp ON l2_block_headers(timestamp);
CREATE INDEX IF NOT EXISTS l2_block_headers_number ON l2_block_headers(number);

CREATE TABLE IF NOT EXISTS l1_contract_events (
    guid             VARCHAR PRIMARY KEY,
    block_hash       VARCHAR   REFERENCES l1_block_headers(hash) ON DELETE CASCADE,
    contract_address VARCHAR  ,
    transaction_hash VARCHAR  ,
    log_index        INTEGER  DEFAULT 0,
    event_signature  VARCHAR  ,
    timestamp        INTEGER  DEFAULT 0 CHECK (timestamp > 0),
    rlp_bytes VARCHAR
);
CREATE INDEX IF NOT EXISTS l1_contract_events_timestamp ON l1_contract_events(timestamp);
CREATE INDEX IF NOT EXISTS l1_contract_events_block_hash ON l1_contract_events(block_hash);
CREATE INDEX IF NOT EXISTS l1_contract_events_event_signature ON l1_contract_events(event_signature);
CREATE INDEX IF NOT EXISTS l1_contract_events_contract_address ON l1_contract_events(contract_address);

CREATE TABLE IF NOT EXISTS l2_contract_events (
    guid             VARCHAR PRIMARY KEY,
    block_hash       VARCHAR   REFERENCES l2_block_headers(hash) ON DELETE CASCADE,
    contract_address VARCHAR  ,
    transaction_hash VARCHAR  ,
    log_index        INTEGER  DEFAULT 0,
    event_signature  VARCHAR  ,
    timestamp        INTEGER  DEFAULT 0 CHECK (timestamp > 0),
    rlp_bytes VARCHAR
);
CREATE INDEX IF NOT EXISTS l2_contract_events_timestamp ON l2_contract_events(timestamp);
CREATE INDEX IF NOT EXISTS l2_contract_events_block_hash ON l2_contract_events(block_hash);
CREATE INDEX IF NOT EXISTS l2_contract_events_event_signature ON l2_contract_events(event_signature);
CREATE INDEX IF NOT EXISTS l2_contract_events_contract_address ON l2_contract_events(contract_address);


CREATE TABLE IF NOT EXISTS l1_transaction_deposits (
    source_hash             VARCHAR PRIMARY KEY,
    l2_transaction_hash     VARCHAR   ,
    initiated_l1_event_guid VARCHAR   UNIQUE REFERENCES l1_contract_events(guid) ON DELETE CASCADE,
    from_address VARCHAR  ,
    to_address   VARCHAR  ,
    eth_amount       UINT256  DEFAULT 0,
    erc20_amount   UINT256  DEFAULT 0,
    gas_limit    UINT256  DEFAULT 0,
    data         VARCHAR  ,
    timestamp    INTEGER  DEFAULT 0 CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS l1_transaction_deposits_timestamp ON l1_transaction_deposits(timestamp);
CREATE INDEX IF NOT EXISTS l1_transaction_deposits_initiated_l1_event_guid ON l1_transaction_deposits(initiated_l1_event_guid);
CREATE INDEX IF NOT EXISTS l1_transaction_deposits_from_address ON l1_transaction_deposits(from_address);

CREATE TABLE IF NOT EXISTS l2_transaction_withdrawals (
    withdrawal_hash         VARCHAR PRIMARY KEY,
    nonce                   UINT256  DEFAULT 0 UNIQUE,
    initiated_l2_event_guid VARCHAR   UNIQUE REFERENCES l2_contract_events(guid) ON DELETE CASCADE,
    proven_l1_event_guid    VARCHAR UNIQUE REFERENCES l1_contract_events(guid) ON DELETE CASCADE,
    finalized_l1_event_guid VARCHAR UNIQUE REFERENCES l1_contract_events(guid) ON DELETE CASCADE,
    succeeded               BOOLEAN,
    from_address VARCHAR  ,
    to_address   VARCHAR  ,
    eth_amount       UINT256  DEFAULT 0,
    erc20_amount   UINT256  DEFAULT 0,
    gas_limit    UINT256  DEFAULT 0,
    data         VARCHAR  ,
    timestamp    INTEGER  DEFAULT 0 CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS l2_transaction_withdrawals_timestamp ON l2_transaction_withdrawals(timestamp);
CREATE INDEX IF NOT EXISTS l2_transaction_withdrawals_initiated_l2_event_guid ON l2_transaction_withdrawals(initiated_l2_event_guid);
CREATE INDEX IF NOT EXISTS l2_transaction_withdrawals_from_address ON l2_transaction_withdrawals(from_address);

CREATE TABLE IF NOT EXISTS l1_bridge_messages(
    message_hash            VARCHAR PRIMARY KEY,
    nonce                   UINT256  DEFAULT 0 UNIQUE,
    transaction_source_hash VARCHAR   ,
    sent_message_event_guid    VARCHAR ,
    relayed_message_event_guid VARCHAR ,
    from_address VARCHAR  ,
    to_address   VARCHAR  ,
    eth_amount       UINT256  DEFAULT 0,
    erc20_amount     UINT256  DEFAULT 0,
    gas_limit    UINT256  DEFAULT 0,
    data         VARCHAR  ,
    timestamp    INTEGER  DEFAULT 0 CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS l1_bridge_messages_timestamp ON l1_bridge_messages(timestamp);
CREATE INDEX IF NOT EXISTS l1_bridge_messages_transaction_source_hash ON l1_bridge_messages(transaction_source_hash);
CREATE INDEX IF NOT EXISTS l1_bridge_messages_from_address ON l1_bridge_messages(from_address);

CREATE TABLE IF NOT EXISTS l2_bridge_messages(
    message_hash                VARCHAR PRIMARY KEY,
    nonce                       UINT256  DEFAULT 0 UNIQUE,
    transaction_withdrawal_hash VARCHAR   UNIQUE REFERENCES l2_transaction_withdrawals(withdrawal_hash) ON DELETE CASCADE,
    sent_message_event_guid    VARCHAR   UNIQUE REFERENCES l2_contract_events(guid) ON DELETE CASCADE,
    relayed_message_event_guid VARCHAR UNIQUE REFERENCES l1_contract_events(guid) ON DELETE CASCADE,
    from_address VARCHAR  ,
    to_address   VARCHAR  ,
    eth_amount       UINT256  DEFAULT 0,
    erc20_amount     UINT256  DEFAULT 0,
    gas_limit    UINT256  DEFAULT 0,
    data         VARCHAR  ,
    timestamp    INTEGER  DEFAULT 0 CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS l2_bridge_messages_timestamp ON l2_bridge_messages(timestamp);
CREATE INDEX IF NOT EXISTS l2_bridge_messages_transaction_withdrawal_hash ON l2_bridge_messages(transaction_withdrawal_hash);
CREATE INDEX IF NOT EXISTS l2_bridge_messages_from_address ON l2_bridge_messages(from_address);

CREATE TABLE IF NOT EXISTS l1_bridged_tokens (
    address        VARCHAR PRIMARY KEY,
    bridge_address VARCHAR  ,

    name     VARCHAR  ,
    symbol   VARCHAR  ,
    decimals INTEGER  DEFAULT 0 CHECK (decimals >= 0 AND decimals <= 18)
);
CREATE TABLE IF NOT EXISTS l2_bridged_tokens (
    address        VARCHAR PRIMARY KEY,
    bridge_address VARCHAR  ,

    l1_token_address VARCHAR REFERENCES l1_bridged_tokens(address) ON DELETE CASCADE,

    name     VARCHAR  ,
    symbol   VARCHAR  ,
    decimals INTEGER  DEFAULT 0 CHECK (decimals >= 0 AND decimals <= 18)
);

CREATE TABLE IF NOT EXISTS l1_bridge_deposits (
    transaction_source_hash   VARCHAR PRIMARY KEY REFERENCES l1_transaction_deposits(source_hash) ON DELETE CASCADE,
    cross_domain_message_hash VARCHAR   UNIQUE REFERENCES l1_bridge_messages(message_hash) ON DELETE CASCADE,
    from_address              VARCHAR  ,
    to_address                VARCHAR  ,
    local_token_address       VARCHAR  ,
    remote_token_address      VARCHAR  ,
    eth_amount                UINT256  DEFAULT 0,
    erc20_amount              UINT256  DEFAULT 0,
    data                      VARCHAR  ,
    timestamp                 INTEGER  DEFAULT 0 CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS l1_bridge_deposits_timestamp ON l1_bridge_deposits(timestamp);
CREATE INDEX IF NOT EXISTS l1_bridge_deposits_cross_domain_message_hash ON l1_bridge_deposits(cross_domain_message_hash);
CREATE INDEX IF NOT EXISTS l1_bridge_deposits_from_address ON l1_bridge_deposits(from_address);

CREATE TABLE IF NOT EXISTS l2_bridge_withdrawals (
    transaction_withdrawal_hash VARCHAR PRIMARY KEY REFERENCES l2_transaction_withdrawals(withdrawal_hash) ON DELETE CASCADE,
    cross_domain_message_hash   VARCHAR   UNIQUE REFERENCES l2_bridge_messages(message_hash) ON DELETE CASCADE,
    from_address                VARCHAR  ,
    to_address                  VARCHAR  ,
    local_token_address         VARCHAR  ,
    remote_token_address        VARCHAR  ,
    eth_amount                  UINT256  DEFAULT 0,
    erc20_amount                UINT256  DEFAULT 0,
    data                        VARCHAR  ,
    timestamp                   INTEGER  DEFAULT 0 CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS l2_bridge_withdrawals_timestamp ON l2_bridge_withdrawals(timestamp);
CREATE INDEX IF NOT EXISTS l2_bridge_withdrawals_cross_domain_message_hash ON l2_bridge_withdrawals(cross_domain_message_hash);
CREATE INDEX IF NOT EXISTS l2_bridge_withdrawals_from_address ON l2_bridge_withdrawals(from_address);

CREATE TABLE IF NOT EXISTS confirm_message (
    data_store_id  UINT256 PRIMARY KEY  UNIQUE REFERENCES confirm_message(data_store_id) ON DELETE CASCADE,
    header_hash    VARCHAR  ,
    timestamp      INTEGER  DEFAULT 0 CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS confirm_message_data_store_id ON confirm_message(data_store_id);
CREATE INDEX IF NOT EXISTS confirm_message_timestamp ON confirm_message(timestamp);

CREATE TABLE IF NOT EXISTS l1_to_l2 (
    guid             VARCHAR PRIMARY KEY,
    l1_block_hash     VARCHAR ,
    queue_index         UINT256,
    l1_transaction_hash VARCHAR  ,
    l2_transaction_hash VARCHAR  ,
    l1_tx_origin        VARCHAR,
    status              SMALlINT ,
    from_address        VARCHAR  ,
    to_address          VARCHAR  ,
    l1_token_address    VARCHAR,
    l2_token_address    VARCHAR,
    eth_amount          UINT256,
    erc20_amount        UINT256,
    gas_limit           UINT256  DEFAULT 0,
    timestamp           INTEGER  DEFAULT 0 CHECK (timestamp > 0),
    asset_type SMALlINT,
    token_amounts VARCHAR  ,
    msg_hash VARCHAR  ,
    token_ids VARCHAR
);
CREATE INDEX IF NOT EXISTS l1_to_l2_timestamp ON l1_to_l2(timestamp);
CREATE INDEX IF NOT EXISTS l1_to_l2_l1_transaction_hash ON l1_to_l2(l1_transaction_hash);
CREATE INDEX IF NOT EXISTS l1_to_l2_l2_transaction_hash ON l1_to_l2(l2_transaction_hash);

CREATE TABLE IF NOT EXISTS l2_to_l1 (
    guid             VARCHAR PRIMARY KEY,
    l2_block_hash              VARCHAR ,
    msg_nonce                    UINT256  DEFAULT 0,
    l2_transaction_hash          VARCHAR  ,
    l2_withdraw_transaction_hash VARCHAR  ,
    l1_prove_tx_hash             VARCHAR,
    l1_finalize_tx_hash          VARCHAR,
    status                       SMALLINT ,
    from_address                 VARCHAR  ,
    to_address                   VARCHAR  ,
    l1_token_address             VARCHAR,
    l2_token_address             VARCHAR,
    eth_amount                   UINT256,
    erc20_amount                 UINT256,
    gas_limit                    UINT256  DEFAULT 0,
    time_left                    UINT256  DEFAULT 0,
    timestamp                    INTEGER  DEFAULT 0 CHECK (timestamp > 0),
    asset_type SMALlINT,
    token_amounts VARCHAR  ,
    msg_hash VARCHAR  ,
    token_ids VARCHAR
);
CREATE INDEX IF NOT EXISTS l2_to_l1_timestamp ON l2_to_l1(timestamp);
CREATE INDEX IF NOT EXISTS l2_to_l1_l2_transaction_hash ON l2_to_l1(l2_transaction_hash);
CREATE INDEX IF NOT EXISTS l2_to_l1_l1_prove_tx_hash ON l2_to_l1(l1_prove_tx_hash);
CREATE INDEX IF NOT EXISTS l2_to_l1_l1_finalize_tx_hash ON l2_to_l1(l1_finalize_tx_hash);

CREATE TABLE IF NOT EXISTS state_root (
    block_hash                VARCHAR PRIMARY KEY  UNIQUE REFERENCES state_root(block_hash) ON DELETE CASCADE,
    transaction_hash          VARCHAR   UNIQUE REFERENCES state_root(transaction_hash) ON DELETE CASCADE,
    l1_block_number           UINT256 DEFAULT 0,
    l2_block_number           UINT256 DEFAULT 0,
    output_index              UINT256  DEFAULT 0,
    prev_total_elements       UINT256 DEFAULT 0,
    status                    SMALLINT  default 0,
    output_root               VARCHAR   UNIQUE REFERENCES state_root(output_root) ON DELETE CASCADE,
    canonical                 BOOLEAN DEFAULT TRUE,
    batch_size                UINT256  DEFAULT 0,
    timestamp                 INTEGER  DEFAULT 0 CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS state_root_block_hash ON state_root(block_hash);
CREATE INDEX IF NOT EXISTS state_root_transaction_hash ON state_root(transaction_hash);

