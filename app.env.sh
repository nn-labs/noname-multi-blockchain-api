cat > app.env << EOF
PORT=:5000
APP_ENV=development
GRPC_HOST=localhost
BTC_RPC_ENDPOINT_TEST=localhost
BTC_RPC_ENDPOINT_MAIN=localhost
BTC_RPC_USER=user
BTC_RPC_PASSWORD=password
EOF