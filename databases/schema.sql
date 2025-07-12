-- 创建 Blog 表
CREATE TABLE IF NOT EXISTS blogs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content_md TEXT NOT NULL,
    content_html TEXT,
    category VARCHAR(50),
    tags TEXT[], -- pq.StringArray 映射为 PostgreSQL 的 TEXT[]
    status VARCHAR(20) DEFAULT 'draft',
    created_at BIGINT DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT DEFAULT EXTRACT(EPOCH FROM NOW())
);

-- 创建 Question 表
CREATE TABLE IF NOT EXISTS questions (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    answer TEXT NOT NULL,
    category VARCHAR(50),
    tags TEXT[], -- pq.StringArray 映射为 PostgreSQL 的 TEXT[]
    difficulty VARCHAR(20),
    created_at BIGINT DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT DEFAULT EXTRACT(EPOCH FROM NOW())
);