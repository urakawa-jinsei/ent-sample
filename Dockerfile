FROM postgres:latest

# SSL証明書のコピー
COPY ./ssl/server.crt /var/lib/postgresql/server.crt
COPY ./ssl/server.key /var/lib/postgresql/server.key

# PostgreSQLの設定ファイルを上書き
COPY ./postgresql.conf /etc/postgresql/postgresql.conf

# SSL証明書の所有者とパーミッションの変更
RUN chown postgres:postgres /var/lib/postgresql/server.crt
RUN chown postgres:postgres /var/lib/postgresql/server.key
RUN chmod 600 /var/lib/postgresql/server.crt
RUN chmod 600 /var/lib/postgresql/server.key
