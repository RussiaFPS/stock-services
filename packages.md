## Аргументация выбора пакетов в go.mod

> github.com/joho/godotenv -
Она позволяет устанавливать необходимые для приложения переменные окружения из файла, такую как 
информацию для подключения к базе данных. Главный ее плюс это простота, быстрая инициализация переменных. 

> github.com/gorilla/rpc -
это основной пакет для служб RPC, обеспечивающий доступ к экспортированным методам объекта через HTTP-запросы. Позволяет
быстро развернуть rpc сервер.

> github.com/jackc/pgx/v5 - database/sql поддерживает разные РСУБД, а pgx — только PostgreSQL.
За счет этого pgx может обеспечить большую производительность, и предложить больше 
возможностей, специфичных именно для постгреса. 


