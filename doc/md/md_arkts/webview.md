### 依赖导包

```typescript
import { webview } from '@kit.ArkWeb';
```

```dart
//Flutter
  flutter_inappwebview:
    git:
      url: 'https://gitee.com/openharmony-sig/flutter_inappwebview.git'
      path: './flutter_inappwebview'
```

```dart
//Flutter example
void main() => runApp(const MaterialApp(home: HomePage()));
class HomePage extends StatefulWidget {
  const HomePage({super.key});
  @override
  State<HomePage> createState() => _HomePageState();
}
class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: _buildBody(),
    );
  }
  Widget _buildBody() {
    return Column(
      children: [
        Text('WebView'),
      ],
    );
  }
}
```

### 工单地址

[华为ir单](https://issuereporter.developer.huawei.com/detail/240920163159075/comment)

### 参考文档

[ohos.web.webview (Webview)](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-webview-V5)

**[flutter_inappwebview](https://gitee.com/openharmony-sig/flutter_inappwebview)**