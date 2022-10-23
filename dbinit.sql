CREATE TABLE BALANCE(
    USERID BIGINT ,
    USERBALANCE REAL,
    UNIQUE (USERID)
);

CREATE TABLE RESERVED(
    USERID BIGINT ,
    SERVICEID BIGINT ,
    ORDERID BIGINT,
    SUM REAL,
    CREATED DATE DEFAULT NOW()::TIMESTAMP,
    UNIQUE (USERID, SERVICEID, ORDERID)
);

CREATE TABLE FINISHED(
    USERID BIGINT ,
    SERVICEID BIGINT ,
    ORDERID BIGINT,
    SUM REAL,
    UPDATED DATE DEFAULT NOW()::TIMESTAMP,
    UNIQUE (USERID, SERVICEID, ORDERID)
);

CREATE TABLE TRANSACTIONS(
    USERID BIGINT ,
    SERVICEID BIGINT DEFAULT 0,
    ORDERID BIGINT DEFAULT 0,
    SUM REAL,
    CREATED DATE DEFAULT NOW()::TIMESTAMP ,
    SOURSE VARCHAR DEFAULT 'INTERNAL',
    TYPE VARCHAR DEFAULT 'PAYMENT'
)