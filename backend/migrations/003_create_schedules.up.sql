CREATE TABLE IF NOT EXISTS schedule_slots (
    id          SERIAL PRIMARY KEY,
    goal_id     INTEGER REFERENCES goals(id) ON DELETE SET NULL,
    day         VARCHAR(20) NOT NULL,
    start_time  TIME NOT NULL,
    end_time    TIME NOT NULL,
    note        TEXT,
    created_at  TIMESTAMPTZ DEFAULT NOW()
);
