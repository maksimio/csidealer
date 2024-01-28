<p align="center">
  <img src="https://user-images.githubusercontent.com/61945327/201778567-fee234ff-84f1-459e-b3b3-cb0d96cc0a68.png" height="100">
</p>

<p align="center">
  <a href="https://github.com/maksimio/smartwifi">smart Wi-Fi</a>
</p>

# Сервер-менеджер посылок CSI
Сервер принимает, декодирует, записывает и раздает посылки CSI, отправленных по TCP с роутера (функция client_main). На роутере установлена [модифицированная](https://github.com/xieyaxiongfly/Atheros_CSI_tool_OpenWRT_src) прошивка OpenWRT с [Atheros CSI Tool](https://wands.sg/research/wifi/AtherosCSI/)

ПО состоит из *интерфейса*, написанного на React + Typescript и *сервера*, написанного на Go

# Клиент
Для сервера предусмотрен интерфейс, который позволяет управлять записью CSI в файл и маршрутизаторами, а также выводить графики CSI в реальном времени с высокой частото (более 100 Гц) через WebGL.

<p align="center">
  <img src="https://user-images.githubusercontent.com/61945327/201782299-123e2466-f490-4690-8e44-a1d2f42c0b54.png" width="500px">
  <img src="https://user-images.githubusercontent.com/61945327/201782498-819afe7e-220b-4652-bd04-7f6145a8302c.png" width="500px">
</p>

# Сервер
Назначение:
1. Прием и декодирование CSI из сырого TCP-трафика
2. REST API: взять n последних пакетов с CSI (ABS, PHASE, IM, RE), начать / остановить запись в файл и т.д.
3. Передача CSI по WebSocket
4. Фильтрация пакетов CSI

Архитектура сервера на Golang:
<p align="center">
  <img src="https://user-images.githubusercontent.com/61945327/201779656-1fff5106-80fc-4d36-9935-777b8abf2a8e.png" width="800">
</p>

## REST API
Актуальное API можно найти в файле `server\internal\controller\http\apiv1.go` Обращение ко всем запросам идет по адресу /api/v1/. Ключевые запросы:

1. `csi/last_n/:type` - параметры: n.
2. `csi/subcarrier_last_n/:type` - параметры: n.

## Запуск и отладка
1. Запустить сервер командой `go run .\server\cmd\app\main.go`
2. Запустить эмулятор работы client_main командой `go run .\server\test\faketcp.go`
3. Из другого терминала перейти в `client` и ввести `npm i`, `npm start`
