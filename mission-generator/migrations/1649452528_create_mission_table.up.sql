create sequence mission_id_seq;

create table mission (
    id bigint PRIMARY KEY NOT NULL DEFAULT nextval('mission_id_seq'),
    title text NOT NULL,
    description text,
    rules text NOT NULL,
    format int NOT NULL,
    deployment_id int NOT NULL
);

create index on mission using btree (deployment_id);

alter sequence mission_id_seq
owned by mission.id;