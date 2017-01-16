
create table run(
  id serial,
  runid varchar(100),
  starttime timestamp,
  endtime timestamp,
  machinename text
);

create table metadata (
  id serial,
  filepath text,
  lastmodified timestamp,
  checksum varchar(256),
  filename text,
  filesize int,
  extension text,
  runid varchar(100)
);

commit;
