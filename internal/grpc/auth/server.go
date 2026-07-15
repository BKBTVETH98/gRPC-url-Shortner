package auth

import (
	ssov1 "github.com/BKBTVETH98/protos/gen/go/sso"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}
