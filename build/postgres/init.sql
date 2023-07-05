create table costumer(
    id serial primary key,
    name varchar(80),
    surname varchar(80),
    doc_number varchar(30),
    birthdate date
);

insert into costumer (name, surname, doc_number, birthdate) values ('Steven', 'Rogers', '1919234692311', '1918-07-04');
insert into costumer (name, surname, doc_number, birthdate) values ('Diana', 'Prince', '1344414123124', '1941-10-01');
