DELIMITER  $$
# AddColumnIfNotExists
DROP
  PROCEDURE  IF  EXISTS  AddColumnIfNotExists$$
 CREATE  PROCEDURE  `AddColumnIfNotExists` (
    tableName  varchar( 100), columnName  varchar( 100),
    dbType  varchar( 100)  ) BEGIN
 DECLARE  _tableCount  INT;
DECLARE  _columnCount  INT;

SET
  _tableCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. TABLES
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tableName);

SET
  _columnCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. COLUMNS
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tableName
AND  COLUMN_NAME  = columnName);
IF  _tableCount  = 1
AND  _columnCount  = 0
 THEN
SET
  @_sqlText = CONCAT( ' ALTER TABLE `',
  tableName, '` ADD COLUMN `',
  columnName, '` ',
  dbType, ' NULL;
');
PREPARE  stmt1
FROM
  @_sqlText;
EXECUTE  stmt1;
DEALLOCATE  PREPARE  stmt1;
END  IF;
END$$

# CreateIndexIfNotExists
DROP
  PROCEDURE  IF  EXISTS  CreateIndexIfNotExists$$
 CREATE  PROCEDURE  `CreateIndexIfNotExists` (
    tableName  varchar( 100), columnName  varchar( 100)  ) BEGIN
 DECLARE  _tableCount  INT;
DECLARE  _indexCount  INT;

SET
  _tableCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. TABLES
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tableName);

SET
  _indexCount  = (
SELECT
  COUNT( 1)
FROM
  information_schema. statistics
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tableName
AND  INDEX_NAME  = CONCAT( 'IX_',
columnName)
);
IF  _tableCount  = 1
AND  _indexCount  = 0
 THEN
SET
  @_sqlText = CONCAT( ' CREATE INDEX `IX_',
  columnName, '` ON `',
  tableName, '`(`',
  columnName, '` ASC);'
); PREPARE  stmt1
FROM
  @_sqlText; EXECUTE  stmt1; DEALLOCATE  PREPARE  stmt1; END  IF; END$$

# CreateIndexIfNotExistsWithColumns
DROP
  PROCEDURE  IF  EXISTS  CreateIndexIfNotExistsWithColumns$$
 CREATE  PROCEDURE  `CreateIndexIfNotExistsWithColumns` (
    IN  tableName  varchar( 200), IN  columnName  VARCHAR( 200)  ) BEGIN
 DECLARE  _tableCount  INT;
DECLARE  _indexCount  INT;

SET
  _tableCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. TABLES
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tableName);

SET
  _indexCount  = (
SELECT
  COUNT( 1)
FROM
  information_schema. statistics
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tableName
AND  INDEX_NAME  = CONCAT( 'IX_',
columnName)
);
IF  _tableCount  = 1
AND  _indexCount  = 0
 THEN
SET
  @_sqlText = CONCAT( ' CREATE INDEX `IX_',
  columnName, '` ON `',
  tableName, '`(',
  columnName, ');
');
PREPARE  stmt1
FROM
  @_sqlText;
EXECUTE  stmt1;
DEALLOCATE  PREPARE  stmt1;
END  IF;
END$$

# CreateUniqueIndexIfNotExistsWithColumns
DROP
  PROCEDURE  IF  EXISTS  CreateUniqueIndexIfNotExistsWithColumns$$
 CREATE  PROCEDURE  `CreateUniqueIndexIfNotExistsWithColumns` (
    IN  tableName  VARCHAR( 200), IN  indexName  VARCHAR( 200),
    IN  columnName  VARCHAR( 200)  ) BEGIN
 DECLARE  _tableCount  INT;
DECLARE  _indexCount  INT;

SET
  _tableCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. TABLES
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tableName);

SET
  _indexCount  = (
SELECT
  COUNT( 1)
FROM
  information_schema. statistics
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tableName
AND  INDEX_NAME  = CONCAT( 'IX_',
indexName)
);
IF  _tableCount  = 1
AND  _indexCount  = 0
 THEN
SET
  @_sqlText = CONCAT( ' CREATE UNIQUE INDEX `IX_',
  indexName, '` ON `',
  tableName, '`(',
  columnName, ');
');
PREPARE  stmt1
FROM
  @_sqlText;
EXECUTE  stmt1;
DEALLOCATE  PREPARE  stmt1;
END  IF;
END$$

# DropColumnIfExists
DROP
  PROCEDURE  IF  EXISTS  DropColumnIfExists$$
 CREATE  PROCEDURE  `DropColumnIfExists` (
    tableName  varchar( 100), columnName  varchar( 100)  ) BEGIN
 DECLARE  _count  INT;

SET
  _count  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. COLUMNS
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tableName
AND  COLUMN_NAME  = columnName);
IF  _count  = 1
 THEN
SET
  @_sqlText = CONCAT( ' ALTER TABLE ',
  tableName, ' DROP COLUMN ',
  columnName, ' ;'
);
PREPARE  stmt1
FROM
  @_sqlText;
EXECUTE  stmt1;
DEALLOCATE  PREPARE  stmt1;
END  IF;
END$$

# DropIndexIfExists
DROP
  PROCEDURE  IF  EXISTS  DropIndexIfExists$$
 CREATE  PROCEDURE  `DropIndexIfExists` (
    tableName  varchar( 100), columnName  varchar( 100)  ) BEGIN
 DECLARE  _count  INT;

SET
  _count  = (
SELECT
  COUNT( 1)
FROM
  information_schema. statistics
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tableName
AND  INDEX_NAME  = CONCAT( 'IX_',
columnName)
);
IF  _count  > 0
 THEN
SET
  @_sqlText = CONCAT( ' DROP INDEX `IX_',
  columnName, '` ON `',
  tableName, '`; '
); PREPARE  stmt1
FROM
  @_sqlText; EXECUTE  stmt1; DEALLOCATE  PREPARE  stmt1; END  IF; END$$

# add_col_if_not_exists
DROP
  PROCEDURE  IF  EXISTS  add_col_if_not_exists$$
 CREATE  PROCEDURE  `add_col_if_not_exists` (
    IN  tbl_name  varchar( 100), IN  col_name  varchar( 100),
    IN  col_definition  varchar( 1000)  ) SQL  SECURITY  INVOKER
 BEGIN
 DECLARE  _tableCount  INT;
DECLARE  _columnCount  INT;

SET
  _tableCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. TABLES
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tbl_name);

SET
  _columnCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. COLUMNS
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tbl_name
AND  COLUMN_NAME  = col_name);
IF  _tableCount  = 1
AND  _columnCount  = 0
 THEN
SET
  @_sqlText = CONCAT( ' ALTER TABLE `',
  tbl_name, '` ADD COLUMN `',
  col_name, '` ',
  col_definition, ';
');
PREPARE  stmt1
FROM
  @_sqlText;
EXECUTE  stmt1;
DEALLOCATE  PREPARE  stmt1;
END  IF;
END$$

# add_index_if_not_exists
DROP
  PROCEDURE  IF  EXISTS  add_index_if_not_exists$$
 CREATE  PROCEDURE  `add_index_if_not_exists` (
    IN  tbl_name  varchar( 64), IN  ix_name  varchar( 64),
    IN  ix_field_names  varchar( 500)  ) SQL  SECURITY  INVOKER
 BEGIN
 DECLARE  _tableCount  INT;
DECLARE  _indexCount  INT;

SET
  _tableCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. TABLES
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tbl_name);

SET
  _indexCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. STATISTICS
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tbl_name
AND  INDEX_NAME  = ix_name);
IF  _tableCount  = 1
AND  _indexCount  = 0
 THEN
SET
  @_sqlText = CONCAT( ' ALTER TABLE `',
  tbl_name, '` ADD INDEX `',
  ix_name, '` (',
  ix_field_names, ');
');
PREPARE  stmt1
FROM
  @_sqlText;
EXECUTE  stmt1;
DEALLOCATE  PREPARE  stmt1;
END  IF;
END$$

# add_unique_if_not_exists
DROP
  PROCEDURE  IF  EXISTS  add_unique_if_not_exists$$
 CREATE  PROCEDURE  `add_unique_if_not_exists` (
    IN  tbl_name  varchar( 64), IN  uk_name  varchar( 64),
    IN  uk_field_names  varchar( 500)  ) BEGIN
 DECLARE  _tableCount  INT;
DECLARE  _indexCount  INT;

SET
  _tableCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. TABLES
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tbl_name);

SET
  _indexCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. STATISTICS
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tbl_name
AND  INDEX_NAME  = uk_name);
IF  _tableCount  = 1
AND  _indexCount  = 0
 THEN
SET
  @_sqlText = CONCAT( ' ALTER TABLE `',
  tbl_name, '` ADD UNIQUE `',
  uk_name, '` (',
  uk_field_names, ');
');
PREPARE  stmt1
FROM
  @_sqlText;
EXECUTE  stmt1;
DEALLOCATE  PREPARE  stmt1;
END  IF;
END$$

# change_col
DROP
  PROCEDURE  IF  EXISTS  change_col$$
 CREATE  PROCEDURE  `change_col` (
    IN  tbl_name  varchar( 100), IN  old_col_name  varchar( 100),
    IN  new_col_name  varchar( 100), IN  col_definition  varchar( 1000)  ) SQL  SECURITY  INVOKER
 BEGIN
SET
  @_sqlText = CONCAT( ' ALTER TABLE `',
  tbl_name, '` CHANGE COLUMN `',
  old_col_name, '` `',
  new_col_name, '` ',
  col_definition, ';
');
PREPARE  stmt1
FROM
  @_sqlText;
EXECUTE  stmt1;
DEALLOCATE  PREPARE  stmt1;
END$$

# change_col_if_exists
DROP
  PROCEDURE  IF  EXISTS  change_col_if_exists$$
 CREATE  PROCEDURE  `change_col_if_exists` (
    IN  tbl_name  varchar( 100), IN  old_col_name  varchar( 100),
    IN  new_col_name  varchar( 100), IN  col_definition  varchar( 1000)  ) SQL  SECURITY  INVOKER
 BEGIN
 DECLARE  _count  INT;

SET
  _count  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. COLUMNS
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tbl_name
AND  COLUMN_NAME  = old_col_name);
IF  _count  = 1
 THEN
SET
  @_sqlText = CONCAT( ' ALTER TABLE `',
  tbl_name, '` CHANGE COLUMN `',
  old_col_name, '` `',
  new_col_name, '` ',
  col_definition, ';
');
PREPARE  stmt1
FROM
  @_sqlText;
EXECUTE  stmt1;
DEALLOCATE  PREPARE  stmt1;
END  IF;
END$$

# drop_col_if_exists
DROP
  PROCEDURE  IF  EXISTS  drop_col_if_exists$$
 CREATE  PROCEDURE  `drop_col_if_exists` (
    IN  tbl_name  varchar( 100), IN  col_name  varchar( 100)  ) SQL  SECURITY  INVOKER
 BEGIN
 DECLARE  _count  INT;

SET
  _count  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. COLUMNS
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
AND  TABLE_NAME  = tbl_name
AND  COLUMN_NAME  = col_name);
IF  _count  = 1
 THEN
SET
  @_sqlText = CONCAT( ' ALTER TABLE ',
  tbl_name, ' DROP COLUMN ',
  col_name, ' ;'
);
PREPARE  stmt1
FROM
  @_sqlText;
EXECUTE  stmt1;
DEALLOCATE  PREPARE  stmt1;
END  IF;
END$$

# drop_index_if_exists
DROP
  PROCEDURE  IF  EXISTS  drop_index_if_exists$$
 CREATE  PROCEDURE  `drop_index_if_exists` (
    IN  tbl_name  varchar( 200), IN  idx_name  varchar( 200)  ) SQL  SECURITY  INVOKER
 BEGIN
 DECLARE  _indexCount  INT;

SET
  _indexCount  = (
SELECT
  COUNT( 1)
FROM
  INFORMATION_SCHEMA. statistics
WHERE
  TABLE_NAME  = tbl_name
  AND  INDEX_NAME  = idx_name
  AND  TABLE_SCHEMA  = (
SELECT
  SCHEMA(
)
)
);
IF  _indexCount  > 0
 THEN
SET
  @str = CONCAT( ' drop index ',
  idx_name, ' on ',
  tbl_name);
PREPARE  stmt1
FROM
  @str;
EXECUTE  stmt1;
DEALLOCATE  PREPARE  stmt1;
END  IF;
END$$

# UpdateColumnIfExists
DROP
  PROCEDURE  IF  EXISTS  UpdateColumnIfExists$$
 CREATE  PROCEDURE  `UpdateColumnIfExists` (
    IN  tableName  varchar( 100) CHARSET  utf8mb4  collate  utf8mb4_general_ci,
    IN  columnName  varchar( 100) CHARSET  utf8mb4  collate  utf8mb4_general_ci,
    IN  dbType  varchar( 100) CHARSET  utf8mb4  collate  utf8mb4_general_ci  ) SQL  SECURITY  INVOKER
 BEGIN
 DECLARE
 _count  INT;
DECLARE
 _sqlText  text;

SET
  _count  = (
SELECT
  COUNT( 1
)
FROM
  INFORMATION_SCHEMA. COLUMNS
WHERE
  TABLE_SCHEMA  = (
SELECT
  SCHEMA  (
)
  )
  AND  TABLE_NAME  = tableName
  AND  COLUMN_NAME  = columnName

);
IF
 _count  = 1  THEN
SET
  @_sqlText = CONCAT( ' ALTER TABLE `',
  tableName, '` MODIFY COLUMN `',
  columnName, '` ',
  dbType
); PREPARE  stmt1
FROM
  @_sqlText; EXECUTE  stmt1; DEALLOCATE  PREPARE  stmt1; END  IF; END$$
 DELIMITER ;