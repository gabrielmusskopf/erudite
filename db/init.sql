-- Criação da tabela Tags
CREATE TABLE Tags (
    Id SERIAL PRIMARY KEY,
    Text TEXT,
    CreationDate TIMESTAMP
);

-- Criação da tabela Questions
CREATE TABLE Questions (
    Id SERIAL PRIMARY KEY,
    Text TEXT,
    Deleted BOOLEAN,
    CreationDate TIMESTAMP
);

-- Criação da tabela Questions
CREATE TABLE QuestionsTags (
    Id SERIAL PRIMARY KEY,
    QuestionId INT,
    TagId INT,
    CreationDate TIMESTAMP
);

-- Criação da tabela Answers
CREATE TABLE Answers (
    Id SERIAL PRIMARY KEY,
    QuestionId INT,
    Text TEXT,
    Correct BOOLEAN
);

-- Adição da chave estrangeira na tabela Answers
ALTER TABLE Answers ADD CONSTRAINT fk_question_answers FOREIGN KEY (QuestionId) REFERENCES Questions(Id);

-- Adição das chaves estrangeira na tabela QuestionsTags
ALTER TABLE QuestionsTags ADD CONSTRAINT fk_question_tags_question FOREIGN KEY (QuestionId) REFERENCES Questions(Id);
ALTER TABLE QuestionsTags ADD CONSTRAINT fk_question_tags_tag FOREIGN KEY (TagId) REFERENCES Tags(Id);

