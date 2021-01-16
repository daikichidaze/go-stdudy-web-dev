# ベースとなるDockerイメージ指定
FROM golang:latest
# 各種インストール
RUN apt update
RUN apt install sudo

# コンテナログイン時のディレクトリ指定
WORKDIR /go/src
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src