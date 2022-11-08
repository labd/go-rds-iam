package rdsiam

import (
	"context"
	"fmt"
	"os"
	"os/user"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
)

type DSN struct {
	host     string
	port     int
	user     string
	database string
}

func createNewDSN(ctx context.Context, dsnString string) (string, error) {
	dsn, err := parseDSN(dsnString)
	if err != nil {
		return "", err
	}

	token, err := GetToken(ctx, dsn)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"user=%s dbname=%s sslmode=require host=%s port=%d password=%s",
		dsn.user,
		dsn.database,
		dsn.host,
		dsn.port,
		token)
	return result, nil
}

func GetToken(ctx context.Context, dsn *DSN) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	endpoint := fmt.Sprintf("%s:%d", dsn.host, dsn.port)
	region := os.Getenv("AWS_REGION")
	return auth.BuildAuthToken(
		ctx, endpoint, region, dsn.user, cfg.Credentials)
}

func parseDSN(dsn string) (*DSN, error) {
	options := make(map[string]string)
	if err := parseOpts(dsn, options); err != nil {
		return nil, err
	}

	result := &DSN{
		host: "localhost",
		port: 5432,
	}

	if val, ok := options["host"]; ok {
		result.host = val
	}

	if val, ok := options["port"]; ok {
		port, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		result.port = port
	}

	if val, ok := options["user"]; ok {
		result.user = val
	} else {
		if userInfo, err := user.Current(); err == nil {
			result.user = userInfo.Username
		}
	}
	if val, ok := options["database"]; ok {
		result.database = val
	} else {
		result.database = result.user
	}
	return result, nil
}
