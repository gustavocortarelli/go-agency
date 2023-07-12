create table costumer(
    id serial primary key,
    name varchar(80),
    surname varchar(80),
    doc_number varchar(30),
    birthdate date
);

insert into costumer (name, surname, doc_number, birthdate) values ('Steven', 'Rogers', '1919234692311', '1918-07-04');
insert into costumer (name, surname, doc_number, birthdate) values ('Diana', 'Prince', '1344414123124', '1941-10-01');

create table country(
    id integer primary key,
    name varchar(100)
);

insert into country (id, name) values (1, 'Brazil'), (2, 'Portugal'), (3, 'Israel');

create table city(
    id integer primary key,
    country_id integer,
    name varchar(100),
    CONSTRAINT fk_country
        FOREIGN KEY(country_id)
            REFERENCES country(id)
);

insert into city (id, name, country_id) values
    (1, 'Brasília', 1), (2, 'São Paulo', 1), (3, 'Curitiba', 1),
    (4, 'Lisboa', 2), (5, 'Porto', 2), (6, 'Braga', 2),
    (7, 'Tel Aviv', 3), (8, 'Haifa', 3), (9, 'Rishon LeZion', 3);

create table costumer_address(
    id serial primary key,
    city_id integer not null,
    costumer_id integer not null,
    address varchar(200),
    zip_code varchar(10),
    CONSTRAINT fk_city
        FOREIGN KEY(city_id)
            REFERENCES city(id),
    CONSTRAINT fk_costumer
        FOREIGN KEY(costumer_id)
            REFERENCES costumer(id)
)

insert into costumer_address (city_id, costumer_id, address, zip_code) values (3, 1, 'Rua XV de Novembro', '80045-270');
insert into costumer_address (city_id, costumer_id, address, zip_code) values (3, 1, 'Avenida Paulista', '01311-100');
insert into costumer_address (city_id, costumer_id, address, zip_code) values (3, 2, 'Avenida da Liberdade', '4715-037');