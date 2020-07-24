# Wol-on-Web

あらかじめ登録したコンピューターに、Wake on LANを送信することができるWebアプリケーションです。
Go言語で記述したAPIサーバーとHTMLの簡易フロントエンドを組み合わせているため、少ないリソースで動作します。

# 必要な環境

  * Docker
  * Docker-Compose

# インストール方法

Docker上で使用する場合、このリポジトリをgit cloneした上でDocker-Composeで動かすのが簡単です。

```bash
git clone https://github.com/ytsurui/wol-on-web
cd wol-on-web
docker-compose up -d --build
```

# 設定

コンテナの起動前に、config.jsonを編集して動作を制御できます。

  * httpport: ポート80番以外で動作させる場合、この番号を変更してください。
  * readonly: trueにすると、コンピューターの追加機能が常にエラーとなり、動作しなくなります。
  * machines: 登録されているコンピューターのリストです。
