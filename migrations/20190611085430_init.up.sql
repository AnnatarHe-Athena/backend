CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    src INTEGER UNIQUE NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS cells (
    id SERIAL PRIMARY KEY,
    img VARCHAR(255) UNIQUE NOT NULL,
    text VARCHAR(255) NOT NULL,
    cate INTEGER REFERENCES categories (id),
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    createdBy integer REFERENCES users (id),
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- 2  -> public 共有的，谁都可以看
    -- 3  -> 受保护的，只有发布者自己能看
    -- 4+ -> 暂未定义
    premission SMALLINT NOT NULL DEFAULT 2,
    likes BIGINT NOT NULL DEFAULT 0,
    from_url VARCHAR(255) NOT NULL DEFAULT '',
    from_id VARCHAR(255) NOT NULL DEFAULT '',
    content TEXT NOT NULL DEFAULT '',
    md5 VARCHAR(255) NOT NULL DEFAULT ''
);
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(48) UNIQUE NOT NULL,
    name VARCHAR(36) UNIQUE NOT NULL,
    pwd VARCHAR(255) NOT NULL,
    avatar VARCHAR(255) NOT NULL DEFAULT '',
    bio VARCHAR(255) NOT NULL DEFAULT '',
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- 用户级别 0-19 超级管理员 20-39 一般管理员 40-59 高级用户 60-79 一般用户 80-99 受限用户
    role SMALLINT NOT NULL DEFAULT 90
);
-- a user has many collection
CREATE TABLE IF NOT EXISTS collections (
    id SERIAL PRIMARY KEY,
    cell INTEGER REFERENCES cells(id),
    owner INTEGER REFERENCES users(id)
);

-- 2017-10-15 version check for mobile platform
CREATE TABLE IF NOT EXISTS versions(
    id SERIAL PRIMARY KEY,
    platform VARCHAR(32) NOT NULL DEFAULT '',
    version INTEGER NOT NULL DEFAULT 0,
    published_by VARCHAR(32) NOT NULL DEFAULT '',
    link VARCHAR(255) NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    title VARCHAR(32) NOT NULL DEFAULT ''
);

-- 待确认
-- 2017-11-21 tags
CREATE TABLE IF NOT EXISTS tags(
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL DEFAULT '',
    desc VARCHAR(255) NOT NULL DEFAULT '',
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- begin
-- tags 和 girls 多对多的关系，用来渐渐替换掉 categories 
-- 暂时并未启用，有大块时间的时候，写个工具，做一套数据迁移
CREATE TABLE NOT EXISTS tags_girls(
    id SERIAL PRIMARY KEY,
    tag_id INTEGER REFERENCES tags(id),
    cell_id INTEGER REFERENCES cells(id)
)

--- end
