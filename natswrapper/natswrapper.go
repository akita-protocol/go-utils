package natswrapper

import (
	"context"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	StandardTimeout = (time.Second * 30)
)

type Events string

const (
	REMOVED_ADDRESS                Events = "REMOVED_DISCORD_ADDRESS"
	JOINED_DISCORD_SERVER          Events = "JOINED_DISCORD_SERVER"
	LEFT_DISCORD_SERVER            Events = "LEFT_DISCORD_SERVER"
	NEW_TWEET                      Events = "NEW_TWEET"
	NEW_APP                        Events = "NEW_APP"
	UPDATE_APP                     Events = "UPDATE_APP"
	NEW_ASSET                      Events = "NEW_ASSET"
	UPDATE_ASSET                   Events = "UPDATE_ASSET"
	NEW_TRANSACTION                Events = "NEW_TRANSACTION"
	NEW_ADDRESS                    Events = "NEW_ADDRESS"
	NEW_ADDRESS_HOLDINGS_PROCESSED Events = "NEW_ADDRESS_HOLDINGS_PROCESSED"
	NEW_COMMUNITY                  Events = "NEW_COMMUNITY"
	NEW_USER                       Events = "NEW_USER"
)

var NamespaceActions map[Namespace][]Events = map[Namespace][]Events{
	NamespaceDiscordBot: {
		NEW_USER,
		NEW_ADDRESS,
		REMOVED_ADDRESS,
		JOINED_DISCORD_SERVER,
		LEFT_DISCORD_SERVER,
	},
	NamespaceFoundationWatch: {
		NEW_TWEET,
	},
	NamespaceSubscriptions: {},
	NamespaceNFD:           {},
	NamespaceBlockWatcher: {
		NEW_APP,
		UPDATE_APP,
		NEW_ASSET,
		UPDATE_ASSET,
		NEW_TRANSACTION,
		NEW_ADDRESS,
		NEW_ADDRESS_HOLDINGS_PROCESSED,
	},
	NamespaceSeed: {},
	NamespaceGateway: {
		NEW_USER,
		NEW_ADDRESS,
	},
}

type Namespace string

const (
	NamespaceDiscordBot      Namespace = "discordbot"
	NamespaceFoundationWatch Namespace = "foundationwatch"
	NamespaceSubscriptions   Namespace = "subscriptions"
	NamespaceNFD             Namespace = "nfd"
	NamespaceBlockWatcher    Namespace = "blockwatcher"
	NamespaceSeed            Namespace = "seed"
	NamespaceGateway         Namespace = "gateway"
)

type NatsWrapper struct {
	*nats.EncodedConn
	Subs      []*nats.Subscription
	Env       string
	Namespace string
}

func New(n *nats.EncodedConn, env string, namespace Namespace) *NatsWrapper {
	return &NatsWrapper{
		EncodedConn: n,
		Subs:        []*nats.Subscription{},
		Env:         env,
		Namespace:   string(namespace),
	}
}

type MsgHandlerFunc = func() nats.MsgHandler

type HandlerFunc = func() nats.Handler

// EncodedBindRecvChan binds a channel for receive operations from NATS.
func (n *NatsWrapper) EncodedBindRecvChan(subject string, channel any) (*nats.Subscription, error) {
	return n.EncodedConn.BindRecvChan(subject, channel)
}

// EncodedBindRecvQueueChan binds a channel for queue-based receive operations from NATS.
func (n *NatsWrapper) EncodedBindRecvQueueChan(subject string, queue string, channel any) (*nats.Subscription, error) {
	return n.EncodedConn.BindRecvQueueChan(subject, queue, channel)
}

// EncodedBindSendChan binds a channel for send operations to NATS.
func (n *NatsWrapper) EncodedBindSendChan(subject string, channel any) error {
	return n.EncodedConn.BindSendChan(subject, channel)
}

// Publish publishes the data argument to the given subject.
func (n *NatsWrapper) Publish(subject string, data []byte) error {
	return n.EncodedConn.Conn.Publish(subject, data)
}

// PublishRequest will perform a Publish() excpecting a response on the reply subject.
func (n *NatsWrapper) PublishRequest(subject string, reply string, data []byte) error {
	return n.EncodedConn.Conn.PublishRequest(subject, reply, data)
}

// EncodedPublish publishes the data argument to the given subject.
func (n *NatsWrapper) EncodedPublish(subject string, v any) error {
	return n.EncodedConn.Publish(subject, v)
}

// EncodedPublishRequest will perform a Publish() expecting a response on the reply subject.
func (n *NatsWrapper) EncodedPublishRequest(subject string, reply string, payload any) error {
	return n.EncodedConn.PublishRequest(subject, reply, payload)
}

// ChanSubscribe will express interest in the given subject and place all messages received on the channel.
func (n *NatsWrapper) ChanSubscribe(subject string, ch chan *nats.Msg) (*nats.Subscription, error) {
	sub, err := n.EncodedConn.Conn.ChanSubscribe(subject, ch)
	if err != nil {
		return nil, err
	}

	n.Subs = append(n.Subs, sub)

	return sub, err
}

// ChanQueueSubscribe will express interest in the given subject.
func (n *NatsWrapper) ChanQueueSubscribe(subject string, group string, ch chan *nats.Msg) (*nats.Subscription, error) {
	sub, err := n.EncodedConn.Conn.ChanQueueSubscribe(subject, group, ch)
	if err != nil {
		return nil, err
	}

	n.Subs = append(n.Subs, sub)
	return sub, err
}

// QueueSubscribe creates an asynchronous queue subscriber on the given subject.
func (n *NatsWrapper) QueueSubscribe(subject string, queue string, handler MsgHandlerFunc) (*nats.Subscription, error) {
	sub, err := n.EncodedConn.Conn.QueueSubscribe(subject, queue, handler())
	if err != nil {
		return nil, err
	}

	n.Subs = append(n.Subs, sub)

	return sub, err
}

// QueueSubscribeSync creates a synchronous queue subscriber on the given subject.
func (n *NatsWrapper) QueueSubscribeSync(subject string, queue string) (*nats.Subscription, error) {
	sub, err := n.EncodedConn.Conn.QueueSubscribeSync(subject, queue)
	if err != nil {
		return nil, err
	}

	return sub, err
}

// QueueSubscribeSyncWithChan will express interest in the given subject.
func (n *NatsWrapper) QueueSubscribeSyncWithChan(subject string, queue string, ch chan *nats.Msg) (*nats.Subscription, error) {
	sub, err := n.EncodedConn.Conn.QueueSubscribeSyncWithChan(subject, queue, ch)
	if err != nil {
		return nil, err
	}

	n.Subs = append(n.Subs, sub)

	return sub, err
}

// EncodedQueueSubscribe will create a queue subscription on the given subject and process incoming messages using the specified Handler.
func (n *NatsWrapper) EncodedQueueSubscribe(subject string, queue string, handler HandlerFunc) (*nats.Subscription, error) {
	sub, err := n.EncodedConn.QueueSubscribe(subject, queue, handler())
	if err != nil {
		return nil, err
	}

	n.Subs = append(n.Subs, sub)

	return sub, err
}

// Request subscribes to a subject and expects one response
func (n *NatsWrapper) Request(subject string, payload []byte) (*nats.Msg, error) {
	return n.EncodedConn.Conn.Request(subject, payload, StandardTimeout)
}

// RequestWithContext takes a context, a subject and payload in bytes and request expecting a single response.
func (n *NatsWrapper) RequestWithContext(ctx context.Context, subject string, data []byte) (*nats.Msg, error) {
	return n.EncodedConn.Conn.RequestWithContext(ctx, subject, data)
}

// EncodedRequest subscribes to a subject and expects one response which modifies the out interface
func (n *NatsWrapper) EncodedRequest(subject string, in, out any) error {
	return n.EncodedConn.Request(subject, in, out, StandardTimeout)
}

// EncodedRequestWithContext will create an Inbox and perform a Request using the provided cancellation context with the Inbox reply for the data v.
func (n *NatsWrapper) EncodedRequestWithContext(ctx context.Context, subject string, in, out any) error {
	return n.EncodedConn.RequestWithContext(ctx, subject, in, out)
}

// Subscribe will express interest in the given subject.
func (n *NatsWrapper) Subscribe(subject string, handler MsgHandlerFunc) (*nats.Subscription, error) {
	return n.EncodedConn.Conn.Subscribe(subject, handler())
}

// SubscribeSync will express interest on the given subject.
func (n *NatsWrapper) SubscribeSync(subject string) (*nats.Subscription, error) {
	return n.EncodedConn.Conn.SubscribeSync(subject)
}

func (n *NatsWrapper) EncodedSubscribe(subject string, handler HandlerFunc) (*nats.Subscription, error) {
	return n.EncodedConn.Subscribe(subject, handler())
}

func (n *NatsWrapper) BuildSubject(values ...string) string {
	return strings.Join(values, ".")
}
