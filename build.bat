rem del .\pkg\* /s/q
go build -o ./bin/ngadmin.exe ./apps/ngadmin
go build -o ./bin/client.exe ./apps/client
go build -o ./bin/login.exe ./apps/ngadmin/services/login
go build -o ./bin/nest.exe ./apps/ngadmin/services/nest
go build -o ./bin/region.exe ./apps/ngadmin/services/region
go build -o ./bin/store.exe ./apps/ngadmin/services/store
go build -o ./bin/world.exe ./apps/ngadmin/services/world
