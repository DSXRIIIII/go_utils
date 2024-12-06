package discovery

import (
	"context"
	"github.com/DSXRIIIII/go-utils/go-consul/discovery/consul"
	"github.com/sirupsen/logrus"
	"time"
)

func RegisterToConsul(ctx context.Context, serviceName string) (func() error, error) {
	registry, err := consul.New("116.198.246.11:8500")
	if err != nil {
		return func() error { return nil }, err
	}
	instanceID := GenerateInstanceID(serviceName)
	grpcAddr := "127.0.0.1:5678"
	if err := registry.Register(ctx, instanceID, serviceName, grpcAddr); err != nil {
		return func() error { return nil }, err
	}
	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				logrus.Panicf("no heartbeat from %s to registry, err=%v", serviceName, err)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	logrus.WithFields(logrus.Fields{
		"serviceName": serviceName,
		"addr":        grpcAddr,
	}).Info("registered to consul")
	return func() error {
		return registry.Deregister(ctx, instanceID, serviceName)
	}, nil
}
