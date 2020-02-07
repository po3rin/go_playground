# multiple gorutine shutdown

起動している複数のgorutineが１つでもエラーを出したら、もしくはSIGTERMなどのシグナルを受け取ったら、起動している全てのgorutineをきっちり終了させるパターンの実装&考察

* counter : errチャネルのカウントによる実装
* wait : sync.WaitGroupによる実装

どちらもコンテキストキャンセルしたらgorutineで起動しているプロセスが確実に終了することが前提の実装になっているが、waitの方がコードがシンプルになって良い。

```
# コンテキストタイムアウト(関数がエラーを出したと想定)
start!
received context cancel
2020/02/08 02:01:47 context deadline exceeded
2020/02/08 02:01:47 context deadline exceeded
2020/02/08 02:01:47 http: Server closed
done

# Ctrl+C で止める
start!
^Creceived signal
2020/02/08 02:01:42 context canceled
2020/02/08 02:01:42 context canceled
2020/02/08 02:01:42 http: Server closed
done
```