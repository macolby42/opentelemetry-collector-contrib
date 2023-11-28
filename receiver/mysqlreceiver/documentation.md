[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# mysql

## Default Metrics

The following metrics are emitted by default. Each of them can be disabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: false
```

### mysql.buffer_pool.data_pages

The number of data pages in the InnoDB buffer pool.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| status | The status of buffer pool data. | Str: ``dirty``, ``clean`` |

### mysql.buffer_pool.limit

The configured size of the InnoDB buffer pool.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| By | Sum | Int | Cumulative | false |

### mysql.buffer_pool.operations

The number of operations on the InnoDB buffer pool.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| operation | The buffer pool operations types. | Str: ``read_ahead_rnd``, ``read_ahead``, ``read_ahead_evicted``, ``read_requests``, ``reads``, ``wait_free``, ``write_requests`` |

### mysql.buffer_pool.page_flushes

The number of requests to flush pages from the InnoDB buffer pool.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

### mysql.buffer_pool.pages

The number of pages in the InnoDB buffer pool.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| kind | The buffer pool pages types. | Str: ``data``, ``free``, ``misc`` |

### mysql.buffer_pool.usage

The number of bytes in the InnoDB buffer pool.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| By | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| status | The status of buffer pool data. | Str: ``dirty``, ``clean`` |

### mysql.double_writes

The number of writes to the InnoDB doublewrite buffer.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| kind | The doublewrite types. | Str: ``pages_written``, ``writes`` |

### mysql.handlers

The number of requests to various MySQL handlers.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| kind | The handler types. | Str: ``commit``, ``delete``, ``discover``, ``external_lock``, ``mrr_init``, ``prepare``, ``read_first``, ``read_key``, ``read_last``, ``read_next``, ``read_prev``, ``read_rnd``, ``read_rnd_next``, ``rollback``, ``savepoint``, ``savepoint_rollback``, ``update``, ``write`` |

### mysql.index.io.wait.count

The total count of I/O wait events for an index.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| operation | The io_waits operation type. | Str: ``delete``, ``fetch``, ``insert``, ``update`` |
| table | Table name for event or process. | Any Str |
| schema | The schema of the object. | Any Str |
| index | The name of the index. | Any Str |

### mysql.index.io.wait.time

The total time of I/O wait events for an index.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| ns | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| operation | The io_waits operation type. | Str: ``delete``, ``fetch``, ``insert``, ``update`` |
| table | Table name for event or process. | Any Str |
| schema | The schema of the object. | Any Str |
| index | The name of the index. | Any Str |

### mysql.locks

The number of MySQL locks.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| kind | The table locks type. | Str: ``immediate``, ``waited`` |

### mysql.log_operations

The number of InnoDB log operations.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| operation | The log operation types. | Str: ``waits``, ``write_requests``, ``writes`` |

### mysql.mysqlx_connections

The number of mysqlx connections.

This metric is specific for MySQL working as Document Store (X-Plugin). [more docs](https://dev.mysql.com/doc/refman/8.0/en/document-store.html)

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| status | The connection status. | Str: ``accepted``, ``closed``, ``rejected`` |

### mysql.opened_resources

The number of opened resources.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| kind | The kind of the resource. | Str: ``file``, ``table_definition``, ``table`` |

### mysql.operations

The number of InnoDB operations.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| operation | The operation types. | Str: ``fsyncs``, ``reads``, ``writes`` |

### mysql.page_operations

The number of InnoDB page operations.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| operation | The page operation types. | Str: ``created``, ``read``, ``written`` |

### mysql.prepared_statements

The number of times each type of prepared statement command has been issued.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| command | The prepare statement command types. | Str: ``execute``, ``close``, ``fetch``, ``prepare``, ``reset``, ``send_long_data`` |

### mysql.row_locks

The number of InnoDB row locks.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| kind | The row lock type. | Str: ``waits``, ``time`` |

### mysql.row_operations

The number of InnoDB row operations.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| operation | The row operation type. | Str: ``deleted``, ``inserted``, ``read``, ``updated`` |

### mysql.sorts

The number of MySQL sorts.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| kind | The sort count type. | Str: ``merge_passes``, ``range``, ``rows``, ``scan`` |

### mysql.table.io.wait.count

The total count of I/O wait events for a table.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| operation | The io_waits operation type. | Str: ``delete``, ``fetch``, ``insert``, ``update`` |
| table | Table name for event or process. | Any Str |
| schema | The schema of the object. | Any Str |

### mysql.table.io.wait.time

The total time of I/O wait events for a table.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| ns | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| operation | The io_waits operation type. | Str: ``delete``, ``fetch``, ``insert``, ``update`` |
| table | Table name for event or process. | Any Str |
| schema | The schema of the object. | Any Str |

### mysql.threads

The state of MySQL threads.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| kind | The thread count type. | Str: ``cached``, ``connected``, ``created``, ``running`` |

### mysql.tmp_resources

The number of created temporary resources.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| resource | The kind of temporary resources. | Str: ``disk_tables``, ``files``, ``tables`` |

### mysql.uptime

The number of seconds that the server has been up.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| s | Sum | Int | Cumulative | true |

## Optional Metrics

The following metrics are not emitted by default. Each of them can be enabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: true
```

### mysql.client.network.io

The number of transmitted bytes between server and clients.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| By | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| kind | The name of the transmission direction. | Str: ``received``, ``sent`` |

### mysql.commands

The number of times each type of command has been executed.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| command | The command types. | Str: ``delete``, ``insert``, ``select``, ``update`` |

### mysql.connection.count

The number of connection attempts (successful or not) to the MySQL server.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

### mysql.connection.errors

Errors that occur during the client connection process.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| error | The connection error type. | Str: ``accept``, ``internal``, ``max_connections``, ``peer_address``, ``select``, ``tcpwrap``, ``aborted``, ``aborted_clients``, ``locked`` |

### mysql.joins

The number of joins that perform table scans.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| kind | The kind of join. | Str: ``full``, ``full_range``, ``range``, ``range_check``, ``scan`` |

### mysql.mysqlx_worker_threads

The number of worker threads available.

This metric is specific for MySQL working as Document Store (X-Plugin). [more docs](https://dev.mysql.com/doc/refman/8.0/en/document-store.html)

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| kind | The worker thread count kind. | Str: ``available``, ``active`` |

### mysql.query.client.count

The number of statements executed by the server. This includes only statements sent to the server by clients.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

### mysql.query.count

The number of statements executed by the server.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

### mysql.query.slow.count

The number of slow queries.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

### mysql.replica.sql_delay

The number of seconds that the replica must lag the source.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| s | Sum | Int | Cumulative | false |

### mysql.replica.time_behind_source

This field is an indication of how “late” the replica is.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| s | Sum | Int | Cumulative | false |

### mysql.statement_event.count

Summary of current and recent statement events.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| schema | The schema of the object. | Any Str |
| digest | Digest. | Any Str |
| digest_text | Text before digestion. | Any Str |
| kind | Possible event states. | Str: ``errors``, ``warnings``, ``rows_affected``, ``rows_sent``, ``rows_examined``, ``created_tmp_disk_tables``, ``created_tmp_tables``, ``sort_merge_passes``, ``sort_rows``, ``no_index_used`` |

### mysql.statement_event.wait.time

The total wait time of the summarized timed events.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| ns | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| schema | The schema of the object. | Any Str |
| digest | Digest. | Any Str |
| digest_text | Text before digestion. | Any Str |

### mysql.table.lock_wait.read.count

The total table lock wait read events.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| schema | The schema of the object. | Any Str |
| table | Table name for event or process. | Any Str |
| kind | Read operation types. | Str: ``normal``, ``with_shared_locks``, ``high_priority``, ``no_insert``, ``external`` |

### mysql.table.lock_wait.read.time

The total table lock wait read events times.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| ns | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| schema | The schema of the object. | Any Str |
| table | Table name for event or process. | Any Str |
| kind | Read operation types. | Str: ``normal``, ``with_shared_locks``, ``high_priority``, ``no_insert``, ``external`` |

### mysql.table.lock_wait.write.count

The total table lock wait write events.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| schema | The schema of the object. | Any Str |
| table | Table name for event or process. | Any Str |
| kind | Write operation types. | Str: ``allow_write``, ``concurrent_insert``, ``low_priority``, ``normal``, ``external`` |

### mysql.table.lock_wait.write.time

The total table lock wait write events times.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| ns | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| schema | The schema of the object. | Any Str |
| table | Table name for event or process. | Any Str |
| kind | Write operation types. | Str: ``allow_write``, ``concurrent_insert``, ``low_priority``, ``normal``, ``external`` |

### mysql.table_open_cache

The number of hits, misses or overflows for open tables cache lookups.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| status | The status of cache access. | Str: ``hit``, ``miss``, ``overflow`` |

## Resource Attributes

| Name | Description | Values | Enabled |
| ---- | ----------- | ------ | ------- |
| mysql.instance.endpoint | Endpoint of the MySQL instance. | Any Str | true |