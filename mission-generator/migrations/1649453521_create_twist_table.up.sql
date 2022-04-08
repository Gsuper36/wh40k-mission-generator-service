create sequence twist_id_seq;

create table twist (
    id bigint NOT NULL PRIMARY KEY DEFAULT nextval('twist_id_seq'),
    title text NOT NULL,
    description text,
    rules text NOT NULL
);

alter sequence twist_id_seq
owned by twist.id;