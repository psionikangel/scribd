
create table run(
  id varchar(36) PRIMARY KEY,
  starttime date,
  endtime date,
  machinename text
);

create table metadata (
  id serial,
  filepath text,
  lastmodified date,
  checksum varchar(256),
  filename text,
  filesize int,
  extension text,
  runid varchar(36) references run(id)
);

commit;
