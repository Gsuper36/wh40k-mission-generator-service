create sequence objective_id_seq;

create table objective (
    id bigint NOT NULL PRIMARY KEY DEFAULT nextval('objective_id_seq'),
    title text NOT NULL,
    description text,
    rules text
);

alter sequence objective_id_seq
owned by objective.id;