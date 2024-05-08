-- Criação da tabela Tags
CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    text TEXT,
    creation_date TIMESTAMP
);

-- Criação da tabela Questions
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    text TEXT,
    deleted BOOLEAN,
    creation_date TIMESTAMP
);

-- Criação da tabela QuestionsTags
CREATE TABLE question_tags (
    id SERIAL PRIMARY KEY,
    question_id INT,
    tag_id INT,
    creation_date TIMESTAMP
);

-- Criação da tabela Answers
CREATE TABLE answers (
    id SERIAL PRIMARY KEY,
    question_id INT,
    text TEXT,
    correct BOOLEAN
);

-- Criação da tabela QuestionAnswers
CREATE TABLE question_answers (
    id SERIAL PRIMARY KEY,
    question_id INT,
    answer_id INT
);

-- Adição da chave estrangeira na tabela Answers
ALTER TABLE answers ADD CONSTRAINT fk_question_answers FOREIGN KEY (question_id) REFERENCES questions(id);

-- Adição das chaves estrangeira na tabela QuestionsTags
ALTER TABLE question_tags ADD CONSTRAINT fk_question_tags_question FOREIGN KEY (question_id) REFERENCES questions(id);
ALTER TABLE question_tags ADD CONSTRAINT fk_question_tags_tag FOREIGN KEY (tag_id) REFERENCES tags(id);

-- Adição das chaves estrangeira na tabela QuestionsAnswers
ALTER TABLE question_answers ADD CONSTRAINT fk_question_answers_question FOREIGN KEY (question_id) REFERENCES questions(id);
ALTER TABLE question_answers ADD CONSTRAINT fk_question_answers_answer FOREIGN KEY (answer_id) REFERENCES answers(id);
