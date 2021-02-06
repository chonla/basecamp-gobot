CREATE TABLE phonebook (
    user_id bigint,
    phone_number varchar(50) not null
);

CREATE INDEX phonebook_idx ON phonebook (
    user_id
);