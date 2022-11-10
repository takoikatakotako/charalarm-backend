package config

const (
	AWSRegion          = "ap-northeast-1"
	LocalstackEndpoint = "http://localhost:4566"

	LocalVoIPPushQueueURL           = "http://localhost:4566/000000000000/voip-push-queue.fifo"
	LocalVoIPPushDeadLetterQueueURL = "http://localhost:4566/000000000000/voip-push-dead-letter-queue.fifo"

	VoIPPushQueueURLKey           = "VOIP_PUSH_QUEUE_URL"
	VoIPPushDeadLetterQueueURLKey = "VOIP_PUSH_DEAD_LETTER_QUEUE_URL"
)
