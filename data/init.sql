insert into t_protocol_mapping (internal_protocol_name, internal_protocol_id, protocol_pools, created_at, updated_at)
values ('Curve - Ironbank (USDT)', '1', '[{
  "address": "0x818bBE45B55c1A933f55dC9eb36b2A899586367e",
  "chain_id": "eth",
  "protocol_id": "curve",
  "pool_id": "0x5282a4ef67d9c33135340fb3289cc1711c13638c"
}]', now(), now()),
       ('Euler (+sUSDe,+USD1,-USD1)', '2', '[{
  "address": "0x818bBE45B55c1A933f55dC9eb36b2A899586367e",
  "chain_id": "bsc",
  "protocol_id": "bsc_euler2",
  "pool_id": "0xb2e5a73cee08593d1a076a2ae7a6e02925a640ea"
}]', now(), now())