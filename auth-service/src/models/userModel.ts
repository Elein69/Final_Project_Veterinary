import { db } from '../db';

export const createUserTable = async () => {
  await db.query(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`);

  await db.query(`
    CREATE TABLE IF NOT EXISTS users (
      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      name VARCHAR(100) NOT NULL,
      email VARCHAR(100) NOT NULL UNIQUE,
      password TEXT NOT NULL,
      role VARCHAR(20) DEFAULT 'cliente',
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )
  `);
};