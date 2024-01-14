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

CREATE TABLE IF NOT EXISTS template_block_headers (
    hash        VARCHAR PRIMARY KEY,
    parent_hash VARCHAR NOT NULL UNIQUE,
    number      UINT256 NOT NULL UNIQUE,
    timestamp   INTEGER NOT NULL UNIQUE ,
    rlp_bytes   VARCHAR NOT NULL
);
CREATE INDEX IF NOT EXISTS template_block_headers_timestamp ON template_block_headers(timestamp);
CREATE INDEX IF NOT EXISTS template_block_headers_number ON template_block_headers(number);


CREATE TABLE IF NOT EXISTS template_contract_events (
    guid             VARCHAR PRIMARY KEY,
    block_hash       VARCHAR NOT NULL REFERENCES template_block_headers(hash) ON DELETE CASCADE,
    contract_address VARCHAR NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    log_index        INTEGER NOT NULL,
    event_signature  VARCHAR NOT NULL,
    timestamp        INTEGER NOT NULL ,
    rlp_bytes        VARCHAR NOT NULL
);
CREATE INDEX IF NOT EXISTS template_contract_events_timestamp ON template_contract_events(timestamp);
CREATE INDEX IF NOT EXISTS template_contract_events_block_hash ON template_contract_events(block_hash);
CREATE INDEX IF NOT EXISTS template_contract_events_event_signature ON template_contract_events(event_signature);
CREATE INDEX IF NOT EXISTS template_contract_events_contract_address ON template_contract_events(contract_address);

CREATE TABLE IF NOT EXISTS template_transactions (
     guid                      VARCHAR PRIMARY KEY,
     block_hash                VARCHAR NOT NULL,
     block_number              UINT256 DEFAULT 0,
     from_address              VARCHAR NOT NULL,
     to_address                VARCHAR,
     gas                       UINT256 NOT NULL,
     gas_price                 UINT256 NOT NULL,
     hash                      VARCHAR NOT NULL,
     input_data                VARCHAR,
     max_fee_per_gas           UINT256 DEFAULT 0,
     max_priority_fee_per_gas  UINT256 DEFAULT 0,
     gas_used                  UINT256 DEFAULT 0,
     cumulative_gas_used       UINT256 DEFAULT 0,
     effective_gas_price       UINT256 DEFAULT 0,
     l1_fee                    UINT256 DEFAULT 0,
     L1_gas_used               UINT256 DEFAULT 0,
     l1_gas_price              UINT256 DEFAULT 0,
     nonce                     UINT256,
     transaction_index         UINT256 NOT NULL,
     tx_type                   SMALLINT NOT NULL default 0,
     r                         VARCHAR NOT NULL,
     s                         VARCHAR NOT NULL,
     v                         VARCHAR NOT NULL,
     status                    SMALLINT NOT NULL default 0,
     contract_address          VARCHAR,
     amount                    UINT256 DEFAULT 0,
     y_parity                  VARCHAR NOT NULL,
     timestamp                 INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS template_transactions_block_number ON template_transactions(block_number);
CREATE INDEX IF NOT EXISTS template_transactions_hash ON template_transactions(hash);
CREATE INDEX IF NOT EXISTS template_transactions_timestamp ON template_transactions(timestamp);


CREATE TABLE IF NOT EXISTS template_l1_to_l2 (
    guid                    VARCHAR PRIMARY KEY,
    l1_block_number         UINT256 NOT NULL,
    l2_block_number         UINT256,
    queue_index             UINT256,
    l1_transaction_hash     VARCHAR NOT NULL,
    l2_transaction_hash     VARCHAR NOT NULL,
    transaction_source_hash VARCHAR NOT NULL,
    message_hash            VARCHAR,
    l1_tx_origin            VARCHAR,
    token_ids               VARCHAR,
    claimed_index           UINT256,
    status                  SMALlINT NOT NULL,
    from_address            VARCHAR NOT NULL,
    to_address              VARCHAR NOT NULL,
    l1_token_address        VARCHAR,
    l2_token_address        VARCHAR,
    asset_type              SMALLINT NOT NULL,
    eth_amount              UINT256,
    token_amounts           VARCHAR,
    gas_limit               UINT256 NOT NULL,
    timestamp               INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_timestamp ON template_l1_to_l2(timestamp);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_l1_transaction_hash ON template_l1_to_l2(l1_transaction_hash);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_l2_transaction_hash ON template_l1_to_l2(l2_transaction_hash);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_from_address ON template_l1_to_l2(from_address);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_to_address ON template_l1_to_l2(to_address);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_message_hash ON template_l1_to_l2(message_hash);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_transaction_source_hash ON template_l1_to_l2(transaction_source_hash);

CREATE TABLE IF NOT EXISTS template_withdraw_proven (
   guid                          VARCHAR PRIMARY KEY,
   block_number                  UINT256 NOT NULL,
   withdraw_hash                 VARCHAR NOT NULL,
   message_hash                  VARCHAR,
   proven_transaction_hash       VARCHAR NOT NULL,
   l1_token_address              VARCHAR,
   l2_token_address              VARCHAR,
   eth_amount                    UINT256,
   erc20_amount                  UINT256,
   related                       BOOLEAN DEFAULT FALSE,
   timestamp                     INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS template_withdraw_proven_withdrawal_hash ON template_withdraw_proven(withdraw_hash);
CREATE INDEX IF NOT EXISTS template_withdraw_proven_timestamp ON template_withdraw_proven(timestamp);

CREATE TABLE IF NOT EXISTS template_withdraw_finalized (
  guid                          VARCHAR PRIMARY KEY,
  block_number                  UINT256 NOT NULL,
  withdraw_hash                 VARCHAR NOT NULL,
  message_hash                  VARCHAR,
  finalized_transaction_hash    VARCHAR NOT NULL,
  l1_token_address              VARCHAR,
  l2_token_address              VARCHAR,
  eth_amount                    UINT256,
  erc20_amount                  UINT256,
  related                       BOOLEAN DEFAULT FALSE,
  timestamp                     INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS template_withdraw_finalized_withdrawal_hash ON template_withdraw_finalized(withdraw_hash);
CREATE INDEX IF NOT EXISTS template_withdraw_finalized_timestamp ON template_withdraw_finalized(timestamp);


CREATE TABLE IF NOT EXISTS template_l2_to_l1 (
    guid                         VARCHAR PRIMARY KEY,
    l1_block_number              UINT256,
    l2_block_number              UINT256 NOT NULL,
    msg_nonce                    UINT256 NOT NULL,
    l2_transaction_hash          VARCHAR NOT NULL,
    withdraw_transaction_hash    VARCHAR NOT NULL,
    message_hash                 VARCHAR,
    l1_prove_tx_hash             VARCHAR,
    l1_finalize_tx_hash          VARCHAR,
    token_ids                    VARCHAR,
    claimed_index                UINT256,
    status                       SMALLINT NOT NULL,
    from_address                 VARCHAR NOT NULL,
    to_address                   VARCHAR NOT NULL,
    l1_token_address             VARCHAR,
    l2_token_address             VARCHAR,
    asset_type                   SMALLINT NOT NULL,
    eth_amount                   UINT256,
    token_amounts                VARCHAR,
    gas_limit                    UINT256 NOT NULL,
    time_left                    UINT256 NOT NULL,
    timestamp                    INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_timestamp ON template_l2_to_l1(timestamp);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_l2_transaction_hash ON template_l2_to_l1(l2_transaction_hash);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_l1_prove_tx_hash ON template_l2_to_l1(l1_prove_tx_hash);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_l1_finalize_tx_hash ON template_l2_to_l1(l1_finalize_tx_hash);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_from_address ON template_l2_to_l1(from_address);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_to_address ON template_l2_to_l1(to_address);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_message_hash ON template_l2_to_l1(message_hash);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_withdraw_transaction_hash ON template_l2_to_l1(withdraw_transaction_hash);


CREATE TABLE IF NOT EXISTS template_batches (
   guid                      VARCHAR PRIMARY KEY,
   batch_index               UINT256 NOT NULL,
   block_hash                VARCHAR NOT NULL,
   transaction_hash          VARCHAR NOT NULL,
   l1_block_number           UINT256 DEFAULT 0,
   l2_block_number           UINT256 DEFAULT 0,
   tx_size                   UINT256 NOT NULL,
   block_size                UINT256 NOT NULL,
   timestamp                 INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS template_batches_block_hash ON template_batches(block_hash);
CREATE INDEX IF NOT EXISTS template_batches_transaction_hash ON template_batches(transaction_hash);


CREATE TABLE IF NOT EXISTS template_state_root (
      guid                      VARCHAR PRIMARY KEY,
      block_hash                VARCHAR NOT NULL,
      transaction_hash          VARCHAR NOT NULL,
      l1_block_number           UINT256 DEFAULT 0,
      l2_block_number           UINT256 DEFAULT 0,
      output_index              UINT256 NOT NULL,
      prev_total_elements       UINT256 DEFAULT 0,
      status                    SMALLINT NOT NULL default 0,
      output_root               VARCHAR NOT NULL,
      canonical                 BOOLEAN DEFAULT TRUE,
      batch_size                UINT256 NOT NULL,
      block_size                UINT256 NOT NULL,
      timestamp                 INTEGER NOT NULL
 );
CREATE INDEX IF NOT EXISTS template_state_root_block_hash ON template_state_root(block_hash);
CREATE INDEX IF NOT EXISTS template_state_root_transaction_hash ON template_state_root(transaction_hash);

