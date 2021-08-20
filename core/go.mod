module github.com/kimbellG/tournament/core

go 1.16

replace github.com/kimbellG/kerror => ../../tournament-error

require (
	github.com/google/uuid v1.3.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/joho/godotenv v1.3.0
	github.com/kimbellG/kerror v0.0.0-20210819100523-8eb79808c2bd // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d // indirect
	golang.org/x/sys v0.0.0-20210820121016-41cdb8703e55 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210811021853-ddbe55d93216 // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)