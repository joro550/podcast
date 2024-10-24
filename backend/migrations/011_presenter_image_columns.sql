alter table presenters
    add column image_url varchar(100) NOT NULL DEFAULT '',
    add column alt_text varchar(100) NOT NULL DEFAULT '',
    add column sha varchar(255) NOT NULL DEFAULT ''
