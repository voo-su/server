# VooSu Server

<img src="./assets/logo.svg" align="left" width=150 height=150 alt="VooSu logo">

[VooSu](https://voo.su) - это открытая платформа для мгновенного обмена сообщениями, разработанная для бизнеса, который
ценит прозрачность и гибкость.

Платформа предлагает возможность настройки под любые бизнес-процессы, обеспечивая легкость внедрения и масштабируемость.

[VooSu](https://voo.su) поддерживает интеграцию с другими системами и предоставляет полную свободу в управлении
коммуникациями,
что делает её отличной альтернативой коммерческим мессенджерам для команд и организаций любого размера.

---

### Репозитории

[Репозиторий веб-версии](https://github.com/voo-su/web)

[Репозиторий приложений (Android, iOS, Desktop)](https://github.com/voo-su/app)

---

### Процедура установки

```bash
git clone https://github.com/voo-su/server.git
```

```bash
cd server
```

```bash
mkdir web/web-client && git clone https://github.com/voo-su/web.git web/web-client
```

```bash
docker-compose up -d
```
