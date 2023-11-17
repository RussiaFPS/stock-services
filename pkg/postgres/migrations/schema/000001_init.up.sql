CREATE TABLE IF NOT EXISTS Stocks (
    Name varchar(255) NOT NULL PRIMARY KEY,
    Availability boolean NOT NULL Default true
);


CREATE TABLE IF NOT EXISTS Products (
    Id varchar(36) NOT NULL PRIMARY KEY,
    Name varchar(255) NOT NULL,
    Size varchar(255)
);

CREATE TABLE IF NOT EXISTS Balance_stocks (
    ProductsId varchar(36),
    StocksId varchar(255),
    Amount int NOT NULL Default 0,
    CountReserv int NOT NULL Default 0,
    FOREIGN KEY (ProductsId)  REFERENCES Products (Id),
    FOREIGN KEY (StocksId)  REFERENCES Stocks (Name),
    CONSTRAINT ProdStock PRIMARY KEY (ProductsId,StocksId)
);

insert into stocks (name) values ('t1');

insert into products (id,name) values ('4795435a-a814-4203-8dc6-4bf093bec52e','ball');
insert into products (id,name) values ('4795435a-a814-4203-8dc6-4bf093bec52r','ked');

insert into balance_stocks (productsid, stocksid, amount) values ('4795435a-a814-4203-8dc6-4bf093bec52e','t1',10);
insert into balance_stocks (productsid, stocksid, amount) values ('4795435a-a814-4203-8dc6-4bf093bec52r','t1',17);