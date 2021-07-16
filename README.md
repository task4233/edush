# edush


## シーケンス図(雑)
- シェルっぽい画面からwebsocketでコマンドを送信する。
![](./images/edush-2.jpg)

## 実現の課題点
- ユーザーアカウントの作成
- 対戦の実現
    - 最終的には1対1
    - 今は一人でゲームができる状態にする。
- その他
    - セキュリティ
       - 権限の制限
    - パフォーマンス
        - コンテナぽこぽこ立てて大丈夫なもんなのかな

## とりあえずの方向
### フェーズ1
コマンドを送信して、判定結果が返ってくるまでの流れを作成する。
- [x] websocketからコマンドを送信を実装
- [x] ユーザーそれぞれに異なる仮想環境を与える
- [] 文字列一致による簡単なjudgeを作成

### フェーズ2
ユーザーの識別をして、誰が勝ったかわかるようにする。
- [] 正答したユーザにwin,その他のユーザにlooseを送信する。

### フェーズ3
ユーザのコマンドをどこで実行させるのかを検討する。
- ルームに適応したコンテナを生成し、そちらのシェルで実行させたい。
    - 実現方法がイメージできないので要調査
