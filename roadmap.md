# Road Map развития
Хочу тут собрать что нужно знать, уметь на определенном уровне. (Это всего лишь мое скромное мнение, не Must have). 
(Полезные понятные материалы приветствуются)

# June
- git
    - https://git-scm.com/book/ru/v2
    - https://tproger.ru/translations/beginner-git-cheatsheet/
    - https://tproger.ru/translations/beginner-git-cheatsheet/
- бд, 
    - писать простые запросы - https://postgrespro.ru/education/books/sqlprimer
    - курс - https://stepik.org/course/63054/promo
    - что такое транзакция принципы ACID
- http протокол
- grpc протокол
- jira - 
    - что такое и зачем - https://timeweb.com/ru/community/articles/kak-rabotat-v-jira
    - настроить борду со своими задачами - https://support.atlassian.com/jira-software-cloud/docs/create-a-board/
- представление что такое очередь сообщений
- linux базовые команды (cd, ls, cat, grep, tail, top) crontab, переменные окружения
    - переменные окружения - https://losst.ru/peremennye-okruzheniya-v-linux
    - crontab - выполнение задач по расписанию - https://tproger.ru/translations/guide-to-cron-jobs/
    - curl - отправить http (и не только) запрос https://losst.ru/kak-polzovatsya-curl#%D0%A7%D1%82%D0%BE_%D1%82%D0%B0%D0%BA%D0%BE%D0%B5_curl
    ```curl -i -XPOST 'http://<some-url>' -d '{<somme-data>}' -H 'header:<some header>'```
        опция -i выводит заголовки ответа, -d - данные, -H заголовок, -X - метод.


### DEV
- базовое знание go
    - тур по go https://tour.golang.org/welcome/1
    - https://golang.org/doc/code
    - go mod, шпаргалка - https://encore.dev/guide/go.mod
    - курс - https://www.coursera.org/learn/golang-webservices-1/home/info
    - курс от озона - https://learning.ozon.ru/319/lp/18-programmy-obucheniya/3521-ozon-golang-school-2020?from=continue-learning
    - контекст - https://betterprogramming.pub/understanding-context-in-golang-7f574d9d94e0
### QA
- пирамида тестирования
    - https://tproger.ru/articles/integracionnye-testy-v-mikroservisah/
- основы питона 
    - курс - https://stepik.org/course/67/promo
    - еще один курс - https://stepik.org/course/512/info
    - декоратор что такое - https://habr.com/ru/post/141411/
    - генератор 
    - управление зависимостями requirements.txt, pip - менеджер пакетов
        - PyCharm и requirements.txt - https://www.jetbrains.com/help/pycharm/managing-dependencies.html#apply_dependencies
        - https://habr.com/ru/post/491916/
- pytest фраемворк для тестирования
    - https://docs.pytest.org/en/6.2.x/contents.html
...


--------
# Middle
- метрики, что такое зачем, как собирается, как пользоваться, prometheus, PromQL, grafana 
- искать логи, пользоваться graylog-ом, jager-ом, логи пода, логи на базы
- kafka
    - книжка - https://github.com/jitendra3109/ApacheKafka/blob/master/Docs/confluent-kafka-definitive-guide-complete.pdf
    - гарантии доставки
    - kafka vs RabbitMQ
- Кеш. в частности - Redis

### DEV
- linux, bash скрипты, pipe, ...
- уверенное знание go, concurrency
    - три патерна concurrency Generator, Futures, Fan-in Fan-out- https://medium.com/@thejasbabu/concurrency-patterns-golang-5c5e1bcd0833
- блокировки в бд, какие бывают, как можно получить блокировку, что такое 
    - https://postgrespro.ru/docs/postgresql/11/explicit-locking
    - alter может блочить таблицу - https://postgrespro.ru/docs/postgresql/11/sql-altertable
- отптимизация SQL запросов
    - explain analyze - https://postgrespro.ru/docs/postgresql/10/using-explain#USING-EXPLAIN-ANALYZE
    - индексы, простенько про индексы - https://highload.today/indeksy-v-mysql/
    - книга - Ганс-Юрген Шёниг **PostgreSQL 11 Мастерство разработки.**
- ? патерны, метрики
- ? профилирование
### QA
- книга > Роман Савин teстирование dot.COM
- уверенное знание питона, создание автотестов
- теория тестирования

--------
# Senior 
- находить проблемы производительности базы данных
    - pg_stat_statements
- понимать/представлять что такое шардирование, репликация
- ? разбираться в тонкостях работы бд, как она работает под капотом, разные виды индексов когда и зачем,

-------
# Team lead
? быть хорошим человеком)
- решать управленческие задачи: резать скоп, находить ресурсы, направлять заказчиков, ...

# Сылки на роадмапы по лучше
https://vladislaveremeev.github.io/#h.aazzr5roamr2

