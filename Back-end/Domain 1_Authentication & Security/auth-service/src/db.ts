import { Pool } from 'pg';
import { DB_CONFIG, IS_CI } from './config/env';

export const db = IS_CI
  ? { connect: async () => console.log('⚠️ CI: DB connection mocked') }
  : new Pool(DB_CONFIG);

({
  host: process.env.DB_HOST || 'localhost', // fallback para CI
  port: Number(process.env.DB_PORT),
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
});