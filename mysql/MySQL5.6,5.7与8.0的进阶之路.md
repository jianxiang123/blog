## MySQL5.6,5.7与8.0的进阶之路

### MySQL5.6新特性

1. **支持GTID复制**
2. **支持延迟复制(备用库)**
3. **支持基于库级别的并行复制(一个实例存在多个schema,对从库复制提升有帮助)**
4. **mysqlbinlog命令支持远程备份binlog**(很关键的一个功能,多集群备份)
5. 对TIME, DATETIME和TIMESTAMP进行了重构，可支持小数秒。DATETIME的空间需求从8byte减少到5byte
6. **支持Online DDL**。ALTER操作不再阻塞DML。(三种方式copy online,inplace online,inplace but offline) **小数据量级可以,大表DDL操作会导致从库严重延时**
7. 支持可传输表空间(transportable tablespaces),`alter table t1 import tablespace`
8. 支持统计信息的持久化。避免主从之间或数据库重启后，同一个SQL的执行计划有差异
   1. 统计信息解释:MySQL通过采样，统计出来表和索引相关信息，例如，表的记录数、聚集索引page个数
   2. 用处：生成执行计划时,需要根据索引的统计信息进行估算,计算出最低代价（或者说是最小开销）的执行计划
   3. 统计信息参数:`show variables like 'innodb_stats%';`
   4. 存在哪里?:
      1. `持久化数据：mysql.innodb_index_stats和mysql.innodb_table_stats`
      2. 非持久化数据(存在内存中):存在内存表MEMORY Table 即`INFORMATION_SCHEMA.TABLES、INFORMATION_SCHEMA.STATISTICS、INNODB_INDEXES`
9. 支持支持全文索引(基本用不到,全文存储es更合适)
10. 支持InnoDB Memcached plugin(内存数据库,用不到 有 redis)
11. EXPLAIN可用来查看DELETE，INSERT，REPLACE，UPDATE等DML操作的执行计划，在此之前，只支持SELECT操作
12. 分区表的增强，包括最大可用分区数增加至8192，支持分区和非分区表之间的数据交换，操作时显式指定分区
13. Redo Log总大小的限制从之前的4G扩展至512G
14. **Undo Log可保存在独立表空间中，因其是随机IO，更适合放到SSD中。但仍然不支持空间的自动回收**
15. **支持在线和关闭MySQL时dump和load Buffer pool的状态，避免数据库重启后需要较长的预热时间**
16. InnoDB内部的性能提升，包括拆分kernel mutex，引入独立的刷新线程，可设置多个purge线程
17. 优化器性能提升，引入了ICP，MRR，BKA等特性，针对子查询进行了优化

### MySQL5.7新特性

1. **支持组复制和InnoDB Cluster(MGR)**
2. 支持多源复制(业务场景较少，传统企业还有)
3. **支持增强半同步（AFTER_SYNC)即无损复制**
4. **支持基于表级别(LOGICAL_CLOCK)的并行复制**
5. **支持在线开启GTID复制**
6. 支持在线设置复制过滤规则
7. **支持在线修改Buffer pool的大小**
8. 支持在同一长度编码字节内，修改VARCHAR的大小只需修改表的元数据，无需创建临时表
9. **支持可设置NUMA架构的内存分配策略(innodb_numa_interleave)**
10. 支持透明页压缩(Transparent Page Compression)
11. **支持UNDO表空间的自动回收**
12. 支持查询优化器的增强和重构
13. 可查看当前session正在执行的SQL的执行计划(EXPLAIN FOR CONNECTION)
14. 引入了查询改写插件（Query Rewrite Plugin），可在服务端对查询进行改写
15. **EXPLAIN FORMAT=JSON会显示成本信息，这样可直观的比较两种执行计划的优劣**
16. 引入了虚拟列，类似于Oracle中的函数索引(定制场景使用)
17. 新实例不再默认创建test数据库及匿名用户
18. **引入ALTER USER命令，可用来修改用户密码，密码的过期策略，及锁定用户等**
19. **mysql.user表中存储密码的字段从password修改为authentication_string**
20. 支持表空间加密
21. 优化了Performance Schema，其内存使用减少
22. Performance Schema引入了众多instrumentation。常用的有Memory usage instrumentation，可用来查看MySQL的内存使用情况，Metadata Locking Instrumentation， 可用来查看MDL的持有情况，Stage Progress instrumentation，可用来查看Online DDL的进度
23. 同一触发事件（INSERT，DELETE，UPDATE），同一触发时间（BEFORE，AFTER），允许创建多个触发器。在此之前，只允许创建一个触发器
24. InnoDB原生支持分区表，在此之前，是通过ha_partition接口来实现的
25. 分区表支持可传输表空间特性。
26. **集成了SYS数据库，简化了MySQL的管理及异常问题的定位**
27. 原生支持JSON类型，并引入了众多JSON函数
28. 引入了新的逻辑备份工具mysqlpump，支持表级别的多线程备份
29. 引入了新的客户端工具mysqlsh，其支持三种语言：JavaScript, Python and SQL。两种API：X DevAPI，AdminAPI，其中，前者可将MySQL作为文档型数据库进行操作，后者用于管理InnoDB Cluster
30. mysql_install_db被mysqld --initialize代替，用来进行实例的初始化
31. 原生支持systemd
32. 引入了super_read_only选项
33. **可设置SELECT操作的超时时长（max_execution_time）**
34. **可通过SHUTDOWN命令关闭MySQL实例**  之前是mysqladmin shutdown（调用mysql_shutdown()API）
35. **引入了innodb_deadlock_detect选项**，在高并发场景下，可使用该选项来关闭死锁检测
36. 引入了Optimizer Hints，可在语句级别控制优化器的行为，如是否开启ICP，MRR等，在此之前，只有Index Hints
37. GIS的增强，包括使用Boost.Geometry替代之前的GIS算法，InnoDB开始支持空间索引

### MySQL8.0新特性

1. **引入了原生的，基于InnoDB的数据字典。**数据字典表位于mysql库中，对用户不可见，同mysql库的其它系统表一样，保存在数据目录下的mysql.ibd文件中。不再置于mysql目录下
2. 重构了INFORMATION_SCHEMA，其中部分表已重构为基于数据字典的视图，在此之前，其为临时表
3. PERFORMANCE_SCHEMA查询性能提升，其已内置多个索引
4. **InnoDB存储引擎支持原子DDL**
5. **支持不可见索引(Invisible index)**
6. **支持降序索引**
7. **优化器加入了直方图功能，对比Oracle**
8. 支持公用表表达式(Common table expressions)
9. **支持窗口函数(Window functions)。**
10. 支持角色（Role）功能，对比Oracle
11. 支持资源组（Resource Groups），可用来控制线程的优先级及其能使用的资源，目前，能被管理的资源只有CPU
12. 引入了innodb_dedicated_server选项，可基于服务器的内存来动态设置innodb_buffer_pool_size，innodb_log_file_size和innodb_flush_method
13. **支持秒加字段(Instant add column)功能**
14. JSON字段的部分更新（JSON Partial Updates）
15. 支持自增主键的持久化
16. **支持可持久化全局变量**（SET PERSIST）
17. **默认字符集由latin1修改为utf8mb4**
18. **默认开启UNDO表空间**，且支持在线调整数量（innodb_undo_tablespaces）。在MySQL 5.7中，默认不开启，若要开启，只能初始化时设置
19. 支持备份锁
20. Redo Log的优化，包括允许多个用户线程并发写入log buffer，可动态修改innodb_log_buffer_size的大小
21. **默认的认证插件由mysql_native_password更改为caching_sha2_password**
22. 默认的内存临时表由MEMORY引擎更改为TempTable引擎，相比于前者，后者支持以变长方式存储VARCHAR，VARBINARY等变长字段。从MySQL 8.0.13开始，TempTable引擎支持BLOB字段
23. **Grant不再隐式创建用户** `grant all privileges on test.* to joe@localhost identified by '123456';` 失效
24. SELECT ... FOR SHARE和SELECT ... FOR UPDATE语句中引入NOWAIT和SKIP LOCKED选项，解决电商场景热点行问题
25. 正则表达式的增强，新增了4个相关函数，REGEXP_INSTR()，REGEXP_LIKE()，REGEXP_REPLACE()，REGEXP_SUBSTR()
26. **查询优化器在制定执行计划时，判定数据是否在Buffer Pool中**`。`之前版本，假设数据都在磁盘中
27. ha_partition接口从代码层移除，如果要使用分区表，只能使用InnoDB存储引擎
28. 引入了更多细粒度的权限来替代SUPER权限，现在授予SUPER权限会提示warning
29. **GROUP BY语句不再隐式排序**
30. **information_schema中的innodb_locks和innodb_lock_waits表被移除**，取而代之的是performance_schema中的data_locks和data_lock_waits表
31. **引入performance_schema.variables_info表**，记录了参数的来源及修改情况
32. 增加了对于客户端报错信息的统计（performance_schema.events_errors_summary_xxx）
33. 可统计查询的响应时间分布（call sys.ps_statement_avg_latency_histogram()）
34. **支持直接修改列名**（ALTER TABLE ... RENAME COLUMN old_name TO new_name）
35. 用户密码可设置重试策略（Reuse Policy）
36. **移除PASSWORD()函数**。这就意味着无法通过“SET PASSWORD ... =PASSWORD('auth_string') ”命令修改用户密码
37. **代码层移除Query Cache模块**，故Query Cache相关的变量和操作均不再支持
38. BLOB, TEXT, GEOMETRY和JSON字段允许设置默认值
39. **可通过RESTART命令重启MySQL实例**

