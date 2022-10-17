create schema if not exists flightBooking;

create table if not exists flightBooking.customers
(
    id          serial primary key,
    name        char(60) not null,
    surname     char(60) not null,
    middle_name char(60)
    );

create table if not exists flightBooking.customers
(
    id          serial primary key,
    name        char(60) not null,
    surname     char(60) not null,
    middle_name char(60)
    );

create table if not exists flightBooking.docs
(
    id       serial primary key,
    doc_type char(50) not null
    );

create table if not exists flightBooking.customer_docs
(
    id          serial primary key,
    customer_id int
    constraint customer_docs_customers_id_fk
    references flightBooking.customers (id)
    );

create table if not exists flightBooking.flights
(
    id           serial primary key,
    ticketsTotal int         not null,
    ticketsFree  int         not null check ( ticketsFree <= flights.ticketsTotal ),
    cost         money       not null,
    departureAt  timestamptz not null,
    arrivalAt    timestamptz not null /* Сюда прооверка не нужна т.к. самолет по часовому поясу может прилететь быстрее, короче в прошлое улетел*/
    );

create table if not exists flightBooking.booked_flights
(
    id          serial primary key,
    flight_id   int
    constraint booked_flights_flights_id_fk
    references flightBooking.flights (id),
    customer_id int
    constraint booked_flights_customer_id_fk
    references flightBooking.customers (id)
    )

