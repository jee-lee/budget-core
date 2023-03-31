-- migrate:up
CREATE function update_updated_at() RETURNS trigger AS $update_updated_at$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$update_updated_at$ LANGUAGE plpgsql;

CREATE TABLE cycle_types (
                             id SERIAL PRIMARY KEY,
                             name TEXT
);

CREATE UNIQUE INDEX idx_cycle_types_name ON cycle_types (name);

CREATE TABLE account_types (
                               id SERIAL PRIMARY KEY,
                               name TEXT
);

CREATE UNIQUE INDEX idx_account_types_name ON account_types (name);

CREATE TABLE transaction_types (
                                   id SERIAL PRIMARY KEY,
                                   name text
);

CREATE UNIQUE INDEX idx_transaction_types_name ON transaction_types (name);

CREATE TABLE accounts (
                          id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                          user_id uuid NOT NULL,
                          account_type_id int NOT NULL REFERENCES account_types,
                          name TEXT,
                          is_joint BOOLEAN,
                          account_number INTEGER,
                          routing_number INTEGER,
                          initial_balance NUMERIC(10, 2),
                          credit_limit INTEGER,
                          statement_end_date DATE,
                          payment_due_date DATE,
                          created_at timestamptz NOT NULL DEFAULT NOW(),
                          updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_accounts_user_id ON accounts(user_id);

CREATE TRIGGER update_accounts_updated_at
    BEFORE UPDATE
    ON accounts
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();


CREATE TABLE joint_accounts (
                                id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                                account_id uuid NOT NULL REFERENCES accounts(id),
                                owner_id uuid NOT NULL,
                                joint_user_id uuid NOT NULL
);

CREATE TABLE categories (
                            id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                            user_id uuid NOT NULL,
                            name TEXT NOT NULL,
                            parent_category_id uuid REFERENCES categories(id),
                            maximum NUMERIC(10,2),
                            cycle_type_id int REFERENCES cycle_types(id),
                            rollover BOOLEAN,
                            joint_user_id uuid,
                            created_at timestamptz NOT NULL DEFAULT NOW(),
                            updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_categories_user_id ON categories(user_id);
CREATE INDEX idx_categories_joint_user_id ON categories(joint_user_id);
CREATE INDEX idx_categories_parent_category_id ON categories(parent_category_id);

CREATE TRIGGER update_categories_updated_at
    BEFORE UPDATE
    ON categories
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TABLE transactions (
                              transaction_date timestamptz NOT NULL DEFAULT NOW(),
                              id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                              user_id uuid NOT NULL,
                              budget_id uuid REFERENCES categories(id),
                              description TEXT,
                              transaction_type_id int NOT NULL REFERENCES transaction_types(id),
                              account_id uuid REFERENCES accounts(id),
                              amount NUMERIC(10,2),
                              currency TEXT NOT NULL DEFAULT 'USD',
                              comments TEXT,
                              created_at timestamptz NOT NULL DEFAULT NOW(),
                              updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_transactions_budget_id ON transactions(budget_id);
CREATE INDEX idx_transactions_account_id ON transactions(account_id);
CREATE INDEX idx_transactions_user_id ON transactions(user_id);

CREATE TRIGGER update_transaction_updated_at
    BEFORE UPDATE
    ON transactions
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

-- migrate:down

DROP TABLE transactions;
DROP TABLE categories;
DROP TABLE accounts;
DROP TABLE transaction_types;
DROP TABLE account_types;
DROP TABLE cycle_types;

DROP FUNCTION update_updated_at();