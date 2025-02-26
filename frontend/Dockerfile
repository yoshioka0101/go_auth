# ベースイメージ
FROM node:18

# 作業ディレクトリの設定
WORKDIR /app

# package.json と package-lock.json をコピー
COPY package.json package-lock.json ./

# 依存関係をインストール
RUN npm install

# アプリのソースコードをコピー
COPY . .

# ポートの設定
EXPOSE 3000

# アプリケーションの起動
CMD ["npm", "start"]
