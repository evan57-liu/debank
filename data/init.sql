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
}]', now(), now()),
    ('Magpie - Pendle - Hyperwave - Hyperliquid - 25SEP25 (USDC)', '3', '[{
  "address": "0x818bBE45B55c1A933f55dC9eb36b2A899586367e",
  "chain_id": "eth",
  "protocol_id": "magpiexyz",
  "pool_id": "0x16296859c15289731521f199f0a5f762df6347d0:267"
}]', now(), now()),
       ('StakeDAO - Curve  (deUSD,USDC)', '4', '[{
  "address": "0x818bBE45B55c1A933f55dC9eb36b2A899586367e",
  "chain_id": "eth",
  "protocol_id": "stakedao",
  "pool_id": "0xeefb505349e9c306f5b93000138509fbd0681fc8"
}]', now(), now()),
       ('StakeDAO - Curve (crvUSD, USDaf)', '5', '[{
  "address": "0x818bBE45B55c1A933f55dC9eb36b2A899586367e",
  "chain_id": "eth",
  "protocol_id": "stakedao",
  "pool_id": "0xdc147ba5abd134f631a67190deb97b7828b4afb7"
}]', now(), now()),
       ('Fluid', '6', '[{
  "address": "0x818bBE45B55c1A933f55dC9eb36b2A899586367e",
  "chain_id": "eth",
  "protocol_id": "fluid",
  "pool_id": "0x324c5dc1fc42c7a4d43d92df1eba58a54d13bf2d"
}]', now(), now()),
       ('StakeDAO - Curve (USDC, USDf)', '7', '[{
  "address": "0x818bBE45B55c1A933f55dC9eb36b2A899586367e",
  "chain_id": "eth",
  "protocol_id": "stakedao",
  "pool_id": "0x5f1f2b52221b6091294865bdace2fb139aa031ad"
}]', now(), now())

