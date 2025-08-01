## Server

### 构建

```go
type Server struct {
    Addr                         string
    Handler                      Handler
    DisableGeneralOptionsHandler bool
    TLSConfig                    *tls.Config
    ReadTimeout                  time.Duration
    ReadHeaderTimeout            time.Duration
    WriteTimeout                 time.Duration
    IdleTimeout                  time.Duration
    MaxHeaderBytes               int
    TLSNextProto                 map[string]func(*Server, *tls.Conn, Handler)
    ConnState                    func(net.Conn, ConnState)
    ErrorLog                     *log.Logger
    BaseContext                  func(net.Listener) context.Context
    ConnContext                  func(ctx context.Context, c net.Conn) context.Context
    HTTP2                        *HTTP2Config
    Protocols                    *Protocols
    inShutdown                   atomic.Bool
    disableKeepAlives            atomic.Bool
    nextProtoOnce                sync.Once
    nextProtoErr                 error
    mu                           sync.Mutex
    listeners                    map[*net.Listener]struct{}
    activeConn                   map[*conn]struct{}
    onShutdown                   []func()
    listenerGroup                sync.WaitGroup
}
```

### 方法

```go
//方法对象: (*Server):
newConn(rwc net.Conn) *http.conn
maxHeaderBytes() int
initialReadLimitSize() int64
tlsHandshakeTimeout() time.Duration
Close() error
Shutdown(ctx context.Context) error
RegisterOnShutdown(f func())
closeIdleConns() bool
closeListenersLocked() error
ListenAndServe() error
shouldConfigureHTTP2ForServe() bool
Serve(l net.Listener) error
ServeTLS(l net.Listener, certFile string, keyFile string) error
protocols() http.Protocols
trackListener(ln *net.Listener, add bool) bool
trackConn(c *http.conn, add bool)
idleTimeout() time.Duration
readHeaderTimeout() time.Duration
doKeepAlives() bool
shuttingDown() bool
SetKeepAlivesEnabled(v bool)
logf(format string, args ...any)
ListenAndServeTLS(certFile string, keyFile string) error
setupHTTP2_ServeTLS() error
setupHTTP2_Serve() error
onceSetNextProtoDefaults_Serve()
onceSetNextProtoDefaults()
ExportAllConnsIdle() bool
ExportAllConnsByState() map[http.ConnState]int
```

