# 获取symbol 列表并按照价格降序排序
GET {{host}}/cm/api/v1/symbols
cm-api-key: {{cm_api_key}}
cm-user-id: hurl
cm-trace-id: 1001

[Query]
sort_field: PRICE
sort_order: DESC
page_size: 10
page_token:

HTTP 200

[Asserts]
jsonpath "$.message" == "success"
jsonpath "$.data.symbols" count == 10

# 获取symbol 列表并按照symbol升序排序
GET {{host}}/cm/api/v1/symbols
cm-api-key: {{cm_api_key}}
cm-user-id: hurl
cm-trace-id: 1002

[Query]
sort_field: SYMBOL
sort_order: ASC
page_size: 20
page_token:

HTTP 200

[Asserts]
jsonpath "$.message" == "success"
jsonpath "$.data.symbols" count == 20

# 获取symbol 列表并按照symbol升序排序，分页
GET {{host}}/cm/api/v1/symbols
cm-api-key: {{cm_api_key}}
cm-user-id: hurl
cm-trace-id: 1003

[Query]
sort_field: SYMBOL
sort_order: ASC
page_size: 10
page_token: BTC

HTTP 200

[Asserts]
jsonpath "$.message" == "success"
jsonpath "$.data.symbols" count == 10

# 获取symbol 列表并按照24小时价格变化率降序排序
GET {{host}}/cm/api/v1/symbols
cm-api-key: {{cm_api_key}}
cm-user-id: hurl
cm-trace-id: 1004

[Query]
sort_field: PRICE_CHANGE_PERCENT_IN_24H
sort_order: ASC
page_size: 10
page_token:

HTTP 200

[Asserts]
jsonpath "$.message" == "success"
jsonpath "$.data.symbols" count == 10


# 通过symbols 获取symbol 列表并按照symbol降序排序
GET {{host}}/cm/api/v1/symbols
cm-api-key: {{cm_api_key}}
cm-user-id: hurl
cm-trace-id: 1005

[Query]
sort_field: SYMBOL
sort_order: ASC
page_size: 10
symbols: BTC
symbols: ETH

HTTP 200

[Asserts]
jsonpath "$.message" == "success"
jsonpath "$.data.symbols" count == 2
jsonpath "$.data.symbols[0].symbol" == "BTC"

# 通过symbol_pairs 获取symbol 列表并按照symbol降序排序
GET {{host}}/cm/api/v1/symbols
cm-api-key: {{cm_api_key}}
cm-user-id: hurl
cm-trace-id: 1006

[Query]
sort_field: SYMBOL
sort_order: ASC
page_size: 10
symbol_pairs: BTCUSDT
symbol_pairs: ETHUSDT

HTTP 200

[Asserts]
jsonpath "$.message" == "success"
jsonpath "$.data.symbols" count == 2
jsonpath "$.data.symbols[0].symbol_pair" == "BTCUSDT"


# 获取支持API交易的symbol 列表并按照symbol降序排序
GET {{host}}/cm/api/v1/symbols
cm-api-key: {{cm_api_key}}
cm-user-id: hurl
cm-trace-id: 1007

[Query]
sort_field: SYMBOL
sort_order: ASC
page_size: 10
is_support_api: true

HTTP 200

[Asserts]
jsonpath "$.message" == "success"
jsonpath "$.data.symbols" count == 10
jsonpath "$.data.symbols[0].is_support_api" == true
jsonpath "$.data.symbols[9].is_support_api" == true