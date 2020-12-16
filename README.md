# エラー機構

エラーの処理を共通化するための参考実装です。

# 使い方

このerrorsパッケージを自身アプリケーション以下にコピーしてください。

```
├── src
│   ├── api
│   ├── app
│   ├── domain
│   └── infra
└── errors ←
```

## 定義する方法

新しいエラーは define_xxx.go に定義してください

```
var エラー定義名 = newXXX(エラーコード, メッセージ)
```

```
var (
    ErrInvalid = newBadRequest("invalid", "ユーザ側のエラーです")
    ErrSystem = newInternalServerError("system", "サーバ側のエラーです")
)
```

## 呼び出す方法

### ラップしたい場合

```
if err := Update(); err != nil {
    return errors.Wrap(err)
}

or  (独自で定義したエラーを使う)
if err != nil {
    return errors.UpdateUserData.Wrap(err, "UpdateUserData")
}
```

### 新しくエラーを発生させたい場合

```
if device != `` {
    return errors.New("unknown device")
}

or

if device != `` {
    return errors.Errorf("AuthService device != ``")
}

or

if device != `` {
    return errors.InvalidDevice.New("AuthService device != ``")
}
```

[エラーコード表](define)
