## Module net

### TcpListener

#### 方法bind

```rust
pub fn bind<A: ToSocketAddrs>(addr: A) -> Result<TcpListener>
```

```rust
//方法bind
/*
Creates a TCP listener bound to 127.0.0.1:80. 
If that fails, create a TCP listener bound to 127.0.0.1:443:
*/
use std::net::{SocketAddr, TcpListener};
let addrs = [
    SocketAddr::from(([127, 0, 0, 1], 80)),
    SocketAddr::from(([127, 0, 0, 1], 443)),
];
let listener = TcpListener::bind(&addrs[..]).unwrap();
```

#### 方法local_addr

```rust
pub fn local_addr(&self) -> Result<SocketAddr>
```

```rust
use std::net::{Ipv4Addr, SocketAddr, SocketAddrV4, TcpListener};

let listener = TcpListener::bind("127.0.0.1:8080").unwrap();
assert_eq!(listener.local_addr().unwrap(),
           SocketAddr::V4(SocketAddrV4::new(Ipv4Addr::new(127, 0, 0, 1), 8080)));
```

#### 方法try_clone

```rust
pub fn try_clone(&self) -> Result<TcpListener>
```

```rust
use std::net::TcpListener;
let listener = TcpListener::bind("127.0.0.1:8080").unwrap();
let listener_clone = listener.try_clone().unwrap();
```

#### 方法accept

```
pub fn accept(&self) -> Result<(TcpStream, SocketAddr)>
```

```rust
use std::net::TcpListener;
let listener = TcpListener::bind("127.0.0.1:8080").unwrap();
match listener.accept() {
    Ok((_socket, addr)) => println!("new client: {addr:?}"),
    Err(e) => println!("couldn't get client: {e:?}"),
}
```

#### 方法incoming

```
pub fn incoming(&self) -> Incoming<'_> 
```

```rust
use std::net::{TcpListener, TcpStream};
fn handle_connection(stream: TcpStream) {
   //...
}
fn main() -> std::io::Result<()> {
    let listener = TcpListener::bind("127.0.0.1:80")?;
    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                handle_connection(stream);
            }
            Err(e) => { /* connection failed */ }
        }
    }
    Ok(())
}
```

#### 方法into_incoming

```
pub fn into_incoming(self) -> IntoIncoming
```

```rust
#![feature(tcplistener_into_incoming)]
use std::net::{TcpListener, TcpStream};
fn listen_on(port: u16) -> impl Iterator<Item = TcpStream> {
    let listener = TcpListener::bind(("127.0.0.1", port)).unwrap();
    listener.into_incoming()
        .filter_map(Result::ok) /* Ignore failed connections */
}
fn main() -> std::io::Result<()> {
    for stream in listen_on(80) {
        /* handle the connection here */
    }
    Ok(())
}
```

#### 方法set_ttl

```
pub fn set_ttl(&self, ttl: u32) -> Result<()>
```

```rust
use std::net::TcpListener;
let listener = TcpListener::bind("127.0.0.1:80").unwrap();
listener.set_ttl(100).expect("could not set TTL");
```

#### 方法ttl

```
pub fn ttl(&self) -> Result<u32>
```

```
use std::net::TcpListener;
let listener = TcpListener::bind("127.0.0.1:80").unwrap();
listener.set_ttl(100).expect("could not set TTL");
assert_eq!(listener.ttl().unwrap_or(0), 100);
```

#### 方法take_error

```
pub fn take_error(&self) -> Result<Option<Error>>
```

```rust
use std::net::TcpListener;
let listener = TcpListener::bind("127.0.0.1:80").unwrap();
listener.take_error().expect("No error was expected");
```

#### 方法set_nonblocking

```
pub fn set_nonblocking(&self, nonblocking: bool) -> Result<()>
```

```rust
use std::io;
use std::net::TcpListener;
let listener = TcpListener::bind("127.0.0.1:7878").unwrap();
listener.set_nonblocking(true).expect("Cannot set non-blocking");
for stream in listener.incoming() {
    match stream {
        Ok(s) => {
            // do something with the TcpStream
            handle_connection(s);
        }
        Err(ref e) if e.kind() == io::ErrorKind::WouldBlock => {
            // wait until network socket is ready, typically implemented
            // via platform-specific APIs such as epoll or IOCP
            wait_for_fd();
            continue;
        }
        Err(e) => panic!("encountered IO error: {e}"),
    }
}
```

#### 其他方法

```rust
pub fn set_only_v6(&self, only_v6: bool) -> Result<()>
pub fn only_v6(&self) -> Result<bool>
```

#### 实现其他接口

```
```



#### 简例01

```rust
use std::net::{TcpListener, TcpStream};

fn handle_client(stream: TcpStream) {
    // ...
}

fn main() -> std::io::Result<()> {
    let listener = TcpListener::bind("127.0.0.1:80")?;

    // accept connections and process them serially
    for stream in listener.incoming() {
        handle_client(stream?);
    }
    Ok(())
}
```

#### 

## 参考文档

[net官方文档](https://doc.rust-lang.org/std/net/index.html)