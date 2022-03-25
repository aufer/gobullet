module github.com/aufer/gobullet/services/store

go 1.18

replace github.com/aufer/gobullet/model v0.0.0 => ../model/

require github.com/aufer/gobullet/model v0.0.0

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
)
