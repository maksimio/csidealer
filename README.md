## Назначение
Сервер предназначен для приема и декодирования посылок с CSI, отправленных по TCP с роутера. На роутере установлена [модифицированная](https://github.com/xieyaxiongfly/Atheros_CSI_tool_OpenWRT_src) прошивка OpenWRT с [Atheros CSI Tool](https://wands.sg/research/wifi/AtherosCSI/). Для отправки отправки пакетов с CSI на роутере используется функция client_main [ip] [port].

## В планах
1. Tcp сервер
2. Чтение CSI
3. Хранение CSI
4. Сохранение CSI в файл
5. Сохранение CSI в БД
6. HTTP сервер для выдачи CSI
7. WebSocket сервер для выдачи CSI
8. Web-интерфейс на React для рисования графиков и вывода статистики, который обращается к HTTP и WebSocket серверам
