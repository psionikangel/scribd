
create table run(
  id serial,
  runid varchar(100),
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
  runid varchar(100) references run(runid)
);

commit;
