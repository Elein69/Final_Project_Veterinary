import dotenv from 'dotenv';
dotenv.config();

export const PORT = process.env.PORT || 3000;

export const DB_CONFIG = {
  host: process.env.DB_HOST || 'localhost',
  port: Number(process.env.DB_PORT) || 5432,
  user: process.env.DB_USER || 'postgres',
  password: process.env.DB_PASSWORD || '',
  database: process.env.DB_NAME || '',
};

export const JWT_SECRET = process.env.JWT_SECRET || 'secret';
