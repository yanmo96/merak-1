package constants

const (
	TEMPORAL_ADDRESS         = "temporaltest-frontend.default.svc.cluster.local"
	TEMPORAL_PORT            = 7233
	TEMPORAL_ENV             = "TEMPORAL"
	TEMPORAL_RPS_ENV         = "RPS"
	TEMPORAL_CONCURRENCY_ENV = "CONCURRENCY"
	LOCALHOST                = "127.0.0.1"
	TEMPORAL_NAMESPACE       = "merak"

	COMPUTE_GRPC_SERVER_ADDRESS         = "merak-compute-service.merak.svc.cluster.local"
	COMPUTE_GRPC_SERVER_ADDRESS_DEFAULT = "merak-compute-service.default.svc.cluster.local"
	TOPLOGY_GRPC_SERVER_PORT            = 40052
	TOPLOGY_GRPC_SERVER_ADDRESS         = "merak-topology-service.merak.svc.cluster.local"
	NETWORK_GRPC_SERVER_PORT            = 40053
	NETWORK_GRPC_SERVER_ADDRESS         = "merak-network-service.merak.svc.cluster.local"
	AGENT_GRPC_SERVER_PORT              = 40054
	COMPUTE_GRPC_SERVER_PORT            = 40051

	COMPUTE_REDIS_ADDRESS = "compute-redis-main.merak.svc.cluster.local"
	COMPUTE_REDIS_PORT    = 30051

	COMPUTE_REDIS_NODE_IP_SET = "NodeIPSet"
	COMPUTE_REDIS_VM_SET      = "VMSet"

	ALCOR_PORT_MANAGER_PORT        = 30006
	ALCOR_PORT_ID_SUBSTRING_LENGTH = 11

	HTTP_OK             = 200
	HTTP_CREATE_SUCCESS = 201

	MIN_PORT = 1
	MAX_PORT = 65535

	GRPC_MAX_RECV_MSG_SIZE = 1024 * 1024 * 500
	GRPC_MAX_SEND_MSG_SIZE = 1024 * 1024 * 500

	LOG_LOCATION = "/var/log/merak/"

	AGENT_MODE_ENV             = "mode"
	AGENT_MODE_STANDALONE      = "standalone"
	AGENT_STANDALONE_IP        = "10.0.0.2"
	AGENT_STANDALONE_MAC       = "aa:bb:cc:dd:ee:ff"
	AGENT_STANDALONE_REMOTE_ID = "NO ALCOR"
	AGENT_STANDALONE_GW        = "10.0.0.1"
	AGENT_STANDALONE_CIDR      = "10.0.0.0/8"
)
