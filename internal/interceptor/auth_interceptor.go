package interceptor

import (
	"context"
	"encoding/base64"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"homework/pkg/hash"
	"homework/pkg/response"
)

func (i *Interceptor) Auth(ctx context.Context, req interface{}, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		i.logger.Error("failed to get metadata")
		return nil, status.Errorf(codes.Internal, response.ErrInternal.Error())
	}

	authHeaders := md.Get("authorization")
	if len(authHeaders) == 0 {
		i.logger.Error("no authorization header in metadata")
		return nil, status.Errorf(codes.Unauthenticated, response.ErrInternal.Error())
	}

	username, password, ok := parseBasicAuth(authHeaders[0])
	if !ok {

		return nil, status.Errorf(codes.Unauthenticated, "can not parse username and password from metadata")
	}

	isSuccessAuth, err := isAuthDataCorrect(username, password)
	if err != nil {
		i.logger.Errorf("error in getting hash password: %v", err)
		return nil, status.Errorf(codes.Internal, response.ErrInternal.Error())
	}
	if !isSuccessAuth {
		i.logger.Errorf("wrong username/password passed")
		return nil, status.Errorf(codes.Unauthenticated, "username/password is wrong")
	}

	return handler(ctx, req)
}
func parseBasicAuth(authHeader string) (username, password string, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", "", false
	}
	decoded, err := base64.StdEncoding.DecodeString(authHeader[len(prefix):])
	if err != nil {
		return "", "", false
	}
	auth := strings.SplitN(string(decoded), ":", 2)
	if len(auth) != 2 {
		return "", "", false
	}
	return auth[0], auth[1], true
}

func isAuthDataCorrect(username, password string) (bool, error) {
	hashedPassword, err := hash.GetHash(password)
	if err != nil {
		return false, err
	}
	return username == os.Getenv("USERNAME") && hashedPassword == os.Getenv("PASSWORD"), nil
}
