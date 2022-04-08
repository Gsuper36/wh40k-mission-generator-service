create sequence deployment_id_seq;

create table deployment (
    id biginteger PRIMARY_KEY NOT NULL DEFAULT nextval('deployment_id_seq'),
    image_url text NOT NULL
);

alter sequence deployment_id_seq
owned by deployment.id;