## Rcp 远场通信协议

### 依赖导包

```typescript
import { rcp } from '@kit.RemoteCommunicationKit';
```

### 接口类SessionConfiguration

```typescript
//源码
export interface SessionConfiguration {
  interceptors?: Interceptor[];
  requestConfiguration?: Configuration;
  baseAddress?: URLOrString;
  headers?: RequestHeaders;
  cookies?: RequestCookies;
  sessionListener?: SessionListener;
  connectionConfiguration?: ConnectionConfiguration;
}
```

### 类Interceptor

```typescript
export interface Interceptor {
  intercept(context: RequestContext, next: RequestHandler): Promise<Response>;
}
```

### 类Configuration

```typescript
export interface Configuration {
  transfer?: TransferConfiguration;
  tracing?: TracingConfiguration;
  proxy?: ProxyConfiguration;
  dns?: DnsConfiguration;
  security?: SecurityConfiguration;
  processing?: ProcessingConfiguration;
}
```

### 类TransferConfiguration

```typescript
export interface TransferConfiguration {
  autoRedirect:boolean
  maxAutoRedirects?: number;
  timeout?: Timeout;
  assumesHTTP3Capable?: boolean;
  pathPreference?: PathPreference;
  serviceType?: ServiceType;
  pausePolicy?: PausePolicy;
}
```

### 类型URLOrString

```
export type URLOrString = URL | string;
```

### 类URL

```
export type URL = url.URL;
```

```typescript
class URL {
  constructor(url: string, base?: string | URL);
  constructor();
  static parseURL(url: string, base?: string | URL): URL;
  toString(): string;
  toJSON(): string;
  hash: string;
  host: string;
  hostname: string;
  href: string;
  readonly origin: string;
  password: string;
  pathname: string;
  port: string;
  protocol: string;
  search: string;
  readonly searchParams: URLSearchParams;
  readonly params: URLParams;
  username: string;
}
```

### 类型RequestHeaders

```typescript
export type RequestHeaders = {
        [k: string]: string | string[] | undefined;
        'authorization'?: string;
        'accept'?: ContentType | ContentType[];
        'accept-charset'?: string | string[];
        'accept-encoding'?: ContentCoding | ContentCoding[];
        'accept-language'?: string | string[];
        'cache-control'?: string | string[];
        'cookie'?: string | string[];
        'range'?: string | string[];
        'upgrade'?: string | string[];
        'user-agent'?: string;
        'content-type'?: ContentType;
    };
```

### 类型ResponseHeaders

```typescript
export type ResponseHeaders = {
        [k: string]: string | string[] | undefined;
        'accept-ranges'?: 'none' | 'bytes' | (string & NonNullable<unknown>);
        'allow'?: HttpMethod | HttpMethod[];
        'cache-control'?: string | string[];
        'content-encoding'?: ContentCoding;
        'content-range'?: string;
        'content-type'?: ContentType;
        'date'?: string;
        'etag'?: string;
        'expires'?: string;
        'location'?: string;
        'retry-after'?: string;
        'set-cookie'?: string | string[];
        'server'?: string;
        'www-authenticate'?: string | string[];
    };
```

### 类型HttpMethod

```typescript
export type HttpMethod = 'GET' | 'POST' | 'HEAD' | 'PUT' | 'DELETE' | 'PATCH' | 'OPTIONS' | (string & NonNullable<unknown>);
```

### 接口RequestCookies

```typescript
export interface RequestCookies {
        [name: string]: string;
    }
```

### 类Request

```typescript
export class Request {
  readonly id: string;
  url: URL;
  method: HttpMethod;
  headers?: RequestHeaders;
  content?: RequestContent;
  cookies?: RequestCookies;
  transferRange?: TransferRange | TransferRange[];
  configuration?: Configuration;
  destination?: ResponseBodyDestination;
  constructor(url: URLOrString, method?: HttpMethod, headers?: RequestHeaders, content?: RequestContent, cookies?: RequestCookies, transferRange?: TransferRange | TransferRange[], configuration?: Configuration);
}
```

### 接口类SessionListener

```typescript
export interface SessionListener {
  onCanceled?: OnCanceled;
  onClosed?: OnClosed;
}
```

```typescript
export type OnCanceled = () => void;
export type OnClosed = () => void;
```

### 类ConnectionConfiguration

```typescript
export interface ConnectionConfiguration {
  readonly maxConnectionsPerHost?: number;
  readonly maxTotalConnections?: number;
}
```

