# Taraxa-exporter

Экспортер метрик для узла сети Taraxa.

### Метрики который отдает экспортер

    SyncStatus - Статус Синхронизации ноды
    SyncTime -  Время в течении которого нода синхронизирована

### Как собрать бинарный файл и запустить экспортер

```bash
cp .env.example .env
go build -o taraxa-exporter .
./taraxa-exporter
```

##### TODO:

1.  Добавить больше метрик
2.  Описать docker-образ
