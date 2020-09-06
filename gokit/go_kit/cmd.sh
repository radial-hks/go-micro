# register service
curl \
--request PUT \
--data @node.json \
http://127.0.0.1:8500/v1/agent/service/register

# register
http://127.0.0.1:8500/v1/agent/service/deregister/14