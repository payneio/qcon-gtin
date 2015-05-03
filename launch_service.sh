#!/bin/bash

fleetctl start gtin@{1..3}
fleetctl start gtin.proxy.private@{1..3}
fleetctl start gtin.proxy.private.frontend
fleetctl start gtin.proxy.public@{1..3}
fleetctl start gtin.proxy.public.frontend

echo "Health check: http://api-blue.qcon.demo/stores/health/gtins/check"
echo "Health check: http://api.qcon.demo/stores/health/gtins/check"
