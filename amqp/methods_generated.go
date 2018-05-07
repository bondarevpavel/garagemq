package amqp

import (
	"errors"
	"io"
)

type Method interface {
	Name() string
	FrameType() byte
	ClassIdentifier() uint16
	MethodIdentifier() uint16
	Read(reader io.Reader) (err error)
	Write(writer io.Writer) (err error)
}

// Connection methods

type ConnectionStart struct {
	VersionMajor byte

	VersionMinor byte

	ServerProperties *Table

	Mechanisms []byte

	Locales []byte
}

func (method *ConnectionStart) Name() string {
	return "ConnectionStart"
}

func (method *ConnectionStart) FrameType() byte {
	return 1
}

func (method *ConnectionStart) ClassIdentifier() uint16 {
	return 10
}

func (method *ConnectionStart) MethodIdentifier() uint16 {
	return 10
}

func (method *ConnectionStart) Read(reader io.Reader) (err error) {

	method.VersionMajor, err = ReadOctet(reader)
	if err != nil {
		return err
	}

	method.VersionMinor, err = ReadOctet(reader)
	if err != nil {
		return err
	}

	method.ServerProperties, err = ReadTable(reader)
	if err != nil {
		return err
	}

	method.Mechanisms, err = ReadLongstr(reader)
	if err != nil {
		return err
	}

	method.Locales, err = ReadLongstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ConnectionStart) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ConnectionStartOk struct {
	ClientProperties *Table

	Mechanism string

	Response []byte

	Locale string
}

func (method *ConnectionStartOk) Name() string {
	return "ConnectionStartOk"
}

func (method *ConnectionStartOk) FrameType() byte {
	return 1
}

func (method *ConnectionStartOk) ClassIdentifier() uint16 {
	return 10
}

func (method *ConnectionStartOk) MethodIdentifier() uint16 {
	return 11
}

func (method *ConnectionStartOk) Read(reader io.Reader) (err error) {

	method.ClientProperties, err = ReadTable(reader)
	if err != nil {
		return err
	}

	method.Mechanism, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.Response, err = ReadLongstr(reader)
	if err != nil {
		return err
	}

	method.Locale, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ConnectionStartOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ConnectionSecure struct {
	Challenge []byte
}

func (method *ConnectionSecure) Name() string {
	return "ConnectionSecure"
}

func (method *ConnectionSecure) FrameType() byte {
	return 1
}

func (method *ConnectionSecure) ClassIdentifier() uint16 {
	return 10
}

func (method *ConnectionSecure) MethodIdentifier() uint16 {
	return 20
}

func (method *ConnectionSecure) Read(reader io.Reader) (err error) {

	method.Challenge, err = ReadLongstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ConnectionSecure) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ConnectionSecureOk struct {
	Response []byte
}

func (method *ConnectionSecureOk) Name() string {
	return "ConnectionSecureOk"
}

func (method *ConnectionSecureOk) FrameType() byte {
	return 1
}

func (method *ConnectionSecureOk) ClassIdentifier() uint16 {
	return 10
}

func (method *ConnectionSecureOk) MethodIdentifier() uint16 {
	return 21
}

func (method *ConnectionSecureOk) Read(reader io.Reader) (err error) {

	method.Response, err = ReadLongstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ConnectionSecureOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ConnectionTune struct {
	ChannelMax uint16

	FrameMax uint32

	Heartbeat uint16
}

func (method *ConnectionTune) Name() string {
	return "ConnectionTune"
}

func (method *ConnectionTune) FrameType() byte {
	return 1
}

func (method *ConnectionTune) ClassIdentifier() uint16 {
	return 10
}

func (method *ConnectionTune) MethodIdentifier() uint16 {
	return 30
}

func (method *ConnectionTune) Read(reader io.Reader) (err error) {

	method.ChannelMax, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.FrameMax, err = ReadLong(reader)
	if err != nil {
		return err
	}

	method.Heartbeat, err = ReadShort(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ConnectionTune) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ConnectionTuneOk struct {
	ChannelMax uint16

	FrameMax uint32

	Heartbeat uint16
}

func (method *ConnectionTuneOk) Name() string {
	return "ConnectionTuneOk"
}

func (method *ConnectionTuneOk) FrameType() byte {
	return 1
}

func (method *ConnectionTuneOk) ClassIdentifier() uint16 {
	return 10
}

func (method *ConnectionTuneOk) MethodIdentifier() uint16 {
	return 31
}

func (method *ConnectionTuneOk) Read(reader io.Reader) (err error) {

	method.ChannelMax, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.FrameMax, err = ReadLong(reader)
	if err != nil {
		return err
	}

	method.Heartbeat, err = ReadShort(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ConnectionTuneOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ConnectionOpen struct {
	VirtualHost string

	Reserved1 string

	Reserved2 bool
}

func (method *ConnectionOpen) Name() string {
	return "ConnectionOpen"
}

func (method *ConnectionOpen) FrameType() byte {
	return 1
}

func (method *ConnectionOpen) ClassIdentifier() uint16 {
	return 10
}

func (method *ConnectionOpen) MethodIdentifier() uint16 {
	return 40
}

func (method *ConnectionOpen) Read(reader io.Reader) (err error) {

	method.VirtualHost, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.Reserved1, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.Reserved2 = bits&(1<<0) != 0

	return
}

func (method *ConnectionOpen) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ConnectionOpenOk struct {
	Reserved1 string
}

func (method *ConnectionOpenOk) Name() string {
	return "ConnectionOpenOk"
}

func (method *ConnectionOpenOk) FrameType() byte {
	return 1
}

func (method *ConnectionOpenOk) ClassIdentifier() uint16 {
	return 10
}

func (method *ConnectionOpenOk) MethodIdentifier() uint16 {
	return 41
}

func (method *ConnectionOpenOk) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ConnectionOpenOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ConnectionClose struct {
	ReplyCode uint16

	ReplyText string

	ClassId uint16

	MethodId uint16
}

func (method *ConnectionClose) Name() string {
	return "ConnectionClose"
}

func (method *ConnectionClose) FrameType() byte {
	return 1
}

func (method *ConnectionClose) ClassIdentifier() uint16 {
	return 10
}

func (method *ConnectionClose) MethodIdentifier() uint16 {
	return 50
}

func (method *ConnectionClose) Read(reader io.Reader) (err error) {

	method.ReplyCode, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.ReplyText, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.ClassId, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.MethodId, err = ReadShort(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ConnectionClose) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ConnectionCloseOk struct {
}

func (method *ConnectionCloseOk) Name() string {
	return "ConnectionCloseOk"
}

func (method *ConnectionCloseOk) FrameType() byte {
	return 1
}

func (method *ConnectionCloseOk) ClassIdentifier() uint16 {
	return 10
}

func (method *ConnectionCloseOk) MethodIdentifier() uint16 {
	return 51
}

func (method *ConnectionCloseOk) Read(reader io.Reader) (err error) {

	return
}

func (method *ConnectionCloseOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

// Channel methods

type ChannelOpen struct {
	Reserved1 string
}

func (method *ChannelOpen) Name() string {
	return "ChannelOpen"
}

func (method *ChannelOpen) FrameType() byte {
	return 1
}

func (method *ChannelOpen) ClassIdentifier() uint16 {
	return 20
}

func (method *ChannelOpen) MethodIdentifier() uint16 {
	return 10
}

func (method *ChannelOpen) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ChannelOpen) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ChannelOpenOk struct {
	Reserved1 []byte
}

func (method *ChannelOpenOk) Name() string {
	return "ChannelOpenOk"
}

func (method *ChannelOpenOk) FrameType() byte {
	return 1
}

func (method *ChannelOpenOk) ClassIdentifier() uint16 {
	return 20
}

func (method *ChannelOpenOk) MethodIdentifier() uint16 {
	return 11
}

func (method *ChannelOpenOk) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadLongstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ChannelOpenOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ChannelFlow struct {
	Active bool
}

func (method *ChannelFlow) Name() string {
	return "ChannelFlow"
}

func (method *ChannelFlow) FrameType() byte {
	return 1
}

func (method *ChannelFlow) ClassIdentifier() uint16 {
	return 20
}

func (method *ChannelFlow) MethodIdentifier() uint16 {
	return 20
}

func (method *ChannelFlow) Read(reader io.Reader) (err error) {

	bits, err := ReadOctet(reader)

	method.Active = bits&(1<<0) != 0

	return
}

func (method *ChannelFlow) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ChannelFlowOk struct {
	Active bool
}

func (method *ChannelFlowOk) Name() string {
	return "ChannelFlowOk"
}

func (method *ChannelFlowOk) FrameType() byte {
	return 1
}

func (method *ChannelFlowOk) ClassIdentifier() uint16 {
	return 20
}

func (method *ChannelFlowOk) MethodIdentifier() uint16 {
	return 21
}

func (method *ChannelFlowOk) Read(reader io.Reader) (err error) {

	bits, err := ReadOctet(reader)

	method.Active = bits&(1<<0) != 0

	return
}

func (method *ChannelFlowOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ChannelClose struct {
	ReplyCode uint16

	ReplyText string

	ClassId uint16

	MethodId uint16
}

func (method *ChannelClose) Name() string {
	return "ChannelClose"
}

func (method *ChannelClose) FrameType() byte {
	return 1
}

func (method *ChannelClose) ClassIdentifier() uint16 {
	return 20
}

func (method *ChannelClose) MethodIdentifier() uint16 {
	return 40
}

func (method *ChannelClose) Read(reader io.Reader) (err error) {

	method.ReplyCode, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.ReplyText, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.ClassId, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.MethodId, err = ReadShort(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ChannelClose) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ChannelCloseOk struct {
}

func (method *ChannelCloseOk) Name() string {
	return "ChannelCloseOk"
}

func (method *ChannelCloseOk) FrameType() byte {
	return 1
}

func (method *ChannelCloseOk) ClassIdentifier() uint16 {
	return 20
}

func (method *ChannelCloseOk) MethodIdentifier() uint16 {
	return 41
}

func (method *ChannelCloseOk) Read(reader io.Reader) (err error) {

	return
}

func (method *ChannelCloseOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

// Exchange methods

type ExchangeDeclare struct {
	Reserved1 uint16

	Exchange string

	Type string

	Passive bool

	Durable bool

	Reserved2 bool

	Reserved3 bool

	NoWait bool

	Arguments *Table
}

func (method *ExchangeDeclare) Name() string {
	return "ExchangeDeclare"
}

func (method *ExchangeDeclare) FrameType() byte {
	return 1
}

func (method *ExchangeDeclare) ClassIdentifier() uint16 {
	return 40
}

func (method *ExchangeDeclare) MethodIdentifier() uint16 {
	return 10
}

func (method *ExchangeDeclare) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.Exchange, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.Type, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.Passive = bits&(1<<0) != 0

	method.Durable = bits&(1<<1) != 0

	method.Reserved2 = bits&(1<<2) != 0

	method.Reserved3 = bits&(1<<3) != 0

	method.NoWait = bits&(1<<4) != 0

	method.Arguments, err = ReadTable(reader)
	if err != nil {
		return err
	}

	return
}

func (method *ExchangeDeclare) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ExchangeDeclareOk struct {
}

func (method *ExchangeDeclareOk) Name() string {
	return "ExchangeDeclareOk"
}

func (method *ExchangeDeclareOk) FrameType() byte {
	return 1
}

func (method *ExchangeDeclareOk) ClassIdentifier() uint16 {
	return 40
}

func (method *ExchangeDeclareOk) MethodIdentifier() uint16 {
	return 11
}

func (method *ExchangeDeclareOk) Read(reader io.Reader) (err error) {

	return
}

func (method *ExchangeDeclareOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ExchangeDelete struct {
	Reserved1 uint16

	Exchange string

	IfUnused bool

	NoWait bool
}

func (method *ExchangeDelete) Name() string {
	return "ExchangeDelete"
}

func (method *ExchangeDelete) FrameType() byte {
	return 1
}

func (method *ExchangeDelete) ClassIdentifier() uint16 {
	return 40
}

func (method *ExchangeDelete) MethodIdentifier() uint16 {
	return 20
}

func (method *ExchangeDelete) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.Exchange, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.IfUnused = bits&(1<<0) != 0

	method.NoWait = bits&(1<<1) != 0

	return
}

func (method *ExchangeDelete) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type ExchangeDeleteOk struct {
}

func (method *ExchangeDeleteOk) Name() string {
	return "ExchangeDeleteOk"
}

func (method *ExchangeDeleteOk) FrameType() byte {
	return 1
}

func (method *ExchangeDeleteOk) ClassIdentifier() uint16 {
	return 40
}

func (method *ExchangeDeleteOk) MethodIdentifier() uint16 {
	return 21
}

func (method *ExchangeDeleteOk) Read(reader io.Reader) (err error) {

	return
}

func (method *ExchangeDeleteOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

// Queue methods

type QueueDeclare struct {
	Reserved1 uint16

	Queue string

	Passive bool

	Durable bool

	Exclusive bool

	AutoDelete bool

	NoWait bool

	Arguments *Table
}

func (method *QueueDeclare) Name() string {
	return "QueueDeclare"
}

func (method *QueueDeclare) FrameType() byte {
	return 1
}

func (method *QueueDeclare) ClassIdentifier() uint16 {
	return 50
}

func (method *QueueDeclare) MethodIdentifier() uint16 {
	return 10
}

func (method *QueueDeclare) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.Queue, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.Passive = bits&(1<<0) != 0

	method.Durable = bits&(1<<1) != 0

	method.Exclusive = bits&(1<<2) != 0

	method.AutoDelete = bits&(1<<3) != 0

	method.NoWait = bits&(1<<4) != 0

	method.Arguments, err = ReadTable(reader)
	if err != nil {
		return err
	}

	return
}

func (method *QueueDeclare) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type QueueDeclareOk struct {
	Queue string

	MessageCount uint32

	ConsumerCount uint32
}

func (method *QueueDeclareOk) Name() string {
	return "QueueDeclareOk"
}

func (method *QueueDeclareOk) FrameType() byte {
	return 1
}

func (method *QueueDeclareOk) ClassIdentifier() uint16 {
	return 50
}

func (method *QueueDeclareOk) MethodIdentifier() uint16 {
	return 11
}

func (method *QueueDeclareOk) Read(reader io.Reader) (err error) {

	method.Queue, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.MessageCount, err = ReadLong(reader)
	if err != nil {
		return err
	}

	method.ConsumerCount, err = ReadLong(reader)
	if err != nil {
		return err
	}

	return
}

func (method *QueueDeclareOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type QueueBind struct {
	Reserved1 uint16

	Queue string

	Exchange string

	RoutingKey string

	NoWait bool

	Arguments *Table
}

func (method *QueueBind) Name() string {
	return "QueueBind"
}

func (method *QueueBind) FrameType() byte {
	return 1
}

func (method *QueueBind) ClassIdentifier() uint16 {
	return 50
}

func (method *QueueBind) MethodIdentifier() uint16 {
	return 20
}

func (method *QueueBind) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.Queue, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.Exchange, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.RoutingKey, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.NoWait = bits&(1<<0) != 0

	method.Arguments, err = ReadTable(reader)
	if err != nil {
		return err
	}

	return
}

func (method *QueueBind) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type QueueBindOk struct {
}

func (method *QueueBindOk) Name() string {
	return "QueueBindOk"
}

func (method *QueueBindOk) FrameType() byte {
	return 1
}

func (method *QueueBindOk) ClassIdentifier() uint16 {
	return 50
}

func (method *QueueBindOk) MethodIdentifier() uint16 {
	return 21
}

func (method *QueueBindOk) Read(reader io.Reader) (err error) {

	return
}

func (method *QueueBindOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type QueueUnbind struct {
	Reserved1 uint16

	Queue string

	Exchange string

	RoutingKey string

	Arguments *Table
}

func (method *QueueUnbind) Name() string {
	return "QueueUnbind"
}

func (method *QueueUnbind) FrameType() byte {
	return 1
}

func (method *QueueUnbind) ClassIdentifier() uint16 {
	return 50
}

func (method *QueueUnbind) MethodIdentifier() uint16 {
	return 50
}

func (method *QueueUnbind) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.Queue, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.Exchange, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.RoutingKey, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.Arguments, err = ReadTable(reader)
	if err != nil {
		return err
	}

	return
}

func (method *QueueUnbind) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type QueueUnbindOk struct {
}

func (method *QueueUnbindOk) Name() string {
	return "QueueUnbindOk"
}

func (method *QueueUnbindOk) FrameType() byte {
	return 1
}

func (method *QueueUnbindOk) ClassIdentifier() uint16 {
	return 50
}

func (method *QueueUnbindOk) MethodIdentifier() uint16 {
	return 51
}

func (method *QueueUnbindOk) Read(reader io.Reader) (err error) {

	return
}

func (method *QueueUnbindOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type QueuePurge struct {
	Reserved1 uint16

	Queue string

	NoWait bool
}

func (method *QueuePurge) Name() string {
	return "QueuePurge"
}

func (method *QueuePurge) FrameType() byte {
	return 1
}

func (method *QueuePurge) ClassIdentifier() uint16 {
	return 50
}

func (method *QueuePurge) MethodIdentifier() uint16 {
	return 30
}

func (method *QueuePurge) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.Queue, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.NoWait = bits&(1<<0) != 0

	return
}

func (method *QueuePurge) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type QueuePurgeOk struct {
	MessageCount uint32
}

func (method *QueuePurgeOk) Name() string {
	return "QueuePurgeOk"
}

func (method *QueuePurgeOk) FrameType() byte {
	return 1
}

func (method *QueuePurgeOk) ClassIdentifier() uint16 {
	return 50
}

func (method *QueuePurgeOk) MethodIdentifier() uint16 {
	return 31
}

func (method *QueuePurgeOk) Read(reader io.Reader) (err error) {

	method.MessageCount, err = ReadLong(reader)
	if err != nil {
		return err
	}

	return
}

func (method *QueuePurgeOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type QueueDelete struct {
	Reserved1 uint16

	Queue string

	IfUnused bool

	IfEmpty bool

	NoWait bool
}

func (method *QueueDelete) Name() string {
	return "QueueDelete"
}

func (method *QueueDelete) FrameType() byte {
	return 1
}

func (method *QueueDelete) ClassIdentifier() uint16 {
	return 50
}

func (method *QueueDelete) MethodIdentifier() uint16 {
	return 40
}

func (method *QueueDelete) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.Queue, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.IfUnused = bits&(1<<0) != 0

	method.IfEmpty = bits&(1<<1) != 0

	method.NoWait = bits&(1<<2) != 0

	return
}

func (method *QueueDelete) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type QueueDeleteOk struct {
	MessageCount uint32
}

func (method *QueueDeleteOk) Name() string {
	return "QueueDeleteOk"
}

func (method *QueueDeleteOk) FrameType() byte {
	return 1
}

func (method *QueueDeleteOk) ClassIdentifier() uint16 {
	return 50
}

func (method *QueueDeleteOk) MethodIdentifier() uint16 {
	return 41
}

func (method *QueueDeleteOk) Read(reader io.Reader) (err error) {

	method.MessageCount, err = ReadLong(reader)
	if err != nil {
		return err
	}

	return
}

func (method *QueueDeleteOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

// Basic methods

type BasicQos struct {
	PrefetchSize uint32

	PrefetchCount uint16

	Global bool
}

func (method *BasicQos) Name() string {
	return "BasicQos"
}

func (method *BasicQos) FrameType() byte {
	return 1
}

func (method *BasicQos) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicQos) MethodIdentifier() uint16 {
	return 10
}

func (method *BasicQos) Read(reader io.Reader) (err error) {

	method.PrefetchSize, err = ReadLong(reader)
	if err != nil {
		return err
	}

	method.PrefetchCount, err = ReadShort(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.Global = bits&(1<<0) != 0

	return
}

func (method *BasicQos) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicQosOk struct {
}

func (method *BasicQosOk) Name() string {
	return "BasicQosOk"
}

func (method *BasicQosOk) FrameType() byte {
	return 1
}

func (method *BasicQosOk) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicQosOk) MethodIdentifier() uint16 {
	return 11
}

func (method *BasicQosOk) Read(reader io.Reader) (err error) {

	return
}

func (method *BasicQosOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicConsume struct {
	Reserved1 uint16

	Queue string

	ConsumerTag string

	NoLocal bool

	NoAck bool

	Exclusive bool

	NoWait bool

	Arguments *Table
}

func (method *BasicConsume) Name() string {
	return "BasicConsume"
}

func (method *BasicConsume) FrameType() byte {
	return 1
}

func (method *BasicConsume) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicConsume) MethodIdentifier() uint16 {
	return 20
}

func (method *BasicConsume) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.Queue, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.ConsumerTag, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.NoLocal = bits&(1<<0) != 0

	method.NoAck = bits&(1<<1) != 0

	method.Exclusive = bits&(1<<2) != 0

	method.NoWait = bits&(1<<3) != 0

	method.Arguments, err = ReadTable(reader)
	if err != nil {
		return err
	}

	return
}

func (method *BasicConsume) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicConsumeOk struct {
	ConsumerTag string
}

func (method *BasicConsumeOk) Name() string {
	return "BasicConsumeOk"
}

func (method *BasicConsumeOk) FrameType() byte {
	return 1
}

func (method *BasicConsumeOk) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicConsumeOk) MethodIdentifier() uint16 {
	return 21
}

func (method *BasicConsumeOk) Read(reader io.Reader) (err error) {

	method.ConsumerTag, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *BasicConsumeOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicCancel struct {
	ConsumerTag string

	NoWait bool
}

func (method *BasicCancel) Name() string {
	return "BasicCancel"
}

func (method *BasicCancel) FrameType() byte {
	return 1
}

func (method *BasicCancel) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicCancel) MethodIdentifier() uint16 {
	return 30
}

func (method *BasicCancel) Read(reader io.Reader) (err error) {

	method.ConsumerTag, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.NoWait = bits&(1<<0) != 0

	return
}

func (method *BasicCancel) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicCancelOk struct {
	ConsumerTag string
}

func (method *BasicCancelOk) Name() string {
	return "BasicCancelOk"
}

func (method *BasicCancelOk) FrameType() byte {
	return 1
}

func (method *BasicCancelOk) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicCancelOk) MethodIdentifier() uint16 {
	return 31
}

func (method *BasicCancelOk) Read(reader io.Reader) (err error) {

	method.ConsumerTag, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *BasicCancelOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicPublish struct {
	Reserved1 uint16

	Exchange string

	RoutingKey string

	Mandatory bool

	Immediate bool
}

func (method *BasicPublish) Name() string {
	return "BasicPublish"
}

func (method *BasicPublish) FrameType() byte {
	return 1
}

func (method *BasicPublish) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicPublish) MethodIdentifier() uint16 {
	return 40
}

func (method *BasicPublish) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.Exchange, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.RoutingKey, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.Mandatory = bits&(1<<0) != 0

	method.Immediate = bits&(1<<1) != 0

	return
}

func (method *BasicPublish) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicReturn struct {
	ReplyCode uint16

	ReplyText string

	Exchange string

	RoutingKey string
}

func (method *BasicReturn) Name() string {
	return "BasicReturn"
}

func (method *BasicReturn) FrameType() byte {
	return 1
}

func (method *BasicReturn) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicReturn) MethodIdentifier() uint16 {
	return 50
}

func (method *BasicReturn) Read(reader io.Reader) (err error) {

	method.ReplyCode, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.ReplyText, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.Exchange, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.RoutingKey, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *BasicReturn) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicDeliver struct {
	ConsumerTag string

	DeliveryTag uint64

	Redelivered bool

	Exchange string

	RoutingKey string
}

func (method *BasicDeliver) Name() string {
	return "BasicDeliver"
}

func (method *BasicDeliver) FrameType() byte {
	return 1
}

func (method *BasicDeliver) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicDeliver) MethodIdentifier() uint16 {
	return 60
}

func (method *BasicDeliver) Read(reader io.Reader) (err error) {

	method.ConsumerTag, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.DeliveryTag, err = ReadLonglong(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.Redelivered = bits&(1<<0) != 0

	method.Exchange, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.RoutingKey, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *BasicDeliver) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicGet struct {
	Reserved1 uint16

	Queue string

	NoAck bool
}

func (method *BasicGet) Name() string {
	return "BasicGet"
}

func (method *BasicGet) FrameType() byte {
	return 1
}

func (method *BasicGet) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicGet) MethodIdentifier() uint16 {
	return 70
}

func (method *BasicGet) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShort(reader)
	if err != nil {
		return err
	}

	method.Queue, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.NoAck = bits&(1<<0) != 0

	return
}

func (method *BasicGet) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicGetOk struct {
	DeliveryTag uint64

	Redelivered bool

	Exchange string

	RoutingKey string

	MessageCount uint32
}

func (method *BasicGetOk) Name() string {
	return "BasicGetOk"
}

func (method *BasicGetOk) FrameType() byte {
	return 1
}

func (method *BasicGetOk) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicGetOk) MethodIdentifier() uint16 {
	return 71
}

func (method *BasicGetOk) Read(reader io.Reader) (err error) {

	method.DeliveryTag, err = ReadLonglong(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.Redelivered = bits&(1<<0) != 0

	method.Exchange, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.RoutingKey, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	method.MessageCount, err = ReadLong(reader)
	if err != nil {
		return err
	}

	return
}

func (method *BasicGetOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicGetEmpty struct {
	Reserved1 string
}

func (method *BasicGetEmpty) Name() string {
	return "BasicGetEmpty"
}

func (method *BasicGetEmpty) FrameType() byte {
	return 1
}

func (method *BasicGetEmpty) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicGetEmpty) MethodIdentifier() uint16 {
	return 72
}

func (method *BasicGetEmpty) Read(reader io.Reader) (err error) {

	method.Reserved1, err = ReadShortstr(reader)
	if err != nil {
		return err
	}

	return
}

func (method *BasicGetEmpty) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicAck struct {
	DeliveryTag uint64

	Multiple bool
}

func (method *BasicAck) Name() string {
	return "BasicAck"
}

func (method *BasicAck) FrameType() byte {
	return 1
}

func (method *BasicAck) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicAck) MethodIdentifier() uint16 {
	return 80
}

func (method *BasicAck) Read(reader io.Reader) (err error) {

	method.DeliveryTag, err = ReadLonglong(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.Multiple = bits&(1<<0) != 0

	return
}

func (method *BasicAck) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicReject struct {
	DeliveryTag uint64

	Requeue bool
}

func (method *BasicReject) Name() string {
	return "BasicReject"
}

func (method *BasicReject) FrameType() byte {
	return 1
}

func (method *BasicReject) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicReject) MethodIdentifier() uint16 {
	return 90
}

func (method *BasicReject) Read(reader io.Reader) (err error) {

	method.DeliveryTag, err = ReadLonglong(reader)
	if err != nil {
		return err
	}

	bits, err := ReadOctet(reader)

	method.Requeue = bits&(1<<0) != 0

	return
}

func (method *BasicReject) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicRecoverAsync struct {
	Requeue bool
}

func (method *BasicRecoverAsync) Name() string {
	return "BasicRecoverAsync"
}

func (method *BasicRecoverAsync) FrameType() byte {
	return 1
}

func (method *BasicRecoverAsync) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicRecoverAsync) MethodIdentifier() uint16 {
	return 100
}

func (method *BasicRecoverAsync) Read(reader io.Reader) (err error) {

	bits, err := ReadOctet(reader)

	method.Requeue = bits&(1<<0) != 0

	return
}

func (method *BasicRecoverAsync) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicRecover struct {
	Requeue bool
}

func (method *BasicRecover) Name() string {
	return "BasicRecover"
}

func (method *BasicRecover) FrameType() byte {
	return 1
}

func (method *BasicRecover) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicRecover) MethodIdentifier() uint16 {
	return 110
}

func (method *BasicRecover) Read(reader io.Reader) (err error) {

	bits, err := ReadOctet(reader)

	method.Requeue = bits&(1<<0) != 0

	return
}

func (method *BasicRecover) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type BasicRecoverOk struct {
}

func (method *BasicRecoverOk) Name() string {
	return "BasicRecoverOk"
}

func (method *BasicRecoverOk) FrameType() byte {
	return 1
}

func (method *BasicRecoverOk) ClassIdentifier() uint16 {
	return 60
}

func (method *BasicRecoverOk) MethodIdentifier() uint16 {
	return 111
}

func (method *BasicRecoverOk) Read(reader io.Reader) (err error) {

	return
}

func (method *BasicRecoverOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

// Tx methods

type TxSelect struct {
}

func (method *TxSelect) Name() string {
	return "TxSelect"
}

func (method *TxSelect) FrameType() byte {
	return 1
}

func (method *TxSelect) ClassIdentifier() uint16 {
	return 90
}

func (method *TxSelect) MethodIdentifier() uint16 {
	return 10
}

func (method *TxSelect) Read(reader io.Reader) (err error) {

	return
}

func (method *TxSelect) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type TxSelectOk struct {
}

func (method *TxSelectOk) Name() string {
	return "TxSelectOk"
}

func (method *TxSelectOk) FrameType() byte {
	return 1
}

func (method *TxSelectOk) ClassIdentifier() uint16 {
	return 90
}

func (method *TxSelectOk) MethodIdentifier() uint16 {
	return 11
}

func (method *TxSelectOk) Read(reader io.Reader) (err error) {

	return
}

func (method *TxSelectOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type TxCommit struct {
}

func (method *TxCommit) Name() string {
	return "TxCommit"
}

func (method *TxCommit) FrameType() byte {
	return 1
}

func (method *TxCommit) ClassIdentifier() uint16 {
	return 90
}

func (method *TxCommit) MethodIdentifier() uint16 {
	return 20
}

func (method *TxCommit) Read(reader io.Reader) (err error) {

	return
}

func (method *TxCommit) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type TxCommitOk struct {
}

func (method *TxCommitOk) Name() string {
	return "TxCommitOk"
}

func (method *TxCommitOk) FrameType() byte {
	return 1
}

func (method *TxCommitOk) ClassIdentifier() uint16 {
	return 90
}

func (method *TxCommitOk) MethodIdentifier() uint16 {
	return 21
}

func (method *TxCommitOk) Read(reader io.Reader) (err error) {

	return
}

func (method *TxCommitOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type TxRollback struct {
}

func (method *TxRollback) Name() string {
	return "TxRollback"
}

func (method *TxRollback) FrameType() byte {
	return 1
}

func (method *TxRollback) ClassIdentifier() uint16 {
	return 90
}

func (method *TxRollback) MethodIdentifier() uint16 {
	return 30
}

func (method *TxRollback) Read(reader io.Reader) (err error) {

	return
}

func (method *TxRollback) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}

type TxRollbackOk struct {
}

func (method *TxRollbackOk) Name() string {
	return "TxRollbackOk"
}

func (method *TxRollbackOk) FrameType() byte {
	return 1
}

func (method *TxRollbackOk) ClassIdentifier() uint16 {
	return 90
}

func (method *TxRollbackOk) MethodIdentifier() uint16 {
	return 31
}

func (method *TxRollbackOk) Read(reader io.Reader) (err error) {

	return
}

func (method *TxRollbackOk) Write(writer io.Writer) (err error) {
	return errors.New("to do")
}