CREATE TABLE tasks (
  task_id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL,
  title varchar(50) NOT NULL,
  description varchar(2000),
  due_date TEXT,
  importance smallint
);

INSERT INTO tasks (
  user_id,
  title,
  description,
  due_date,
  importance
) VALUES (
  1,
  "Test task number 1",
  "This is a longer description of the test task numero uno.",
  "2015-02-01 00:00:00",
  0
);

INSERT INTO tasks (
  user_id,
  title,
  due_date,
  importance
) VALUES (
  2,
  "This is Tanis' first task",
  "2015-02-01 00:00:00",
  0
);

INSERT INTO tasks (
  user_id,
  title,
  description,
  due_date,
  importance
) VALUES (
  1,
  "Test task number 2",
  "This is a longer description of the test task numero dos.",
  "2015-02-01 00:00:00",
  0
);