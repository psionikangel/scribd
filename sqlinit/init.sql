
create table run(
  id varchar(36),
  starttime date,
  endtime date,
  machinename text
);

create table metadata (
  filepath text,
  lastmodified date,
  checksum varchar(256),
  filename text,
  filesize int,
  extension text,
  runid varchar(36) references run(id)
);

commit;
