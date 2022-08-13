
create table question
(
    id       integer default nextval('question_id_seq'::regclass) not null
        constraint question_pk
            primary key,
    name     text                                                 not null,
    response text                                                 not null,
    serie_id integer
        constraint question_serie_id_fk
            references serie
)

create table serie
(
    id   integer default nextval('serie_id_seq'::regclass) not null
        constraint serie_pk
            primary key,
    name text                                              not null
)