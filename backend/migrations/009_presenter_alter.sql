alter table presenters
    add column description varchar(100) NOT NULL DEFAULT '',
    add column socials JSONB NOT NULL DEFAULT '{}'::jsonb

