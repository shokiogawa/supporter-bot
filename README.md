# supporter-bot
## 開発
DockerフォルダのDocker.developを使用。
ホットリロードするため。

## 本番(Fargateに上げる用)
ルートディレクトリのDockerfileを使用。
マルチステージビルドをし、本番用にイメージの容量を軽量化しているため。
