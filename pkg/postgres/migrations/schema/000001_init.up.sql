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