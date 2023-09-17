package pubInner

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/i-Things/things/shared/clients"
	"github.com/i-Things/things/shared/conf"
	"github.com/i-Things/things/shared/devices"
	"github.com/i-Things/things/shared/events"
	"github.com/i-Things/things/shared/events/topics"
	"github.com/i-Things/things/shared/utils"
	"github.com/nats-io/nats.go"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	NatsClient struct {
		client *nats.Conn
	}
)

func newNatsClient(conf conf.NatsConf) (PubInner, error) {
	nc, err := clients.NewNatsClient(conf)
	if err != nil {
		return nil, err
	}
	return &NatsClient{client: nc}, nil
}

func (n *NatsClient) DevPubGateway(ctx context.Context, publishMsg *devices.DevPublish) error {
	pubStr, _ := json.Marshal(publishMsg)
	return n.publish(ctx,
		fmt.Sprintf(topics.DeviceUpMsg, publishMsg.Handle, publishMsg.ProductID, publishMsg.DeviceName), pubStr)
}

func (n *NatsClient) DevPubMsg(ctx context.Context, publishMsg *devices.DevPublish) error {
	pubStr, _ := json.Marshal(publishMsg)
	err := n.publish(ctx,
		fmt.Sprintf(topics.DeviceUpMsg, publishMsg.Handle, publishMsg.ProductID, publishMsg.DeviceName), pubStr)
	if err != nil {
		logx.Errorf("%s.publish  err:%v", utils.FuncName(), err)
		return err
	}
	return err
}

func (n *NatsClient) DevPubThing(ctx context.Context, publishMsg *devices.DevPublish) error {
	pubStr, _ := json.Marshal(publishMsg)
	err := n.publish(ctx,
		fmt.Sprintf(topics.DeviceUpThing, publishMsg.ProductID, publishMsg.DeviceName), pubStr)
	if err != nil {
		logx.Errorf("%s.publish  err:%v", utils.FuncName(), err)
		return err
	}
	return err
}

func (n *NatsClient) DevPubOta(ctx context.Context, publishMsg *devices.DevPublish) error {
	pubStr, _ := json.Marshal(publishMsg)
	err := n.publish(ctx,
		fmt.Sprintf(topics.DeviceUpOta, publishMsg.ProductID, publishMsg.DeviceName), pubStr)
	if err != nil {
		logx.Errorf("%s.publish  err:%v", utils.FuncName(), err)
		return err
	}
	return err
}

func (n *NatsClient) DevPubConfig(ctx context.Context, publishMsg *devices.DevPublish) error {
	pubStr, _ := json.Marshal(publishMsg)
	err := n.publish(ctx,
		fmt.Sprintf(topics.DeviceUpConfig, publishMsg.ProductID, publishMsg.DeviceName), pubStr)
	if err != nil {
		logx.Errorf("%s.publish  err:%v", utils.FuncName(), err)
		return err
	}
	return err
}

func (n *NatsClient) DevPubShadow(ctx context.Context, publishMsg *devices.DevPublish) error {
	pubStr, _ := json.Marshal(publishMsg)
	err := n.publish(ctx,
		fmt.Sprintf(topics.DeviceUpShadow, publishMsg.ProductID, publishMsg.DeviceName), pubStr)
	if err != nil {
		logx.Errorf("%s.publish  err:%v", utils.FuncName(), err)
		return err
	}
	return err
}

func (n *NatsClient) DevPubSDKLog(ctx context.Context, publishMsg *devices.DevPublish) error {
	pubStr, _ := json.Marshal(publishMsg)
	err := n.publish(ctx,
		fmt.Sprintf(topics.DeviceUpSDKLog, publishMsg.ProductID, publishMsg.DeviceName), pubStr)
	if err != nil {
		logx.Errorf("%s.publish  err:%v", utils.FuncName(), err)
		return err
	}
	return err
}

func (n *NatsClient) PubConn(ctx context.Context, conn ConnType, info *devices.DevConn) error {
	str, _ := json.Marshal(info)
	var err error
	switch conn {
	case Connect:
		err = n.publish(ctx, topics.DeviceUpStatusConnected, str)
	case DisConnect:
		err = n.publish(ctx, topics.DeviceUpStatusDisconnected, str)
	default:
		panic("not support conn type")
	}
	if err != nil {
		logx.Errorf("%s.publish  err:%v", utils.FuncName(), err)
		return err
	}
	return err
}

func (n *NatsClient) publish(ctx context.Context, topic string, payload []byte) error {
	err := n.client.Publish(topic, events.NewEventMsg(ctx, payload))
	if err != nil {
		logx.WithContext(ctx).Errorf("%s nats publish failure err:%v topic:%v", err, topic)
	}
	return err
}
