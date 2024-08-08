CREATE TABLE articles(
    id SERIAL PRIMARY KEY, 
    title TEXT,
    description TEXT,
    url TEXT,
    source TEXT,
    imageUrl TEXT,
    category 
    language TEXT,
    country TEXT,
    published DATE,    
);