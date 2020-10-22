package config

import "time"

const (
	SERVER_LISTENING_ADDRESS = ":8888"
	WS_READ_BUF_SIZE         = 2048
	WS_WRITE_BUF_SIZE        = 2048
	WS_READ_LIMIT            = 1024
	WS_SEND_WAIT             = 5 * time.Second
	CHAN_SEND_BUF_SIZE       = 16
	PONG_WAIT                = 70 * time.Second
	PING_PERIOD              = 60 * time.Second
	CONFIG_FILE              = "./configs/lessismore.toml"
	DOTENV_FILE              = "./configs/.lessismore.env"
)
