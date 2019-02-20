SET client_encoding = 'UTF8';

-- 触发器

-- 操作事项记录
create table if not exists operation (
    id BIGSERIAL primary key, 
    user_id int not null default 0,
    controller varchar(50) not null default '',
    action varchar(50) not null default '',
    create_time timestamp without time zone not null,
    description text not null default ''
);
---- 用户登录
--create table if not exists user (
--    id serial primary key,
--    user_name varchar(50) not null default '',
--    password char(32) not null default '',
--    salt varchar(10) not null default '',
--    status int not null default 0,
--    create_time timestamp without time zone not null,
--    last_login timestamp without time zone not null,
--    update_time timestamp without time zone not null,
--    email varchar(50) not null default '',
--    last_ip char(16)  not null default '',
--    sex int not null default 0,
--    CHECK (email ~* '^\w+@\w+[.]\w+$')
--);
---- 角色表
--create table if not exists role (
--    id serial primary key,
--    role_name varchar(50) not null default '',
--    description varchar(255) not null default '',
--    create_time timestamp without time zone not null,
--    update_time timestamp without time zone not null 
--);
--
--create table if not exists user_role (
--    id serial primary key,
--    user_id int not null default 0 REFERENCES user(id),
--    role_id int not null default 0 REFERENCES role(id)
--)
--
--create table if not exists perm (
--    id serial primary key,
--    module varchar(50) not null default '',
--    action varchar(50) not null default ''
--);
--
--create table if not exists role_perm (
--    id serial primary key,
--    role_id int not null default 0 REFERENCES role(id),
--    perm_id int not null default 0 REFERENCES perm(id)
--);
-- 审核表
create table if not exists review (
    id serial primary key,
    checker_id int not null default 0,   -- 审批人
    applicant_id int not null default 0, -- 申请人
    item_id int not null default 0, -- 审核事项
    type int not null default 0, -- 合同，充值等审核
    status int not null default 0, 
    description text default ''
);

create table if not exists customer (
    id serial primary key,
    company_name varchar(50) not null default '',
    company_type int not null default 0, -- 个人独资；合伙企业；股份公司；国有企业；外资企业；合资企业
    company_type_desc varchar(50) default '', -- 如：非上市
    register_fund decimal(10,2) not null default 0.00, -- 万元为单位
    license_type int not null default 0, -- 社会统一信用代码证;营业执照
    license_no varchar(50) not null default '', -- 号码
    license_valid_time timestamp without time zone not null, 
    main_business varchar(255) not null default '',
    phone_usedFor varchar(255) not null default '', -- 400电话用途
    register_address varchar(255) not null default '', -- 注册地址
    postcode int not null default 0, 
    company_reprensent varchar(50) not null default '', -- 法人
    company_reprensent_id char(18) not null default '', -- 法人身份证
    is_operator int not null default 0, -- 是否有经办人
    operator varchar(50) not null default '', -- 经办人姓名
    operator_id char(18) not null default '', -- 经办人身份证
    represent_id_valid_date timestamp without time zone not null,  -- 法人身份证到期时间
    operator_id_valid_date timestamp without time zone not null, -- 经办人身份证到期时间
    operator_mobile char(11) not null default '',
    operator_telephone varchar(20) not null default '',
    operator_email varchar(50) not null default '',
    account_type int not null default 0, -- 公司账户，法人账户
    account_bank varchar(255) not null default '', -- 开户行
    account_no varchar(50) not null default '', -- 银行账户号码
    kefu_id int not null default 0, 
    shangwu_id int not null default 0, 
    qudao_id int not null default 0, 
    qq varchar(20) not null default '',
    company_website varchar(50) not null default '',
    fax varchar(50) not null default '',
    fund decimal(10,2) not null default 0.00, -- 充值到账金额
    status int not null default 0,
    description text not null default ''
);

create table if not exists customer_uploads (
    id serial primary key, 
    customer_id int not null default 0,
    represent_id_img varchar(50) not null default '',
    operator_id_img varchar(50) not null default '', 
    license_img varchar(50) not null default '', 
    shoulidan_img varchar(50) not null default '', 
    service_procotol varchar(50) not null default '', 
    safe_duty_promise varchar(50) not null default '',
    operator_authorize varchar(50) not null default '', -- 经办人授权书
    open_account_license varchar(50) not null default '',
    contract varchar(50) not null default '',
    ring_require varchar(50) not null default '',
    more varchar(255) not null default '',
    UNIQUE(customer_id)
);

create table if not exists phone_400 (
    id serial primary key, 
    phone char(10) not null default '' UNIQUE,
    small_phone varchar(20) not null default '' UNIQUE,
    customer_id int not null default 0 REFERENCES customer(id),
    telFee_pkg_id int not null default 0, -- 话费套餐
    func_pkg_id int not null default 0, -- 功能套餐
    call_type int not null default 0, -- 直连 or ivr
    response_type int not null default 0, -- 顺序接听 or 随机接听
    pre_answer int not null default 0, -- 是否预摘机
    status int not null default 0, -- 开通 or 暂停 or ...
    contact_date timestamp without time zone not null, -- 联通合同日期
    valide_date timestamp without time zone not null, -- 我们合同日期
    worked_date timestamp without time zone not null, -- 生效日期
    tel_surplus decimal(10, 2) not null default 0.00, -- 话费余额
    tel_invalid_date timestamp without time zone not null, -- 电话失效日期
    func_invalid_date timestamp without time zone not null -- 功能失效日期
);

create table if not exists bind_phoner (
    id serial primary key,
    bind_phone varchar(20) not null default '',
    customer_id int not null default 0 REFERENCES customer(id),
    phone_400_number char(10) not null default '' REFERENCES phone_400(phone), 
    bind_phone_img varchar(50) not null default '',
    privilege int not null default 0, -- 优先级
    jobnum varchar(20) not null default '', -- 工号
    gateway_amount int not null default 0, -- 中继数量，默认无
    wait_time int not null default 15, -- 默认拨号时间15秒
    start_time timestamp without time zone not null,
    end_time timestamp without time zone not null
);

create table if not exists blacklist(
    id serial primary key,
    phone_400_number char(10) not null default '' REFERENCES phone_400(phone), 
    caller_number varchar(20) not null default ''
);

create table if not exists talk_package(
    id serial primary key,
    name varchar(50) not null default '',
    period int not null default 0,  -- 月份为单位
    fee decimal(10, 2) not null default 0.00,
    rate decimal(10, 2) not null default 0.00,
    min_consume decimal(10, 2) not null default 0.00,
    description text not null default ''
);


create table if not exists func_package (
    id serial primary key, 
    name varchar(50) not null default '', 
    period int not null default 0, -- 月份为单位
    type int not null default 0, -- 功能套餐类型
    fee decimal(10, 2) not null default 0.00,
    description text default ''
);

create table if not exists func_setting (
    id serial primary key, 
    phone_400_number char(10) not null default '' REFERENCES phone_400(phone), 
    func_package_id int not null default 0 REFERENCES func_package(id),
    plan_timing int not null default 0, -- 时间计划的类型，枚举如：工作日||非工作日，节假日||非节假日
    start_time timestamp without time zone not null,
    end_time timestamp without time zone not null
);

create table if not exists music (
    id BIGSERIAL primary key,
    name varchar(50) not null default '',
    path varchar(50) not null default '',
    create_time timestamp without time zone not null
);

create table if not exists call_ring (
    id serial primary key, 
    phone_400_number char(10) not null default '' REFERENCES phone_400(phone), 
    music_id int not null default 0 REFERENCES music(id),
    start_time timestamp without time zone not null,
    end_time timestamp without time zone not null,
    description text default ''
);

create table if not exists special_ring (
    id serial primary key, 
    phone_400_number char(10) not null default '' REFERENCES phone_400(phone), 
    music_id int not null default 0 REFERENCES music(id),
    start_time timestamp without time zone not null,
    end_time timestamp without time zone not null,
    description text default ''
);

create table if not exists ivr (
    id serial primary key, 
    phone_400_number char(10) not null default '' REFERENCES phone_400(phone), 
    music_id int not null default 0 REFERENCES music(id),
    ivr_level int not null default 0,
    department varchar(50) not null default '',
    key_length int not null default 1,
    start_time timestamp without time zone not null,
    end_time timestamp without time zone not null,
    description text default ''
);

create table if not exists ivr_action (
    id serial primary key, 
    ivr_id int not null default 0 REFERENCES ivr(id),
    keys int not null default 0,
    action varchar(50) not null default '',
    value varchar(50) not null default '',
    description text not null default ''
);
-- 充值记录
create table if not exists payment (
    id serial primary key, 
    customer_id int not null default 0 REFERENCES customer(id),
    fee decimal(10, 2) not null default 0.00,
    create_time timestamp without time zone not null,
    status int not null default 0,
    description text not null default ''
);

-- 消费记录
create table if not exists consume (
    id serial primary key, 
    customer_id int not null default 0 REFERENCES customer(id),
    phone_400_number char(10) not null default '' REFERENCES phone_400(phone), 
    type int not null default 0, -- 话费、IVR、彩铃等
    type_id int not null default 0, -- 消费类型所对应的套餐id
    fee decimal(10, 2) not null default 0.00,
    create_time timestamp without time zone not null,
    status int not null default 0,
    description text not null default ''
);

-- 话单
create table if not exists call_pg_cdr (
    id BIGSERIAL primary key, 
    local_ip_v4 char(16) not null default '',
    caller_id_name varchar(50) not null default '',
    caller_id_number varchar(50) not null default '',
    outbound_caller_id_number varchar(50) not null default '',
    destination_number varchar(50) not null default '',
    context varchar(50) not null default '',
    start_stamp timestamp without time zone not null,
    answer_stamp timestamp without time zone not null,
    end_stamp timestamp without time zone not null,
    duration int not null default 0,
    billsec int not null default 0,
    hangup_cause varchar(30) not null default '',
    uuid char(36) not null default '' UNIQUE,
    bleg_uuid char(36) not null default '' UNIQUE,
    account_code varchar(50) not null default '',
    read_codec varchar(10) not null default '',
    write_codec varchar(10) not null default '',
    record_file varchar(255) not null default '',
    direction varchar(50) not null default '',
    sip_hangup_disposition varchar(50) not null default '',
    origination_uuid varchar(100) not null default '',
    sip_gateway_name varchar(50) not null default '',
    rate decimal(10,2) not null default 0.00
);

-- 插入话单的时候计算扣费


-- 录音
create table if not exists call_record (
    id BIGSERIAL primary key,
    uuid char(36) not null default '' REFERENCES call_pg_cdr(uuid),
    path varchar(50) not null default ''
);

create table if not exists area_phone (
    id serial primary key,
    mobile_segment character varying(20) NOT NULL,
    district character varying(50),
    telphone_segment character varying(10)
);

create table if not exists phone_area (
    id serial primary key, 
    phone_400_number char(10) not null default '' REFERENCES phone_400(phone), 
    area_id int not null default 0
);

COPY area_phone (mobile_segment, district, telphone_segment) FROM '/tmp/result.txt';

CREATE TRIGGER insert_tbl_call_pg_cdr BEFORE INSERT ON call_pg_cdr FOR EACH ROW EXECUTE PROCEDURE auto_insert_into_call_pg_cdr('start_stamp');