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

type NatsWrapper struct {
	*nats.EncodedConn
	Subs      []*nats.Subscription
	Env       string
	Namespace string
}

func New(n *nats.EncodedConn, env, namespace string) *NatsWrapper {
	return &NatsWrapper{
		EncodedConn: n,
		Subs:        []*nats.Subscription{},
		Env:         env,
		Namespace:   namespace,
	}
}

type MsgHandlerFunc = func() nats.MsgHandler

type HandlerFunc = func() nats.Handler

// EncodedBindRecvChan binds a channel for receive operations from NATS.
func (n *NatsWrapper) EncodedBindRecvChan(subject string, channel interface{}) (*nats.Subscription, error) {
	return n.EncodedConn.BindRecvChan(n.dotJoin(subject), channel)
}

// EncodedBindRecvQueueChan binds a channel for queue-based receive operations from NATS.
func (n *NatsWrapper) EncodedBindRecvQueueChan(subject string, queue string, channel interface{}) (*nats.Subscription, error) {
	return n.EncodedConn.BindRecvQueueChan(n.dotJoin(subject), queue, channel)
}

// EncodedBindSendChan binds a channel for send operations to NATS.
func (n *NatsWrapper) EncodedBindSendChan(subject string, channel interface{}) error {
	return n.EncodedConn.BindSendChan(n.dotJoin(subject), channel)
}

// Publish publishes the data argument to the given subject.
func (n *NatsWrapper) Publish(subject string, data []byte) error {
	return n.EncodedConn.Conn.Publish(n.dotJoin(subject), data)
}

// PublishRequest will perform a Publish() excpecting a response on the reply subject.
func (n *NatsWrapper) PublishRequest(subject string, reply string, data []byte) error {
	return n.EncodedConn.Conn.PublishRequest(n.dotJoin(subject), reply, data)
}

// EncodedPublish publishes the data argument to the given subject.
func (n *NatsWrapper) EncodedPublish(subject string, v any) error {
	return n.EncodedConn.Publish(n.dotJoin(subject), v)
}

// EncodedPublishRequest will perform a Publish() expecting a response on the reply subject.
func (n *NatsWrapper) EncodedPublishRequest(subject string, reply string, payload []byte) error {
	return n.EncodedConn.PublishRequest(n.dotJoin(subject), reply, payload)
}

// ChanSubscribe will express interest in the given subject and place all messages received on the channel.
func (n *NatsWrapper) ChanSubscribe(subject string, ch chan *nats.Msg) (*nats.Subscription, error) {
	sub, err := n.EncodedConn.Conn.ChanSubscribe(n.dotJoin(subject), ch)
	if err != nil {
		return nil, err
	}

	n.Subs = append(n.Subs, sub)

	return sub, err
}

// ChanQueueSubscribe will express interest in the given subject.
func (n *NatsWrapper) ChanQueueSubscribe(subject string, group string, ch chan *nats.Msg) (*nats.Subscription, error) {
	sub, err := n.EncodedConn.Conn.ChanQueueSubscribe(n.dotJoin(subject), group, ch)
	if err != nil {
		return nil, err
	}

	n.Subs = append(n.Subs, sub)
	return sub, err
}

// QueueSubscribe creates an asynchronous queue subscriber on the given subject.
func (n *NatsWrapper) QueueSubscribe(subject string, queue string, handler MsgHandlerFunc) (*nats.Subscription, error) {
	sub, err := n.EncodedConn.Conn.QueueSubscribe(n.dotJoin(subject), queue, handler())
	if err != nil {
		return nil, err
	}

	n.Subs = append(n.Subs, sub)

	return sub, err
}

// QueueSubscribeSync creates a synchronous queue subscriber on the given subject.
func (n *NatsWrapper) QueueSubscribeSync(subject string, queue string) (*nats.Subscription, error) {
	sub, err := n.EncodedConn.Conn.QueueSubscribeSync(n.dotJoin(subject), queue)
	if err != nil {
		return nil, err
	}

	return sub, err
}

// QueueSubscribeSyncWithChan will express interest in the given subject.
func (n *NatsWrapper) QueueSubscribeSyncWithChan(subject string, queue string, ch chan *nats.Msg) (*nats.Subscription, error) {
	sub, err := n.EncodedConn.Conn.QueueSubscribeSyncWithChan(n.dotJoin(subject), queue, ch)
	if err != nil {
		return nil, err
	}

	n.Subs = append(n.Subs, sub)

	return sub, err
}

// EncodedQueueSubscribe will create a queue subscription on the given subject and process incoming messages using the specified Handler.
func (n *NatsWrapper) EncodedQueueSubscribe(subject string, queue string, handler MsgHandlerFunc) (*nats.Subscription, error) {
	return n.EncodedConn.QueueSubscribe(n.dotJoin(subject), queue, handler)
}

// Request subscribes to a subject and expects one response
func (n *NatsWrapper) Request(subject string, payload []byte) (*nats.Msg, error) {
	return n.EncodedConn.Conn.Request(n.dotJoin(subject), payload, StandardTimeout)
}

// RequestWithContext takes a context, a subject and payload in bytes and request expecting a single response.
func (n *NatsWrapper) RequestWithContext(ctx context.Context, subject string, data []byte) (*nats.Msg, error) {
	return n.EncodedConn.Conn.RequestWithContext(ctx, n.dotJoin(subject), data)
}

// EncodedRequest subscribes to a subject and expects one response which modifies the out interface
func (n *NatsWrapper) EncodedRequest(subject string, in, out any) error {
	return n.EncodedConn.Request(n.dotJoin(subject), in, out, StandardTimeout)
}

// EncodedRequestWithContext will create an Inbox and perform a Request using the provided cancellation context with the Inbox reply for the data v.
func (n *NatsWrapper) EncodedRequestWithContext(ctx context.Context, subject string, in, out any) error {
	return n.EncodedConn.RequestWithContext(ctx, n.dotJoin(subject), in, out)
}

// Subscribe will express interest in the given subject.
func (n *NatsWrapper) Subscribe(subject string, handler MsgHandlerFunc) (*nats.Subscription, error) {
	return n.EncodedConn.Conn.Subscribe(n.dotJoin(subject), handler())
}

// SubscribeSync will express interest on the given subject.
func (n *NatsWrapper) SubscribeSync(subject string) (*nats.Subscription, error) {
	return n.EncodedConn.Conn.SubscribeSync(n.dotJoin(subject))
}

func (n *NatsWrapper) EncodedSubscribe(subject string, handler MsgHandlerFunc) (*nats.Subscription, error) {
	return n.EncodedConn.Subscribe(n.dotJoin(subject), handler)
}

func (n *NatsWrapper) dotJoin(values ...string) string {
	vals := []string{n.Env}
	vals = append(vals, values...)
	return strings.Join(vals, ".")
}
