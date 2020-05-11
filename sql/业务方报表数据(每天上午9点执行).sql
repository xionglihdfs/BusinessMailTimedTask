select 
  create_time '创建时间',
  name '姓名',
  age '年龄',
  email '邮箱'
 from test
 where
    -- 当天0:00到上午9点的数据(每天上午9点执行)
	create_time BETWEEN SUBDATE(DATE_FORMAT(NOW(),'%Y-%m-%d'),interval 0 minute)
	and SUBDATE(DATE_FORMAT(NOW(),'%Y-%m-%d'),interval -540 minute);