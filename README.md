# Wol-on-Web

あらかじめ登録したコンピューターに、Wake on LANを送信することができるWebアプリケーションです。  
Go言語で記述したAPIサーバーとHTMLの簡易フロントエンドを組み合わせているため、少ないリソースで動作します。

# 必要な環境

Docker上で動かす場合

  * Docker
  * Docker-Compose (あると便利)

スタンドアロンで動かす場合

  * Go (1.14.6で動作確認)

# インストール方法

  * Docker上で使用する場合

このリポジトリをgit cloneした上でDocker-Composeで動かすのが簡単です。

```bash
git clone https://github.com/ytsurui/wol-on-web
cd wol-on-web
docker-compose up -d --build
```

  * スタンドアロンで使用する場合

このリポジトリをgit cloneした上で、main.goをビルドしてください。
ビルド前にgoの環境が整っていることが前提です。

```bash
git clone https://github.com/ytsurui/wol-on-web
cd wol-on-web/webapi
go build -o wolonweb
```

下記のライブラリが必要となりますので、事前にインストールしてください。

```bash
go get github.com/gorilla/mux
go get github.com/mdlayher/wol
go get github.com/sparrc/go-ping
```

# 必要なライブラリ

  * Docker上で使用する場合

ビルド時に自動的に組み込まれるため、独自にインストールする必要はありません。

  * スタンドアロンで使用する場合

github.com/gorilla/mux
github.com/mdlayher/wol
github.com/sparrc/go-ping

# 設定

コンテナの起動前に、config.jsonを編集して動作を制御できます。

  * httpport: ポート80番以外で動作させる場合、この番号を変更してください。
  * readonly: trueにすると、コンピューターの追加機能が常にエラーとなり、動作しなくなります。
  * machines: 登録されているコンピューターのリストです。
