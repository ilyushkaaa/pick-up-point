## Домашнее задание №2 «Хранение информации о ПВЗ»
### Цель:

Модифицировать ваш сервис, который вы делали в первом домашнем задании, добавить возможность поддержания двух потоков. Данные при работе в двух потоках должны быть потокобезопасные. Приложение должно работать в интерактивном режиме.

### Задание:

- **Модифицируйте Go-приложение**, добавьте (если не было) простую систему для хранения ПВЗ (пунктов выдачи заказов) и предоставлять возможность записи и чтения данных.
- Реализуйте функционал для **записи данных о ПВЗ в файл (мапку)**. Каждый ПВЗ должен содержать информацию, такую как название, адрес и контактные данные.
- Реализуйте функционал для **чтения данных из файла (мапки)** и вывода информации о ПВЗ.
- Приложение должно поддерживать возможность работы **двух потоков одновременно:** один для записи, другой для чтения.
- Приложение должно использовать **механизм блокировки** (например, мьютекс) для синхронизации доступа к данным между процессами записи и чтения.

### Дополнительные задания на 10 баллов:

- Добавьте обработку системных сигналов, таких как SIGINT (Ctrl+C) и SIGTERM, чтобы корректно завершить работу приложения.
- Приложение должно выводить информацию о текущем состоянии для каждого потока. Реализовать в отдельном потоке.

### Подсказки:

- Для реагирования на системные сигналы, вы можете использовать пакет os/signal в Go.
- Для синхронизации доступа к файлу между процессами записи и чтения, вы можете использовать пакет sync и его мьютексы.
- Для работы с файлами в Go, вы можете использовать пакет os или io/ioutil.

### Deadline
16 марта, 23:59 (сдача) / 19 марта, 23:59 (проверка)
