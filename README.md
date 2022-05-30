📚[Ссылка](https://github.com/maksimio/csi_classification) на главную проекта

## Назначение
Сервер предназначен для приема и декодирования посылок с CSI, отправленных по TCP с роутера. На роутере установлена [модифицированная](https://github.com/xieyaxiongfly/Atheros_CSI_tool_OpenWRT_src) прошивка OpenWRT с [Atheros CSI Tool](https://wands.sg/research/wifi/AtherosCSI/). Для отправки отправки пакетов с CSI на роутере используется функция client_main [ip] [port].

## Задачи


## REST API
Обращение ко всем запросам идет по адресу /api/v1/

1. Число пакетов, считанное с начала работы сервера
2. Сырые значения CSI
3. Амплитудные значения CSI
4. Фазовые значения CSI
