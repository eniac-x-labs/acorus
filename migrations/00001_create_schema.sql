DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'uint256') THEN
            CREATE DOMAIN UINT256 AS NUMERIC
                CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
        ELSE
            ALTER DOMAIN UINT256 DROP CONSTRAINT uint256_check;
            ALTER DOMAIN UINT256 ADD
                CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
        END IF;
    END
$$;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" cascade;
-- CREATE EXTENSION "uuid-ossp";

CREATE TABLE IF NOT EXISTS template_block_headers
(
    guid        text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    hash        VARCHAR NOT NULL,
    parent_hash VARCHAR NOT NULL,
    number      UINT256 NOT NULL,
    timestamp   INTEGER NOT NULL,
    rlp_bytes   VARCHAR NOT NULL
);
CREATE INDEX IF NOT EXISTS template_block_headers_timestamp ON template_block_headers (timestamp);
CREATE INDEX IF NOT EXISTS template_block_headers_number ON template_block_headers (number);


CREATE TABLE IF NOT EXISTS template_contract_events
(
    guid             text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    block_hash       VARCHAR NOT NULL,
    contract_address VARCHAR NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    log_index        INTEGER NOT NULL,
    block_number     UINT256 NOT NULL,
    event_signature  VARCHAR NOT NULL,
    timestamp        INTEGER NOT NULL,
    rlp_bytes        VARCHAR NOT NULL
);
CREATE INDEX IF NOT EXISTS template_contract_events_number ON template_contract_events (block_number);
CREATE INDEX IF NOT EXISTS template_contract_events_timestamp ON template_contract_events (timestamp);
CREATE INDEX IF NOT EXISTS template_contract_events_block_hash ON template_contract_events (block_hash);
CREATE INDEX IF NOT EXISTS template_contract_events_event_signature ON template_contract_events (event_signature);
CREATE INDEX IF NOT EXISTS template_contract_events_contract_address ON template_contract_events (contract_address);
CREATE INDEX IF NOT EXISTS template_contract_events_transaction_hash ON template_contract_events (transaction_hash);

CREATE TABLE IF NOT EXISTS template_transactions
(
    guid                     text PRIMARY KEY  DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    block_hash               VARCHAR  NOT NULL,
    block_number             UINT256           DEFAULT 0,
    from_address             VARCHAR  NOT NULL,
    to_address               VARCHAR,
    gas                      UINT256  NOT NULL,
    gas_price                UINT256  NOT NULL,
    hash                     VARCHAR  NOT NULL,
    input_data               VARCHAR,
    max_fee_per_gas          UINT256           DEFAULT 0,
    max_priority_fee_per_gas UINT256           DEFAULT 0,
    gas_used                 UINT256           DEFAULT 0,
    cumulative_gas_used      UINT256           DEFAULT 0,
    effective_gas_price      UINT256           DEFAULT 0,
    l1_fee                   UINT256           DEFAULT 0,
    L1_gas_used              UINT256           DEFAULT 0,
    l1_gas_price             UINT256           DEFAULT 0,
    nonce                    UINT256,
    transaction_index        UINT256  NOT NULL,
    tx_type                  SMALLINT NOT NULL default 0,
    r                        VARCHAR  NOT NULL,
    s                        VARCHAR  NOT NULL,
    v                        VARCHAR  NOT NULL,
    status                   SMALLINT NOT NULL default 0,
    contract_address         VARCHAR,
    amount                   UINT256           DEFAULT 0,
    y_parity                 VARCHAR  NOT NULL,
    timestamp                INTEGER  NOT NULL
);
CREATE INDEX IF NOT EXISTS template_transactions_block_number ON template_transactions (block_number);
CREATE INDEX IF NOT EXISTS template_transactions_hash ON template_transactions (hash);
CREATE INDEX IF NOT EXISTS template_transactions_timestamp ON template_transactions (timestamp);


CREATE TABLE IF NOT EXISTS template_l1_to_l2
(
    guid                    text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    l1_block_number         UINT256  NOT NULL,
    l2_block_number         UINT256,
    queue_index             UINT256,
    l1_transaction_hash     VARCHAR  NOT NULL,
    l2_transaction_hash     VARCHAR  NOT NULL,
    transaction_source_hash VARCHAR  NOT NULL,
    message_hash            VARCHAR,
    l1_tx_origin            VARCHAR,
    token_ids               VARCHAR,
    claimed_index           UINT256,
    status                  SMALlINT NOT NULL,
    from_address            VARCHAR  NOT NULL,
    to_address              VARCHAR  NOT NULL,
    l1_token_address        VARCHAR,
    l2_token_address        VARCHAR,
    asset_type              SMALLINT NOT NULL,
    eth_amount              UINT256,
    token_amounts           VARCHAR,
    gas_limit               UINT256  NOT NULL,
    timestamp               INTEGER  NOT NULL
);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_timestamp ON template_l1_to_l2 (timestamp);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_l1_transaction_hash ON template_l1_to_l2 (l1_transaction_hash);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_l2_transaction_hash ON template_l1_to_l2 (l2_transaction_hash);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_from_address ON template_l1_to_l2 (from_address);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_to_address ON template_l1_to_l2 (to_address);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_message_hash ON template_l1_to_l2 (message_hash);
CREATE INDEX IF NOT EXISTS template_l1_to_l2_transaction_source_hash ON template_l1_to_l2 (transaction_source_hash);

CREATE TABLE IF NOT EXISTS template_withdraw_proven
(
    guid                    text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    block_number            UINT256 NOT NULL,
    withdraw_hash           VARCHAR NOT NULL,
    message_hash            VARCHAR,
    proven_transaction_hash VARCHAR NOT NULL,
    l1_token_address        VARCHAR,
    l2_token_address        VARCHAR,
    eth_amount              UINT256,
    erc20_amount            UINT256,
    related                 BOOLEAN          DEFAULT FALSE,
    timestamp               INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS template_withdraw_proven_withdrawal_hash ON template_withdraw_proven (withdraw_hash);
CREATE INDEX IF NOT EXISTS template_withdraw_proven_timestamp ON template_withdraw_proven (timestamp);

CREATE TABLE IF NOT EXISTS template_withdraw_finalized
(
    guid                       text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    block_number               UINT256 NOT NULL,
    withdraw_hash              VARCHAR NOT NULL,
    message_hash               VARCHAR,
    finalized_transaction_hash VARCHAR NOT NULL,
    l1_token_address           VARCHAR,
    l2_token_address           VARCHAR,
    eth_amount                 UINT256,
    erc20_amount               UINT256,
    related                    BOOLEAN          DEFAULT FALSE,
    timestamp                  INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS template_withdraw_finalized_withdrawal_hash ON template_withdraw_finalized (withdraw_hash);
CREATE INDEX IF NOT EXISTS template_withdraw_finalized_timestamp ON template_withdraw_finalized (timestamp);


CREATE TABLE IF NOT EXISTS template_l2_to_l1
(
    guid                      text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    l1_block_number           UINT256,
    l2_block_number           UINT256  NOT NULL,
    msg_nonce                 UINT256  NOT NULL,
    l2_transaction_hash       VARCHAR  NOT NULL,
    withdraw_transaction_hash VARCHAR  NOT NULL,
    message_hash              VARCHAR,
    l1_prove_tx_hash          VARCHAR,
    l1_finalize_tx_hash       VARCHAR,
    token_ids                 VARCHAR,
    claimed_index             UINT256,
    status                    SMALLINT NOT NULL,
    from_address              VARCHAR  NOT NULL,
    to_address                VARCHAR  NOT NULL,
    l1_token_address          VARCHAR,
    l2_token_address          VARCHAR,
    asset_type                SMALLINT NOT NULL,
    eth_amount                UINT256,
    token_amounts             VARCHAR,
    gas_limit                 UINT256  NOT NULL,
    time_left                 UINT256  NOT NULL,
    version                   INTEGER          DEFAULT 0,
    timestamp                 INTEGER  NOT NULL
);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_timestamp ON template_l2_to_l1 (timestamp);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_l2_transaction_hash ON template_l2_to_l1 (l2_transaction_hash);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_l1_prove_tx_hash ON template_l2_to_l1 (l1_prove_tx_hash);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_l1_finalize_tx_hash ON template_l2_to_l1 (l1_finalize_tx_hash);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_from_address ON template_l2_to_l1 (from_address);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_to_address ON template_l2_to_l1 (to_address);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_message_hash ON template_l2_to_l1 (message_hash);
CREATE INDEX IF NOT EXISTS template_l2_to_l1_withdraw_transaction_hash ON template_l2_to_l1 (withdraw_transaction_hash);


CREATE TABLE IF NOT EXISTS template_batches
(
    guid             text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    batch_index      UINT256 NOT NULL,
    block_hash       VARCHAR NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    l1_block_number  UINT256          DEFAULT 0,
    l2_block_number  UINT256          DEFAULT 0,
    tx_size          UINT256 NOT NULL,
    block_size       UINT256 NOT NULL,
    timestamp        INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS template_batches_block_hash ON template_batches (block_hash);
CREATE INDEX IF NOT EXISTS template_batches_transaction_hash ON template_batches (transaction_hash);


CREATE TABLE IF NOT EXISTS template_state_root
(
    guid                text PRIMARY KEY  DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    block_hash          VARCHAR  NOT NULL,
    transaction_hash    VARCHAR  NOT NULL,
    l1_block_number     UINT256           DEFAULT 0,
    l2_block_number     UINT256           DEFAULT 0,
    output_index        UINT256  NOT NULL,
    prev_total_elements UINT256           DEFAULT 0,
    status              SMALLINT NOT NULL default 0,
    output_root         VARCHAR  NOT NULL,
    canonical           BOOLEAN           DEFAULT TRUE,
    batch_size          UINT256  NOT NULL,
    block_size          UINT256  NOT NULL,
    timestamp           INTEGER  NOT NULL
);
CREATE INDEX IF NOT EXISTS template_state_root_block_hash ON template_state_root (block_hash);
CREATE INDEX IF NOT EXISTS template_state_root_transaction_hash ON template_state_root (transaction_hash);


create table if not exists template_msg_sent
(
    guid              text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash           varchar,
    msg_hash          varchar,
    msg_hash_relation boolean          default false,
    to_l1_l2_table    boolean          default false,
    layer_type        smallint not null,
    data              varchar
);
CREATE INDEX IF NOT EXISTS template_msg_sent_tx_hash ON template_msg_sent (tx_hash);
CREATE INDEX IF NOT EXISTS template_msg_sent_msg_hash ON template_msg_sent (msg_hash);

create table if not exists template_relay
(
    guid           text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash        varchar,
    msg_hash       varchar,
    layer_type     smallint not null,
    l1_l2_relation boolean          default false,
    block_number   UINT256          default 0
);
CREATE INDEX IF NOT EXISTS template_relay_tx_hash ON template_relay (tx_hash);
CREATE INDEX IF NOT EXISTS template_relay_msg_hash ON template_relay (msg_hash);
create table if not exists template_msg_hash
(
    guid     text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash  varchar,
    msg_hash varchar
);
CREATE INDEX IF NOT EXISTS template_msg_hash_tx_hash ON template_msg_hash (tx_hash);
CREATE INDEX IF NOT EXISTS template_msg_hash_msg_hash ON template_msg_hash (msg_hash);


CREATE TABLE IF NOT EXISTS template_relay_message
(
    guid                   text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    block_number           UINT256 NOT NULL,
    relay_transaction_hash VARCHAR NOT NULL,
    message_hash           VARCHAR,
    l1_token_address       VARCHAR,
    l2_token_address       VARCHAR,
    eth_amount             UINT256,
    erc20_amount           UINT256,
    related                BOOLEAN          DEFAULT FALSE,
    timestamp              INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS template_relay_message_message_hash ON template_relay_message (message_hash);
CREATE INDEX IF NOT EXISTS template_relay_message_timestamp ON template_relay_message (timestamp);

CREATE TABLE IF NOT EXISTS staking_record
(
    guid          text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash       VARCHAR  NOT NULL,
    block_number  UINT256  NOT NULL,
    user_address  VARCHAR  not null,
    token         VARCHAR,
    amount        UINT256  not null,
    reward        UINT256,
    start_pool_id uint256,
    end_pool_id   uint256,
    status        smallint not null,
    tx_type       smallint not null,
    asset_type    SMALLINT NOT NULL,
    timestamp     INTEGER  NOT NULL CHECK (timestamp > 0),
    is_points     boolean          default false
);
CREATE INDEX IF NOT EXISTS staking_tx_hash ON staking_record (tx_hash);
CREATE INDEX IF NOT EXISTS staking_block_number ON staking_record (block_number);
CREATE INDEX IF NOT EXISTS staking_user_address ON staking_record (user_address);
CREATE INDEX IF NOT EXISTS staking_token ON staking_record (token);
CREATE INDEX IF NOT EXISTS staking_status ON staking_record (status);
CREATE INDEX IF NOT EXISTS staking_asset_type ON staking_record (asset_type);
CREATE INDEX IF NOT EXISTS staking_timestamp ON staking_record (timestamp);


CREATE TABLE IF NOT EXISTS bridge_record
(
    guid                   text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    source_chain_id        varchar,
    dest_chain_id          varchar,
    source_tx_hash         VARCHAR,
    dest_tx_hash           VARCHAR,
    source_block_number    UINT256,
    dest_block_number      UINT256,
    source_token_address   VARCHAR,
    dest_token_address     VARCHAR,
    msg_hash               varchar,
    from_address           varchar,
    to_address             varchar,
    status                 smallint not null,
    amount                 UINT256,
    nonce                  UINT256,
    fee                    UINT256,
    asset_type             SMALLINT NOT NULL,
    bridge_record_relation boolean          default false,
    msg_sent_timestamp     INTEGER,
    claim_timestamp        INTEGER,
    is_points              boolean          default false
);
CREATE INDEX IF NOT EXISTS bridge_record_source_chain_id ON bridge_record (source_chain_id);
CREATE INDEX IF NOT EXISTS bridge_record_dest_chain_id ON bridge_record (dest_chain_id);
CREATE INDEX IF NOT EXISTS bridge_record_source_tx_hash ON bridge_record (source_tx_hash);
CREATE INDEX IF NOT EXISTS bridge_record_dest_tx_hash ON bridge_record (dest_tx_hash);
CREATE INDEX IF NOT EXISTS bridge_record_msg_hash ON bridge_record (msg_hash);
CREATE INDEX IF NOT EXISTS bridge_record_source_block_number ON bridge_record (source_block_number);
CREATE INDEX IF NOT EXISTS bridge_record_dest_block_number ON bridge_record (dest_block_number);
CREATE INDEX IF NOT EXISTS bridge_record_source_token_address ON bridge_record (source_token_address);
CREATE INDEX IF NOT EXISTS bridge_record_dest_token_address ON bridge_record (dest_token_address);
CREATE INDEX IF NOT EXISTS bridge_record_from ON bridge_record (from_address);
CREATE INDEX IF NOT EXISTS bridge_record_to ON bridge_record (to_address);
CREATE INDEX IF NOT EXISTS bridge_record_status ON bridge_record (status);
CREATE INDEX IF NOT EXISTS bridge_record_asset_type ON bridge_record (asset_type);
CREATE INDEX IF NOT EXISTS bridge_record_msg_sent_timestamp ON bridge_record (msg_sent_timestamp);
CREATE INDEX IF NOT EXISTS bridge_record_claim_timestamp ON bridge_record (claim_timestamp);


create table if not exists bridge_msg_sent
(
    guid                   text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash                varchar,
    msg_hash               varchar,
    dest_hash              varchar,
    dest_block_number      UINT256          default 0,
    dest_timestamp         INTEGER,
    dest_token             varchar,
    fee                    UINT256          default 0,
    msg_nonce              UINT256          default 0,
    msg_hash_relation      boolean          default false,
    bridge_relation        boolean          default false,
    to_bridge_record       boolean          default false,
    to_change_trans_status boolean          default false,
    to_cross_status        boolean          default false,
    data                   varchar
);

CREATE INDEX IF NOT EXISTS bridge_msg_sent_tx_hash ON bridge_msg_sent (tx_hash);
CREATE INDEX IF NOT EXISTS bridge_msg_sent_msg_hash ON bridge_msg_sent (msg_hash);

create table if not exists bridge_msg_hash
(
    guid      text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash   varchar,
    fee       UINT256          default 0,
    msg_nonce UINT256          default 0,
    msg_hash  varchar
);
CREATE INDEX IF NOT EXISTS bridge_msg_hash_tx_hash ON bridge_msg_hash (tx_hash);
CREATE INDEX IF NOT EXISTS bridge_msg_hash_msg_hash ON bridge_msg_hash (msg_hash);

create table if not exists bridge_claim
(
    guid           text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash        varchar,
    msg_hash       varchar,
    dest_token     varchar,
    token_relation boolean          default false,
    timestamp      INTEGER CHECK (timestamp > 0),
    block_number   UINT256          default 0
);

CREATE INDEX IF NOT EXISTS bridge_claim_tx_hash ON bridge_claim (tx_hash);
CREATE INDEX IF NOT EXISTS bridge_claim_msg_hash ON bridge_claim (msg_hash);

create table if not exists bridge_finalize
(
    guid       text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash    varchar,
    dest_token varchar
);

CREATE INDEX IF NOT EXISTS bridge_finalize_tx_hash ON bridge_claim (tx_hash);


create table if not exists bridge_block_listener
(
    guid         text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    chain_id     varchar,
    block_number UINT256          default 0,
    created      INTEGER CHECK (created > 0),
    updated      INTEGER CHECK (updated > 0)
);

create table if not exists bridge_funding_pool_update
(
    guid            text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash         varchar,
    source_chain_id varchar,
    dest_chain_id   varchar,
    receive_address varchar,
    amount          varchar,
    token_address   varchar,
    layer_type      int,
    on_send         boolean          default false
);

CREATE INDEX IF NOT EXISTS bridge_funding_pool_update_tx_hash ON bridge_funding_pool_update (tx_hash);
CREATE INDEX IF NOT EXISTS bridge_funding_pool_update_receive_address ON bridge_funding_pool_update (receive_address);

create table if not exists ac_last_block
(
    guid         text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    chain_id     varchar,
    block_number UINT256          default 0,
    created      INTEGER CHECK (created > 0),
    updated      INTEGER CHECK (updated > 0)
);
CREATE INDEX IF NOT EXISTS ac_last_block_chain_id ON ac_last_block (chain_id);

create table if not exists ac_chain_unstake
(
    guid            text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    block_number    UINT256          default 0,
    tx_hash         varchar,
    eth_amount      UINT256          default 0,
    locked_amount   UINT256          default 0,
    claim_tx_hash   varchar,
    l2_strategy     varchar,
    staker          varchar,
    bridge          varchar,
    source_chain_id varchar,
    dest_chain_id   varchar,
    unstake_nonce UINT256 default 0,
    status          int              default 0,
    notify_relayer  boolean          default false,
    migrate_notify  boolean          default false,
    created         INTEGER CHECK (created > 0),
    updated         INTEGER,
    is_points              boolean          default false

    );
CREATE INDEX IF NOT EXISTS ac_chain_unstake_tx_hash ON ac_chain_unstake (tx_hash);
CREATE INDEX IF NOT EXISTS ac_chain_unstake_l2_strategy ON ac_chain_unstake (l2_strategy);
CREATE INDEX IF NOT EXISTS ac_chain_unstake_staker ON ac_chain_unstake (staker);
CREATE INDEX IF NOT EXISTS ac_chain_unstake_bridge ON ac_chain_unstake (bridge);
CREATE INDEX IF NOT EXISTS ac_chain_unstake_claim_tx_hash ON ac_chain_unstake (claim_tx_hash);

create table if not exists ac_chain_stake
(
    guid            text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    block_number    UINT256          default 0,
    tx_hash         varchar,
    shares      UINT256          default 0,
    token_address   varchar,
    strategy_address     varchar,
    staker          varchar,
    chain_id varchar,
    created         INTEGER CHECK (created > 0),
    is_points              boolean          default false
    );
CREATE INDEX IF NOT EXISTS ac_chain_stake_tx_hash ON ac_chain_stake (tx_hash);
CREATE INDEX IF NOT EXISTS ac_chain_stake_strategy ON ac_chain_stake (strategy_address);
CREATE INDEX IF NOT EXISTS ac_chain_stake_chain_id ON ac_chain_stake (chain_id);
CREATE INDEX IF NOT EXISTS ac_chain_stake_staker ON ac_chain_stake (staker);

create table if not exists ac_operator_shares_increased(
    guid            text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    block_number      UINT256          default 0,
    tx_hash         varchar,
    shares      UINT256          default 0,
    use_shares      UINT256          default 0,
    strategy_address     varchar,
    operator varchar,
    staker          varchar,
    status int              default 0,
    chain_id varchar,
    created         INTEGER CHECK (created > 0)
    );
CREATE INDEX IF NOT EXISTS ac_operator_shares_increased_tx_hash ON ac_operator_shares_increased (tx_hash);
CREATE INDEX IF NOT EXISTS ac_operator_shares_increased_operator ON ac_operator_shares_increased (operator);
CREATE INDEX IF NOT EXISTS ac_operator_shares_increased_strategy ON ac_operator_shares_increased (strategy_address);
CREATE INDEX IF NOT EXISTS ac_operator_shares_increased_chain_id ON ac_operator_shares_increased (chain_id);
CREATE INDEX IF NOT EXISTS ac_operator_shares_increased_staker ON ac_operator_shares_increased (staker);

create table if not exists ac_chain_increase_batch(
    guid            text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    strategy_address     varchar,
    staker          varchar,
    operator varchar,
    shares      UINT256          default 0,
    chain_id varchar,
    batch_id varchar,
    notify_relayer  boolean          default false,
    created         INTEGER CHECK (created > 0)
);

CREATE INDEX IF NOT EXISTS ac_chain_increase_batch_operator ON ac_chain_increase_batch (operator);
CREATE INDEX IF NOT EXISTS ac_chain_increase_batch_strategy ON ac_chain_increase_batch (strategy_address);
CREATE INDEX IF NOT EXISTS ac_chain_increase_batch_chain_id ON ac_chain_increase_batch (chain_id);
CREATE INDEX IF NOT EXISTS ac_chain_increase_batch_staker ON ac_chain_increase_batch (staker);
CREATE INDEX IF NOT EXISTS ac_chain_increase_batch_batch_id ON ac_chain_increase_batch (batch_id);

create table if not exists ac_chain_dapplink_bridge(
    guid            text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash varchar,
    source_chain_id     varchar,
    dest_chain_id          varchar,
    bridge varchar,
    staking_manager_address      varchar,
    token_address varchar,
    amount UINT256          default 0,
    batch_id  UINT256          default 0,
    nonce         UINT256          default 0,
    notify_relayer  boolean          default false,
    created INTEGER CHECK (created > 0)
);
CREATE INDEX IF NOT EXISTS ac_chain_dapplink_bridge_batch_id ON ac_chain_dapplink_bridge (batch_id);
CREATE INDEX IF NOT EXISTS ac_chain_dapplink_bridge_tx_hash ON ac_chain_dapplink_bridge (tx_hash);

create table if not exists ac_migrate_shares(
    guid            text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash varchar,
    staker     varchar,
    strategy          varchar,
    chain_id varchar,
    shares UINT256 default 0,
    block_number UINT256          default 0,
    unstake_nonce         UINT256          default 0,
    notify_relayer  boolean          default false,
    created INTEGER CHECK (created > 0)
);

CREATE INDEX IF NOT EXISTS ac_migrate_shares_tx_hash ON ac_migrate_shares (tx_hash);


create table if not exists ac_chain_withdraw
(
    guid            text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    block_number    UINT256          default 0,
    tx_hash         varchar,
    shares      UINT256          default 0,
    strategy     varchar,
    operator varchar,
    staker          varchar,
    chain_id varchar,
    created         INTEGER CHECK (created > 0)
    );
CREATE INDEX IF NOT EXISTS ac_chain_withdraw_tx_hash ON ac_chain_withdraw (tx_hash);
CREATE INDEX IF NOT EXISTS ac_chain_withdraw_operator ON ac_chain_withdraw (operator);
CREATE INDEX IF NOT EXISTS ac_chain_withdraw_strategy ON ac_chain_withdraw (strategy);
CREATE INDEX IF NOT EXISTS ac_chain_withdraw_chain_id ON ac_chain_withdraw (chain_id);
CREATE INDEX IF NOT EXISTS ac_chain_withdraw_staker ON ac_chain_withdraw (staker);



create table if not exists ac_chain_deth_shares
(
    guid            text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    shares      UINT256          default 0,
    token_address   varchar,
    strategy_address     varchar,
    staker          varchar,
    chain_id varchar,
    created         INTEGER CHECK (created > 0)
    );
CREATE INDEX IF NOT EXISTS ac_chain_deth_shares_strategy ON ac_chain_deth_shares (strategy_address);
CREATE INDEX IF NOT EXISTS ac_chain_deth_shares_chain_id ON ac_chain_deth_shares (chain_id);
CREATE INDEX IF NOT EXISTS ac_chain_deth_shares_staker ON ac_chain_deth_shares (staker);


create table if not exists ac_chain_deth_transfer
(
    guid            text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    tx_hash      varchar        ,
    block_number    UINT256          default 0,
    from_address     varchar,
    to_address          varchar,
    shares UINT256         default 0,
    claimed int              default 0,
    message_nonce UINT256         default 0,
    notify_relayer boolean default false,
    created         INTEGER CHECK (created > 0)
    );
CREATE INDEX IF NOT EXISTS ac_chain_deth_transfer_tx_hash ON ac_chain_deth_transfer (tx_hash);
CREATE INDEX IF NOT EXISTS ac_chain_deth_transfer_from_address ON ac_chain_deth_transfer (from_address);
CREATE INDEX IF NOT EXISTS ac_chain_deth_transfer_to_address ON ac_chain_deth_transfer (to_address);


